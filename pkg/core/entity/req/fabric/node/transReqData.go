package node

import (
	"github.com/rts-gordon/bsn-sdk-go/pkg/core/entity/base"
	"github.com/rts-gordon/bsn-sdk-go/pkg/core/trans"
)

type TransReqData struct {
	base.BaseReqModel
	Body TransReqDataBody `json:"body"`
}

type TransReqDataBody struct {
	UserName string `json:"userName"`
	Nonce    string `json:"nonce"`

	ChainCode    string            `json:"chainCode"`
	FuncName     string            `json:"funcName"`
	Args         []string          `json:"args"`
	TransientMap map[string]string `json:"transientData"`
}

func (f *TransReqData) GetEncryptionValue() string {

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.UserName
	fp = fp + f.Body.Nonce
	fp = fp + f.Body.ChainCode
	fp = fp + f.Body.FuncName

	for _, a := range f.Body.Args {
		fp = fp + a
	}

	for k, v := range f.Body.TransientMap {
		fp = fp + k + v
	}

	return fp

}

func (t *TransReqDataBody) GetTransRequest(channelId string) *trans.TransRequest {
	request := &trans.TransRequest{
		ChannelId:   channelId,
		ChaincodeId: t.ChainCode,
		Fcn:         t.FuncName,
	}

	for _, a := range t.Args {
		request.Args = append(request.Args, []byte(a))
	}
	request.TransientMap = make(map[string][]byte)
	for k, v := range t.TransientMap {
		request.TransientMap[k] = []byte(v)
	}

	return request
}
