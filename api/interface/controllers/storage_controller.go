package controllers

import (
	"instagram/api/domain/model"
	"instagram/api/usecase/service"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
)

type storageController struct {
	storageService service.StorageService
}

type StorageController interface {
	UploadFile(imagefile io.Reader, id string, contenType string) error
}

func NewStorageController(ss service.StorageService) StorageController {
	return &storageController{ss}
}

func (storageController *storageController) UploadFile(imagefile io.Reader, id string, contentType string) error {
	uploadImage := &model.UploadImage{
		ACL:         aws.String("public-read"),
		ContentType: aws.String(contentType),
		Bucket:      aws.String(os.Getenv("BUCKET_NAME")),
		Key:         aws.String("go-instagram/" + id),
		Body:        imagefile,
	}
	return storageController.storageService.UploadFile(uploadImage)
}
