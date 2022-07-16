package middleware

import (
	"bytes"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/internal/session"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"encoding/json"
	clog "github.com/ucanme/fastgo/library/log"
)


var skipUrls = map[string]int{
	"/v1/login" :1,
	"/v1/report":1,
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoginAuth(c *gin.Context)  {
	if skipUrls[c.Request.RequestURI] == 0 {
		cookie, err := c.Cookie("login_session")
		sid, _ := url.QueryUnescape(cookie)
		if sid == "" {
			response.Fail(c, consts.ACCOUTN_NOT_LOGIN, "请登陆")
			c.Abort()
			return
		}
		ck, err := session.Manager.Read(sid)
		clog.LogNotice(map[string]interface{}{"cookie_notice":ck,"err":err})
		if err != nil {
			err = session.Manager.SessionDestroy(c)
			response.Fail(c, consts.ACCOUTN_NOT_LOGIN, "请登陆")
			c.Abort()
			return
		}
	}
}

func Access() gin.HandlerFunc {
	return func(c *gin.Context) {
		LoginAuth(c)
		blw := &bodyLogWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = blw
		start := time.Now()
		fields := map[string]interface{}{
			"log_type": "access",
			"method":   c.Request.Method,
			"latency":  time.Since(start) / time.Microsecond,
			"path":     c.Request.URL.String(),
		}

		//origin := c.Request.Header.Get("Origin")
		//if origin != ""  {

		//	c.Header("Access-Control-Allow-Origin", "http://localhost:18089")
		////	c.Header("Access-Control-Allow-Origin", "*")
		//	c.Header("Access-Control-Max-Age", "3600")
		//	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		//	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		//	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		//	c.Header("Access-Control-Allow-Credentials", "true")
		//	c.Set("content-type", "application/json")
		////}

		if c.Request.Method == "OPTIONS"{
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		//s := session.Manager.SessionStart(c,"")
		//if s.Value == ""{
		//	response.Response(c,400,"没有登陆",nil)
		//	return
		//}
		//sVal := v1.SessionData{}
		//err := json.Unmarshal([]byte(s.Value),&sVal)
		//if err != nil{
		//	response.Response(c,400,"没有登陆",nil)
		//	return
		//}
		//c.Set("session",sVal)

		if c.Request.Method == "POST" {
			reqBody, err := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
			if err != nil {
				fields["err_info"] = "dump_request_error" + err.Error()
			} else {
				fields["request"] = string(reqBody)
			}
		}

		clog.LogNotice(map[string]interface{}{"cookie_start":1})
		for _,v := range c.Request.Cookies(){
			clog.LogNotice(map[string]interface{}{"name":v.Name,"value":v.Value,"domain":v.Domain})
		}
		clog.LogNotice(map[string]interface{}{"cookie_end":1})

		cookie, err := c.Cookie("login_session")
		clog.LogNotice(map[string]interface{}{"cookie":cookie,"err":err})
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
