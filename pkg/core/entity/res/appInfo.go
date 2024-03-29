package res

import (
	"strconv"

	"github.com/rts-gordon/bsn-sdk-go/pkg/core/entity/base"
)

type AppInfoResData struct {
	base.BaseResModel
	Body AppInfoResDataBody `json:"body"`
}

type AppInfoResDataBody struct {
	AppName       string `json:"appName"`
	AppType       string `json:"appType"`
	CaType        int    `json:"caType"`
	AlgorithmType int    `json:"algorithmType"`
	MspId         string `json:"mspId"`
	ChannelId     string `json:"channelId"`
}

func (f *AppInfoResData) GetEncryptionValue() string {

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.AppName
	fp = fp + f.Body.AppType
	fp = fp + strconv.Itoa(f.Body.CaType)
	fp = fp + strconv.Itoa(f.Body.AlgorithmType)
	fp = fp + f.Body.MspId
	fp = fp + f.Body.ChannelId
	return fp
}
