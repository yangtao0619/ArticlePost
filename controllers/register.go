package controllers

import (
	"github.com/astaxie/beego"
	models2 "shanghai2/models"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) ShowRegister() {
	//展示界面
	c.TplName = "register.html"
}

func (c *RegisterController) HandleRegister() {
	//这里需要将注册的数据写入到数据库中，写入之前需要进行orm操作
	//先取出数据
	username := c.GetString("userName")
	passwd := c.GetString("password")
	if username == "" || passwd == "" {
		c.Data["errMsg"] = "用户名或者密码不能为空"
		beego.Error("用户名或者密码不能为空")
		c.TplName = "register.html"
		return
	}
	//插入数据到数据库
	var student models2.Student
	student.Name = username
	student.PassWd = passwd
	newOrm := orm.NewOrm()
	_, err := newOrm.Insert(&student)
	if err != nil {
		c.Data["errMsg"] = "数据错误，请重新输入！"
		beego.Error("数据错误，请重新输入！")
		c.TplName = "register.html"
		return
	}
	//c.Ctx.WriteString("注册成功")
	c.Redirect("login.html",302)
}
