package handler

import (
	"fmt"
	"instagram/api/infrastructure/utils"
	"instagram/api/interface/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type likeHandler struct {
	likeController controllers.LikeController
	postController controllers.PostController
}

type LikeHandler interface {
	CreateLike(c echo.Context) error
	DeleteLike(c echo.Context) error
	GetLike(c echo.Context) error
}

func NewLikeHandler(lc controllers.LikeController, pc controllers.PostController) LikeHandler {
	return &likeHandler{lc, pc}
}

// モデルとは異なるLikeを設定
type LikeData struct {
	UserID uint `json:"user_id" form:"user_id" query:"user_id"`
	PostID uint `json:"post_id" form:"post_id" query:"post_id"`
}

// いいねしたユーザーのidと投稿のidが送られることが前提
func (likeHandler *likeHandler) CreateLike(c echo.Context) error {
	fmt.Println("呼ばれたよ")

	like := new(LikeData)
	if err := c.Bind(like); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if err := likeHandler.likeController.CreateLike(like.UserID, like.PostID); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

// いいね削除
func (likeHandler *likeHandler) DeleteLike(c echo.Context) error {

	like := new(LikeData)
	if err := c.Bind(like); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if err := likeHandler.likeController.DeleteLike(like.UserID, like.PostID); err != nil {
		return c.NoContent(http.StatusOK)
	} else {
		return c.NoContent(http.StatusInternalServerError)
	}
}

// 投稿のid(post_id)に対するいいねしたユーザーの一覧を返す
func (likeHandler *likeHandler) GetLike(c echo.Context) error {

	// post_idを取得
	postID := utils.PathParamToUint(c, "post_id")

	post, err := likeHandler.postController.GetPostByID(postID)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	users := post.LikeUsers

	return c.JSON(http.StatusOK, users)
}
