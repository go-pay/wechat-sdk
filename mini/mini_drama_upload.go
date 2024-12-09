package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// MediaAssetSingleFileUpload 媒资上传-单个文件上传
// 注意：errcode = 0 为成功
// 说明：上传媒体（和封面）文件，上传小文件（小于10MB）时使用。上传大文件请使用分片上传接口。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-1-%E5%8D%95%E4%B8%AA%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0
func (s *SDK) MediaAssetSingleFileUpload(c context.Context, body bm.BodyMap) (rsp *MediaAssetSingleFileUploadRsp, err error) {
	path := "/wxa/sec/vod/singlefileupload?access_token=" + s.accessToken
	rsp = &MediaAssetSingleFileUploadRsp{}
	if _, err = s.DoRequestPostFile(c, path, body, rsp); err != nil {
		return nil, err
	}
	if rsp.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", rsp.Errcode, rsp.Errmsg)
	}
	return rsp, nil
}

// MediaAssetPullUpload 媒资上传-拉取上传
// 注意：errcode = 0 为成功
// 说明：该接口用于将一个网络上的视频拉取上传到平台。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-2-%E6%8B%89%E5%8F%96%E4%B8%8A%E4%BC%A0
func (s *SDK) MediaAssetPullUpload(c context.Context, body bm.BodyMap) (rsp *MediaAssetPullUploadRsp, err error) {
	path := "/wxa/sec/vod/pullupload?access_token=" + s.accessToken
	rsp = &MediaAssetPullUploadRsp{}
	if _, err = s.DoRequestPost(c, path, body, rsp); err != nil {
		return nil, err
	}
	if rsp.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", rsp.Errcode, rsp.Errmsg)
	}
	return rsp, nil
}

// MediaAssetGetTask 媒资上传-查询任务
// 注意：errcode = 0 为成功
// 说明：该接口用于查询拉取上传的任务状态。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-3-%E6%9F%A5%E8%AF%A2%E4%BB%BB%E5%8A%A1
func (s *SDK) MediaAssetGetTask(c context.Context, body bm.BodyMap) (rsp *MediaAssetGetTaskRsp, err error) {
	path := "/wxa/sec/vod/gettask?access_token=" + s.accessToken
	rsp = &MediaAssetGetTaskRsp{}
	if _, err = s.DoRequestPost(c, path, body, rsp); err != nil {
		return nil, err
	}
	if rsp.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", rsp.Errcode, rsp.Errmsg)
	}
	return rsp, nil
}

// MediaAssetApplyUpload 媒资上传-申请分片上传
// 注意：errcode = 0 为成功
// 说明：上传大文件时需使用分片上传方式，分为 3 个步骤：
// 1、申请分片上传，确定文件名、格式类型，返回 upload_id，唯一标识本次分片上传。
// 2、上传分片，多次调用上传文件分片，需要携带 part_number 和 upload_id，其中 part_number 为分片的编号，支持乱序上传。当传入 part_number 和 upload_id 都相同的时候，后发起上传请求的分片将覆盖之前的分片。
// 3、确认分片上传，当上传完所有分片后，需要完成整个文件的合并。请求体中需要给出每一个分片的 part_number 和 etag，用来校验分片的准确性，最后返回文件的 media_id。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-4-%E7%94%B3%E8%AF%B7%E5%88%86%E7%89%87%E4%B8%8A%E4%BC%A0
func (s *SDK) MediaAssetApplyUpload(c context.Context, body bm.BodyMap) (rsp *MediaAssetApplyUploadRsp, err error) {
	path := "/wxa/sec/vod/applyupload?access_token=" + s.accessToken
	rsp = &MediaAssetApplyUploadRsp{}
	if _, err = s.DoRequestPost(c, path, body, rsp); err != nil {
		return nil, err
	}
	if rsp.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", rsp.Errcode, rsp.Errmsg)
	}
	return rsp, nil
}

// MediaAssetUploadPart 媒资上传-上传分片
// 注意：errcode = 0 为成功
// 说明：将文件的其中一个分片上传到平台，最多支持100个分片，每个分片大小为5MB，最后一个分片可以小于5MB。该接口适用于视频和封面图片。视频最大支持500MB，封面图片最大支持10MB。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-5-%E4%B8%8A%E4%BC%A0%E5%88%86%E7%89%87
func (s *SDK) MediaAssetUploadPart(c context.Context, body bm.BodyMap) (rsp *MediaAssetUploadPartRsp, err error) {
	path := "/wxa/sec/vod/uploadpart?access_token=" + s.accessToken
	rsp = &MediaAssetUploadPartRsp{}
	if _, err = s.DoRequestPostFile(c, path, body, rsp); err != nil {
		return nil, err
	}
	if rsp.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", rsp.Errcode, rsp.Errmsg)
	}
	return rsp, nil
}

// MediaAssetCommitUpload 媒资上传-确认上传
// 注意：errcode = 0 为成功
// 说明：该接口用于完成整个分片上传流程，合并所有文件分片，确认媒体文件（和封面图片文件）上传到平台的结果，返回文件的 ID。请求中需要给出每一个分片的 part_number 和 etag，用来校验分片的准确性。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-6-%E7%A1%AE%E8%AE%A4%E4%B8%8A%E4%BC%A0
func (s *SDK) MediaAssetCommitUpload(c context.Context, body bm.BodyMap) (rsp *MediaAssetCommitUploadRsp, err error) {
	path := "/wxa/sec/vod/commitupload?access_token=" + s.accessToken
	rsp = &MediaAssetCommitUploadRsp{}
	if _, err = s.DoRequestPost(c, path, body, rsp); err != nil {
		return nil, err
	}
	if rsp.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", rsp.Errcode, rsp.Errmsg)
	}
	return rsp, nil
}
