#! /bin/bash

# 定义要停止的端口号列表
PORTS_TO_STOP=(4006 8600 8601 8700 8701 8800 8801 8802 8810 8888 9000 9001 9002)
# 循环处理每个端口号
for PORT_TO_STOP in "${PORTS_TO_STOP[@]}"
do
    # 查找监听指定端口的进程并获取其PID
    PID=$(lsof -t -i :$PORT_TO_STOP)

    if [ -z "$PID" ]; then
        echo "No process found listening on port $PORT_TO_STOP"
    else
        # 使用kill命令终止进程
        echo "Stopping process on port $PORT_TO_STOP (PID: $PID)"
        kill $PID
    fi
done



scriptDir=$(cd "$(dirname "$0")" && pwd)

homeDir="$scriptDir/.."
serviceDir="$homeDir/service"

echo "--------admin-rpc--------"
appPath="$serviceDir/admin/rpc"
cd "$appPath" || exit 0
go run admin.go &

echo "--------user-rpc--------"
appPath="$serviceDir/user/rpc"
cd "$appPath" || exit 0
go run user.go &

echo "--------portal-rpc--------"
appPath="$serviceDir/portal/rpc"
cd "$appPath" || exit 0
go run portal.go &

echo "--------shop-rpc--------"
appPath="$serviceDir/shop/rpc"
cd "$appPath" || exit 0
go run shop.go &

echo "--------lowcode-rpc--------"
appPath="$serviceDir/lowcode/rpc"
cd "$appPath" || exit 0
go run lowcode.go &

echo "--------tenant-rpc--------"
appPath="$serviceDir/tenant/rpc"
cd "$appPath" || exit 0
go run tenant.go &

echo "--------admin-api--------"
appPath="$serviceDir/admin/api"
cd "$appPath" || exit 0
go run admin.go &

echo "--------user-api--------"
appPath="$serviceDir/user/api"
cd "$appPath" || exit 0
go run user.go &

echo "--------portal-api--------"
appPath="$serviceDir/portal/api"
cd "$appPath" || exit 0
go run portal.go &

#echo "--------lowcode-api--------"
#appPath="$serviceDir/lowcode/api"
#cd "$appPath" || exit 0
#go run lowcode.go &

echo "--------tenant-api--------"
appPath="$serviceDir/tenant/api"
cd "$appPath" || exit 0
go run tenant.go &

#echo "--------shop-api--------"
#appPath="$serviceDir/shop/api"
#cd "$appPath" || exit 0
#go run shop.go &

echo "--------wechat-api--------"
appPath="$serviceDir/wechat/api"
cd "$appPath" || exit 0
go run wechat.go &