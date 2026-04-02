package public

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// DraftAddDraft 新建草稿
// articles：图文消息，一个图文消息支持1到8条图文
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Add_draft.html
func (s *SDK) DraftAddDraft(c context.Context, articles []bm.BodyMap) (result *DraftAddRsp, err error) {
	path := "/cgi-bin/draft/add?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("articles", articles)

	result = &DraftAddRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// DraftGetDraft 获取草稿
// mediaId：草稿的media_id
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Get_draft.html
func (s *SDK) DraftGetDraft(c context.Context, mediaId string) (result *DraftGetRsp, err error) {
	path := "/cgi-bin/draft/get?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("media_id", mediaId)

	result = &DraftGetRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// DraftDelDraft 删除草稿
// mediaId：草稿的media_id
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Delete_draft.html
func (s *SDK) DraftDelDraft(c context.Context, mediaId string) (err error) {
	path := "/cgi-bin/draft/delete?access_token=" + s.accessToken
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

// DraftUpdateDraft 修改草稿
// mediaId：草稿的media_id
// index：要更新的文章在图文消息中的位置（多图文消息时，此字段才有意义），第一篇为0
// articles：要修改的图文消息
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Update_draft.html
func (s *SDK) DraftUpdateDraft(c context.Context, mediaId string, index int, articles bm.BodyMap) (err error) {
	path := "/cgi-bin/draft/update?access_token=" + s.accessToken
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

// FreepublishSubmit 发布接口
// mediaId：草稿的media_id
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Publish/Publish.html
func (s *SDK) FreepublishSubmit(c context.Context, mediaId string) (result *FreepublishSubmitRsp, err error) {
	path := "/cgi-bin/freepublish/submit?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("media_id", mediaId)

	result = &FreepublishSubmitRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// FreepublishGet 获取发布详情
// publishId：发布任务的id
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Publish/Get_status.html
func (s *SDK) FreepublishGet(c context.Context, publishId string) (result *FreepublishGetRsp, err error) {
	path := "/cgi-bin/freepublish/get?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("publish_id", publishId)

	result = &FreepublishGetRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// FreepublishDelete 删除发布
// articleId：成功发布时返回的article_id
// index：要删除的文章在图文消息中的位置，第一篇编号为1，该字段不填或填0会删除全部文章
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Publish/Delete_posts.html
func (s *SDK) FreepublishDelete(c context.Context, articleId string, index int) (err error) {
	path := "/cgi-bin/freepublish/delete?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("article_id", articleId)
	if index > 0 {
		body.Set("index", index)
	}

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// FreepublishGetArticle 通过article_id获取已发布文章
// articleId：要获取的草稿的article_id
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Publish/Get_article_from_id.html
func (s *SDK) FreepublishGetArticle(c context.Context, articleId string) (result *FreepublishArticleRsp, err error) {
	path := "/cgi-bin/freepublish/getarticle?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("article_id", articleId)

	result = &FreepublishArticleRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// FreepublishBatchGet 获取成功发布列表
// offset：从全部素材的该偏移位置开始返回，0表示从第一个素材
// count：返回素材的数量，取值在1到20之间
// noContent：1表示不返回content字段，0表示正常返回，默认为0
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Publish/Get_publication_records.html
func (s *SDK) FreepublishBatchGet(c context.Context, offset, count, noContent int) (result *FreepublishBatchGetRsp, err error) {
	path := "/cgi-bin/freepublish/batchget?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("offset", offset)
	body.Set("count", count)
	body.Set("no_content", noContent)

	result = &FreepublishBatchGetRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// CommentOpen 打开已群发文章评论
// msgDataId：群发返回的msg_data_id
// index：多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (s *SDK) CommentOpen(c context.Context, msgDataId int64, index int) (err error) {
	path := "/cgi-bin/comment/open?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("msg_data_id", msgDataId)
	if index >= 0 {
		body.Set("index", index)
	}

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// CommentClose 关闭已群发文章评论
// msgDataId：群发返回的msg_data_id
// index：多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (s *SDK) CommentClose(c context.Context, msgDataId int64, index int) (err error) {
	path := "/cgi-bin/comment/close?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("msg_data_id", msgDataId)
	if index >= 0 {
		body.Set("index", index)
	}

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// CommentList 查看指定文章的评论数据
// msgDataId：群发返回的msg_data_id
// index：多图文时，用来指定第几篇图文，从0开始，不带默认返回该msg_data_id的第一篇图文
// begin：起始位置
// count：获取数目（>=50会被拒绝）
// commentType：type=0 普通评论&精选评论 type=1 普通评论 type=2 精选评论
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (s *SDK) CommentList(c context.Context, msgDataId int64, index, begin, count, commentType int) (result *CommentListRsp, err error) {
	path := "/cgi-bin/comment/list?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("msg_data_id", msgDataId)
	body.Set("index", index)
	body.Set("begin", begin)
	body.Set("count", count)
	body.Set("type", commentType)

	result = &CommentListRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// CommentMarkElect 将评论标记精选
// msgDataId：群发返回的msg_data_id
// index：多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
// userCommentId：用户评论id
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (s *SDK) CommentMarkElect(c context.Context, msgDataId int64, index int, userCommentId int64) (err error) {
	path := "/cgi-bin/comment/markelect?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("msg_data_id", msgDataId)
	body.Set("index", index)
	body.Set("user_comment_id", userCommentId)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// CommentUnmarkElect 将评论取消精选
// msgDataId：群发返回的msg_data_id
// index：多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
// userCommentId：用户评论id
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (s *SDK) CommentUnmarkElect(c context.Context, msgDataId int64, index int, userCommentId int64) (err error) {
	path := "/cgi-bin/comment/unmarkelect?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("msg_data_id", msgDataId)
	body.Set("index", index)
	body.Set("user_comment_id", userCommentId)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// CommentDelete 删除评论
// msgDataId：群发返回的msg_data_id
// index：多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
// userCommentId：用户评论id
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (s *SDK) CommentDelete(c context.Context, msgDataId int64, index int, userCommentId int64) (err error) {
	path := "/cgi-bin/comment/delete?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("msg_data_id", msgDataId)
	body.Set("index", index)
	body.Set("user_comment_id", userCommentId)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// CommentReplyAdd 回复评论
// msgDataId：群发返回的msg_data_id
// index：多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
// userCommentId：用户评论id
// content：回复内容
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (s *SDK) CommentReplyAdd(c context.Context, msgDataId int64, index int, userCommentId int64, content string) (err error) {
	path := "/cgi-bin/comment/reply/add?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("msg_data_id", msgDataId)
	body.Set("index", index)
	body.Set("user_comment_id", userCommentId)
	body.Set("content", content)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}
