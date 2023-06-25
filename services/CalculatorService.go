package services

import "gitub.com/mrkresnofatih/gofastapi/models/calculator"

type ICalculatorService interface {
	Add(addRequest calculator.CalculatorAddRequest) (calculator.CalculatorAddResponse, error)
	Subtract(subtractRequest calculator.CalculatorSubtractRequest) (calculator.CalculatorSubtractResponse, error)
}

type CalculatorService struct{}

func (c *CalculatorService) Add(addRequest calculator.CalculatorAddRequest) (calculator.CalculatorAddResponse, error) {
	result := addRequest.VariableOne + addRequest.VariableTwo
	return calculator.CalculatorAddResponse{
		Result: result,
	}, nil
}

func (c *CalculatorService) Subtract(subtractRequest calculator.CalculatorSubtractRequest) (calculator.CalculatorSubtractResponse, error) {
	result := subtractRequest.VariableOne - subtractRequest.VariableTwo
	return calculator.CalculatorSubtractResponse{
		Result: result,
	}, nil
}
