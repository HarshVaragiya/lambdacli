package cmd

import (
	"fmt"
	"github.com/HarshVaragiya/lambdacli/lib"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

func setLogLevel() {
	if debug {
		log.SetLevel(logrus.TraceLevel)
	}
}

func getLambdaFunction(cmd *cobra.Command) *lib.Function {

	functionName, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")
	codeUri, _ := cmd.Flags().GetString("code-uri")
	handler, _ := cmd.Flags().GetString("handler")
	timeout, _ := cmd.Flags().GetString("timeout")
	runtime, _ := cmd.Flags().GetString("runtime")

	return &lib.Function{
		Name:           functionName,
		Description:    description,
		CodeUri:        codeUri,
		Handler:        handler,
		TimeoutSeconds: timeout,
		Runtime:        runtime,
	}
}

func getLambdaEvent(cmd *cobra.Command) *lib.Event {
	eventData, _ := cmd.Flags().GetString("event")
	context, _ := cmd.Flags().GetString("context")
	return &lib.Event{EventData: eventData, Context: context}
}

func validateCreateLambda(function *lib.Function) error {
	if function.Name == "" || function.Handler == "" || function.CodeUri == "" || function.Runtime == "" {
		return fmt.Errorf("cannot create function with empty fields in [name, handler, code-uri, runtime]")
	} else if !strings.HasSuffix(strings.ToLower(function.CodeUri), ".zip") {
		return fmt.Errorf("code-uri has to be a location to a zip archive")
	}
	return nil
}

func validateFunctionBasics(function *lib.Function) error {
	if function.Name == "" {
		return fmt.Errorf("function name cannot be blank")
	}
	return nil
}

func validateEventBasics(event *lib.Event) error {
	if event.EventData == "" {
		return fmt.Errorf("event data cannot be blank")
	}
	return nil
}
