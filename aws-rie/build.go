package awsrie

import (
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

type Build struct {
	Logger scribe.Logger
}

func (b *Build) Build(packit.BuildContext) (packit.BuildResult, error) {
	return packit.BuildResult{}, nil
}
