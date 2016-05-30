// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"log"

	"github.com/blacktop/vm-proxy/drivers/virtualbox"
	"github.com/spf13/cobra"
)

// controlvmCmd represents the controlvm command
var controlvmCmd = &cobra.Command{
	Use:   "controlvm",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines`,
	Run: func(cmd *cobra.Command, args []string) {
		d := virtualbox.NewDriver("", "")
		outList, err := d.StopVM(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(outList)
	},
}

func init() {
	RootCmd.AddCommand(controlvmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// controlvmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// controlvmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}