package main

import (
	"test/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	p1, errCr := domain.NewTargetConfiguration("project1")
	require.NoError(t, errCr)

	content, errRead := _fs.ReadFile("configs/config-project1.json")
	require.NoError(t, errRead)

	t.Log(string(content))

	_, errCo := p1.Read(content)
	require.NoError(t, errCo)

	t.Log(p1)

	reconstructedValue, errGet := p1.GetValue(domain.OSVariableName("var01"))
	require.NoError(t, errGet)
	require.Equal(t, domain.OSVariableValue("1234"), reconstructedValue)
}
