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
	beego.Info("connection success")
	//1、从session拿数据
	//2、从数据库拿数据
	//3、打包成json数据返回
	result := new(commons.Result)
	var areas []models.Area
	o := orm.NewOrm()
	qs := o.QueryTable("area")
	num, err := qs.All(&areas)
	if err != nil {
		result.Code = models.RECODE_DBERR
		result.Msg = models.RecodeText(models.RECODE_DBERR)
		return
	}
	if num == 0 {
		result.Code = models.RECODE_NODATA
		result.Msg = models.RecodeText(models.RECODE_NODATA)
		return
	}
	result.Code = models.RECODE_OK
	result.Msg = models.RecodeText(models.RECODE_OK)
	result.Data = areas

	this.Data["json"] = result
	this.ServeJSON()
}
