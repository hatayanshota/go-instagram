package service

type strageService struct {
	storageRepository repository.storageRepository
}

type StorageService interface {
}

// ファイルのアップロード
func (strageService *strageService) UploadFile(uploadFile *model.UploadFile) error {
	return storageService.storageRepository.Upload(uploadFile)
}
