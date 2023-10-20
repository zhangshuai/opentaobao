package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAccountDaycostGetAPIRequest 查询今日消耗 API请求
// alibaba.scbp.account.daycost.get
//
// 查询今日消耗
type AlibabaScbpAccountDaycostGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpAccountDaycostGetRequest 初始化AlibabaScbpAccountDaycostGetAPIRequest对象
func NewAlibabaScbpAccountDaycostGetRequest() *AlibabaScbpAccountDaycostGetAPIRequest {
	return &AlibabaScbpAccountDaycostGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountDaycostGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.daycost.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountDaycostGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAccountDaycostGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
