package downloader

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"example.com/modules"

	"fmt"
	"os"
)

func S3Download(bucket, key string) (string, error) {
	localPath := "/tmp/" + modules.GenRndName()

	file, err := os.Create(localPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
	if err != nil {
		return "", err
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	return localPath, nil
}
