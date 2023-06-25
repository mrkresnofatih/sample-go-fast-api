package calculator

import (
	"github.com/labstack/echo/v4"
	GoFast "github.com/mrkresnofatih/gofast"
	"gitub.com/mrkresnofatih/gofastapi/models/calculator"
	"gitub.com/mrkresnofatih/gofastapi/services"
	"log"
	"net/http"
)

type CalculatorSubtractEndpoint struct {
	CalculatorService services.ICalculatorService
}

func (c *CalculatorSubtractEndpoint) GetHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		addReq := new(calculator.CalculatorSubtractRequest)
		_ = ctx.Bind(addReq)

		subtractResult, err := c.CalculatorService.Subtract(*addReq)
		if err != nil {
			log.Println("failed to add")
			return GoFast.SendBadResponse(ctx, "failed to add")
		}

		log.Println("subtract success")
		return GoFast.SendGoodResponse[calculator.CalculatorSubtractResponse](ctx, subtractResult)
	}
}

func (c *CalculatorSubtractEndpoint) GetMethod() string {
	return http.MethodPost
}

func (c *CalculatorSubtractEndpoint) GetPath() string {
	return "/subtract"
}

func (c *CalculatorSubtractEndpoint) Register(group *echo.Group) {
	group.Match([]string{c.GetMethod()}, c.GetPath(), c.GetHandler())
}
