package model

const (
	Mini = "mini"
	Open = "open"

	DebugOff = 0
	DebugOn  = 1
)

type Config struct {
	Appid       string
	Secret      string
	AccessToken string
	Host        string
}

type ErrorCode struct {
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}
