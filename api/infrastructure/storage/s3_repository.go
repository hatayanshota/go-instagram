package storage

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type storageRepository struct {
	s3Config *aws.Config
}

type StorageRepository interface {
	UploadFile(imagefile io.Reader, id string, contenType string) error
}

// s3にファイルをアップロード
func (storageRepository *storageRepository) UploadFile(uploadImage *model.UploadImage) error {

	sess := session.Must(session.NewSession(storageRepository.s3Config))

	uploader := s3manager.NewUploader(sess)

	cparams := &s3manager.UploadInput{
		ACL:         uploadImage.ACL,
		ContentType: uploadImage.ContentType,
		Bucket:      uploadImage.Bucket,
		Key:         uploadImage.Key,
		Body:        uploadImage.Body,
	}

	_, err := uploader.Upload(cparams)

	return err
}
