package aliyun

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// SlbaliyuncscomdescribeBackendServers20130221APIResponse 后端服务器健康检查 API返回值
// slb.aliyuncs.com.DescribeBackendServers.2013-02-21
//
// 后端服务器健康检查，对SLB实例的后端服务器进行健康检查，返回后端服务器的健康状况。后端服务器的健康状况为normal,abnormal和Unavailable三种。其中Unavailable表示这个SLB实例没有配置健康检查，无法获取后端服务器的健康状况。如果没有传入ListenerPort，则表示对这个SLB实例下的所有Listener后端服务器进行健康检查。
type SlbaliyuncscomdescribeBackendServers20130221APIResponse struct {
	model.CommonResponse
	SlbaliyuncscomdescribeBackendServers20130221APIResponseModel
}

// SlbaliyuncscomdescribeBackendServers20130221APIResponseModel is 后端服务器健康检查 成功返回结果
type SlbaliyuncscomdescribeBackendServers20130221APIResponseModel struct {
	XMLName xml.Name `xml:"slb_aliyuncs_com_DescribeBackendServers_2013-02-21_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Listeners
	Listeners []Listener `json:"listeners,omitempty" xml:"listeners>listener,omitempty"`
	// request id
	Requestid string `json:"requestid,omitempty" xml:"requestid,omitempty"`
	// LoadBalancerId
	Loadbalancerid string `json:"loadbalancerid,omitempty" xml:"loadbalancerid,omitempty"`
}