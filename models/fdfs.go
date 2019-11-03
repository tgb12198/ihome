package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tedcy/fdfs_client"
	"mime/multipart"
	"path"
)

//上传文件
func Upload(file multipart.File, h *multipart.FileHeader) (url string, err error) {
	client, err := fdfs_client.NewClientWithConfig("conf/client.conf") //D:/Program Files/GO/goWorkSpace/bin/src/ihome/
	defer client.Destory()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	suffix := path.Ext(h.Filename)
	buffer, e := getFileBuffer(file, h)
	if e != nil {
		fmt.Println(e.Error())
	}
	fileId, err := client.UploadByBuffer(buffer, suffix[1:])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return fileId, err
}

func getFileBuffer(file multipart.File, header *multipart.FileHeader) (buffer []byte, err error) {
	buffer = make([]byte, header.Size)
	_, e := file.Read(buffer)
	if e != nil {
		beego.Error("获取文件数据失败", e)
		return
	}
	return buffer, e
}
