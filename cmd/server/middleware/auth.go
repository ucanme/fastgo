package middleware

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"github.com/ucanme/fastgo/conf"
	"io/ioutil"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	clog "github.com/ucanme/fastgo/library/log"
)

const (
	HeaderKeyApiKey    = "apiKey"
	HeaderKeySign      = "sign"
	HeaderKeyTimestamp = "timestamp"
)

func Auth() gin.HandlerFunc {
	secret := conf.Config.Auth.Secret
	accounts := conf.Config.Auth.Accounts
	skips := conf.Config.Auth.Skips
	var skip map[string]struct{}
	if length := len(skips); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range skips {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		if s, ok := c.GetQuery("secret"); ok && s != "" && s == secret {
			c.Next()
			return
		}

		if _, ok := skip[c.Request.URL.Path]; ok {
			c.Next()
			return
		}

		apiKey := c.GetHeader(HeaderKeyApiKey)
		sign := c.GetHeader(HeaderKeySign)
		ts := c.GetHeader(HeaderKeyTimestamp)

		apiSecret, ok := accounts[apiKey]
		fmt.Println("sign", apiKey, sign, ts, apiSecret)
		if apiKey == "" || sign == "" || !ok {
			renderAuthFail(c)
			c.Abort()
			return
		}

		var (
			reqBody []byte
			err     error
		)

		switch c.Request.Method {
		case "POST", "PUT":
			reqBody, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				renderAuthFail(c)
				c.Abort()
				return
			}
		default:
			values := c.Request.URL.Query()
			keys := make([]string, 0, len(values))
			for key := range values {
				keys = append(keys, key)
			}

			sort.Strings(keys)
			buf := bytes.Buffer{}
			for i, key := range keys {
				if val, ok := c.GetQuery(key); ok {
					buf.WriteString(key + "=" + val)
					if i != len(keys)-1 {
						buf.WriteString("&")
					}
				}
			}

			reqBody = buf.Bytes()
		}

		_ts, err := strconv.ParseInt(ts, 10, 64)
		if err != nil {
			renderAuthFail(c)
			c.Abort()
			return
		}
		tm := time.Unix(_ts, 0)
		if time.Now().Sub(tm).Seconds() > 60 {
			renderAuthFail(c)
			c.Abort()
			return
		}

		signing := bytes.Buffer{}
		signing.Write(reqBody)
		signing.WriteString(ts)
		signing.WriteString(apiSecret)

		// 计算签名
		genSign := fmt.Sprintf("%x", sha1.Sum(signing.Bytes()))
		fmt.Println("sign", genSign)
		if genSign != sign {
			renderAuthFail(c)
			c.Abort()
			return
		}

		// restore Body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		c.Next()
	}
}

func renderAuthFail(c *gin.Context) {
	jsonObj := map[string]interface{}{
		"request_id": clog.GetRequestID(c),
		"error_no":   400,
		"error_msg":  "auth failed",
	}
	c.JSON(200, jsonObj)
}
