package public

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// MediaUpload 新增临时素材
// mediaType：媒体文件类型，分别有图片（image）、语音（voice）、视频（video）、缩略图（thumb）
// media：form-data中媒体文件标识，有filename、filelength、content-type等信息
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html
func (s *SDK) MediaUpload(c context.Context, mediaType string, media *bm.File) (result *MediaUploadRsp, err error) {
	path := "/cgi-bin/media/upload?access_token=" + s.accessToken + "&type=" + mediaType

	result = &MediaUploadRsp{}
	if _, err = s.doRequestUpload(c, path, "media", media, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// MediaGet 获取临时素材
// mediaId：媒体文件ID
// 返回文件的二进制内容
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_temporary_materials.html
func (s *SDK) MediaGet(c context.Context, mediaId string) (media []byte, err error) {
	path := "/cgi-bin/media/get?access_token=" + s.accessToken + "&media_id=" + mediaId

	return s.doRequestGetMedia(c, path)
}

// MaterialAddNews 新增永久图文素材
// articles：图文消息，一个图文消息支持1到8条图文
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (s *SDK) MaterialAddNews(c context.Context, articles []bm.BodyMap) (result *MaterialAddNewsRsp, err error) {
	path := "/cgi-bin/material/add_news?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("articles", articles)

	result = &MaterialAddNewsRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// MaterialUploadImg 上传图文消息内的图片
// media：form-data中媒体文件标识
// 返回图片URL，可放置在图文消息中使用
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (s *SDK) MaterialUploadImg(c context.Context, media *bm.File) (result *MaterialUploadImgRsp, err error) {
	path := "/cgi-bin/media/uploadimg?access_token=" + s.accessToken

	result = &MaterialUploadImgRsp{}
	if _, err = s.doRequestUpload(c, path, "media", media, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// MaterialAddMaterial 新增其他类型永久素材
// mediaType：媒体文件类型，分别有图片（image）、语音（voice）、视频（video）、缩略图（thumb）
// media：form-data中媒体文件标识
// description：视频素材的描述，BodyMap格式：{"title": "标题", "introduction": "描述"}，非视频素材传nil
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (s *SDK) MaterialAddMaterial(c context.Context, mediaType string, media *bm.File, description bm.BodyMap) (result *MaterialAddMaterialRsp, err error) {
	path := "/cgi-bin/material/add_material?access_token=" + s.accessToken + "&type=" + mediaType

	result = &MaterialAddMaterialRsp{}
	if description != nil {
		if _, err = s.doRequestUploadWithForm(c, path, "media", media, "description", description, result); err != nil {
			return nil, err
		}
	} else {
		if _, err = s.doRequestUpload(c, path, "media", media, result); err != nil {
			return nil, err
		}
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// MaterialGetMaterial 获取永久素材
// mediaId：媒体文件ID
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (s *SDK) MaterialGetMaterial(c context.Context, mediaId string) (result *MaterialGetMaterialRsp, err error) {
	path := "/cgi-bin/material/get_material?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("media_id", mediaId)

	result = &MaterialGetMaterialRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// MaterialDelMaterial 删除永久素材
// mediaId：媒体文件ID
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Deleting_Permanent_Assets.html
func (s *SDK) MaterialDelMaterial(c context.Context, mediaId string) (err error) {
	path := "/cgi-bin/material/del_material?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("media_id", mediaId)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// MaterialUpdateNews 修改永久图文素材
// mediaId：要修改的图文消息的id
// index：要更新的文章在图文消息中的位置（多图文消息时，此字段才有意义），第一篇为0
// articles：要修改的图文消息
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Editing_Permanent_Rich_Media_Assets.html
func (s *SDK) MaterialUpdateNews(c context.Context, mediaId string, index int, articles bm.BodyMap) (err error) {
	path := "/cgi-bin/material/update_news?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("media_id", mediaId)
	body.Set("index", index)
	body.Set("articles", articles)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// MaterialGetMaterialCount 获取素材总数
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_the_total_of_all_materials.html
func (s *SDK) MaterialGetMaterialCount(c context.Context) (result *MaterialGetMaterialCountRsp, err error) {
	path := "/cgi-bin/material/get_materialcount?access_token=" + s.accessToken

	result = &MaterialGetMaterialCountRsp{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// MaterialBatchGetMaterial 获取素材列表
// materialType：素材的类型，图片（image）、视频（video）、语音（voice）、图文（news）
// offset：从全部素材的该偏移位置开始返回，0表示从第一个素材
// count：返回素材的数量，取值在1到20之间
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html
func (s *SDK) MaterialBatchGetMaterial(c context.Context, materialType string, offset, count int) (result *MaterialBatchGetMaterialRsp, err error) {
	path := "/cgi-bin/material/batchget_material?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("type", materialType)
	body.Set("offset", offset)
	body.Set("count", count)

	result = &MaterialBatchGetMaterialRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
