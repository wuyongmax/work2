package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"log"
//	"flag"
//	"github.com/golang/glog"
)

func main() {
//	flag.Set("v",4)
//	glog.V(2).Info("in main")
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/GetComputeInfo", GetComputeInfo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("server error 500")
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200 ok")

}

func GetComputeInfo(w http.ResponseWriter, r *http.Request) {
	//手动构造的VERSION
	version := os.Getenv("VERSION")
	addr := r.RemoteAddr
	ar := strings.Split(addr, ":")
	r.Header.Get("")
	ip := ar[0]
	// go func(){}()
	fmt.Println("")
	port := ar[1]
	fmt.Println(fmt.Sprintf("client ip is %s port %s", ip, port))
	for k, v := range r.Header {
		w.Header().Set(k, v[0])

	}
	w.Header().Set("version", version)

}
