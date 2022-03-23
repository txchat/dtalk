package model

type AppConfig struct {
	BasicConfig BasicConfig `yaml:"basic_config"`
	Request     Request     `yaml:"request"`
	Response    Response    `yaml:"response"`
}

type BasicConfig struct {
	Method string `yaml:"method"`
	Url    string `yaml:"url"`
}

type Request struct {
	TokenName string `yaml:"token_name"`
	//FieldsOfHeader []string                    `yaml:"fields_of_header"`
	//FieldsOfBody map[interface{}]interface{} `yaml:"fields_of_body"`
}
type Response struct {
	//ContentType   string                      `yaml:"content_type"`
	UidName       string                      `yaml:"uid_name"`
	SuccessResult map[interface{}]interface{} `yaml:"success_result"`
}
