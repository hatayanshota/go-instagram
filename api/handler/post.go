package handler

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"instagram/api/model"
	"instagram/api/public"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//新規投稿(user_id, caption, imageでPOST)
func CreatePost(c echo.Context) error {
	public.AllowCORS(c)

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

	//io.Readerに変換
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
	if public.CaptionValidate(caption) {
		return c.String(401, "文字制限オーバー")
	}

	//画像サイズの判定
	if file.Size > 62914560 {
		return c.String(402, "画像サイズオーバー")
	}

	//画像形式の判定
	content_type, flag := public.ImageFormatValidate(format)
	if flag {
		return c.String(403, "無効な画像形式です")
	}

	id_uint := model.GetLastPostID()
	id_int := int(id_uint) + 1
	id := strconv.Itoa(id_int)

	image_url := "http://s3.ap-northeast-1.amazonaws.com/cms-intern-module/go-instagram/" + id

	model.CreatePost(user_id, image_url, caption)

	//S3に画像保存
	if err := public.UploadFile(imagefile, id, content_type); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

type PostIndex struct {
	Posts       []model.Post `json:"posts"`
	LoginUserID uint         `json:"login_user_id"`
	MaxPage     int          `json:"max_page"`
}

//投稿を10件取得
func GetPostIndex(c echo.Context) error {
	public.AllowCORS(c)

	page_num, _ := strconv.Atoi(c.QueryParam("page_num"))

	var pi PostIndex
	pi.Posts = *model.GetPost(page_num)
	login_user, is_login := model.LoginUser(c)
	if is_login {
		pi.LoginUserID = login_user.ID
	}

	post_count := model.GetPostCount()
	max_page := post_count / 10
	if post_count%10 != 0 {
		max_page += 1
	}
	pi.MaxPage = max_page

	return c.JSON(http.StatusOK, pi)
}

//指定したidの投稿を削除
func DeletePost(c echo.Context) error {
	public.AllowCORS(c)
	post_id := public.PathParamToUint(c, "post_id")

	if model.DeletePost(post_id) {
		return c.NoContent(http.StatusOK)
	} else {
		return c.NoContent(http.StatusInternalServerError)
	}
}
