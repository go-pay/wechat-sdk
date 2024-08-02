package public

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/bm"
	"github.com/go-pay/wechat-sdk"
	"github.com/go-pay/xlog"
)

var (
	ctx       = context.Background()
	publicSDK *SDK
	err       error
	// 测试时，将自己的Appid和Secret填入，此appid和secret为测试号
	Appid  = "wxcfad67697020fc14"
	Secret = "c104683b3067ceac97b680aa5bf62b69"
)

func TestMain(m *testing.M) {
	xlog.SetLevel(xlog.DebugLevel)
	// 初始化微信公众号 SDK
	//	Appid：Appid
	//	Secret：appSecret
	//	autoManageToken：是否自动获取并自动维护刷新 AccessToken
	publicSDK, err = New(Appid, Secret, true)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 打开Debug开关，输出日志
	publicSDK.DebugSwitch = wechat.DebugOn

	// 若 autoManageToken 为 false，需要手动设置 Token
	// publicSDK.SetPublicAccessToken("access_token")

	// 首次获取AccessToken请通过此方法获取，之后请通过下面的回调方法获取
	at := publicSDK.GetPublicAccessToken()
	xlog.Infof("at: %s", at)

	// 每次刷新 accessToken 后，此方法回调返回 accessToken 和 有效时间（秒）
	publicSDK.SetPublicAccessTokenCallback(func(appid, accessToken string, expireIn int, err error) {
		if err != nil {
			xlog.Errorf("refresh access token error(%+v)", err)
			return
		}
		xlog.Infof("appid:%s , accessToken: %s", appid, accessToken)
		xlog.Infof("expireIn: %d", expireIn)
	})
	os.Exit(m.Run())
}

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

	rsp, err := publicSDK.QRCodeCreate(ctx, body)
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
	rsp2, err := publicSDK.QRCodeCreate(ctx, body)
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

	rsp, err := publicSDK.ShortKeyGen(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp)
}

func TestShortKeyFetch(t *testing.T) {
	rsp, err := publicSDK.ShortKeyFetch(ctx, "PwOQoY7mfqpXyFn")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp)
}

func TestUserTagCreate(t *testing.T) {
	rsp, err := publicSDK.UserTagCreate(ctx, "test_tag4")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Infof("rsp:%+v", rsp)
	xlog.Infof("rsp.Tag:%+v", rsp.Tag)
}

func TestUserTagList(t *testing.T) {
	rsp, err := publicSDK.UserTagList(ctx)
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
	err = publicSDK.UserTagUpdate(ctx, 100, "test_tag_update")
	if err != nil {
		xlog.Error(err)
		return
	}
}
