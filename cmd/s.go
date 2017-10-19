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
	. "tsp/tsputils"
)

// sCmd represents the s command
var sCmd = &cobra.Command{
	Use:   "s",
	Short: "为指定项目在指定端口启用一个HTTP Serve,并在默认浏览器打开",
	Long: ` 为指定项目在指定端口启用一个HTTP Serve,并在默认浏览器打开:
	
	tsp s dirnameForProjectA portForProjectA
	 `,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("参数不足")

		} else {

			ServeProject(args)
		}
	},
}

func init() {
	RootCmd.AddCommand(sCmd)

}
