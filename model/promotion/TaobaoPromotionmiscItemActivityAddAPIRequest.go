package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建无条件单品优惠活动 API请求
taobao.promotionmisc.item.activity.add

创建无条件单品优惠活动。1、可以选择是全店参加或者部分商品参加：participate_range：0表示全部参与； 1表示部分商品参与。<br/>2、如果是部分商品参加，则需要通过taobao.promotionmisc.activity.range.add接口来指定需要参加的商品。<br/>3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
*/
type TaobaoPromotionmiscItemActivityAddAPIRequest struct {
    model.Params
    // 活动名称，超过5个汉字时，商品详情中显示的优惠名称为：卖家优惠。
    _name   string
    // 活动范围：0表示全部参与； 1表示部分商品参与。
    _participateRange   int64
    // 活动开始时间。
    _startTime   string
    // 活动结束时间。
    _endTime   string
    // 是否指定用户标签。
    _isUserTag   bool
    // 用户标签。当is_user_tag为true时，该值才有意义。
    _userTag   string
    // 是否有减钱行为。
    _isDecreaseMoney   bool
    // 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
    _decreaseAmount   int64
    // 是否有打折行为。
    _isDiscount   bool
    // 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
    _discountRate   int64
}

// 初始化TaobaoPromotionmiscItemActivityAddAPIRequest对象
func NewTaobaoPromotionmiscItemActivityAddRequest() *TaobaoPromotionmiscItemActivityAddAPIRequest{
    return &TaobaoPromotionmiscItemActivityAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscItemActivityAddAPIRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.item.activity.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscItemActivityAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 活动名称，超过5个汉字时，商品详情中显示的优惠名称为：卖家优惠。
func (r *TaobaoPromotionmiscItemActivityAddAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoPromotionmiscItemActivityAddAPIRequest) GetName() string {
    return r._name
}
// ParticipateRange Setter
// 活动范围：0表示全部参与； 1表示部分商品参与。
func (r *TaobaoPromotionmiscItemActivityAddAPIRequest) SetParticipateRange(_participateRange int64) error {
    r._participateRange = _participateRange
    r.Set("participate_range", _participateRange)
    return nil
}

// ParticipateRange Getter
func (r TaobaoPromotionmiscItemActivityAddAPIRequest) GetParticipateRange() int64 {
    return r._participateRange
}
// StartTime Setter
// 活动开始时间。
func (r *TaobaoPromotionmiscItemActivityAddAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoPromotionmiscItemActivityAddAPIRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 活动结束时间。
func (r *TaobaoPromotionmiscItemActivityAddAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoPromotionmiscItemActivityAddAPIRequest) GetEndTime() string {
    return r._endTime
}
// IsUserTag Setter
// 是否指定用户标签。
func (r *TaobaoPromotionmiscItemActivityAddAPIRequest) SetIsUserTag(_isUserTag bool) error {
    r._isUserTag = _isUserTag
    r.Set("is_user_tag", _isUserTag)
    return nil
}

// IsUserTag Getter
func (r TaobaoPromotionmiscItemActivityAddAPIRequest) GetIsUserTag() bool {
    return r._isUserTag
}
// UserTag Setter
// 用户标签。当is_user_tag为true时，该值才有意义。
func (r *TaobaoPromotionmiscItemActivityAddAPIRequest) SetUserTag(_userTag string) error {
    r._userTag = _userTag
    r.Set("user_tag", _userTag)
    return nil
}

// UserTag Getter
func (r TaobaoPromotionmiscItemActivityAddAPIRequest) GetUserTag() string {
    return r._userTag
}
// IsDecreaseMoney Setter
// 是否有减钱行为。
func (r *TaobaoPromotionmiscItemActivityAddAPIRequest) SetIsDecreaseMoney(_isDecreaseMoney bool) error {
    r._isDecreaseMoney = _isDecreaseMoney
    r.Set("is_decrease_money", _isDecreaseMoney)
    return nil
}

// IsDecreaseMoney Getter
func (r TaobaoPromotionmiscItemActivityAddAPIRequest) GetIsDecreaseMoney() bool {
    return r._isDecreaseMoney
}
// DecreaseAmount Setter
// 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
func (r *TaobaoPromotionmiscItemActivityAddAPIRequest) SetDecreaseAmount(_decreaseAmount int64) error {
    r._decreaseAmount = _decreaseAmount
    r.Set("decrease_amount", _decreaseAmount)
    return nil
}

// DecreaseAmount Getter
func (r TaobaoPromotionmiscItemActivityAddAPIRequest) GetDecreaseAmount() int64 {
    return r._decreaseAmount
}
// IsDiscount Setter
// 是否有打折行为。
func (r *TaobaoPromotionmiscItemActivityAddAPIRequest) SetIsDiscount(_isDiscount bool) error {
    r._isDiscount = _isDiscount
    r.Set("is_discount", _isDiscount)
    return nil
}

// IsDiscount Getter
func (r TaobaoPromotionmiscItemActivityAddAPIRequest) GetIsDiscount() bool {
    return r._isDiscount
}
// DiscountRate Setter
// 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
func (r *TaobaoPromotionmiscItemActivityAddAPIRequest) SetDiscountRate(_discountRate int64) error {
    r._discountRate = _discountRate
    r.Set("discount_rate", _discountRate)
    return nil
}

// DiscountRate Getter
func (r TaobaoPromotionmiscItemActivityAddAPIRequest) GetDiscountRate() int64 {
    return r._discountRate
}