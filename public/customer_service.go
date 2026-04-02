package public

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// KfAccountAdd 添加客服账号
// kfAccount：客服账号，格式为：账号前缀@公众号微信号
// nickname：客服昵称
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (s *SDK) KfAccountAdd(c context.Context, kfAccount, nickname string) (err error) {
	path := "/customservice/kfaccount/add?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("kf_account", kfAccount)
	body.Set("nickname", nickname)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// KfAccountUpdate 修改客服账号
// kfAccount：客服账号
// nickname：客服昵称
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (s *SDK) KfAccountUpdate(c context.Context, kfAccount, nickname string) (err error) {
	path := "/customservice/kfaccount/update?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("kf_account", kfAccount)
	body.Set("nickname", nickname)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// KfAccountDel 删除客服账号
// kfAccount：客服账号
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (s *SDK) KfAccountDel(c context.Context, kfAccount string) (err error) {
	path := "/customservice/kfaccount/del?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("kf_account", kfAccount)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// KfAccountUploadHeadImg 设置客服账号的头像
// kfAccount：客服账号
// media：媒体文件
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (s *SDK) KfAccountUploadHeadImg(c context.Context, kfAccount string, media *bm.File) (err error) {
	path := "/customservice/kfaccount/uploadheadimg?access_token=" + s.accessToken + "&kf_account=" + kfAccount

	result := &ErrorCode{}
	if _, err = s.doRequestUpload(c, path, "media", media, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// KfAccountGetList 获取所有客服账号
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (s *SDK) KfAccountGetList(c context.Context) (result *KfAccountListRsp, err error) {
	path := "/cgi-bin/customservice/getkflist?access_token=" + s.accessToken

	result = &KfAccountListRsp{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// KfAccountInviteWorker 邀请绑定客服账号
// kfAccount：完整客服账号，格式为：账号前缀@公众号微信号
// inviteWx：接收绑定邀请的客服微信号
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (s *SDK) KfAccountInviteWorker(c context.Context, kfAccount, inviteWx string) (err error) {
	path := "/customservice/kfaccount/inviteworker?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("kf_account", kfAccount)
	body.Set("invite_wx", inviteWx)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// KfSessionCreate 创建会话
// kfAccount：完整客服账号，格式为：账号前缀@公众号微信号
// openid：粉丝的openid
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (s *SDK) KfSessionCreate(c context.Context, kfAccount, openid string) (err error) {
	path := "/customservice/kfsession/create?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("kf_account", kfAccount)
	body.Set("openid", openid)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// KfSessionClose 关闭会话
// kfAccount：完整客服账号，格式为：账号前缀@公众号微信号
// openid：粉丝的openid
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (s *SDK) KfSessionClose(c context.Context, kfAccount, openid string) (err error) {
	path := "/customservice/kfsession/close?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("kf_account", kfAccount)
	body.Set("openid", openid)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// KfSessionGetSession 获取客户会话状态
// openid：粉丝的openid
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (s *SDK) KfSessionGetSession(c context.Context, openid string) (result *KfSessionGetRsp, err error) {
	path := "/customservice/kfsession/getsession?access_token=" + s.accessToken + "&openid=" + openid

	result = &KfSessionGetRsp{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// KfSessionGetSessionList 获取客服会话列表
// kfAccount：完整客服账号，格式为：账号前缀@公众号微信号
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (s *SDK) KfSessionGetSessionList(c context.Context, kfAccount string) (result *KfSessionListRsp, err error) {
	path := "/customservice/kfsession/getsessionlist?access_token=" + s.accessToken + "&kf_account=" + kfAccount

	result = &KfSessionListRsp{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// KfSessionGetWaitCase 获取未接入会话列表
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (s *SDK) KfSessionGetWaitCase(c context.Context) (result *KfSessionWaitCaseRsp, err error) {
	path := "/customservice/kfsession/getwaitcase?access_token=" + s.accessToken

	result = &KfSessionWaitCaseRsp{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// CustomSend 发送客服消息
// touser：普通用户openid
// msgtype：消息类型，文本为text，图片为image，语音为voice，视频为video，音乐为music，图文为news，菜单为msgmenu
// content：消息内容，BodyMap格式，根据msgtype不同而不同
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html
func (s *SDK) CustomSend(c context.Context, touser, msgtype string, content bm.BodyMap) (err error) {
	path := "/cgi-bin/message/custom/send?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("touser", touser)
	body.Set("msgtype", msgtype)
	body.Set(msgtype, content)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// CustomTyping 客服输入状态
// touser：普通用户openid
// command：Typing（对用户下发"正在输入"状态）、CancelTyping（取消对用户的"正在输入"状态）
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html
func (s *SDK) CustomTyping(c context.Context, touser, command string) (err error) {
	path := "/cgi-bin/message/custom/typing?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("touser", touser)
	body.Set("command", command)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// MsgRecordList 获取聊天记录
// starttime：起始时间，unix时间戳
// endtime：结束时间，unix时间戳，每次查询时段不能超过24小时
// msgid：消息id顺序从小到大，从1开始
// number：每次获取条数，最多10000条
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Obtain_chat_transcript.html
func (s *SDK) MsgRecordList(c context.Context, starttime, endtime int64, msgid int64, number int) (result *MsgRecordListRsp, err error) {
	path := "/customservice/msgrecord/getmsglist?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("starttime", starttime)
	body.Set("endtime", endtime)
	body.Set("msgid", msgid)
	body.Set("number", number)

	result = &MsgRecordListRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
