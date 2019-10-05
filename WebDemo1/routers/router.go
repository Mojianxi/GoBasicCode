package routers

import (
	"WebDemo1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test", &controllers.TestController{},"get,post:Test")
    beego.Router("/test/form", &controllers.TestController{},"get,post:TestForm")
    // beego.Router("/test/form", &controllers.TestController{},"get,post:TestForm")
    beego.Router("/test_model",&controllers.TestModelController{},"get:Get;post:Post")
    beego.Router("/test_view",&controllers.TestViewController{},"get:Get;post:Post")
	beego.Router("/test_httplib",&controllers.TestHttpLibController{},"get:Get;post:Post")
	beego.Router("/test_context",&controllers.TestContextController{},"get:Get;post:Post")
}
