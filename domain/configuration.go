package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"sync"
	"test/app/apperrors"
)

type TargetValue struct {
	TargetID TargetID
	Value    OSVariableValue
}

type Configuration struct {
	// slice chosen as map returns random order
	Targets       []TargetID
	CurrentTarget TargetID

	Data map[OSVariableName]map[TargetID]OSVariableValue `json:"configuration"`

	configFiles map[string][]byte
	sync.Mutex
}

func newConfiguration() *Configuration {
	return &Configuration{
		Data:        make(map[OSVariableName]map[TargetID]OSVariableValue),
		configFiles: make(map[string][]byte),
	}
}

func NewTargets(targets ...string) []TargetID {
	res := make([]TargetID, len(targets))

	for ix, target := range targets {
		res[ix] = TargetID(target)
	}

	return res
}

func NewConfigurationFrom(configFiles map[string][]byte) (*Configuration, error) {
	if len(configFiles) == 0 {
		return nil, errors.New("no config files contents passed")
	}

	res := Configuration{
		configFiles: configFiles,
	}

	if errInit := res.init(); errInit != nil {
		return nil, errInit
	}

	return &res, nil
}

func (cfg *Configuration) init() error {
	tempConfiguration := Configuration{
		Data: make(map[OSVariableName]map[TargetID]OSVariableValue),
	}

	for target, contents := range cfg.configFiles {
		bufProject := TargetConfiguration{
			TargetID: TargetID(target),
		}

		if _, errPro := bufProject.Read(contents); errPro != nil {
			return errPro
		}

		if errAdd := tempConfiguration.AddProject(&bufProject); errAdd != nil {
			return apperrors.ErrDomain{
				Caller:     "init",
				NameMethod: "AddProject",
				Issue:      errAdd,
			}
		}
	}

	cfg.Lock()
	defer cfg.Unlock()

	cfg.CurrentTarget = tempConfiguration.Targets[0]
	cfg.Targets = tempConfiguration.Targets
	cfg.Data = tempConfiguration.Data

	return nil
}

func (cfg *Configuration) WriteTo(w io.Writer) (int, error) {
	data, errMa := json.MarshalIndent(cfg, "", " ")
	if errMa != nil {
		return 0, errMa
	}

	return w.Write(data)
}

// HasTarget has locking.
func (cfg *Configuration) HasTarget(target TargetID) bool {
	cfg.Lock()
	defer cfg.Unlock()

	for ix := range cfg.Targets {
		if cfg.Targets[ix] == target {
			return true
		}
	}

	return false
}

func (cfg *Configuration) AddProject(project *TargetConfiguration) error {
	if cfg.HasTarget(project.TargetID) {
		return fmt.Errorf("configuration already has target '%s' data", project.TargetID)
	}

	temp := cfg.Data

	for vname, vvalue := range project.Entries {
		if projectValues, exists := temp[vname]; exists {
			if _, valueExists := projectValues[project.TargetID]; valueExists {
				return apperrors.ErrValidation{
					Caller: "AddProject",
					Issue:  fmt.Errorf("for entry:'%s' entry already exists for target:%s", vname, project.TargetID),
				}
			}
		}

		if temp[vname] == nil {
			temp[vname] = make(map[TargetID]OSVariableValue)
		}

		temp[vname][project.TargetID] = vvalue
	}

	cfg.Lock()
	defer cfg.Unlock()

	cfg.Targets = append(cfg.Targets, project.TargetID)

	cfg.Data = temp

	return nil
}

// TODO: get vars for target
func (cfg *Configuration) GetVariables() []OSVariableName {
	var res []OSVariableName

	cfg.Lock()
	defer cfg.Unlock()

	for vname := range cfg.Data {
		res = append(res, vname)
	}

	sort.Sort(JustWords(res))

	return res
}

func (cfg *Configuration) CompareProjects(project1, project2 string) string {
	return ""
}
