package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ihome/commons"
	"ihome/models"
)

type HousesController struct {
	beego.Controller
}

func (this *HousesController) GetHouses() {
	var result commons.Result
	userId := this.GetSession("user_id")
	var houses []models.House

	o := orm.NewOrm()
	qs := o.QueryTable("house")
	num, err := qs.Filter("user__id", userId.(int)).All(&houses)
	if err != nil {
		result.Code = models.RECODE_DBERR
		result.Code = models.RecodeText(models.RECODE_DBERR)
		return
	}
	if num == 0 {
		result.Code = models.RECODE_NODATA
		result.Msg = models.RecodeText(models.RECODE_NODATA)
		return
	}

	result.Code = models.RECODE_OK
	result.Msg = models.RecodeText(models.RECODE_OK)
	result.Data = houses
	this.Data["json"] = result
	this.ServeJSON()
}

//发布房源
func (this *HousesController)PublishHouse()  {

}
