package open

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/wechat-sdk"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

var (
	ctx     = context.Background()
	openSDK *SDK
	// 测试时，将自己的Appid和Secret填入，此appid和secret为测试号
	Appid  = "wxcfad67697020fc14"
	Secret = "c104683b3067ceac97b680aa5bf62b69"
)

func TestMain(m *testing.M) {
	// New 初始化微信开放平台 SDK
	// Appid：Appid
	// Secret：appSecret
	// autoManageToken：是否自动维护刷新 AccessToken（用户量较少时推荐使用，默认10分钟轮询检测一次，发现有效期小于1.5倍轮询时间时，自动刷新）
	openSDK = New(Appid, Secret, true)

	// 打开Debug开关，输出日志
	openSDK.DebugSwitch = wechat.DebugOn

	// 可自行设置 AccessToken 刷新间隔
	//openSDK.SetAccessTokenRefreshInternal(5 * time.Minute)

	// 此方法回调返回 AccessToken
	openSDK.SetAccessTokenCallback(func(at *AT, err error) {
		if err != nil {
			xlog.Errorf("call back access token err:%+v", err)
			return
		}
		xlog.Infof("call back access token: %v", at)
	})
	os.Exit(m.Run())
}

func TestCode2AccessToken(t *testing.T) {
	at, err := openSDK.Code2AccessToken(ctx, "xxx")
	if err != nil {
		xlog.Errorf("Code2AccessToken,err:%v", err)
		return
	}
	xlog.Infof("at: %v", at)
}

func TestRefreshAccessToken(t *testing.T) {
	at, err := openSDK.RefreshAccessToken(ctx, "refreshToken")
	if err != nil {
		xlog.Errorf("RefreshAccessToken,err:%v", err)
		return
	}
	xlog.Infof("at: %v", at)
}

func TestCheckAccessToken(t *testing.T) {
	err := openSDK.CheckAccessToken(ctx, "accessToken", "openid")
	if err != nil {
		xlog.Errorf("CheckAccessToken,err:%v", err)
		return
	}
}

func TestUserInfo(t *testing.T) {
	rsp, err := openSDK.UserInfo(ctx, "access_token", "openid", "zh_CN")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp)
}
