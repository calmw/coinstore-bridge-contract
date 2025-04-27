/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tronCmd represents the tron command
var tronCmd = &cobra.Command{
	Use:   "tron",
	Short: "TRON链合约初始化设置",
	Long: `功能描述：合约部署后，设置合约
使用示例: ./tb tron --admin_address '0xa...' --fee_address '0xa...' --server_address '0xa...' --key 'ee...' --relayer_one_address  '0x1...'   --relayer_two_address  '0x0...' --relayer_three_address '0x2...' --passphrase 'xxxxx' `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tron called")
	},
}

func init() {
	rootCmd.AddCommand(tronCmd)

	tronCmd.PersistentFlags().StringVarP(&TronKeyStorePassphrase, "passphrase", "p", "123456", "tron keystore passphrase")
}
