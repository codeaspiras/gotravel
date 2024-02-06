package stdio

import "fmt"

func (*io) Echo(template string, args ...interface{}) {
	fmt.Printf(template+"\n", args...)
}
