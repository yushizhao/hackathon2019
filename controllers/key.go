package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type KeyController struct {
	beego.Controller
}

// @Title Display
// @Description Display addressed
// @Success 200 {string} hello
// @router /display [get]
func (c *KeyController) Display() {

	c.Data["json"] = "hello"
	c.ServeJSON()
}
