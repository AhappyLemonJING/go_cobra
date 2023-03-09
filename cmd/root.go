package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mycobra", //根目录其实用不上 随便叫啥都行
	Short: "简短的描述",
	Long:  "详细的描述",
}

func Execute() {
	rootCmd.Execute()
}

var userLicense string

func init() {
	rootCmd.PersistentFlags().Bool("viper", true, "是否采用viper读取配置文件")
	rootCmd.PersistentFlags().StringP("author", "a", "WangZJ", "作者名称")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "授权信息")
	rootCmd.Flags().StringP("source", "s", "", "来源")
}
