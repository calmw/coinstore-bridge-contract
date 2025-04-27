/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"coinstore/cmd/deploy/eth"
	"fmt"

	"github.com/spf13/cobra"
)

// ethCmd represents the eth command
var ethCmd = &cobra.Command{
	Use:   "bsc",
	Short: "ETH链合约初始化设置",
	Long: `功能描述：合约部署后，设置合约
使用示例: ./tb bsc --admin_address '0xa...' --fee_address '0xa...' --server_address '0xa...' --key 'ee...' --relayer_one_address  '0x1...'   --relayer_two_address  '0x0...' --relayer_three_address '0x2...' `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("设置ETH链合约...")
		eth.InitTt(PrivateKey, AdminAddress, FeeAddress, ServerAddress, RelayerOne, RelayerTwo, RelayerThree, Fee)
	},
}

func init() {
	rootCmd.AddCommand(ethCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ethCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ethCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
