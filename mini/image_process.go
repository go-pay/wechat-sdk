package mini

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-pay/bm"
	"github.com/go-pay/util"
	"github.com/go-pay/wechat-sdk"
	"github.com/go-pay/xhttp"
)

// AiCrop 图片智能裁剪，根据图片内容返回智能剪裁区域
// 注意：errcode = 0 为成功
// imgUrl：要检测的图片 url，传这个则不用传 img 参数
// img：form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传 img_url
// ratios：裁剪比例，可选 "1:1"、"3:4"、"4:3"、"16:9"、"9:16"，不传默认为 "3:4"
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr-identify/img-proc/aiCrop.html
func (s *SDK) AiCrop(c context.Context, imgUrl string, img *bm.File, ratios []string) (result *AiCropRsp, err error) {
	path := "/cv/img/aicrop?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}
	if img != nil {
		body.SetFormFile("img", img)
	}
	if len(ratios) > 0 {
		body.Set("ratios", ratios)
	}

	result = &AiCropRsp{}
	if img != nil {
		// 使用 multipart/form-data
		if _, err = s.DoRequestPostFile(c, path, body, result); err != nil {
			return nil, err
		}
	} else {
		// 使用 JSON
		if _, err = s.DoRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}

	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ScanQRCode 本接口提供基于小程序的条码/二维码识别的API
// 注意：errcode = 0 为成功
// imgUrl：要检测的图片 url，传这个则不用传 img 参数
// img：form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传 img_url
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr-identify/img-proc/scanQRCode.html
func (s *SDK) ScanQRCode(c context.Context, imgUrl string, img *bm.File) (result *ScanQRCodeRsp, err error) {
	path := "/cv/img/qrcode?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}
	if img != nil {
		body.SetFormFile("img", img)
	}

	result = &ScanQRCodeRsp{}
	if img != nil {
		// 使用 multipart/form-data
		if _, err = s.DoRequestPostFile(c, path, body, result); err != nil {
			return nil, err
		}
	} else {
		// 使用 JSON
		if _, err = s.DoRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}

	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// SuperResolution 图片高清化，提高图片分辨率，返回处理后的图片二进制内容
// 注意：errcode = 0 为成功
// imgUrl：要检测的图片 url，传这个则不用传 img 参数
// img：form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传 img_url
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr-identify/img-proc/superResolution.html
func (s *SDK) SuperResolution(c context.Context, imgUrl string, img *bm.File) (image []byte, err error) {
	path := "/cv/img/superresolution?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}
	if img != nil {
		body.SetFormFile("img", img)
	}

	if img != nil {
		// 使用 multipart/form-data
		image, err = s.doRequestPostFileForImage(c, path, body)
	} else {
		// 使用 JSON
		image, err = s.doRequestPostForImage(c, path, body)
	}

	if err != nil {
		return nil, err
	}
	return image, nil
}

// OcrIdCard 身份证OCR识别
// 注意：errcode = 0 为成功
// imgUrl：要检测的图片 url，传这个则不用传 img 参数
// img：form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传 img_url
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr-identify/ocr/idcard.html
func (s *SDK) OcrIdCard(c context.Context, imgUrl string, img *bm.File) (result *OcrIdCardRsp, err error) {
	path := "/cv/ocr/idcard?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}
	if img != nil {
		body.SetFormFile("img", img)
	}

	result = &OcrIdCardRsp{}
	if img != nil {
		// 使用 multipart/form-data
		if _, err = s.DoRequestPostFile(c, path, body, result); err != nil {
			return nil, err
		}
	} else {
		// 使用 JSON
		if _, err = s.DoRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}

	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// OcrBankCard 银行卡OCR识别
// 注意：errcode = 0 为成功
// imgUrl：要检测的图片 url，传这个则不用传 img 参数
// img：form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传 img_url
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr-identify/ocr/bankcard.html
func (s *SDK) OcrBankCard(c context.Context, imgUrl string, img *bm.File) (result *OcrBankCardRsp, err error) {
	path := "/cv/ocr/bankcard?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}
	if img != nil {
		body.SetFormFile("img", img)
	}

	result = &OcrBankCardRsp{}
	if img != nil {
		// 使用 multipart/form-data
		if _, err = s.DoRequestPostFile(c, path, body, result); err != nil {
			return nil, err
		}
	} else {
		// 使用 JSON
		if _, err = s.DoRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}

	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// OcrDriving 驾驶证OCR识别
// 注意：errcode = 0 为成功
// imgUrl：要检测的图片 url，传这个则不用传 img 参数
// img：form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传 img_url
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr-identify/ocr/driving.html
func (s *SDK) OcrDriving(c context.Context, imgUrl string, img *bm.File) (result *OcrDrivingRsp, err error) {
	path := "/cv/ocr/driving?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}
	if img != nil {
		body.SetFormFile("img", img)
	}

	result = &OcrDrivingRsp{}
	if img != nil {
		// 使用 multipart/form-data
		if _, err = s.DoRequestPostFile(c, path, body, result); err != nil {
			return nil, err
		}
	} else {
		// 使用 JSON
		if _, err = s.DoRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}

	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// OcrVehicleLicense 行驶证OCR识别
// 注意：errcode = 0 为成功
// imgUrl：要检测的图片 url，传这个则不用传 img 参数
// img：form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传 img_url
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr-identify/ocr/vehicleLicense.html
func (s *SDK) OcrVehicleLicense(c context.Context, imgUrl string, img *bm.File) (result *OcrVehicleLicenseRsp, err error) {
	path := "/cv/ocr/drivinglicense?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}
	if img != nil {
		body.SetFormFile("img", img)
	}

	result = &OcrVehicleLicenseRsp{}
	if img != nil {
		// 使用 multipart/form-data
		if _, err = s.DoRequestPostFile(c, path, body, result); err != nil {
			return nil, err
		}
	} else {
		// 使用 JSON
		if _, err = s.DoRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}

	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// OcrBusinessLicense 营业执照OCR识别
// 注意：errcode = 0 为成功
// imgUrl：要检测的图片 url，传这个则不用传 img 参数
// img：form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传 img_url
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr-identify/ocr/businessLicense.html
func (s *SDK) OcrBusinessLicense(c context.Context, imgUrl string, img *bm.File) (result *OcrBusinessLicenseRsp, err error) {
	path := "/cv/ocr/bizlicense?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}
	if img != nil {
		body.SetFormFile("img", img)
	}

	result = &OcrBusinessLicenseRsp{}
	if img != nil {
		// 使用 multipart/form-data
		if _, err = s.DoRequestPostFile(c, path, body, result); err != nil {
			return nil, err
		}
	} else {
		// 使用 JSON
		if _, err = s.DoRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}

	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// OcrCommon 通用印刷体OCR识别
// 注意：errcode = 0 为成功
// imgUrl：要检测的图片 url，传这个则不用传 img 参数
// img：form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传 img_url
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr-identify/ocr/common.html
func (s *SDK) OcrCommon(c context.Context, imgUrl string, img *bm.File) (result *OcrCommonRsp, err error) {
	path := "/cv/ocr/comm?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}
	if img != nil {
		body.SetFormFile("img", img)
	}

	result = &OcrCommonRsp{}
	if img != nil {
		// 使用 multipart/form-data
		if _, err = s.DoRequestPostFile(c, path, body, result); err != nil {
			return nil, err
		}
	} else {
		// 使用 JSON
		if _, err = s.DoRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}

	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// OcrPlateNumber 车牌OCR识别
// 注意：errcode = 0 为成功
// imgUrl：要检测的图片 url，传这个则不用传 img 参数
// img：form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传 img_url
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr-identify/ocr/plateNumber.html
func (s *SDK) OcrPlateNumber(c context.Context, imgUrl string, img *bm.File) (result *OcrPlateNumberRsp, err error) {
	path := "/cv/ocr/platenum?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	if imgUrl != "" {
		body.Set("img_url", imgUrl)
	}
	if img != nil {
		body.SetFormFile("img", img)
	}

	result = &OcrPlateNumberRsp{}
	if img != nil {
		// 使用 multipart/form-data
		if _, err = s.DoRequestPostFile(c, path, body, result); err != nil {
			return nil, err
		}
	} else {
		// 使用 JSON
		if _, err = s.DoRequestPost(c, path, body, result); err != nil {
			return nil, err
		}
	}

	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// doRequestPostFileForImage 处理返回图片的 POST 文件上传请求
func (s *SDK) doRequestPostFileForImage(c context.Context, path string, body bm.BodyMap) (image []byte, err error) {
	uri := s.Host + path
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Mini_SDK_URI: %s", uri)
	}

	req := s.hc.Req(xhttp.TypeMultipartFormData)
	req.Header.Add(wechat.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := req.Post(uri).SendMultipartBodyMap(body).EndBytes(c)
	if err != nil {
		return nil, fmt.Errorf("http.request(POST, %s), err:%w", uri, err)
	}

	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Mini_SDK_Response: [%d] -> %d bytes", res.StatusCode, len(bs))
	}

	// 判断返回的是图片还是错误信息
	contentType := res.Header.Get("Content-Type")
	if contentType == "image/jpeg" || contentType == "image/png" {
		return bs, nil
	}

	// 如果不是图片，尝试解析错误信息
	ec := &ErrorCode{}
	if err = json.Unmarshal(bs, ec); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if ec.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}

	return bs, nil
}
