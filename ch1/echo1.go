package main

import (
	"fmt"
	"os"
)

func main() {
	// 声明了俩个string类型的变量s和sep，如果没有被初始化，他将隐式地初始化为这个类型的空值
	var s, sep string
	// 循环的索引变量i在for循环开始处声明 := 用于短变量声明，可声明一个或多个
	// for initialization; condition ; post { .. }
	// for condition {} , for {.. break return }
	for i := 1; i < len(os.Args); i++ {
		// same to s = s + sep + os.Args[i]
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
