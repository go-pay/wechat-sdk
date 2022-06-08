package wechat

import (
	"testing"

	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

func TestUserInfo(t *testing.T) {
	rsp, err := openSDK.UserInfo(ctx, "openid", "zh_CN")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp)
}
