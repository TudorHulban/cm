package domain

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func (cfg *Configuration) DeleteTargetConfiguration(target string) error {
	if !cfg.HasTarget(TargetID(target)) {
		return fmt.Errorf("configuration does not have target '%s' data", target)
	}

	cfg.Lock()
	defer cfg.Unlock()

	index := slices.Index(cfg.Targets, TargetID(target))

	cfg.Targets = append(cfg.Targets[:index], cfg.Targets[index+1:]...)

	if cfg.CurrentTarget == TargetID(target) {
		if len(cfg.Targets) > 0 {
			cfg.CurrentTarget = cfg.Targets[0]
		} else {
			cfg.CurrentTarget = ""
		}
	}

	for _, projectValues := range cfg.Data {
		delete(projectValues, TargetID(target))
	}

	return nil
}
