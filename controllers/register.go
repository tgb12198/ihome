package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ihome/commons"
	"ihome/models"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Register() {
	var result commons.Result
	var resp = make(map[string]interface{})

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &resp); err == nil {
		o := orm.NewOrm()

		exist := o.QueryTable("user").Filter("mobile", resp["mobile"].(string)).Exist()
		if exist {
			result.Code = models.RECODE_ACCOUNTEXIST
			result.Msg = models.RecodeText(models.RECODE_ACCOUNTEXIST)
		} else {
			u := models.User{}
			u.Name = resp["mobile"].(string)
			u.Password_hash = resp["password"].(string)
			u.Mobile = resp["mobile"].(string)
			_, err := o.Insert(&u)
			if err == nil {
				result.Code = models.RECODE_OK
				result.Msg = models.RecodeText(models.RECODE_OK)
				this.SetSession("name", u.Name)
			}
		}
	} else {
		result.Code = models.RECODE_DBERR
		result.Msg = models.RecodeText(models.RECODE_DBERR)
	}
	this.Data["json"] = result
	this.ServeJSON()
}
