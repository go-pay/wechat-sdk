package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// SubscribeMessageSend 发送订阅消息
// 注意：errcode = 0 为成功
// toUser：接收者（用户）的 openid
// templateId：所需下发的订阅模板id
// page：点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
// data：模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
// miniprogramState：跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
// lang：进入小程序查看"的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/subscribe-message/sendMessage.html
func (s *SDK) SubscribeMessageSend(c context.Context, toUser, templateId, page string, data bm.BodyMap, miniprogramState, lang string) (err error) {
	path := "/cgi-bin/message/subscribe/send?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("touser", toUser).
		Set("template_id", templateId).
		Set("data", data)

	if page != "" {
		body.Set("page", page)
	}
	if miniprogramState != "" {
		body.Set("miniprogram_state", miniprogramState)
	}
	if lang != "" {
		body.Set("lang", lang)
	}

	rsp := &SubscribeMessageSendRsp{}
	if _, err = s.DoRequestPost(c, path, body, rsp); err != nil {
		return err
	}
	if rsp.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", rsp.Errcode, rsp.Errmsg)
	}
	return nil
}

// GetTemplateList 获取当前帐号下的个人模板列表
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/subscribe-message/getTemplateList.html
func (s *SDK) GetTemplateList(c context.Context) (templates *TemplateListRsp, err error) {
	path := "/wxaapi/newtmpl/gettemplate?access_token=" + s.accessToken
	templates = &TemplateListRsp{}
	if _, err = s.DoRequestGet(c, path, templates); err != nil {
		return nil, err
	}
	if templates.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", templates.Errcode, templates.Errmsg)
	}
	return templates, nil
}

// GetPubTemplateTitleList 获取帐号所属类目下的公共模板标题
// 注意：errcode = 0 为成功
// ids：类目 id，多个用逗号隔开
// start：用于分页，表示从 start 开始。从 0 开始计数。
// limit：用于分页，表示拉取 limit 条记录。最大为 30。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/subscribe-message/getPubTemplateTitleList.html
func (s *SDK) GetPubTemplateTitleList(c context.Context, ids string, start, limit int) (titles *PubTemplateTitleListRsp, err error) {
	path := "/wxaapi/newtmpl/getpubtemplatetitles?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("ids", ids).
		Set("start", start).
		Set("limit", limit)

	titles = &PubTemplateTitleListRsp{}
	if _, err = s.DoRequestPost(c, path, body, titles); err != nil {
		return nil, err
	}
	if titles.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", titles.Errcode, titles.Errmsg)
	}
	return titles, nil
}

// AddTemplate 组合模板并添加至帐号下的个人模板库
// 注意：errcode = 0 为成功
// tid：模板标题 id，可通过接口获取
// kidList：开发者自行组合好的模板关键词列表，关键词顺序可以自由搭配（例如 [3,5,4] 或 [4,5,3]），最多支持5个，最少2个关键词组合
// sceneDesc：服务场景描述，15个字以内
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/subscribe-message/addTemplate.html
func (s *SDK) AddTemplate(c context.Context, tid string, kidList []int, sceneDesc string) (result *AddTemplateRsp, err error) {
	path := "/wxaapi/newtmpl/addtemplate?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("tid", tid).
		Set("kidList", kidList).
		Set("sceneDesc", sceneDesc)

	result = &AddTemplateRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// DeleteTemplate 删除帐号下的个人模板
// 注意：errcode = 0 为成功
// priTmplId：要删除的模板id
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/subscribe-message/deleteTemplate.html
func (s *SDK) DeleteTemplate(c context.Context, priTmplId string) (err error) {
	path := "/wxaapi/newtmpl/deltemplate?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("priTmplId", priTmplId)

	rsp := &DeleteTemplateRsp{}
	if _, err = s.DoRequestPost(c, path, body, rsp); err != nil {
		return err
	}
	if rsp.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", rsp.Errcode, rsp.Errmsg)
	}
	return nil
}

// GetCategory 获取小程序账号的类目
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/subscribe-message/getCategory.html
func (s *SDK) GetCategory(c context.Context) (categories *CategoryListRsp, err error) {
	path := "/wxaapi/newtmpl/getcategory?access_token=" + s.accessToken
	categories = &CategoryListRsp{}
	if _, err = s.DoRequestGet(c, path, categories); err != nil {
		return nil, err
	}
	if categories.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", categories.Errcode, categories.Errmsg)
	}
	return categories, nil
}
