package controller

import (
	"encoding/json"
	"strconv"
	"test/app/services"
)

func (co *ControllerWeb) GetCurrentTarget() ([]byte, error) {
	reconstructed := co.ServiceMain.GetCurrentTarget()

	res := make(map[string]any, 2)
	res["success"] = true
	res["current-target"] = reconstructed

	return json.Marshal(res)
}

func (co *ControllerWeb) GetTargets() ([]byte, error) {
	reconstructed, currentTarget := co.ServiceMain.GetTargets()

	res := make(map[string]any, 2)
	res["success"] = true

	targets := make(map[string]string, len(reconstructed))

	for ix, target := range reconstructed {
		if target == currentTarget {
			targets[strconv.Itoa(ix+1)+"*"] = string(target)

			continue
		}

		targets[strconv.Itoa(ix+1)] = string(target)
	}

	res["targets"] = targets

	return json.Marshal(res)
}

func (co *ControllerWeb) GetVariableValues(params *services.ParamsGetVariableValues) ([]byte, error) {
	reconstructed, errFind := co.ServiceMain.GetVariableValues(params)
	if errFind != nil {
		return nil, errFind
	}

	res := make(map[string]any, 2)
	res["success"] = true
	res["targets"] = reconstructed

	return json.Marshal(res)
}

func (co *ControllerWeb) GetTargetConfiguration(params *services.ParamsFindTargetConfiguration) ([]byte, error) {
	reconstructed, errFind := co.ServiceMain.GetTargetConfiguration(params)
	if errFind != nil {
		return nil, errFind
	}

	res := make(map[string]any, 2)
	res["success"] = true

	mapReconstructed := make(map[string]any, 2)
	mapReconstructed["targetID"] = reconstructed.TargetID

	values := make(map[string]string, len(reconstructed.Entries))

	for vname, vvalue := range reconstructed.Entries {
		values[string(vname)] = string(vvalue)
	}

	mapReconstructed["entries"] = values

	res["target"] = mapReconstructed

	return json.Marshal(res)
}

// se https://stackoverflow.com/questions/70634715/golang-map-to-json-but-preserve-the-key-order
func (co *ControllerWeb) GetTargetConfigurationWSlice(params *services.ParamsFindTargetConfiguration) ([]byte, error) {
	reconstructed, errFind := co.ServiceMain.GetTargetConfiguration(params)
	if errFind != nil {
		return nil, errFind
	}

	result := []byte("{")

	jSuccess, _ := json.Marshal("success")
	result = append(result, jSuccess...)

	result = append(result, ":"...)

	jTrue, _ := json.Marshal(true)
	result = append(result, jTrue...)
	result = append(result, ","...)

	jTarget, _ := json.Marshal("target")
	result = append(result, jTarget...)
	result = append(result, ":"...)

	result = append(result, "{"...)

	jTargetID, _ := json.Marshal("targetID")

	result = append(result, jTargetID...)
	result = append(result, ":"...)

	jTargetValue, _ := json.Marshal(reconstructed.TargetID)

	result = append(result, jTargetValue...)
	result = append(result, ","...)

	jEntries, _ := json.Marshal("entries")

	result = append(result, jEntries...)
	result = append(result, ":"...)

	values := make(map[string]string, len(reconstructed.Entries))

	for vname, vvalue := range reconstructed.Entries {
		values[string(vname)] = string(vvalue)
	}

	jValues, _ := json.Marshal(values)
	result = append(result, jValues...)
	result = append(result, "}"...)
	result = append(result, "}"...)

	return result, nil
}
