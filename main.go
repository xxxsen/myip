package main

import (
	"flag"
	"log"
	"net/http"
)

var bind = flag.String("bind", ":5578", "bind address")

func getip(rw http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-real-IP")
	log.Printf("Recv request from ip:%s", ip)
	rw.Write([]byte(ip))
}

func main() {
	flag.Parse()
	http.HandleFunc("/", getip)
	log.Printf("Server bind on %s", *bind)
	err := http.ListenAndServe(*bind, nil)
	if err != nil {
		log.Fatalf("Bind ipaddr fail, ip:%s, err:%v", *bind, err)
	}
}
