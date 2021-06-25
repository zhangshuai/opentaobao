package fenxiao

// FenxiaoGrade 
type FenxiaoGrade struct {

    // 主键
    GradeId   int64 `json:"grade_id,omitempty"`

    // 分销商等级名称
    Name   string `json:"name,omitempty"`

    // 记录创建时间
    Created   string `json:"created,omitempty"`

    // 记录最后修改时间
    Modified   string `json:"modified,omitempty"`

}