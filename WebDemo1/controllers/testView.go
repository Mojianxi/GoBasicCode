package controllers
import (
	"github.com/astaxie/beego"
	"WebDemo1/models"
)
type TestViewController struct{
	beego.Controller
}
func (c *TestViewController) Get(){
	c.Data["Title"]="hello golang"
	c.Data["IsDisplay"]=true
	var users []models.Test
	models.ReadUserInfo(&users)
	c.Data["users"]=users
	c.Data["Content"]="golang view example"
	c.TplName="test_view.tpl"
}