package open

import "context"

const (
	Success = 0

	HostDefault = "https://api.weixin.qq.com"
)

type Config struct {
	Ctx         context.Context
	Appid       string
	Secret      string
	AccessToken string
	Host        string
}

type AT struct {
	AccessToken  string // 获取到的凭证
	ExpiresIn    int    // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	RefreshToken string // 用户刷新 access_token
	Openid       string // 授权用户唯一标识
	Scope        string // 用户授权的作用域，使用逗号（,）分隔
	Unionid      string // 当且仅当该移动应用已获得该用户的 userinfo 授权时，才会出现该字段
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

type ErrorCode struct {
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}

type UserInfo struct {
	Openid     string   `json:"openid,omitempty"`     // 普通用户的标识，对当前开发者帐号唯一
	Nickname   string   `json:"nickname,omitempty"`   // 普通用户昵称
	Sex        int      `json:"sex,omitempty"`        // 普通用户性别，1 为男性，2 为女性
	Province   string   `json:"province,omitempty"`   // 普通用户个人资料填写的省份
	City       string   `json:"city,omitempty"`       // 普通用户个人资料填写的城市
	Country    string   `json:"country,omitempty"`    // 国家，如中国为 CN
	Headimgurl string   `json:"headimgurl,omitempty"` // 用户头像，最后一个数值代表正方形头像大小（有 0、46、64、96、132 数值可选，0 代表 640*640 正方形头像），用户没有头像时该项为空
	Privilege  []string `json:"privilege,omitempty"`  // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	Unionid    string   `json:"unionid,omitempty"`    // 用户统一标识。针对一个微信开放平台帐号下的应用，同一用户的 unionid 是唯一的。
	Errcode    int      `json:"errcode,omitempty"`    // 错误码
	Errmsg     string   `json:"errmsg,omitempty"`     // 错误信息
}
