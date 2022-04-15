#!/usr/bin/env bash
set -eux

TimeTag=$(date +%Y%m%d-%H%M%S%z)
CommitId=$(git rev-parse --short HEAD)
BuildMachine=$(whoami)@$(hostname)

GbTag="github.com/g1022/gbtag"
LdFlags="\
-X '${GbTag}.Time=$TimeTag' \
-X '${GbTag}.Commit=$CommitId'  \
-X '${GbTag}.Machine=$BuildMachine' "

go build -ldflags "${LdFlags}"   -o ../build/demo main.go

echo "run build bin..."
../build/demo

