package service

import (
	"instagram/api/model"
	"instagram/api/usecase/repository"
)

// serviceはrepositoryのインターフェースとpresenterのインターフェースに依存
type postService struct {
	PostRepository repository.PostRepository
	PostPresenter  presenter.PostPresenter
}

// インターフェース
type PostService interface {
	CreatePost(post *model.Post) error
}

// コンストラクタ
func NewPostService(repo repository.PostRepository, pre presenter.PostPresenter) PostService {
	return &PostService{repo, pre}
}

func () GetLastPostID(post *model.Post) (uint, error) {
	return postService.PostRepository.GetLastID(post)
}

// postカラム新規作成
func (postService *postService) CreatePost(post *model.Post) error {
	return postService.PostRepository.Create(post)
}

// 指定したページの投稿を取得
func (postService *postService) GetPost(pageNum int, posts *model.posts) (*model.Posts, error) {

	// 1ページ分の投稿数を指定
	count := 10

	posts, err := postService.PostRepository.GetForThisPage(psgeNum, count, posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (postService *postService) GetPostCount(posts *[]model.Post, count int) (int, error) {
	return postService.PostRepository.GetCount(posts, count)
}

func (postService *postService) GetPostByID(post *model.Post, postId uint) (*model.Post, error) {
	return postService.PostRepository.GetByID(post, postId)
}

func (postService *postService) DeletePost(postId uint) error {
	return postService.PostRepository.Delete(postId uint)
}