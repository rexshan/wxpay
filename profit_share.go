package wxpay

const (
	AddReceiver = "/pay/profitsharingaddreceiver"
	DelReceiver = "/pay/profitsharingremovereceiver"
	FinProfit = "/secapi/pay/profitsharingfinish"
	StartMulProfit = "/secapi/pay/multiprofitsharing"
	StartProfit = "/secapi/pay/profitsharing"
	QueryProfit = "/pay/profitsharingquery"

	RefundProfit = "/secapi/pay/profitsharingreturn"
	RefundQuery = "/pay/profitsharingreturnquery"
)


// ProfitSharingAddReceiver 添加分账接受方
func (this *Client) ProfitSharingAddReceiver(params Receiver) (result *AddOrDelProfitRsp, err error) {
	if err = this.doRequest("POST", this.BuildAPI(AddReceiver), params, &result); err != nil {
		return nil, err
	}
	return result, err
}

// ProfitSharingRemoveReceiver 删除分账接收方
func (this *Client) ProfitSharingRemoveReceiver(params Receiver) (result *AddOrDelProfitRsp, err error) {
	if err = this.doRequest("POST", this.BuildAPI(DelReceiver), params, &result); err != nil {
		return nil, err
	}
	return result, err
}

// ProfitSharingFinish 完成分账
func (this *Client) ProfitSharingFinish(params ProfitSharingFinishParam) (result *ProfitSharingRsp, err error) {
	if err = this.doRequest("POST", this.BuildAPI(FinProfit), params, &result); err != nil {
		return nil, err
	}
	return result, err
}

// ProfitSharing 开始分账 默认为单次分账 option=multi为多次分账
func (this *Client) ProfitSharing(params ProfitSharingParam, option string) (result *ProfitSharingRsp,err error) {
	var apiPath = StartProfit
	if option == "multi" {
		apiPath = StartMulProfit
	}
	if err = this.doRequest("POST", this.BuildAPI(apiPath), params, &result); err != nil {
		return nil, err
	}
	return result, err
}

// ProfitSharingQuery 查询分账结果
func (this *Client) ProfitSharingQuery(params ProfitSharingQueryParam) (result *QueryProfitRes, err error) {
	if err = this.doRequest("POST", this.BuildAPI(QueryProfit), params, &result); err != nil {
		return nil, err
	}
	return result, err
}

// ProfitSharingReturn 分账回退
/*
func (this *Client) ProfitSharingReturn(params ProfitSharingReturn) (map[string]string, error) {
	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())
	m["return_account_type"] = "MERCHANT_ID "
	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/secapi/pay/profitsharingreturn", strings.NewReader(utils.MAP2XML(m)), true)
}

// ProfitSharingReturnQuery 分账回退结果查询
func (this *Client) ProfitSharingReturnQuery(params ProfitSharingReturnQuery) (map[string]string, error) {
	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())
	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/pay/profitsharingreturnquery", strings.NewReader(utils.MAP2XML(m)), false)
}

*/
