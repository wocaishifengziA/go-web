/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// gogo, err := cmd.Flags().GetBool("toggle")
		// if err != nil {
		// 	fmt.Println("gogo:", err)
		// }
		// fmt.Println("ok gogo: ", gogo)
		fstatus, _ := cmd.Flags().GetBool("float")
		if fstatus { // 如果为 true，则调用 floatAdd 函数
			floatAdd(args)
		} else {
			intAdd(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}

func intAdd(args []string) {
	var sum int
	// 循环 args 参数，循环的第一个值为 args 的索引，这里我们不需要，所以用 _ 忽略掉
	for _, ival := range args {
		// 将 string 转换成 int 类型
		temp, err := strconv.Atoi(ival)
		if err != nil {
			panic(err)
		}
		sum = sum + temp
	}
	fmt.Printf("Addition of numbers %s is %d\n", args, sum)
}

func floatAdd(args []string) {
	var sum float64
	for _, fval := range args {
		// 将字符串转换成 float64 类型
		temp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			panic(err)
		}
		sum = sum + temp
	}
	fmt.Printf("Sum of floating numbers %s is %f\n", args, sum)
}
