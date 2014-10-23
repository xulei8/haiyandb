package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	m "github.com/xulei8/haiyandb/models"
)

type Agent struct {
	beego.Controller
}

type CallAction struct {
	beego.Controller
}
type CallLog struct {
	beego.Controller
	object m.DqAction
}

func (this *Agent) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *CallAction) Get() {
	id, _ := this.GetInt("cid")
	if id < 1 {
		return
	}
	this.Data["Cid"] = id
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "callaction.tpl"
}

func (this *CallAction) Post() {
	id, _ := this.GetInt("id")
	aid, _ := this.GetInt("actiontype")
	sid, _ := this.GetInt("statuid")
	nid, _ := this.GetInt("nv")
	note := this.GetString("note")
	o := orm.NewOrm()
	oneC := m.DqContact{Id: id}
	//	o.Read(&oneC)
	oneC.Called = 10
	o.Update(&oneC, "Called")

	u := new(m.DqAction)
	u.Notes = note
	u.ActionID = int(aid)
	u.ConID = int(id)
	u.StatuID = int(sid)
	u.NotesInt = int(nid)
	u.Uname = "10000"

	o.Insert(u)

	this.Data["Cid"] = id
	this.TplNames = "callactionok.tpl"

}

func (this *CallLog) Get() {
	cid := this.GetString("cid")
	var rs []m.DqAction

	o := orm.NewOrm()
	qs := o.QueryTable("dq_action").Filter("ConID", cid).OrderBy("-id")
	//fmt.Println("%q",qs)

	if err := this.SetObjects(qs, &rs); err != nil {
		this.Data["Error"] = err
		beego.Error(err)
	}
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "calllog.tpl"
}

func (this *CallLog) SetObjects(qs orm.QuerySeter, objects interface{}) error {
	_, err := qs.Count()
	//fmt.Println("Count %d",cnt)
	if err != nil {
		return err
	}
	// create paginator

	if cnt, err := qs.Limit(30, 0).All(objects); err != nil {
		return err
	} else {
		this.Data["Objects"] = objects
		this.Data["ObjectsCnt"] = cnt

	}
	return nil
}

type AgentCall struct {
	beego.Controller
}

func (this *AgentCall) Get() {

	id, _ := this.GetInt("cid")
	if id < 1 {
		var onec []m.DqContact
		o := orm.NewOrm()
		if _, err := o.QueryTable("dq_contact").Filter("Called", 0).Limit(1, 0).All(&onec); err != nil {
			return
		} else {
			if len(onec) > 0 {

				this.Data["Website"] = "beego.me"
				this.Data["Email"] = "astaxie@gmail.com"
				this.Data["C"] = onec[0]
				this.TplNames = "agentcall.tpl"
				return
			} else {
				return
			}
		}

	}

	u := m.DqContact{Id: id}
	o := orm.NewOrm()
	o.Read(&u)
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["C"] = u
	this.Data["Cid"] = id

	this.TplNames = "agentcall.tpl"
}
