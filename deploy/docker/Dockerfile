FROM golang:alpine AS builder

LABEL stage=gobuilder

ARG SERVICE_TYPE
ARG SERVICE_NAME

ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .
RUN go get -u github.com/ve-weiyi/vkit@latest
RUN go mod tidy
RUN go build -ldflags="-s -w" -o /app/server service/${SERVICE_NAME}/${SERVICE_TYPE}/*.go

FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/server /app/server

CMD ["./server"]
