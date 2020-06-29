package event

import "github.com/chcp/bsn-sdk-go/pkg/core/entity/base"

type RemoveReqData struct {
	base.BaseReqModel
	Body RemoveReqDataBody `json:"body"`

}

type RemoveReqDataBody struct {
	EventId string `json:"eventId"`
}

func (f *RemoveReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.EventId

}
