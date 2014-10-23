package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表
type DqContact struct {
	Id      int64
	Sex     int    `orm:" null"  `
	Uname   string `orm:"size(16);null"  `
	Piname  string `orm:"size(16);null"  `
	Tel     string `orm:"size(22);null;index"  `
	Tel2    string `orm:"size(22);null"  `
	Comname string `orm:"size(33);null"  `
	Mname   string `orm:"size(16);null"  `
	Address string `orm:"size(32);null"  `
	Note    string `orm:"size(152);null"  `
	Email   string `orm:"size(32);null"  `
	Qq      string `orm:"size(32);null"  `

	Called        int       `orm:" null"  `
	Hits          int       `orm:" null"  `
	CallHits      int       `orm:" null"  `
	Statu         int       `orm:" null"  `
	Lockit        int       `orm:" null"  `
	Addtime       time.Time `orm:"auto_now;type(datetime); null"`
	FirstCallTime time.Time `orm:"auto_now;type(datetime); null"`
	LastCalltime  time.Time `orm:"auto_now;type(datetime); null"`
	EditTime      time.Time `orm:"auto_now;type(datetime); null"`
}

func init() {
	orm.RegisterModel(new(DqContact))
}

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
