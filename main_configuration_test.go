package main

import (
	"os"
	"test/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConfiguration(t *testing.T) {
	fsEntries, errRead := loadFSEntries()
	require.NoError(t, errRead)

	configFiles, errLoad := loadVarsFiles(fsEntries)
	require.NoError(t, errLoad)
	require.NotZero(t, configFiles)
	require.Len(t, configFiles, 3)

	t.Log(configFiles)

	config, errCr := domain.NewConfigurationFrom(configFiles)
	require.NoError(t, errCr)
	require.NotZero(t, config)
	require.Len(t, config.Targets, 3)
	require.Len(t, config.Data, 3)

	config.WriteTo(os.Stdout)
}
