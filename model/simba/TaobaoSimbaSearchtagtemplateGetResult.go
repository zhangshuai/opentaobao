package simba

// TaobaoSimbaSearchtagtemplateGetResult 结构体
type TaobaoSimbaSearchtagtemplateGetResult struct {
	// DimDtOs
	DimList []DimDtOs `json:"dim_list,omitempty" xml:"dim_list>dim_dt_os,omitempty"`
	// 人群模版名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 人群模版id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 人群类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
