package mini

import (
	"testing"

	"github.com/go-pay/bm"
	"github.com/go-pay/xlog"
)

func TestMediaAssetListMedia(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("drama_id", 20001).
		Set("limit", 20)

	rsp, err := miniSDK.MediaAssetListMedia(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("rsp:%+v", rsp)
}

func TestMediaAssetGetMedia(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("media_id", 20001)

	rsp, err := miniSDK.MediaAssetGetMedia(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("rsp:%+v", rsp)
}

func TestMediaAssetGetMediaLink(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("media_id", 28918028).
		Set("t", 1689990878)

	rsp, err := miniSDK.MediaAssetGetMediaLink(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("rsp:%+v", rsp)
}

func TestMediaAssetDeleteMedia(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("media_id", 28918028)

	rsp, err := miniSDK.MediaAssetDeleteMedia(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("rsp:%+v", rsp)
}
