#! /bin/bash

scriptDir=$(cd "$(dirname "$0")" && pwd)

homeDir="$scriptDir/.."
serviceDir="$homeDir/service"

GOOS=linux
GOARCH=amd64

appPath="$serviceDir/admin/api"

echo "--------$appPath--------"

cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main admin.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/admin-api:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log


echo "--------$appPath--------"

appPath="$serviceDir/admin/rpc"
cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main admin.go

echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/admin-rpc:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log

echo "--------$appPath--------"

appPath="$serviceDir/user/api"
cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main user.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/user-api:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log

echo "--------$appPath--------"

appPath="$serviceDir/user/rpc"
cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main user.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/user-rpc:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log

echo "--------$appPath--------"

appPath="$serviceDir/portal/api"
cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main portal.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/portal-api:latest -f "$appPath/Dockerfile" .
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log

echo "--------$appPath--------"

appPath="$serviceDir/portal/rpc"
cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main portal.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/portal-rpc:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log

echo "--------$appPath--------"

appPath="$serviceDir/tenant/api"
cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main tenant.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/tenant-api:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log

echo "--------$appPath--------"

appPath="$serviceDir/tenant/rpc"
cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main tenant.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/tenant-rpc:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log

echo "--------$appPath--------"

appPath="$serviceDir/lowcode/api"
cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main lowcode.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/lowcode-api:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log

echo "--------$appPath--------"

appPath="$serviceDir/lowcode/rpc"
cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main lowcode.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/lowcode-rpc:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log

echo "--------$appPath--------"

appPath="$serviceDir/shop/api"
cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main shop.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/shop-api:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log

echo "--------$appPath--------"

appPath="$serviceDir/shop/rpc"
cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main shop.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/shop-rpc:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log

echo "--------$appPath--------"

appPath="$serviceDir/wechat/api"
cd "$appPath" || exit 0
GOOS=$GOOS GOARCH=$GOARCH go build -o main wechat.go
echo "build app finished at $(date +%Y-%m-%d\ %H:%M:%S)" > build.log
cd "$homeDir" || exit 0
docker build -t return1996/wechat-api:latest -f "$appPath/Dockerfile" .
echo "build dockerfile finished at $(date +%Y-%m-%d\ %H:%M:%S)" >> build.log

echo "--------$appPath--------"