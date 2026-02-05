package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type HelmChartPath struct {
	RepoURL   string `json:"repo_url"`
	ChartPath string `json:"chart_path"`
}

func (h HelmChartPath) ToBase64Id() (string, error) {
	jsonDataBytes, err := json.Marshal(h)
	if err != nil {
		return "", err
	}
	str := base64.StdEncoding.EncodeToString(jsonDataBytes)
	return str, nil
}

func Base64StringToHelmChart(base64Data string) (HelmChartPath, error) {
	dataBytes, err := base64.StdEncoding.DecodeString(base64Data)
	fmt.Printf("%s\n", dataBytes)
	if err != nil {
		return HelmChartPath{}, err
	}
	var h HelmChartPath
	err = json.Unmarshal(dataBytes, &h)
	fmt.Printf("%v\n", h)
	if err != nil {
		return HelmChartPath{}, err
	}
	return h, nil
}
