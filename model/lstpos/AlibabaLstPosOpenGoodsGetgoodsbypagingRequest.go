package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询用户全量的门店域商品接口(每页最多20条) API请求
alibaba.lst.pos.open.goods.getgoodsbypaging

分页查询用户全量的门店域商品接口(每页最多20条)
*/
type AlibabaLstPosOpenGoodsGetgoodsbypagingRequest struct {
    model.Params
    // 当前页码
    page   int64
    // 当前主账号userId
    userId   int64
}

// 初始化AlibabaLstPosOpenGoodsGetgoodsbypagingRequest对象
func NewAlibabaLstPosOpenGoodsGetgoodsbypagingRequest() *AlibabaLstPosOpenGoodsGetgoodsbypagingRequest{
    return &AlibabaLstPosOpenGoodsGetgoodsbypagingRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenGoodsGetgoodsbypagingRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.goods.getgoodsbypaging"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenGoodsGetgoodsbypagingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Page Setter
// 当前页码
func (r *AlibabaLstPosOpenGoodsGetgoodsbypagingRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaLstPosOpenGoodsGetgoodsbypagingRequest) GetPage() int64 {
    return r.page
}
// UserId Setter
// 当前主账号userId
func (r *AlibabaLstPosOpenGoodsGetgoodsbypagingRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaLstPosOpenGoodsGetgoodsbypagingRequest) GetUserId() int64 {
    return r.userId
}
