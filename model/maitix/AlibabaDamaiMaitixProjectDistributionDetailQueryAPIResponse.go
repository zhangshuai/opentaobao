package maitix

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixProjectDistributionDetailQueryAPIResponse 大麦分销项目内容详情查询 API返回值
// alibaba.damai.maitix.project.distribution.detail.query
//
// 大麦分销项目内容详情查询，提供项目的内容详情
type AlibabaDamaiMaitixProjectDistributionDetailQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixProjectDistributionDetailQueryAPIResponseModel
}

// AlibabaDamaiMaitixProjectDistributionDetailQueryAPIResponseModel is 大麦分销项目内容详情查询 成功返回结果
type AlibabaDamaiMaitixProjectDistributionDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_project_distribution_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaDamaiMaitixProjectDistributionDetailQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
