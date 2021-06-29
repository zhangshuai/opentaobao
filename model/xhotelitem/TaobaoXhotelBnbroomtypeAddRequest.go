package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
民宿新增房源 API请求
taobao.xhotel.bnbroomtype.add

添加民宿房源
*/
type TaobaoXhotelBnbroomtypeAddRequest struct {
    model.Params
    // 销售渠道,默认taobao
    vendor   string
    // 房型id, 这是卖家自己系统中的ID
    outerId   string
    // 外部门店id
    outHid   string
    // 房型名
    name   string
    // 房型英文名
    nameE   string
    // 房型类型,见https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    productType   int64
    // 有无资质执照 0 没有 1有
    hasLicense   int64
    // 单间面积，单位平方米
    houseSize   int64
    // 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
    floor   string
    // 房型介绍
    introduction   string
    // 亮点描述
    brightspot   string
    // 位置描述
    localInfo   string
    // 周边描述
    surroundInfo   string
    // 品牌名称
    brand   string
    // 开业时间，格式为2015-01-01
    openingTime   string
    // 装修时间，格式为2015-01-01装修时间
    decorateTime   string
    // 装修等级 1 精装；2普通；3简装
    decorateLevel   int64
    // 装修风格https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    decorateStyle   int64
    // 风景类型(枚举)https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    scenicFeature   int64
    // 出租类型，1整租；2分租。3床位 默认整租，该字段不能更新
    rentType   int64
    // 单间面积，单位平方米
    rentSize   int64
    // 是否与房东同住 0 不同住 1同住
    hasLandlord   int64
    // 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
    tel   string
    // 真实联系方式
    realTel   string
    // 状态 0：在线 -1：不在线 -2:停售
    status   *model.File
    // 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
    settlementCurrency   string
    // 是否支持IM聊天 0不支持 1支持
    supportIm   int64
    // 是否开启闪订 0不开启 1开启
    quickOrder   int64
    // 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347
    bedInfo   string
    // 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房,bedroom和livingroom不能为空
    houseModel   string
    // 窗型-1.有窗；2.无窗；3.部分有窗
    windowType   int64
    // 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
    pics   []BnbPictureDTO
    // 房型外部标签 标签信息，逗号(,)分隔
    outerTags   string
    // 是否使用实拍图片 0不使用 1使用
    isUseShootImage   int64
    // 视频地址
    videoUrl   string
    // 民宿房源位置信息
    location   *BnbLocationDto
    // 最大入住人数 1-50
    maxOccupancy   int64
    // 民宿入住要求&附加信息
    bnbBookingTime   *BnbBookingTimeDto
    // 入住须知
    checkInNotes   string
    // 0：不限制，1：只限男性，2：只限女性'
    guestGender   int64
    // 是否接待儿童、老人；成年人必接待，详见“可接待客人”https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    guestAge   int64
    // 是否可接待外宾 0：否 1：是
    receiveForeigners   int64
    // “打扫类型1(1客1扫/换),2(1天1扫/换),https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    cleaningFrequency   int64
    // 详见“允许活动”：https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    activitiesAllowed   string
    // 0-n；若不可加床，值为0
    extraBedsNum   int64
    // 押金类型0.线下；1.线上
    depositType   int64
    // 是否信用免押金0：否 1：是
    supportcredit   int64
    // 押金金额
    depositAmount   int64
    // 加人收费信息
    charge   *BnbChargeDto
    // 清洁费是否收取 0：否 1：是
    cleaningCharge   int64
    // 清洁费类型 0.线下；1.线上
    cleaningType   int64
    // 清洁费金额；整数[1，9999999]
    extraCleaningCharge   int64
    // 发票，0：卖家提供发票，1：房东提供发票
    invoice   int64
    // 可提供发票类型，1.专票 2.纸质普票 3.电子普票
    invoiceType   int64
    // 是否有前台 0没有 1有
    hasFrontDesk   int64
    // 如果要变更商品房型编码请使用该字段。
    newOuterId   string
    // 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}，见https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    service   string
}

// 初始化TaobaoXhotelBnbroomtypeAddRequest对象
func NewTaobaoXhotelBnbroomtypeAddRequest() *TaobaoXhotelBnbroomtypeAddRequest{
    return &TaobaoXhotelBnbroomtypeAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbroomtypeAddRequest) GetApiMethodName() string {
    return "taobao.xhotel.bnbroomtype.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbroomtypeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Vendor Setter
// 销售渠道,默认taobao
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetVendor() string {
    return r.vendor
}
// OuterId Setter
// 房型id, 这是卖家自己系统中的ID
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetOuterId() string {
    return r.outerId
}
// OutHid Setter
// 外部门店id
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetOutHid(outHid string) error {
    r.outHid = outHid
    r.Set("out_hid", outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetOutHid() string {
    return r.outHid
}
// Name Setter
// 房型名
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetName() string {
    return r.name
}
// NameE Setter
// 房型英文名
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetNameE(nameE string) error {
    r.nameE = nameE
    r.Set("name_e", nameE)
    return nil
}

// NameE Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetNameE() string {
    return r.nameE
}
// ProductType Setter
// 房型类型,见https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetProductType(productType int64) error {
    r.productType = productType
    r.Set("product_type", productType)
    return nil
}

// ProductType Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetProductType() int64 {
    return r.productType
}
// HasLicense Setter
// 有无资质执照 0 没有 1有
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetHasLicense(hasLicense int64) error {
    r.hasLicense = hasLicense
    r.Set("has_license", hasLicense)
    return nil
}

// HasLicense Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetHasLicense() int64 {
    return r.hasLicense
}
// HouseSize Setter
// 单间面积，单位平方米
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetHouseSize(houseSize int64) error {
    r.houseSize = houseSize
    r.Set("house_size", houseSize)
    return nil
}

// HouseSize Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetHouseSize() int64 {
    return r.houseSize
}
// Floor Setter
// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetFloor(floor string) error {
    r.floor = floor
    r.Set("floor", floor)
    return nil
}

// Floor Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetFloor() string {
    return r.floor
}
// Introduction Setter
// 房型介绍
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetIntroduction(introduction string) error {
    r.introduction = introduction
    r.Set("introduction", introduction)
    return nil
}

// Introduction Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetIntroduction() string {
    return r.introduction
}
// Brightspot Setter
// 亮点描述
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetBrightspot(brightspot string) error {
    r.brightspot = brightspot
    r.Set("brightspot", brightspot)
    return nil
}

// Brightspot Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetBrightspot() string {
    return r.brightspot
}
// LocalInfo Setter
// 位置描述
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetLocalInfo(localInfo string) error {
    r.localInfo = localInfo
    r.Set("local_info", localInfo)
    return nil
}

// LocalInfo Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetLocalInfo() string {
    return r.localInfo
}
// SurroundInfo Setter
// 周边描述
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetSurroundInfo(surroundInfo string) error {
    r.surroundInfo = surroundInfo
    r.Set("surround_info", surroundInfo)
    return nil
}

// SurroundInfo Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetSurroundInfo() string {
    return r.surroundInfo
}
// Brand Setter
// 品牌名称
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetBrand(brand string) error {
    r.brand = brand
    r.Set("brand", brand)
    return nil
}

// Brand Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetBrand() string {
    return r.brand
}
// OpeningTime Setter
// 开业时间，格式为2015-01-01
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetOpeningTime(openingTime string) error {
    r.openingTime = openingTime
    r.Set("opening_time", openingTime)
    return nil
}

// OpeningTime Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetOpeningTime() string {
    return r.openingTime
}
// DecorateTime Setter
// 装修时间，格式为2015-01-01装修时间
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetDecorateTime(decorateTime string) error {
    r.decorateTime = decorateTime
    r.Set("decorate_time", decorateTime)
    return nil
}

// DecorateTime Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetDecorateTime() string {
    return r.decorateTime
}
// DecorateLevel Setter
// 装修等级 1 精装；2普通；3简装
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetDecorateLevel(decorateLevel int64) error {
    r.decorateLevel = decorateLevel
    r.Set("decorate_level", decorateLevel)
    return nil
}

// DecorateLevel Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetDecorateLevel() int64 {
    return r.decorateLevel
}
// DecorateStyle Setter
// 装修风格https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetDecorateStyle(decorateStyle int64) error {
    r.decorateStyle = decorateStyle
    r.Set("decorate_style", decorateStyle)
    return nil
}

// DecorateStyle Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetDecorateStyle() int64 {
    return r.decorateStyle
}
// ScenicFeature Setter
// 风景类型(枚举)https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetScenicFeature(scenicFeature int64) error {
    r.scenicFeature = scenicFeature
    r.Set("scenic_feature", scenicFeature)
    return nil
}

// ScenicFeature Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetScenicFeature() int64 {
    return r.scenicFeature
}
// RentType Setter
// 出租类型，1整租；2分租。3床位 默认整租，该字段不能更新
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetRentType(rentType int64) error {
    r.rentType = rentType
    r.Set("rent_type", rentType)
    return nil
}

// RentType Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetRentType() int64 {
    return r.rentType
}
// RentSize Setter
// 单间面积，单位平方米
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetRentSize(rentSize int64) error {
    r.rentSize = rentSize
    r.Set("rent_size", rentSize)
    return nil
}

// RentSize Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetRentSize() int64 {
    return r.rentSize
}
// HasLandlord Setter
// 是否与房东同住 0 不同住 1同住
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetHasLandlord(hasLandlord int64) error {
    r.hasLandlord = hasLandlord
    r.Set("has_landlord", hasLandlord)
    return nil
}

// HasLandlord Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetHasLandlord() int64 {
    return r.hasLandlord
}
// Tel Setter
// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetTel(tel string) error {
    r.tel = tel
    r.Set("tel", tel)
    return nil
}

// Tel Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetTel() string {
    return r.tel
}
// RealTel Setter
// 真实联系方式
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetRealTel(realTel string) error {
    r.realTel = realTel
    r.Set("real_tel", realTel)
    return nil
}

// RealTel Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetRealTel() string {
    return r.realTel
}
// Status Setter
// 状态 0：在线 -1：不在线 -2:停售
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetStatus(status *model.File) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetStatus() *model.File {
    return r.status
}
// SettlementCurrency Setter
// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetSettlementCurrency(settlementCurrency string) error {
    r.settlementCurrency = settlementCurrency
    r.Set("settlement_currency", settlementCurrency)
    return nil
}

// SettlementCurrency Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetSettlementCurrency() string {
    return r.settlementCurrency
}
// SupportIm Setter
// 是否支持IM聊天 0不支持 1支持
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetSupportIm(supportIm int64) error {
    r.supportIm = supportIm
    r.Set("support_im", supportIm)
    return nil
}

// SupportIm Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetSupportIm() int64 {
    return r.supportIm
}
// QuickOrder Setter
// 是否开启闪订 0不开启 1开启
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetQuickOrder(quickOrder int64) error {
    r.quickOrder = quickOrder
    r.Set("quick_order", quickOrder)
    return nil
}

// QuickOrder Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetQuickOrder() int64 {
    return r.quickOrder
}
// BedInfo Setter
// 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetBedInfo(bedInfo string) error {
    r.bedInfo = bedInfo
    r.Set("bed_info", bedInfo)
    return nil
}

// BedInfo Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetBedInfo() string {
    return r.bedInfo
}
// HouseModel Setter
// 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房,bedroom和livingroom不能为空
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetHouseModel(houseModel string) error {
    r.houseModel = houseModel
    r.Set("house_model", houseModel)
    return nil
}

// HouseModel Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetHouseModel() string {
    return r.houseModel
}
// WindowType Setter
// 窗型-1.有窗；2.无窗；3.部分有窗
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetWindowType(windowType int64) error {
    r.windowType = windowType
    r.Set("window_type", windowType)
    return nil
}

// WindowType Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetWindowType() int64 {
    return r.windowType
}
// Pics Setter
// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetPics(pics []BnbPictureDTO) error {
    r.pics = pics
    r.Set("pics", pics)
    return nil
}

// Pics Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetPics() []BnbPictureDTO {
    return r.pics
}
// OuterTags Setter
// 房型外部标签 标签信息，逗号(,)分隔
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetOuterTags(outerTags string) error {
    r.outerTags = outerTags
    r.Set("outer_tags", outerTags)
    return nil
}

// OuterTags Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetOuterTags() string {
    return r.outerTags
}
// IsUseShootImage Setter
// 是否使用实拍图片 0不使用 1使用
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetIsUseShootImage(isUseShootImage int64) error {
    r.isUseShootImage = isUseShootImage
    r.Set("is_use_shoot_image", isUseShootImage)
    return nil
}

// IsUseShootImage Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetIsUseShootImage() int64 {
    return r.isUseShootImage
}
// VideoUrl Setter
// 视频地址
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetVideoUrl(videoUrl string) error {
    r.videoUrl = videoUrl
    r.Set("video_url", videoUrl)
    return nil
}

// VideoUrl Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetVideoUrl() string {
    return r.videoUrl
}
// Location Setter
// 民宿房源位置信息
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetLocation(location *BnbLocationDto) error {
    r.location = location
    r.Set("location", location)
    return nil
}

// Location Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetLocation() *BnbLocationDto {
    return r.location
}
// MaxOccupancy Setter
// 最大入住人数 1-50
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetMaxOccupancy(maxOccupancy int64) error {
    r.maxOccupancy = maxOccupancy
    r.Set("max_occupancy", maxOccupancy)
    return nil
}

// MaxOccupancy Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetMaxOccupancy() int64 {
    return r.maxOccupancy
}
// BnbBookingTime Setter
// 民宿入住要求&附加信息
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetBnbBookingTime(bnbBookingTime *BnbBookingTimeDto) error {
    r.bnbBookingTime = bnbBookingTime
    r.Set("bnb_booking_time", bnbBookingTime)
    return nil
}

// BnbBookingTime Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetBnbBookingTime() *BnbBookingTimeDto {
    return r.bnbBookingTime
}
// CheckInNotes Setter
// 入住须知
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetCheckInNotes(checkInNotes string) error {
    r.checkInNotes = checkInNotes
    r.Set("check_in_notes", checkInNotes)
    return nil
}

// CheckInNotes Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetCheckInNotes() string {
    return r.checkInNotes
}
// GuestGender Setter
// 0：不限制，1：只限男性，2：只限女性'
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetGuestGender(guestGender int64) error {
    r.guestGender = guestGender
    r.Set("guest_gender", guestGender)
    return nil
}

// GuestGender Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetGuestGender() int64 {
    return r.guestGender
}
// GuestAge Setter
// 是否接待儿童、老人；成年人必接待，详见“可接待客人”https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetGuestAge(guestAge int64) error {
    r.guestAge = guestAge
    r.Set("guest_age", guestAge)
    return nil
}

// GuestAge Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetGuestAge() int64 {
    return r.guestAge
}
// ReceiveForeigners Setter
// 是否可接待外宾 0：否 1：是
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetReceiveForeigners(receiveForeigners int64) error {
    r.receiveForeigners = receiveForeigners
    r.Set("receive_foreigners", receiveForeigners)
    return nil
}

// ReceiveForeigners Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetReceiveForeigners() int64 {
    return r.receiveForeigners
}
// CleaningFrequency Setter
// “打扫类型1(1客1扫/换),2(1天1扫/换),https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetCleaningFrequency(cleaningFrequency int64) error {
    r.cleaningFrequency = cleaningFrequency
    r.Set("cleaning_frequency", cleaningFrequency)
    return nil
}

// CleaningFrequency Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetCleaningFrequency() int64 {
    return r.cleaningFrequency
}
// ActivitiesAllowed Setter
// 详见“允许活动”：https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetActivitiesAllowed(activitiesAllowed string) error {
    r.activitiesAllowed = activitiesAllowed
    r.Set("activities_allowed", activitiesAllowed)
    return nil
}

// ActivitiesAllowed Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetActivitiesAllowed() string {
    return r.activitiesAllowed
}
// ExtraBedsNum Setter
// 0-n；若不可加床，值为0
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetExtraBedsNum(extraBedsNum int64) error {
    r.extraBedsNum = extraBedsNum
    r.Set("extra_beds_num", extraBedsNum)
    return nil
}

// ExtraBedsNum Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetExtraBedsNum() int64 {
    return r.extraBedsNum
}
// DepositType Setter
// 押金类型0.线下；1.线上
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetDepositType(depositType int64) error {
    r.depositType = depositType
    r.Set("deposit_type", depositType)
    return nil
}

// DepositType Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetDepositType() int64 {
    return r.depositType
}
// Supportcredit Setter
// 是否信用免押金0：否 1：是
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetSupportcredit(supportcredit int64) error {
    r.supportcredit = supportcredit
    r.Set("supportcredit", supportcredit)
    return nil
}

// Supportcredit Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetSupportcredit() int64 {
    return r.supportcredit
}
// DepositAmount Setter
// 押金金额
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetDepositAmount(depositAmount int64) error {
    r.depositAmount = depositAmount
    r.Set("deposit_amount", depositAmount)
    return nil
}

// DepositAmount Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetDepositAmount() int64 {
    return r.depositAmount
}
// Charge Setter
// 加人收费信息
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetCharge(charge *BnbChargeDto) error {
    r.charge = charge
    r.Set("charge", charge)
    return nil
}

// Charge Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetCharge() *BnbChargeDto {
    return r.charge
}
// CleaningCharge Setter
// 清洁费是否收取 0：否 1：是
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetCleaningCharge(cleaningCharge int64) error {
    r.cleaningCharge = cleaningCharge
    r.Set("cleaning_charge", cleaningCharge)
    return nil
}

// CleaningCharge Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetCleaningCharge() int64 {
    return r.cleaningCharge
}
// CleaningType Setter
// 清洁费类型 0.线下；1.线上
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetCleaningType(cleaningType int64) error {
    r.cleaningType = cleaningType
    r.Set("cleaning_type", cleaningType)
    return nil
}

// CleaningType Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetCleaningType() int64 {
    return r.cleaningType
}
// ExtraCleaningCharge Setter
// 清洁费金额；整数[1，9999999]
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetExtraCleaningCharge(extraCleaningCharge int64) error {
    r.extraCleaningCharge = extraCleaningCharge
    r.Set("extra_cleaning_charge", extraCleaningCharge)
    return nil
}

// ExtraCleaningCharge Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetExtraCleaningCharge() int64 {
    return r.extraCleaningCharge
}
// Invoice Setter
// 发票，0：卖家提供发票，1：房东提供发票
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetInvoice(invoice int64) error {
    r.invoice = invoice
    r.Set("invoice", invoice)
    return nil
}

// Invoice Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetInvoice() int64 {
    return r.invoice
}
// InvoiceType Setter
// 可提供发票类型，1.专票 2.纸质普票 3.电子普票
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetInvoiceType(invoiceType int64) error {
    r.invoiceType = invoiceType
    r.Set("invoice_type", invoiceType)
    return nil
}

// InvoiceType Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetInvoiceType() int64 {
    return r.invoiceType
}
// HasFrontDesk Setter
// 是否有前台 0没有 1有
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetHasFrontDesk(hasFrontDesk int64) error {
    r.hasFrontDesk = hasFrontDesk
    r.Set("has_front_desk", hasFrontDesk)
    return nil
}

// HasFrontDesk Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetHasFrontDesk() int64 {
    return r.hasFrontDesk
}
// NewOuterId Setter
// 如果要变更商品房型编码请使用该字段。
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetNewOuterId(newOuterId string) error {
    r.newOuterId = newOuterId
    r.Set("new_outer_id", newOuterId)
    return nil
}

// NewOuterId Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetNewOuterId() string {
    return r.newOuterId
}
// Service Setter
// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}，见https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetService(service string) error {
    r.service = service
    r.Set("service", service)
    return nil
}

// Service Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetService() string {
    return r.service
}