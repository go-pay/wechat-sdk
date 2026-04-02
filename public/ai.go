package public

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// VoiceTranslate 语音识别
// format：文件格式（只支持mp3，16k，单声道，最大1M），推荐格式：16k采样率16bit编码单声道，支持格式：16k/8k采样率16bit编码单声道
// voiceId：语音唯一标识
// lang：语言，zh_CN 或 en_US，默认中文
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html
func (s *SDK) VoiceTranslate(c context.Context, format, voiceId, lang string) (result *VoiceTranslateRsp, err error) {
	path := "/cgi-bin/media/voice/translatecontent?access_token=" + s.accessToken + "&format=" + format + "&voice_id=" + voiceId
	if lang != "" {
		path += "&lang=" + lang
	}

	result = &VoiceTranslateRsp{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// QrcodeImgScan 二维码/条码识别
// imgUrl：图片URL，支持http://、https://开头的URL
// img：form-data中媒体文件标识，有filename、filelength、content-type等信息
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html
func (s *SDK) QrcodeImgScan(c context.Context, imgUrl string, img *bm.File) (result *QrcodeImgScanRsp, err error) {
	path := "/cv/img/qrcode?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}
	if img != nil {
		body.Set("img", img)
	}

	result = &QrcodeImgScanRsp{}
	if img != nil {
		if _, err = s.doRequestUpload(c, path, "img", img, result); err != nil {
			return nil, err
		}
	} else {
		if _, err = s.doRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// AiCropImg 图片智能裁剪
// imgUrl：图片URL，支持http://、https://开头的URL
// img：form-data中媒体文件标识
// ratios：裁剪比例，如 "1:1" "16:9"
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html
func (s *SDK) AiCropImg(c context.Context, imgUrl string, img *bm.File, ratios []string) (result *AiCropImgRsp, err error) {
	path := "/cv/img/aicrop?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}
	if len(ratios) > 0 {
		body.Set("ratios", ratios)
	}

	result = &AiCropImgRsp{}
	if img != nil {
		if _, err = s.doRequestUploadWithForm(c, path, "img", img, "ratios", body, result); err != nil {
			return nil, err
		}
	} else {
		if _, err = s.doRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// SuperResolutionImg 图片高清化
// imgUrl：图片URL
// img：form-data中媒体文件标识
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html
func (s *SDK) SuperResolutionImg(c context.Context, imgUrl string, img *bm.File) (image []byte, err error) {
	path := "/cv/img/superresolution?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}

	if img != nil {
		// 文件上传方式返回图片二进制
		return s.doRequestGetMedia(c, path)
	}
	// URL方式也返回图片二进制
	return s.doRequestGetMedia(c, path+"&img_url="+imgUrl)
}

// OcrIdCardImg 身份证OCR识别
// imgUrl：图片URL
// img：form-data中媒体文件标识
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (s *SDK) OcrIdCardImg(c context.Context, imgUrl string, img *bm.File) (result *OcrIdCardRsp, err error) {
	path := "/cv/ocr/idcard?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}

	result = &OcrIdCardRsp{}
	if img != nil {
		if _, err = s.doRequestUpload(c, path, "img", img, result); err != nil {
			return nil, err
		}
	} else {
		if _, err = s.doRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// OcrBankCardImg 银行卡OCR识别
// imgUrl：图片URL
// img：form-data中媒体文件标识
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (s *SDK) OcrBankCardImg(c context.Context, imgUrl string, img *bm.File) (result *OcrBankCardRsp, err error) {
	path := "/cv/ocr/bankcard?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}

	result = &OcrBankCardRsp{}
	if img != nil {
		if _, err = s.doRequestUpload(c, path, "img", img, result); err != nil {
			return nil, err
		}
	} else {
		if _, err = s.doRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// OcrDrivingImg 行驶证OCR识别
// imgUrl：图片URL
// img：form-data中媒体文件标识
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (s *SDK) OcrDrivingImg(c context.Context, imgUrl string, img *bm.File) (result *OcrDrivingRsp, err error) {
	path := "/cv/ocr/driving?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}

	result = &OcrDrivingRsp{}
	if img != nil {
		if _, err = s.doRequestUpload(c, path, "img", img, result); err != nil {
			return nil, err
		}
	} else {
		if _, err = s.doRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// OcrDrivingLicenseImg 驾驶证OCR识别
// imgUrl：图片URL
// img：form-data中媒体文件标识
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (s *SDK) OcrDrivingLicenseImg(c context.Context, imgUrl string, img *bm.File) (result *OcrDrivingLicenseRsp, err error) {
	path := "/cv/ocr/drivinglicense?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}

	result = &OcrDrivingLicenseRsp{}
	if img != nil {
		if _, err = s.doRequestUpload(c, path, "img", img, result); err != nil {
			return nil, err
		}
	} else {
		if _, err = s.doRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// OcrBizLicenseImg 营业执照OCR识别
// imgUrl：图片URL
// img：form-data中媒体文件标识
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (s *SDK) OcrBizLicenseImg(c context.Context, imgUrl string, img *bm.File) (result *OcrBizLicenseRsp, err error) {
	path := "/cv/ocr/bizlicense?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}

	result = &OcrBizLicenseRsp{}
	if img != nil {
		if _, err = s.doRequestUpload(c, path, "img", img, result); err != nil {
			return nil, err
		}
	} else {
		if _, err = s.doRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// OcrCommonImg 通用印刷体OCR识别
// imgUrl：图片URL
// img：form-data中媒体文件标识
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (s *SDK) OcrCommonImg(c context.Context, imgUrl string, img *bm.File) (result *OcrCommonRsp, err error) {
	path := "/cv/ocr/comm?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}

	result = &OcrCommonRsp{}
	if img != nil {
		if _, err = s.doRequestUpload(c, path, "img", img, result); err != nil {
			return nil, err
		}
	} else {
		if _, err = s.doRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
