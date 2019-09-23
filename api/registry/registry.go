package registry

import (
	"instagram/api/infrastructure/api/handler"
	"instagram/api/infrastructure/database"
	"instagram/api/infrastructure/storage"
	"instagram/api/interface/controllers"
	"instagram/api/usecase/repository"
	"instagram/api/usecase/service"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/jinzhu/gorm"
)

// 依存解決用クラス
type interactor struct {
	db       *gorm.DB
	s3Config *aws.Config
}

// インターフェース
type Interactor interface {
	NewAppHandler() handler.AppHandler
}

// コンストラクタ
func NewInteractor(db *gorm.DB, s3Config *aws.Config) Interactor {
	return &interactor{db, s3Config}
}

// Appハンドラ
func (i *interactor) NewAppHandler() handler.AppHandler {
	return handler.NewAppHandler(i.NewUserHandler(), i.NewPostHandler(), i.NewLikeHandler(), i.NewAuthHandler())
}

// userハンドラ
func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserController())
}

// postハンドラ
func (i *interactor) NewPostHandler() handler.PostHandler {
	return handler.NewPostHandler(i.NewPostController(), i.NewUserController(), i.NewStorageController())
}

// likeハンドラ
func (i *interactor) NewLikeHandler() handler.LikeHandler {
	return handler.NewLikeHandler(i.NewLikeController(), i.NewPostController())
}

// authハンドラ
func (i *interactor) NewAuthHandler() handler.AuthHandler {
	return handler.NewAuthHandler(i.NewUserController())
}

// userコントローラ
func (i *interactor) NewUserController() controllers.UserController {
	return controllers.NewUserController(i.NewUserService())
}

// postコントローラ
func (i *interactor) NewPostController() controllers.PostController {
	return controllers.NewPostController(i.NewPostService())
}

// likeコントローラ
func (i *interactor) NewLikeController() controllers.LikeController {
	return controllers.NewLikeController(i.NewLikeService())
}

// storageコントローラ
func (i *interactor) NewStorageController() controllers.StorageController {
	return controllers.NewStorageController(i.NewStorageService())
}

// userサービス
func (i *interactor) NewUserService() service.UserService {
	return service.NewUserService(i.NewUserRepository())
}

// postサービス
func (i *interactor) NewPostService() service.PostService {
	return service.NewPostService(i.NewPostRepository())
}

// likeサービス
func (i *interactor) NewLikeService() service.LikeService {
	return service.NewLikeService(i.NewLikeRepository())
}

// storageサービス
func (i *interactor) NewStorageService() service.StorageService {
	return service.NewStorageService(i.NewStorageRepository())
}

// userリポジトリ
func (i *interactor) NewUserRepository() repository.UserRepository {
	return database.NewUserRepository(i.db)
}

// postリポジトリ
func (i *interactor) NewPostRepository() repository.PostRepository {
	return database.NewPostRepository(i.db)
}

// likeリポジトリ
func (i *interactor) NewLikeRepository() repository.LikeRepository {
	return database.NewLikeRepository(i.db)
}

// storageリポジトリ
func (i *interactor) NewStorageRepository() repository.StorageRepository {
	return storage.NewStorageRepository(i.s3Config)
}
