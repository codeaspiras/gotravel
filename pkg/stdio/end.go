package stdio

import (
	"bufio"
	"fmt"
	"os"
)

func (i *io) End() {
	fmt.Printf("# FIM - Pressione ENTER para encerrar o programa ~ ")
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
}
