package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "hello命令简介",
	Long:  `hello命令详细介绍`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println(cmd.Flag("author").Value)
		fmt.Println(cmd.Flag("name").Value)
	},
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(helloCmd)
	//本地flag
	var Source string
	helloCmd.Flags().StringVarP(&Source, "source", "s", "", "读取文件路径")



}

