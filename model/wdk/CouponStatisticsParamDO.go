package wdk

// CouponStatisticsParamDO 
type CouponStatisticsParamDO struct {

    // 页码，即当前第几页
    PageIndex   int64 `json:"page_index,omitempty"`

    // 每页记录数，不能超过200
    PageSize   int64 `json:"page_size,omitempty"`

    // 日期，格式为yyyymmdd
    StatisticsDate   string `json:"statistics_date,omitempty"`

    // 品牌名称数组
    BrandNames   []String `json:"brand_names,omitempty"`

}