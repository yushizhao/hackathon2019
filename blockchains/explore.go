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

// Free api keys have strict key limits. We use mock data to ensure demo success.
func (e *ETHExplorer) MOCKBalance() (interface{}, error) {
	b := make(map[string]float64)
	b["ETH"] = 12.983
	b["USDC"] = 18848.0
	b["QTUM"] = 3668
	return b, nil
}
