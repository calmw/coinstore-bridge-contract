/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"coinstore/cmd/deploy/tt"
	"fmt"

	"github.com/spf13/cobra"
)

// ttCmd represents the tt command
var ttCmd = &cobra.Command{
	Use:   "tt",
	Short: "Tantin链合约初始化设置",
	Long: `功能描述：合约部署后，设置合约
使用示例: ./tb tt --admin_address '0xa...' --fee_address '0xa...' --server_address '0xa...' --key 'ee...' --relayer_one_address  '0x1...'   --relayer_two_address  '0x0...' --relayer_three_address '0x2...' `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("设置TT链合约...")
		tt.InitTt(PrivateKey, AdminAddress, FeeAddress, ServerAddress, RelayerOne, RelayerTwo, RelayerThree, Fee)
	},
}

func init() {
	rootCmd.AddCommand(ttCmd)
}
