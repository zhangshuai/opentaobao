package axindata

// TaobaoalitriptravelfscrouteapiprojectaddTopResult 结构体
type TaobaoalitriptravelfscrouteapiprojectaddTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 业务数据
	FscProjectAddApiResponse *FscProjectAddApiResponse `json:"fsc_project_add_api_response,omitempty" xml:"fsc_project_add_api_response,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}