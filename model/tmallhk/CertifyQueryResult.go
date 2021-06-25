package tmallhk

// CertifyQueryResult 
type CertifyQueryResult struct {

    // 清关对象
    Module   *ConsigneeCertifyInfo `json:"module,omitempty"`

    // 错误原因
    ErrMsg   string `json:"err_msg,omitempty"`

    // 错误代码
    ErrCode   string `json:"err_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}