package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/astaxie/beego/orm"
)

//用户
type Student struct {
	Id       int
	Name     string
	PassWd   string
	Articles []*Article `orm:"rel(m2m)"` //设置多对多
}

//文章   用户和文章应该是多对多的关系   文章和文章类型是一对多的关系
type Article struct {
	Id       int          `orm:"pk;auto"`
	ArtName  string       `orm:"size(20)"`
	Artime   time.Time    `orm:"auto_now"`
	Acount   int          `orm:"default(0);null"`
	Acontent string       `orm:"size(500)"`
	Aimg     string       `orm:"size(100)"`
	Students []*Student   `orm:"reverse(many)"` //设置多对多的反向关系
	AType    *ArticleType `orm:"rel(fk);on_delete(set_null);null"`       //Article和ArticleType是多对一的关系
}

//文章类型
type ArticleType struct {
	Id       int
	TName    string
	Articles []*Article `orm:"reverse(many)"` //注意这里的many和[]的绑定关系
}

func init() {
	////注册数据库
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/testbase?charset=utf8")
	//注册model
	orm.RegisterModel(new(Student), new(Article), new(ArticleType))
	//同步
	orm.RunSyncdb("default", false, true)

	////测试数据库操作
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testbase?charset=utf8")
	//if err != nil {
	//	beego.Error("open sql err")
	//	return
	//}
	//var count int
	//rows, _ := db.Query(`select count(*) from article`)
	//rows.Scan(&count)
	//beego.Error("count is ", count)
}
