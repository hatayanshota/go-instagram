package service

import (
	"instagram/api/domain/model"
	"instagram/api/usecase/repository"
)

// serviceはrepositoryのインターフェースとpresenterのインターフェースに依存
type likeService struct {
	likeRepository repository.LikeRepository
}

// インターフェース
type LikeService interface {
	CreateLike(like *model.Like) error
	DeleteLike(like *model.Like, userID uint, postID uint) error
}

// コンストラクタ
func NewLikeService(repo repository.LikeRepository) LikeService {
	return &likeService{repo}
}

// 新規いいね
func (likeService *likeService) CreateLike(like *model.Like) error {
	return likeService.likeRepository.Create(like)
}

// いいね削除
func (likeService *likeService) DeleteLike(like *model.Like, userID uint, postID uint) error {
	return likeService.likeRepository.Delete(like, userID, postID)
}
