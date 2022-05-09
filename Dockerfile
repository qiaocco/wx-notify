FROM golang:alpine AS builder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY main.go controllers.go config.toml.bak ./
RUN go build -ldflags="-s -w" -o /app/wx-notify

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/wx-notify /app/wx-notify
COPY --from=builder /build/config.toml.bak /app/config.toml

CMD ["./wx-notify"]
