#! /bin/bash

parentDir=$(cd $(dirname $0);cd ..; pwd)

GOOS=linux
GOARCH=amd64

appPath="$parentDir/rpc/"

cd "$appPath"
GOOS=$GOOS GOARCH=$GOARCH go build -o main tenant.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)"
cd "$parentDir/../.."
docker build -t return1996/tenant-rpc:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)"