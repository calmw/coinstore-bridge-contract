## 镜像构建

``` shell
# build 
docker buildx build --platform linux/amd64 --tag calmw/cs_bridge:0.0.2 --push .
```

``` shell
# build 
docker buildx build --platform linux/amd64 --tag calmw/cs_bridge_api:0.0.2 --push .
```

## X86版本

- calmw/bridge:0.10.51

## Arm版本

- calmw/bridge:0.10.31
