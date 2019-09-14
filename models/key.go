package models

import (
	"github.com/yushizhao/hackathon2019/blockchains"
)

type Key struct {
	BTC *blockchains.BTCStyleKey
	ETH *blockchains.ETHStyleKey
}

var MasterKey Key

func (k *Key) Generate() (err error) {
	k.BTC, err = blockchains.NewBTCStyleKey()
	if err != nil {
		return err
	}
	k.ETH, err = blockchains.NewETHStyleKey()
	if err != nil {
		return err
	}
	return nil
}

func (k *Key) Display() map[string]string {
	addrs := make(map[string]string)
	addrs["BTC"] = k.BTC.Addr
	addrs["ETH"] = k.ETH.Address.Hex()
	return addrs
}
