ARG GO_VERSION=1.23
FROM registry.cn-beijing.aliyuncs.com/qqlx/golang:${GO_VERSION}-alpine AS builder
WORKDIR /app
COPY . .
RUN <<EOF
  set -ex
  export GOPROXY=https://goproxy.cn,direct
  go build -o app -ldflags="-s -w" .
EOF

FROM registry.cn-beijing.aliyuncs.com/qqlx/alpine:3.12
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app", "run"]
