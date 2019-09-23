package service

import (
	"instagram/api/domain/model"
	"instagram/api/usecase/repository"
)

type storageService struct {
	StorageRepository repository.StorageRepository
}

type StorageService interface {
	UploadFile(uploadImage *model.UploadImage) error
}

func NewStorageService(sr repository.StorageRepository) StorageService {
	return &storageService{sr}
}

// ファイルのアップロード
func (storageService *storageService) UploadFile(uploadImage *model.UploadImage) error {
	return storageService.StorageRepository.Upload(uploadImage)
}
