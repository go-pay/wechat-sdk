package mini

const (
	BaseHost   WechatHost = "https://api.weixin.qq.com"
	BaseHost2  WechatHost = "https://api2.weixin.qq.com"
	BaseHostSH WechatHost = "https://sh.api.weixin.qq.com"
	BaseHostSZ WechatHost = "https://sz.api.weixin.qq.com"
	BaseHostHK WechatHost = "https://hk.api.weixin.qq.com"
)

type WechatHost string

type Code2Session struct {
	Openid     string `json:"openid,omitempty"`      // 用户唯一标识
	SessionKey string `json:"session_key,omitempty"` // 会话密钥
	Unionid    string `json:"unionid,omitempty"`     // 用户在开放平台的唯一标识符
	Errcode    int    `json:"errcode,omitempty"`     // 错误码
	Errmsg     string `json:"errmsg,omitempty"`      // 错误信息
}

type AccessToken struct {
	AccessToken string `json:"access_token,omitempty"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in,omitempty"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	Errcode     int    `json:"errcode,omitempty"`      // 错误码
	Errmsg      string `json:"errmsg,omitempty"`       // 错误信息
}

type PaidUnionId struct {
	Unionid string `json:"unionid,omitempty"` // 用户在开放平台的唯一标识符
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}
