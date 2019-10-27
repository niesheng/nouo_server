package main

import (
	"os"

	"github.com/valyala/fasthttp"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())

	if Config_.Ssl {
		_, err := os.Stat(Config_.Cert)
		if err != nil {
			Exit(err) //power error
		}
		_, err = os.Stat(Config_.Key)
		if err != nil {
			Exit(err) //power error
		}

		fasthttp.ListenAndServeTLS(":"+Config_.Port, Config_.Cert, Config_.Key, fastHTTPHandler)
	} else {
		fasthttp.ListenAndServe(":"+Config_.Port, fastHTTPHandler)
	}
	select {}
}
