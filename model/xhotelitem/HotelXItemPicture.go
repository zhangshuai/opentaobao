package xhotelitem

// HotelXitemPicture 结构体
type HotelXitemPicture struct {
	// 图片地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 是否主图
	IsMain bool `json:"is_main,omitempty" xml:"is_main,omitempty"`
}