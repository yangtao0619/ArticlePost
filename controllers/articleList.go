package controllers

import (
	"github.com/astaxie/beego"
	"shanghai2/models"
	"github.com/astaxie/beego/orm"
	"math"
)

type ArticleListController struct {
	beego.Controller
}

//查询数据库内容，返回给前端界面
func (c *ArticleListController) ShowArticleList() {
	//得到select的值
	typeName := c.GetString("select")
	beego.Info("typeName is", typeName)
	//定义一个切片，用于存储文章信息
	var articleList []models.Article
	newOrm := orm.NewOrm()
	artSet := newOrm.QueryTable("article")

	//给数据做分页,首先需要获取所有的记录数,获得总记录数也要根据是否选中下拉列表来区分
	var count int64
	var err error
	if typeName == ""{
		count, err = artSet.RelatedSel("AType").Count()
	}else{
		count,err = artSet.RelatedSel("AType").Filter("AType__TName", typeName).Count()
	}
	if err != nil {
		beego.Error("获取总记录数失败,", err)
		return
	}
	//确定每页显示的数量
	pageSize := 2
	beego.Info("count is ", count, "pageSize is ", pageSize)
	c.Data["colCount"] = count
	c.Data["pageCount"] = int(math.Ceil(float64(count) / float64(pageSize)))

	//可以获得用户请求的页码,默认请求页码是1
	pageIndex, err := c.GetInt("pageIndex", 1)
	beego.Info("request page index is ", pageIndex, err)
	c.Data["pageIndex"] = pageIndex
	//这里需要将对应页码的数据返回给浏览器
	if typeName != "" {
		artSet.Limit(pageSize, (pageIndex-1)*2).RelatedSel("AType").Filter("AType__TName", typeName).All(&articleList)
	} else {
		artSet.Limit(pageSize, (pageIndex-1)*2).RelatedSel("AType").All(&articleList)
	}
	beego.Info("return data success!")
	c.Data["articles"] = articleList
	for _, article := range articleList {
		beego.Info(article.AType)
	}

	//获取cookie中的session
	userName := c.GetSession("username")
	if userName == nil {
		c.Data["username"] = ""
	} else {
		c.Data["username"] = userName
	}

	//还需要查询数据库，给select赋值
	var types []models.ArticleType
	typeSeter := newOrm.QueryTable("ArticleType")
	typeSeter.All(&types)
	c.Data["types"] = types
	c.Data["typeName"] = typeName
	c.Layout = "layout.html"
	c.TplName = "index.html"
}
