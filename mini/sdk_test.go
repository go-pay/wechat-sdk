package mini

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/wechat-sdk"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

var (
	ctx   = context.Background()
	wxsdk *SDK
	err   error
	// 测试时，将自己的Appid和Secret填入
	Appid  = ""
	Secret = ""
)

func TestMain(m *testing.M) {
	// NewSDK 初始化微信小程序 SDK
	//	appid：小程序 appid
	//	secret：小程序 appSecret
	//	accessToken：微信小程序AccessToken，若此参数为空，则自动获取并自动维护刷新
	wxsdk, err = NewSDK(Appid, Secret)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 可替换host节点
	// wxsdk.SetHost(wechat.HostSH)

	// New完SDK，首次获取AccessToken请通过此方法获取，之后请通过下面的回调方法获取
	at := wxsdk.GetAccessToken()
	xlog.Infof("at: %s", at)

	// 每次刷新 AccessToken 后，此方法回调返回 AccessToken 和 有效时间（秒）
	wxsdk.SetAccessTokenCallback(func(accessToken string, expireIn int, err error) {
		if err != nil {
			xlog.Errorf("refresh access token error(%+v)", err)
		}
		xlog.Infof("accessToken: %s", accessToken)
		xlog.Infof("expireIn: %d", expireIn)
	})

	// 打开Debug开关，输出日志
	wxsdk.DebugSwitch = wechat.DebugOff

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

func TestCode2Session(t *testing.T) {
	session, err := wxsdk.Code2Session(ctx, "wxCode")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("at:%+v", session)
}

func TestVerifyDecryptOpenData(t *testing.T) {
	rwData := `{"nickName":"Band","gender":1,"language":"zh_CN","city":"Guangzhou","province":"Guangdong","country":"CN","avatarUrl":"http://wx.qlogo.cn/mmopen/vi_32/1vZvI39NWFQ9XM4LtQpFrQJ1xlgZxx3w7bQxKARol6503Iuswjjn6nIGBiaycAjAtpujxyzYsrztuuICqIM5ibXQ/0"}`
	sign := "75e81ceda165f4ffa64f4068af58c64b8f54b88c"
	sessionKey := "HyVFkGl5F5OQWJZZaNzBBg=="
	ok := wxsdk.VerifyDecryptOpenData(rwData, sign, sessionKey)
	xlog.Debugf("verify result: %v", ok)
}

func TestDecryptOpenData(t *testing.T) {
	data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
	iv := "Cds8j3VYoGvnTp1BrjXdJg=="
	session := "lyY4HPQbaOYzZdG+JcYK9w=="

	//微信小程序 手机号
	phone := new(UserPhone)

	err = wxsdk.DecryptOpenData(data, iv, session, phone)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("PhoneNumber:", phone.PhoneNumber)
	xlog.Debug("PurePhoneNumber:", phone.PurePhoneNumber)
	xlog.Debug("CountryCode:", phone.CountryCode)
	xlog.Debug("Watermark:", phone.Watermark)

	sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
	encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
	iv2 := "r7BXXKkLb8qrSNn05n0qiA=="

	// 微信小程序 用户信息
	userInfo := new(UserInfo)

	err = wxsdk.DecryptOpenData(encryptedData, iv2, sessionKey, userInfo)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("NickName:", userInfo.NickName)
	xlog.Debug("AvatarUrl:", userInfo.AvatarUrl)
	xlog.Debug("Country:", userInfo.Country)
	xlog.Debug("Province:", userInfo.Province)
	xlog.Debug("City:", userInfo.City)
	xlog.Debug("Gender:", userInfo.Gender)
	xlog.Debug("OpenId:", userInfo.OpenId)
	xlog.Debug("UnionId:", userInfo.UnionId)
	xlog.Debug("Watermark:", userInfo.Watermark)
}
