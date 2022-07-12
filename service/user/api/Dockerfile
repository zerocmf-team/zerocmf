FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates
ENV TZ Asia/Shanghai

WORKDIR /app

COPY service/user/api/etc /app/etc
COPY service/user/api/main /app/main

CMD ["./main", "-f", "etc/user.yaml"]