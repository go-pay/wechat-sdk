package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// GetDailySummary 获取用户访问小程序数据概况
// 注意：errcode = 0 为成功
// beginDate：开始日期，为周一日期，格式为 yyyyMMdd，如：20170306
// endDate：结束日期，为周日日期，限定查询一周数据，格式为 yyyyMMdd，如：20170312
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getDailySummary.html
func (s *SDK) GetDailySummary(c context.Context, beginDate, endDate string) (result *DailySummaryRsp, err error) {
	path := "/datacube/getweanalysisappiddailysummarytrend?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate).
		Set("end_date", endDate)

	result = &DailySummaryRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetDailyVisitTrend 获取用户访问小程序数据日趋势
// 注意：errcode = 0 为成功
// beginDate：开始日期，格式为 yyyyMMdd
// endDate：结束日期，限定查询1天数据，允许设置的最大值为昨日，格式为 yyyyMMdd
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-trend/getDailyVisitTrend.html
func (s *SDK) GetDailyVisitTrend(c context.Context, beginDate, endDate string) (result *VisitTrendRsp, err error) {
	path := "/datacube/getweanalysisappiddailyvisittrend?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate).
		Set("end_date", endDate)

	result = &VisitTrendRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetWeeklyVisitTrend 获取用户访问小程序数据周趋势
// 注意：errcode = 0 为成功
// beginDate：开始日期，为周一日期，格式为 yyyyMMdd，如：20170306
// endDate：结束日期，为周日日期，格式为 yyyyMMdd，如：20170312
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-trend/getWeeklyVisitTrend.html
func (s *SDK) GetWeeklyVisitTrend(c context.Context, beginDate, endDate string) (result *VisitTrendRsp, err error) {
	path := "/datacube/getweanalysisappidweeklyvisittrend?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate).
		Set("end_date", endDate)

	result = &VisitTrendRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetMonthlyVisitTrend 获取用户访问小程序数据月趋势
// 注意：errcode = 0 为成功
// beginDate：开始日期，为自然月第一天，格式为 yyyyMMdd，如：20170201
// endDate：结束日期，为自然月最后一天，格式为 yyyyMMdd，如：20170228
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-trend/getMonthlyVisitTrend.html
func (s *SDK) GetMonthlyVisitTrend(c context.Context, beginDate, endDate string) (result *VisitTrendRsp, err error) {
	path := "/datacube/getweanalysisappidmonthlyvisittrend?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate).
		Set("end_date", endDate)

	result = &VisitTrendRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetDailyRetain 获取用户访问小程序日留存
// 注意：errcode = 0 为成功
// beginDate：开始日期，格式为 yyyyMMdd
// endDate：结束日期，限定查询1天数据，允许设置的最大值为昨日，格式为 yyyyMMdd
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-retain/getDailyRetain.html
func (s *SDK) GetDailyRetain(c context.Context, beginDate, endDate string) (result *RetainInfoRsp, err error) {
	path := "/datacube/getweanalysisappiddailyretaininfo?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate).
		Set("end_date", endDate)

	result = &RetainInfoRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetWeeklyRetain 获取用户访问小程序周留存
// 注意：errcode = 0 为成功
// beginDate：开始日期，为周一日期，格式为 yyyyMMdd，如：20170306
// endDate：结束日期，为周日日期，格式为 yyyyMMdd，如：20170312
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-retain/getWeeklyRetain.html
func (s *SDK) GetWeeklyRetain(c context.Context, beginDate, endDate string) (result *RetainInfoRsp, err error) {
	path := "/datacube/getweanalysisappidweeklyretaininfo?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate).
		Set("end_date", endDate)

	result = &RetainInfoRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetMonthlyRetain 获取用户访问小程序月留存
// 注意：errcode = 0 为成功
// beginDate：开始日期，为自然月第一天，格式为 yyyyMMdd，如：20170201
// endDate：结束日期，为自然月最后一天，格式为 yyyyMMdd，如：20170228
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-retain/getMonthlyRetain.html
func (s *SDK) GetMonthlyRetain(c context.Context, beginDate, endDate string) (result *RetainInfoRsp, err error) {
	path := "/datacube/getweanalysisappidmonthlyretaininfo?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate).
		Set("end_date", endDate)

	result = &RetainInfoRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetVisitPage 获取访问页面数据
// 注意：errcode = 0 为成功
// beginDate：开始日期，格式为 yyyyMMdd
// endDate：结束日期，限定查询1天数据，允许设置的最大值为昨日，格式为 yyyyMMdd
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getVisitPage.html
func (s *SDK) GetVisitPage(c context.Context, beginDate, endDate string) (result *VisitPageRsp, err error) {
	path := "/datacube/getweanalysisappidvisitpage?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate).
		Set("end_date", endDate)

	result = &VisitPageRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUserPortrait 获取小程序用户画像分布数据
// 注意：errcode = 0 为成功
// beginDate：开始日期，格式为 yyyyMMdd
// endDate：结束日期，开始日期与结束日期相差的天数限定为0/6/29，分别表示查询最近1/7/30天数据，允许设置的最大值为昨日，格式为 yyyyMMdd
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getUserPortrait.html
func (s *SDK) GetUserPortrait(c context.Context, beginDate, endDate string) (result *UserPortraitRsp, err error) {
	path := "/datacube/getweanalysisappiduserportrait?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate).
		Set("end_date", endDate)

	result = &UserPortraitRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetPerformanceData 获取小程序性能数据
// 注意：errcode = 0 为成功
// time：查询数据的时间，格式为 yyyyMMdd
// module：查询的模块，枚举值：["networkRam", "networkRom", "networkPackage", "networkCpu"]
// params：查询参数，JSON 字符串
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getPerformanceData.html
func (s *SDK) GetPerformanceData(c context.Context, time, module string, params bm.BodyMap) (result *PerformanceDataRsp, err error) {
	path := "/wxa/business/performance/boot?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("time", time).
		Set("module", module)
	if params != nil {
		body.Set("params", params)
	}

	result = &PerformanceDataRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetVisitDistribution 获取用户小程序访问分布数据
// 注意：errcode = 0 为成功
// beginDate：开始日期，格式为 yyyyMMdd
// endDate：结束日期，限定查询1天数据，允许设置的最大值为昨日，格式为 yyyyMMdd
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getVisitDistribution.html
func (s *SDK) GetVisitDistribution(c context.Context, beginDate, endDate string) (result *VisitDistributionRsp, err error) {
	path := "/datacube/getweanalysisappidvisitdistribution?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate).
		Set("end_date", endDate)

	result = &VisitDistributionRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
