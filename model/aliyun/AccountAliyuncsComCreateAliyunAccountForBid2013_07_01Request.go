package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
为bid用户创建账号 API请求
account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01

给指定的bid创建账号，同时账号打上owner bid的标记
*/
type AccountAliyuncsComCreateAliyunAccountForBid2013_07_01Request struct {
    model.Params
}

// 初始化AccountAliyuncsComCreateAliyunAccountForBid2013_07_01Request对象
func NewAccountAliyuncsComCreateAliyunAccountForBid2013_07_01Request() *AccountAliyuncsComCreateAliyunAccountForBid2013_07_01Request{
    return &AccountAliyuncsComCreateAliyunAccountForBid2013_07_01Request{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComCreateAliyunAccountForBid2013_07_01Request) GetApiMethodName() string {
    return "account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01"
}

// IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComCreateAliyunAccountForBid2013_07_01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}