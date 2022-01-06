package main

import (
	"TestLambda/greeting" //自作のTestLambda Moduleの、greeting Packageを使うということ

	"github.com/aws/aws-lambda-go/lambda"
)

func executeFunction() {
	greeting.SayHello()
}

func main() {
	lambda.Start(executeFunction)
}
