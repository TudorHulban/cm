package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddEntry(t *testing.T) {
	config1, errCr1 := NewTargetConfiguration("project1")
	require.NoError(t, errCr1)
	require.NotZero(t, config1)

	config1.AddEntry(&Entry{
		Name:  OSVariableName("var01"),
		Value: OSVariableValue("1234"),
	})

	config1.AddEntry(&Entry{
		Name:  OSVariableName("var02"),
		Value: OSVariableValue("abcd"),
	})

	_, errWr1 := config1.WriteToFile()
	require.NoError(t, errWr1)

	config2, errCr2 := NewTargetConfiguration("project2")
	require.NoError(t, errCr2)
	require.NotZero(t, config2)

	config2.AddEntry(&Entry{
		Name:  OSVariableName("var01"),
		Value: OSVariableValue("5678"),
	})

	config2.AddEntry(&Entry{
		Name:  OSVariableName("var02"),
		Value: OSVariableValue("efgh"),
	})

	_, errWr2 := config2.WriteToFile()
	require.NoError(t, errWr2)
}
