package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type NewSQLFiled struct {
	Label string
	Name  string
	List  bool
	Add   bool
}

type NewSQLModel struct {
	ID          int
	AppName     string
	AppKey      string
	AllowAdd    bool
	AllowDelete bool
	Allowupdate bool

	TitleFiled string

	Fileds []NewSQLFiled
}

var mNewSQLModel = map[string]NewSQLModel{}

type NewSQLSave struct {
	beego.Controller
}

type NewSQLDelete struct {
	beego.Controller
}

type NewSQLUpdate struct {
	beego.Controller
}

type NewSQLGet struct {
	beego.Controller
}

func init() {
	var f1s []NewSQLFiled
	newm := NewSQLModel{}
	newm.AllowAdd = true
	newm.AllowDelete = true
	newm.Allowupdate = true
	newm.AppKey = "app"
	newm.AppName = "用户"
	newm.TitleFiled = "firstname"

	nf := NewSQLFiled{}
	nf.Add = true
	nf.List = true
	nf.Label = "姓名"
	nf.Name = "firstname"

	nf1 := NewSQLFiled{}
	nf1.Add = true
	nf1.List = true
	nf1.Label = "昵称"
	nf1.Name = "lastname"

	f1s = append(f1s, nf, nf1)
	newm.Fileds = f1s
	mNewSQLModel["app"] = newm

}
func (this *NewSQLSave) Post() {
	o := orm.NewOrm()
	sql := ""

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
	this.Ctx.WriteString(`{"success":true }`)
	fmt.Print(sql)

}

func (this *NewSQLDelete) Post() {
	o := orm.NewOrm()
	sql := ""
	id := this.GetString("id")
	sql = " update app set  deleted = 1 where id =  " + id + " limit 1 "

	o.Raw(sql).Exec()
	this.Ctx.WriteString(`{"success":true }`)
	fmt.Print(sql)

}

func (this *NewSQLUpdate) Post() {
	o := orm.NewOrm()
	sql := ""
	id := this.GetString("id")

	formData := this.Input()
	i := 0
	str := ""

	for k, v := range formData {
		if i == 0 {
			str += "\"" + k + "\"," + "\"" + v[0] + "\""

		} else {
			str += ",\"" + k + "\"," + "\"" + v[0] + "\""
		}
		i++
	}

	sql = " update app set data =  COLUMN_CREATE(" + str + ")  where id = " + id + "  limit 1 ;"

	o.Raw(sql).Exec()

	this.Ctx.WriteString(`{"success":true }`)
	fmt.Print(sql)

}

type NewsqlData struct {
	Id    int
	Title string

	Createtime time.Time `orm:"auto_now;type(datetime); null"`
	Data       string
}

func (this *NewSQLGet) Post() {
	o := orm.NewOrm()
	var dd []NewsqlData

	rows, _ := this.GetInt("rows")
	page, _ := this.GetInt("page")

	if page > 10000 || page < 1 {
		page = 1
	}

	if rows > 500 || rows < 1 {
		rows = 10
	}
	page--
	skip := rows * page
	sql := "  select id    from app where type = 'demo' and deleted = 0    "

	allnum, _ := o.Raw(sql).QueryRows(&dd)
	skips := fmt.Sprintf("%d", skip)
	rowss := fmt.Sprintf("%d", rows)
	sql = "  select id  , title ,createtime, column_json(data) as data  from app where type = 'demo' and deleted = 0   order by id desc limit " + skips + " ," + rowss

	nums, err := o.Raw(sql).QueryRows(&dd)

	if err == nil && nums > 0 {

		//	fmt.Print(dd)
	}
	i := 0
	strj := ""
	for _, v := range dd {
		s := v.Data
		dataid := fmt.Sprintf("%d", v.Id)
		s = s[0:len(s)-1] + ",\"id\":" + dataid + " }"
		if i == 0 {
			strj += s
		} else {
			strj += "," + s
		}
		i++
	}

	nums_str := fmt.Sprintf("%d", allnum)
	this.Ctx.WriteString(`{"rows":[` + strj + `],"total":` + nums_str + `}`)

	//{"rows":[{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":0},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":1},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":2},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":3},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":4},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":5},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":6},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":7},{"email":"ss@qq.com","phone":"2222","lastname":"dddds ","firstname":"dddd  ddddddddddd","id":8}],"total":5}

	fmt.Print("hello")

}
