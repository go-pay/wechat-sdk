package public

import "github.com/go-pay/wechat-sdk/pkg/xtime"

const (
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

type QRCodeRsp struct {
	Errcode       int    `json:"errcode,omitempty"`
	Errmsg        string `json:"errmsg,omitempty"`
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds,omitempty"`
	Url           string `json:"url"`
}

type ShortKeyGenRsp struct {
	Errcode  int    `json:"errcode,omitempty"`
	Errmsg   string `json:"errmsg,omitempty"`
	ShortKey string `json:"short_key"`
}

type ShortKeyFetchRsp struct {
	Errcode       int        `json:"errcode,omitempty"`
	Errmsg        string     `json:"errmsg,omitempty"`
	LongData      string     `json:"long_data"`
	CreateTime    xtime.Time `json:"create_time"`
	ExpireSeconds int        `json:"expire_seconds"`
}

type UserTagRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Tag     *Tag   `json:"tag,omitempty"`
}

type Tag struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"` // 此标签下粉丝数
}

type UserTagListRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Tags    []*Tag `json:"tags,omitempty"`
}
type UserTagFansListRsp struct {
	Errcode    int       `json:"errcode,omitempty"`
	Errmsg     string    `json:"errmsg,omitempty"`
	Count      int       `json:"count"`
	NextOpenid string    `json:"next_openid,omitempty"`
	Data       *FansData `json:"data"`
}

type FansData struct {
	Openid []string `json:"openid"`
}

type UserTagIdListRsp struct {
	Errcode   int    `json:"errcode,omitempty"`
	Errmsg    string `json:"errmsg,omitempty"`
	TagidList []int  `json:"tagid_list,omitempty"`
}
