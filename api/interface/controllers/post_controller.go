package controllers

import (
	"instagram/api/domain/model"
	"instagram/api/usecase/service"
)

// controlleはserviceに依存
type postController struct {
	postService service.PostService
}

// インターフェース宣言
type PostController interface {
	GetLastPostID() (uint, error)
	CreatePost(userID uint, imageUrl string, caption string) error
	GetPost(pageNum int) (*[]model.Post, error)
	GetPostCount() (int, error)
	GetPostByID(postID uint) (*model.Post, error)
	DeletePost(postID uint) error
}

// serviceに依存したcontrollerを生成
func NewPostController(ps service.PostService) PostController {
	return &postController{ps}
}

func (postController *postController) GetLastPostID() (uint, error) {
	post := &model.Post{}
	return postController.postService.GetLastPostID(post)
}

func (postController *postController) CreatePost(userID uint, imageUrl string, caption string) error {
	// usecaseで扱いやすい形に変換
	post := &model.Post{
		UserID:   userID,
		ImageURL: imageUrl,
		Caption:  caption,
	}
	return postController.postService.CreatePost(post)
}

// 指定したページの投稿を取得
func (postController *postController) GetPost(pageNum int) (*[]model.Post, error) {
	posts := &[]model.Post{}

	return postController.postService.GetPost(pageNum, posts)
}

// 投稿の数を取得
func (postController *postController) GetPostCount() (int, error) {
	posts := &[]model.Post{}
	var c int

	count, err := postController.postService.GetPostCount(posts, c)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (postController *postController) GetPostByID(postID uint) (*model.Post, error) {
	post := &model.Post{}

	return postController.postService.GetPostByID(post, postID)
}

func (postController *postController) DeletePost(postID uint) error {
	return postController.postService.DeletePost(postID)
}
