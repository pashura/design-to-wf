package jackalope_service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func S3Service(filename string, bucket string, key string) {

	// Initial credentials loaded from SDK's default credential chain, such as
	// the environment, shared credentials (~/.aws/credentials)
	// Role. These credentials are used to make the AWS STS Assume Role API.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	downloader := s3manager.NewDownloader(sess)

	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("failed to create file %q, %v", filename, err)
	}

	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Printf("failed to download file, %v", err)
	}
	fmt.Printf("file %v downloaded, %d bytes\n", filename, n)
}
