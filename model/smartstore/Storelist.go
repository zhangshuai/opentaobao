package smartstore

// Storelist 
type Storelist struct {
    // 需要的设备code列表
    NeedDeviceCodeList   []Needdevicecodelist `json:"need_device_code_list,omitempty" xml:"need_device_code_list>needdevicecodelist,omitempty"`
    // 商场介绍
    MallIntroduce   string `json:"mall_introduce,omitempty" xml:"mall_introduce,omitempty"`
    // 商场所在省份
    MallProvince   string `json:"mall_province,omitempty" xml:"mall_province,omitempty"`
    // 门店名称
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    // 门店地址
    StoreAddress   string `json:"store_address,omitempty" xml:"store_address,omitempty"`
    // 商场地址
    MallAddress   string `json:"mall_address,omitempty" xml:"mall_address,omitempty"`
    // 商场所在区
    MallArea   string `json:"mall_area,omitempty" xml:"mall_area,omitempty"`
    // 门店id
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 商场所在城市
    MallCity   string `json:"mall_city,omitempty" xml:"mall_city,omitempty"`
    // isv appKey
    AppKey   string `json:"app_key,omitempty" xml:"app_key,omitempty"`
    // 商场名称
    MallName   string `json:"mall_name,omitempty" xml:"mall_name,omitempty"`
    // 已有的设备code列表
    HasDeviceCodeList   []Hasdevicecodelist `json:"has_device_code_list,omitempty" xml:"has_device_code_list>hasdevicecodelist,omitempty"`
}
