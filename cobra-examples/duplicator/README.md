## How to Implement CLI App Using Cobra?

Golang is great to build the Command Line Interface (CLI) app. In this tutorial, I'm going to present step by step how to implement that CLI app using Go lang and Cobra package.

I assume that You have already installed a Golang and set up a development environment for it. The library which is the best match in the case of the CLI app is Cobra.

Before we start using the Cobra library we need to get to know the concepts behind it.

"Cobra is built on a structure of commands, arguments & flags. Commands represent actions, Args are things and Flags are modifiers for those actions." - Cobra github

The structure will be like

```
APPNAME Command Args --flags
```

Example app execute in console

- docker ps -a

- Hugo new site example.com

If you want to learn golang, this tutorial will show you how you can tackle a specific challenge considering CLI apps, that as developers we often encounter. Enjoy and feel free to let me know in the comment section if anything requires further elaboration!

Troubleshooting

T01
```
Trouble:
../../../../spf13/viper/viper.go:40:2: cannot find package "github.com/hashicorp/hcl/printer"
Solution:
export GO111MODULE=on
```

T02

```
Trouble:
cannot find module providing package github.com/researchlab/gbp/cobra-examples/dupliator/cmd: working directory is not part of a module
Solution:
step 01
➜  dupliator git:(master) ✗ go mod init duplicator
go: creating new go.mod: module duplicator
➜  dupliator git:(master) ✗ go test
go: finding module for package github.com/researchlab/gbp/cobra-examples/duplicator/cmd
go: downloading github.com/researchlab/gbp v0.0.0-20210410032833-a8fb0b4f8872
main.go:17:8: module github.com/researchlab/gbp@latest found (v0.0.0-20210410032833-a8fb0b4f8872), but does not contain package github.com/researchlab/gbp/cobra-examples/duplicator/cmd

step02
➜  gbp git:(master) ✗ git add cobra-examples/
➜  gbp git:(master) ✗ git commit -m"duplicator cli"
[master e9abb80] duplicator cli
➜  gbp git:(master) ✗ git push origin master -vvv
➜  gbp git:(master) ✗ cd cobra-examples/duplicator
➜  duplicator git:(master) ✗ ls
LICENSE cmd     go.mod  go.sum  main.go
➜  duplicator git:(master) ✗ go build -o duplicator main.go && duplicator duplicate -f ext
go: finding module for package github.com/researchlab/gbp/cobra-examples/duplicator/cmd
go: downloading github.com/researchlab/gbp/cobra-examples/duplicator v0.0.0-20210411163813-e9abb809d735
go: found github.com/researchlab/gbp/cobra-examples/duplicator/cmd in github.com/researchlab/gbp/cobra-examples/duplicator v0.0.0-20210411163813-e9abb809d735
go: github.com/researchlab/gbp/cobra-examples/duplicator/cmd: github.com/researchlab/gbp/cobra-examples/duplicator@v0.0.0-20210411163813-e9abb809d735: parsing go.mod:
	module declares its path as: duplicator
	        but was required as: github.com/researchlab/gbp/cobra-examples/duplicator

step03
➜  duplicator git:(master) ✗ ls
LICENSE cmd     go.mod  go.sum  main.go
➜  duplicator git:(master) ✗ rm go.*
➜  duplicator git:(master) ✗ go mod init github.com/researchlab/gbp/cobra-examples/duplicator
go: creating new go.mod: module github.com/researchlab/gbp/cobra-examples/duplicator
➜  duplicator git:(master) ✗ go build -o duplicator main.go && ./duplicator duplicate -e ext
go: finding module for package github.com/spf13/cobra
go: finding module for package github.com/spf13/viper
go: downloading github.com/spf13/cobra v1.1.3
go: found github.com/spf13/cobra in github.com/spf13/cobra v1.1.3
go: found github.com/spf13/viper in github.com/spf13/viper v1.7.1

duplicate called
ext

```

Step 0:  Install cobra package

```
➜  dev go get -u github.com/spf13/cobra
```

Step 1: Creating the project in Golang


Cobra library provides initializator command for the new project. Let's init base CLI structure using cobra.

```
➜  src cobra init duplicator -b github.com/researchlab/gbp/cobra-examples/
Your Cobra application is ready at
/Users/lihong/workbench/dev/src/github.com/researchlab/gbp/cobra-examples/duplicator
Give it a try by going there and running `go run main.go`
Add commands to it by running `cobra add [cmdname]`
➜  src cd github.com/researchlab/gbp/cobra-examples/
➜  src pwd
/Users/lihong/workbench/dev/src
➜  src ls github.com/researchlab/gbp/cobra-examples/duplicator/
LICENSE  cmd/     main.go
```

Insider the project dir

```
➜  src cd github.com/researchlab/gbp/cobra-examples/duplicator
➜  dupliator git:(master) ✗ tree
.
├── LICENSE
├── cmd
│   └── root.go
└── main.go

1 directory, 3 files
```

Add new action duplicate

```
➜  dupliator git:(master) ✗ cobra add duplicate
duplicate created at /Users/lihong/workbench/dev/src/github.com/researchlab/gbp/cobra-examples/duplicator/cmd/duplicate.go
➜  dupliator git:(master) ✗ tree
.
├── LICENSE
├── cmd
│   ├── duplicate.go
│   └── root.go
└── main.go

1 directory, 4 files
```

The main.go is the entry point of the CLI. Inside the main.go it is calling the Execute function of the

```
➜  duplicator git:(master) ✗ cat main.go
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

package main

import "github.com/researchlab/gbp/cobra-examples/duplicator/cmd"

func main() {
	cmd.Execute()
}
```

Let's look to root.go The main part of root.go is our command rootCmd where we define it.

```
➜  duplicator git:(master) ✗ cat cmd/root.go
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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dupliator",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
// Uncomment the following line if your bare application
// has an action associated with it:
//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dupliator.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".dupliator") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
```

Build the app without defining any code for root command

```
go build -o  duplicator main.go && ./duplicator duplicate
```
As a result we get only print in terminal

```
duplicate called
```
Step 2: Let's define a flag for our duplicate command

Define var 'fileExt' to store file extension which kind of file we are going to copy.
```
var fileExt string

	// duplicateCmd represents the duplicate command
	var duplicateCmd = &cobra.Command{
```
Add definition flag to init() function in duplicate.go file. Look that the flag extension is required.
```
duplicateCmd.Flags().StringVarP(&fileExt, "extension", "e", "", "file extension is required")
duplicateCmd.MarkFlagRequired("extension")
```
And print fileExt var in Run section of our command
```
Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("duplicate called")
		fmt.Println(fileExt)
	},
```
Let's build and run it.
```
go build -o  duplicator main.go && ./duplicator duplicate -f txt
```
As result we get print in console
```
 duplicate called
  txt
```
 
Step 3: Let's define a new flag 'dirName' which store handle dir name where the file will be duplicated.
```
var dirName string

func init() {
	fmt.Println("init called")
	duplicateCmd.Flags().StringVarP(&fileExt, "extension", "e", "", "file extension is required")
	duplicateCmd.MarkFlagRequired("extension")
	duplicateCmd.PersistentFlags().StringVarP(&dirName, "dirname", "d", "copied_files", "dir to copie")
	rootCmd.AddCommand(duplicateCmd)
}
```
 
We see that the flag -d is not required and has default value copied_files for the dir.

Step 4: Let's add a new command which prints the version of our app
```
cobra add version
```
 
Add version function to version.go file.
```
func version(){
	fmt.Println("version called")
}
```
And call it in Run section.

Run: func(cmd *cobra.Command, args []string) { fmt.Println("version called") version() },
```
➜  duplicator git:(master) ✗ go build -o duplicator main.go && ./duplicator version
version called
```
Summary and conclusion
 
I have decided to use the Cobra package since it is really easy to integrate it with the base code of any application. The best in Cobra package is that it organizes commands into dedicated command file ie. duplicate.go The package also auto-generate man pages for your command. I see big potential in this package.

Tips

01 Parameters Usage
```
cat cmd/add.go
// 持久命令,意思是rootCmd命令和其子命令都能够使用这个参数
rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
// 本地命令,意思是只有rootCmd命令才能够使用该参数,如果其他命令设置了改参数则会报错
rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
// 标记 region 是必须的,如果没有该参数,则会报错
rootCmd.MarkFlagRequired("region")

// 其他命令的参数
addCmd.Flags().Int32VarP(&Num, "num", "n", 0, "Number")
```

02 Validator
```
//参数验证器
//cobra.NoArgs 如果任何位置有参数,就报错
//cobra.ArbitraryArgs 接收任何参数
//cobra.MinimumNArgs(int) 如果少于n个参数就报错
//cobra.MaximumNArgs(int) 如果大于n个参数就报错
//cobra.ExactArgs(int) 如果不等于n个参数就报错

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Args:  cobra.NoArgs,
```
