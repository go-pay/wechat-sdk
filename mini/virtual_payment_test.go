package mini

import (
	"testing"

	"github.com/go-pay/bm"
	"github.com/go-pay/xlog"
)

func TestXpayQueryUserBalance(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("openid", "openid").
		Set("env", 0).
		Set("user_ip", "127.0.0.1")
	result, err := miniSDK.XpayQueryUserBalance(ctx, "signature", "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("result: %+v", result)
}

func TestXpayCurrencyPay(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("openid", "openid").
		Set("env", 0).
		Set("user_ip", "127.0.0.1").
		Set("amount", 100).
		Set("order_id", "test_order_001").
		Set("payitem", "item_info").
		Set("remark", "test")
	result, err := miniSDK.XpayCurrencyPay(ctx, "signature", "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("result: %+v", result)
}

func TestXpayQueryOrder(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("openid", "openid").
		Set("env", 0).
		Set("order_id", "test_order_001")
	result, err := miniSDK.XpayQueryOrder(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("result: %+v", result)
}

func TestXpayCancelCurrencyPay(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("openid", "openid").
		Set("env", 0).
		Set("user_ip", "127.0.0.1").
		Set("pay_order_id", "test_order_001").
		Set("order_id", "refund_order_001").
		Set("amount", 100)
	result, err := miniSDK.XpayCancelCurrencyPay(ctx, "signature", "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("result: %+v", result)
}

func TestXpayNotifyProvideGoods(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("order_id", "test_order_001").
		Set("env", 0)
	err := miniSDK.XpayNotifyProvideGoods(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("notify provide goods success")
}

func TestXpayPresentCurrency(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("openid", "openid").
		Set("env", 0).
		Set("order_id", "present_order_001").
		Set("amount", 50)
	result, err := miniSDK.XpayPresentCurrency(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("result: %+v", result)
}

func TestXpayDownloadBill(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("begin_ds", 20230801).
		Set("end_ds", 20230810)
	result, err := miniSDK.XpayDownloadBill(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("download_url: %s", result.Url)
}

func TestXpayRefundOrder(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("openid", "openid").
		Set("order_id", "test_order_001").
		Set("refund_order_id", "refund_001").
		Set("left_fee", 100).
		Set("refund_fee", 50).
		Set("biz_meta", "test_meta").
		Set("refund_reason", "1").
		Set("req_from", "1").
		Set("env", 0)
	result, err := miniSDK.XpayRefundOrder(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("result: %+v", result)
}

func TestXpayCreateWithdrawOrder(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("withdraw_no", "withdraw_001").
		Set("withdraw_amount", "0.01").
		Set("env", 0)
	result, err := miniSDK.XpayCreateWithdrawOrder(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("result: %+v", result)
}

func TestXpayQueryWithdrawOrder(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("withdraw_no", "withdraw_001").
		Set("env", 0)
	result, err := miniSDK.XpayQueryWithdrawOrder(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("result: %+v", result)
}

func TestXpayStartUploadGoods(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("upload_item", []map[string]any{
		{
			"id":       "item_001",
			"name":     "测试道具",
			"price":    100,
			"remark":   "test",
			"item_url": "https://example.com/item.png",
		},
	}).Set("env", 0)
	err := miniSDK.XpayStartUploadGoods(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("start upload goods success")
}

func TestXpayQueryUploadGoods(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("env", 0)
	result, err := miniSDK.XpayQueryUploadGoods(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("status: %d, items: %+v", result.Status, result.UploadItem)
}

func TestXpayStartPublishGoods(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("publish_item", []map[string]any{
		{"id": "item_001"},
	}).Set("env", 0)
	err := miniSDK.XpayStartPublishGoods(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("start publish goods success")
}

func TestXpayQueryPublishGoods(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("env", 0)
	result, err := miniSDK.XpayQueryPublishGoods(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("status: %d, items: %+v", result.Status, result.PublishItem)
}

func TestXpayQueryBizBalance(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("env", 0)
	result, err := miniSDK.XpayQueryBizBalance(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("balance: %+v", result.BalanceAvailable)
}

func TestXpayQueryTransferAccount(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("env", 0)
	result, err := miniSDK.XpayQueryTransferAccount(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("acct_list: %+v", result.AcctList)
}

func TestXpayQueryAdverFunds(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("page", 1).
		Set("page_size", 10).
		Set("env", 0)
	result, err := miniSDK.XpayQueryAdverFunds(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("total_page: %d, list: %+v", result.TotalPage, result.AdverFundsList)
}

func TestXpayCreateFundsBill(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("transfer_amount", 100).
		Set("transfer_account_uid", 123456).
		Set("transfer_account_name", "test_account").
		Set("transfer_account_agency_id", 789).
		Set("request_id", "req_001").
		Set("settle_begin", 1690848000).
		Set("settle_end", 1691452800).
		Set("env", 0).
		Set("authorize_advertise", 1).
		Set("fund_type", 0)
	result, err := miniSDK.XpayCreateFundsBill(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("bill_id: %s", result.BillId)
}

func TestXpayBindTransferAccount(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("transfer_account_uid", 123456).
		Set("transfer_account_org_name", "test_org").
		Set("env", 0)
	err := miniSDK.XpayBindTransferAccount(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("bind transfer account success")
}

func TestXpayQueryFundsBill(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("page", 1).
		Set("page_size", 10).
		SetBodyMap("filter", func(b bm.BodyMap) {
			b.Set("oper_time_begin", 1690848000).
				Set("oper_time_end", 1691452800)
		}).
		Set("env", 0)
	result, err := miniSDK.XpayQueryFundsBill(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("total_page: %d, list: %+v", result.TotalPage, result.BillList)
}

func TestXpayQueryRecoverBill(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("page", 1).
		Set("page_size", 10).
		SetBodyMap("filter", func(b bm.BodyMap) {
			b.Set("recover_time_begin", 1690848000).
				Set("recover_time_end", 1691452800)
		}).
		Set("env", 0)
	result, err := miniSDK.XpayQueryRecoverBill(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("total_page: %d, list: %+v", result.TotalPage, result.BillList)
}

func TestXpayQuerySubscribeContract(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("openid", "openid").
		Set("product_id", "product_001").
		Set("out_contract_code", "contract_001")
	result, err := miniSDK.XpayQuerySubscribeContract(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("authorization_state: %s", result.AuthorizationState)
}

func TestXpaySendSubscribePrePayment(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("openid", "openid").
		Set("deduct_price", 100).
		Set("product_id", "product_001").
		Set("out_contract_code", "contract_001")
	err := miniSDK.XpaySendSubscribePrePayment(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("send subscribe pre payment success")
}

func TestXpaySubmitSubscribePayOrder(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("openid", "openid").
		Set("offer_id", "offer_001").
		Set("buy_quantity", 1).
		Set("env", 0).
		Set("currency_type", "CNY").
		Set("product_id", "product_001").
		Set("deduct_price", 100).
		Set("order_id", "subscribe_order_001").
		Set("attach", "test_attach")
	err := miniSDK.XpaySubmitSubscribePayOrder(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("submit subscribe pay order success")
}

func TestXpayCancelSubscribeContract(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("openid", "openid").
		Set("termination_reason", "test reason").
		Set("product_id", "product_001").
		Set("out_contract_code", "contract_001")
	err := miniSDK.XpayCancelSubscribeContract(ctx, body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("cancel subscribe contract success")
}

func TestXpayGetComplaintList(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("env", 0).
		Set("begin_date", "2023-08-01").
		Set("end_date", "2023-08-10").
		Set("offset", 0).
		Set("limit", 10)
	result, err := miniSDK.XpayGetComplaintList(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("total: %d, complaints: %+v", result.Total, result.Complaints)
}

func TestXpayGetComplaintDetail(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("env", 0).
		Set("complaint_id", "complaint_001")
	result, err := miniSDK.XpayGetComplaintDetail(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("complaint: %+v", result.Complaint)
}

func TestXpayGetNegotiationHistory(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("env", 0).
		Set("complaint_id", "complaint_001").
		Set("offset", 0).
		Set("limit", 10)
	result, err := miniSDK.XpayGetNegotiationHistory(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("total: %d, history: %+v", result.Total, result.History)
}

func TestXpayResponseComplaint(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("env", 0).
		Set("complaint_id", "complaint_001").
		Set("response_content", "已处理").
		Set("response_images", []string{})
	err := miniSDK.XpayResponseComplaint(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("response complaint success")
}

func TestXpayCompleteComplaint(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("env", 0).
		Set("complaint_id", "complaint_001")
	err := miniSDK.XpayCompleteComplaint(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("complete complaint success")
}

func TestXpayUploadVpFile(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("env", 0).
		Set("img_url", "https://example.com/image.png").
		Set("file_name", "test.png")
	result, err := miniSDK.XpayUploadVpFile(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("file_id: %s", result.FileId)
}

func TestXpayGetUploadFileSign(t *testing.T) {
	body := make(bm.BodyMap)
	body.Set("env", 0).
		Set("wxpay_url", "https://api.mch.weixin.qq.com/v3/merchant-service/images/xxx").
		Set("convert_cos", true).
		Set("complaint_id", "complaint_001")
	result, err := miniSDK.XpayGetUploadFileSign(ctx, "pay_sig", body)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("sign: %s, cos_url: %s", result.Sign, result.CosUrl)
}
