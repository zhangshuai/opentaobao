package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryWorkerModifyAPIRequest 送货入户并安装修改师傅信息 API请求
// alibaba.ascp.industry.worker.modify
//
// 送货入户并安装修改师傅信息
type AlibabaAscpIndustryWorkerModifyAPIRequest struct {
	model.Params
	// 请求对象
	_omsServiceModifyWorkerParameter *OmsServiceModifyWorkerParameter
}

// NewAlibabaAscpIndustryWorkerModifyRequest 初始化AlibabaAscpIndustryWorkerModifyAPIRequest对象
func NewAlibabaAscpIndustryWorkerModifyRequest() *AlibabaAscpIndustryWorkerModifyAPIRequest {
	return &AlibabaAscpIndustryWorkerModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryWorkerModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.worker.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryWorkerModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpIndustryWorkerModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOmsServiceModifyWorkerParameter is OmsServiceModifyWorkerParameter Setter
// 请求对象
func (r *AlibabaAscpIndustryWorkerModifyAPIRequest) SetOmsServiceModifyWorkerParameter(_omsServiceModifyWorkerParameter *OmsServiceModifyWorkerParameter) error {
	r._omsServiceModifyWorkerParameter = _omsServiceModifyWorkerParameter
	r.Set("oms_service_modify_worker_parameter", _omsServiceModifyWorkerParameter)
	return nil
}

// GetOmsServiceModifyWorkerParameter OmsServiceModifyWorkerParameter Getter
func (r AlibabaAscpIndustryWorkerModifyAPIRequest) GetOmsServiceModifyWorkerParameter() *OmsServiceModifyWorkerParameter {
	return r._omsServiceModifyWorkerParameter
}
