FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk update && apk upgrade

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY service/admin/api/data /app/data
COPY service/admin/api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/admin service/admin/api/admin.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/admin /app/admin
COPY --from=builder /app/data /app/data
COPY --from=builder /app/etc /app/etc

CMD ["./admin", "-f", "etc/admin.yaml"]
