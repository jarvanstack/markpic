/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dengjiawen8955/du/tools/regs"
	"github.com/dengjiawen8955/du/tools/upload"

	"github.com/spf13/cobra"
)

// uCmd represents the u command
var uCmd = &cobra.Command{
	Use:   "u",
	Short: "将 markdown 中的所有本地图片通过 picgo 上传到图床",
	Long: `将 markdown 中的所有本地图片通过 picgo 上传到图床. For example:

du u --from README.md `,
	Run: func(cmd *cobra.Command, args []string) {
		from := cmd.Flag("from").Value.String()
		fmt.Println("[上传] ", from)

		to := from + uploadFilePrefix
		err := u(from, to)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("[上传完成]", from, to)
	},
}

func u(from, to string) error {
	// 输入
	formFile, err := os.Open(from)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer formFile.Close()
	fromBuf := bufio.NewReader(formFile)

	// 输出
	// 输出
	toFile, err := os.OpenFile(to, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer toFile.Close()
	toBuf := bufio.NewWriter(toFile)

	// 下载器
	uploader := upload.NewUploader()

	// 读取
	for {
		line, err := fromBuf.ReadString('\n')
		if err != nil {
			break
		}

		// 获取 URL
		urls := regs.GetLocalImg(line)
		for _, url := range urls {
			// 下载
			newUrl, err := uploader.Upload(url)
			if err != nil {
				fmt.Println(err)
				return err
			}

			// 替换
			line = strings.ReplaceAll(line, url, newUrl)
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
