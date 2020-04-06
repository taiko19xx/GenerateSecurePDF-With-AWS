package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/pdfcpu/pdfcpu/pkg/api"

	"example.com/modules"
	"example.com/modules/downloader"
	"example.com/modules/types/event"
	"example.com/modules/uploader"

	"os"
)

func handleRequest(ctx context.Context, evnt event.Event) (event.Event, error) {
	item := evnt.Path
	newItem := modules.GenRndPDFName()
	bucket := os.Getenv("BUCKET_NAME")

	newItemLocalPath := "/tmp/" + newItem
	newItemObjectKey := "tmp/" + newItem

	itemPath, err := downloader.S3Download(bucket, item)
	if err != nil {
		return evnt, err
	}

	properties := map[string]string{"author": "Author name", "GenerateTo": evnt.Email}
	api.AddPropertiesFile(itemPath, newItemLocalPath, properties, nil)

	err = uploader.S3Upload(newItemLocalPath, bucket, newItemObjectKey)
	if err != nil {
		return evnt, err
	}

	resp := event.Event{
		Email: evnt.Email,
		Path:  newItemObjectKey,
	}

	return resp, nil
}

func main() {
	lambda.Start(handleRequest)
}
