package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// SearchSubmitPages 提交小程序页面url及参数信息
// 注意：errcode = 0 为成功
// pages：页面信息列表，BodyMap 数组格式，每个元素包含 path 和 query
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/search-management/site-search/submitPages.html
func (s *SDK) SearchSubmitPages(c context.Context, pages []bm.BodyMap) (result *SearchSubmitPagesRsp, err error) {
	path := "/wxa/search/wxaapi_submitpages?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("pages", pages)

	result = &SearchSubmitPagesRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// SearchDeletePage 删除已提交的小程序页面
// 注意：errcode = 0 为成功
// pagePath：页面路径
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/search-management/site-search/deletePages.html
func (s *SDK) SearchDeletePage(c context.Context, pagePath string) (result *SearchDeletePageRsp, err error) {
	path := "/wxa/search/wxaapi_deletepage?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("path", pagePath)

	result = &SearchDeletePageRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
