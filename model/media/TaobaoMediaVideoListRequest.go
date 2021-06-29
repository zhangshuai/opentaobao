package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家视频列表 API请求
taobao.media.video.list

用于获取授权商家的视频列表
*/
type TaobaoMediaVideoListRequest struct {
    model.Params
    // 搜索条件
    searchCondition   *VideoSearchCondition2
}

// 初始化TaobaoMediaVideoListRequest对象
func NewTaobaoMediaVideoListRequest() *TaobaoMediaVideoListRequest{
    return &TaobaoMediaVideoListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMediaVideoListRequest) GetApiMethodName() string {
    return "taobao.media.video.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMediaVideoListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SearchCondition Setter
// 搜索条件
func (r *TaobaoMediaVideoListRequest) SetSearchCondition(searchCondition *VideoSearchCondition2) error {
    r.searchCondition = searchCondition
    r.Set("search_condition", searchCondition)
    return nil
}

// SearchCondition Getter
func (r TaobaoMediaVideoListRequest) GetSearchCondition() *VideoSearchCondition2 {
    return r.searchCondition
}
