package node

import (
	"strconv"

	"github.com/rts-gordon/bsn-sdk-go/pkg/core/entity/base"
)

type LedgerResData struct {
	base.BaseResModel
	Body *LedgerResDataBody `json:"body"`
}

type LedgerResDataBody struct {
	BlockHash    string `json:"blockHash"`
	PreBlockHash string `json:"preBlockHash"`
	Height       uint64 `json:"height"`
}

func (f *LedgerResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.BlockHash
	fp = fp + f.Body.PreBlockHash
	fp = fp + strconv.FormatUint(f.Body.Height, 10)

	return fp
}
