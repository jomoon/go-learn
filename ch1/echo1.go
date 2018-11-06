package main

import (
	"fmt"
	"os"
)

func main() {
	// 声明了俩个string类型的变量s和sep，如果没有被初始化，他将隐式地初始化为这个类型的空值
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
