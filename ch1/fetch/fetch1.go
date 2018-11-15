package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Fetch1() {
	for _, url := range os.Args[1:] {
		httpPre := "http://"
		hasHTTPPre := strings.HasPrefix(url, httpPre)
		if !hasHTTPPre {
			url = httpPre + url
		}
		resp, err := http.Get(url)

		if err != nil {
			// %v 内置格式的任意值
			fmt.Fprintf(os.Stderr, "Fetch:%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		fmt.Println(resp.Status)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
