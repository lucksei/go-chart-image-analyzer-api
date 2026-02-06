package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sync"
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

type HelmChartResult struct {
	RepoURL   string        `json:"repo_url"`
	ChartPath string        `json:"chart_path"`
	Images    []ImageResult `json:"images"`
}

type ResultStore struct {
	mu    sync.Mutex
	store map[string]HelmChartResult
}

func NewResultStore() *ResultStore {
	resultStore := &ResultStore{
		store: map[string]HelmChartResult{},
	}
	return resultStore
}

func (r *ResultStore) Get(key string) (HelmChartResult, bool) {
	r.mu.Lock()
	orig, ok := r.store[key]
	r.mu.Unlock()
	if !ok {
		return HelmChartResult{}, false
	}
	// TODO: Use a damn library for deep copying next time...
	newCopy := HelmChartResult{
		RepoURL:   orig.RepoURL,
		ChartPath: orig.ChartPath,
	}
	for _, img := range orig.Images {
		newCopy.Images = append(newCopy.Images, ImageResult{
			img.Name,
			img.Size,
			img.LayerNumber,
		})
	}
	return newCopy, true
}

func (r *ResultStore) Put(key string, v HelmChartResult) error {
	r.mu.Lock()

	newCopy := HelmChartResult{
		RepoURL:   v.RepoURL,
		ChartPath: v.ChartPath,
	}
	images := []ImageResult{}
	for _, img := range v.Images {
		images = append(images, ImageResult{
			Name:        img.Name,
			Size:        img.Size,
			LayerNumber: img.LayerNumber,
		})
	}
	newCopy.Images = images
	r.store[key] = newCopy

	r.mu.Unlock()
	return nil
}
