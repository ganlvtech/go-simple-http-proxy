package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

var (
	remoteHost      string
	listenHost      string
	enableHTTPS     bool
	certPath        string
	keyPath         string
	simpleHostProxy = httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = remoteHost
			req.Host = remoteHost
		},
	}
)

func init() {
	flag.StringVar(&remoteHost, "remote", "127.0.0.1", "Proxy remote host")
	flag.StringVar(&listenHost, "listen", ":8080", "Listening port")
	flag.BoolVar(&enableHTTPS, "https", false, "Enable HTTPS")
	flag.StringVar(&certPath, "cert", "server.crt", "HTTPS cert path")
	flag.StringVar(&keyPath, "key", "server.key", "HTTPS key path")
}

func main() {
	flag.Parse()

	r := gin.Default()
	r.Any("/*action", func(ctx *gin.Context) {
		simpleHostProxy.ServeHTTP(ctx.Writer, ctx.Request)
	})

	var err error
	if enableHTTPS {
		err = r.RunTLS(listenHost, certPath, keyPath)
	} else {
		err = r.Run(listenHost)
	}

	if err != nil {
		log.Fatal(err)
	}
}
