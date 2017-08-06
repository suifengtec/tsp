// Copyright © 2017 suifengtec <suifengtec@qq.com>
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

// wCmd represents the w command
var wCmd = &cobra.Command{
	Use:   "w",
	Short: "打包项目,并检测文件变化",
	Long: `打包项目,并检测文件变化:
	tsp w dirnameForProjectA
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("参数不足")

		} else {

			WatchProject(args)
		}
	},
}

func init() {
	RootCmd.AddCommand(wCmd)
}
