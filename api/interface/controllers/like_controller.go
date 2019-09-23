package controllers

import (
	"instagram/api/domain/model"
	"instagram/api/usecase/service"
)

// controlleはserviceに依存
type likeController struct {
	likeService service.LikeService
}

// インターフェース宣言
type LikeController interface {
	CreateLike(userID uint, postID uint) error
	DeleteLike(userID uint, postID uint) error
}

// serviceに依存したcontrollerを生成
func NewLikeController(ls service.LikeService) LikeController {
	return &likeController{ls}
}

// 新規いいね
func (likeController *likeController) CreateLike(userID uint, postID uint) error {
	like := &model.Like{
		UserID: userID,
		PostID: postID,
	}

	return likeController.likeService.CreateLike(like)
}

// いいね削除
func (likeController *likeController) DeleteLike(userID uint, postID uint) error {
	like := &model.Like{}
	return likeController.likeService.DeleteLike(like, userID, postID)
}
