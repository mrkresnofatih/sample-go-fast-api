package controllers

import (
	"github.com/labstack/echo/v4"
	GoFast "github.com/mrkresnofatih/gofast"
	"gitub.com/mrkresnofatih/gofastapi/endpoints/calculator"
	calculator2 "gitub.com/mrkresnofatih/gofastapi/models/calculator"
	"gitub.com/mrkresnofatih/gofastapi/services"
)

type CalculatorController struct {
	CalculatorService services.ICalculatorService
}

func (c CalculatorController) Register(echo *echo.Echo) {
	controllerRouter := GoFast.ControllerRouter{
		MainRouter: echo,
		PathPrefix: "/calculator",
	}

	addEndpoint := &calculator.CalculatorAddEndpoint{
		CalculatorService: c.CalculatorService,
	}
	addWithValidation := &GoFast.RequireValidationDecorator[calculator2.CalculatorAddRequest]{
		Endpoint: addEndpoint,
	}
	controllerRouter.AddEndpoint(addWithValidation)

	subtractEndpoint := &calculator.CalculatorSubtractEndpoint{
		CalculatorService: c.CalculatorService,
	}
	subtractWithValidation := &GoFast.RequireValidationDecorator[calculator2.CalculatorSubtractRequest]{
		Endpoint: subtractEndpoint,
	}
	controllerRouter.AddEndpoint(subtractWithValidation)

	controllerRouter.Build()
}
