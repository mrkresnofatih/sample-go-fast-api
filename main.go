package main

import (
	"gitub.com/mrkresnofatih/gofastapi/controllers"
	"sync"
)

func main() {
	var appRunState sync.WaitGroup
	appRunState.Add(1)
	controllers.InitHttpServer(&appRunState)
	appRunState.Wait()
}
