package routers

import (
	"shanghai2/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//增加过滤器
	beego.Info("执行router的init")
	beego.InsertFilter("/article/*", beego.BeforeExec, LoginCheck)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{}, "get:ShowRegister;post:HandleRegister")
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/article/addArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/article/showArticleList", &controllers.ArticleListController{}, "get:ShowArticleList;post:Post")
	beego.Router("/article/showArticleContent", &controllers.ArticleContentController{}, "get:ShowArticleContent;post:Post")
	beego.Router("/article/editArticle", &controllers.ArticleEditController{}, "get:EditArticle;post:HandleEditArticle")
	beego.Router("/article/deleteArticle", &controllers.ArticleDeleteController{}, "get:DeleteArticle;post:Post")
	beego.Router("/article/addType", &controllers.AddTypeController{}, "get:ShowAddType;post:HandleAddType")
	beego.Router("/article/logout", &controllers.LogoutController{}, "get:Logout;post:Post")
	beego.Router("/article/deleteType", &controllers.TypeDeleteController{}, "get:DeleteType;post:Post")

}

var LoginCheck = func(c *context.Context) {
	//先取出用户名查看是否存在，不存在的话就重定向到登录界面
	userName := c.Input.Session("username")
	beego.Info("username is", userName)
	if userName == nil {
		beego.Info("username is nil")
		c.Redirect(302, "/login")
	}
}
