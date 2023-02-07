package debug

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/txchat/dtalk/pkg/notify"
)

type Validate struct {
	mockCode string
	real     notify.Validate
}

func NewDebugValidate(code string, v notify.Validate) *Validate {
	return &Validate{
		mockCode: code,
		real:     v,
	}
}

func GetMockCode(mode string) string {
	defCode := "111111"
	reg := regexp.MustCompile(`^FzmRandom(\d+)$`)

	//找出子串
	ret := reg.FindStringSubmatch(mode)
	if len(ret) < 2 {
		return defCode
	}

	num, err := strconv.Atoi(ret[1])
	if err != nil || num == 0 {
		return defCode
	}
	return strings.Repeat("1", num)
}

func (v *Validate) Send(params map[string]string) (interface{}, error) {
	if v.real == nil {
		return nil, errors.New("未注册验证器")
	}
	return v.real.Send(params)
}

func (v *Validate) ValidateCode(param map[string]string) error {
	code := param[notify.ParamCode]
	if v.mockCode == code {
		return nil
	}
	if v.real == nil {
		return errors.New("未注册验证器")
	}
	return v.real.ValidateCode(param)
}
