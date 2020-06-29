/**
 * @Author: Gao Chenxi
 * @Description:
 * @Date: 2020/4/1 4:35 PM
 * @File: main
 */
package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/chcp/bsn-sdk-go/pkg/client/fabric"
	"github.com/chcp/bsn-sdk-go/pkg/core/config"
	req "github.com/chcp/bsn-sdk-go/pkg/core/entity/req/fabric/node"
	"github.com/chcp/bsn-sdk-go/pkg/util/crypto"
)

func main() {
	fmt.Println("github.com/chcp/bsn-sdk-go")

	config, err := config.NewMockConfig()
	if err != nil {
		log.Fatal(err)
	}

	fabricClient, err := fabric.InitFabricClient(config)
	if err != nil {
		log.Fatal(err)
	}

	var args []string
	args = append(args, "test20200406")

	nonce, _ := crypto.GetRandomNonce()
	body := req.TransReqDataBody{
		UserName:     "user001", //不需要设置username
		Nonce:        base64.StdEncoding.EncodeToString(nonce),
		ChainCode:    "cc_base",
		FuncName:     "get",
		Args:         args,
		TransientMap: make(map[string]string),
	}

	res, _ := fabricClient.ReqChainCode(body)

	fmt.Println(res)
}
