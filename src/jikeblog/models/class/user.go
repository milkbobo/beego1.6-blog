package class

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    string `orm:"pk;size(10)"`
	Nick  string
	Email string
}

func TestORM() {
	o := orm.NewOrm()

	u := User{"jike", "geek", "123@q.com"}
	o.Insert(&u)

	ul := User{Id: "jike"}
	o.Read(&ul)
	fmt.Print(ul)

	ul.Nick = "lisan"
	o.Update(&ul)

	u2 := User{Id: "jike"}
	o.Read(&u2)
	fmt.Println(u2)

	o.Delete(&u2)

	o.Read(&u2)
}
