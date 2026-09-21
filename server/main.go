package main

import (
	"time"
)

var startTime = time.Now()

func main() {
	go startGrpc()
	startHttp()
}
