package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	// Sprintf formats according to a format specifier and returns the resulting string
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)} , nil 
}

func main() {
	// Start takes a handler and talks to an internal Lambda endpoint to pass requests to the handler - which is a function
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(HandleLambdaEvent)
}
