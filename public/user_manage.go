package public

import (
	"context"
	"errors"

	"github.com/go-pay/wechat-sdk/pkg/bmap"
)

// UserTagCreate 用户标签创建
//	注意：errcode = 0 为成功
//	文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagCreate(c context.Context, tagName string) (ut *UserTagRsp, err error) {
	path := "/cgi-bin/tags/create?access_token=" + s.Conf.AccessToken
	body := make(bmap.BodyMap)
	body.SetBodyMap("tag", func(b bmap.BodyMap) {
		b.Set("name", tagName)
	})
	ut = &UserTagRsp{}
	if err = s.doRequestPost(c, path, body, ut); err != nil {
		return nil, err
	}
	return
}

// UserTagList 获取已创建的用户标签列表
//	注意：errcode = 0 为成功
//	文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagList(c context.Context) (utl *UserTagListRsp, err error) {
	path := "/cgi-bin/tags/get?access_token=" + s.Conf.AccessToken
	utl = &UserTagListRsp{}
	if err = s.doRequestGet(c, path, utl); err != nil {
		return nil, err
	}
	return
}

// UserTagUpdate 用户标签编辑更新
//	注意：errcode = 0 为成功
//	文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagUpdate(c context.Context, tagId int, tagName string) (ec *ErrorCode, err error) {
	path := "/cgi-bin/tags/update?access_token=" + s.Conf.AccessToken
	body := make(bmap.BodyMap)
	body.SetBodyMap("tag", func(b bmap.BodyMap) {
		b.Set("id", tagId)
		b.Set("name", tagName)
	})
	ec = &ErrorCode{}
	if err = s.doRequestPost(c, path, body, ec); err != nil {
		return nil, err
	}
	return
}

// UserTagDelete 用户标签删除
//	注意：errcode = 0 为成功
//	文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagDelete(c context.Context, tagId int) (ec *ErrorCode, err error) {
	path := "/cgi-bin/tags/delete?access_token=" + s.Conf.AccessToken
	body := make(bmap.BodyMap)
	body.SetBodyMap("tag", func(b bmap.BodyMap) {
		b.Set("id", tagId)
	})
	ec = &ErrorCode{}
	if err = s.doRequestPost(c, path, body, ec); err != nil {
		return nil, err
	}
	return
}

// UserTagFansList 获取标签下粉丝列表
//	注意：errcode = 0 为成功
//	openid：第一个拉取的 openid，不填默认从头开始拉取
//	文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagFansList(c context.Context, tagId int, openid string) (utf *UserTagFansListRsp, err error) {
	path := "/cgi-bin/user/tag/get?access_token=" + s.Conf.AccessToken
	body := make(bmap.BodyMap)
	body.Set("tagid", tagId)
	if openid != "" {
		body.Set("next_openid", openid)
	}
	utf = &UserTagFansListRsp{}
	if err = s.doRequestPost(c, path, body, utf); err != nil {
		return nil, err
	}
	return
}

// UserTagBatchTagging 批量为用户打标签
//	注意：errcode = 0 为成功
//	文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagBatchTagging(c context.Context, tagId int, openidList []string) (ec *ErrorCode, err error) {
	if len(openidList) <= 0 {
		return nil, errors.New("openid_list is empty")
	}
	path := "/cgi-bin/tags/members/batchtagging?access_token=" + s.Conf.AccessToken
	body := make(bmap.BodyMap)
	body.Set("tagid", tagId)
	body.Set("openid_list", openidList)
	ec = &ErrorCode{}
	if err = s.doRequestPost(c, path, body, ec); err != nil {
		return nil, err
	}
	return
}

// UserTagBatchUnTagging 批量为用户取消标签
//	注意：errcode = 0 为成功
//	文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagBatchUnTagging(c context.Context, tagId int, openidList []string) (ec *ErrorCode, err error) {
	if len(openidList) <= 0 {
		return nil, errors.New("openid_list is empty")
	}
	path := "/cgi-bin/tags/members/batchuntagging?access_token=" + s.Conf.AccessToken
	body := make(bmap.BodyMap)
	body.Set("tagid", tagId)
	body.Set("openid_list", openidList)
	ec = &ErrorCode{}
	if err = s.doRequestPost(c, path, body, ec); err != nil {
		return nil, err
	}
	return
}

// UserTagIdList 获取用户身上的标签列表
//	注意：errcode = 0 为成功
//	文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagIdList(c context.Context, openid string) (uti *UserTagIdListRsp, err error) {
	if openid == "" {
		return nil, errors.New("openid is empty")
	}
	path := "/cgi-bin/tags/getidlist?access_token=" + s.Conf.AccessToken
	body := make(bmap.BodyMap)
	body.Set("openid", openid)
	uti = &UserTagIdListRsp{}
	if err = s.doRequestPost(c, path, body, uti); err != nil {
		return nil, err
	}
	return
}
