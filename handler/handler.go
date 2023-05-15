package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func HandleGetIP(hs []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleGetIP(w, r, hs, false)
	}
}

func HandleDebug(rw http.ResponseWriter, r *http.Request) {
	if err := r.Header.Write(rw); err != nil {
		log.Printf("write debug msg fail, err:%v", err)
	}
}

func HandleGetIPJson(hs []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleGetIP(w, r, hs, true)
	}
}

func findIP(headers []string, r *http.Request) string {
	var ip string
	for _, h := range headers {
		ip = r.Header.Get(h)
		if len(ip) > 0 {
			break
		}
	}
	idx := strings.Index(ip, ",")
	if idx < 0 {
		return ip
	}
	return ip[0:idx]
}

func handleGetIP(w http.ResponseWriter, r *http.Request, headers []string, useJson bool) {
	ip := findIP(headers, r)
	render := renderText
	if useJson {
		render = renderJson
	}
	render(w, ip)
}

func renderText(w http.ResponseWriter, ip string) {
	_, _ = w.Write([]byte(ip))
}

func renderJson(w http.ResponseWriter, ip string) {
	m := &IPModel{
		IP: ip,
	}
	raw, _ := json.Marshal(m)
	_, _ = w.Write(raw)
}
