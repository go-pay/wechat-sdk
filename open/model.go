package open

import "context"

type Config struct {
	Ctx         context.Context
	Appid       string
	Secret      string
	AccessToken string
	Host        string
}

type AccessToken struct {
	AccessToken  string `json:"access_token,omitempty"`  // 获取到的凭证
	ExpiresIn    int    `json:"expires_in,omitempty"`    // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	RefreshToken string `json:"refresh_token,omitempty"` // 用户刷新 access_token
	Openid       string `json:"openid,omitempty"`        // 授权用户唯一标识
	Scope        string `json:"scope,omitempty"`         // 用户授权的作用域，使用逗号（,）分隔
	Unionid      string `json:"unionid,omitempty"`       // 当且仅当该移动应用已获得该用户的 userinfo 授权时，才会出现该字段
	Errcode      int    `json:"errcode,omitempty"`       // 错误码
	Errmsg       string `json:"errmsg,omitempty"`        // 错误信息
}
