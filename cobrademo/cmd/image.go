package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)

//todo  go run main.go  image p1 p2 p3
//todo  go run main.go  image  times  -t=3  p1  p2  p3

//定义一个imageCmd命令行
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Print images information",
	Long: "Print all images information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("image one is ubuntu 16.04")
		fmt.Println("image two is ubuntu 18.04")
		fmt.Println("image args are : " + strings.Join(args, " "))
	},
}

//定义一个cmdTimes命令行
//@todo init中将cmdTimes命令添加为imageCmd的子命令
var echoTimes int //循环打印多少次
var cmdTimes = &cobra.Command{
	Use:   "times [string to echo]",
	Short: "Echo anything to the screen more times",
	Long: `echo things multiple times back to the user by providing a count and a string.`,
	Args: cobra.MinimumNArgs(1), //至少要有 1 个位置参数，否则报错
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < echoTimes; i++ {
			fmt.Println("Echo: " + strings.Join(args, " "))
		}
	},
}






func init() {
	rootCmd.AddCommand(imageCmd)

	//这里配置flags以及命令行的配置信息
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input") //p *int, name, shorthand, default value, usage
	imageCmd.AddCommand(cmdTimes)



	////支持永久性标志，该标志将对此命令和所有子命令起作用，例如：
	//imageCmd.PersistentFlags().String("foo", "", "A help for foo")
	////支持仅在直接调用此命令时运行的本地标志，例如：
    //imageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
