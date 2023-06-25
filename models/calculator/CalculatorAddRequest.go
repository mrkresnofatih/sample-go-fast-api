package calculator

type CalculatorAddRequest struct {
	VariableOne int `json:"variableOne" validate:"required"`
	VariableTwo int `json:"variableTwo" validate:"required"`
}
