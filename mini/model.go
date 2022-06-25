package mini

import "github.com/go-pay/wechat-sdk/pkg/xtime"

const (
	Success = 0

	MsgTypeText     = 0 // 文本消息
	MsgTypeImage    = 1 // 图片消息
	MsgTypeLink     = 2 // 图文链接
	MsgTypeMiniPage = 3 // 小程序卡片

	TypingTyping = 0 // 对用户下发"正在输入"状态，
	TypingCancel = 1 // 取消对用户的"正在输入"状态

	HostDefault = "https://api.weixin.qq.com"
)

type MsgType int8

type TypingStatus int8

//type Config struct {
//	Appid       string
//	Secret      string
//	AccessToken string
//	Host        string
//}

type ErrorCode struct {
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}

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

type CheckEncryptedResult struct {
	Vaild      bool   `json:"vaild"`             // 是否是合法的数据
	CreateTime int    `json:"create_time"`       // 加密数据生成的时间戳
	Errcode    int    `json:"errcode,omitempty"` // 错误码
	Errmsg     string `json:"errmsg,omitempty"`  // 错误信息
}

// 微信小程序解密后 用户手机号
type UserPhone struct {
	PhoneNumber     string         `json:"phoneNumber,omitempty"`
	PurePhoneNumber string         `json:"purePhoneNumber,omitempty"`
	CountryCode     string         `json:"countryCode,omitempty"`
	Watermark       *watermarkInfo `json:"watermark,omitempty"`
}

// 微信小程序解密后 用户信息
type UserInfo struct {
	OpenId    string         `json:"openId,omitempty"`
	NickName  string         `json:"nickName,omitempty"`
	Gender    int            `json:"gender,omitempty"`
	City      string         `json:"city,omitempty"`
	Province  string         `json:"province,omitempty"`
	Country   string         `json:"country,omitempty"`
	AvatarUrl string         `json:"avatarUrl,omitempty"`
	UnionId   string         `json:"unionId,omitempty"`
	Watermark *watermarkInfo `json:"watermark,omitempty"`
}

type watermarkInfo struct {
	Appid     string `json:"appid,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
}

type UploadTempMedia struct {
	MediaId   string     `json:"media_id"`          // 媒体文件上传后，获取标识，3天内有效。
	Type      string     `json:"type"`              // 文件类型
	CreatedAt xtime.Time `json:"created_at"`        // 媒体文件上传时间戳
	Errcode   int        `json:"errcode,omitempty"` // 错误码
	Errmsg    string     `json:"errmsg,omitempty"`  // 错误信息
}
