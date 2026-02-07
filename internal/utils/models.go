package utils

type ImageAnalysis struct {
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	LayerNumber int    `json:"layer_number"`
}

type HelmChartSource struct {
	RepoURL  string `json:"repo_url"`
	ChartRef string `json:"chart_ref"`
}

type HelmChartAnalysis struct {
	RepoURL  string          `json:"repo_url"`
	ChartRef string          `json:"chart_path"`
	Images   []ImageAnalysis `json:"images"`
}
