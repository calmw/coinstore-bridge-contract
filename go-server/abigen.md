## 生成绑定文件

- 生成绑定go文件命令

#### 生成userLocation合约代码

```shell
# Bridge
./abigen --abi ../abi/Bridge.json --pkg bridge --type Bridge --out ./pkg/binding/bridge/bridge.go
```

```shell
# Vote
./abigen --abi ../abi/Vote.json --pkg bridge --type Vote --out ./pkg/binding/bridge/vote.go
```

```shell
# Tantin
./abigen --abi ../abi/TantinBridge.json --pkg bridge --type Tantin --out ./pkg/binding/bridge/tantin.go
```
