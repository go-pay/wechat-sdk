package open

import (
	"context"

	"github.com/go-pay/wechat-sdk/model"
	"github.com/go-pay/wechat-sdk/pkg/bm"
)

// UserTagCreate 用户标签创建
//	注意：errcode = 0 为成功
//	文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagCreate(c context.Context, tagName string) (ut *model.UserTagRsp, err error) {
	path := "/cgi-bin/tags/create?access_token=" + s.Conf.AccessToken
	body := make(bm.BodyMap)
	body.SetBodyMap("tag", func(b bm.BodyMap) {
		b.Set("name", tagName)
	})
	ut = &model.UserTagRsp{}
	if err = s.doRequestPost(c, path, body, ut); err != nil {
		return nil, err
	}
	return
}

// UserTagList 获取已创建的用户标签
//	注意：errcode = 0 为成功
//	文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagList(c context.Context) (utl *model.UserTagListRsp, err error) {
	path := "/cgi-bin/tags/get?access_token=" + s.Conf.AccessToken
	utl = &model.UserTagListRsp{}
	if err = s.doRequestGet(c, path, utl); err != nil {
		return nil, err
	}
	return
}

// UserTagUpdate 用户标签编辑更新
//	注意：errcode = 0 为成功
//	文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (s *SDK) UserTagUpdate(c context.Context, tagId int, tagName string) (ec *model.ErrorCode, err error) {
	path := "/cgi-bin/tags/update?access_token=" + s.Conf.AccessToken
	body := make(bm.BodyMap)
	body.SetBodyMap("tag", func(b bm.BodyMap) {
		b.Set("id", tagId)
		b.Set("name", tagName)
	})
	ec = &model.ErrorCode{}
	if err = s.doRequestPost(c, path, body, ec); err != nil {
		return nil, err
	}
	return
}
