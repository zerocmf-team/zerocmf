#! /bin/bash

parentDir=$(cd $(dirname $0);cd ..; pwd)

GOOS=linux
GOARCH=amd64

appPath="$parentDir/admin/api"
cd "$appPath"
GOOS=$GOOS GOARCH=$GOARCH go build -o main admin.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$parentDir/.."
docker build -t return1996/admin-api:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log
docker push return1996/admin-api:latest

appPath="$parentDir/admin/rpc"
cd "$appPath"
GOOS=$GOOS GOARCH=$GOARCH go build -o main admin.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$parentDir"/..
docker build -t return1996/admin-rpc:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log
docker push return1996/admin-rpc:latest

appPath="$parentDir/user/api"
cd "$appPath"
GOOS=$GOOS GOARCH=$GOARCH go build -o main user.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$parentDir"/..
docker build -t return1996/user-api:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log
docker push return1996/user-api:latest

appPath="$parentDir/user/rpc"
cd "$appPath"
GOOS=$GOOS GOARCH=$GOARCH go build -o main user.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$parentDir"/..
docker build -t return1996/user-rpc:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log
docker push return1996/user-rpc:latest

appPath="$parentDir/portal/api"
cd "$appPath"
GOOS=$GOOS GOARCH=$GOARCH go build -o main portal.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$parentDir"/..
docker build -t return1996/portal-api:latest -f "$appPath/Dockerfile" .
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
docker push return1996/portal-api:latest

appPath="$parentDir/portal/rpc"
cd "$appPath"
GOOS=$GOOS GOARCH=$GOARCH go build -o main portal.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$parentDir"/..
docker build -t return1996/portal-rpc:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log
docker push return1996/portal-rpc:latest

appPath="$parentDir/tenant/api"
cd "$appPath"
GOOS=$GOOS GOARCH=$GOARCH go build -o main tenant.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$parentDir"/..
docker build -t return1996/tenant-api:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log
docker push return1996/tenant-api:latest

appPath="$parentDir/tenant/rpc"
cd "$appPath"
GOOS=$GOOS GOARCH=$GOARCH go build -o main tenant.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$parentDir"/..
docker build -t return1996/tenant-rpc:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log
docker push return1996/tenant-rpc:latest