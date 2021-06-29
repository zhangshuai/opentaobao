package c2m

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺商品接口 API请求
taobao.txp.item.itemlistget

淘小铺商品的查询服务。
*/
type TaobaoTxpItemItemlistgetRequest struct {
    model.Params
    // 第几页
    beginPage   int64
    // 每页多少条
    pageSize   int64
}

// 初始化TaobaoTxpItemItemlistgetRequest对象
func NewTaobaoTxpItemItemlistgetRequest() *TaobaoTxpItemItemlistgetRequest{
    return &TaobaoTxpItemItemlistgetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTxpItemItemlistgetRequest) GetApiMethodName() string {
    return "taobao.txp.item.itemlistget"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTxpItemItemlistgetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BeginPage Setter
// 第几页
func (r *TaobaoTxpItemItemlistgetRequest) SetBeginPage(beginPage int64) error {
    r.beginPage = beginPage
    r.Set("begin_page", beginPage)
    return nil
}

// BeginPage Getter
func (r TaobaoTxpItemItemlistgetRequest) GetBeginPage() int64 {
    return r.beginPage
}
// PageSize Setter
// 每页多少条
func (r *TaobaoTxpItemItemlistgetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTxpItemItemlistgetRequest) GetPageSize() int64 {
    return r.pageSize
}