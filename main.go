package main

import (
	"flag"
	"log"
	"myip/handler"
	"net/http"
	"strings"
)

var bind = flag.String("bind", ":5578", "bind address")
var header = flag.String("use_headers", "X-real-IP", "specify ip header")

func main() {
	flag.Parse()

	headers := strings.Split(*header, ";")
	if len(headers) == 0 {
		panic("no valid headers found")
	}

	http.HandleFunc("/", handler.HandleGetIP(headers))
	http.HandleFunc("/debug", handler.HandleDebug)
	http.HandleFunc("/json", handler.HandleGetIPJson(headers))

	log.Printf("Server bind on %s", *bind)
	err := http.ListenAndServe(*bind, nil)
	if err != nil {
		log.Fatalf("Bind ipaddr fail, ip:%s, err:%v", *bind, err)
	}
}
