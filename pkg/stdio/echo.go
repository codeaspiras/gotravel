package stdio

import "fmt"

func Echo(template string, args ...interface{}) {
	fmt.Printf(template+"\n", args...)
}
