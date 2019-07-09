package routers

import (
	"MIX_GRPC/src/begoo-api/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/regist", &controllers.UserController{}, "POST:Regist")
    beego.Router("/getuinfobyemail", &controllers.UserController{}, "*:Show")
    beego.Router("/getalluinfo", &controllers.UserController{}, "*:GetAll")

    beego.Router("/addbuyitem", &controllers.SaleController{}, "POST:Addbuyitem")
	beego.Router("/getitemsbyemail", &controllers.SaleController{}, "*:GetBuyitemsByEmail")
	beego.Router("/getallitems", &controllers.SaleController{}, "*:GetAllBuyitems")

}
