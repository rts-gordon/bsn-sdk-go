package event

import "github.com/chcp/bsn-sdk-go/pkg/core/entity/base"

type QueryReqData struct {
	base.BaseReqModel
	Body interface{} `json:"body"`

}

func (f *QueryReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue()

}
