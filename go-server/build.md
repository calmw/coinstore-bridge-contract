## 镜像构建

``` shell
# relayer node 
docker buildx build --platform linux/amd64 --tag calmw/tt_bridge:0.1.1 --push .
```

``` shell
# api 
docker buildx build --platform linux/amd64 --tag calmw/tt_bridge_api:0.0.38 --push .
```


## 部署程序构建

``` shell
# relayer node 
go build -o tb  -trimpath cmd/deploy/main.go
```

