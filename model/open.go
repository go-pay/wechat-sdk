package model

import "github.com/go-pay/wechat-sdk/pkg/xtime"

type QRCodeRsp struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds,omitempty"`
	Url           string `json:"url"`
}

type ShortKeyGenRsp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	ShortKey string `json:"short_key"`
}

type ShortKeyFetchRsp struct {
	Errcode       int        `json:"errcode"`
	Errmsg        string     `json:"errmsg"`
	LongData      string     `json:"long_data"`
	CreateTime    xtime.Time `json:"create_time"`
	ExpireSeconds int        `json:"expire_seconds"`
}
