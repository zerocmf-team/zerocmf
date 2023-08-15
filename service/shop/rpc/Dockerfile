FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates
ENV TZ Asia/Shanghai

WORKDIR /app

COPY service/shop/rpc/etc /app/etc
COPY service/shop/rpc/main /app/main

CMD ["./main", "-f", "etc/shop.yaml"]