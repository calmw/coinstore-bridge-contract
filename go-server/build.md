## 镜像构建

``` shell
# build 
docker buildx build --platform linux/amd64 --tag calmw/tt_bridge:0.0.10 --push .
```

``` shell
# build 
docker buildx build --platform linux/amd64 --tag calmw/tt_bridge_api:0.0.15 --push .
```

