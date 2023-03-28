package phpserverclient

import "time"

const HttpReqTimeout = 20 * time.Second

const (
	Phone   = 0
	Email   = 1
	Address = 2
)

const (
	ParamSvcURL   = "serviceUrl"
	ParamCodeType = "codetype"
	ParamMobile   = "mobile"
	ParamEmail    = "email"
	ParamMsg      = "msg"
	ParamTicket   = "ticket"
	ParamBizId    = "biz"
	ParamCode     = "code"
)

const (
	Quick  = "quick"
	Bind   = "bind"
	Export = "export"
)
