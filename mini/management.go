package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// GetAccountBasicInfo 获取小程序基本信息
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/basic-info/getAccountBasicInfo.html
func (s *SDK) GetAccountBasicInfo(c context.Context) (result *GetAccountBasicInfoRsp, err error) {
	path := "/cgi-bin/account/getaccountbasicinfo?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	result = &GetAccountBasicInfoRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetPage 获取已上传的代码的页面列表
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/code-management/code-management/getPage.html
func (s *SDK) GetPage(c context.Context) (result *GetPageRsp, err error) {
	path := "/wxa/get_page?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	result = &GetPageRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetCategory 获取授权小程序帐号的可选类目
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/code-management/category/getCategory.html
func (s *SDK) GetMiniCategory(c context.Context) (result *GetCategoryRsp, err error) {
	path := "/wxa/get_category?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	result = &GetCategoryRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetExtConfig 获取小程序的第三方提交代码的页面配置
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/code-management/code-management/getExtConfig.html
func (s *SDK) GetExtConfig(c context.Context) (result *GetExtConfigRsp, err error) {
	path := "/wxa/get_ext_config?access_token=" + s.accessToken

	result = &GetExtConfigRsp{}
	if _, err = s.DoRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// SetExtConfig 设置小程序的第三方提交代码的页面配置
// 注意：errcode = 0 为成功
// extConfig：第三方自定义的配置，JSON 字符串
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/code-management/code-management/setExtConfig.html
func (s *SDK) SetExtConfig(c context.Context, extConfig string) (result *SetExtConfigRsp, err error) {
	path := "/wxa/set_ext_config?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("ext_config", extConfig)

	result = &SetExtConfigRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
