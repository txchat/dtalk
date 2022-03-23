package sms

import (
	"github.com/txchat/dtalk/service/backup/model"
	"testing"
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

	sms := NewSMS(url, appkey, secretKey, msg)
	params := map[string]string{
		model.ParamMobile:   "15763946517",
		model.ParamCodeType: "bind_policebook",
	}
	rlt, err := sms.Send(params)
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

	sms := NewSMS(url, appkey, secretKey, msg)
	params := map[string]string{
		model.ParamMobile:   "15763946517",
		model.ParamCode:     "04037",
		model.ParamCodeType: "bind_policebook",
	}
	err := sms.ValidateCode(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success")
}
