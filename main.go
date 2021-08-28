package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"poring/pkg"
)

type Item struct {
	SensorId  string  `dynamodbav:"sensorID" json:"sensorID"`
	TimeStamp int     `dynamodbav:"timestamp" json:"timestamp"`
	Co2       int     `dynamodbav:"co2" json:"co2"`
	Temp      float64 `dynamodbav:"temp" json:"temp"`
	Hum       float64 `dynamodbav:"hum" json:"hum"`
}

type Response struct {
	Co2       int     `json:"co2"`
	Temp      float64 `json:"temp"`
	Hum       float64 `json:"hum"`
	Message  string `json:"message"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	pathparam := request.PathParameters["deviceID"]
	// DB接続
	svc := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	getParamPerson := &dynamodb.QueryInput{
		TableName: aws.String("SensorData"),
		ExpressionAttributeNames: map[string]*string{
			"#ID": aws.String("sensorID"), // alias付けれたりする
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":sensorID": {
				S: aws.String(pathparam),
			},
		},
		KeyConditionExpression: aws.String("#ID = :sensorID"),
		ScanIndexForward:       aws.Bool(false),
		Limit:                  aws.Int64(1),
	}

	// 検索
	getData, err := svc.Query(getParamPerson)
	if err != nil {
		fmt.Println("[Query Error]", err)
		return events.APIGatewayProxyResponse{
				Body:       err.Error(),
				StatusCode: 500,
			},
			err
	}

	sensorData := []Item{}
	err = dynamodbattribute.UnmarshalListOfMaps(getData.Items, &sensorData)
	if err != nil {
		fmt.Println(err.Error())
	}

	comfortLevel := pkg.CheckComfortLevel(sensorData[0].Temp,sensorData[0].Hum)

	co2Level:= pkg.CheckCo2Level(sensorData[0].Co2)

	message := pkg.CreateMessage(comfortLevel,co2Level)


	response := Response{
		Temp: sensorData[0].Temp,
		Hum: sensorData[0].Hum,
		Co2: sensorData[0].Co2,
		Message: message,
	}

	responseJson, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
			"Content-Type":                 "application/json",
		},
		Body:       string(responseJson),
		StatusCode: 200,
	}, nil

}
func main() {
	lambda.Start(handler)
}
