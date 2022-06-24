package open

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

var (
	ctx     = context.Background()
	openSDK *SDK
	err     error
	// 测试时，将自己的Appid和Secret填入，此appid和secret为测试号
	Appid  = "wxcfad67697020fc14"
	Secret = "c104683b3067ceac97b680aa5bf62b69"
)

func TestMain(m *testing.M) {
	// 初始化微信开放平台 SDK
	//	Appid：Appid
	//	Secret：appSecret
	//	autoManageToken：是否自动获取并自动维护刷新 AccessToken
	openSDK = New(Appid, Secret, true)

	// 打开Debug开关，输出日志
	openSDK.DebugSwitch = DebugOn

	// 如果 自行维护 AccessToken，请需要手动设置 Token
	// openSDK.SetOpenAccessToken("access_token")

	// 注意：必须优先换取 开放平台 AccessToken，否则会导致部分接口调用失败
	at, err := openSDK.Code2AccessToken(ctx, "xxx")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("at: %s", at)

	// 每次刷新 accessToken 后，此方法回调返回 accessToken 和 有效时间（秒）
	openSDK.SetOpenAccessTokenCallback(func(at *AccessToken, err error) {
		if err != nil {
			xlog.Errorf("refresh access token error(%+v)", err)
			return
		}
		xlog.Infof("AccessToken: %+v", at)
	})
	os.Exit(m.Run())
}

func TestUserInfo(t *testing.T) {
	rsp, err := openSDK.UserInfo(ctx, "openid", "zh_CN")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp)
}
