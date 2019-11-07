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

//修改用户名
func (this *UserController) ModifyUserName() {
	var result commons.Result
	mapName := make(map[string]string)
	//1、获取修改的信息
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &mapName); err != nil {
		result.Code = models.RECODE_NODATA
		result.Msg = models.RecodeText(models.RECODE_NODATA)
		return
	}
	name := mapName["name"]
	//2、根据session中的userid获取用户信息
	userId := this.GetSession("user_id")
	user := models.User{Id: userId.(int)}
	o := orm.NewOrm()
	if err := o.Read(&user); err != nil {
		result.Code = models.RECODE_SESSIONERR
		result.Msg = models.RecodeText(models.RECODE_SESSIONERR)
		return
	}
	//3、修改用户信息
	user.Name = name
	if _, err := o.Update(&user); err != nil {
		result.Code = models.RECODE_DBERR
		result.Msg = models.RecodeText(models.RECODE_DBERR)
		return
	}
	//4、更新session
	this.SetSession("name", name)
	//5、返回前端页面数据
	result.Code = models.RECODE_OK
	result.Msg = models.RecodeText(models.RECODE_OK)
	result.Data = mapName

	this.Data["json"] = result
	this.ServeJSON()
}

func (this *UserController) SaveAuthor() {
	var result commons.Result
	authorMap := make(map[string]string)
	//1、获取修改的实名认证信息
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &authorMap); err != nil {
		result.Code = models.RECODE_NODATA
		result.Msg = models.RecodeText(models.RECODE_NODATA)
		return
	}
	//2、根据session中的userid获取用户信息
	userId := this.GetSession("user_id")
	user := models.User{Id: userId.(int)}
	o := orm.NewOrm()
	if err := o.Read(&user); err != nil {
		result.Code = models.RECODE_DBERR
		result.Msg = models.RecodeText(models.RECODE_DBERR)
		return
	}
	//3、保存实名认证信息
	user.Real_name = authorMap["real_name"]
	user.Id_card = authorMap["id_card"]
	if _, err := o.Update(&user); err != nil {
		result.Code = models.RECODE_DBERR
		result.Msg = models.RecodeText(models.RECODE_DBERR)
		return
	}
	//4、返回前端数据
	result.Code = models.RECODE_OK
	result.Msg = models.RecodeText(models.RECODE_OK)
	result.Data = authorMap

	this.Data["json"] = result
	this.ServeJSON()
}
