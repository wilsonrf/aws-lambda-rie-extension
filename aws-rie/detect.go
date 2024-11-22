package awsrie

import (
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

const (
	PlanEntryCustomRuntimeEmulator = "aws-custom-runtime-emulator-extension"
)

type Detect struct {
	Logger scribe.Logger
}

func (d *Detect) Detect(context packit.DetectContext) (packit.DetectResult, error) {

	return packit.DetectResult{
		Plan: packit.BuildPlan{
			Provides: []packit.BuildPlanProvision{
				{Name: PlanEntryCustomRuntimeEmulator},
			},
		},
	}, nil
}
