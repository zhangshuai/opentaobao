package mozi

// PageAllAccountsResult 
type PageAllAccountsResult struct {
    // 返回的记录总数
    TotalSize   int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
    // 返回的数据
    DataList   []AlibabaMoziBucAccountPageallT `json:"data_list,omitempty" xml:"data_list>alibaba_mozi_buc_account_pageall_t,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 请求的序列化
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 每页多少条数据
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 当前第几页
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    // 返回的状态消息
    ResponseMessage   string `json:"response_message,omitempty" xml:"response_message,omitempty"`
    // 返回附加值
    ResponseMetaData   string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
    // 返回码
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`
}
