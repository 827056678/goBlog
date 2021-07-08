package controllers

import beego "github.com/beego/beego/v2/server/web"

type IndexController struct {
	beego.Controller
}

func (i *IndexController) Get() {
	i.Ctx.WriteString("index.get")
}
