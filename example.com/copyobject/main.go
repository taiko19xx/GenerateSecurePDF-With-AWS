package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"context"

	"example.com/modules/types/event"
	"example.com/modules"

	"os"
	
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func handleRequest(ctx context.Context, evnt event.Event) (event.Event, error) {
    item := evnt.Path
    newItem := modules.GenRndPDFName()
    bucket := os.Getenv("BUCKET_NAME")

	newItemObjectKey := "public/" + newItem

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	svc := s3.New(sess)

	input := &s3.CopyObjectInput{
		Bucket: aws.String(bucket),
		CopySource: aws.String(bucket + "/" + item),
		Key: aws.String(newItemObjectKey),
	}

	_, err := svc.CopyObject(input)
	if err != nil {
		return evnt,  err
	}
	
	resp := event.Event{
		Email: evnt.Email,
		Path: newItemObjectKey,
	}

	return resp, nil
}

func main() {
	lambda.Start(handleRequest)
}