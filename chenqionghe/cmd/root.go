package cmd
import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)
//命令行的定义
var rootCmd = &cobra.Command{
	Use:   "chenqionghe",
	Short: "getting muscle is not easy",
	Long: `let's do it, yeah buddy light weight baby!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello chenqionghe")
	},
}

//绑定配置
//添加一个initConfig方法
func initConfig() {
	fmt.Println("initConfig")
	//获取项目的执行路径
	path,_:= os.Getwd()
	println(path)

	viper.AddConfigPath("./")
	viper.AddConfigPath("./conf")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}


//todo 选项的配置就需要在init方法中设置了
func init(){
	cobra.OnInitialize(initConfig) //这会在运行每个子命令之前运行

	//选项author
	var Author string
	rootCmd.PersistentFlags().StringVar(&Author, "author", "defaultAuthor", "作者名")
	//这将把viper配置和flag绑定，如果用户不设置-author选项，将从配置中查找
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))

	//选项verbose
	var Verbose  bool
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "全局版本")

	//选项name
	var Name  string
	rootCmd.Flags().StringVarP(&Name, "name", "n", "", "你的名字")
	rootCmd.MarkFlagRequired("name")

}


//应用入口会执行rootCmd的Execute方法额
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}