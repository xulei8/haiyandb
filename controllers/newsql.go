package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type NewSQLSave struct {
	beego.Controller
}

type NewSQLGet struct {
	beego.Controller
}

func (this *NewSQLSave) Post() {
	o := orm.NewOrm()
	sql := ""

	this.Ctx.WriteString("aaaa")
	formData := this.Input()
	i := 0
	str := ""
	title := ""
	for k, v := range formData {
		if i == 0 {
			str += "\"" + k + "\"," + "\"" + v[0] + "\""
			title = v[0]
		} else {
			str += ",\"" + k + "\"," + "\"" + v[0] + "\""
		}
		i++
	}
	tt := time.Now()

	t := tt.String()

	sql = "insert into app(title ,type,createorid,ownerid ,data, createtime) values('" + title + "','demo',1,1,COLUMN_CREATE(" + str + ") ,'" + t + "');"

	o.Raw(sql).Exec()
	fmt.Print(sql)

}

type NewsqlData struct {
	Id    int
	Title string

	Createtime time.Time `orm:"auto_now;type(datetime); null"`
}

func (this *NewSQLGet) Post() {
	o := orm.NewOrm()
	var dd []NewsqlData
	sql := "  select id  , title ,createtime  from app where type = 'demo' order by id desc "

	nums, err := o.Raw(sql).QueryRows(&dd)

	if err == nil && nums > 0 {

		fmt.Print(dd)
	}

	this.Ctx.WriteString(`{"rows":[  {"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":8}],"total":5}`)

	//{"rows":[{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":0},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":1},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":2},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":3},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":4},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":5},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":6},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":7},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":8}],"total":5}

	fmt.Print("hello")

}
