/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "合约初始化设置",
	Long: `合约部署后，初始化设置合约,使用示例:
tt --chain TT/BSC/TRON/ETH  --key abc123...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploy called")
		switch ChainName {
		case "Tantin":
		case "BSC":
		case "TRON":
		case "ETH":

		}
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.PersistentFlags().StringVarP(&ChainName, "chain", "c", "TT", "链名称")
	deployCmd.PersistentFlags().StringVarP(&PrivateKey, "key", "k", "abc123...", "私钥")
}
