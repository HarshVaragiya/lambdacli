/*
Copyright © 2021 Harsh Varagiya <harsh8v@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create lambda function on the LambdaFn application server",
	Long:  `creating lambda function on the LambdaFn application server`,
	Run: func(cmd *cobra.Command, args []string) {
		function := getLambdaFunction(cmd)
		if err := validateCreateLambda(function); err != nil {
			log.Errorf("error validating create request. error = %v", err)
			return
		}
		client := newLambdaFnClient(serverUrl)
		log.Infof("attempting to create lambda function [%s]", function.Name)
		err := client.createLambdaFunction(function)
		if err != nil {
			log.Errorf("error creating lambda function")
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
