#!/bin/bash

#set -x

pledge::util:build_docker_images() {
  local docker_registry=$1
  local docker_tag=$2
  local base_image="alpine:3.7"

  query=$(pledge::util::find_changes)

  if [ "$query" == "" ]; then
    pledge::util::log "no change and exit..."
    exit 0
  fi

  for b in ${query}; do
    b=${b//\/\/src/"/src"}

    local binary_file_path=$(pledge::util::find_binary "$b")
    local binary_name=$(pledge::util::get_binary_name "$b")
    local docker_build_path="dockerbuild/${binary_name}"
    local docker_file_path="${docker_build_path}/Dockerfile"
    local docker_image_tag="${docker_registry}/${binary_name}:${docker_tag}"

    pledge::util::log "Starting docker build for image: ${binary_name}"
    (
      rm -rf "${docker_build_path}"
      mkdir -p "${docker_build_path}"
      cp "${binary_file_path}" "${docker_build_path}/${binary_name}"
      cat <<EOF >"${docker_file_path}"
FROM ${base_image}
COPY ${binary_name} /usr/local/bin/${binary_name}
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
  && apk update --no-cache \
  && apk add --no-cache tzdata ca-certificates \
  && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
  && echo "Asia/Shanghai" > /etc/timezone
  
ENTRYPOINT ["/usr/local/bin/${binary_name}"]
EOF
      docker build -q -t "${docker_image_tag}" "${docker_build_path}"
      docker push ${docker_image_tag}
    )

    cat <<EOF >>".gitlab/deploy.yaml"
${binary_name}:
  stage: deploy
  script:
    - bash build/deploy.sh ${binary_name} ${docker_tag}
  only:
    - tags
  when: manual
  environment:
    name: production
  tags:
    - docker

EOF

  done

  pledge::util::log "Docker builds done"
}

pledge::util::find_changes() {
  files=$(git diff $(git rev-list --tags --max-count=1) ${CI_COMMIT_SHA} --name-only --diff-filter=ACM | grep -E -i ".go$")
  # files=$(git diff 733abfd1ea5d992c08bd2864cf58b1fb840406aa 6c517d77200edf6e16e1ae0084445b4ccc7612af --name-only --diff-filter=ACM | grep -E -i ".go$")
  # echo "change files:\n${files}\n"

  paths=""
  for file in ${files}; do
    if [[ "${paths}" != "" ]]; then
      paths="${paths} union allpaths(//src/..., ${file})"
    else
      paths="allpaths(//src/..., ${file})"
    fi
  done

  query=$(bazel query "${paths}" --keep_going | grep //src/ | grep -v go_default_library | grep -v .go)
  echo -e "${query}"
}

pledge::util::find_binary() {
  local -r lookfor="$1"
  local -r platform=$(pledge::util::host_platform)

  IFS=': ' read -r -a array <<<"${lookfor}"

  local bin=$(find "bazel-bin/" -type f -path "*${array[0]}/${platform/\//_}*/${array[1]}" 2>/dev/null || true)

  echo -n $bin
}

pledge::util::get_binary_name() {
  local -r lookfor="$1"
  IFS=': ' read -r -a array <<<"${lookfor}"
  name=${array[0]//\/\/src\//""}
  name=${name//\/src\//""}
  name=${name//\/cmd/""}
  name=${name//"service"/"svc"}
  name=${name//\//"-"}
  echo $name
}

pledge::util::auto_tag() {
  version=$(git describe --tags $(git rev-list --tags --max-count=1) 2>/dev/null || true)

  major=0
  minor=0
  build=0

  regex="([0-9]+).([0-9]+).([0-9]+)"
  if [[ $version =~ $regex ]]; then
    major="${BASH_REMATCH[1]}"
    minor="${BASH_REMATCH[2]}"
    build="${BASH_REMATCH[3]}"
  fi

  level=$(echo $CI_COMMIT_MESSAGE | cut -d' ' -f 3 | sed "s/^'\(.*\)'$/\1/" | cut -d'/' -f 1)

  if [[ $level == "feature" ]]; then
    minor=$((minor + 1))
    build=0
  elif [[ $level == "hotfix" ]]; then
    build=$((build + 1))
  elif [[ $level == "release" ]]; then
    major=$((major + 1))
    minor=0
    build=0
  else
    build=$((build + 1))
  fi

  echo "${major}.${minor}.${build}"
}

pledge::util::host_platform() {
  local host_os
  local host_arch
  case "$(uname -s)" in
  Darwin)
    host_os=darwin
    ;;
  Linux)
    host_os=linux
    ;;
  *)
    kratos::log::error "Unsupported host OS.  Must be Linux or Mac OS X."
    exit 1
    ;;
  esac

  case "$(uname -m)" in
  x86_64*)
    host_arch=amd64
    ;;
  i?86_64*)
    host_arch=amd64
    ;;
  amd64*)
    host_arch=amd64
    ;;
  aarch64*)
    host_arch=arm64
    ;;
  arm64*)
    host_arch=arm64
    ;;
  arm*)
    host_arch=arm
    ;;
  i?86*)
    host_arch=x86
    ;;
  s390x*)
    host_arch=s390x
    ;;
  ppc64le*)
    host_arch=ppc64le
    ;;
  *)
    pledge::util::log "Unsupported host arch. Must be x86_64, 386, arm, arm64, s390x or ppc64le."
    exit 1
    ;;
  esac
  echo "${host_os}/${host_arch}"
}

pledge::util::log() {
  timestamp=$(date +"[%m%d %H:%M:%S]")
  echo "+++ ${timestamp} ${1}"
  shift
  for message; do
    echo "    ${message}"
  done
}