package debug

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/txchat/dtalk/service/backup/model"
)

type DebugValidate struct {
	mockCode string
	real     model.Validate
}

func NewDebugValidate(code string, v model.Validate) *DebugValidate {
	return &DebugValidate{
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

func (v *DebugValidate) Send(params map[string]string) (interface{}, error) {
	if v.real == nil {
		return nil, errors.New("未注册验证器")
	}
	return v.real.Send(params)
}

func (v *DebugValidate) ValidateCode(param map[string]string) error {
	code := param[model.ParamCode]
	if v.mockCode == code {
		return nil
	}
	if v.real == nil {
		return errors.New("未注册验证器")
	}
	return v.real.ValidateCode(param)
}
