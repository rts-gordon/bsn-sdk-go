package node

import "github.com/rts-gordon/bsn-sdk-go/pkg/core/entity/base"

type SdkTransReqData struct {
	base.BaseReqModel
	Body SdkTransReqDataBody `json:"body"`
}

type SdkTransReqDataBody struct {
	TransData string `json:"transData"`
}

func (f *SdkTransReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.TransData

}
