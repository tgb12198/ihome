package main

import (
	"fmt"
	_ "ihome/routers"
	"github.com/astaxie/beego"
	_"ihome/models"
	"net/http"
	"strings"
	"github.com/astaxie/beego/context"
)

func main() {
	ignoreStaticPath()
	beego.Run()
}

func ignoreStaticPath() {

	//透明static
	beego.SetStaticPath("group1/M00/","fdfs/storage_data/data/")

	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	fmt.Println("哈哈："+orpath)
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html"+ctx.Request.URL.Path)
}

