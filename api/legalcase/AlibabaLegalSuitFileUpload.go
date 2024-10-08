package legalcase

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalSuitFileUpload 诉讼文件上传接口
// alibaba.legal.suit.file.upload
//
// 上传文件接口
func AlibabaLegalSuitFileUpload(ctx context.Context, clt *core.SDKClient, req *legalcase.AlibabaLegalSuitFileUploadAPIRequest, resp *legalcase.AlibabaLegalSuitFileUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
