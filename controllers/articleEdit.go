package controllers

import (
	"github.com/astaxie/beego"
	"shanghai2/models"
	"github.com/astaxie/beego/orm"
	"path"
	"time"
	"strconv"
)

type ArticleEditController struct {
	beego.Controller
}

func (c *ArticleEditController) EditArticle() {
	//这里处理编辑文章的逻辑
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
	c.Data["article"] = article
	c.Layout = "layout.html"
	c.TplName = "update.html"
}

/*
处理编辑文章之后的表单提交请求
 */
func (c *ArticleEditController) HandleEditArticle() {
	beego.Info("抓取post请求的内容")
	//这里需要重新的更新内容,和addArticle的步骤类似
	//文章名和文章的内容不能为空
	articleName := c.GetString("articleName")
	content := c.GetString("content")
	id, err := c.GetInt("id", 1)
	beego.Info("id is ", id, "name is ", articleName, "content is ", content)
	//得到图片存放的路径
	imgPath := HandleUploadFile(&c.Controller, "uploadname")
	if articleName == "" || content == "" || imgPath == "" {
		c.Data["errmsg"] = "信息填写错误"
		beego.Error("信息填写错误!")
		c.Redirect("/article/editArticle?id="+strconv.Itoa(id), 302)
		return
	}

	//如果保存文件正常的话，需要将文章信息保存到数据库中
	newOrm := orm.NewOrm()
	var article models.Article
	article.Id = id
	//更新数据之前需要先查询一遍数据库
	//注意read的时候需要先给对象传入id
	readErr := newOrm.Read(&article, "id")
	if readErr != nil {
		c.Data["errmsg"] = "要修改的文章不存在"
		beego.Error("要修改的文章不存在")
		c.Redirect("/article/editArticle?id="+strconv.Itoa(id), 302)
		return
	}
	//read之后才能给对象赋值
	article.Acontent = content
	article.ArtName = articleName
	if imgPath != "NoImg" {
		article.Aimg = imgPath
	}
	_, err = newOrm.Update(&article)
	if err != nil {
		c.Data["errmsg"] = "insert data err"
		beego.Error("insert data err:", err)
		c.Redirect("/article/editArticle?id="+strconv.Itoa(id), 302)
		return
	}
	//如果什么错误都没有的话,重定向到展示新闻列表页面
	c.Redirect("/article/showArticleList", 302)
}
func HandleUploadFile(c *beego.Controller, filePath string) string {
	//这里需要处理以下上传文件的操作，首先需要更改表单的一个属性
	file, header, err := c.GetFile("uploadname")
	if header.Filename == "" {
		return "NoImg"
	}
	defer file.Close()
	if err != nil {
		c.Data["errmsg"] = "get file err"
		beego.Error("get file err:", err)
		return ""
	}
	//过滤文件名
	ext := path.Ext(header.Filename)
	if ".jpg" != ext && ".jpeg" != ext && ".png" != ext {
		//文件格式不对
		c.Data["errmsg"] = "文件格式不正确"
		beego.Error("文件格式不正确！ext is ", ext)
		return ""
	}
	//过滤文件大小
	if header.Size > 5000000 {
		c.Data["errmsg"] = "文件超出大小限制"
		beego.Error("文件超出大小限制！")
		return ""
	}
	fileName := time.Now().Format("2006-01-02-15-04-05")
	err = c.SaveToFile(filePath, "./static/img/"+fileName+ext)
	if err != nil {
		c.Data["errmsg"] = "上传文件失败"
		beego.Error("上传文件失败！")
		return ""
	}
	return "/static/img/" + fileName + ext

}
