package stdio

type IO interface {
	Echo(template string, arguments ...interface{})
	End()
}

func New() IO {
	return &io{}
}

type io struct{}
