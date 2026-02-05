package models

type HelmChartRequest struct {
	RepoURL   string `json:"repo_url"`
	ChartPath string `json:"chart_path" binding:"required"` // TODO: Add to documentation later on. For remote, Can be a .tgz URI, a oci:// URI, or if RepoURL is set, a chart name from the repo specified
}

type HelmChartResponse struct {
	RepoURL   string       `json:"repo_url"`
	ChartPath string       `json:"chart_path"`
	Images    []ChartImage `json:"images"`
}

type ChartImage struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}
