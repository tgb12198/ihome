package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ihome/commons"
	"ihome/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Register() {
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
			u.Name = resp["name"].(string)
			u.Password_hash = commons.EncryptData(resp["password"].(string))
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

func (this *UserController) Login() {
	var result commons.Result
	var resp = make(map[string]interface{})
	//1、得到用户信息
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	mobile := resp["mobile"].(string)
	password := commons.EncryptData(resp["password"].(string))
	//2、判断是否合法
	if mobile == "" || password == "" {
		result.Code = models.RECODE_NOTNULL
		result.Msg = models.RecodeText(models.RECODE_NOTNULL)
		return
	}

	//3、与数据库匹配判断帐号密码是否正确
	o := orm.NewOrm()
	var user models.User
	if err := o.QueryTable("user").Filter("mobile", mobile).One(&user); err != nil {
		result.Code = models.RECODE_DBERR
		result.Msg = models.RecodeText(models.RECODE_DBERR)
		return
	}

	if user.Mobile == "" {
		result.Code = models.RECODE_USERERR
		result.Msg = models.RecodeText(models.RECODE_USERERR)
		return
	}
	if user.Password_hash != password {
		result.Code = models.RECODE_PWDERR
		result.Msg = models.RecodeText(models.RECODE_PWDERR)
		return
	}

	//4、设置session
	this.SetSession("name", user.Name)
	this.SetSession("mobile", user.Mobile)
	this.SetSession("user_id", user.Id)
	this.SetSession("password", user.Password_hash)

	//5、返回登录信息
	result.Code = models.RECODE_OK
	result.Msg = models.RecodeText(models.RECODE_OK)
	this.Data["json"] = result
	this.ServeJSON()
}

//获取用户信息
func (this *UserController) GetUser() {
	var result commons.Result
	var user models.User
	userId := this.GetSession("user_id")

	//4.2更新user表头像字段
	o := orm.NewOrm()
	user.Id = userId.(int)
	dbErr := o.Read(&user)
	if dbErr != nil {
		beego.Error("获取数据库数据失败", dbErr)
		result.Code = models.RECODE_SESSIONERR
		result.Msg = models.RecodeText(models.RECODE_SESSIONERR)
	}
	result.Code = models.RECODE_OK
	result.Msg = models.RecodeText(models.RECODE_OK)
	result.Data = user
	this.Data["json"] = result
	this.ServeJSON()

}
