package notify

const (
	Debug   = "debug"
	Release = "release"
)

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

type Config struct {
	Surl      string
	AppKey    string
	SecretKey string
	Msg       string
	Env       string
	CodeTypes map[string]string
}

type Whitelist struct {
	Account string
	Code    string
	Enable  bool
}

type Validate interface {
	Send(map[string]string) (interface{}, error)
	ValidateCode(map[string]string) error
}
