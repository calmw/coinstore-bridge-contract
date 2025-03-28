FROM golang AS builder
WORKDIR /app
ADD . /app
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
RUN apt-get update && apt-get install -y gcc g++ make
RUN go mod tidy
RUN go build -o /tt_bridge cmd/bridge/main.go

FROM golang
WORKDIR /app
COPY --from=builder /tt_bridge ./
RUN chmod +x ./tt_bridge
RUN mkdir -p "/src/tron_keystore/keystore"

ENTRYPOINT ["./tt_bridge"]