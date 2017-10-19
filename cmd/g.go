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

// gCmd represents the g command
var gCmd = &cobra.Command{
	Use:   "g",
	Short: "A brief description of your command",
	Long: `据给定的参数生成一个项目:
	tsp g dirnameForProjectA ClassForProjectA NameOfProjectA
	`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			tsp g [projectDirName] [projectClassName] [projectNameWithoutSpace]
		*/
		if len(args) < 3 {
			fmt.Println("参数不足")

		} else {

			GenerateProject(args)
		}
	},
}

func init() {
	RootCmd.AddCommand(gCmd)

}
