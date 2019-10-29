package main

import (
	"os"
	"runtime"

	"github.com/valyala/fasthttp"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if Config_.Server.Tls {
		_, err := os.Stat(Config_.Server.Cert)
		if err != nil {
			Exit(err) //power error
		}
		_, err = os.Stat(Config_.Server.Key)
		if err != nil {
			Exit(err) //power error
		}

		fasthttp.ListenAndServeTLS(":"+Config_.Server.Port, Config_.Server.Cert, Config_.Server.Key, router_handle)
	} else {
		fasthttp.ListenAndServe(":"+Config_.Server.Port, router_handle)
	}

	select {}
}
