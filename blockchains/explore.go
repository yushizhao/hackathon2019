package blockchains

import "github.com/astaxie/beego/httplib"

type ETHExplorer struct {
	URL    string
	APIKEY string
}

var Ethplorer ETHExplorer

func init() {
	Ethplorer.URL = "http://api.ethplorer.io/"
	Ethplorer.APIKEY = "freekey"
}

func (e *ETHExplorer) Balance(addr string) ([]byte, error) {
	endpoint := "getAddressInfo/" + addr + "?apiKey=" + Ethplorer.APIKEY
	req := httplib.Get(e.URL + endpoint)
	b, err := req.Bytes()
	if err != nil {
		return nil, err
	}
	return b, nil
}
