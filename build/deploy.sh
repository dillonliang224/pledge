#!/bin/bash

name=$1
tag=$2

repository="${CI_REGISTRY_IMAGE}/${name}"

echo "deploy ${image}"

deploy="helm upgrade --install --force --wait --tiller-namespace production --namespace production"

config="charts/values/${name}.yaml"
if [ -f "$config" ]; then
	deploy+=" -f $config "
fi

deploy+=" --set image.repository=${repository} --set image.tag=${tag} $name ./charts"

eval $deploy