package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest 零售药店查询本企业上游企业出库单据信息 API请求
// alibaba.alihealth.drugtrace.top.lsyd.listupout
//
// 查询货主/本企业上游企业出库单据信息
type AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 开始日期（不写时分秒）
	_beginDate string
	// 结束日期（不写时分秒）
	_endDate string
	// 发货单位
	_fromUserId string
	// 生产批号
	_produceBatchNo string
	// 药品ID
	_drugEntBaseInfoId string
	// 单据类型
	_billType string
	// 药品类型
	_physicType string
	// 状态
	_status string
	// 单据号
	_billCode string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
}

// NewAlibabaAlihealthDrugtraceTopLsydListupoutRequest 初始化AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydListupoutRequest() *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest {
	return &AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest{
		Params: model.NewParams(12),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) Reset() {
	r._refEntId = ""
	r._beginDate = ""
	r._endDate = ""
	r._fromUserId = ""
	r._produceBatchNo = ""
	r._drugEntBaseInfoId = ""
	r._billType = ""
	r._physicType = ""
	r._status = ""
	r._billCode = ""
	r._pageSize = 0
	r._page = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.listupout"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBeginDate is BeginDate Setter
// 开始日期（不写时分秒）
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束日期（不写时分秒）
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetFromUserId is FromUserId Setter
// 发货单位
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetProduceBatchNo is ProduceBatchNo Setter
// 生产批号
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) SetProduceBatchNo(_produceBatchNo string) error {
	r._produceBatchNo = _produceBatchNo
	r.Set("produce_batch_no", _produceBatchNo)
	return nil
}

// GetProduceBatchNo ProduceBatchNo Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetProduceBatchNo() string {
	return r._produceBatchNo
}

// SetDrugEntBaseInfoId is DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
	r._drugEntBaseInfoId = _drugEntBaseInfoId
	r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
	return nil
}

// GetDrugEntBaseInfoId DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetDrugEntBaseInfoId() string {
	return r._drugEntBaseInfoId
}

// SetBillType is BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetBillType() string {
	return r._billType
}

// SetPhysicType is PhysicType Setter
// 药品类型
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) SetPhysicType(_physicType string) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetPhysicType() string {
	return r._physicType
}

// SetStatus is Status Setter
// 状态
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetStatus() string {
	return r._status
}

// SetBillCode is BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) GetPage() int64 {
	return r._page
}

var poolAlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopLsydListupoutRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydListupoutRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest
func GetAlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest() *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest 将 AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest(v *AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest.Put(v)
}
