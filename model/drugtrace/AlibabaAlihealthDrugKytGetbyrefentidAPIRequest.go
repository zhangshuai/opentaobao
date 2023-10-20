package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytgetbyrefentidAPIRequest 根据企业唯一标识查看企业详细信息 API请求
// alibaba.alihealth.drug.kyt.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaalihealthdrugkytgetbyrefentidAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
	_destRefEntId string
}

// NewAlibabaalihealthdrugkytgetbyrefentidRequest 初始化AlibabaalihealthdrugkytgetbyrefentidAPIRequest对象
func NewAlibabaalihealthdrugkytgetbyrefentidRequest() *AlibabaalihealthdrugkytgetbyrefentidAPIRequest {
	return &AlibabaalihealthdrugkytgetbyrefentidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytgetbyrefentidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getbyrefentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytgetbyrefentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytgetbyrefentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaalihealthdrugkytgetbyrefentidAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytgetbyrefentidAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetDestRefEntId is DestRefEntId Setter
// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
func (r *AlibabaalihealthdrugkytgetbyrefentidAPIRequest) SetDestRefEntId(_destRefEntId string) error {
	r._destRefEntId = _destRefEntId
	r.Set("dest_ref_ent_id", _destRefEntId)
	return nil
}

// GetDestRefEntId DestRefEntId Getter
func (r AlibabaalihealthdrugkytgetbyrefentidAPIRequest) GetDestRefEntId() string {
	return r._destRefEntId
}
