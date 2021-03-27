package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ucanme/fastgo/conf"
	"github.com/ucanme/fastgo/controller/response"
	"path"
	"time"
)


type uploadResp struct {
	Url string `json:"url"`
}
//文件上传
func DeviceUpload(c *gin.Context) {
	var err error
	f, err := c.FormFile("file")
	if err != nil {
		response.Fail(c,400,"请选择上传文件")
		return
	}

	name :=  time.Now().Format("20060102150405")+f.Filename
	filename := path.Join(conf.Config.UploadDir.Dir,name)
	err = c.SaveUploadedFile(f, filename)
	if err != nil {
		fmt.Println(err)
		response.Fail(c,400,"文件保存失败")
		return
	}

	uploadResp := uploadResp{Url: conf.Config.UploadDir.Host+name}
	response.Success(c, uploadResp)
}



