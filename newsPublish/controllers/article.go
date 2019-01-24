package controllers

import (
	"github.com/astaxie/beego"

	"time"
	"path"
	"github.com/astaxie/beego/orm"
	"newsPublish/models"
	"math"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) ShowIndex() {
	o := orm.NewOrm()
	typeName :=this.GetString("select")
	pageIndex, err := this.GetInt("pageIndex")
	if err != nil {
		pageIndex = 1
	}

	var articles []models.Article
	var count int64
	pageSize := 3
	start := pageSize * (pageIndex - 1)

	if typeName==""{

		qs := o.QueryTable("Article")
		qs.RelatedSel("ArticleType").Limit(pageSize, start).All(&articles)
		count, _= qs.RelatedSel("ArticleType").Count()
	}else {

		qs := o.QueryTable("Article")
		qs.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).All(&articles)
		count,_=qs.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).Count()
	}


	pageCount := math.Ceil(float64(count) / float64(pageSize))

	var articleTypes []models.ArticleType
	o.QueryTable("ArticleType").All(&articleTypes)
	this.Data["articleTypes"] = articleTypes

	beego.Error(typeName)
    this.Data["typeName"]=typeName
	this.Data["pageIndex"] = pageIndex
	this.Data["count"] = count
	this.Data["pageCount"] = pageCount
	this.Data["articles"] = articles

	this.TplName = "index.html"

}
func (this *ArticleController) ShowAdd() {
	o := orm.NewOrm()
	var articleTypes []models.ArticleType
	o.QueryTable("ArticleType").All(&articleTypes)
	this.Data["articleTypes"] = articleTypes
	this.TplName = "add.html"
}
func (this *ArticleController) HandleAdd() {

	articleName := this.GetString("articleName")
	content := this.GetString("content")
	file, head, err := this.GetFile("uploadname")
	if articleName == "" || content == "" || err != nil {
		beego.Error("获取用户添加数据失败")
		this.TplName = "add.html"
		return
	}
	defer file.Close()
	if head.Size > 5000000 {
		beego.Error("图片太大，接收不了！")
		this.TplName = "add.html"
		return
	}
	ext := path.Ext(head.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		beego.Error("文件格式不正确")
		this.TplName = "add.html"
		return
	}
	fileName := time.Now().Format("20060102150405")

	this.SaveToFile("uploadname", "./static/img/"+fileName+ext)

	o := orm.NewOrm()
	var article models.Article
	article.Title = articleName
	article.Content = content
	article.Img = "/static/img/" + fileName + ext
	typeName := this.GetString("select")
	var articleType models.ArticleType

	articleType.TypeName = typeName

	o.Read(&articleType,"TypeName")

	article.ArticleType = &articleType

	o.Insert(&article)
	this.Redirect("/index", 302)

}
func (this *ArticleController) ShowContent() {
	id, err := this.GetInt("id")
	if err != nil {
		beego.Error("获取数据错误")
		this.TplName = "index.html"
		return
	}
	o := orm.NewOrm()
	var article models.Article
	article.Id2 = id
	o.Read(&article)
	article.ReadCount += 1
	o.Update(&article)
	this.Data["article"] = article
	this.TplName = "content.html"
}
func (this *ArticleController) ShowEditArticle() {
	id, err := this.GetInt("id")
	if err != nil {
		beego.Error("获取数据错误", err)
		this.TplName = "index.html"
		return
	}
	o := orm.NewOrm()
	var article models.Article
	article.Id2 = id
	o.Read(&article)
	this.Data["article"] = article
	this.TplName = "update.html"
}
func UploadFunc(this *ArticleController, fileName string) string {
	file, head, err := this.GetFile(fileName)
	if err != nil {
		beego.Error("获取用户添加数据失败")
		this.TplName = "add.html"
		return ""
	}
	defer file.Close()
	if head.Size > 5000000 {
		beego.Error("图片太大，接收不了！")
		this.TplName = "add.html"
		return ""
	}
	ext := path.Ext(head.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		beego.Error("文件格式不正确")
		this.TplName = "add.html"
		return ""
	}
	filePath := time.Now().Format("20060102150405")

	this.SaveToFile(fileName, "./static/img/"+filePath+ext)
	return "/static/img/" + filePath + ext
}

func (this *ArticleController) HandleEditArticle() {
	id, err := this.GetInt("id")
	content := this.GetString("content")
	articleName := this.GetString("articleName")
	filePath := UploadFunc(this, "uploadname")
	if err != nil || articleName == "" || content == "" || filePath == "" {
		beego.Error("获取数据错误")
		this.TplName = "update.html"
		return
	}
	o := orm.NewOrm()
	var article models.Article
	article.Id2 = id
	err = o.Read(&article)
	if err != nil {
		beego.Error("更新数据不存在")
		this.TplName = "update.html"
		return
	}
	article.Title = articleName
	article.Content = content
	article.Img = filePath
	o.Update(&article)
	this.Redirect("/index", 302)
}
func (this*ArticleController) HandleDelete() {
	id,err :=this.GetInt("id")
	if err != nil{
		beego.Error("删除请求数据错误")
		this.TplName = "index.html"
		return
	}
	o :=orm.NewOrm()
	var article models.Article
	article.Id2 = id
	_,err = o.Delete(&article)
	if err!=nil{
		beego.Error("删除失败")
		this.TplName = "index.html"
		return
	}
	this.Redirect("/index",302)
}
func (this*ArticleController)ShowAddType(){
	o := orm.NewOrm()
	qs := o.QueryTable("ArticleType")
	var articleTypes []models.ArticleType
	qs.All(&articleTypes)
	this.Data["articleTypes"] = articleTypes
	this.TplName = "addType.html"
}
func (this*ArticleController)HandleAddType(){
	typeName := this.GetString("typeName")
	if typeName == ""{
		beego.Error("数据类型不能为空")
		this.TplName = "addType.html"
		return
	}
	o :=orm.NewOrm()
	var articleType models.ArticleType
	articleType.TypeName = typeName
	o.Insert(&articleType)
	this.Redirect("/addType",302)
}
