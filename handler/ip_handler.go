package handler

import (
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type IPHandler struct {
	headerlist []string
}

func NewIPHandler(headerlist []string) *IPHandler {
	return &IPHandler{
		headerlist: headerlist,
	}
}

func (h *IPHandler) HandleGetIP(c *gin.Context) {
	ip := h.findIP(c.Request)
	if h.isAllowJsonResponse(c.Request) {
		c.JSON(http.StatusOK, gin.H{"ip": ip})
		return
	}
	c.String(http.StatusOK, ip)
}

func (h *IPHandler) isAllowJsonResponse(r *http.Request) bool {
	accept := r.Header.Get("Accept")
	if len(accept) == 0 {
		return false
	}
	return strings.Contains(accept, "application/json")
}

func (h *IPHandler) findIP(r *http.Request) string {
	var ip string
	for _, h := range h.headerlist {
		ip = r.Header.Get(h)
		if len(ip) > 0 {
			break
		}
	}
	idx := strings.Index(ip, ",")
	if idx > 0 {
		ip = ip[0:idx]
	}
	ip = strings.TrimSpace(ip)
	if len(ip) == 0 { //header都找不到, 那么就从 RemoteAddr 中获取
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}
