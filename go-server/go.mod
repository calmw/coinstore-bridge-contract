module coinstore

go 1.23.0

toolchain go1.23.7

replace github.com/fbsobreira/gotron-sdk v0.0.0-20230907131216-1e824406fe8c => github.com/sunbankio/gotron-sdk v0.0.0-20231003155243-a269b0d040c3

replace github.com/codahale/hdrhistogram => github.com/HdrHistogram/hdrhistogram-go v1.1.2

require (
	github.com/btcsuite/btcd/btcec/v2 v2.3.2
	github.com/calmw/clog v0.0.3
	github.com/didip/tollbooth/v7 v7.0.2
	github.com/didip/tollbooth_gin v0.0.0-20250112173845-11eddec067c4
	github.com/ethereum/go-ethereum v1.14.3
	github.com/ethersphere/bee/v2 v2.5.0
	github.com/fbsobreira/gotron-sdk v0.0.0-20230907131216-1e824406fe8c
	github.com/forgoer/openssl v1.6.0
	github.com/gin-gonic/gin v1.10.0
	github.com/jasonlvhit/gocron v0.0.1
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.4.0
	github.com/status-im/keycard-go v0.2.0
	github.com/stretchr/testify v1.10.0
	golang.org/x/crypto v0.36.0
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.36.5
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.17.0 // indirect
	github.com/bytedance/sonic v1.13.1 // indirect
	github.com/bytedance/sonic/loader v0.2.4 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.5 // indirect
	github.com/consensys/bavard v0.1.22 // indirect
	github.com/consensys/gnark-crypto v0.14.0 // indirect
	github.com/crate-crypto/go-ipa v0.0.0-20240724233137-53bbb0ceb27a // indirect
	github.com/crate-crypto/go-kzg-4844 v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/deckarep/golang-set/v2 v2.6.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/ethereum/c-kzg-4844 v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/gin-contrib/sse v1.0.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-pkgz/expirable-cache/v3 v3.0.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.25.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.10 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.6.0 // indirect
	github.com/prometheus/common v0.47.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rjeczalik/notify v0.9.3 // indirect
	github.com/shengdoushi/base58 v1.0.0 // indirect
	github.com/shirou/gopsutil v3.21.5+incompatible // indirect
	github.com/supranational/blst v0.3.14 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	github.com/zondax/hid v0.9.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/arch v0.15.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto v0.0.0-20230526161137-0005af68ea54 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230525234035-dd9d682886f9 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
