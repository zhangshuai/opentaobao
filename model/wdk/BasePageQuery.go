package wdk

// BasePageQuery 
type BasePageQuery struct {

    // 页面大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 当前分页，从1开始
    Current   int64 `json:"current,omitempty"`

}