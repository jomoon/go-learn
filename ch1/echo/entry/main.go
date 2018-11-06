package main

import (
	"fmt"
	"time"

	"hjj.com/study/ch1/echo"
)

func main() {
	start := time.Now()
	echo.Echo1()
	end := time.Now()
	fmt.Println(end.Sub(start))

	start = time.Now()
	echo.Echo2()
	end = time.Now()
	fmt.Println(end.Sub(start))

	start = time.Now()
	echo.Echo3()
	end = time.Now()
	fmt.Println(end.Sub(start))

	// echo.Practice1()
	// echo.Practice2()
}
