package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/yushizhao/hackathon2019/controllers:BlockchainController"] = append(beego.GlobalControllerRouter["github.com/yushizhao/hackathon2019/controllers:BlockchainController"],
        beego.ControllerComments{
            Method: "Balance",
            Router: `/balance`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yushizhao/hackathon2019/controllers:BlockchainController"] = append(beego.GlobalControllerRouter["github.com/yushizhao/hackathon2019/controllers:BlockchainController"],
        beego.ControllerComments{
            Method: "Transaction",
            Router: `/transaction`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yushizhao/hackathon2019/controllers:KeyController"] = append(beego.GlobalControllerRouter["github.com/yushizhao/hackathon2019/controllers:KeyController"],
        beego.ControllerComments{
            Method: "Display",
            Router: `/display`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
