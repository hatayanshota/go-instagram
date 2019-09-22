package storage

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func NewS3() *aws.Config {

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_KEY"), ""),
		Region:           aws.String("ap-northeast-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}

	return s3Config
}
