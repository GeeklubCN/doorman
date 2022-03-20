package proxy

import (
	"bufio"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func NewGinHandler(realAddr []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// step 1: resolve proxy address, change scheme and host in request
		req := c.Request
		proxy, err := url.Parse(getLoadBalanceAddr(realAddr))
		if err != nil {
			log.Printf("error in parse addr: %v", err)
			c.String(500, "error")
			return
		}
		req.URL.Scheme = proxy.Scheme
		req.URL.Host = proxy.Host

		// step 2: use http.Transport to do request to real server.
		transport := http.DefaultTransport
		resp, err := transport.RoundTrip(req)
		if err != nil {
			log.Printf("error in roundtrip: %v", err)
			c.String(500, "error")
			return
		}

		// step 3: return real server response to upstream.
		for k, vv := range resp.Header {
			for _, v := range vv {
				c.Header(k, v)
			}
		}
		defer func() { _ = resp.Body.Close() }()
		_, _ = bufio.NewReader(resp.Body).WriteTo(c.Writer)
		c.Abort()
		return
	}
}

func getLoadBalanceAddr(realAddr []string) string {
	return realAddr[0]
}
