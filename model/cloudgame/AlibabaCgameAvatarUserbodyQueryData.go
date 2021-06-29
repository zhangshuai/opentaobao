package cloudgame

// AlibabaCgameAvatarUserbodyQueryData 
type AlibabaCgameAvatarUserbodyQueryData struct {
    // body traceID
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 脸部数据
    FaceDataJson   string `json:"face_data_json,omitempty" xml:"face_data_json,omitempty"`
    // 性别， 1-male, 2-female
    Gender   int64 `json:"gender,omitempty" xml:"gender,omitempty"`
    // 请求唯一标识ID
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 用户mixUserId
    MixUserId   string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
    // 扩展数据
    ExtInfo   string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
}
