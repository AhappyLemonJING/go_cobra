package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hookRootCmd = &cobra.Command{
	Use: "hookroot",
	Run: func(cmd *cobra.Command, args []string) {
		// 第三个被执行
		fmt.Println("钩子函数 run")

	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// run函数之前执行  可被继承  第一个被执行
		fmt.Println("PersistentPreRun")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// run函数之后执行  可被继承   第五个被执行
		fmt.Println("PersistentPostRun")
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// run函数之前执行  不可被继承  第二个被执行
		fmt.Println("PreRun")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		// run函数之后执行  不可被继承  第四个被执行
		fmt.Println("PostRun")
	},
}

var hookSubCmd = &cobra.Command{
	Use: "hooksub",
	Run: func(cmd *cobra.Command, args []string) {
		// 第三个被执行
		fmt.Println("钩子函数 run")

	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// run函数之后执行  可被继承   第五个被执行
		fmt.Println("PersistentPostRun")
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// run函数之前执行  不可被继承  第二个被执行
		fmt.Println("PreRun")
	},
}

func init() {
	rootCmd.AddCommand(hookRootCmd)
	hookRootCmd.AddCommand(hookSubCmd)
}
