package viapi

// Element 
type Element struct {
    // 任务Id
    TaskId   string `json:"task_id,omitempty" xml:"task_id,omitempty"`
    // 单个文本的检测结果
    Results   []AliyunViapiImageauditScantextResult `json:"results,omitempty" xml:"results>aliyun_viapi_imageaudit_scantext_result,omitempty"`
}