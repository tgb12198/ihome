package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"ihome/commons"
	"ihome/models"
	"time"
)

type AreaController struct {
	beego.Controller
}

const AREASKEY string = "areas"

func (this *AreaController) GetArea() {

	redisConn, err := cache.NewCache("redis", `{"key":"ihome","conn":":6379","dbNum":"0"}`)
	if err != nil {
		beego.Info("cache_conn连接错误：", err)
	}
	/*cacheErr := redisConn.Put("aaa", "bbb", time.Second*3600)
	if cacheErr != nil {
		beego.Info("set数据异常：", cacheErr)
	}
	beego.Info("获取key值：", redisConn.Get("aaa"))

	beego.Info("connection success")*/
	//1、从redis缓存拿数据
	//2、从数据库拿数据
	//3、打包成json数据返回
	result := new(commons.Result)
	var areas []models.Area

	data := redisConn.Get(AREASKEY)
	//beego.Info("data111:", data)

	if data != nil {
		/*var s = fmt.Sprintf("%s", data)
		json.Unmarshal([]byte(s), &areas)
		fmt.Printf("%v", areas)*/
		//beego.Info("redis数据")
		var s = fmt.Sprintf("%s", data)
		json.Unmarshal([]byte(s), &areas)
		result.Code = models.RECODE_OK
		result.Msg = models.RecodeText(models.RECODE_OK)
		result.Data = areas

	} else {
		beego.Info("数据库数据")
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
		//beego.Info("areas:", areas)
		jsonArea, err := json.Marshal(areas)
		//beego.Info("json:", areas)
		if err != nil {
			beego.Info("转换失败")
		}
		cacheErr := redisConn.Put(AREASKEY, jsonArea, time.Second*3600)
		if cacheErr != nil {
			beego.Info("set数据异常：", cacheErr)
		}
	}

	this.Data["json"] = result
	this.ServeJSON()
}
