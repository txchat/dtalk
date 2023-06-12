package whitelist

import (
	"errors"

	"github.com/txchat/dtalk/internal/notify"
)

type Validate struct {
	whitelist map[string]string
	real      notify.Validate
}

func NewWhitelistValidate(list []notify.Whitelist, v notify.Validate) *Validate {
	whitelist := make(map[string]string)
	for _, item := range list {
		if item.Enable {
			whitelist[item.Account] = item.Code
		}
	}
	return &Validate{
		whitelist: whitelist,
		real:      v,
	}
}

func (v *Validate) Send(params map[string]string) (interface{}, error) {
	if v.real == nil {
		return nil, errors.New("未注册验证器")
	}
	return v.real.Send(params)
}

func (v *Validate) ValidateCode(param map[string]string) error {
	code := param[notify.Code]
	acc := param[notify.Account]
	if c, ok := v.whitelist[acc]; ok && c == code {
		return nil
	}
	if v.real == nil {
		return errors.New("未注册验证器")
	}
	return v.real.ValidateCode(param)
}
