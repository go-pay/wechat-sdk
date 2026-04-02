package public

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-pay/bm"
)

// UserTagCreate 用户标签创建
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagCreate(c context.Context, tagName string) (ut *UserTagRsp, err error) {
	path := "/cgi-bin/tags/create?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.SetBodyMap("tag", func(b bm.BodyMap) {
		b.Set("name", tagName)
	})
	ut = &UserTagRsp{}
	if _, err = s.doRequestPost(c, path, body, ut); err != nil {
		return nil, err
	}
	if ut.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", ut.Errcode, ut.Errmsg)
	}
	return ut, nil
}

// UserTagList 获取已创建的用户标签列表
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagList(c context.Context) (utl *UserTagListRsp, err error) {
	path := "/cgi-bin/tags/get?access_token=" + s.accessToken
	utl = &UserTagListRsp{}
	if _, err = s.doRequestGet(c, path, utl); err != nil {
		return nil, err
	}
	if utl.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", utl.Errcode, utl.Errmsg)
	}
	return utl, nil
}

// UserTagUpdate 用户标签编辑更新
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagUpdate(c context.Context, tagId int, tagName string) (err error) {
	path := "/cgi-bin/tags/update?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.SetBodyMap("tag", func(b bm.BodyMap) {
		b.Set("id", tagId)
		b.Set("name", tagName)
	})
	ec := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// UserTagDelete 用户标签删除
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagDelete(c context.Context, tagId int) (err error) {
	path := "/cgi-bin/tags/delete?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.SetBodyMap("tag", func(b bm.BodyMap) {
		b.Set("id", tagId)
	})
	ec := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// UserTagFansList 获取标签下粉丝列表
// 注意：errcode = 0 为成功
// openid：第一个拉取的 openid，不填默认从头开始拉取
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagFansList(c context.Context, tagId int, openid string) (utf *UserTagFansListRsp, err error) {
	path := "/cgi-bin/user/tag/get?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("tagid", tagId)
	if openid != "" {
		body.Set("next_openid", openid)
	}
	utf = &UserTagFansListRsp{}
	if _, err = s.doRequestPost(c, path, body, utf); err != nil {
		return nil, err
	}
	if utf.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", utf.Errcode, utf.Errmsg)
	}
	return utf, nil
}

// UserTagBatchTagging 批量为用户打标签
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagBatchTagging(c context.Context, tagId int, openidList []string) (err error) {
	if len(openidList) <= 0 {
		return errors.New("openid_list is empty")
	}
	path := "/cgi-bin/tags/members/batchtagging?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("tagid", tagId)
	body.Set("openid_list", openidList)
	ec := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// UserTagBatchUnTagging 批量为用户取消标签
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagBatchUnTagging(c context.Context, tagId int, openidList []string) (err error) {
	if len(openidList) <= 0 {
		return errors.New("openid_list is empty")
	}
	path := "/cgi-bin/tags/members/batchuntagging?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("tagid", tagId)
	body.Set("openid_list", openidList)
	ec := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// UserTagIdList 获取用户身上的标签列表
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagIdList(c context.Context, openid string) (uti *UserTagIdListRsp, err error) {
	if openid == "" {
		return nil, errors.New("openid is empty")
	}
	path := "/cgi-bin/tags/getidlist?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("openid", openid)
	uti = &UserTagIdListRsp{}
	if _, err = s.doRequestPost(c, path, body, uti); err != nil {
		return nil, err
	}
	if uti.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", uti.Errcode, uti.Errmsg)
	}
	return uti, nil
}

// UserInfoUpdateRemark 设置用户备注名
// openid：用户openid
// remark：新的备注名，长度必须小于30字符
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/Configuring_user_notes.html
func (s *SDK) UserInfoUpdateRemark(c context.Context, openid, remark string) (err error) {
	path := "/cgi-bin/user/info/updateremark?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("openid", openid)
	body.Set("remark", remark)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// UserInfoGet 获取用户基本信息
// openid：用户的标识，对当前公众号唯一
// lang：返回国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html
func (s *SDK) UserInfoGet(c context.Context, openid, lang string) (result *UserInfoRsp, err error) {
	path := "/cgi-bin/user/info?access_token=" + s.accessToken + "&openid=" + openid
	if lang != "" {
		path += "&lang=" + lang
	}

	result = &UserInfoRsp{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// UserInfoBatchGet 批量获取用户基本信息
// userList：用户openid列表，BodyMap数组格式：[{"openid": "otvxTs4dckWG7imySrJd6jSi0CWE", "lang": "zh_CN"}]
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html
func (s *SDK) UserInfoBatchGet(c context.Context, userList []bm.BodyMap) (result *UserInfoBatchRsp, err error) {
	path := "/cgi-bin/user/info/batchget?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("user_list", userList)

	result = &UserInfoBatchRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// UserGet 获取用户列表
// nextOpenid：第一个拉取的OPENID，不填默认从头开始拉取
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (s *SDK) UserGet(c context.Context, nextOpenid string) (result *UserListRsp, err error) {
	path := "/cgi-bin/user/get?access_token=" + s.accessToken
	if nextOpenid != "" {
		path += "&next_openid=" + nextOpenid
	}

	result = &UserListRsp{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// BlackListGetList 获取公众号的黑名单列表
// beginOpenid：当 begin_openid 为空时，默认从开头拉取
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
func (s *SDK) BlackListGetList(c context.Context, beginOpenid string) (result *BlackListRsp, err error) {
	path := "/cgi-bin/tags/members/getblacklist?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	if beginOpenid != "" {
		body.Set("begin_openid", beginOpenid)
	}

	result = &BlackListRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// BlackListBatchBlackList 拉黑用户
// openidList：需要拉入黑名单的用户的openid，一次拉黑最多允许20个
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
func (s *SDK) BlackListBatchBlackList(c context.Context, openidList []string) (err error) {
	path := "/cgi-bin/tags/members/batchblacklist?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("openid_list", openidList)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// BlackListBatchUnBlackList 取消拉黑用户
// openidList：需要取消拉黑的用户的openid，一次取消拉黑最多允许20个
// 文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
func (s *SDK) BlackListBatchUnBlackList(c context.Context, openidList []string) (err error) {
	path := "/cgi-bin/tags/members/batchunblacklist?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("openid_list", openidList)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}
