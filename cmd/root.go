/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	downloadFilePrefix = ".download.md"
	uploadFilePrefix   = ".upload.md"
)

var (
	from string
	dir  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "du [command][flags]",
	Short: "markdown 图片下载上传工具.",
	Long: `markdown 图片下载上传工具. 例如:

# 一键下载 markdown 中图片, 并上传图片到图床并替换链接
du test.md`,
	Args: cobra.MatchAll(cobra.ArbitraryArgs), // 任意参数
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var from, dir string

		if len(args) > 0 {
			from = args[0]
		} else {
			from = cmd.Flag("from").Value.String()
		}

		if len(args) > 1 {
			dir = args[1]
		} else {
			dir = cmd.Flag("dir").Value.String()
		}

		err := du(from, dir)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&from, "from", "f", "source.md",
		"需要处理的源文件")
	rootCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "images/",
		"图片存放的目录")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
