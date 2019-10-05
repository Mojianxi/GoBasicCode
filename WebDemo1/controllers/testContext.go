package controllers
import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/httplib"
	"strconv"
)
type TestContextController struct{
	beego.Controller
}
func (c *TestContextController) Get(){
	c.Ctx.WriteString(c.Ctx.Input.IP()+":"+strconv.Itoa(c.Ctx.Input.Port()))
}