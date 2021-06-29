package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品库存只读接口(最多20条库存信息) API请求
alibaba.lst.pos.open.inventory.getinventorydata

商品库存只读接口(最多20条库存信息)
*/
type AlibabaLstPosOpenInventoryGetinventorydataRequest struct {
    model.Params
    // ISV商品Id列表
    isvGoodsIdList   []string
    // 门店主账号Id
    userId   int64
}

// 初始化AlibabaLstPosOpenInventoryGetinventorydataRequest对象
func NewAlibabaLstPosOpenInventoryGetinventorydataRequest() *AlibabaLstPosOpenInventoryGetinventorydataRequest{
    return &AlibabaLstPosOpenInventoryGetinventorydataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenInventoryGetinventorydataRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.inventory.getinventorydata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenInventoryGetinventorydataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsvGoodsIdList Setter
// ISV商品Id列表
func (r *AlibabaLstPosOpenInventoryGetinventorydataRequest) SetIsvGoodsIdList(isvGoodsIdList []string) error {
    r.isvGoodsIdList = isvGoodsIdList
    r.Set("isv_goods_id_list", isvGoodsIdList)
    return nil
}

// IsvGoodsIdList Getter
func (r AlibabaLstPosOpenInventoryGetinventorydataRequest) GetIsvGoodsIdList() []string {
    return r.isvGoodsIdList
}
// UserId Setter
// 门店主账号Id
func (r *AlibabaLstPosOpenInventoryGetinventorydataRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaLstPosOpenInventoryGetinventorydataRequest) GetUserId() int64 {
    return r.userId
}