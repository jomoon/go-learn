package echo

import (
	"fmt"
	"os"
)

func Practice1() {
	// os.Args[0]是命令的名字
	// 得出结果是  /var/folders/pw/2x750brs2z58vvqh6tqls99w0000gn/T/go-build393628209/b001/exe/main
	fmt.Println(os.Args[0])
}
