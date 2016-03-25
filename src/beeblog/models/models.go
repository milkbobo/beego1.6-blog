package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      time.Time
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index;null"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func AddTopic(title, content string) error {

	o := orm.NewOrm()

	topic := &Topic{
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}

	_, err := o.Insert(topic)
	return err
}

func RegisterDB() {

	orm.Debug = true
	// 注册模型
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@tcp(localhost:3306)/main?charset=utf8")

}

func AddCategory(name string) error {
	o := orm.NewOrm()

	fmt.Println(name)

	cate := &Category{Title: name}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)

	//找不到返回nil
	if err == nil {
		fmt.Println("qs.Filter!OK!")
		return err

	}

	fmt.Println("Filter!OK!")

	_, err = o.Insert(cate)
	//有错误的话
	if err != nil {
		return err
	}

	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetAllTopics(isDese bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error
	if isDese {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	fmt.Sprint(qs)
	return cates, err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)
	return topic, err

}

func ModifyTopic(tid, title, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}

	return nil
}
