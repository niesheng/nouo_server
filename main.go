package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// http.Handle("/", httpServer(Config_.Static))
	// if Config_.Ssl {
	// 	http.ListenAndServeTLS(":"+Config_.Port, Config_.Cert, Config_.Key, nil)
	// } else {
	// 	http.ListenAndServe(":"+Config_.Port, nil)
	// }
}
