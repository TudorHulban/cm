package main

import (
	"os"
	"test/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConfiguration(t *testing.T) {
	configFiles, errLoad := loadConfigFiles()
	require.NoError(t, errLoad)
	require.NotZero(t, configFiles)
	require.Len(t, configFiles, 2)

	t.Log(configFiles)

	config, errCr := domain.NewConfigurationFrom(configFiles)
	require.NoError(t, errCr)
	require.NotZero(t, config)
	require.Len(t, config.Targets, 2)
	require.Len(t, config.Data, 2)

	config.WriteTo(os.Stdout)
}
