FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates
ENV TZ Asia/Shanghai

WORKDIR /app

COPY service/tenant/api/etc /app/etc
COPY service/tenant/api/main /app/main

CMD ["./main", "-f", "etc/tenant.yaml"]