FROM  golang AS builder
ADD .. /src
WORKDIR /src
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN cd cmd/server && go build -o /bridge_api .

# final stage
FROM ubuntu
RUN apt-get -y update && apt-get -y upgrade
COPY --from=builder /bridge_api ./
RUN chmod +x ./bridge_api

ENTRYPOINT ["./bridge_api"]
