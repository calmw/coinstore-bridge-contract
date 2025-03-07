## 生成绑定文件

- 生成绑定go文件命令

#### 生成userLocation合约代码

```shell
# Bridge
./abigen --abi ../abi/Bridge.json --pkg core --type Bridge --out ./pkg/binding/core/core.go
```

```shell
# Vote
./abigen --abi ../abi/Vote.json --pkg core --type Vote --out ./pkg/binding/core/vote.go
```

```shell
# Tantin
./abigen --abi ../abi/TantinBridge.json --pkg core --type Tantin --out ./pkg/binding/core/tantin.go
```
