package mini

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/wechat-sdk"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

var (
	ctx    = context.Background()
	wxsdk  *SDK
	err    error
	Appid  = ""
	Secret = ""
)

func TestMain(m *testing.M) {
	// NewSDK 初始化微信小程序 SDK
	//	appid：小程序 appid
	//	secret：小程序 appSecret
	wxsdk, err = NewSDK(Appid, Secret)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 打开Debug开关，输出日志
	wxsdk.DebugSwitch = wechat.DebugOff

	os.Exit(m.Run())
}

func TestGetAccessToken(t *testing.T) {
	/*
		AccessToken:50_pqQ93MxTCupSi9Ih2uYVPY_nYrLGfeDlh1yVtjek265u4KnlQkbES7WaYT0f5jTwLwxgdkrGbHssWKd83HiNEuIwixC5bt8OM1bAbVsTioHQI4ldO3JnBCz2LFd0BAZIbQToQgY2u9KmpcpmSIVgABABRE
		ExpiresIn:7200
		Errcode:0
		Errmsg:
	*/
	at := wxsdk.GetAccessToken()
	xlog.Debugf("at:%+v", at)
}

func TestCode2Session(t *testing.T) {
	session, err := wxsdk.Code2Session(ctx, "wxCode")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("at:%+v", session)
}
