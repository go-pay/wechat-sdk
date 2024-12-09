package mini

import "github.com/go-pay/xtime"

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

type PhoneNumberRsp struct {
	PhoneInfo *UserPhone `json:"phone_info,omitempty"` // 用户手机号信息
	Errcode   int        `json:"errcode,omitempty"`    // 错误码
	Errmsg    string     `json:"errmsg,omitempty"`     // 错误信息
}

type PluginOpenPid struct {
	Openpid string `json:"openpid,omitempty"` // 插件用户的唯一标识
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}

type UserEncryptKey struct {
	Errcode     int        `json:"errcode"`
	Errmsg      string     `json:"errmsg"`
	KeyInfoList []*KeyInfo `json:"key_info_list"`
}

type KeyInfo struct {
	EncryptKey string `json:"encrypt_key"`
	Version    int    `json:"version"`
	ExpireIn   int    `json:"expire_in"`
	Iv         string `json:"iv"`
	CreateTime int    `json:"create_time"`
}

type MediaAssetSingleFileUploadRsp struct {
	MediaId int    `json:"media_id"` // 媒体文件id。
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type MediaAssetPullUploadRsp struct {
	TaskId  int    `json:"task_id"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type MediaAssetGetTaskRsp struct {
	TaskInfo *PullUploadTaskInfo `json:"task_info"`
	Errcode  int                 `json:"errcode"`
	Errmsg   string              `json:"errmsg"`
}

type PullUploadTaskInfo struct {
	Id         int    `json:"id"`
	TaskType   int    `json:"task_type"`
	Status     int    `json:"status"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	CreateTime int    `json:"create_time"`
	FinishTime int    `json:"finish_time"`
	MediaId    int    `json:"media_id"`
}

type MediaAssetApplyUploadRsp struct {
	UploadId string `json:"upload_id"`
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
}

type MediaAssetUploadPartRsp struct {
	Etag    string `json:"etag"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type MediaAssetCommitUploadRsp struct {
	MediaId int    `json:"media_id"` // 媒体文件id。
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type MediaAssetListMediaRsp struct {
	Errcode       int          `json:"errcode"`
	Errmsg        string       `json:"errmsg"`
	MediaInfoList []*MediaInfo `json:"media_info_list"`
}

type MediaInfo struct {
	MediaId     int          `json:"media_id"`
	CreateTime  int          `json:"create_time"`
	ExpireTime  int          `json:"expire_time"`
	DramaId     int          `json:"drama_id"`
	FileSize    string       `json:"file_size"`
	Duration    int          `json:"duration"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CoverUrl    string       `json:"cover_url"`
	OriginalUrl string       `json:"original_url"`
	Mp4Url      string       `json:"mp4_url"`
	HlsUrl      string       `json:"hls_url"`
	AuditDetail *AuditDetail `json:"audit_detail"`
}

type AuditDetail struct {
	Status                 int      `json:"status"`
	CreateTime             int      `json:"create_time"`
	AuditTime              int      `json:"audit_time"`
	Reason                 string   `json:"reason"`
	EvidenceMaterialIdList []string `json:"evidence_material_id_list"`
}

type MediaAssetGetMediaRsp struct {
	Errcode   int        `json:"errcode"`
	Errmsg    string     `json:"errmsg"`
	MediaInfo *MediaInfo `json:"media_info"`
}

type MediaAssetGetMediaLinkRsp struct {
	Errcode   int                `json:"errcode"`
	Errmsg    string             `json:"errmsg"`
	MediaInfo *MediaPlaybackInfo `json:"media_info"`
}

type MediaPlaybackInfo struct {
	MediaId     int    `json:"media_id"`
	Duration    int    `json:"duration"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CoverUrl    string `json:"cover_url"`
	Mp4Url      string `json:"mp4_url"`
	HlsUrl      string `json:"hls_url"`
}

type MediaAssetDeleteMediaRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
