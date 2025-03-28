## 生成绑定文件

- 生成绑定go文件命令

#### 生成userLocation合约代码

```shell
# Bridge
./abigen --abi ../abi/Bridge.json --pkg binding --type Bridge --out ./binding/bridge.go
```

```shell
# Vote
./abigen --abi ../abi/Vote.json --pkg binding --type Vote --out ./binding/vote.go
```

```shell
# Tantin
./abigen --abi ../abi/TantinBridge.json --pkg binding --type Tantin --out ./binding/tantin.go
```
