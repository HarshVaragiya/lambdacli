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
	"time"
)

// invokeCmd represents the invoke command
var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "invoke a LambdaFn function",
	Long:  `invoke a LambdaFn function on the application server`,
	Run: func(cmd *cobra.Command, args []string) {
		function := getLambdaFunction(cmd)
		event := getLambdaEvent(cmd)
		if err := validateFunctionBasics(function); err != nil {
			log.Errorf("error validating create request. error = %v", err)
			return
		}
		if err := validateEventBasics(event); err != nil {
			log.Warnf("event looks invalid. proceeding in 5 seconds. error = %v", err)
			time.Sleep(time.Second * 5)
		}
		client := newLambdaFnClient(serverUrl)
		log.Infof("invoking lambda function [%s]", function.Name)
		err := client.invokeLambda(function, event)
		if err != nil {
			log.Errorf("error invoking lambda function")
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(invokeCmd)
}
