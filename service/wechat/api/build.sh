#! /bin/bash

parentDir=$(cd $(dirname $0);cd ..; pwd)

GOOS=linux
GOARCH=amd64

appPath="$parentDir/api/"

cd "$appPath"
GOOS=$GOOS GOARCH=$GOARCH go build -o main wechat.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)"
cd "$parentDir/../.."
docker build -t return1996/wechat-api:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)"

docker push return1996/wechat-api:latest