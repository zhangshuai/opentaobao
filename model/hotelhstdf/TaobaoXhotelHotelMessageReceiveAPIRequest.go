package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelhotelmessagereceiveAPIRequest 接收道消息接口 API请求
// taobao.xhotel.hotel.message.receive
//
// 接收道消息接口
type TaobaoxhotelhotelmessagereceiveAPIRequest struct {
	model.Params
	// 文件在OSS的存储路径，不包含文件名，只包含文件夹路径
	_ossPath string
	// 上传的文件个数
	_fileCount int64
	// 文件记录总数量
	_recordCount int64
	// 数据的类型，枚举值见文档
	_dataType int64
}

// NewTaobaoxhotelhotelmessagereceiveRequest 初始化TaobaoxhotelhotelmessagereceiveAPIRequest对象
func NewTaobaoxhotelhotelmessagereceiveRequest() *TaobaoxhotelhotelmessagereceiveAPIRequest {
	return &TaobaoxhotelhotelmessagereceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelhotelmessagereceiveAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.hotel.message.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelhotelmessagereceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelhotelmessagereceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOssPath is OssPath Setter
// 文件在OSS的存储路径，不包含文件名，只包含文件夹路径
func (r *TaobaoxhotelhotelmessagereceiveAPIRequest) SetOssPath(_ossPath string) error {
	r._ossPath = _ossPath
	r.Set("oss_path", _ossPath)
	return nil
}

// GetOssPath OssPath Getter
func (r TaobaoxhotelhotelmessagereceiveAPIRequest) GetOssPath() string {
	return r._ossPath
}

// SetFileCount is FileCount Setter
// 上传的文件个数
func (r *TaobaoxhotelhotelmessagereceiveAPIRequest) SetFileCount(_fileCount int64) error {
	r._fileCount = _fileCount
	r.Set("file_count", _fileCount)
	return nil
}

// GetFileCount FileCount Getter
func (r TaobaoxhotelhotelmessagereceiveAPIRequest) GetFileCount() int64 {
	return r._fileCount
}

// SetRecordCount is RecordCount Setter
// 文件记录总数量
func (r *TaobaoxhotelhotelmessagereceiveAPIRequest) SetRecordCount(_recordCount int64) error {
	r._recordCount = _recordCount
	r.Set("record_count", _recordCount)
	return nil
}

// GetRecordCount RecordCount Getter
func (r TaobaoxhotelhotelmessagereceiveAPIRequest) GetRecordCount() int64 {
	return r._recordCount
}

// SetDataType is DataType Setter
// 数据的类型，枚举值见文档
func (r *TaobaoxhotelhotelmessagereceiveAPIRequest) SetDataType(_dataType int64) error {
	r._dataType = _dataType
	r.Set("data_type", _dataType)
	return nil
}

// GetDataType DataType Getter
func (r TaobaoxhotelhotelmessagereceiveAPIRequest) GetDataType() int64 {
	return r._dataType
}