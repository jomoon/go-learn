package simp

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Getstr() string {
	filename := os.Args[1]
	if filename != "" {
		filename = "/Users/zain/test.txt"
	}
	data,err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr,"%v",err)

	}
	result := ""
	for _,line := range strings.Split(string(data),"\n") {
		result += "("+strings.TrimSpace(line)+",2,201,188,12,'',0,0,'2018-11-07 14:27:11','','','2018-11-07 14:27:11','hjj','hjj'),"
	}
	result = strings.TrimRight(result,",")
	return result
}