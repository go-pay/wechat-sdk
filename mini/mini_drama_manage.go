package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// MediaAssetListMedia 媒资管理-获取媒资列表
// 注意：errcode = 0 为成功
// 说明：该接口用于获取已上传到平台的媒资列表。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_2-1-%E8%8E%B7%E5%8F%96%E5%AA%92%E8%B5%84%E5%88%97%E8%A1%A8
func (s *SDK) MediaAssetListMedia(c context.Context, body bm.BodyMap) (rsp *MediaAssetListMediaRsp, err error) {
	path := "/wxa/sec/vod/listmedia?access_token=" + s.accessToken
	rsp = &MediaAssetListMediaRsp{}
	if _, err = s.DoRequestPost(c, path, body, rsp); err != nil {
		return nil, err
	}
	if rsp.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", rsp.Errcode, rsp.Errmsg)
	}
	return rsp, nil
}

// MediaAssetGetMedia 媒资管理-获取媒资详细信息
// 注意：errcode = 0 为成功
// 说明：该接口用于获取已上传到平台的指定媒资信息，用于开发者后台管理使用。用于给用户客户端播放的链接应该使用getmedialink接口获取。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_2-2-%E8%8E%B7%E5%8F%96%E5%AA%92%E8%B5%84%E8%AF%A6%E7%BB%86%E4%BF%A1%E6%81%AF
func (s *SDK) MediaAssetGetMedia(c context.Context, body bm.BodyMap) (rsp *MediaAssetGetMediaRsp, err error) {
	path := "/wxa/sec/vod/getmedia?access_token=" + s.accessToken
	rsp = &MediaAssetGetMediaRsp{}
	if _, err = s.DoRequestPost(c, path, body, rsp); err != nil {
		return nil, err
	}
	if rsp.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", rsp.Errcode, rsp.Errmsg)
	}
	return rsp, nil
}

// MediaAssetGetMediaLink 媒资管理-获取媒资播放链接
// 注意：errcode = 0 为成功
// 说明：该接口用于获取视频临时播放链接，用于给用户的播放使用。只有审核通过的视频才能通过该接口获取播放链接。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_2-3-%E8%8E%B7%E5%8F%96%E5%AA%92%E8%B5%84%E6%92%AD%E6%94%BE%E9%93%BE%E6%8E%A5
func (s *SDK) MediaAssetGetMediaLink(c context.Context, body bm.BodyMap) (rsp *MediaAssetGetMediaLinkRsp, err error) {
	path := "/wxa/sec/vod/getmedialink?access_token=" + s.accessToken
	rsp = &MediaAssetGetMediaLinkRsp{}
	if _, err = s.DoRequestPost(c, path, body, rsp); err != nil {
		return nil, err
	}
	if rsp.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", rsp.Errcode, rsp.Errmsg)
	}
	return rsp, nil
}

// MediaAssetDeleteMedia 媒资管理-删除媒资
// 注意：errcode = 0 为成功
// 说明：该接口用于删除指定媒资。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_2-4-%E5%88%A0%E9%99%A4%E5%AA%92%E8%B5%84
func (s *SDK) MediaAssetDeleteMedia(c context.Context, body bm.BodyMap) (rsp *MediaAssetDeleteMediaRsp, err error) {
	path := "/wxa/sec/vod/deletemedia?access_token=" + s.accessToken
	rsp = &MediaAssetDeleteMediaRsp{}
	if _, err = s.DoRequestPost(c, path, body, rsp); err != nil {
		return nil, err
	}
	if rsp.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", rsp.Errcode, rsp.Errmsg)
	}
	return rsp, nil
}
