package stdio

type IO interface {
	Ask(template string, arguments ...interface{}) (string, error)
	Echo(template string, arguments ...interface{})
}

func New() IO {
	return &io{}
}

type io struct{}
