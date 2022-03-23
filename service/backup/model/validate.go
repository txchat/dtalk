package model

type Validate interface {
	Send(map[string]string) (interface{}, error)
	ValidateCode(map[string]string) error
}
