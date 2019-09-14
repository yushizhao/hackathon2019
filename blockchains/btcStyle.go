package blockchains

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

var NET = &chaincfg.MainNetParams

type BTCStyleKey struct {
	Master  *hdkeychain.ExtendedKey
	ChildID uint32
	Addr    string
}

// keyToAddr maps the passed private to corresponding p2pkh address.
func keyToAddr(key *btcec.PrivateKey, net *chaincfg.Params) (btcutil.Address, error) {
	serializedKey := key.PubKey().SerializeCompressed()
	pubKeyAddr, err := btcutil.NewAddressPubKey(serializedKey, net)
	if err != nil {
		return nil, err
	}
	return pubKeyAddr.AddressPubKeyHash(), nil
}

func NewBTCStyleKey() (*BTCStyleKey, error) {
	var key BTCStyleKey
	var err error
	seed := "MOCK hackathon2019 MOCK"
	key.Master, err = hdkeychain.NewMaster([]byte(seed), NET)
	if err != nil {
		return &key, err
	}

	key.ChildID = 1
	child, err := key.Master.Child(key.ChildID)
	key.ChildID++
	if err != nil {
		return &key, err
	}

	pk, err := child.ECPrivKey()
	if err != nil {
		return &key, err
	}

	addr, err := keyToAddr(pk, NET)
	if err != nil {
		return &key, err
	}
	key.Addr = addr.EncodeAddress()

	return &key, nil
}
