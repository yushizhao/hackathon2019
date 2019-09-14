package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yushizhao/hackathon2019/models"
)

type KeyController struct {
	beego.Controller
}

// @Title Display
// @Description Display addressed
// @Success 200
// @router /display [get]
func (c *KeyController) Display() {

	c.Data["json"] = models.MasterKey.Display()
	c.ServeJSON()
}
