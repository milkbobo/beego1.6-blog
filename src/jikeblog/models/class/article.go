package class

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id      int
	Title   string `orm:"size(80)"`
	Content string `orm:"type(text)"`
	Author  *User  `orm:"rel(fk);"`
	Replys  int
	Views   int

	Tags []*Tag `orm:"rel(m2m)"`

	Time time.Time `orm:"auto_now_add;type(datatime)"`

	Defunct bool
}

func (a *Article) ReadDB() (err error) {
	o := orm.NewOrm()
	if err = o.Read(a); err != nil {
		beego.Info(err)
	}

	return
}

func (a Article) Create() (n int64, err error) {
	o := orm.NewOrm()
	if n, err = o.Insert(&a); err != nil {
		beego.Info(err)
	}
	return
}

func (a Article) Update() (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(&a); err != nil {
		beego.Info(err)
	}
	return
}

func (a Article) Delete() (err error) {
	a.Defunct = true
	err = a.Update()
	return
}

//高级查询，筛选用户
func (a Article) Gets() (ret []Article) {
	o := orm.NewOrm()
	o.QueryTable("article").Filter("Author", a.Author).Filter("defunct", 0).Filter("defunct", 0).All(&ret)
	return
}
