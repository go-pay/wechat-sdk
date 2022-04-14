package wechat

import (
	"testing"

	"github.com/go-pay/wechat-sdk/mini"
	"github.com/go-pay/wechat-sdk/pkg/bmap"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

func TestCode2Session(t *testing.T) {
	session, err := miniSDK.Code2Session(ctx, "wxCode")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("at:%+v", session)
}

func TestUniformMessageSend(t *testing.T) {
	body := make(bmap.BodyMap)
	bb := make(bmap.BodyMap)
	bb.Set("appid", "APPID").
		Set("template_id", "TEMPLATE_ID").SetBodyMap("miniprogram", func(b bmap.BodyMap) {
		b.Set("appid", "xiaochengxuappid12345").Set("pagepath", "index?foo=bar")
	})

	body.Set("touser", "Openid").
		Set("mp_template_msg", bb)

	//xlog.Debugf("%s", body.JsonBody())

	rsp, err := miniSDK.UniformMessageSend(ctx, "Openid", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("rsp:%+v", rsp)
}

func TestVerifyDecryptOpenData(t *testing.T) {
	rwData := `{"nickName":"Band","gender":1,"language":"zh_CN","city":"Guangzhou","province":"Guangdong","country":"CN","avatarUrl":"http://wx.qlogo.cn/mmopen/vi_32/1vZvI39NWFQ9XM4LtQpFrQJ1xlgZxx3w7bQxKARol6503Iuswjjn6nIGBiaycAjAtpujxyzYsrztuuICqIM5ibXQ/0"}`
	sign := "75e81ceda165f4ffa64f4068af58c64b8f54b88c"
	sessionKey := "HyVFkGl5F5OQWJZZaNzBBg=="
	ok := miniSDK.VerifyDecryptOpenData(rwData, sign, sessionKey)
	xlog.Debugf("verify result: %t", ok)
}

func TestDecryptOpenData(t *testing.T) {
	data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
	iv := "Cds8j3VYoGvnTp1BrjXdJg=="
	session := "lyY4HPQbaOYzZdG+JcYK9w=="

	// 微信小程序 手机号
	phone := new(mini.UserPhone)

	err = miniSDK.DecryptOpenData(data, iv, session, phone)
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
	userInfo := new(mini.UserInfo)

	err = miniSDK.DecryptOpenData(encryptedData, iv2, sessionKey, userInfo)
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
