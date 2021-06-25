package simba

// CreativePage 
type CreativePage struct {

    // 每页数据大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 返回第几页的数据从1开始
    PageNo   int64 `json:"page_no,omitempty"`

    // 所查询的数据总数，只有当返回第一页数据时有值，当要求返回的数据非第一页时，此返回值无效
    TotalItem   int64 `json:"total_item,omitempty"`

    // 广告创意列表
    CreativeList   []Creative `json:"creative_list,omitempty"`

}