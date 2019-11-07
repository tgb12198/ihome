package routers

import (
	"github.com/astaxie/beego"
	"ihome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetArea")
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{}, "get:GetHouseIndex")
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetSessionData;delete:DeleteSessionData")
	beego.Router("/api/v1.0/register", &controllers.UserController{}, "post:Register")
	beego.Router("/api/v1.0/sessions", &controllers.UserController{}, "post:Login")
	beego.Router("/api/v1.0/upload", &controllers.UserController{}, "get:GetUser")
	beego.Router("/api/v1.0/user", &controllers.UserController{}, "get:GetUser")
	beego.Router("/api/v1.0/upload/avatar",&controllers.UploadController{},"post:UploadFile")
	beego.Router("/api/v1.0/user/name",&controllers.UserController{},"put:ModifyUserName")
	beego.Router("api/v1.0/user/auth",&controllers.UserController{},"get:GetUser;post:SaveAuthor")
}
