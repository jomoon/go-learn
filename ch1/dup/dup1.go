package dup

import (
	"bufio"
	"fmt"
	"os"
)

func Dup1() {
	// map 存储一个键值对的集合
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 等价于 line := input.Text()
		// counts[line] = counts[line] + 1
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
