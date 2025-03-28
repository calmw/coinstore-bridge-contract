# 第一阶段：构建应用
FROM golang AS builder
WORKDIR /app
ADD . /app
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
RUN apt-get update && apt-get install -y gcc g++ make
RUN go mod tidy
RUN go build -o /app/bridge cmd/bridge/main.go

# 第二阶段：创建最小运行时镜像
FROM golang
WORKDIR /app
COPY --from=builder /app/bridge /app/bridge
RUN chmod +x /app/bridge
RUN mkdir -p "/src/tron_keystore/keystore"

ENTRYPOINT ["./bridge"]
