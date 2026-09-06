#!/bin/bash

set -e
usage() {
  echo "Usage: $(basename "$0") [-h] [-l]"
  echo " -l  : build linux amd64 binaries only, default to build for all platforms"
}

build=all
while getopts "hl" arg; do
  case $arg in
  h)
    usage
    exit 0
    ;;
  l)
    build=linux_amd64
    ;;
  *)
    usage
    exit 1
    ;;
  esac
done
shift $((OPTIND - 1))

program=json2table
source=./cmd/json2table

# shellcheck disable=SC1091
. ./release.env

build_date=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
git_commit=$(git rev-parse HEAD)

build(){
  local GOOS=$1
  local GOARCH=$2
  local extension=$3
  echo "Building for $GOOS/$GOARCH"
  go build \
  -ldflags "-X github.com/siakhooi/json2table/internal/versioninfo.Version=$RELEASE_VERSION -X github.com/siakhooi/json2table/internal/versioninfo.Date=$build_date -X github.com/siakhooi/json2table/internal/versioninfo.Commit=$git_commit" \
  -o bin/"${program}-${GOOS}-${GOARCH}${extension}" $source
}

if [[ "$build" == "linux_amd64" ]]; then
  build linux amd64 ""
else
  build linux amd64 ""
  build linux arm64 ""
  build windows amd64 ".exe"
  build windows 386 ".exe"
  build darwin amd64 ""
  build darwin arm64 ""
  build freebsd amd64 ""
  build freebsd arm64 ""
  build netbsd amd64 ""
  build netbsd arm64 ""
  build openbsd amd64 ""
  build openbsd arm64 ""
fi
