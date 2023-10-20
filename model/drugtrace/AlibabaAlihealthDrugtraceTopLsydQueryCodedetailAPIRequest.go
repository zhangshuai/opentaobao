package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest 根据码查询码信息 API请求
// alibaba.alihealth.drugtrace.top.lsyd.query.codedetail
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest struct {
	model.Params
	// 码列表【多个码用逗号拼接的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
	_codes []string
	// 企业ref_ent_id
	_refEntId string
}

// NewAlibabaalihealthdrugtracetoplsydquerycodedetailRequest 初始化AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest对象
func NewAlibabaalihealthdrugtracetoplsydquerycodedetailRequest() *AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest {
	return &AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.query.codedetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表【多个码用逗号拼接的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
func (r *AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业ref_ent_id
func (r *AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}
