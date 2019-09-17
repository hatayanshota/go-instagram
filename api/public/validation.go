package public

import (
	"unicode/utf8"
)

//文字数が200文字以内であることを確認
func CaptionValidate(caption string) bool {
	return utf8.RuneCountInString(caption) > 200
}

//画像形式がjpgかpngかgifであることを確認
func ImageFormatValidate(ext string) (string, bool) {
	switch ext {
	case "jpg":
		return "image/jpeg", false
	case "jpeg":
		return "image/jpeg", false
	case "png":
		return "image/png", false
	case "gif":
		return "image/gif", false
	default:
		return "", true
	}
}
