/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra-app",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello app!")
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")             // 仅当前命令可用
	rootCmd.PersistentFlags().BoolP("gogo", "g", false, "Help message for persistent") // 当前及其子命令
	cobra.OnInitialize(initConfig1, initConfig2)
}

func initConfig1() {
	fmt.Println("initConfig 1")
}

func initConfig2() {
	fmt.Println("initConfig 2")
}

var cfgFile = ""

func InitConfig() {
	if cfgFile != "" {
		// 使用 flag 标志中传递的配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 获取 Home 目录
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// 在 Home 目录下面查找名为 ".my-calc" 的配置文件
		viper.AddConfigPath(home)
		viper.SetConfigName(".my-calc")
	}
	// 读取匹配的环境变量
	viper.AutomaticEnv()
	// 如果有配置文件，则读取它
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
