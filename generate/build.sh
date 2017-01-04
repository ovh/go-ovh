#!/bin/bash

set -e

ovhapi2openapi -i ovhapi-eu.yaml -o ovhapi-eu.gen.yaml

echo "TO RUN: java -Dmodels -DmodelTests=false -DmodelDocs=false -jar modules/swagger-codegen-cli/target/swagger-codegen-cli.jar generate --model-package types -i $(pwd)/ovhapi-eu.gen.yaml -l go -o $(pwd)/../ovh/types/"
