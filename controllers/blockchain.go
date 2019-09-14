package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/yushizhao/hackathon2019/blockchains"
)

// Operations about Users
type BlockchainController struct {
	beego.Controller
}

// @Title Balance
// @Description check balance
// @Success 200 {string} hello
// @router /balance [get]
func (c *BlockchainController) Balance() {
	// b, err := blockchains.Ethplorer.Balance(models.MasterKey.ETH.Address.Hex())
	// Free api keys have strict key limits. We use mock data to ensure demo success.
	b, err := blockchains.Ethplorer.MOCKBalance()

	if err != nil {
		logs.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = b
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
