package domain

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func TestAddProject(t *testing.T) {
	configMain := newConfiguration()
	require.NotZero(t, configMain)
	require.Zero(t, configMain.Targets)

	configProjectA, errCr1 := NewTargetConfiguration("projectA")
	require.NoError(t, errCr1)
	require.NotZero(t, configProjectA)

	configProjectA.AddEntry(&Entry{
		Name:  OSVariableName("var01"),
		Value: OSVariableValue("1234"),
	})

	configProjectA.AddEntry(&Entry{
		Name:  OSVariableName("var02"),
		Value: OSVariableValue("abcd"),
	})

	require.NoError(t, configMain.AddProject(configProjectA))

	t.Log("targets", configMain.Targets)
	t.Log("data", configMain.Data)

	require.Len(t, configMain.Targets, 1)

	configMain.WriteTo(os.Stdout)

	configProjectB, errCr2 := NewTargetConfiguration("projectB")
	require.NoError(t, errCr2)
	require.NotZero(t, configProjectB)

	configProjectB.AddEntry(&Entry{
		Name:  OSVariableName("var01"),
		Value: OSVariableValue("5678"),
	})

	configProjectB.AddEntry(&Entry{
		Name:  OSVariableName("var02"),
		Value: OSVariableValue("efgh"),
	})

	configProjectB.AddEntry(&Entry{
		Name:  OSVariableName("var03"),
		Value: OSVariableValue("xyz"),
	})

	require.NoError(t, configMain.AddProject(configProjectB))
	require.Len(t, configMain.Targets, 2)

	configMain.WriteTo(os.Stdout)
}

func TestGetProject(t *testing.T) {
	configMain := newConfiguration()
	require.NotZero(t, configMain)
	require.Zero(t, configMain.Targets)

	configProjectA, errCr1 := NewTargetConfiguration("projectA")
	require.NoError(t, errCr1)
	require.NotZero(t, configProjectA)

	configProjectA.AddEntry(&Entry{
		Name:  OSVariableName("var01"),
		Value: OSVariableValue("1234"),
	})

	configProjectA.AddEntry(&Entry{
		Name:  OSVariableName("var02"),
		Value: OSVariableValue("abcd"),
	})

	require.NoError(t, configMain.AddProject(configProjectA))

	reconstructedProjectConfig, errGet := configMain.FindTargetConfiguration(string(configProjectA.TargetID))
	require.NoError(t, errGet)
	require.Zero(t, deep.Equal(configProjectA, reconstructedProjectConfig))

	reconstructedProjectConfig.WriteTo(os.Stdout)

	fmt.Println("-----------")
	fmt.Println(configMain.GetVariables())
}
