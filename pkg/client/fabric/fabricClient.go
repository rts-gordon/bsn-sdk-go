package fabric

import (
	"github.com/chcp/bsn-sdk-go/pkg/client"
	"github.com/chcp/bsn-sdk-go/pkg/core/config"
	"github.com/chcp/bsn-sdk-go/pkg/core/entity/msp"
	"github.com/chcp/bsn-sdk-go/pkg/util/keystore"
	"github.com/chcp/bsn-sdk-go/pkg/util/userstore"
	"github.com/wonderivan/logger"
)

//initialize client requested by fabric 
func InitFabricClient(config *config.Config) (*FabricClient, error) {

	//initialize configuration information 
	if err := config.Init(); err != nil {
		logger.Error("Configuration initialization failed")
		return nil, err
	}
	//generate a private key handler
	ks, err := keystore.NewFileBasedKeyStore(nil, config.GetKSPath(), false)

	if err != nil {
		logger.Error("keystore initialization failed")
		return nil, err
	}
	//generate a cert handler 
	us := userstore.NewUserStore(config.GetUSPath())

	client := client.Client{
		Ks:     ks,
		Us:     us,
		Config: config,
		Users:  make(map[string]*msp.UserData),
	}

	fabricClient := &FabricClient{client}
	//configure the algorithm type of user signature and generate a unified signature verification handler
	err = fabricClient.SetAlgorithm(config.GetAppInfo().AlgorithmType, config.GetAppCert().AppPublicCert, config.GetAppCert().UserAppPrivateCert)

	if err != nil {
		logger.Error("signHandle initialization failed")
		return nil, err
	}
	//load the client's info generated locally 
	fabricClient.LoadUser()

	return fabricClient, nil
}

type FabricClient struct {
	client.Client
}
