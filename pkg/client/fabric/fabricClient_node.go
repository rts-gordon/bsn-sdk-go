package fabric

import (
	nodereq "github.com/rts-gordon/bsn-sdk-go/pkg/core/entity/req/fabric/node"
	noderes "github.com/rts-gordon/bsn-sdk-go/pkg/core/entity/res/fabric/node"
	"github.com/wonderivan/logger"

	"encoding/json"

	"github.com/rts-gordon/bsn-sdk-go/pkg/core/trans"
	"github.com/rts-gordon/bsn-sdk-go/pkg/util/http"
)

func (c *FabricClient) SdkTran(body nodereq.TransReqDataBody) (*noderes.TranResData, error) {

	url := c.GetURL("/api/fabric/v1/node/trans")

	user, err := c.GetUser(body.UserName)
	if err != nil {
		return nil, err
	}

	request := body.GetTransRequest(c.Config.GetAppInfo().ChannelId)

	transData, _, err := trans.CreateRequest(user, request)

	if err != nil {
		return nil, err
	}

	data := &nodereq.SdkTransReqData{}
	data.Header = c.GetHeader()
	data.Body = nodereq.SdkTransReqDataBody{
		TransData: transData,
	}
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &noderes.TranResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil

}

func (c *FabricClient) ReqChainCode(body nodereq.TransReqDataBody) (*noderes.TranResData, error) {
	url := c.GetURL("/api/fabric/v1/node/reqChainCode")

	data := &nodereq.TransReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())
	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}
	res := &noderes.TranResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

func (c *FabricClient) GetTransInfo(body nodereq.TxTransReqDataBody) (*noderes.TransactionResData, error) {
	url := c.GetURL("/api/fabric/v1/node/getTransaction")

	data := &nodereq.TxTransReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())
	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}
	res := &noderes.TransactionResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

func (c *FabricClient) GetBlockInfo(body nodereq.BlockReqDataBody) (*noderes.BlockResData, error) {

	url := c.GetURL("/api/fabric/v1/node/getBlockInfo")

	data := &nodereq.BlockReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &noderes.BlockResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

func (c *FabricClient) GetLedgerInfo() (*noderes.LedgerResData, error) {

	url := c.GetURL("/api/fabric/v1/node/getLedgerInfo")

	data := &nodereq.LedgerReqData{}
	data.Header = c.GetHeader()
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &noderes.LedgerResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}
