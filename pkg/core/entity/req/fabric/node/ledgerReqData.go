package node

import "github.com/chcp/bsn-sdk-go/pkg/core/entity/base"

type LedgerReqData struct {
	base.BaseReqModel
	Body LedgerReqDataBody `json:"body"`
}

type LedgerReqDataBody struct {
}

func (f *LedgerReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue()

}
