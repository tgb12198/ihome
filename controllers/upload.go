package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ihome/commons"
	"ihome/models"
	//"path"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) UploadFile() {
	var result commons.Result
	//1、获取文件信息
	f, h, err := this.GetFile("avatar")
	if err != nil {
		beego.Info("获取文件失败", err.Error())
		result.Code = models.RECODE_REQERR
		result.Msg = models.RecodeText(models.RECODE_DBERR)
	}
	defer f.Close()
	//2、获取文件后辍名
	//suffix:=path.Ext(h.Filename)
	//3、上传文件到fastdfs并获取到文件url
	url, err := models.Upload(f, h)
	if err != nil {
		beego.Error("上传文件失败", err)
		result.Code = models.RECODE_REQERR
		result.Msg = models.RecodeText(models.RECODE_REQERR)
		return
	}
	//4、更新User表

	//4.1从session中取得userId
	userId:= this.GetSession("user_id")
	//4.2更新user表头像字段
	o := orm.NewOrm()
	var user models.User
	user.Id = userId.(int)
	if err := o.Read(&user); err != nil {
		result.Code = models.RECODE_DBERR
		result.Msg = models.RecodeText(models.RECODE_DBERR)
		return
	}
	user.Avatar_url = url
	num, err := o.Update(&user)
	if err != nil {
		beego.Error("更新数据失败", err)
		result.Code = models.RECODE_DBERR
		result.Msg = models.RecodeText(models.RECODE_DBERR)
		return
	}
	if num > 0 {
		result.Code = models.RECODE_OK
		result.Msg = models.RecodeText(models.RECODE_OK)
		result.Data = commons.BASEPATH + url
	}
	this.Data["json"] = result
	this.ServeJSON()

}
