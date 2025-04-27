/*
Package cmd
Copyright © 2025 calm.wang@hotmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	PrivateKey             string
	AdminAddress           string
	FeeAddress             string
	ServerAddress          string
	RelayerOne             string
	RelayerTwo             string
	RelayerThree           string
	TronKeyStorePassphrase string
	Fee                    uint64
)

var rootCmd = &cobra.Command{
	Use:   "tb",
	Short: "合约初始化设置",
	Long: `功能描述：合约部署后，设置合约
使用示例: ./tb bsc --admin_address '0xa...' --fee_address '0xa...' --server_address '0xa...' --key 'ee...' --relayer_one_address  '0x1...'   --relayer_two_address  '0x0...' --relayer_three_address '0x2...' `,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&AdminAddress, "admin_address", "a", "", "管理员角色账户")
	rootCmd.PersistentFlags().StringVarP(&FeeAddress, "fee_address", "f", "", "跨链费接受地址")
	rootCmd.PersistentFlags().StringVarP(&ServerAddress, "server_address", "s", "", "服务端价格签名地址")
	rootCmd.PersistentFlags().StringVarP(&PrivateKey, "key", "k", "", "default admin账户私钥")
	rootCmd.PersistentFlags().StringVarP(&RelayerOne, "relayer_one_address", "l", "", "relayer 1 账户地址")
	rootCmd.PersistentFlags().StringVarP(&RelayerTwo, "relayer_two_address", "m", "", "relayer 2 账户地址")
	rootCmd.PersistentFlags().StringVarP(&RelayerThree, "relayer_three_address", "n", "", "relayer 3 账户地址")
	rootCmd.PersistentFlags().Uint64VarP(&Fee, "fee", "e", 4, "跨链费，折合U的数量")
}
