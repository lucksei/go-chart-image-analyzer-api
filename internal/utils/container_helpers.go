package utils

import (
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func PullImageAndParseAPIInfo(image string) (ImageAnalysis, error) {
	imageAnalysis := ImageAnalysis{}

	ref, err := name.ParseReference(image)
	if err != nil {
		return ImageAnalysis{}, err
	}
	imageAnalysis.Name = ref.Name()

	img, err := remote.Image(ref, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		return ImageAnalysis{}, err
	}
	manifest, err := img.Manifest()
	if err != nil {
		return ImageAnalysis{}, err
	}

	imageAnalysis.LayerNumber = len(manifest.Layers)

	for _, layer := range manifest.Layers {
		imageAnalysis.Size += layer.Size
	}

	return imageAnalysis, err
}
