## 镜像构建

``` shell
# build 
docker buildx build --platform linux/amd64 --tag calmw/cs_bridge:0.0.1 --push .
```

``` shell
# build 
docker buildx build --platform linux/amd64 --tag calmw/cs_bridge_api:0.0.1 --push .
```

