package public

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// MenuCreate 创建自定义菜单
// button：菜单数组，最多包含3个一级菜单，每个一级菜单最多包含5个二级菜单
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html
func (s *SDK) MenuCreate(c context.Context, button []bm.BodyMap) (err error) {
	path := "/cgi-bin/menu/create?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("button", button)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// MenuGet 查询自定义菜单
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Querying_Custom_Menus.html
func (s *SDK) MenuGet(c context.Context) (result *MenuGetRsp, err error) {
	path := "/cgi-bin/menu/get?access_token=" + s.accessToken

	result = &MenuGetRsp{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// MenuDelete 删除自定义菜单
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Deleting_Custom-Defined_Menu.html
func (s *SDK) MenuDelete(c context.Context) (err error) {
	path := "/cgi-bin/menu/delete?access_token=" + s.accessToken

	result := &ErrorCode{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// MenuAddConditional 创建个性化菜单
// button：菜单数组
// matchrule：菜单匹配规则
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html
func (s *SDK) MenuAddConditional(c context.Context, button []bm.BodyMap, matchrule bm.BodyMap) (result *MenuAddConditionalRsp, err error) {
	path := "/cgi-bin/menu/addconditional?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("button", button)
	body.Set("matchrule", matchrule)

	result = &MenuAddConditionalRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// MenuDelConditional 删除个性化菜单
// menuid：个性化菜单ID
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html
func (s *SDK) MenuDelConditional(c context.Context, menuid string) (err error) {
	path := "/cgi-bin/menu/delconditional?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("menuid", menuid)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// MenuTryMatch 测试个性化菜单匹配结果
// userid：用户OpenID或微信号
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html
func (s *SDK) MenuTryMatch(c context.Context, userid string) (result *MenuTryMatchRsp, err error) {
	path := "/cgi-bin/menu/trymatch?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("user_id", userid)

	result = &MenuTryMatchRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetCurrentSelfMenuInfo 获取自定义菜单配置
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Getting_Custom_Menu_Configurations.html
func (s *SDK) GetCurrentSelfMenuInfo(c context.Context) (result *CurrentSelfMenuInfoRsp, err error) {
	path := "/cgi-bin/get_current_selfmenu_info?access_token=" + s.accessToken

	result = &CurrentSelfMenuInfoRsp{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
