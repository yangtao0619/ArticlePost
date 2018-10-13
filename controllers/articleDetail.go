package controllers

import (
	"github.com/astaxie/beego"
	"shanghai2/models"
	"github.com/astaxie/beego/orm"
)

type ArticleContentController struct {
	beego.Controller
}

func (c *ArticleContentController) ShowArticleContent() {
	//展示文章的内容界面,需要先获取文章的id
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error("get id err:", err)
		return
	}
	beego.Info("get article id is:", id)
	//拿到文章的id之后需要查询一遍数据库
	var article models.Article
	article.Id = id
	newOrm := orm.NewOrm()
	readErr := newOrm.Read(&article)
	if readErr != nil {
		beego.Error("read date err:", readErr)
		return
	}
	//c.Data["TName"] = tName
	//在显示登录界面的时候判断cookie
	userName := c.GetSession("username")
	if userName == nil {
		c.Data["username"] = ""
	} else {
		c.Data["username"] = userName
	}

	article.Acount += 1
	newOrm.Update(&article)

	//将得到的文章详情返回给前端
	beego.Info(article.Acontent)

	/*
	下面是多表插入的操作，步骤如下：
	1.确定要插入的表.向article表中插入students字段
	2.新建student对象，获得插入的信息，给student赋值
	3.使用add方法将student插入到article中
	 */

	//开始执行插入多对多数据的代码
	m2m := newOrm.QueryM2M(&article, "Students")
	var student models.Student
	student.Name = userName.(string)
	newOrm.Read(&student, "Name")
	m2m.Add(student)
	/*
	下面是多表查询的操作
	1.确定要查询的数据，这里要查询浏览了该文章的学生
	2.所以首先获取QuerySeter对象的时候要使用“Student“这个表名
	3.查询的时候要进行过滤，指定Student表中的Articles字段和Article表中Id通过Article Id进行关联
	4.Distinct可以去除冗余
	5.All操作将查询的数据赋值给students
	 */
	//确定要查询的表

	var stus []*models.Student
	studentSeter := newOrm.QueryTable("Student")
	studentSeter.Filter("Articles__Article__Id", article.Id).Distinct().All(&stus)
	article.Students = stus
	//newOrm.LoadRelated(&article,"Students")   //加载所有的数据到article中
	c.Data["article"] = article

	var aType models.ArticleType
	aType.Id = article.AType.Id
	//根据id找name
	newOrm.Read(&aType)
	c.Data["TName"] = aType.TName
	beego.Info("TName is ", aType.TName)

	//拿到username的时候需要将访问数据存到文章数据库中
	//var student models.Student
	//student.Name = userName

	c.Layout = "layout.html"
	c.TplName = "content.html"

}
