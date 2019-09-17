package public

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Githubから得体情報を構造体で定義
type GithubUser struct {
	UserIcon string `json:"avatar_url"`
	UserName string `json:"name"`
	UserId   uint   `json:"id"`
}

// Githubからユーザ情報を取得する
func GetGithubUser(token string) (user_icon, user_name string, user_id uint) {
	res, err := http.Get("https://api.github.com/user?access_token=" + token)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var gu GithubUser

	err = json.Unmarshal(body, &gu)
	if err != nil {
		log.Fatal(err)
	}

	return gu.UserIcon, gu.UserName, gu.UserId
}
