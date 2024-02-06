package stdio

import (
	"bufio"
	"fmt"
	"os"
)

func (*io) Ask(template string, args ...interface{}) (string, error) {
	fmt.Printf(template, args...)
	buf := bufio.NewScanner(os.Stdin)
	if buf.Scan() {
		return buf.Text(), nil
	}

	if err := buf.Err(); err != nil {
		return "", err
	}

	return "", nil
}
