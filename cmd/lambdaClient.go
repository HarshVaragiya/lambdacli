package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HarshVaragiya/lambdacli/lib"
	"io/ioutil"
	"net/http"
)

type lambdaFnClient struct {
	httpClient    *http.Client
	serverBaseUrl string
}

func newLambdaFnClient(serverUrl string) *lambdaFnClient {
	httpClient := &http.Client{}
	return &lambdaFnClient{httpClient: httpClient, serverBaseUrl: serverUrl}
}

func (client *lambdaFnClient) createLambdaFunction(function *lib.Function) error {
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
	var response lib.LambdaRestApiResponse
	log.Debugf("attempting to unmarshal response body")
	respBytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		log.Errorf("error processing response. error = %v", err)
		log.Errorf("response data: %v", string(respBytes))
		return err
	}
	if resp.StatusCode == http.StatusCreated {
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

func (client *lambdaFnClient) invokeLambda(function *lib.Function, event *lib.Event) error {
	setLogLevel()
	log.Debugf("attempting to invoke lambda function [%s] on LambdaFn backend [%s]", function.Name, client.serverBaseUrl)
	methodUrl := fmt.Sprintf("%s/lambda/%s/invoke", client.serverBaseUrl, function.Name)
	log.Tracef("method URL: %s", methodUrl)
	eventJson, err := json.Marshal(event)
	if err != nil {
		log.Errorf("error converting event to JSON object. error = %v", err)
		return err
	}
	log.Debugf("event JSON : %s", string(eventJson))
	req, err := http.NewRequest(http.MethodGet, methodUrl, bytes.NewBuffer(eventJson))
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
	var response lib.Response
	log.Debugf("attempting to unmarshal lambda response body")
	respBytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		log.Errorf("error processing response. error = %v", err)
		log.Errorf("response data: %v", string(respBytes))
		return err
	}
	if resp.StatusCode == http.StatusOK {
		log.Printf("function [ %s ] invoked. response status code : %v ", function.Name, response.StatusCode)
		log.Printf("EventId : %v", response.EventId)
		log.Printf("Data    : %v", response.Data)
		log.Printf("Message : %v", response.Message)
		log.Printf("Stderr  : %v", response.Stderr)
		return nil
	} else {
		log.Warnf("something went wrong with the request for [%s]", function.Name)
		log.Warnf("EventId : %v , StatusCode : %v", response.EventId, response.StatusCode)
		log.Warnf("Message : %v", response.Message)
		log.Warnf("Stderr  : %v", response.Stderr)
		log.Warnf("Data    : %v", response.Data)
		return fmt.Errorf("error processing request")
	}
}

func (client *lambdaFnClient) deleteLambda(function *lib.Function) error {
	setLogLevel()
	log.Debugf("attempting to remove lambda function [%s] on LambdaFn backend [%s]", function.Name, client.serverBaseUrl)
	methodUrl := fmt.Sprintf("%s/lambda", client.serverBaseUrl)
	log.Tracef("method URL: %s", methodUrl)
	functionJson, err := json.Marshal(function)
	if err != nil {
		log.Errorf("unable to convert function to JSON object. error = %v", err)
		return err
	}
	log.Debugf("request JSON : %s", string(functionJson))
	req, err := http.NewRequest(http.MethodDelete, methodUrl, bytes.NewBuffer(functionJson))
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
	var response lib.LambdaRestApiResponse
	log.Debugf("attempting to unmarshal lambda response body")
	respBytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		log.Errorf("error processing response. error = %v", err)
		log.Errorf("response data: %v", string(respBytes))
		return err
	}
	if resp.StatusCode == http.StatusCreated {
		log.Printf("LambdaFn [ %s ] deleted. details :", function.Name)
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
