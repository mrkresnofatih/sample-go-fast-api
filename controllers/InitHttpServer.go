package controllers

import (
	GoFast "github.com/mrkresnofatih/gofast"
	"gitub.com/mrkresnofatih/gofastapi/services"
	"sync"
)

func InitHttpServer(appRunState *sync.WaitGroup) {
	go func() {
		httpServerObj := &GoFast.ApplicationServer{}

		calculatorController := &CalculatorController{
			CalculatorService: &services.CalculatorService{},
		}
		httpServerObj.AddController(calculatorController)

		httpServerObj.Initialize()
		httpServerObj.Router.Logger.Fatal(httpServerObj.Router.Start(":1323"))
		appRunState.Done()
	}()
}
