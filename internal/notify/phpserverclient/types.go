package phpserverclient

type SendResult struct {
	IsShow     int
	IsValidate int
	Data       map[string]interface{}
}

type Error struct {
	Message string
	Err     string
	Code    string
}

func (e *Error) Error() string {
	return e.Message
}
