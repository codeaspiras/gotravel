package stdio

type IO interface {
	Ask(template string, arguments ...interface{}) (string, error)
	AskFloat(template string, arguments ...interface{}) (float64, error)
	Echo(template string, arguments ...interface{})
	End()
}

func New() IO {
	return &io{}
}

type io struct{}
