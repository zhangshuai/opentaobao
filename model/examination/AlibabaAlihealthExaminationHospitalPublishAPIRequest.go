package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationHospitalPublishAPIRequest 体检机构对接_门店发布／更新 API请求
// alibaba.alihealth.examination.hospital.publish
//
// 第三方B端有新的门店发布，或者老的门店更新的时候，使用这个接口
type AlibabaAlihealthExaminationHospitalPublishAPIRequest struct {
	model.Params
	// 门店简介
	_detail string
	// 门店联系电话
	_tel string
	// 门店所属城市
	_cityName string
	// 门店城市code（国标）
	_cityCode string
	// 操作类型: publish=发布，update=更新
	_type string
	// 医院等级，三甲、
	_keyWord string
	// “须知”使用下面note_category字段
	_examNotice string
	// 门店位置经度高德 坐标系
	_pointX string
	// 门店位置纬度高德 坐标系
	_pointY string
	// 门店地址
	_address string
	// 工作时间
	_workTime string
	// 门店名称
	_hospitalName string
	// 门店code，机构保证唯一
	_hospitalCode string
	// 交通线路，通过\r\n 进行换行
	_routes string
	// http://images.aliyun.com/image?id=123
	_logo string
	// 是否支持在线报告。0:不支持;1:支持
	_onlineReport int64
	// 社会统一信用代码
	_socialCreditCode string
	// 线下报告获取说明（必填）
	_reportWay string
	// 线上体检报告几天出具（如果有电子报告必填）
	_reportWayOnline string
	// 环境图片(json字符串数组)，第一张是头图；（传图前先找运营同学要图片规范，别瞎传）
	_envImgsUrl string
	// 免费停车场,绿色VIP通道,免费早餐,3天出报告,1V1导检,接待引导,独家签约,专家会诊,当天出报告；多个逗号分隔
	_specialTagsCode string
	// 通知信息
	_notify string
	// 不同种类的预约须知；
	_noteCategory string
	// 经营模式 0自营模式、1平台模式
	_mode string
	// 门店与医院协议
	_agreement string
	// 营业执照
	_businessLicense string
	// 医疗经营许可
	_medicalLicense string
	// 类目:1=体检；2=核酸；3=上门；4=健康证；多个类目以逗号分割
	_category string
}

// NewAlibabaAlihealthExaminationHospitalPublishRequest 初始化AlibabaAlihealthExaminationHospitalPublishAPIRequest对象
func NewAlibabaAlihealthExaminationHospitalPublishRequest() *AlibabaAlihealthExaminationHospitalPublishAPIRequest {
	return &AlibabaAlihealthExaminationHospitalPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.hospital.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Detail Setter
// 门店简介
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetDetail(_detail string) error {
	r._detail = _detail
	r.Set("detail", _detail)
	return nil
}

// Get Detail Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetDetail() string {
	return r._detail
}

// Set is Tel Setter
// 门店联系电话
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetTel(_tel string) error {
	r._tel = _tel
	r.Set("tel", _tel)
	return nil
}

// Get Tel Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetTel() string {
	return r._tel
}

// Set is CityName Setter
// 门店所属城市
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetCityName(_cityName string) error {
	r._cityName = _cityName
	r.Set("city_name", _cityName)
	return nil
}

// Get CityName Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetCityName() string {
	return r._cityName
}

// Set is CityCode Setter
// 门店城市code（国标）
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// Get CityCode Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetCityCode() string {
	return r._cityCode
}

// Set is Type Setter
// 操作类型: publish=发布，update=更新
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetType() string {
	return r._type
}

// Set is KeyWord Setter
// 医院等级，三甲、
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetKeyWord(_keyWord string) error {
	r._keyWord = _keyWord
	r.Set("key_word", _keyWord)
	return nil
}

// Get KeyWord Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetKeyWord() string {
	return r._keyWord
}

// Set is ExamNotice Setter
// “须知”使用下面note_category字段
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetExamNotice(_examNotice string) error {
	r._examNotice = _examNotice
	r.Set("exam_notice", _examNotice)
	return nil
}

// Get ExamNotice Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetExamNotice() string {
	return r._examNotice
}

// Set is PointX Setter
// 门店位置经度高德 坐标系
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetPointX(_pointX string) error {
	r._pointX = _pointX
	r.Set("point_x", _pointX)
	return nil
}

// Get PointX Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetPointX() string {
	return r._pointX
}

// Set is PointY Setter
// 门店位置纬度高德 坐标系
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetPointY(_pointY string) error {
	r._pointY = _pointY
	r.Set("point_y", _pointY)
	return nil
}

// Get PointY Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetPointY() string {
	return r._pointY
}

// Set is Address Setter
// 门店地址
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// Get Address Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetAddress() string {
	return r._address
}

// Set is WorkTime Setter
// 工作时间
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetWorkTime(_workTime string) error {
	r._workTime = _workTime
	r.Set("work_time", _workTime)
	return nil
}

// Get WorkTime Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetWorkTime() string {
	return r._workTime
}

// Set is HospitalName Setter
// 门店名称
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetHospitalName(_hospitalName string) error {
	r._hospitalName = _hospitalName
	r.Set("hospital_name", _hospitalName)
	return nil
}

// Get HospitalName Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetHospitalName() string {
	return r._hospitalName
}

// Set is HospitalCode Setter
// 门店code，机构保证唯一
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetHospitalCode(_hospitalCode string) error {
	r._hospitalCode = _hospitalCode
	r.Set("hospital_code", _hospitalCode)
	return nil
}

// Get HospitalCode Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetHospitalCode() string {
	return r._hospitalCode
}

// Set is Routes Setter
// 交通线路，通过\r\n 进行换行
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetRoutes(_routes string) error {
	r._routes = _routes
	r.Set("routes", _routes)
	return nil
}

// Get Routes Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetRoutes() string {
	return r._routes
}

// Set is Logo Setter
// http://images.aliyun.com/image?id=123
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetLogo(_logo string) error {
	r._logo = _logo
	r.Set("logo", _logo)
	return nil
}

// Get Logo Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetLogo() string {
	return r._logo
}

// Set is OnlineReport Setter
// 是否支持在线报告。0:不支持;1:支持
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetOnlineReport(_onlineReport int64) error {
	r._onlineReport = _onlineReport
	r.Set("online_report", _onlineReport)
	return nil
}

// Get OnlineReport Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetOnlineReport() int64 {
	return r._onlineReport
}

// Set is SocialCreditCode Setter
// 社会统一信用代码
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetSocialCreditCode(_socialCreditCode string) error {
	r._socialCreditCode = _socialCreditCode
	r.Set("social_credit_code", _socialCreditCode)
	return nil
}

// Get SocialCreditCode Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetSocialCreditCode() string {
	return r._socialCreditCode
}

// Set is ReportWay Setter
// 线下报告获取说明（必填）
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetReportWay(_reportWay string) error {
	r._reportWay = _reportWay
	r.Set("report_way", _reportWay)
	return nil
}

// Get ReportWay Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetReportWay() string {
	return r._reportWay
}

// Set is ReportWayOnline Setter
// 线上体检报告几天出具（如果有电子报告必填）
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetReportWayOnline(_reportWayOnline string) error {
	r._reportWayOnline = _reportWayOnline
	r.Set("report_way_online", _reportWayOnline)
	return nil
}

// Get ReportWayOnline Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetReportWayOnline() string {
	return r._reportWayOnline
}

// Set is EnvImgsUrl Setter
// 环境图片(json字符串数组)，第一张是头图；（传图前先找运营同学要图片规范，别瞎传）
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetEnvImgsUrl(_envImgsUrl string) error {
	r._envImgsUrl = _envImgsUrl
	r.Set("env_imgs_url", _envImgsUrl)
	return nil
}

// Get EnvImgsUrl Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetEnvImgsUrl() string {
	return r._envImgsUrl
}

// Set is SpecialTagsCode Setter
// 免费停车场,绿色VIP通道,免费早餐,3天出报告,1V1导检,接待引导,独家签约,专家会诊,当天出报告；多个逗号分隔
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetSpecialTagsCode(_specialTagsCode string) error {
	r._specialTagsCode = _specialTagsCode
	r.Set("special_tags_code", _specialTagsCode)
	return nil
}

// Get SpecialTagsCode Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetSpecialTagsCode() string {
	return r._specialTagsCode
}

// Set is Notify Setter
// 通知信息
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetNotify(_notify string) error {
	r._notify = _notify
	r.Set("notify", _notify)
	return nil
}

// Get Notify Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetNotify() string {
	return r._notify
}

// Set is NoteCategory Setter
// 不同种类的预约须知；
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetNoteCategory(_noteCategory string) error {
	r._noteCategory = _noteCategory
	r.Set("note_category", _noteCategory)
	return nil
}

// Get NoteCategory Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetNoteCategory() string {
	return r._noteCategory
}

// Set is Mode Setter
// 经营模式 0自营模式、1平台模式
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetMode(_mode string) error {
	r._mode = _mode
	r.Set("mode", _mode)
	return nil
}

// Get Mode Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetMode() string {
	return r._mode
}

// Set is Agreement Setter
// 门店与医院协议
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetAgreement(_agreement string) error {
	r._agreement = _agreement
	r.Set("agreement", _agreement)
	return nil
}

// Get Agreement Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetAgreement() string {
	return r._agreement
}

// Set is BusinessLicense Setter
// 营业执照
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetBusinessLicense(_businessLicense string) error {
	r._businessLicense = _businessLicense
	r.Set("business_license", _businessLicense)
	return nil
}

// Get BusinessLicense Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetBusinessLicense() string {
	return r._businessLicense
}

// Set is MedicalLicense Setter
// 医疗经营许可
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetMedicalLicense(_medicalLicense string) error {
	r._medicalLicense = _medicalLicense
	r.Set("medical_license", _medicalLicense)
	return nil
}

// Get MedicalLicense Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetMedicalLicense() string {
	return r._medicalLicense
}

// Set is Category Setter
// 类目:1=体检；2=核酸；3=上门；4=健康证；多个类目以逗号分割
func (r *AlibabaAlihealthExaminationHospitalPublishAPIRequest) SetCategory(_category string) error {
	r._category = _category
	r.Set("category", _category)
	return nil
}

// Get Category Getter
func (r AlibabaAlihealthExaminationHospitalPublishAPIRequest) GetCategory() string {
	return r._category
}
