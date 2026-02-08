package utils

type ImageAnalysis struct {
	Name           string `json:"name"`
	Size           int64  `json:"size"`
	NumberOfLayers int    `json:"no_of_layers"`
}

type HelmChartSource struct {
	RepoURL  string `json:"repo_url"`
	ChartRef string `json:"chart_ref" binding:"required"` // TODO: Add to documentation later on. For remote, Can be a .tgz URI, a oci:// URI, or if RepoURL is set, a chart name from the repo specified
}

type HelmChartAnalysis struct {
	RepoURL  string          `json:"repo_url"`
	ChartRef string          `json:"chart_path"`
	Images   []ImageAnalysis `json:"images"`
}
