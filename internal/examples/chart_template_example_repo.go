package examples

import (
	"fmt"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/engine"
	"helm.sh/helm/v3/pkg/registry"
)

func chartTemplateExampleRepo() {
	settings := cli.New()
	actionConfig := &action.Configuration{}
	if err := actionConfig.Init(
		settings.RESTClientGetter(),
		settings.Namespace(),
		"",
		nil,
	); err != nil {
		panic(err)
	}

	if actionConfig.RegistryClient == nil {
		regClient, err := registry.NewClient(
			registry.ClientOptWriter(os.Stderr),
		)
		if err != nil {
			panic(err)
		}
		actionConfig.RegistryClient = regClient
	}

	client := action.NewInstall(actionConfig)

	client.DryRun = true
	client.ClientOnly = true
	client.RepoURL = "https://helm.github.io/examples"

	chartPath, err := client.LocateChart("hello-world", settings)
	if err != nil {
		panic(err)
	}

	chrt, err := loader.Load(chartPath)
	if err != nil {
		panic(err)
	}

	valsToRender, err := chartutil.ToRenderValues(chrt, chrt.Values, chartutil.ReleaseOptions{}, chartutil.DefaultCapabilities)

	rendered, err := engine.Render(chrt, valsToRender)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", rendered)
}
