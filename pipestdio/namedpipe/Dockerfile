FROM golang:1.10.3-alpine3.8 AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories

RUN apk update && \
    apk add git build-base && \
    rm -rf /var/cache/apk/* && \
    mkdir -p "$GOPATH/src/gitlab.alipay-inc.com/basement/cube/logcollector"

ADD . "$GOPATH/src/gitlab.alipay-inc.com/basement/cube/logcollector"

RUN cd "$GOPATH/src/gitlab.alipay-inc.com/basement/cube/logcollector" && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --installsuffix cgo --ldflags="-s" -o /logcollector ./

FROM alpine:3.8
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories
RUN apk add --update ca-certificates

COPY --from=builder /logcollector /bin/logcollector

ENTRYPOINT ["/bin/logcollector"]