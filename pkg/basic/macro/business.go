package macro

const (
	StatusAuthFailed = 1000000
)

var ErrMsg = map[int64]string{
	StatusAuthFailed: "登陆校验失败",
}

type Error struct {
	Code int64
	Msg  string
}
