package main

import (
	"os"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	awsrie "github.com/wilsonrf/aws-rie-extension/aws-rie"
)

func main() {
	logger := scribe.NewLogger(os.Stdout)
	detect := awsrie.Detect{Logger: logger}
	generate := awsrie.Generate{Logger: logger}
	packit.RunExtension(detect.Detect, generate.Generate)
}
