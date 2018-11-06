package echo

import (
	"fmt"
	"os"
)

func Practice2() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, " -- ", arg)
	}
}
