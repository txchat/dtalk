package whitelist

import (
	"errors"

	"github.com/txchat/dtalk/service/backup/config"
	"github.com/txchat/dtalk/service/backup/model"
)

type WhitelistValidate struct {
	whitelist map[string]string
	real      model.Validate
}

func NewWhitelistValidate(list []config.SMSEmailWhitelist, v model.Validate) *WhitelistValidate {
	whitelist := make(map[string]string)
	for _, item := range list {
		if item.Enable {
			whitelist[item.Account] = item.Code
		}
	}
	return &WhitelistValidate{
		whitelist: whitelist,
		real:      v,
	}
}

func (v *WhitelistValidate) Send(params map[string]string) (interface{}, error) {
	if v.real == nil {
		return nil, errors.New("未注册验证器")
	}
	return v.real.Send(params)
}

func (v *WhitelistValidate) ValidateCode(param map[string]string) error {
	code := param[model.ParamCode]
	phone := param[model.ParamMobile]
	email := param[model.ParamEmail]
	acc := phone
	if phone == "" {
		acc = email
	}
	if c, ok := v.whitelist[acc]; ok && c == code {
		return nil
	}
	if v.real == nil {
		return errors.New("未注册验证器")
	}
	return v.real.ValidateCode(param)
}
