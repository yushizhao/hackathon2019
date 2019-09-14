package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	explore "github.com/yushizhao/hackathon2019/explorers"
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
	b, err := explore.Ethplorer.MOCKBalance()

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
	// b, err := blockchains.Ethplorer.Transaction("0xb297cacf0f91c86dd9d2fb47c6d12783121ab780")
	// Free api keys have strict key limits. We use mock data to ensure demo success.
	b, err := explore.Ethplorer.MOCKTransaction()

	if err != nil {
		logs.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = b
	c.ServeJSON()
}
