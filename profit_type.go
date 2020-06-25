package wxpay

import (
	"encoding/json"
	"net/url"
)
type CommonRes struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppId      string `xml:"appid"`
	MCHId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
}
//分账
type ProfitSharingParam struct {
	TransactionID string     `json:"transaction_id,omitempty"` // 微信支付ID
	OutOrderNo    string     `json:"out_order_no,omitempty"`   // 商户订单ID
	Receivers     []Receiver `json:"receivers,omitempty"`      // 接收方列表
}

type Receiver struct {
	Type           string `json:"type,omitempty"`            // 账号类型
	Account        string `json:"account,omitempty"`         // 接收方账号
	Name           string `json:"name,omitempty"`            // 商户全称
	CustomRelation string `json:"custom_relation,omitempty"` // 子商户与接收方具体的关系，本字段最多10个字
	RelationType   string `json:"relation_type,omitempty"`   // 分账关系
	Amount         int    `json:"amount,omitempty"`          // 分账金额
	Description    string `json:"description,omitempty"`     // 分账描述
}

func (this Receiver)Params()url.Values {
	var m = make(url.Values)
	var receiver, err = json.Marshal(this)
	if err == nil {
		m.Set("receivers",string(receiver))
	}
	return m
}

func (this ProfitSharingParam) Params() url.Values {
	var m = make(url.Values)
	if len(this.TransactionID) > 0 {
		m.Set("transaction_id",this.TransactionID)
	}
	if len(this.OutOrderNo) > 0 {
		m.Set("out_order_no",this.OutOrderNo)
	}
	if len(this.Receivers) > 0 {
		var receiver, err = json.Marshal(this.Receivers)
		if err == nil {
			m.Set("receivers",string(receiver))
		}
	}
	return m
}



type ProfitSharingRsp struct {
	CommonRes
	PrepayId   string `xml:"transaction_id"`
	TradeType  string `xml:"out_order_no"`
	CodeURL    string `xml:"order_id"`
}

type AddOrDelProfitRsp struct {
	CommonRes
	Receiver   string 	`xml:"receiver"`
}

type QueryProfitRes struct {
	CommonRes
	Receiver   		string 	`xml:"receiver"`
	TransactionID   string 	`xml:"transaction_id"`
	OutOrderNo   	string 	`xml:"out_order_no"`
	OrderID   		string 	`xml:"order_id"`
	Status   		string 	`xml:"status"`
	CloseReason   	string 	`xml:"close_reason"`
	Amount     		int 	`xml:"amount"`
	Description   	string 	`xml:"description"`
}

type ProfitSharingFinishParam struct {
	TransactionID string `json:"transaction_id,omitempty"` // 微信支付ID
	OutOrderNo    string `json:"out_order_no,omitempty"`   // 商户订单ID
	Description   string `json:"description,omitempty"`    // 分账完结的原因描述
}

func (this ProfitSharingFinishParam) Params() url.Values {
	var m = make(url.Values)

	return m
}

type ProfitSharingQueryParam struct {
	TransactionID 	string  		// 微信支付ID
	OutOrderNo    	string    	// 商户订单ID
}

func (this ProfitSharingQueryParam) Params() url.Values {
	var m = make(url.Values)
	if len(this.TransactionID) > 0 {
		m.Set("transaction_id",this.TransactionID)
	}
	if len(this.OutOrderNo) > 0 {
		m.Set("out_order_no",this.TransactionID)
	}
	return m
}


type ProfitSharingReturnParam struct {
	OrderID       string `json:"order_id,omitempty"`       // 原发起分账请求时，微信返回的微信分账单号，与商户分账单号一一对应 2选1
	OutOrderNo    string `json:"out_order_no,omitempty"`   // 原发起分账请求时使用的商户后台系统的分账单号 2选1
	OutReturnNo   string `json:"out_return_no,omitempty"`  // 商户在自己后台生成的一个新的回退单号
	ReturnAccount string `json:"return_account,omitempty"` // 回退方类型是MERCHANT_ID时，填写商户ID
	ReturnAmount  int    `json:"return_amount,omitempty"`  // 回退金额
	Description   string `json:"description,omitempty"`    // 分账回退的原因描述
}

func (this ProfitSharingReturnParam) Params() url.Values {
	var m = make(url.Values)

	return m
}

type ProfitSharingReturnQueryParam struct {
	OrderID     string `json:"order_id,omitempty"`      // 原发起分账请求时，微信返回的微信分账单号，与商户分账单号一一对应 2选1
	OutOrderNo  string `json:"out_order_no,omitempty"`  // 原发起分账请求时使用的商户后台系统的分账单号 2选1
	OutReturnNo string `json:"out_return_no,omitempty"` // 商户系统内部的回退单号
}

func (this ProfitSharingReturnQueryParam) Params() url.Values {
	var m = make(url.Values)

	return m
}