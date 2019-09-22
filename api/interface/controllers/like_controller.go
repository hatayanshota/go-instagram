package controllers

// controlleはserviceに依存
type likeController struct {
	likeService service.LikeService
}

// インターフェース宣言
type LikeController interface {
	CreateLike(userId uint, postId uint) error
	DeleteLike(userId uint, postId uint) error
}

// serviceに依存したcontrollerを生成
func NewLikeController(ls service.LikeService) LikeController {
	return &likeController{ls}
}

// 新規いいね
func (likeController *likeController) CreateLike(userId uint, postId uint) error {
	like := &model.Like{
		UserID: userId,
		PostID: postId
	}

	return likeController.likeService.CreateLike(like)
}

// いいね削除
func (likeController *likeController) DeleteLike(userId uint, postId uint) error {
	like := &Like{}
	return likeController.LikeService.DeleteLike(like, userId, postId)
}