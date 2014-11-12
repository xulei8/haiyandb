package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

type NewSQLApp struct {
	beego.Controller
}

func (this *NewSQLApp) Get() {

	mod := this.GetString("app")
	this.Data["AppName"] = mNewSQLModel[mod].AppName
	this.Data["ModName"] = mod

	this.Data["ToolNew"] = mNewSQLModel[mod].AllowAdd
	this.Data["ToolEdit"] = mNewSQLModel[mod].Allowupdate
	this.Data["ToolDelete"] = mNewSQLModel[mod].AllowDelete
	this.Data["FS"] = mNewSQLModel[mod].Fileds

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *MainController) Get() {
	ReloadMySQLMod()
	mod := "app"
	this.Data["AppName"] = mNewSQLModel[mod].AppName
	this.Data["ModName"] = mod

	this.Data["ToolNew"] = mNewSQLModel[mod].AllowAdd
	this.Data["ToolEdit"] = mNewSQLModel[mod].Allowupdate
	this.Data["ToolDelete"] = mNewSQLModel[mod].AllowDelete
	this.Data["FS"] = mNewSQLModel[mod].Fileds

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"

	this.TplNames = "index.tpl"
}

type HelpController struct {
	beego.Controller
}

func NewSQLAPP(app string) string {

	str := ""

	return str
}

func (this *HelpController) Get() {

	this.Data["AppName"] = mNewSQLModel["app"].AppName
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
