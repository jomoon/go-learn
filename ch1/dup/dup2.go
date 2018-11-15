package dup

import (
	"bufio"
	"fmt"
	"os"
)

// Dup2 这个方法是乱码的 ？？ 原因？？
func Dup2() {
	counts := make(map[string]int)
	files := os.Args
	fmt.Println(files)
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			fmt.Println(arg)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
