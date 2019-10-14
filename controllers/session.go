package controllers

import (
	"github.com/astaxie/beego"
	"ihome/commons"
	"ihome/models"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) GetSessionData() {
	result := new(commons.Result)
	user := models.User{}

	name := this.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		result.Code = models.RECODE_OK
		result.Msg = models.RecodeText(models.RECODE_OK)
		result.Data = user
	} else {
		result.Code = models.RECODE_DBERR
		result.Msg = models.RecodeText(models.RECODE_DBERR)
	}

	this.Data["json"] = result
	this.ServeJSON()
}

func (this *SessionController) DeleteSessionData() {
	result := new(commons.Result)
	this.DelSession("name")
	result.Code = models.RECODE_OK
	result.Msg = models.RecodeText(models.RECODE_OK)
	this.Data["json"] = result
	this.ServeJSON()
}
