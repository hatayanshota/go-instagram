package handler

import (
	"instagram/api/model"
	"instagram/api/public"
	"net/http"

	"github.com/labstack/echo/v4"
)

// モデルとは異なるLikeを設定
type LikeData struct {
	UserID uint `json:"user_id" form:"user_id" query:"user_id"`
	PostID uint `json:"post_id" form:"post_id" query:"post_id"`
}

// いいねしたユーザーのidと投稿のidが送られることが前提
func CreateLike(c echo.Context) error {
	public.AllowCORS(c)

	like := new(LikeData)
	if err := c.Bind(like); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	model.CreateLike(like.UserID, like.PostID)
	return c.NoContent(http.StatusOK)
}

// いいね削除
func DeleteLike(c echo.Context) error {
	public.AllowCORS(c)

	like := new(LikeData)
	if err := c.Bind(like); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if model.DeleteLike(like.UserID, like.PostID) {
		return c.NoContent(http.StatusOK)
	} else {
		return c.NoContent(http.StatusInternalServerError)
	}
}

// 投稿のid(post_id)に対するいいねしたユーザーの一覧を返す
func GetLike(c echo.Context) error {
	public.AllowCORS(c)

	post_id := public.PathParamToUint(c, "post_id")

	users := model.GetPostByID(post_id).LikeUsers

	return c.JSON(http.StatusOK, users)
}
