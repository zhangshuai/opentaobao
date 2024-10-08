package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomtypeAddAPIRequest 房型新增接口（ID重复变更新） API请求
// taobao.xhotel.roomtype.add
//
// 房型添加或更新
type TaobaoXhotelRoomtypeAddAPIRequest struct {
	model.Params
	// 房型名称。不能超过200个字符
	_name string
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
	// 卖家房型ID，不能重复建议格式是:酒店code_房型code
	_outerId string
	// 系统商，无申请不可使用
	_vendor string
	// （必传）商家酒店ID，指明该房型属于哪家酒店
	_outHid string
	// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
	_pics string
	// 卖家房型英文名称
	_nameE string
	// 操作人信息
	_operator string
	// main_bed_type母床型,sub_bed_type子床型。详情参见文档： https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.yN2mES&docType=1&articleId=118712&previewCode=1DABB73EA935608455E203BA06CF3566
	_bedInfo string
	// 酒店房型设施
	_standardRoomFacilities string
	// 窗型，窗型（windowType）： 0=无窗 1=有窗 2=部分有窗  窗型缺陷（windowTypeDefect）： 0=窗户不可打开通风 1=窗外有遮挡 2=窗外为酒店内景观  特殊窗（windowTypeSpecial）： 0=落地窗 1=飘窗 2=天窗 3=小窗。  当为有窗或部分有窗时，窗型缺陷可多选，特殊窗型需单选
	_windowDesc string
	// 房型加床政策。 费用(fee) 说明(desc)
	_addBed string
	// 儿童政策
	_childrenPolicy string
	// （已废弃）请使用outHid
	_hid int64
	// 最大入住人数，默认2（1-99）
	_maxOccupancy int64
	// 0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗
	_windowType int64
	// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
	_srid int64
	// 属性值为1: 含义是非直连房型
	_connectionType int64
}

// NewTaobaoXhotelRoomtypeAddRequest 初始化TaobaoXhotelRoomtypeAddAPIRequest对象
func NewTaobaoXhotelRoomtypeAddRequest() *TaobaoXhotelRoomtypeAddAPIRequest {
	return &TaobaoXhotelRoomtypeAddAPIRequest{
		Params: model.NewParams(24),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelRoomtypeAddAPIRequest) Reset() {
	r._name = ""
	r._area = ""
	r._floor = ""
	r._bedType = ""
	r._bedSize = ""
	r._internet = ""
	r._service = ""
	r._extend = ""
	r._outerId = ""
	r._vendor = ""
	r._outHid = ""
	r._pics = ""
	r._nameE = ""
	r._operator = ""
	r._bedInfo = ""
	r._standardRoomFacilities = ""
	r._windowDesc = ""
	r._addBed = ""
	r._childrenPolicy = ""
	r._hid = 0
	r._maxOccupancy = 0
	r._windowType = 0
	r._srid = 0
	r._connectionType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.roomtype.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 房型名称。不能超过200个字符
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetName() string {
	return r._name
}

// SetArea is Area Setter
// 具体面积大小，请按照正确格式填写。两种格式，例如：40或者 10-20
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetArea(_area string) error {
	r._area = _area
	r.Set("area", _area)
	return nil
}

// GetArea Area Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetArea() string {
	return r._area
}

// SetFloor is Floor Setter
// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetFloor(_floor string) error {
	r._floor = _floor
	r.Set("floor", _floor)
	return nil
}

// GetFloor Floor Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetFloor() string {
	return r._floor
}

// SetBedType is BedType Setter
// 床型。按链接中床型列表定义值存储 http://open.taobao.com/docs/doc.htm?&amp;docType=1&amp;articleId=105610
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetBedType(_bedType string) error {
	r._bedType = _bedType
	r.Set("bed_type", _bedType)
	return nil
}

// GetBedType BedType Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetBedType() string {
	return r._bedType
}

// SetBedSize is BedSize Setter
// 床宽。按自己定义存储，比如：2.1米
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetBedSize(_bedSize string) error {
	r._bedSize = _bedSize
	r.Set("bed_size", _bedSize)
	return nil
}

// GetBedSize BedSize Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetBedSize() string {
	return r._bedSize
}

// SetInternet is Internet Setter
// 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetInternet(_internet string) error {
	r._internet = _internet
	r.Set("internet", _internet)
	return nil
}

// GetInternet Internet Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetInternet() string {
	return r._internet
}

// SetService is Service Setter
// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {&#34;bar&#34;:false,&#34;catv&#34;:false,&#34;ddd&#34;:false,&#34;idd&#34;:false,&#34;pubtoilet&#34;:false,&#34;toilet&#34;:false}
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetService(_service string) error {
	r._service = _service
	r.Set("service", _service)
	return nil
}

// GetService Service Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetService() string {
	return r._service
}

// SetExtend is Extend Setter
// 不要使用
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetExtend() string {
	return r._extend
}

// SetOuterId is OuterId Setter
// 卖家房型ID，不能重复建议格式是:酒店code_房型code
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetVendor is Vendor Setter
// 系统商，无申请不可使用
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOutHid is OutHid Setter
// （必传）商家酒店ID，指明该房型属于哪家酒店
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetPics is Pics Setter
// 房型图片只支持远程图片，格式如下：[{&#34;url&#34;:&#34;http://taobao.com/123.jpg&#34;,&#34;ismain&#34;:&#34;true&#34;},{&#34;url&#34;:&#34;http://taobao.com/456.jpg&#34;,&#34;ismain&#34;:&#34;false&#34;},{&#34;url&#34;:&#34;http://taobao.com/789.jpg&#34;,&#34;ismain&#34;:&#34;false&#34;}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetPics(_pics string) error {
	r._pics = _pics
	r.Set("pics", _pics)
	return nil
}

// GetPics Pics Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetPics() string {
	return r._pics
}

// SetNameE is NameE Setter
// 卖家房型英文名称
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetNameE(_nameE string) error {
	r._nameE = _nameE
	r.Set("name_e", _nameE)
	return nil
}

// GetNameE NameE Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetNameE() string {
	return r._nameE
}

// SetOperator is Operator Setter
// 操作人信息
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetOperator() string {
	return r._operator
}

// SetBedInfo is BedInfo Setter
// main_bed_type母床型,sub_bed_type子床型。详情参见文档： https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.yN2mES&amp;docType=1&amp;articleId=118712&amp;previewCode=1DABB73EA935608455E203BA06CF3566
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetBedInfo(_bedInfo string) error {
	r._bedInfo = _bedInfo
	r.Set("bed_info", _bedInfo)
	return nil
}

// GetBedInfo BedInfo Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetBedInfo() string {
	return r._bedInfo
}

// SetStandardRoomFacilities is StandardRoomFacilities Setter
// 酒店房型设施
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetStandardRoomFacilities(_standardRoomFacilities string) error {
	r._standardRoomFacilities = _standardRoomFacilities
	r.Set("standard_room_facilities", _standardRoomFacilities)
	return nil
}

// GetStandardRoomFacilities StandardRoomFacilities Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetStandardRoomFacilities() string {
	return r._standardRoomFacilities
}

// SetWindowDesc is WindowDesc Setter
// 窗型，窗型（windowType）： 0=无窗 1=有窗 2=部分有窗  窗型缺陷（windowTypeDefect）： 0=窗户不可打开通风 1=窗外有遮挡 2=窗外为酒店内景观  特殊窗（windowTypeSpecial）： 0=落地窗 1=飘窗 2=天窗 3=小窗。  当为有窗或部分有窗时，窗型缺陷可多选，特殊窗型需单选
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetWindowDesc(_windowDesc string) error {
	r._windowDesc = _windowDesc
	r.Set("window_desc", _windowDesc)
	return nil
}

// GetWindowDesc WindowDesc Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetWindowDesc() string {
	return r._windowDesc
}

// SetAddBed is AddBed Setter
// 房型加床政策。 费用(fee) 说明(desc)
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetAddBed(_addBed string) error {
	r._addBed = _addBed
	r.Set("add_bed", _addBed)
	return nil
}

// GetAddBed AddBed Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetAddBed() string {
	return r._addBed
}

// SetChildrenPolicy is ChildrenPolicy Setter
// 儿童政策
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetChildrenPolicy(_childrenPolicy string) error {
	r._childrenPolicy = _childrenPolicy
	r.Set("children_policy", _childrenPolicy)
	return nil
}

// GetChildrenPolicy ChildrenPolicy Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetChildrenPolicy() string {
	return r._childrenPolicy
}

// SetHid is Hid Setter
// （已废弃）请使用outHid
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetHid() int64 {
	return r._hid
}

// SetMaxOccupancy is MaxOccupancy Setter
// 最大入住人数，默认2（1-99）
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetMaxOccupancy(_maxOccupancy int64) error {
	r._maxOccupancy = _maxOccupancy
	r.Set("max_occupancy", _maxOccupancy)
	return nil
}

// GetMaxOccupancy MaxOccupancy Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetMaxOccupancy() int64 {
	return r._maxOccupancy
}

// SetWindowType is WindowType Setter
// 0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetWindowType(_windowType int64) error {
	r._windowType = _windowType
	r.Set("window_type", _windowType)
	return nil
}

// GetWindowType WindowType Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetWindowType() int64 {
	return r._windowType
}

// SetSrid is Srid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetSrid(_srid int64) error {
	r._srid = _srid
	r.Set("srid", _srid)
	return nil
}

// GetSrid Srid Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetSrid() int64 {
	return r._srid
}

// SetConnectionType is ConnectionType Setter
// 属性值为1: 含义是非直连房型
func (r *TaobaoXhotelRoomtypeAddAPIRequest) SetConnectionType(_connectionType int64) error {
	r._connectionType = _connectionType
	r.Set("connection_type", _connectionType)
	return nil
}

// GetConnectionType ConnectionType Getter
func (r TaobaoXhotelRoomtypeAddAPIRequest) GetConnectionType() int64 {
	return r._connectionType
}

var poolTaobaoXhotelRoomtypeAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelRoomtypeAddRequest()
	},
}

// GetTaobaoXhotelRoomtypeAddRequest 从 sync.Pool 获取 TaobaoXhotelRoomtypeAddAPIRequest
func GetTaobaoXhotelRoomtypeAddAPIRequest() *TaobaoXhotelRoomtypeAddAPIRequest {
	return poolTaobaoXhotelRoomtypeAddAPIRequest.Get().(*TaobaoXhotelRoomtypeAddAPIRequest)
}

// ReleaseTaobaoXhotelRoomtypeAddAPIRequest 将 TaobaoXhotelRoomtypeAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelRoomtypeAddAPIRequest(v *TaobaoXhotelRoomtypeAddAPIRequest) {
	v.Reset()
	poolTaobaoXhotelRoomtypeAddAPIRequest.Put(v)
}
