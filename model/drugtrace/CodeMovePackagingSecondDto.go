package drugtrace

// CodeMovePackagingSecondDto 结构体
type CodeMovePackagingSecondDto struct {
	// 内层小对象
	RecoverCodeMovePackagingResultDtos []RecoverCodeMovePackagingResultDto `json:"recover_code_move_packaging_result_dtos,omitempty" xml:"recover_code_move_packaging_result_dtos>recover_code_move_packaging_result_dto,omitempty"`
	// 替换父码
	SourceParentCode string `json:"source_parent_code,omitempty" xml:"source_parent_code,omitempty"`
	// 替换子码
	SourceChildCodes string `json:"source_child_codes,omitempty" xml:"source_child_codes,omitempty"`
	// 被替换父码
	TargetParentCode string `json:"target_parent_code,omitempty" xml:"target_parent_code,omitempty"`
	// 被替换子码
	TargetChildCodes string `json:"target_child_codes,omitempty" xml:"target_child_codes,omitempty"`
	// 码拼箱状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 码拼箱信息
	Info string `json:"info,omitempty" xml:"info,omitempty"`
}