package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest 根据码获取码信息-监管 API请求
// alibaba.alihealth.drug.code.list.code.supervise
//
// 服务描述
// 医保鉴证核查是基于在两定机构的药品管理（入库、出库或盘点）环节，增加扫码匹配
// 与验证鉴核流程；
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息；
// 若所传的监管码是非最小包装监管码，且存在药品混包的情况，则此接口不支持。这种
// 情况下，需要分多次调用该接口。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest struct {
	model.Params
	// 追溯码
	_codeList []string
	// ISV开放平台帐号标识
	_certIsvNo string
	// 调用方式：formal-正式、test-测试
	_invocation string
	// 终端类型 1005100-零售药店 ；10052-医疗机构    10053-建行智能pos机，10054-建行裕农通，10055-建行手机银行，10056-建行龙易行，10057-建行APP，10058-建行自助设备，10059-建行扫码设备
	_terminalType string
	// 调用零售药店名称
	_terminalName string
	// 门店所属的行政区域ID
	_bureauId string
	// 零售终端id
	_terminalEntId string
}

// NewAlibabaAlihealthDrugCodeListCodeSuperviseRequest 初始化AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest对象
func NewAlibabaAlihealthDrugCodeListCodeSuperviseRequest() *AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest {
	return &AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) Reset() {
	r._codeList = r._codeList[:0]
	r._certIsvNo = ""
	r._invocation = ""
	r._terminalType = ""
	r._terminalName = ""
	r._bureauId = ""
	r._terminalEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.list.code.supervise"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodeList is CodeList Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) SetCodeList(_codeList []string) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// GetCodeList CodeList Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) GetCodeList() []string {
	return r._codeList
}

// SetCertIsvNo is CertIsvNo Setter
// ISV开放平台帐号标识
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) SetCertIsvNo(_certIsvNo string) error {
	r._certIsvNo = _certIsvNo
	r.Set("cert_isv_no", _certIsvNo)
	return nil
}

// GetCertIsvNo CertIsvNo Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) GetCertIsvNo() string {
	return r._certIsvNo
}

// SetInvocation is Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) SetInvocation(_invocation string) error {
	r._invocation = _invocation
	r.Set("invocation", _invocation)
	return nil
}

// GetInvocation Invocation Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) GetInvocation() string {
	return r._invocation
}

// SetTerminalType is TerminalType Setter
// 终端类型 1005100-零售药店 ；10052-医疗机构    10053-建行智能pos机，10054-建行裕农通，10055-建行手机银行，10056-建行龙易行，10057-建行APP，10058-建行自助设备，10059-建行扫码设备
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetTerminalName is TerminalName Setter
// 调用零售药店名称
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// GetTerminalName TerminalName Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// SetBureauId is BureauId Setter
// 门店所属的行政区域ID
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) SetBureauId(_bureauId string) error {
	r._bureauId = _bureauId
	r.Set("bureau_id", _bureauId)
	return nil
}

// GetBureauId BureauId Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) GetBureauId() string {
	return r._bureauId
}

// SetTerminalEntId is TerminalEntId Setter
// 零售终端id
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) SetTerminalEntId(_terminalEntId string) error {
	r._terminalEntId = _terminalEntId
	r.Set("terminal_ent_id", _terminalEntId)
	return nil
}

// GetTerminalEntId TerminalEntId Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) GetTerminalEntId() string {
	return r._terminalEntId
}

var poolAlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugCodeListCodeSuperviseRequest()
	},
}

// GetAlibabaAlihealthDrugCodeListCodeSuperviseRequest 从 sync.Pool 获取 AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest
func GetAlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest() *AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest {
	return poolAlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest.Get().(*AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest)
}

// ReleaseAlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest 将 AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest(v *AlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeListCodeSuperviseAPIRequest.Put(v)
}
