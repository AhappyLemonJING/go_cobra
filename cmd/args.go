package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// 自定义验证
var curArgsCheckCmd = &cobra.Command{
	Use:   "cus",
	Long:  "",
	Short: "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("至少输入一个参数")
		}
		if len(args) > 2 {
			return errors.New("至多输入两个个参数")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("自定义参数验证 start")
		fmt.Println(args)
		fmt.Println("自定义参数验证 end")
	},
}

// 无参数验证
var noArgsCmd = &cobra.Command{
	Use:  "noargs",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("无参数验证 start")
		fmt.Println(args)
		fmt.Println("无参数验证 end")
		return nil
	},
}

// 可以接受任何参数
var arbitrayArgCmd = &cobra.Command{
	Use:  "ab",
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("可接受任何参数 start")
		fmt.Println(args)
		fmt.Println("可接受任何参数 end")
		return nil
	},
}

var onlyArgsCmd = &cobra.Command{
	Use:       "only",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"123", "456"},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("only start")
		fmt.Println(args)
		fmt.Println("only end")
		return nil
	},
}

var exactArgsCmd = &cobra.Command{
	Use:  "exact",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("exact start")
		fmt.Println(args)
		fmt.Println("exact end")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(curArgsCheckCmd)
	rootCmd.AddCommand(noArgsCmd)
	rootCmd.AddCommand(arbitrayArgCmd)
	rootCmd.AddCommand(onlyArgsCmd)
	rootCmd.AddCommand(exactArgsCmd)
}
