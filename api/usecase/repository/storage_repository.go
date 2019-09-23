package repository

import "instagram/api/domain/model"

type StorageRepository interface {
	Upload(uploadImage *model.UploadImage) error
}
