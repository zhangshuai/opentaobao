package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
系统升级查询 API返回值 
yunos.tvpubadmin.device.osupgradequery

系统升级查询
*/
type YunosTvpubadminDeviceOsupgradequeryAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceOsupgradequeryResponse
}

// 系统升级查询 成功返回结果
type YunosTvpubadminDeviceOsupgradequeryResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_osupgradequery_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 信息结构
    ObjectList   *PaginationDO `json:"object_list,omitempty" xml:"object_list,omitempty"`
}
