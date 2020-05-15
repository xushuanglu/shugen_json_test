package routers

import (
	"shugen_json_test/controllers"

	"github.com/astaxie/beego"
)



func init() {

	beego.Router("/iroot/json", &controllers.LoginController{}, "*:GetJson")
}
