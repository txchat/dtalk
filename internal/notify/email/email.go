package email

import (
	"github.com/txchat/dtalk/internal/notify"
	"github.com/txchat/dtalk/internal/notify/phpserverclient"
)

type Email struct {
	exec *phpserverclient.Client
}

func NewEmail(url, appKey, secretKey, msg string) *Email {
	return &Email{
		exec: phpserverclient.NewClient(url, appKey, secretKey, msg),
	}
}

func (e *Email) Send(param map[string]string) (interface{}, error) {
	return e.exec.SendEmail(&phpserverclient.Config{
		Email:      param[notify.Account],
		Ticket:     param[phpserverclient.ParamTicket],
		BusinessId: param[phpserverclient.ParamBizId],
		CodeType:   param[phpserverclient.ParamCodeType],
	})
}

func (e *Email) ValidateCode(param map[string]string) error {
	return e.exec.ValidateEmail(&phpserverclient.Config{
		Email:    param[notify.Account],
		CodeType: param[phpserverclient.ParamCodeType],
		Code:     param[notify.Code],
	})
}
