package jackalope_service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pashura/design-to-wf/api/properties"
	"os"
	"strings"
)

func S3Service(keys ...string) {

	// Initial credentials loaded from SDK's default credential chain, such as
	// the environment, shared credentials (~/.aws/credentials)
	// Role. These credentials are used to make the AWS STS Assume Role API.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	downloader := s3manager.NewDownloader(sess)

	for i := 0; i < len(keys); i++ {
		var filename string
		var bucket string

		switch {
		case strings.HasPrefix(keys[i], "XSD"):
			filename = "api/jackalope_service/enums.json"
			bucket = properties.S3bucketMasterTempletes
		case strings.HasPrefix(keys[i], "txn"):
			filename = "api/jackalope_service/schema.xsd"
			bucket = properties.S3bucketJackalope
		default:
			filename = ""
			bucket = ""
		}

		f, err := os.Create(filename)
		if err != nil {
			fmt.Printf("failed to create file %q, %v", filename, err)
		}

		operationDone := make(chan bool)
		go func() {
			_, err := downloader.Download(f, &s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(keys[i]),
			})
			if err != nil {
				fmt.Printf("failed to download file, %v", err)
			}

			operationDone <- true
		}()

		<-operationDone

		fmt.Printf("file %v downloaded\n", filename)
	}

}
