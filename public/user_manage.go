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
