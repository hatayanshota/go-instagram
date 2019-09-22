package controllers

import (
	"instagram/api/model"
	"instagram/api/usecase/usecase/service"
)

// controlleはserviceに依存
type postController struct {
	postService service.PostService
}

// インターフェース宣言
type PostController interface {
	CreatePost(post *model.Post) error
}

// serviceに依存したcontrollerを生成
func NewPostController(ps service.PostService) PostController {
	return &postController{ps}
}

func (postController *postController) GetLastPostID() (uint, error) {
	post := &model.Post{}
	return postController.postService.GetLastPostID(post)
}

func (postController *postController) CreatePost(userId, imageUrl, caption) error {
	// usecaseで扱いやすい形に変換
	post := &model.Post{
		UserID: userId
		ImageURL: imageUrl
		Caption: caption
	}
	return postController.PostService.CreatePost(post)
}

// 指定したページの投稿を取得
func (postController *postController) GetPost(pageNum int) (*model.Post, error) {
	posts := &[]model.Post{}

	return postController.PostService.GetPost(pageNum, posts)
}

// 投稿の数を取得
func (postController *postController) GetPostCount() (int, error) {
	posts := &[]model.Post{}
	var c int
	
	count, err := postController.PostService.GetPostCount(posts, c)
	if err != nil {
		return nil, err
	}

	return count, nil
}

func (postController *postController) GetPostByID(postId uint) (*model.Post, error) {
	post := &model.Post{}

	return postController.PostService.GetPostByID(post, postId)
}

func (postController *postController) DeletePost(postId uint) error {
	return postController.PostService.DeletePost(postId)
}