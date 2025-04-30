/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"coinstore/cmd/deploy/tron"
	"github.com/spf13/cobra"
)

// tronCmd represents the tron command
var tronCmd = &cobra.Command{
	Use:   "tron",
	Short: "TRON链合约初始化设置",
	Long: `功能描述：合约部署后，设置合约
使用示例: ./tb tron --admin_address '0xa...' --fee_address '0xa...' --server_address '0xa...' --key 'ee...' --relayer_one_address  '0x1...'   --relayer_two_address  '0x0...' --relayer_three_address '0x2...' --passphrase 'xxxxx' `,
	Run: func(cmd *cobra.Command, args []string) {
		tron.InitTron(PrivateKey, AdminAddress, FeeAddress, ServerAddress, RelayerOne, RelayerTwo, RelayerThree, Fee)
	},
}

func init() {
	rootCmd.AddCommand(tronCmd)
}
