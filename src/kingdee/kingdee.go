package kingdee

import (
	"fmt"

	"github.com/deep-project/kingdee"
	"github.com/deep-project/kingdee/adapters"
)

/*
	{
	    "acctID": "675e79d7cb38d3",
	    "username": "订单同步接口用户",
	    "password":"",
	    "appid": "322694_31cO3aDIUqoex+ys1dQo561KRuXb6POo",
	    "appsecret": "83f36550c7c24dd0987f371c639ea085",
	    "lcid": 2052
	}
*/
func KingdeeGetInStock() error {
	// https://app.tolobio.com/k3cloud/Kingdee.BOS.WebApi.ServicesStub.AuthService.ValidateUser.common.kdsvc
	cli, err := kingdee.New("https://app.tolobio.com/K3Cloud/", &adapters.LoginBySign{
		AccountID:  "675e79d7cb38d3",
		Username:   "订单同步接口用户",
		AppID:      "322694_31cO3aDIUqoex+ys1dQo561KRuXb6POo",
		AppSecret:  "83f36550c7c24dd0987f371c639ea085",
		LanguageID: "2052",
	})
	if err != nil {
		return err
	}
	raw, err := cli.View("SAL_SaleOrder", map[string]any{"Number": "CGRK00019"})
	if err != nil {
		return err
	}
	fmt.Println(string(raw))
	return nil
}
