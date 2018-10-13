package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shanghai2/models"
)

type AddTypeController struct {
	beego.Controller
}

func (c *AddTypeController) ShowAddType() {
	//展示数据之前需要读取数据并将数据写给前端界面
	newOrm := orm.NewOrm()
	var types []models.ArticleType
	typeSeter := newOrm.QueryTable("article_type")
	//将查询到的数据全部复制给切片
	typeSeter.All(&types)
	c.Data["articleTypes"] = types
	//在显示登录界面的时候判断cookie
	userName := c.GetSession("username")
	beego.Info("username is",userName)
	if userName == "" {
		c.Data["username"] = ""
		c.Data["checked"] = ""
	} else {
		c.Data["username"] = userName
		c.Data["checked"] = "checked"
	}
	c.Layout = "layout.html"
	c.TplName = "addType.html"
}

func (c *AddTypeController) HandleAddType() {
	//接收post请求
	beego.Info("get post data")
	//这里需要获取浏览器传递过来的数据
	typeName := c.GetString("typeName")
	if typeName == "" {
		informAndRedirect(c, "类型字段不能为空", "addType.html")
		return
	}
	//插入数据库
	newOrm := orm.NewOrm()
	var aType models.ArticleType
	aType.TName = typeName
	_, err := newOrm.Insert(&aType)
	if err != nil {
		informAndRedirect(c, "插入类型数据失败", "addType.html")
		return
	}
	//如果插入正确的话，重定向到本页面
	informAndRedirect(c, "插入成功", "addType.html")
}

//处理通知信息的函数
func informAndRedirect(c *AddTypeController, errMsg string, pageName string) {
	c.Data["errMsg"] = errMsg
	beego.Error(errMsg)
	c.Redirect(pageName, 302)
}
