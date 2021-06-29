package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加产品规格 API请求
tmall.product.spec.add

增加产品规格
*/
type TmallProductSpecAddRequest struct {
    model.Params
    // 产品ID
    productId   int64
    // 存放产品规格认证类型-认证图片url映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为调用tmall.product.spec.pic.upload返回的认证图片url文本
    certifiedPicStr   string
    // 产品规格吊牌价，以分为单位，无默认值，上限999999999
    labelPrice   int64
    // 存放产品规格认证类型-认证文本映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为认证文本值
    certifiedTxtStr   string
    // 产品的规格属性
    specProps   string
    // 规格属性别名,只允许传颜色别名
    specPropsAlias   string
    // 用户自定义销售属性，结构：pid1:value1;pid2:value2。在
    customerSpecProps   string
    // 产品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1
    changeProp   string
    // 产品图片
    image   []*model.File
    // 产品二维码
    barcode   string
    // 产品货号
    productCode   string
    // 产品上市时间
    marketTime   string
}

// 初始化TmallProductSpecAddRequest对象
func NewTmallProductSpecAddRequest() *TmallProductSpecAddRequest{
    return &TmallProductSpecAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductSpecAddRequest) GetApiMethodName() string {
    return "tmall.product.spec.add"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductSpecAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TmallProductSpecAddRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TmallProductSpecAddRequest) GetProductId() int64 {
    return r.productId
}
// CertifiedPicStr Setter
// 存放产品规格认证类型-认证图片url映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为调用tmall.product.spec.pic.upload返回的认证图片url文本
func (r *TmallProductSpecAddRequest) SetCertifiedPicStr(certifiedPicStr string) error {
    r.certifiedPicStr = certifiedPicStr
    r.Set("certified_pic_str", certifiedPicStr)
    return nil
}

// CertifiedPicStr Getter
func (r TmallProductSpecAddRequest) GetCertifiedPicStr() string {
    return r.certifiedPicStr
}
// LabelPrice Setter
// 产品规格吊牌价，以分为单位，无默认值，上限999999999
func (r *TmallProductSpecAddRequest) SetLabelPrice(labelPrice int64) error {
    r.labelPrice = labelPrice
    r.Set("label_price", labelPrice)
    return nil
}

// LabelPrice Getter
func (r TmallProductSpecAddRequest) GetLabelPrice() int64 {
    return r.labelPrice
}
// CertifiedTxtStr Setter
// 存放产品规格认证类型-认证文本映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为认证文本值
func (r *TmallProductSpecAddRequest) SetCertifiedTxtStr(certifiedTxtStr string) error {
    r.certifiedTxtStr = certifiedTxtStr
    r.Set("certified_txt_str", certifiedTxtStr)
    return nil
}

// CertifiedTxtStr Getter
func (r TmallProductSpecAddRequest) GetCertifiedTxtStr() string {
    return r.certifiedTxtStr
}
// SpecProps Setter
// 产品的规格属性
func (r *TmallProductSpecAddRequest) SetSpecProps(specProps string) error {
    r.specProps = specProps
    r.Set("spec_props", specProps)
    return nil
}

// SpecProps Getter
func (r TmallProductSpecAddRequest) GetSpecProps() string {
    return r.specProps
}
// SpecPropsAlias Setter
// 规格属性别名,只允许传颜色别名
func (r *TmallProductSpecAddRequest) SetSpecPropsAlias(specPropsAlias string) error {
    r.specPropsAlias = specPropsAlias
    r.Set("spec_props_alias", specPropsAlias)
    return nil
}

// SpecPropsAlias Getter
func (r TmallProductSpecAddRequest) GetSpecPropsAlias() string {
    return r.specPropsAlias
}
// CustomerSpecProps Setter
// 用户自定义销售属性，结构：pid1:value1;pid2:value2。在
func (r *TmallProductSpecAddRequest) SetCustomerSpecProps(customerSpecProps string) error {
    r.customerSpecProps = customerSpecProps
    r.Set("customer_spec_props", customerSpecProps)
    return nil
}

// CustomerSpecProps Getter
func (r TmallProductSpecAddRequest) GetCustomerSpecProps() string {
    return r.customerSpecProps
}
// ChangeProp Setter
// 产品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1
func (r *TmallProductSpecAddRequest) SetChangeProp(changeProp string) error {
    r.changeProp = changeProp
    r.Set("change_prop", changeProp)
    return nil
}

// ChangeProp Getter
func (r TmallProductSpecAddRequest) GetChangeProp() string {
    return r.changeProp
}
// Image Setter
// 产品图片
func (r *TmallProductSpecAddRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

// Image Getter
func (r TmallProductSpecAddRequest) GetImage() []*model.File {
    return r.image
}
// Barcode Setter
// 产品二维码
func (r *TmallProductSpecAddRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

// Barcode Getter
func (r TmallProductSpecAddRequest) GetBarcode() string {
    return r.barcode
}
// ProductCode Setter
// 产品货号
func (r *TmallProductSpecAddRequest) SetProductCode(productCode string) error {
    r.productCode = productCode
    r.Set("product_code", productCode)
    return nil
}

// ProductCode Getter
func (r TmallProductSpecAddRequest) GetProductCode() string {
    return r.productCode
}
// MarketTime Setter
// 产品上市时间
func (r *TmallProductSpecAddRequest) SetMarketTime(marketTime string) error {
    r.marketTime = marketTime
    r.Set("market_time", marketTime)
    return nil
}

// MarketTime Getter
func (r TmallProductSpecAddRequest) GetMarketTime() string {
    return r.marketTime
}
