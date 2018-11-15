package fetch

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func FetchAll() {
	start := time.Now()
	// 创建字符串通道
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // 从通道中接受
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 不要泄漏资源

	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		return
	}

	secs := time.Since(start).Seconds
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
