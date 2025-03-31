#!/usr/bin/env bash

set -eo pipefail

mkdir -p bin
cd bin

# Download maven (required by OpenAPI generator CLI).
curl -O https://dlcdn.apache.org/maven/maven-3/3.9.9/binaries/apache-maven-3.9.9-bin.tar.gz
tar xvzf apache-maven-3.9.9-bin.tar.gz

# Download OpenAPI generator CLI.
curl -O https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh
chmod +x ./openapi-generator-cli.sh

cd ..

# Generate OpenAPI client.
export PATH="${PWD}"/apache-maven-3.9.9/bin:"${PATH}"
export PATH=/opt/bin/java/jdk21/bin:"${PATH}"
# It's assumed that the latest JSON spec is downloaded locally in bin/test_selection_services.json.
bin/openapi-generator-cli.sh generate -i bin/test_selection_services.json -g go -o ./ --git-user-id evergreen-ci --git-repo-id test-selection-client
