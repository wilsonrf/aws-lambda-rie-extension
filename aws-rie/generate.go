/*
 * Copyright 2012-2024 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
	RUN curl -Lo ~/.aws-lambda-rie/aws-lambda-rie https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie
	RUN chmod +x ~/.aws-lambda-rie/aws-lambda-rie`)

	return packit.GenerateResult{
		RunDockerfile: r,
	}, nil
}
