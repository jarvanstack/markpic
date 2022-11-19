/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"du/tools/download"
	"du/tools/regs"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// uCmd represents the u command
var uCmd = &cobra.Command{
	Use:   "u",
	Short: "将 markdown 中的本地图片通过 picgo 上传到图床",
	Long: `将 markdown 中的本地图片通过 picgo 上传到图床. For example:

du u --from README.md --picgo C:\picgo.exe`,
	Run: func(cmd *cobra.Command, args []string) {
		from := cmd.Flag("from").Value.String()
		picgo := cmd.Flag("picgo").Value.String()
		fmt.Println("[上传] ", from, picgo)

		err := u(from, picgo)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("[上传完成]", from, picgo)
	},
}

func u(from, picgo string) error {
	// 输入
	formFile, err := os.Open(from)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer formFile.Close()
	fromBuf := bufio.NewReader(formFile)

	// 输出
	toFile, err := os.OpenFile(from+downloadFilePrefix, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer toFile.Close()
	toBuf := bufio.NewWriter(toFile)

	// 下载器
	downloader := download.NewDownLoader(dir)

	// 读取
	for {
		line, err := fromBuf.ReadString('\n')
		if err != nil {
			break
		}

		// 获取 URL
		urls := regs.GetRemoteImg(line)
		if len(urls) > 0 {
			for _, url := range urls {
				// 下载
				newUrl, err := downloader.DownLoad(url)
				if err != nil {
					fmt.Println(err)
					return err
				}

				// 替换
				line = strings.ReplaceAll(line, url, newUrl)
			}
		}

		toBuf.WriteString(line)
	}
	toBuf.Flush()

	return nil
}

func init() {
	rootCmd.AddCommand(uCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
