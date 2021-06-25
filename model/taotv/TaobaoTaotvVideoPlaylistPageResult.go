package taotv

// TaobaoTaotvVideoPlaylistPageResult 
type TaobaoTaotvVideoPlaylistPageResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // 状态码
    HttpStatusCode   int64 `json:"http_status_code,omitempty"`

    // 播单信息分页信息
    Model   *TaobaoTaotvVideoPlaylistPageModel `json:"model,omitempty"`

}