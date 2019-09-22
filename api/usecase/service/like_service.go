package service

import (
	"instagram/api/model"
	"instagram/api/usecase/repository"
)

// serviceはrepositoryのインターフェースとpresenterのインターフェースに依存
type likeService struct {
	LikeRepository repository.LikeRepository
	LikePresenter  presenter.LikePresenter
}

// インターフェース
type LikeService interface {
	CreateLike(like *model.Like) error
	DeleteLike(like *model.Like, userId uint, postId uint) error
}

// コンストラクタ
func NewLikeService(repo repository.LikeRepository, pre presenter.LikePresenter) LikeService {
	return &LikeService{repo, pre}
}

// 新規いいね
func (likeService *likeService) CreateLike(like *model.Like) error {
	return likeService.LikeRepository.Create(like)
}

// いいね削除
func (likeService *likeService) DeleteLike(like *model.Like, userId uint, postId uint) error {
	return likeService.LikeRepository.Delete(like, userId, postId)
}
