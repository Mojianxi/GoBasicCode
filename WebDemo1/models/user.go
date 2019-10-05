package models
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)
var (
	o orm.Ormer
)
//UseInfo操作的表是user_info 
type Test struct{
	Id int64
	TName string
	TTime string
}
func init(){
	orm.Debug=true//是否开启调试模式,开启输出sql
	orm.RegisterDataBase("default","mysql","root:123456@tcp(localhost:3306)/cost?charset=utf8")
	orm.RegisterModel(new(Test))
	o=orm.NewOrm()

}
func AddUser(test *Test)(int64 ,error){
	o :=orm.NewOrm()
	id ,err:=o.Insert(test)
	return id,err
}
func ReadUserInfo(users *[]Test){
	qb, _:=orm.NewQueryBuilder("mysql")
	qb.Select("*").From("test")
	sql :=qb.String()
	o.Raw(sql).QueryRows(users)
}