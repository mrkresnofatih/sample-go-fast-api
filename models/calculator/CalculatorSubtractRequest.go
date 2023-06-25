package calculator

type CalculatorSubtractRequest struct {
	VariableOne int `json:"variableOne" validate:"required"`
	VariableTwo int `json:"variableTwo" validate:"required"`
}
