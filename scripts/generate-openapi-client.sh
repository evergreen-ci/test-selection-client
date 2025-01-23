#!/usr/bin/env bash

# kim: TODO: make sure this bin is at the root directory
mkdir -p bin
cd bin || return

# Download maven
curl -LO https://dlcdn.apache.org/maven/maven-3/3.9.9/binaries/apache-maven-3.9.9-bin.tar.gz >apache-maven-3.9.9-bin.tar.gz
tar xvzf apache-maven-3.9.9-bin.tar.gz

# Download OpenAPI generator CLI.
export PATH="${PWD}"/apache-maven-3.9.9/bin:"${PATH}"
export PATH=/opt/bin/java/jdk21/bin:"${PATH}"
curl -LO https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh
chmod +x ./openapi-generator-cli.sh

curl -LO https://mciuploads.s3.amazonaws.com/evergreen/latest/swagger.json

# Generate OpenAPI client
cd ..
# kim: TODO: fill in actual test selection spec URL
bin/openapi-generator-cli.sh generate -i bin/swagger.json -g go -o ./ --git-user-id evergreen-ci --git-repo-id test-selection-client
