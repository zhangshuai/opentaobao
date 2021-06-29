package admarket

// Monitor 
type Monitor struct {
    // 曝光地址
    ViewUrlList   []string `json:"view_url_list,omitempty" xml:"view_url_list>string,omitempty"`
    // 结束播放地址
    EndPlayUrlList   []string `json:"end_play_url_list,omitempty" xml:"end_play_url_list>string,omitempty"`
    // 开始播放地址
    StartPlayUrlList   []string `json:"start_play_url_list,omitempty" xml:"start_play_url_list>string,omitempty"`
    // 点击地址
    ClickUrlList   []string `json:"click_url_list,omitempty" xml:"click_url_list>string,omitempty"`
    // 事件上报地址
    EventUrlList   []string `json:"event_url_list,omitempty" xml:"event_url_list>string,omitempty"`
}