package wechat

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/wechat-sdk/mini"
	"github.com/go-pay/wechat-sdk/open"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

var (
	ctx     = context.Background()
	wxsdk   *SDK
	miniSDK *mini.SDK
	openSDK *open.SDK
	err     error
	// 测试时，将自己的Appid和Secret填入，此appid和secret为测试号
	Appid  = "wxcfad67697020fc14"
	Secret = "c104683b3067ceac97b680aa5bf62b69"
)

func TestMain(m *testing.M) {
	// NewSDK 初始化微信 SDK
	//	Appid：Appid
	//	Secret：appSecret
	//	accessToken：AccessToken，若此参数为空，则自动获取并自动维护刷新
	wxsdk, err = NewSDK(Appid, Secret)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 可替换host节点
	//wxsdk.SetHost(HostSH)
	// 打开Debug开关，输出日志
	wxsdk.DebugSwitch = DebugOff

	// New完SDK，首次获取AccessToken请通过此方法获取，之后请通过下面的回调方法获取
	at := wxsdk.GetAccessToken()
	xlog.Infof("at: %s", at)

	// 每次刷新 accessToken 后，此方法回调返回 accessToken 和 有效时间（秒）
	wxsdk.SetAccessTokenCallback(func(accessToken string, expireIn int, err error) {
		if err != nil {
			xlog.Errorf("refresh access token error(%+v)", err)
		}
		xlog.Infof("accessToken: %s", accessToken)
		xlog.Infof("expireIn: %d", expireIn)
	})

	// New 微信小程序 SDK
	miniSDK = wxsdk.NewMini()
	//miniSDK.DebugSwitch = DebugOn

	// New 微信公众号 SDK
	openSDK = wxsdk.NewOpen()
	openSDK.DebugSwitch = DebugOff

	os.Exit(m.Run())
}

func TestGetAccessToken(t *testing.T) {
	/*
		accessToken:50_pqQ93MxTCupSi9Ih2uYVPY_nYrLGfeDlh1yVtjek265u4KnlQkbES7WaYT0f5jTwLwxgdkrGbHssWKd83HiNEuIwixC5bt8OM1bAbVsTioHQI4ldO3JnBCz2LFd0BAZIbQToQgY2u9KmpcpmSIVgABABRE
		ExpiresIn:7200
		Errcode:0
		Errmsg:
	*/
	xlog.Debugf("at:%s", wxsdk.GetAccessToken())
}
