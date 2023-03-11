package domain

type OSVariableName string
type OSVariableValue string

type Entry struct {
	Name  OSVariableName
	Value OSVariableValue
}
