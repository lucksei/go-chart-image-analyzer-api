package examples

import (
	"fmt"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func containerRegistryCustomExample() {
	fmt.Println("Parsing image reference...")
	ref, err := name.ParseReference("docker.io/kooldev/pause:latest")
	if err != nil {
		panic(err)
	}
	fmt.Println("Pulling image...")
	img, err := remote.Image(ref, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		panic(err)
	}

	manifest, err := img.Manifest()
	for _, layer := range manifest.Layers {
		digest := layer.Digest
		fmt.Println(digest.String())
	}
}
