package handler

type appHandler struct {
	UserHandler
	PostHandler
	LikeHandler
	AuthHandler
}

// すべてのハンドラのインターフェースを持つ
type AppHandler interface {
	UserHandler
	PostHandler
	LikeHandler
	AuthHandler
}

func NewAppHandler(uh UserHandler, ph PostHandler, lh LikeHandler, ah AuthHandler) appHandler {
	return appHandler{uh, ph, lh, ah}
}
