#!/bin/bash

GIT_COMMIT=$(git rev-parse --short HEAD)
TAG=$(git describe --exact-match --abbrev=0 --tags ${COMMIT} 2> /dev/null || true)
DATE=$(date +'%Y-%m-%d')

echo "Building binaries"
echo Git commit: $GIT_COMMIT Version: $TAG Build date: $DATE

go generate

#LINUX
export GOARCH="amd64"
export GOOS="linux"
export CGO_ENABLED=0
go build -ldflags "-X github.com/simon-harlow/docgen/cmd.GitCommit=$GIT_COMMIT -X github.com/simon-harlow/docgen/cmd.Version=$TAG -X github.com/simon-harlow/docgen/cmd.BuildDate=$DATE" -o bin/linux_amd64 -v

export GOARCH="arm64"
export GOOS="linux"
export CGO_ENABLED=0
go build -ldflags "-X github.com/simon-harlow/docgen/cmd.GitCommit=$GIT_COMMIT -X github.com/simon-harlow/docgen/cmd.Version=$TAG -X github.com/simon-harlow/docgen/cmd.BuildDate=$DATE" -o bin/linux_arm64 -v

echo "Build complete"