// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long:  `为指定目录下的项目运行一个简单的 HTTP Server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")

		serveProject(args)
	},
}

/*

多个命令包中公用的代码,该如何处理?
*/
func serveProject(args []string) {
	/*var (
		dirName string = filepath.Join(args[1], "dist")
		port    string = args[2]
	)

	valid, _ := isPortValid(port)

	if !valid {
		P(WARING, "Invalid port %v .", port)
	}

	if !IsPathExist(joinPath(dirName)) {

		P(WARING, "dir %v not exist.", dirName)
		return
	}
	if err := os.Chdir(dirName); err != nil {

		P(WARING, "切换工作目录时出错")
		return
	}

	cmd := exec.Command("http-server", "-p", port, "-a", "127.0.0.1", "--cors", "-o")

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Print("服务器已启动")
	fmt.Print(string(cmdOutput.Bytes()))*/

}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
