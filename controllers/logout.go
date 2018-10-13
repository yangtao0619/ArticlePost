package controllers

import "github.com/astaxie/beego"

type LogoutController struct{
	beego.Controller
}

func (c *LogoutController)Logout(){
	//这里处理退出的逻辑
	c.DelSession("username")
	//然后重定向
	c.Redirect("/login",302)
}