package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex
var count int

func Server2() {
	http.HandleFunc("/",handler2)
	http.HandleFunc("/count",counter)
	http.HandleFunc("/server",handler3)
	http.HandleFunc("/lissajous",lissHandler)

	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler2(w http.ResponseWriter,r *http.Request) {
	mu.Lock()
	count ++
	mu.Unlock()
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
}

func getRequestParam(r * http.Request,query string) string {
	split := strings.Split(r.URL.RawQuery, "&")
	for _,e := range split {
		contains := strings.Contains(e, query)
		if contains {
			return strings.Split(e,"=")[1]
		}
	}
	return ""
}


func lissHandler(w http.ResponseWriter,r *http.Request) {
	param := getRequestParam(r, "cycles")
	cycles,err := strconv.Atoi(param)
	if err != nil {
		lissajous(w)
	}
	lissajousCir(w,float64(cycles))
}

func handler3(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"%s %s %s\n",r.Method,r.URL,r.Proto)

	for k,v:=range r.Header  {
		fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
	}

	fmt.Fprintf(w,"Host = %q\n",r.Host)
	fmt.Fprintf(w,"RemoteAddr = %q\n",r.RemoteAddr)

	if err := r.ParseForm();err != nil {
		log.Print(err)
	}

	for k,v := range r.Form {
		fmt.Fprintf(w,"Form[%q] = %q\n",k,v)
	}
}

func counter(w http.ResponseWriter,r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w,"Count %d\n",count)
	mu.Unlock()
}