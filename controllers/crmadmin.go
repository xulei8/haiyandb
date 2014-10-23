package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	m "github.com/xulei8/haiyandb/models"
	"github.com/xulei8/haiyandb/utils"
	"strconv"
	"strings"
)

var o orm.Ormer

type baseRouter struct {
	beego.Controller

	isLogin bool
}

type CCT struct {
	ID int
	NN string
}
type CrmAdmin struct {
	//beego.Controller
	object m.DqContact
	baseRouter
}

func (this *CCT) Test() (link string) {

	return "testcct"
}
func (this *baseRouter) SetPaginator(per int, nums int64) *utils.Paginator {
	p := utils.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["Paginator"] = p
	this.Data["Tests"] = "aaaaaaaaaaaa"
	cct := CCT{ID: 2, NN: "ssss"}
	this.Data["Cc"] = cct
	return p
}

type CrmImport struct {
	beego.Controller
}

func (this *CrmAdmin) Object() interface{} {
	return &this.object
}

func (this *CrmAdmin) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	var articles []m.DqContact

	o := orm.NewOrm()

	qs := o.QueryTable("dq_contact").OrderBy("-id")
	//fmt.Println("%q",qs)

	if err := this.SetObjects(qs, &articles); err != nil {
		this.Data["Error"] = err
		beego.Error(err)
	}

	this.TplNames = "admin.tpl"
}

func (this *CrmAdmin) Post() {
	fa := this.GetStrings("fa")
	for _, v := range fa {

		id, _ := strconv.Atoi(v)

		o := orm.NewOrm()
		oneC := m.DqContact{Id: int64(id)}
		//	o.Read(&oneC)
		oneC.Called = 0
		o.Update(&oneC, "Called")

		//fmt.Printf("%d\n", id)
	}

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	var articles []m.DqContact

	o := orm.NewOrm()

	qs := o.QueryTable("dq_contact").OrderBy("-id")
	//fmt.Println("%q",qs)

	if err := this.SetObjects(qs, &articles); err != nil {
		this.Data["Error"] = err
		beego.Error(err)
	}

	this.TplNames = "admin.tpl"
}

func (this *CrmAdmin) SetObjects(qs orm.QuerySeter, objects interface{}) error {
	cnt, err := qs.Count()
	//fmt.Println("Count %d",cnt)
	if err != nil {
		return err
	}
	// create paginator
	pp := this.SetPaginator(10, cnt)

	if cnt, err := qs.Limit(pp.PerPageNums, pp.Offset()).All(objects); err != nil {
		return err
	} else {
		this.Data["Objects"] = objects
		this.Data["ObjectsCnt"] = cnt

	}
	return nil
}

func (this *CrmImport) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "import.tpl"
}

// 导入数据
func (this *CrmImport) Post() {
	nums := this.GetString("nums")
	line := strings.Split(nums, "\n")
	l := len(line)
	i := 0

	o = orm.NewOrm()
	for i = 0; i < l; i++ {

		a := strings.Trim(line[i], "\n")
		b := strings.Trim(a, " ")
		if len(b) > 3 {
			b = b + ",,,,,"
			arr := strings.Split(b, ",")
			u := new(m.DqContact)
			u.Mname = "10000"
			u.Uname = arr[1]
			u.Tel = arr[0]

			u.Tel2 = arr[2]

			if len(u.Tel) < 6 {
				continue
			} //电话号码长度

			if u.Uname == "" {
				u.Uname = "No Name"
			}
			//fmt.Println("line " + ":" + b)
			//m.AddDqContact(u)
			//o = orm.NewOrm()

			o.Insert(u)
		}
	}

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "import.tpl"
}
