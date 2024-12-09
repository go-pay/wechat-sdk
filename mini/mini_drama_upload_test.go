package mini

import (
	"os"
	"testing"

	"github.com/go-pay/bm"
	"github.com/go-pay/xlog"
)

func TestMediaAssetSingleFileUpload(t *testing.T) {
	filePath := "/Users/xxx/Downloads/test.mp4"

	file, err := os.ReadFile(filePath)
	if err != nil {
		xlog.Errorf("os.ReadFile err: %v", err)
		return
	}

	md := &bm.File{
		Name:    "test.mp4",
		Content: file,
	}

	body := make(bm.BodyMap)
	body.Set("media_name", "我的演艺 - 第1集").
		Set("media_type", "MP4").
		SetFormFile("media_data", md)

	rsp, err := miniSDK.MediaAssetSingleFileUpload(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("rsp:%+v", rsp)
}

func TestMediaAssetPullUpload(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("media_name", "我的演艺 - 第1集").
		Set("media_url", "https://developers.weixin.qq.com/test.mp4").
		Set("cover_url", "https://developers.weixin.qq.com/test.jpg")

	rsp, err := miniSDK.MediaAssetPullUpload(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("rsp:%+v", rsp)
}

func TestMediaAssetGetTask(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("task_id", 8412368)
	rsp, err := miniSDK.MediaAssetGetTask(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("rsp:%+v", rsp)
}

func TestMediaAssetApplyUpload(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("media_name", "我的演艺 - 第1集").
		Set("media_type", "MP4").
		Set("cover_type", "JPG")

	rsp, err := miniSDK.MediaAssetApplyUpload(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("rsp:%+v", rsp)
}

func TestMediaAssetUploadPart(t *testing.T) {
	filePath := "/Users/xxx/Downloads/test.mp4"

	file, err := os.ReadFile(filePath)
	if err != nil {
		xlog.Errorf("os.ReadFile err: %v", err)
		return
	}

	md := &bm.File{
		Name:    "test.mp4",
		Content: file,
	}

	body := make(bm.BodyMap)
	body.Set("upload_id", "9457878").
		Set("part_number", 1).
		Set("resource_type", 1).
		SetFormFile("data", md)

	rsp, err := miniSDK.MediaAssetUploadPart(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("rsp:%+v", rsp)
}

type PartInfo struct {
	PartNumber int    `json:"part_number"`
	Etag       string `json:"etag"`
}

func TestMediaAssetCommitUpload(t *testing.T) {
	pis := []*PartInfo{
		{
			PartNumber: 1,
			Etag:       "d899fbd1e06109ea2e4550f5751c88d6",
		},
		{
			PartNumber: 2,
			Etag:       "jfb9892jfnhda2e4550f5bvhju9392af",
		},
		{
			PartNumber: 3,
			Etag:       "bifh9u92wjefvjhytvn9u2898ef9uhea",
		},
	}

	body := make(bm.BodyMap)
	body.Set("upload_id", "9457878").
		Set("media_part_infos", pis)

	rsp, err := miniSDK.MediaAssetCommitUpload(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("rsp:%+v", rsp)
}
