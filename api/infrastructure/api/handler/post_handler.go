package handler

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"instagram/api/domain/model"
	"instagram/api/infrastructure/utils"
	"instagram/api/interface/controllers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type postHandler struct {
	postController    controllers.PostController
	userController    controllers.UserController
	storageController controllers.StorageController
}

type PostHandler interface {
	CreatePost(c echo.Context) error
	GetPostIndex(c echo.Context) error
	DeletePost(c echo.Context) error
}

func NewPostHandler(pc controllers.PostController, uc controllers.UserController, sc controllers.StorageController) PostHandler {
	return &postHandler{pc, uc, sc}
}

//新規投稿(user_id, caption, imageでPOST)
func (postHandler *postHandler) CreatePost(c echo.Context) error {

	//user_idを取得
	user_id_string := c.FormValue("user_id")
	user_id_int, _ := strconv.Atoi(user_id_string)
	user_id := uint(user_id_int)

	//captionを取得
	caption := c.FormValue("caption")

	//imageを取得
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	//imageをio.Readerに変換
	imagefile, err := file.Open()
	if err != nil {
		return err
	}
	defer imagefile.Close()

	//拡張子取得
	_, format, err := image.DecodeConfig(imagefile)
	if err != nil {
		return c.String(400, "これは画像ファイルではありません")
	}

	//文字数判定
	if utils.CaptionValidate(caption) {
		return c.String(401, "文字制限オーバー")
	}

	//画像サイズの判定
	if file.Size > 62914560 {
		return c.String(402, "画像サイズオーバー")
	}

	//画像形式の判定
	content_type, flag := utils.ImageFormatValidate(format)
	if flag {
		return c.String(403, "無効な画像形式です")
	}

	// image_urlの末尾はpostテーブルに保存されるべきidとする
	id_uint, err := postHandler.postController.GetLastPostID()
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	id_int := int(id_uint) + 1
	id := strconv.Itoa(id_int)
	image_url := "http://s3.ap-northeast-1.amazonaws.com/go-instagram/" + id

	// storageに画像保存
	if err := postHandler.storageController.UploadFile(imagefile, id, content_type); err != nil {
		return err
	}

	// controllerにデータを送信
	if err := postHandler.postController.CreatePost(user_id, image_url, caption); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

type PostIndex struct {
	Posts       *[]model.Post `json:"posts"`
	LoginUserID uint          `json:"login_user_id"`
	MaxPage     int           `json:"max_page"`
}

// ページ番号を指定して投稿を取得
func (postHandler *postHandler) GetPostIndex(c echo.Context) error {

	// ページ番号を取得
	pageNum, _ := strconv.Atoi(c.QueryParam("page_num"))

	// 投稿を取得
	posts, err := postHandler.postController.GetPost(pageNum)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	githubToken := utils.ReadGithuTokenCookie(c)
	loginUser, isLogin, err := postHandler.userController.GetLoginUser(githubToken)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	var loginUserID uint
	if isLogin {
		loginUserID = loginUser.ID
	}

	postCount, err := postHandler.postController.GetPostCount()
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	maxPage := postCount / 10
	if postCount%10 != 0 {
		maxPage += 1
	}

	pi := &PostIndex{
		posts,
		loginUserID,
		maxPage,
	}

	return c.JSON(http.StatusOK, pi)

}

// 指定したidの投稿を削除
func (postHandler *postHandler) DeletePost(c echo.Context) error {

	// post_id取得
	postID := utils.PathParamToUint(c, "post_id")

	// 投稿削除要求
	if err := postHandler.postController.DeletePost(postID); err != nil {
		return c.NoContent(http.StatusOK)
	} else {
		return c.NoContent(http.StatusInternalServerError)
	}
}
