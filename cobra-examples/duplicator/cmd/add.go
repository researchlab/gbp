// Copyright © 2021 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

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
		// TODO: Work your own magic here
		fmt.Println("add called")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 持久命令,意思是rootCmd命令和其子命令都能够使用这个参数
	//	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	//	// 本地命令,意思是只有rootCmd命令才能够使用该参数,如果其他命令设置了改参数则会报错
	//	rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	//	rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
	//	// 标记 region 是必须的,如果没有该参数,则会报错
	//	rootCmd.MarkFlagRequired("region")
	//
	//	// 其他命令的参数
	//	addCmd.Flags().Int32VarP(&Num, "num", "n", 0, "Number")
}
