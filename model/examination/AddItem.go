package examination

// AddItem 
type AddItem struct {
    // 版本号，isv需要进行校验版本号是否过期，判断加项信息是否已更改，健康未同步
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
    // 加项id
    IsvItemId   string `json:"isv_item_id,omitempty" xml:"isv_item_id,omitempty"`
}
