package calculator

import (
	"github.com/labstack/echo/v4"
	GoFast "github.com/mrkresnofatih/gofast"
	"gitub.com/mrkresnofatih/gofastapi/models/calculator"
	"gitub.com/mrkresnofatih/gofastapi/services"
	"log"
	"net/http"
)

type CalculatorAddEndpoint struct {
	CalculatorService services.ICalculatorService
}

func (c *CalculatorAddEndpoint) GetHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		addReq := new(calculator.CalculatorAddRequest)
		_ = ctx.Bind(addReq)

		addResult, err := c.CalculatorService.Add(*addReq)
		if err != nil {
			log.Println("failed to add")
			return GoFast.SendBadResponse(ctx, "failed to add")
		}

		log.Println("add success")
		return GoFast.SendGoodResponse[calculator.CalculatorAddResponse](ctx, addResult)
	}
}

func (c *CalculatorAddEndpoint) GetMethod() string {
	return http.MethodPost
}

func (c *CalculatorAddEndpoint) GetPath() string {
	return "/add"
}

func (c *CalculatorAddEndpoint) Register(group *echo.Group) {
	group.Match([]string{c.GetMethod()}, c.GetPath(), c.GetHandler())
}
