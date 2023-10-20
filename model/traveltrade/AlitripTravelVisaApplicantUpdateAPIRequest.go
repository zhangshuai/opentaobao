package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelvisaapplicantupdateAPIRequest 飞猪度假-普通签证-申请人进度推进接口 API请求
// alitrip.travel.visa.applicant.update
//
// 普通签证订单的申请人进度推进接口，用于商家代用户填写申请人基本信息 或 推进单个申请人的签证进度。
type AlitriptravelvisaapplicantupdateAPIRequest struct {
	model.Params
	// 特殊必填，申请人基本信息（证件号，姓名，手机号）列表。当operType为1或2或4时必填
	_applicantInfos []NormalVisaApplicantInfo
	// 特殊必填，更多材料
	_documentInfos []NormalVisaDocumentInfo
	// 必填，子订单id
	_subOrderId string
	// 证件照文件类型
	_photoType string
	// 护照文件类型
	_passportType string
	// 酒店预订文件类型
	_hotelBookingFormType string
	// 机票预订文件类型
	_flightBookingFormType string
	// 特殊必填，签证申请人进度推进操作（目前只支持对单个申请人状态进行推进）。当operType为3时必填
	_applicantOp *NormalVisaApplicantOperation
	// 必填，操作类型。1-上传新申请人基本信息（商家代填申请人），2-更新已有申请人基本信息，3-更新已有申请人的签证进度，4-商家代传申请人信息和材料(使馆直连订单)
	_operType int64
	// 特殊必填，pdf文件字节流。用于上传电子签pdf结果 或者 预约面试信pdf文件。
	_fileBytes *model.File
	// 特殊必填，文件字节流，用于上传证件照，必须和photoType同时传
	_photoBytes *model.File
	// 特殊必填，文件字节流，用于上传护照，必须和passportType同时传
	_passportBytes *model.File
	// 特殊必填，文件字节流，用于上传酒店预订，必须和hotelBookingFormType同时传
	_hotelBookingFormBytes *model.File
	// 特殊必填，文件字节流，用于上传机票预订，必须和flightBookingFormType同时传
	_flightBookingFormBytes *model.File
}

// NewAlitriptravelvisaapplicantupdateRequest 初始化AlitriptravelvisaapplicantupdateAPIRequest对象
func NewAlitriptravelvisaapplicantupdateRequest() *AlitriptravelvisaapplicantupdateAPIRequest {
	return &AlitriptravelvisaapplicantupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.visa.applicant.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplicantInfos is ApplicantInfos Setter
// 特殊必填，申请人基本信息（证件号，姓名，手机号）列表。当operType为1或2或4时必填
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetApplicantInfos(_applicantInfos []NormalVisaApplicantInfo) error {
	r._applicantInfos = _applicantInfos
	r.Set("applicant_infos", _applicantInfos)
	return nil
}

// GetApplicantInfos ApplicantInfos Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetApplicantInfos() []NormalVisaApplicantInfo {
	return r._applicantInfos
}

// SetDocumentInfos is DocumentInfos Setter
// 特殊必填，更多材料
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetDocumentInfos(_documentInfos []NormalVisaDocumentInfo) error {
	r._documentInfos = _documentInfos
	r.Set("document_infos", _documentInfos)
	return nil
}

// GetDocumentInfos DocumentInfos Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetDocumentInfos() []NormalVisaDocumentInfo {
	return r._documentInfos
}

// SetSubOrderId is SubOrderId Setter
// 必填，子订单id
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetSubOrderId(_subOrderId string) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetSubOrderId() string {
	return r._subOrderId
}

// SetPhotoType is PhotoType Setter
// 证件照文件类型
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetPhotoType(_photoType string) error {
	r._photoType = _photoType
	r.Set("photo_type", _photoType)
	return nil
}

// GetPhotoType PhotoType Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetPhotoType() string {
	return r._photoType
}

// SetPassportType is PassportType Setter
// 护照文件类型
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetPassportType(_passportType string) error {
	r._passportType = _passportType
	r.Set("passport_type", _passportType)
	return nil
}

// GetPassportType PassportType Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetPassportType() string {
	return r._passportType
}

// SetHotelBookingFormType is HotelBookingFormType Setter
// 酒店预订文件类型
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetHotelBookingFormType(_hotelBookingFormType string) error {
	r._hotelBookingFormType = _hotelBookingFormType
	r.Set("hotel_booking_form_type", _hotelBookingFormType)
	return nil
}

// GetHotelBookingFormType HotelBookingFormType Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetHotelBookingFormType() string {
	return r._hotelBookingFormType
}

// SetFlightBookingFormType is FlightBookingFormType Setter
// 机票预订文件类型
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetFlightBookingFormType(_flightBookingFormType string) error {
	r._flightBookingFormType = _flightBookingFormType
	r.Set("flight_booking_form_type", _flightBookingFormType)
	return nil
}

// GetFlightBookingFormType FlightBookingFormType Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetFlightBookingFormType() string {
	return r._flightBookingFormType
}

// SetApplicantOp is ApplicantOp Setter
// 特殊必填，签证申请人进度推进操作（目前只支持对单个申请人状态进行推进）。当operType为3时必填
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetApplicantOp(_applicantOp *NormalVisaApplicantOperation) error {
	r._applicantOp = _applicantOp
	r.Set("applicant_op", _applicantOp)
	return nil
}

// GetApplicantOp ApplicantOp Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetApplicantOp() *NormalVisaApplicantOperation {
	return r._applicantOp
}

// SetOperType is OperType Setter
// 必填，操作类型。1-上传新申请人基本信息（商家代填申请人），2-更新已有申请人基本信息，3-更新已有申请人的签证进度，4-商家代传申请人信息和材料(使馆直连订单)
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetOperType(_operType int64) error {
	r._operType = _operType
	r.Set("oper_type", _operType)
	return nil
}

// GetOperType OperType Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetOperType() int64 {
	return r._operType
}

// SetFileBytes is FileBytes Setter
// 特殊必填，pdf文件字节流。用于上传电子签pdf结果 或者 预约面试信pdf文件。
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetFileBytes(_fileBytes *model.File) error {
	r._fileBytes = _fileBytes
	r.Set("file_bytes", _fileBytes)
	return nil
}

// GetFileBytes FileBytes Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetFileBytes() *model.File {
	return r._fileBytes
}

// SetPhotoBytes is PhotoBytes Setter
// 特殊必填，文件字节流，用于上传证件照，必须和photoType同时传
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetPhotoBytes(_photoBytes *model.File) error {
	r._photoBytes = _photoBytes
	r.Set("photo_bytes", _photoBytes)
	return nil
}

// GetPhotoBytes PhotoBytes Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetPhotoBytes() *model.File {
	return r._photoBytes
}

// SetPassportBytes is PassportBytes Setter
// 特殊必填，文件字节流，用于上传护照，必须和passportType同时传
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetPassportBytes(_passportBytes *model.File) error {
	r._passportBytes = _passportBytes
	r.Set("passport_bytes", _passportBytes)
	return nil
}

// GetPassportBytes PassportBytes Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetPassportBytes() *model.File {
	return r._passportBytes
}

// SetHotelBookingFormBytes is HotelBookingFormBytes Setter
// 特殊必填，文件字节流，用于上传酒店预订，必须和hotelBookingFormType同时传
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetHotelBookingFormBytes(_hotelBookingFormBytes *model.File) error {
	r._hotelBookingFormBytes = _hotelBookingFormBytes
	r.Set("hotel_booking_form_bytes", _hotelBookingFormBytes)
	return nil
}

// GetHotelBookingFormBytes HotelBookingFormBytes Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetHotelBookingFormBytes() *model.File {
	return r._hotelBookingFormBytes
}

// SetFlightBookingFormBytes is FlightBookingFormBytes Setter
// 特殊必填，文件字节流，用于上传机票预订，必须和flightBookingFormType同时传
func (r *AlitriptravelvisaapplicantupdateAPIRequest) SetFlightBookingFormBytes(_flightBookingFormBytes *model.File) error {
	r._flightBookingFormBytes = _flightBookingFormBytes
	r.Set("flight_booking_form_bytes", _flightBookingFormBytes)
	return nil
}

// GetFlightBookingFormBytes FlightBookingFormBytes Getter
func (r AlitriptravelvisaapplicantupdateAPIRequest) GetFlightBookingFormBytes() *model.File {
	return r._flightBookingFormBytes
}
