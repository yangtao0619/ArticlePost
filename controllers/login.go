package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shanghai2/models"
	"encoding/base64"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) ShowLogin() {
	//在显示登录界面的时候判断cookie
	userName := c.Ctx.GetCookie("username")
	//使用base64加密解密
	usernameBytes, _ := base64.StdEncoding.DecodeString(userName)
	userName = string(usernameBytes)
	if userName == ""{
		c.Data["username"] = ""
		c.Data["checked"] = ""
	}else{
		c.Data["username"] = userName
		c.Data["checked"] = "checked"
	}
	c.TplName = "login.html"
}

func (c *LoginController) HandleLogin() {
	//需要先获取用户输入的数据，然后和数据库内容进行校验
	username := c.GetString("userName")
	passwd := c.GetString("password")
	if username == "" || passwd == ""{
		c.Data["errMsg"] = "用户名或者密码不能为空"
		beego.Error("用户名或者密码不能为空")
		c.TplName = "login.html"
		return
	}
	//需要先查找一边数据库
	newOrm := orm.NewOrm()
	var student models.Student
	student.Name = username
	beego.Error("查找前的passwd：",student.PassWd)

	readErr := newOrm.Read(&student, "Name")

	beego.Error("查找后的passwd：",student.PassWd)
	if readErr != nil{
		c.Data["errMsg"] = "用户不存在"
		beego.Error("用户不存在")
		c.TplName = "login.html"
		return
	}

	if student.PassWd != passwd{
		c.Data["errMsg"] = "密码错误，请重新输入！"
		beego.Error("密码错误，请重新输入！")
		c.TplName = "login.html"
		return
	}

	remember := c.GetString("remember")
	beego.Info("remember is",remember)
	if remember == "on"{
		c.Ctx.SetCookie("username",base64.StdEncoding.EncodeToString([]byte(username)),1000000)
	}else{
		c.Ctx.SetCookie("username",username,-1)
	}
	//登录成功的时候，服务器要写入session
	c.SetSession("username",username)
	c.Redirect("/article/showArticleList",302)
}
