package wxpay

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)


// GetTradeNotification https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_7&index=3
func (this *Client) GetTradeNotification(req *http.Request) (*TradeNotification, error) {
	key, err := this.getKey()
	if err != nil {
		return nil, err
	}
	return GetTradeNotification(req, key)
}

func GetTradeNotification(req *http.Request, key string) (noti *TradeNotification, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}

	var data, _ = ioutil.ReadAll(req.Body)

	if _, err := VerifyResponseData(data, key); err != nil {
		return nil, err
	}

	if err = xml.Unmarshal(data, &noti); err != nil {
		return nil, err
	}
	return noti, err
}

func GetTradeNotificationByAppid(req *http.Request,GetWxPayClient func(appid, MchId string) (*Client))(wxClient *Client,noti *TradeNotification, err error) {
	if req == nil {
		return nil,nil, errors.New("request 参数不能为空")
	}

	var data, _ = ioutil.ReadAll(req.Body)
	if err = xml.Unmarshal(data, &noti); err != nil {
		return nil,nil, err
	}

	wxClient = GetWxPayClient(noti.AppId, noti.MCHId)
	if wxClient == nil {
		return nil, nil, errors.New("此appId 没有配置:"+noti.AppId)
	}
	key, err := wxClient.getKey()
	if err != nil {
		return nil, nil, err
	}

	if _, err := VerifyResponseData(data, key); err != nil {
		return nil,nil, err
	}
	return wxClient,noti,nil
}


func (this *Client) AckNotification(w http.ResponseWriter) {
	AckNotification(w)
}

func AckNotification(w http.ResponseWriter) {
	var v = url.Values{}
	v.Set("return_code", "SUCCESS")
	v.Set("return_msg", "OK")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(URLValueToXML(v)))
}
