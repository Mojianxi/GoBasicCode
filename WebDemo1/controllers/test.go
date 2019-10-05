package controllers
import (
	"github.com/astaxie/beego"
)
type TestController struct{
	beego.Controller
}
type User struct{
    UserName string
    Password string
}
func (c *TestController) Test() {
	c.Ctx.WriteString("this TestController")
	// id :=c.GetString("id")
	// c.Ctx.WriteString(id)
}
func (c *TestController) TestForm(){
	c.Ctx.WriteString(`<html><form action="http://localhost:8080/test/form" method="post"> 
        	<input type="text" name="UserName"/>
        	<input type="password" name="Password"/>
        	<input type="submit" value="提交"/>
        </form></html>`)
}
func (c *TestController) Post(){
    u :=User{}
    if err:=c.ParseForm(&u); err!=nil{
        //process err
    }
    c.Ctx.WriteString("Username"+u.UserName)
}
