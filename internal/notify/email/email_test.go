package email

import (
	"testing"

	"github.com/txchat/dtalk/internal/notify"
	"github.com/txchat/dtalk/internal/notify/phpserverclient"
)

/*
[smsConfig]
url="http://<host:port>/send/sms2"
codeType="chat_notice"
mobile=[""]
*/
func Test_Send(t *testing.T) {
	url := "http://118.31.52.32"
	appkey := "chat33pro"
	secretKey := "eQXXMphNFHQL4YeW"
	msg := "FzmRandom5"

	email := NewEmail(url, appkey, secretKey, msg)
	params := map[string]string{
		notify.Account:                "815904261@qq.com",
		phpserverclient.ParamCodeType: "bind_policebook",
	}
	rlt, err := email.Send(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success", "rlt", rlt)
}

func Test_ValidateCode(t *testing.T) {
	url := "http://118.31.52.32"
	appkey := "chat33pro"
	secretKey := "eQXXMphNFHQL4YeW"
	msg := "FzmRandom5"

	email := NewEmail(url, appkey, secretKey, msg)
	params := map[string]string{
		notify.Account:                "815904261@qq.com",
		notify.Code:                   "17091",
		phpserverclient.ParamCodeType: "bind_policebook",
	}
	err := email.ValidateCode(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success")
}
