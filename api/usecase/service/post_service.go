package service

import (
	"instagram/api/domain/model"
	"instagram/api/usecase/repository"
)

// serviceはrepositoryのインターフェースとpresenterのインターフェースに依存
type postService struct {
	postRepository repository.PostRepository
}

// インターフェース
type PostService interface {
	GetLastPostID(post *model.Post) (uint, error)
	CreatePost(post *model.Post) error
	GetPost(pageNum int, posts *[]model.Post) (*[]model.Post, error)
	GetPostCount(posts *[]model.Post, count int) (int, error)
	GetPostByID(post *model.Post, postID uint) (*model.Post, error)
	DeletePost(postID uint) error
}

// コンストラクタ
func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo}
}

func (postService *postService) GetLastPostID(post *model.Post) (uint, error) {
	return postService.postRepository.GetLastID(post)
}

// postカラム新規作成
func (postService *postService) CreatePost(post *model.Post) error {
	return postService.postRepository.Create(post)
}

// 指定したページの投稿を取得
func (postService *postService) GetPost(pageNum int, posts *[]model.Post) (*[]model.Post, error) {

	// 1ページ分の投稿数を指定
	count := 10

	var offset = 0
	if pageNum != 0 {
		offset = (pageNum - 1) * 10
	}

	posts, err := postService.postRepository.GetForThisPage(count, offset, posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (postService *postService) GetPostCount(posts *[]model.Post, count int) (int, error) {
	return postService.postRepository.GetCount(posts, count)
}

func (postService *postService) GetPostByID(post *model.Post, postID uint) (*model.Post, error) {
	return postService.postRepository.GetByID(post, postID)
}

func (postService *postService) DeletePost(postID uint) error {
	return postService.postRepository.Delete(postID)
}
