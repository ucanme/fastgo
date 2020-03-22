package log

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/ucanme/fastgo/conf"
	"path"
	"path/filepath"
	"reflect"
	"time"
)

type logInfo struct {
	RequestId string                 `json:"request_id"`
	Level     string                 `json:"level"`
	Info      map[string]interface{} `json:"info"`
}

//var LogChan = make(chan (logInfo), 10000)

const REQUEST_ID = "request_id"

//设置请求request_id
func SetRequestID(c *gin.Context) string {
	requestId := c.GetHeader("request_id")
	if len(requestId) < 20 {
		requestId = uuid.NewV4().String()
	}

	c.Set("request_id", requestId)
	return requestId
}

//获取请求request_id
func GetRequestID(c *gin.Context) string {
	v, ok := c.Get(REQUEST_ID)
	if ok && reflect.TypeOf(v).Kind() == reflect.String {
		return v.(string)
	}
	requestId := uuid.NewV4().String()
	c.Set(REQUEST_ID, requestId)
	return requestId
}

var noticeLogger *logrus.Logger
var warningLogger *logrus.Logger
var errorLogger *logrus.Logger

func Init() {
	// 日志文件

	conf.Config.Log.FilePath, _ = filepath.Abs(conf.Config.Log.FilePath)
	noticefileName := path.Join(conf.Config.Log.FilePath, conf.Config.Log.FileName+".log")
	warningfileName := path.Join(conf.Config.Log.FilePath, conf.Config.Log.FileName+".wf.log")
	errorfileName := path.Join(conf.Config.Log.FilePath, conf.Config.Log.FileName+".error.log")
	//
	//// 写入文件
	//noticeFile, err := os.OpenFile(noticefileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	//if err != nil {
	//	panic(err)
	//}
	//warningFile, err := os.OpenFile(warningfileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	//if err != nil {
	//	panic(err)
	//}
	//errorFile, err := os.OpenFile(errorfileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	//if err != nil {
	//	panic(err)
	//}

	//if err != nil {
	//	fmt.Println("err", err)
	//}
	// 实例化
	noticeLogger = logrus.New()
	noticeLogger.SetFormatter(&logrus.JSONFormatter{})

	warningLogger = logrus.New()
	warningLogger.SetFormatter(&logrus.JSONFormatter{})

	errorLogger = logrus.New()
	errorLogger.SetFormatter(&logrus.JSONFormatter{})

	//设置日志级别
	//设置输出

	// 设置 rotatelogs
	noticeWriter, err := rotatelogs.New(
		// 分割后的文件名称
		noticefileName+".%Y%m%d",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(noticefileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(1*time.Hour),
	)

	if err != nil {
		panic(err)
	}
	noticeLogger.SetOutput(noticeWriter)

	warningWriter, err := rotatelogs.New(
		// 分割后的文件名称
		warningfileName+".%Y%m%d",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(noticefileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(1*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	warningLogger.SetOutput(warningWriter)

	errorWriter, err := rotatelogs.New(
		// 分割后的文件名称
		errorfileName+".%Y%m%d",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(noticefileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(1*time.Hour),
	)
	if err != nil {
		panic(err)
	}

	errorLogger.SetOutput(errorWriter)
	LogNotice(map[string]interface{}{"hello": "world"})
}

func LogNotice(data map[string]interface{}) {
	fields := logrus.Fields{}
	for i, v := range data {
		fields[i] = v
	}
	noticeLogger.WithFields(fields).Info()
}

func LogError(data map[string]interface{}) {
	fields := logrus.Fields{}
	for i, v := range data {
		fields[i] = v
	}
	errorLogger.WithFields(fields).Warn()
}
func LogWarning(data map[string]interface{}) {
	fields := logrus.Fields{}
	for i, v := range data {
		fields[i] = v
	}
	warningLogger.WithFields(fields).Error()
}
