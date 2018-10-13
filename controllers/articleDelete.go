package controllers

import (
	"github.com/astaxie/beego"
	"shanghai2/models"
	"github.com/astaxie/beego/orm"
)

type ArticleDeleteController struct {
	beego.Controller
}

func (c *ArticleDeleteController) DeleteArticle() {
	//处理删除文章的请求,首先获得要删除文章的id
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error("要删除的文章id不存在")
		c.Redirect("/article/showArticleList", 302)
		return
	}
	//从数据库中删除文章
	var article models.Article
	newOrm := orm.NewOrm()
	article.Id = id
	readErr := newOrm.Read(&article)
	if readErr != nil {
		//读取数据库有错误
		beego.Error("要删除的文章不存在")
		c.Redirect("/article/showArticleList", 302)
		return
	}
	_, err = newOrm.Delete(&article)
	if err != nil {
		beego.Error("文章删除失败!")
		c.Redirect("/article/showArticleList", 302)
		return
	}

	c.Redirect("/article/showArticleList", 302)
}
