package controllers

import "github.com/astaxie/beego"

// Operations about Users
type BlockchainController struct {
	beego.Controller
}

// @Title Balance
// @Description check balance
// @Success 200 {string} hello
// @router /balance [get]
func (c *BlockchainController) Balance() {

	c.Data["json"] = "hello"
	c.ServeJSON()
}

// @Title Transaction
// @Description check transactions
// @Success 200 {string} hello
// @router /transaction [get]
func (c *BlockchainController) Transaction() {

	c.Data["json"] = "hello"
	c.ServeJSON()
}
