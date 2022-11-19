/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// duCmd represents the du command
var duCmd = &cobra.Command{
	Use:   "du",
	Short: "du 一键下载 markdown 中图片, 并上传图片到图床并替换链接",
	Long: `du 一键下载 markdown 中图片, 并上传图片到图床并替换链接. For example:

du du --from README.md`,
	Run: func(cmd *cobra.Command, args []string) {
		from := cmd.Flag("from").Value.String()
		dir := cmd.Flag("dir").Value.String()
		fmt.Println("[下载-上传]")

		toLocal := from + downloadFilePrefix
		err := d(from, toLocal, dir)
		if err != nil {
			fmt.Println(err)
			return
		}

		fromLocal := toLocal
		toRemote := from + uploadFilePrefix
		err = u(fromLocal, toRemote)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("[下载-上传完成]", fromLocal, toRemote)
	},
}

func init() {
	rootCmd.AddCommand(duCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// duCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// duCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
