package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ihome/commons"
	"ihome/models"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) GetArea() {
	result := new(commons.Result)
	var areas []models.Area
	o := orm.NewOrm()
	qs := o.QueryTable("area")
	num, err := qs.All(&areas)
	if err != nil {
		result.ErrNo = 4001
		result.Msg = "查询数据错误"
		return
	}
	if num == 0 {
		result.ErrNo = 4002
		result.Msg = "无查询到数据"
		return
	}
	result.ErrNo = 0
	result.Msg = "OK"
	result.Data = areas

	this.Data["json"] = result
	this.ServeJSON()
}
