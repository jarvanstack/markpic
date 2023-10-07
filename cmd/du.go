/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// duCmd represents the markpic command
var duCmd = &cobra.Command{
	Use:   "du",
	Short: "先 d(下载) 再 u(上传)",
	Run: func(cmd *cobra.Command, args []string) {
		from := cmd.Flag("from").Value.String()
		dir := cmd.Flag("dir").Value.String()

		err := du(from, dir)
		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

func du(from, dir string) error {
	fmt.Println("[下载-上传]")

	toLocal := from + downloadFilePrefix
	err := d(from, toLocal, dir)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fromLocal := toLocal
	toRemote := from + uploadFilePrefix
	err = u(fromLocal, toRemote)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("[下载-上传完成]", fromLocal, toRemote)

	return nil
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
