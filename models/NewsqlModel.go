package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表
type NewsqlData struct {
	Id        int64
	Creatorid int    `orm:" null"  `
	Ownerid   int    `orm:" null"  `
	App       string `orm:"size(20);null"  `
	Title     string `orm:"size(100);null"  `

	Createtime time.Time `orm:"auto_now;type(datetime); null"`

	EditTime time.Time `orm:"type(datetime); null"`
}

func init() {
	orm.RegisterModel(new(NewsqlData))
}

func AddNewsqlData(r *NewsqlData) (int64, error) {

	o := orm.NewOrm()
	c := new(NewsqlData)
	c.Title = r.Title
	id, err := o.Insert(c)
	return id, err
}
