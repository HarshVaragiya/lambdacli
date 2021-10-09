/*
Copyright © 2021 Harsh Varagiya

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
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var log = logrus.New()
var serverUrl string
var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lambdacli",
	Short: "A cli to interact with LambdaFn application server",
	Long:  `A cli to allow user to create/invoke/removing a lambda function on the LambdaFn application server`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&serverUrl, "server", "http://127.0.0.1:8080","LambdaFn Application Server Base URL")
	rootCmd.PersistentFlags().BoolVarP(&debug,"debug","v", false, "enabled debug and trace messages")

	rootCmd.PersistentFlags().StringP("function-name","n","lambda","lambda function name")
	rootCmd.PersistentFlags().StringP("description","d","","lambda function description")
	rootCmd.PersistentFlags().StringP("code-uri","c","","location of code URI on disk (absolute path)")
	rootCmd.PersistentFlags().String("handler","lambda_function.lambda_handler","lambda handler")
	rootCmd.PersistentFlags().String("timeout","3s","lambda function timeout")
	rootCmd.PersistentFlags().String("runtime","python3","lambda runtime")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lambdacli.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".lambdacli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".lambdacli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
