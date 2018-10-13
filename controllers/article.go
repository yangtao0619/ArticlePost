package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shanghai2/models"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) ShowAddArticle() {
	newOrm := orm.NewOrm()
	var types []models.ArticleType
	typeSeter := newOrm.QueryTable("ArticleType")
	//将查询到的数据全部复制给切片
	typeSeter.All(&types)
	//if err != nil {
	//	informAndRedirectInAddArticle(c, "切片写入失败", "add.html")
	//	return
	//}
	//在显示登录界面的时候判断cookie
	userName := c.GetSession("username")
	if userName == ""{
		c.Data["username"] = ""
		c.Data["checked"] = ""
	}else{
		c.Data["username"] = userName
		c.Data["checked"] = "checked"
	}

	c.Data["articleTypes"] = types
	c.Layout = "layout.html"
	c.TplName = "add.html"
}

func (c *ArticleController) HandleAddArticle() {
	//文章名和文章的内容不能为空
	articleName := c.GetString("articleName")
	content := c.GetString("content")
	//获取下拉列表选择的数据
	TName := c.GetString("select")
	beego.Info("select is ", TName)
	imgPath := HandleUploadFile(&c.Controller, "uploadname")
	//判空操作
	if articleName == "" || content == "" || imgPath == "" {
		c.Data["errmsg"] = "文章名称或文章内容或者图片不能为空！"
		beego.Error("文章名称或文章内容或者图片不能为空！")
		c.TplName = "add.html"
		return
	}

	//如果保存文件正常的话，需要将文章信息保存到数据库中
	newOrm := orm.NewOrm()
	var article models.Article
	article.Acontent = content
	article.ArtName = articleName
	article.Aimg = imgPath
	//将数据写入
	var articleType models.ArticleType
	articleType.TName = TName
	newOrm.Read(&articleType, "TName")
	if articleType.Id <= 0{
		articleType.Id = 1
	}
	beego.Info("insert type is", strconv.Itoa(articleType.Id), articleType.TName)
	article.AType = &articleType
	//article.AType = &models.ArticleType{TName: TName}
	//将选择的数据插入数据库
	_, err := newOrm.Insert(&article)
	if err != nil {
		c.Data["errmsg"] = "insert data err"
		beego.Error("insert data err:", err)
		c.TplName = "add.html"
		return
	}
	c.Redirect("/article/showArticleList", 302)
}
