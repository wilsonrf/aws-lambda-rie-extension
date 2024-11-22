package awsrie

import (
	"strings"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

type Generate struct {
	Logger scribe.Logger
}

func (b *Generate) Generate(context packit.GenerateContext) (packit.GenerateResult, error) {

	r := strings.NewReader(`ARG base_image
	FROM ${base_image}
	RUN mkdir -p ~/.aws-lambda-rie
	USER root
	RUN apt-get update && apt-get install -y curl
	RUN curl -Lo ~/.aws-lambda-rie/aws-lambda-rie https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie
	RUN chmod +x ~/.aws-lambda-rie/aws-lambda-rie
	ARG user_id
	USER ${user_id}`)

	return packit.GenerateResult{
		RunDockerfile: r,
	}, nil
}
