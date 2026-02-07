package utils

import (
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/engine"
	"helm.sh/helm/v3/pkg/registry"
)

var Settings *cli.EnvSettings
var Client *action.Install

func InitHelmSettings() error {
	Settings = cli.New()
	actionConfig := &action.Configuration{}
	if err := actionConfig.Init(
		Settings.RESTClientGetter(),
		Settings.Namespace(),
		"",
		nil,
	); err != nil {
		return err
	}

	if actionConfig.RegistryClient == nil {
		regClient, err := registry.NewClient(
			registry.ClientOptWriter(os.Stderr),
		)
		if err != nil {
			return err
		}
		actionConfig.RegistryClient = regClient
	}

	Client = action.NewInstall(actionConfig)

	Client.DryRun = true
	Client.ClientOnly = true

	return nil
}

func RenderHelmTemplate(helmChartSource HelmChartSource) (map[string]string, error) {
	Client.ChartPathOptions.RepoURL = helmChartSource.RepoURL

	chartPath, err := Client.LocateChart(helmChartSource.ChartRef, Settings)
	if err != nil {
		return nil, err
	}

	chrt, err := loader.Load(chartPath)
	if err != nil {
		return nil, err
	}

	valsToRender, err := chartutil.ToRenderValues(
		chrt,
		chrt.Values,
		chartutil.ReleaseOptions{},
		chartutil.DefaultCapabilities,
	)
	if err != nil {
		return nil, err
	}

	render, err := engine.Render(chrt, valsToRender)
	if err != nil {
		return nil, err
	}

	return render, nil
}
