package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerBillUploadAPIRequest 上传入出库单api API请求
// alibaba.alihealth.tracecodeseller.bill.upload
//
// 上传入出库单api
type AlibabaAlihealthTracecodesellerBillUploadAPIRequest struct {
	model.Params
	// 身份认证
	_skeyCode string
	// 商家id
	_entInfoId int64
	// 单据编号
	_billCode string
	// 出入库类型
	_type string
	// 出入库时间 精确到 年 月 日 时 分 秒
	_time string
	// 自己仓库id
	_warehouseId int64
	// 渠道商id
	_entMerchantId int64
	// 把txt格式的文件转成16进制的字符串，txt文件是每个码一行
	_codeInfo string
}

// NewAlibabaAlihealthTracecodesellerBillUploadRequest 初始化AlibabaAlihealthTracecodesellerBillUploadAPIRequest对象
func NewAlibabaAlihealthTracecodesellerBillUploadRequest() *AlibabaAlihealthTracecodesellerBillUploadAPIRequest {
	return &AlibabaAlihealthTracecodesellerBillUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.bill.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SkeyCode Setter
// 身份认证
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetSkeyCode(_skeyCode string) error {
	r._skeyCode = _skeyCode
	r.Set("skey_code", _skeyCode)
	return nil
}

// Get SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetSkeyCode() string {
	return r._skeyCode
}

// Set is EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// Get EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// Set is BillCode Setter
// 单据编号
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// Get BillCode Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetBillCode() string {
	return r._billCode
}

// Set is Type Setter
// 出入库类型
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetType() string {
	return r._type
}

// Set is Time Setter
// 出入库时间 精确到 年 月 日 时 分 秒
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetTime(_time string) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// Get Time Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetTime() string {
	return r._time
}

// Set is WarehouseId Setter
// 自己仓库id
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetWarehouseId(_warehouseId int64) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// Get WarehouseId Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetWarehouseId() int64 {
	return r._warehouseId
}

// Set is EntMerchantId Setter
// 渠道商id
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetEntMerchantId(_entMerchantId int64) error {
	r._entMerchantId = _entMerchantId
	r.Set("ent_merchant_id", _entMerchantId)
	return nil
}

// Get EntMerchantId Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetEntMerchantId() int64 {
	return r._entMerchantId
}

// Set is CodeInfo Setter
// 把txt格式的文件转成16进制的字符串，txt文件是每个码一行
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetCodeInfo(_codeInfo string) error {
	r._codeInfo = _codeInfo
	r.Set("code_info", _codeInfo)
	return nil
}

// Get CodeInfo Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetCodeInfo() string {
	return r._codeInfo
}
