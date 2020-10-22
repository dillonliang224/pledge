#!/bin/bash

set -o errexit
set -o pipefail

source "build/init.sh"
source "build/utils.sh"

bazel build //src/... //library/... 
cat bazel-out/stable-status.txt

tag=$(pledge::util::auto_tag)
echo "build tag: $tag"

cat <<EOF >".gitlab/deploy.yaml"
deploy:
  stage: deploy
  script:
    - echo "deploy"
  only:
    - /^deploy-.*$/
  tags:
    - docker

EOF

pledge::util:build_docker_images  "${CI_REGISTRY_IMAGE}" "${tag}"

git add .
git commit -m "pledge deploy..."
git tag $tag
git push --tags