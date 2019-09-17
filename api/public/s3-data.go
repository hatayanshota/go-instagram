package public

import (
	_ "image/png"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//ファイルのアップロード
func UploadFile(imagefile io.Reader, id string, content_type string) error {

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_KEY"), ""),
		Region:           aws.String("ap-northeast-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}

	sess := session.Must(session.NewSession(s3Config))

	uploader := s3manager.NewUploader(sess)

	bucket := aws.String(os.Getenv("BUCKET_NAME"))

	cparams := &s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		ContentType: aws.String(content_type),
		Bucket:      bucket,
		Key:         aws.String("go-instagram/" + id),
		Body:        imagefile,
	}

	_, err := uploader.Upload(cparams)

	return err
}