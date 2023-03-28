package notify

const (
	Debug   = "debug"
	Release = "release"
)

const (
	Account = "account"
	Code    = "code"
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
