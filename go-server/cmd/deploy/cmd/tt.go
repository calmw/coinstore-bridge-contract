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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("设置TT链合约...")
		tt.InitTt(PrivateKey, AdminAddress, FeeAddress, ServerAddress, RelayerOne, RelayerTwo, RelayerThree, Fee)
	},
}

func init() {
	rootCmd.AddCommand(ttCmd)
	rootCmd.PersistentFlags().StringVarP(&AdminAddress, "admin_address", "a", "", "管理员角色账户")
	rootCmd.PersistentFlags().StringVarP(&FeeAddress, "fee_address", "f", "", "跨链费接受地址")
	rootCmd.PersistentFlags().StringVarP(&ServerAddress, "server_address", "s", "", "服务端价格签名地址")
	rootCmd.PersistentFlags().StringVarP(&RelayerOne, "relayer_one_address", "l", "", "relayer 1 账户地址")
	rootCmd.PersistentFlags().StringVarP(&RelayerTwo, "relayer_two_address", "m", "", "relayer 2 账户地址")
	rootCmd.PersistentFlags().StringVarP(&RelayerThree, "relayer_three_address", "n", "", "relayer 3 账户地址")
	rootCmd.PersistentFlags().Uint64VarP(&Fee, "fee", "e", 4, "跨链费，折合U的数量")
}
