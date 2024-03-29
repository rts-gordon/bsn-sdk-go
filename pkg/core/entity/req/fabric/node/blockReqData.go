package node

import (
	"strconv"

	"github.com/rts-gordon/bsn-sdk-go/pkg/core/entity/base"
)

type BlockReqData struct {
	base.BaseReqModel
	Body BlockReqDataBody `json:"body"`
}

type BlockReqDataBody struct {
	BlockNumber uint64 `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`
	TxId        string `json:"txId"`
}

func (f *BlockReqData) GetEncryptionValue() string {

	fp := f.GetBaseEncryptionValue()
	fp = fp + strconv.FormatUint(f.Body.BlockNumber, 10)
	fp = fp + f.Body.BlockHash
	fp = fp + f.Body.TxId
	return fp

}
