package models

type Account struct {
	Blockchain string
	Device     string
	Address    string
	Function   interface{}
}

var Accounts []Account
