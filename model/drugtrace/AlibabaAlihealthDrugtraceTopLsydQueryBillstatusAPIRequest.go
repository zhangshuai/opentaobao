package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest 上传单据后处理状态查询 API请求
// alibaba.alihealth.drugtrace.top.lsyd.query.billstatus
//
// 单据处理状态查询
type AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest struct {
	model.Params
	// 企业ref_ent_id
	_refEntId string
	// 开始日期（没有时分秒，【单据创建时间】）
	_beginDate string
	// 结束日期（没有时分秒，【单据创建时间】）
	_endDate string
	// 单据类型 A：全部 AI：全部入库 AO：全部出库
	_billType string
	// 单据号
	_billCode string
	// 药品类型
	_drugType string
	// 状态  0, 上传成功     3, 处理成功     4, 处理失败
	_dealStatus string
	// 发货企业【注：是ent_id,不是ref_ent_id】
	_fromUserId string
	// 收货企业【注：是ent_id,不是ref_ent_id】
	_toUserId string
	// 代理商
	_agentRefUserId string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
}

// NewAlibabaalihealthdrugtracetoplsydquerybillstatusRequest 初始化AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest对象
func NewAlibabaalihealthdrugtracetoplsydquerybillstatusRequest() *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest {
	return &AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.query.billstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ref_ent_id
func (r *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBeginDate is BeginDate Setter
// 开始日期（没有时分秒，【单据创建时间】）
func (r *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束日期（没有时分秒，【单据创建时间】）
func (r *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetBillType is BillType Setter
// 单据类型 A：全部 AI：全部入库 AO：全部出库
func (r *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetBillType() string {
	return r._billType
}

// SetBillCode is BillCode Setter
// 单据号
func (r *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetDrugType is DrugType Setter
// 药品类型
func (r *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) SetDrugType(_drugType string) error {
	r._drugType = _drugType
	r.Set("drug_type", _drugType)
	return nil
}

// GetDrugType DrugType Getter
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetDrugType() string {
	return r._drugType
}

// SetDealStatus is DealStatus Setter
// 状态  0, 上传成功     3, 处理成功     4, 处理失败
func (r *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) SetDealStatus(_dealStatus string) error {
	r._dealStatus = _dealStatus
	r.Set("deal_status", _dealStatus)
	return nil
}

// GetDealStatus DealStatus Getter
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetDealStatus() string {
	return r._dealStatus
}

// SetFromUserId is FromUserId Setter
// 发货企业【注：是ent_id,不是ref_ent_id】
func (r *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetToUserId is ToUserId Setter
// 收货企业【注：是ent_id,不是ref_ent_id】
func (r *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) SetToUserId(_toUserId string) error {
	r._toUserId = _toUserId
	r.Set("to_user_id", _toUserId)
	return nil
}

// GetToUserId ToUserId Getter
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetToUserId() string {
	return r._toUserId
}

// SetAgentRefUserId is AgentRefUserId Setter
// 代理商
func (r *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// GetAgentRefUserId AgentRefUserId Getter
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest) GetPage() int64 {
	return r._page
}
