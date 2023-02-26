package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"test/app/apperrors"
)

// TODO: add a service area

type TargetConfiguration struct {
	Entries  map[OSVariableName]OSVariableValue
	TargetID TargetID
}

func NewTargetConfiguration(target string) (*TargetConfiguration, error) {
	if len(target) == 0 {
		return nil, errors.New("passed project ID is empty")
	}

	return &TargetConfiguration{
		TargetID: TargetID(target),
		Entries:  make(map[OSVariableName]OSVariableValue),
	}, nil
}

func (cfg *TargetConfiguration) Read(p []byte) (int, error) {
	var temp TargetConfiguration

	if errUn := json.Unmarshal(p, &temp); errUn != nil {
		return 0, errUn
	}

	*cfg = temp

	return len(p), nil
}

func (cfg *TargetConfiguration) getBackupFileName(folder, prefix string) string {
	return folder + "/" + prefix + "-" + string(cfg.TargetID) + ".json"
}

func (cfg *TargetConfiguration) WriteToFile() (int, error) {
	f, errCr := os.Create(cfg.getBackupFileName(_folderBackupFile, _prefixBackupFile))
	if errCr != nil {
		return 0, errCr
	}
	defer f.Close()

	return cfg.WriteTo(f)
}

func (cfg *TargetConfiguration) WriteTo(w io.Writer) (int, error) {
	data, errMa := json.MarshalIndent(cfg, "", " ")
	if errMa != nil {
		return 0, errMa
	}

	return w.Write(data)
}

func (cfg *TargetConfiguration) AddEntry(entry *Entry) error {
	if _, exists := cfg.Entries[entry.Name]; exists {
		return fmt.Errorf("entry '%s' already exists", entry.Name)
	}

	cfg.Entries[entry.Name] = entry.Value

	return nil
}

func (cfg *TargetConfiguration) GetValue(vname OSVariableName) (OSVariableValue, error) {
	if vvalue, exists := cfg.Entries[vname]; exists {
		return vvalue, nil
	}

	return OSVariableValue(""), apperrors.ErrRecordNotFound{}
}
