package cmd

import (
	"fmt"
	"github.com/HarshVaragiya/LambdaFn/golambda"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func setLogLevel(){
	if debug {
		log.SetLevel(logrus.TraceLevel)
	}
}

func getLambdaFunction(cmd *cobra.Command) *golambda.Function {

	functionName, _ := cmd.Flags().GetString("function-name")
	description, _ := cmd.Flags().GetString("description")
	codeUri, _ := cmd.Flags().GetString("code-uri")
	handler, _ := cmd.Flags().GetString("handler")
	timeout, _ := cmd.Flags().GetString("timeout-seconds")
	runtime, _ := cmd.Flags().GetString("runtime")

	return &golambda.Function{
		Name:           functionName,
		Description:    description,
		CodeUri:        codeUri,
		Handler:        handler,
		TimeoutSeconds: timeout,
		Runtime:        runtime,
	}
}

func ValidateCreateLambda(function *golambda.Function) error {
	if function.Name == "" || function.Handler == "" || function.CodeUri == "" || function.Runtime == "" {
		return fmt.Errorf("cannot create function with empty fields in [name, handler, code-uri, runtime]")
	}
	return nil
}
