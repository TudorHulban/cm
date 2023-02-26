package domain

import (
	"fmt"
)

func (cfg *Configuration) FindTargetConfiguration(target string) (*TargetConfiguration, error) {
	if len(target) == 0 {
		target = string(cfg.CurrentTarget)
	} else {
		if !cfg.HasTarget(TargetID(target)) {
			return nil, fmt.Errorf("configuration does not have target '%s' data", target)
		}
	}

	configurationProject, errCr := NewTargetConfiguration(target)
	if errCr != nil {
		return nil, errCr
	}

	cfg.Lock()
	defer cfg.Unlock()

	for vname, projectValues := range cfg.Data {
		if vvalue, exists := projectValues[configurationProject.TargetID]; exists {
			configurationProject.Entries[vname] = vvalue
		}
	}

	if len(configurationProject.Entries) == 0 {
		return nil, fmt.Errorf("no entries were found for target:'%s'", target)
	}

	return configurationProject, nil
}

func (cfg *Configuration) FindVariableValues(vname string, targets ...string) ([]TargetValue, error) {
	cfg.Lock()
	defer cfg.Unlock()

	vvalues, exists := cfg.Data[OSVariableName(vname)]
	if !exists {
		return nil, fmt.Errorf("no values found for variable:'%s", vname)
	}

	var length int
	var over []TargetID

	if len(targets) > 0 {
		length = len(targets)
		over = NewTargets(targets...)
	} else {
		length = len(cfg.Targets)
		over = cfg.Targets
	}

	res := make([]TargetValue, length)

	for ix, target := range over {
		res[ix] = TargetValue{
			TargetID: target,
		}

		if value, exists := vvalues[target]; exists {
			res[ix].Value = value
		}
	}

	return res, nil
}
