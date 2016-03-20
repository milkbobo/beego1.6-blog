package class

import (
	"github.com/astaxie/beego/orm"
)

type Tag struct {
	ID       int64
	Name     string     `orm:"index"`
	Articles []*Article `orm:"reverse(many)"`
}

func (t Tag) Get() *Tag {
	o := orm.NewOrm()
	o.QueryTable("tag").Filter("Name", t.Name).One(&t)
	return &t
}

func (t Tag) GetOrNew() *Tag {
	o := orm.NewOrm()
	_, _, _ = o.ReadOrCreate(&t, "Name")
	return &t
}
