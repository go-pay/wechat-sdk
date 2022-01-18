package wechat

import (
	"testing"

	"github.com/go-pay/wechat-sdk/pkg/bm"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

func TestQRCodeCreate(t *testing.T) {
	body := make(bm.BodyMap)
	// 临时二维码
	body.Set("expire_seconds", 604800).
		Set("action_name", "QR_SCENE").
		SetBodyMap("action_info", func(b bm.BodyMap) {
			b.SetBodyMap("scene", func(b bm.BodyMap) {
				b.Set("scene_id", 123)
			})
		})

	rsp, err := openSDK.QRCodeCreate(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp)

	body.Reset()

	// 永久二维码
	body.Set("action_name", "QR_LIMIT_SCENE").
		SetBodyMap("action_info", func(b bm.BodyMap) {
			b.SetBodyMap("scene", func(b bm.BodyMap) {
				b.Set("scene_id", 456)
			})
		})
	rsp2, err := openSDK.QRCodeCreate(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp2)
}

func TestShortKeyGen(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("long_data", "loooooong data").
		Set("expire_seconds", 86400)

	rsp, err := openSDK.ShortKeyGen(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp)
}

func TestShortKeyFetch(t *testing.T) {
	rsp, err := openSDK.ShortKeyFetch(ctx, "PwOQoY7mfqpXyFn")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp)
}

func TestUserTagCreate(t *testing.T) {
	rsp, err := openSDK.UserTagCreate(ctx, "test_tag4")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp)
	xlog.Infof("rsp.Tag:%+v", rsp.Tag)
}

func TestUserTagList(t *testing.T) {
	rsp, err := openSDK.UserTagList(ctx)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp)
	for _, v := range rsp.Tags {
		xlog.Infof("rsp.Tag:%+v", v)
	}
}

func TestUserTagUpdate(t *testing.T) {
	rsp, err := openSDK.UserTagUpdate(ctx, 100, "test_tag_update")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp)
}
