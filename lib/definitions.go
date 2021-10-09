package lib

import (
	"time"
)

type Function struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	CodeUri        string `json:"codeuri"`
	Handler        string `json:"handler"`
	Timeout        time.Duration
	TimeoutSeconds string            `json:"timeout-seconds"`
	Runtime        string            `json:"runtime"`
	Arn            string            `json:"arn"`
	Tags           map[string]string `json:"tags"`
	Env            map[string]string `json:"env"`
}

type LambdaRestApiResponse struct {
	Message          string `json:"message"`
	ErrorMessage     string `json:"error-message"`
	FunctionName     string `json:"function-name"`
	ModifiedResource string `json:"modified-resource"`
	DebugMessage     string `json:"debug-message"`
}

type Event struct {
	EventId   string `json:"eventId,omitempty"`
	EventData string `json:"eventData,omitempty"`
	Context   string `json:"context,omitempty"`
}

type Response struct {
	Data       string `json:"data,omitempty"`
	Stderr     string `json:"stderr,omitempty"`
	StatusCode int32  `json:"statusCode,omitempty"`
	Message    string `json:"message,omitempty"`
	EventId    string `json:"eventId,omitempty"`
}
