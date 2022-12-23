package cert

import (
	"encoding/hex"
	"fmt"
	"github.com/cloudflare/cfssl/csr"
	"github.com/rts-gordon/bsn-sdk-go/pkg/util/keystore"
	"github.com/rts-gordon/bsn-sdk-go/third_party/github.com/hyperledger/fabric/bccsp"
	"testing"
)

func TestGetCSRPEM(t *testing.T) {

	name := "abc"
	cr := GetCertificateRequest(name)
	var ks bccsp.KeyStore

	fks, err := keystore.NewFileBasedKeyStore(nil, "./test/msp/keystore", false)

	ks = fks

	key, cspSigner, err := keystore.BCCSPKeyRequestGenerate(ks)

	if err != nil {
		fmt.Println(err)

	}
	csrPEM, err := csr.Generate(cspSigner, cr)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println("key:", hex.EncodeToString(key.SKI()))

	fmt.Println("csrPEM：", string(csrPEM))

}
