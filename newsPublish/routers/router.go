package routers

import (
	"newsPublish/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.UserController{},
    "get:ShowRegister;post:HandleRegister")
    beego.Router("/login",&controllers.UserController{},
    "get:ShowLogin;post:HandleLogin")
    beego.Router("/index",&controllers.ArticleController{},
    "get:ShowIndex")
    beego.Router("/addArticle",&controllers.ArticleController{},
    "get:ShowAdd;post:HandleAdd")
    beego.Router("content",&controllers.ArticleController{},
    "get:ShowContent")
    beego.Router("editArticle",&controllers.ArticleController{},
    "get:ShowEditArticle;post:HandleEditArticle")
    beego.Router("deleteArticle",&controllers.ArticleController{},
    "get:HandleDelete")
    beego.Router("addType",&controllers.ArticleController{},
    "get:ShowAddType;post:HandleAddType")
}
