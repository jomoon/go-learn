package echo

import (
	"fmt"
	"os"
	"strings"
)

func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
