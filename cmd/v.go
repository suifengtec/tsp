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
	"github.com/spf13/cobra"
	. "tsp/tsputils"
)

// vCmd represents the v command
var vCmd = &cobra.Command{
	Use:   "v",
	Short: "当前版本",
	Long: `获取当前版本:
	tsp v
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ShowVersion()
	},
}

func init() {
	RootCmd.AddCommand(vCmd)

}
