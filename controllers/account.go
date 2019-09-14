package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yushizhao/hackathon2019/models"
)

type AccountController struct {
	beego.Controller
}

// @Title Display
// @Description Display accounts
// @Success 200
// @router /display [get]
func (c *AccountController) Display() {

	c.Data["json"] = models.Accounts
	c.ServeJSON()
}
