package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MySQLModConfig struct {
	Id        int
	Aname     string
	Appkeyf   string
	Appname   string
	Apptitlef string

	Allowadd  string
	Allowedit string

	Allowdelete string
}

type MySQLModFieldConfig struct {
	Fname      string
	Aname      string
	Ftype      string
	Frequire   string
	Dataoption string
	Addtip     string
}
type NewSQLFiled struct {
	Label      string
	Name       string
	List       bool
	ListWidth  string
	Add        bool
	AddRequire bool
	AddType    string
	HideInNew  bool
	HideInEdit bool
	HideInList bool
	JuiType    string
	Value      string

	FDataOption string
	LabelAddTip string
}

type NewSQLModel struct {
	ID          int
	AppName     string
	AppKey      string
	AllowAdd    bool
	AllowDelete bool
	Allowupdate bool

	TitleFiled string
	DataKey    string

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
	newm.DataKey = "demo"

	nf := NewSQLFiled{}
	nf.Add = true
	nf.List = true
	nf.Label = "姓名"
	nf.Name = "firstname"
	nf.AddRequire = true
	nf.JuiType = "easyui-textbox"

	nf1 := NewSQLFiled{}
	nf1.Add = true
	nf1.List = true
	nf1.Label = "昵称"
	nf1.Name = "lastname"
	nf1.AddRequire = true
	nf1.ListWidth = "60px"
	nf1.JuiType = "easyui-textbox"

	nf2 := NewSQLFiled{}
	nf2.Add = true
	nf2.List = true
	nf2.Label = "电话"
	nf2.Name = "phone"
	nf2.HideInNew = true
	nf2.JuiType = "easyui-textbox"

	nf3 := NewSQLFiled{}
	nf3.Add = true
	nf3.List = true
	nf3.Label = "邮件"
	nf3.Name = "email"
	nf3.AddType = "email"
	nf3.HideInEdit = true
	nf3.HideInList = false
	nf3.JuiType = "easyui-textbox"

	nf4 := NewSQLFiled{}
	nf4.Add = true
	nf4.List = true
	nf4.Label = "语言"
	nf4.Name = "lang"

	nf4.HideInEdit = false
	nf4.HideInList = false
	nf4.JuiType = "easyui-combobox"
	//nf4.FDataOption = "valueField: 'id', textField: 'text',  data: [{  \"id\":1,  \"text\":\"text1\"},{\"id\":2,\"text\":\"text2\"}] "
	nf4.FDataOption = "valueField: 'text', textField: 'text', url:'/data_select_values/?keys=app'   "

	nf5 := NewSQLFiled{}
	nf5.Add = true
	nf5.List = true
	nf5.Label = "生日"
	nf5.Name = "birth"
	nf5.JuiType = "easyui-datebox"
	f1s = append(f1s, nf, nf1, nf2, nf3, nf4, nf5)
	newm.Fileds = f1s

	newm2 := NewSQLModel{}
	newm2.AllowAdd = true
	newm2.AllowDelete = true
	newm2.Allowupdate = true
	newm2.AppKey = "contact"
	newm2.AppName = "客户"
	newm2.TitleFiled = "aname"
	newm2.DataKey = "custom"

	nff := NewSQLFiled{}
	nff.Add = true
	nff.List = true
	nff.Label = "客户"
	nff.Name = "aname"
	nff.AddRequire = true
	nff.JuiType = "easyui-textbox"
	nff.ListWidth = "60px"

	nff1 := NewSQLFiled{}
	nff1.Add = true
	nff1.List = true
	nff1.Label = "手机"
	nff1.Name = "tel"
	nff1.AddRequire = true

	nff1.JuiType = "easyui-textbox"

	nff4 := NewSQLFiled{}
	nff4.Add = true
	nff4.List = true
	nff4.Label = "电话"
	nff4.Name = "otherphone"
	nff4.JuiType = "easyui-textbox"

	nff2 := NewSQLFiled{}
	nff2.Add = true
	nff2.List = true
	nff2.Label = "地址"
	nff2.Name = "addr"
	nff2.JuiType = "easyui-textbox"

	nfa4 := NewSQLFiled{}
	nfa4.Add = true
	nfa4.List = true
	nfa4.Label = "客户类别"
	nfa4.Name = "acctype"

	nfa4.JuiType = "easyui-combobox"
	nfa4.FDataOption = "valueField: 'text', textField: 'text',  data: [{  \"id\":1,  \"text\":\"成交客户\"},{\"id\":2,\"text\":\"重点客户\"},{\"id\":3,\"text\":\"无效客户\"}] "
	//nf4.FDataOption = "valueField: 'text', textField: 'text', url:'/data_select_values/?keys=app'   "

	nff21 := NewSQLFiled{}
	nff21.Add = true
	nff21.List = true
	nff21.Label = "QQ"
	nff21.Name = "qq"
	nff21.JuiType = "easyui-textbox"

	nf311 := NewSQLFiled{}
	nf311.Add = true
	nf311.List = true
	nf311.Label = "邮件"
	nf311.Name = "email"
	nf311.AddType = "email"
	nf311.JuiType = "easyui-textbox"

	nff21n := NewSQLFiled{}
	nff21n.Add = true
	nff21n.List = true
	nff21n.Label = "备注"
	nff21n.Name = "note"
	nff21n.JuiType = "easyui-textbox"

	var f2s []NewSQLFiled
	f2s = append(f2s, nff, nff1, nff4, nfa4, nff2, nff21, nf311, nff21n)

	newm2.Fileds = f2s

	newm3 := NewSQLModel{}
	newm3.AllowAdd = true
	newm3.AllowDelete = true
	newm3.Allowupdate = true
	newm3.AppKey = "config"
	newm3.AppName = "配置"
	newm3.TitleFiled = "aname"
	newm3.DataKey = "conf"

	cf1 := NewSQLFiled{}
	cf1.Add = true
	cf1.List = true
	cf1.Label = "配置名称"
	cf1.Name = "aname"
	cf1.AddRequire = true
	cf1.JuiType = "easyui-textbox"
	cf1.ListWidth = "60px"

	cf2 := NewSQLFiled{}
	cf2.Add = true
	cf2.List = true
	cf2.Label = "变量"
	cf2.Name = "keys"
	cf2.AddRequire = true
	cf2.JuiType = "easyui-textbox"

	cf3 := NewSQLFiled{}
	cf3.Add = true
	cf3.List = true
	cf3.Label = "值"
	cf3.Name = "value"
	cf3.JuiType = "easyui-textbox"

	cf4 := NewSQLFiled{}
	cf4.Add = true
	cf4.List = true
	cf4.Label = "说明"
	cf4.Name = "note"
	cf4.JuiType = "easyui-textbox"

	var f3s []NewSQLFiled
	f3s = append(f3s, cf1, cf2, cf3, cf4)
	newm3.Fileds = f3s

	newmmod := NewSQLModel{}
	newmmod.AllowAdd = true
	newmmod.AllowDelete = true
	newmmod.Allowupdate = true
	newmmod.AppKey = "modsconfig"
	newmmod.AppName = "模块"
	newmmod.TitleFiled = "aname"
	newmmod.DataKey = "mods"

	modsf1 := NewSQLFiled{}
	modsf1.Add = true
	modsf1.List = true
	modsf1.Label = "模块名称"
	modsf1.Name = "aname"
	modsf1.AddRequire = true
	modsf1.JuiType = "easyui-textbox"
	modsf1.ListWidth = "60px"

	modsf2 := NewSQLFiled{}
	modsf2.Add = true
	modsf2.List = true
	modsf2.Label = "模块名"
	modsf2.Name = "appname"
	modsf2.AddRequire = true
	modsf2.JuiType = "easyui-textbox"

	modsf3 := NewSQLFiled{}
	modsf3.Add = true
	modsf3.List = true
	modsf3.Label = "主键名"
	modsf3.Name = "appkeyf"
	modsf3.AddRequire = true
	modsf3.JuiType = "easyui-textbox"

	modsf4 := NewSQLFiled{}
	modsf4.Add = true
	modsf4.List = true
	modsf4.Label = "主字段"
	modsf4.Name = "apptitlef"
	modsf4.AddRequire = true
	modsf4.JuiType = "easyui-textbox"

	modsf5 := NewSQLFiled{}
	modsf5.Add = true
	modsf5.List = true
	modsf5.Label = "新建工具"
	modsf5.Name = "allowadd"
	modsf5.JuiType = "easyui-combobox"
	modsf5.FDataOption = "valueField: 'id', textField: 'text',  data: [{  \"id\":1,  \"text\":\"允许新建\"},{\"id\":0,\"text\":\"禁止新建\"} ] "
	modsf5.AddRequire = true

	modsf6 := NewSQLFiled{}
	modsf6.Add = true
	modsf6.List = true
	modsf6.Label = "编辑工具"
	modsf6.Name = "allowedit"
	modsf6.JuiType = "easyui-combobox"
	modsf6.FDataOption = "valueField: 'id', textField: 'text',  data: [{  \"id\":1,  \"text\":\"允许编辑\"},{\"id\":0,\"text\":\"禁止编辑\"} ] "
	modsf6.AddRequire = true

	modsf7 := NewSQLFiled{}
	modsf7.Add = true
	modsf7.List = true
	modsf7.Label = "删除工具"
	modsf7.Name = "allowdelete"
	modsf7.JuiType = "easyui-combobox"
	modsf7.FDataOption = "valueField: 'id', textField: 'text',  data: [{  \"id\":1,  \"text\":\"允许删除\"},{\"id\":0,\"text\":\"禁止删除\"} ] "
	modsf7.AddRequire = true

	var f4s []NewSQLFiled
	f4s = append(f4s, modsf1, modsf2, modsf3, modsf4, modsf5, modsf6, modsf7)
	newmmod.Fileds = f4s

	newmf := NewSQLModel{}
	newmf.AllowAdd = true
	newmf.AllowDelete = true
	newmf.Allowupdate = true
	newmf.AppKey = "fsconfig"
	newmf.AppName = "字段"
	newmf.TitleFiled = "aname"
	newmf.DataKey = "fsconfigf"

	mf1 := NewSQLFiled{}
	mf1.Add = true
	mf1.List = true
	mf1.Label = "字段"
	mf1.Name = "aname"
	mf1.AddRequire = true
	mf1.JuiType = "easyui-textbox"
	mf1.ListWidth = "60px"

	mf2 := NewSQLFiled{}
	mf2.Add = true
	mf2.List = true
	mf2.Label = "标签"
	mf2.Name = "fname"
	mf2.AddRequire = true
	mf2.JuiType = "easyui-textbox"

	mf3 := NewSQLFiled{}
	mf3.Add = true
	mf3.List = true
	mf3.Label = "模块"
	mf3.Name = "mods"
	mf3.JuiType = "easyui-combobox"
	//mf3.FDataOption = "valueField: 'id', textField: 'text',  data: [{  \"id\":1,  \"text\":\"成交客户\"},{\"id\":2,\"text\":\"重点客户\"},{\"id\":3,\"text\":\"无效客户\"}] "
	mf3.FDataOption = "valueField: 'id', textField: 'text', url:'/data_select_values/?keys=modsconfig'   "

	mf4 := NewSQLFiled{}
	mf4.Add = true
	mf4.List = true
	mf4.Label = "类型"
	mf4.Name = "ftype"
	mf4.JuiType = "easyui-combobox"
	mf4.FDataOption = `valueField: 'id', textField: 'text',  data: [
	{ "id":"easyui-textbox",  "text":"文本框"}
	,{"id":"easyui-combobox","text":"下拉菜单"}
	,{"id":"easyui-numberspinner","text":"数字下拉"}
	,{"id":"easyui-datebox","text":"日期"}
	
	] `
	//mf4.FDataOption = "valueField: 'text', textField: 'text', url:'/data_select_values/?keys=modsconfig'   "

	mf5 := NewSQLFiled{}
	mf5.Add = true
	mf5.List = true
	mf5.Label = "类型"
	mf5.Name = "frequire"
	mf5.JuiType = "easyui-combobox"
	mf5.FDataOption = "valueField: 'text', textField: 'text',  data: [{  \"id\":1,  \"text\":\"必填\"},{\"id\":2,\"text\":\"非必填\"} ] "

	mfo := NewSQLFiled{}
	mfo.Add = true
	mfo.List = true
	mfo.Label = "排序"
	mfo.Name = "forder"
	mfo.JuiType = "easyui-numberspinner"
	mfo.Value = "50"

	mft := NewSQLFiled{}
	mft.Add = true
	mft.List = true
	mft.Label = "添加帮助"
	mft.Name = "addTip"
	mft.JuiType = "easyui-textbox"

	mf6 := NewSQLFiled{}
	mf6.Add = true
	mf6.List = true
	mf6.Label = "字段选项"
	mf6.Name = "dataoption"
	mf6.JuiType = "easyui-textbox"
	mf6.LabelAddTip = `下拉项例子：valueField: 'id', textField: 'text',  data: [ {id:'5',text :'5rrrr5'}  ]
	模块下拉项：valueField: 'id', textField: 'text', url:'/data_select_values/?keys=modsconfig' 
	`

	var f5s []NewSQLFiled
	f5s = append(f5s, mf1, mf2, mf3, mf4, mf5, mfo, mft, mf6)
	newmf.Fileds = f5s

	mNewSQLModel["app"] = newm
	mNewSQLModel["contact"] = newm2
	mNewSQLModel["config"] = newm3
	mNewSQLModel["modsconfig"] = newmmod

	mNewSQLModel["fsconfig"] = newmf
	InitMySQLMod()
	fmt.Print("\n\n\n", mNewSQLModel, "\n\n\n")

}

func ReloadMySQLMod() {
	InitMySQLMod()
}
func InitMySQLMod() {
	o := orm.NewOrm()
	var dd []NewsqlData
	sql := `select  id  , title ,createtime, column_json(data)   as data
	  from app where type = 'mods' and deleted = 0  
	 order by id `
	/*
		aname   string
		appkeyf string
		appname string

		allowadd    string
		allowedit   string
		apptitlef   string
		allowdelete string
	*/
	fmt.Printf(sql)
	o.Raw(sql).QueryRows(&dd)
	//_, _ :=
	for _, v := range dd {

		fmt.Print("\n\n\n", v, "\n\n\n")
		fmt.Print("\n\n\n", v.Data, "\n\n\n")
		b := []byte(v.Data)
		var modMySQl MySQLModConfig
		json.Unmarshal(b, &modMySQl)
		modMySQl.Id = v.Id
		fmt.Printf("Mod: %+v  \n", modMySQl)

		newModMySQLtmp := NewSQLModel{}
		newModMySQLtmp.AppName = modMySQl.Aname
		newModMySQLtmp.AppKey = modMySQl.Appname
		newModMySQLtmp.DataKey = modMySQl.Appkeyf
		if modMySQl.Allowadd == "1" {
			newModMySQLtmp.AllowAdd = true
		}

		if modMySQl.Allowedit == "1" {
			newModMySQLtmp.Allowupdate = true
		}
		if modMySQl.Allowdelete == "1" {
			newModMySQLtmp.AllowDelete = true
		}
		idstr := fmt.Sprintf("%d", v.Id)
		sqlf := `select  id  , title ,createtime, 
		column_json(data)   as data
	  from app 
	where
	 type = 'fsconfigf' 
	and deleted = 0  
	and  
	 COLUMN_GET(data , 'mods' as int  ) = '` + idstr + `'
	 order by 
	 COLUMN_GET(data , 'forder' as int  )   `

		var dds []NewsqlData
		o.Raw(sqlf).QueryRows(&dds)
		var fieldstmp []NewSQLFiled
		for _, v2 := range dds {
			fmt.Print(v2.Data)
			bs := []byte(v2.Data)
			var fconfig MySQLModFieldConfig
			json.Unmarshal(bs, &fconfig)
			fmt.Print("\nNewConfig:\n", fconfig.Fname)
			mft := NewSQLFiled{}
			mft.Add = true
			mft.List = true
			mft.Label = fconfig.Fname
			mft.Name = fconfig.Aname
			if fconfig.Frequire == "必填" {
				mft.AddRequire = true
			}
			mft.LabelAddTip = fconfig.Addtip
			mft.FDataOption = fconfig.Dataoption
			mft.JuiType = fconfig.Ftype
			mft.ListWidth = "60px"
			fieldstmp = append(fieldstmp, mft)
		}
		newModMySQLtmp.Fileds = fieldstmp
		mNewSQLModel[newModMySQLtmp.AppKey] = newModMySQLtmp
	}
}

type NewSQLSelectValue struct {
	beego.Controller
}

func (this *NewSQLSelectValue) Post() {

	mod := this.GetString("keys")
	if len(mod) < 1 {
		return
	}
	var dd []NewsqlData
	o := orm.NewOrm()
	modKey := mNewSQLModel[mod].DataKey
	sql := "  select  id  , title ,createtime, column_json(data)   from app where type = '" + modKey + "' and deleted = 0   order by id "
	fmt.Printf(sql)
	nums, err := o.Raw(sql).QueryRows(&dd)

	if nums == 0 {
		this.Ctx.WriteString("[]")
		return
	}

	if err == nil && nums > 0 {

		//	fmt.Print(dd)
	}
	i := 0
	strj := "["

	for _, v := range dd {
		if i > 0 {
			strj += ","
		}
		i++
		strid := fmt.Sprintf("%d", v.Id)
		strj += "{ \"id\" : \"" + strid + "\"  , \"text\" : \"" + v.Title + "\" }"
	}
	strj += " ]"
	this.Ctx.WriteString(strj)

}

func (this *NewSQLSave) Post() {
	o := orm.NewOrm()
	sql := ""
	mod := this.GetString("modName")
	modKey := mNewSQLModel[mod].DataKey

	formData := this.Input()
	i := 0
	str := ""
	title := ""
	for k, v := range formData {
		if k == "modName" {
			continue
		}
		if i == 0 {
			str += "\"" + k + "\"," + "\"" + v[0] + "\""
			title = v[0]
		} else {
			str += ",\"" + k + "\"," + "\"" + v[0] + "\""
		}
		i++
	}

	title2 := this.GetString(mNewSQLModel[mod].TitleFiled)
	if len(title2) > 1 {
		title = title2
	}
	tt := time.Now()

	t := tt.String()

	sql = "insert into app(title ,type,createorid,ownerid ,data, createtime) values('" + title + "','" + modKey + "',1,1,COLUMN_CREATE(" + str + ") ,'" + t + "');"

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
	mod := this.GetString("modName")
	formData := this.Input()
	i := 0
	str := ""
	title := ""

	for k, v := range formData {
		if k == "modName" {
			continue
		}
		if i == 0 {
			str += "\"" + k + "\"," + "\"" + v[0] + "\""
			title = v[0]

		} else {
			str += ",\"" + k + "\"," + "\"" + v[0] + "\""
		}
		i++
	}
	title2 := this.GetString(mNewSQLModel[mod].TitleFiled)
	if len(title2) > 1 {
		title = title2
	}
	tt := time.Now()

	t := tt.String()
	sql = " update app set  edittime = '" + t + "',   title = '" + title + "', data =  COLUMN_CREATE(" + str + ")  where id = " + id + "  limit 1 ;"

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

	mod := this.GetString("modName")
	modKey := mNewSQLModel[mod].DataKey
	//fmt.Print(mod)
	//fmt.Println(modKey)

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
	sql := "  select id    from app where type = '" + modKey + "' and deleted = 0    "

	allnum, _ := o.Raw(sql).QueryRows(&dd)
	skips := fmt.Sprintf("%d", skip)
	rowss := fmt.Sprintf("%d", rows)
	sql = "  select id  , title ,createtime, column_json(data) as data  from app where type = '" + modKey + "' and deleted = 0   order by id desc limit " + skips + " ," + rowss

	nums, err := o.Raw(sql).QueryRows(&dd)

	fmt.Println(sql)
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
