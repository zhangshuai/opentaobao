package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse 上传用户实时位置 API返回值
// alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo
//
// 上传用户实时位置
// HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
// HSF方法名称：insertActualUserLocationInfo
type AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponseModel).Reset()
}

// AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponseModel is 上传用户实时位置 成功返回结果
type AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_adminmap_userlocationinfo_insertactualuserlocationinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse)
	},
}

// GetAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse 从 sync.Pool 获取 AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse
func GetAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse() *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse {
	return poolAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse.Get().(*AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse)
}

// ReleaseAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse 将 AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse(v *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse) {
	v.Reset()
	poolAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIResponse.Put(v)
}
