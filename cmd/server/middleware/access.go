package middleware

import (
	"bytes"
	"fmt"
	"github.com/ucanme/fastgo/library/session"
	"io/ioutil"
	"net"
	"time"

	"github.com/gin-gonic/gin"

	"encoding/json"
	clog "github.com/ucanme/fastgo/library/log"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Access() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = blw
		start := time.Now()
		fields := map[string]interface{}{
			"log_type": "access",
			"method":   c.Request.Method,
			"latency":  time.Since(start) / time.Microsecond,
			"path":     c.Request.URL.String(),
		}



		if c.Request.Method == "POST" {
			reqBody, err := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

			fmt.Println("requsetData", reqBody)
			if err != nil {
				fields["err_info"] = "dump_request_error" + err.Error()
			} else {
				fields["request"] = string(reqBody)
			}
		}

		c.Next()
		var out struct {
			RequestID string `json:"request_id"`
		}
		out.RequestID = clog.GetRequestID(c)

		fields["request_id"] = out.RequestID
		fields["response"] = string(blw.body.Bytes())

		json.Unmarshal(blw.body.Bytes(), &out)

		if host, _, err := net.SplitHostPort(c.Request.RemoteAddr); err == nil {
			fields["remote_addr"] = host
		}
		if ua := c.Request.Header.Get("User-Agent"); ua != "" {
			fields["user_agent"] = ua
		}

		clog.LogNotice(fields)
	}
}
