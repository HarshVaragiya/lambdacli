package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HarshVaragiya/LambdaFn/golambda"
	"github.com/HarshVaragiya/LambdaFn/liblambda"
	"io/ioutil"
	"net/http"
)

type lambdaFnClient struct {
	httpClient		*http.Client
	serverBaseUrl	string
}

func newLambdaFnClient(serverUrl string) *lambdaFnClient {
	httpClient := &http.Client{}
	return &lambdaFnClient{httpClient: httpClient, serverBaseUrl: serverUrl}
}

func (client *lambdaFnClient) createLambdaFunction(function *golambda.Function) error {
	setLogLevel()
	log.Debugf("attempting to create lambda function [%s] on LambdaFn backend [%s]", function.Name, client.serverBaseUrl)
	methodUrl := fmt.Sprintf("%s/lambda", client.serverBaseUrl)
	log.Tracef("method URL: %s", methodUrl)
	functionJson, err := json.Marshal(function)
	if err != nil {
		log.Errorf("unable to convert function to JSON object. error = %v", err)
		return err
	}
	log.Debugf("request JSON : %s", string(functionJson))
	req, err := http.NewRequest(http.MethodPut, methodUrl, bytes.NewBuffer(functionJson))
	if err != nil {
		log.Errorf("error creating the API request. error = %v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.httpClient.Do(req)
	if err != nil {
		log.Errorf("error with http response. error = %v", err)
		return err
	}
	var response liblambda.LambdaRestApiResponse
	log.Debugf("attempting to unmarshal response body")
	respBytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		log.Errorf("error processing response. error = %v", err)
		log.Errorf("response data: %v", string(respBytes))
		return err
	}
	if resp.StatusCode == http.StatusOK {
		log.Printf("LambdaFn [ %s ] created. details :", function.Name)
		log.Printf("Arn : %v", response.ModifiedResource)
		log.Printf("Msg : %v", response.Message)
		return nil
	} else {
		log.Warnf("something went wrong with the request for [%s]", function.Name)
		log.Warnf("Msg : %v", response.Message)
		log.Warnf("Err : %v", response.ErrorMessage)
		log.Warnf("Debug : %v", response.DebugMessage)
		return fmt.Errorf("error processing request")
	}
}
