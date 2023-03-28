package sms

import (
	"github.com/txchat/dtalk/internal/notify"
	"github.com/txchat/dtalk/internal/notify/phpserverclient"
)

type SMS struct {
	exec *phpserverclient.Client
}

func NewSMS(url, appKey, secretKey, msg string) *SMS {
	return &SMS{
		exec: phpserverclient.NewClient(url, appKey, secretKey, msg),
	}
}

func (s *SMS) Send(param map[string]string) (interface{}, error) {
	return s.exec.SendSMS(&phpserverclient.Config{
		Phone:      param[notify.Account],
		Ticket:     param[phpserverclient.ParamTicket],
		BusinessId: param[phpserverclient.ParamBizId],
		CodeType:   param[phpserverclient.ParamCodeType],
	})
}

func (s *SMS) ValidateCode(param map[string]string) error {
	return s.exec.ValidateSMS(&phpserverclient.Config{
		Phone:    param[notify.Account],
		CodeType: param[phpserverclient.ParamCodeType],
		Code:     param[notify.Code],
	})
}
