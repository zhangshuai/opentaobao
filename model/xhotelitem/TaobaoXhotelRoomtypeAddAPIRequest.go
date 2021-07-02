package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomtypeAddAPIRequest 房型新增接口（ID重复变更新） API请求
// taobao.xhotel.roomtype.add
//
// 房型添加或更新
type TaobaoXhotelRoomtypeAddAPIRequest struct {
	model.Params
	// 卖家房型ID，不能重复建议格式是:酒店code_房型code
	_outerId string
	// （已废弃）请使用outHid
	_hid int64
	// 房型名称。不能超过30字
	_name string
	// 最大入住人数，默认2（1-99）
	_maxOccupancy int64
	// 具体面积大小，请按照正确格式填写。两种格式，例如：40或者 10-20
	_area string
	// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
	_floor string
	// 床型。按链接中床型列表定义值存储 http://open.taobao.com/docs/doc.htm?&docType=1&articleId=105610
	_bedType string
	// 床宽。按自己定义存储，比如：2.1米
	_bedSize string
	// 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
	_internet string
	// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}
	_service string
	// 不要使用
	_extend string
	// 0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗
	_windowType int64
	// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
	_srid int64
	// （必传）商家酒店ID，指明该房型属于哪家酒店
	_outHid string
	// 系统商，无申请不可使用
	_vendor string
	// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
	_pics string
	// 卖家房型英文名称
	_nameE string
	// 操作人信息
	_operator string
	// 属性值为1: 含义是非直连房型
	_connectionType int64
	// main_bed_type母床型,sub_bed_type子床型。详情参见文档： https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.yN2mES&docType=1&articleId=118712&previewCode=1DABB73EA935608455E203BA06CF3566
	_bedInfo string
	// 酒店房型设施
	_standardRoomFacilities string
}

// NewTaobaoXhotelRoomtypeAddRequest 初始化TaobaoXhotelRoomtypeAddAPIRequest对象
func NewTaobaoXhotelRoomtypeAddRequest() *TaobaoXhotelRoomtypeAddAPIRequest {
	return &TaobaoXhotelRoomtypeAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.roomtype.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OuterId Setter
// 卖家房型ID，不能重复建议格式是:酒店code_房型code
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is Hid Setter
// （已废弃）请使用outHid
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// Get Hid Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetHid() int64 {
	return r._hid
}

// Set is Name Setter
// 房型名称。不能超过30字
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetName() string {
	return r._name
}

// Set is MaxOccupancy Setter
// 最大入住人数，默认2（1-99）
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetMaxOccupancy(_maxOccupancy int64) error {
	r._maxOccupancy = _maxOccupancy
	r.Set("max_occupancy", _maxOccupancy)
	return nil
}

// Get MaxOccupancy Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetMaxOccupancy() int64 {
	return r._maxOccupancy
}

// Set is Area Setter
// 具体面积大小，请按照正确格式填写。两种格式，例如：40或者 10-20
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetArea(_area string) error {
	r._area = _area
	r.Set("area", _area)
	return nil
}

// Get Area Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetArea() string {
	return r._area
}

// Set is Floor Setter
// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetFloor(_floor string) error {
	r._floor = _floor
	r.Set("floor", _floor)
	return nil
}

// Get Floor Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetFloor() string {
	return r._floor
}

// Set is BedType Setter
// 床型。按链接中床型列表定义值存储 http://open.taobao.com/docs/doc.htm?&docType=1&articleId=105610
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetBedType(_bedType string) error {
	r._bedType = _bedType
	r.Set("bed_type", _bedType)
	return nil
}

// Get BedType Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetBedType() string {
	return r._bedType
}

// Set is BedSize Setter
// 床宽。按自己定义存储，比如：2.1米
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetBedSize(_bedSize string) error {
	r._bedSize = _bedSize
	r.Set("bed_size", _bedSize)
	return nil
}

// Get BedSize Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetBedSize() string {
	return r._bedSize
}

// Set is Internet Setter
// 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetInternet(_internet string) error {
	r._internet = _internet
	r.Set("internet", _internet)
	return nil
}

// Get Internet Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetInternet() string {
	return r._internet
}

// Set is Service Setter
// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetService(_service string) error {
	r._service = _service
	r.Set("service", _service)
	return nil
}

// Get Service Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetService() string {
	return r._service
}

// Set is Extend Setter
// 不要使用
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// Get Extend Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetExtend() string {
	return r._extend
}

// Set is WindowType Setter
// 0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetWindowType(_windowType int64) error {
	r._windowType = _windowType
	r.Set("window_type", _windowType)
	return nil
}

// Get WindowType Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetWindowType() int64 {
	return r._windowType
}

// Set is Srid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetSrid(_srid int64) error {
	r._srid = _srid
	r.Set("srid", _srid)
	return nil
}

// Get Srid Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetSrid() int64 {
	return r._srid
}

// Set is OutHid Setter
// （必传）商家酒店ID，指明该房型属于哪家酒店
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// Get OutHid Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetOutHid() string {
	return r._outHid
}

// Set is Vendor Setter
// 系统商，无申请不可使用
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// Get Vendor Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetVendor() string {
	return r._vendor
}

// Set is Pics Setter
// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetPics(_pics string) error {
	r._pics = _pics
	r.Set("pics", _pics)
	return nil
}

// Get Pics Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetPics() string {
	return r._pics
}

// Set is NameE Setter
// 卖家房型英文名称
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetNameE(_nameE string) error {
	r._nameE = _nameE
	r.Set("name_e", _nameE)
	return nil
}

// Get NameE Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetNameE() string {
	return r._nameE
}

// Set is Operator Setter
// 操作人信息
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// Get Operator Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetOperator() string {
	return r._operator
}

// Set is ConnectionType Setter
// 属性值为1: 含义是非直连房型
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetConnectionType(_connectionType int64) error {
	r._connectionType = _connectionType
	r.Set("connection_type", _connectionType)
	return nil
}

// Get ConnectionType Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetConnectionType() int64 {
	return r._connectionType
}

// Set is BedInfo Setter
// main_bed_type母床型,sub_bed_type子床型。详情参见文档： https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.yN2mES&docType=1&articleId=118712&previewCode=1DABB73EA935608455E203BA06CF3566
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetBedInfo(_bedInfo string) error {
	r._bedInfo = _bedInfo
	r.Set("bed_info", _bedInfo)
	return nil
}

// Get BedInfo Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetBedInfo() string {
	return r._bedInfo
}

// Set is StandardRoomFacilities Setter
// 酒店房型设施
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetStandardRoomFacilities(_standardRoomFacilities string) error {
	r._standardRoomFacilities = _standardRoomFacilities
	r.Set("standard_room_facilities", _standardRoomFacilities)
	return nil
}

// Get StandardRoomFacilities Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetStandardRoomFacilities() string {
	return r._standardRoomFacilities
}
