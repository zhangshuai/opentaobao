package film

import (
	"sync"
)

// ResultGeneralModel 结构体
type ResultGeneralModel struct {
	// 调用失败描述信息
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 0代表接口调用成功，其他值表示调用失败，错误信息可详见接口文档
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// returnUrl
	ReturnUrl string `json:"return_url,omitempty" xml:"return_url,omitempty"`
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// returnErrorOper
	ReturnErrorOper string `json:"return_error_oper,omitempty" xml:"return_error_oper,omitempty"`
	// returnErrorSolution
	ReturnErrorSolution string `json:"return_error_solution,omitempty" xml:"return_error_solution,omitempty"`
	// returnErrorStackTrace
	ReturnErrorStackTrace string `json:"return_error_stack_trace,omitempty" xml:"return_error_stack_trace,omitempty"`
	// returnValue
	ReturnValue *TopRefundOrderStatus `json:"return_value,omitempty" xml:"return_value,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolResultGeneralModel = sync.Pool{
	New: func() any {
		return new(ResultGeneralModel)
	},
}

// GetResultGeneralModel() 从对象池中获取ResultGeneralModel
func GetResultGeneralModel() *ResultGeneralModel {
	return poolResultGeneralModel.Get().(*ResultGeneralModel)
}

// ReleaseResultGeneralModel 释放ResultGeneralModel
func ReleaseResultGeneralModel(v *ResultGeneralModel) {
	v.ReturnMessage = ""
	v.ReturnCode = ""
	v.ReturnUrl = ""
	v.RequestId = ""
	v.ReturnErrorOper = ""
	v.ReturnErrorSolution = ""
	v.ReturnErrorStackTrace = ""
	v.ReturnValue = nil
	v.Success = false
	poolResultGeneralModel.Put(v)
}
