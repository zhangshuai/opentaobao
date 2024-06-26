package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest 导出项目和药品目录 API请求
// alibaba.alihealth.drugcode.drugfactory.exportproject
//
// 导出临床项目及其药品目录
type AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
}

// NewAlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest 初始化AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest {
	return &AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest) Reset() {
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.exportproject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest()
	},
}

// GetAlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest 从 sync.Pool 获取 AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest
func GetAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest {
	return poolAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest.Get().(*AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest)
}

// ReleaseAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest 将 AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest(v *AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest.Put(v)
}
