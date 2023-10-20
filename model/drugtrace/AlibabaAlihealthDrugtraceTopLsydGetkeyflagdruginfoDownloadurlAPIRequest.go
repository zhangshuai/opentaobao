package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIRequest 获取重点追溯品种明细下载URL API请求
// alibaba.alihealth.drugtrace.top.lsyd.getkeyflagdruginfo.downloadurl
//
// 获取重点追溯品种明细下载URL
type AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIRequest struct {
	model.Params
	// 调用接口的企业ID
	_refEntId string
}

// NewAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlRequest 初始化AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlRequest() *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIRequest {
	return &AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.getkeyflagdruginfo.downloadurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 调用接口的企业ID
func (r *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIRequest) GetRefEntId() string {
	return r._refEntId
}
