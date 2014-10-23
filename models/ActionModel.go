package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表
type DqAction struct {
	Id       int64
	ConID    int       `orm:" null"  `
	Uname    string    `orm:"size(16);null"  `
	StatuID  int       `orm:" null"  `
	ActionID int       `orm:" null"  `
	Notes    string    `orm:"size(122);null"  `
	NotesInt int       `orm:" null"  `
	Addtime  time.Time `orm:"auto_now;type(datetime); null"`
}

func init() {
	orm.RegisterModel(new(DqAction))
}

/*
func AddDqContact(r *DqContact) (int64, error) {

	o := orm.NewOrm()
	c := new(DqContact)
	c.Uname = r.Uname
	c.Mname = r.Mname
	c.Tel = r.Tel
	c.Tel2 = r.Tel2

	id, err := o.Insert(c)
	return id, err
}
*/
