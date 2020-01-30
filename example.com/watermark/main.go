package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"context"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"

	"example.com/modules/types/event"
	"example.com/modules"
	"example.com/modules/downloader"
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

    onTop := true
	wm, _ := pdfcpu.ParseTextWatermarkDetails(evnt.Email, "rot:0, pos:bc, op:0.5, s:0.5 abs", onTop)
	api.AddWatermarksFile(itemPath, newItemLocalPath, nil, wm, nil)
	
	err = uploader.S3Upload(newItemLocalPath, bucket, newItemObjectKey)
    if err != nil {
		return evnt, err
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