package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"newsPublish/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) ShowRegister() {
	this.TplName = "register.html"
}

//处理注册数据
func (this *UserController) HandleRegister() {
	userName := this.GetString("userName")
	pwd := this.GetString("password")
	if userName == "" || pwd == "" {
		beego.Error("用户名或密码不能为空")
		this.TplName = "register.html"
		return
	}
	//	操作数据
	o := orm.NewOrm()
	var user models.User
	user.Name = userName
	user.Pwd = pwd
	_, err := o.Insert(&user)
	if err != nil {
		beego.Error("用户注册失败")
		this.TplName = "register.html"
		return
	}
	this.Ctx.WriteString("用户注册成功")

}
func (this*UserController)	ShowLogin()  {
	this.TplName = "login.html"
}
func (this*UserController)HandleLogin()  {
	userName:=this.GetString("userName")
	pwd:=this.GetString("password")
	if userName=="" ||pwd==""{
		beego.Error("用户名或密码不能为空")
		this .TplName = "Login.html"
		return
	}
	o := orm.NewOrm()
	var user models.User
	user.Name = userName
	err:=o.Read(&user,"Name")
	if err!=nil{
		beego.Error("用户不全在")
		this.TplName = "login.html"
		return
	}
	if user.Pwd !=pwd {
		beego.Error("输入的密码错误")
		this.TplName = "login.html"
		return
	}
	//this.Ctx.WriteString("登陆成功")
	this.Redirect("/index",302)
}
