package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shanghai2/models"
)

type TypeDeleteController struct {
	beego.Controller
}

func (c *TypeDeleteController) DeleteType() {
	beego.Info("删除类型请求")
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error("获取id失败")
		c.Redirect("/login", 302)
		return
	}

	newOrm := orm.NewOrm()
	var aType models.ArticleType
	aType.Id = id
	newOrm.Delete(&aType)

	c.Redirect("/article/addType",302)
}
