FROM  golang AS builder
ADD . /src
WORKDIR /src
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
RUN go mod tidy
RUN cd cmd/server && go build -o /bridge_api .

# final stage
FROM ubuntu
COPY --from=builder /bridge_api ./
RUN chmod +x ./bridge_api

ENTRYPOINT ["./bridge_api"]
