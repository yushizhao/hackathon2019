package models

import "github.com/astaxie/beego/logs"

// mock data for demo
func init() {
	err := MasterKey.Generate()
	if err != nil {
		logs.Error(err)
	}

	Accounts = make([]Account, 4)

	A1 := Account{
		Blockchain: "ETH",
		Device:     "the car",
		Address:    MasterKey.ETH.Address.Hex(),
	}

	A2 := Account{
		Blockchain: "ETH",
		Device:     "Sunny's phone",
		Address:    "0x2720325facB1e55B412642C67F52bBF23399125c",
		Function:   []string{"user-level car API"},
	}

	A3 := Account{
		Blockchain: "ETH",
		Device:     "Victor's car",
		Address:    "0x306E19D4081267072288b05259233AB49f0Baa09",
		Function:   []string{"say hello", "share location", "vote"},
	}

	A4 := Account{
		Blockchain: "ETH",
		Device:     "Sunny's cloud",
		Address:    "0x5569c89d1382C04fb394d62F8FF95Ccc9CA99122",
		Function:   []string{"cloud drive"},
	}

	Accounts[0] = A1
	Accounts[1] = A2
	Accounts[2] = A3
	Accounts[3] = A4
}
