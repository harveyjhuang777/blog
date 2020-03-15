package main

import (
	//_ "net/http/pprof"

	"github.com/jwjhuang/blog/service"
)

func main() {
	//runtime.SetBlockProfileRate(1)
	//runtime.SetMutexProfileFraction(1)
	//go http.ListenAndServe(":6060", nil) // for debug
	service.Run(service.ServeBlog)
}
