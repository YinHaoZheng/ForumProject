package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Layout = "layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["LayoutHeader"] = "baseLayout/header.html"
	this.LayoutSections["NaviBar"] = "baseLayout/navibar.html"
	this.LayoutSections["SearchWrapper"] = "baseLayout/searchwrapper.html"
	this.LayoutSections["Footer"] = "baseLayout/footer.html"
	this.LayoutSections["Script"] = "baseLayout/script.html"
	this.TplName = "index/index.html"
}

