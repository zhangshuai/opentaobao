package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseDeptSyncnewAPIRequest 直连科室上传 API请求
// alibaba.alihealth.medicalbase.dept.syncnew
//
// 直连科室上传接口
type AlibabaAlihealthMedicalbaseDeptSyncnewAPIRequest struct {
	model.Params
	// 科室状态同步
	_topChannelDeptSyncDTO *TopChannelDeptSyncDto
}

// NewAlibabaAlihealthMedicalbaseDeptSyncnewRequest 初始化AlibabaAlihealthMedicalbaseDeptSyncnewAPIRequest对象
func NewAlibabaAlihealthMedicalbaseDeptSyncnewRequest() *AlibabaAlihealthMedicalbaseDeptSyncnewAPIRequest {
	return &AlibabaAlihealthMedicalbaseDeptSyncnewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseDeptSyncnewAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.dept.syncnew"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseDeptSyncnewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalbaseDeptSyncnewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopChannelDeptSyncDTO is TopChannelDeptSyncDTO Setter
// 科室状态同步
func (r *AlibabaAlihealthMedicalbaseDeptSyncnewAPIRequest) SetTopChannelDeptSyncDTO(_topChannelDeptSyncDTO *TopChannelDeptSyncDto) error {
	r._topChannelDeptSyncDTO = _topChannelDeptSyncDTO
	r.Set("top_channel_dept_sync_d_t_o", _topChannelDeptSyncDTO)
	return nil
}

// GetTopChannelDeptSyncDTO TopChannelDeptSyncDTO Getter
func (r AlibabaAlihealthMedicalbaseDeptSyncnewAPIRequest) GetTopChannelDeptSyncDTO() *TopChannelDeptSyncDto {
	return r._topChannelDeptSyncDTO
}
