## 镜像构建

``` shell
# relayer node 
docker buildx build --platform linux/amd64 --tag calmw/tt_bridge:0.1.1 --push .


docker buildx build --platform linux/amd64 --tag harbor.devops.tantin.com/chain/tt_bridge:0.1.0 --push .
```

``` shell
# api 
docker buildx build --platform linux/amd64 --tag calmw/tt_bridge_api:0.0.39 --push .
```

## 部署程序构建

``` shell
# relayer node 
go build -o tb  -trimpath cmd/deploy/main.go
```

``` shell
# EVM使用示例 
./tb tron --admin_address 'TFBymbm7LrbRreGtByMPRD2HUyneKabsqb' --fee_address 'TFBymbm7LrbRreGtByMPRD2HUyneKabsqb' --server_address 'TFBymbm7LrbRreGtByMPRD2HUyneKabsqb' --relayer_one_address  'TTgY73yj5vzGM2HGHhVt7AR7avMW4jUx6n'   --relayer_two_address  'TSARBFH6PW6jEuf8chd1DxZGW6JEmHuv6g' --relayer_three_address 'TEz4CMzy3mgtVECcYxu5ui9nJfgv3oXhyx' --fee 4 --passphrase '123456' --key 'XXXXXXX'
```

``` shell
# TRON使用示例 
./tb tron --admin_address 'TFBymbm7LrbRreGtByMPRD2HUyneKabsqb' --fee_address 'TFBymbm7LrbRreGtByMPRD2HUyneKabsqb' --server_address 'TFBymbm7LrbRreGtByMPRD2HUyneKabsqb' --relayer_one_address  'TTgY73yj5vzGM2HGHhVt7AR7avMW4jUx6n'   --relayer_two_address  'TSARBFH6PW6jEuf8chd1DxZGW6JEmHuv6g' --relayer_three_address 'TEz4CMzy3mgtVECcYxu5ui9nJfgv3oXhyx' --fee 4 --key 'XXXXXXX'
```

