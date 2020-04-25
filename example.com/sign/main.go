package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/signintech/gopdf"

	"example.com/modules"
	"example.com/modules/downloader"
	"example.com/modules/types/event"
	"example.com/modules/uploader"

	"os"
	"path/filepath"
)

func handleRequest(ctx context.Context, evnt event.Event) (event.Event, error) {
	item := evnt.Path
	newItem := modules.GenRndPDFName()
	bucket := os.Getenv("BUCKET_NAME")

	newItemLocalPath := "/tmp/" + newItem
	newItemObjectKey := "tmp/" + newItem
	signFilePath := "/tmp/" + modules.GenRndPDFName()

	itemPath, err := downloader.S3Download(bucket, item)
	if err != nil {
		return evnt, err
	}

	executablePath := ""
	executablePath, err = os.Executable()
	if err != nil {
		return evnt, err
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA5})
	pdf.AddPage()
	err = pdf.AddTTFFont("Tanuki", filepath.Dir(executablePath) + "/TanukiMagic.ttf")
	if err != nil {
		return evnt, err
	}
	err = pdf.SetFont("Tanuki", "", 12)
	if err != nil {
		return evnt, err
	}
	pdf.Cell(nil, "次の利用者によってダウンロードされました")
	pdf.Br(20)
	pdf.Cell(nil, evnt.Email)
	pdf.WritePdf(signFilePath)

	inFiles := []string{itemPath, signFilePath}

	api.MergeFile(inFiles, newItemLocalPath, nil)

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
