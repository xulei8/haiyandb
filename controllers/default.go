package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	this.Data["AppName"] = mNewSQLModel["app"].AppName
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
