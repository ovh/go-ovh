#!/bin/bash

set -e

# /paas/monitoring is skipped

ovhapi2openapi -i ovhapi-eu.yaml -o ovhapi-eu.gen.yaml

echo "TO RUN: java -Dmodels -DmodelTests=false -DmodelDocs=false -jar modules/swagger-codegen-cli/target/swagger-codegen-cli.jar generate --model-package types -i $(pwd)/ovhapi-eu.gen.yaml -l go -o $(pwd)/../ovh/types/ && go fmt github.com/runabove/go-sdk/ovh/types --type-mappings=paas.Timeseries.PermissionEnum=string,dedicated.Ceph.UserPoolPermSetAll.Permissions=string,paas.Timeseries.Consumption.Item.MetricNameEnum=string,paas.Timeseries.QuotaTypeEnum=string"
