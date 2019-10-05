package controllers
import (
	"github.com/astaxie/beego"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/orm"
	// "fmt"
	"WebDemo1/models"
)
// //UseInfo操作的表是user_info 
// type Test struct{
// 	Id int64
// 	TName string
// 	TTime string
// }
type TestModelController struct{
	beego.Controller
}
func (c *TestModelController) Get(){
	
	// //采用QueryBuilder方式进行读取
	// var users []Test
	// qb, _:=orm.NewQueryBuilder("mysql")
	// qb.Select("t_name").From("Test").Where("id=?").Limit(1)
	// sql :=qb.String()
	// o.Raw(sql,2).QueryRows(&users)
	// c.Ctx.WriteString(fmt.Sprintf("user intfo:%v",users))
	user :=models.Test{Id:13,TName:"cgetm1",TTime:"2019-10-05"}
	models.AddUser(&user)
	c.Ctx.WriteString("call model success!")
}
