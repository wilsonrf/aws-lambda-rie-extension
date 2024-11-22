package awsrie

import (
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

type Generate struct {
	Logger scribe.Logger
}

func (b *Generate) Build(packit.GenerateContext) (packit.GenerateResult, error) {
	return packit.GenerateResult{}, nil
}
