package keystore

import (
	"fmt"
	"github.com/rts-gordon/bsn-sdk-go/pkg/core/entity/msp"
	"github.com/rts-gordon/bsn-sdk-go/pkg/util/userstore"
	"testing"
)

func TestLoadKey(t *testing.T) {

	ks, _ := NewFileBasedKeyStore(nil, "F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/rts-gordon/bsn-sdk-go/test/msp/keystore", false)
	us := userstore.NewUserStore("F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/rts-gordon/bsn-sdk-go/test/msp")

	user := &msp.UserData{
		UserName: "sdktest",
		AppCode:  "app0006202004071529586812466",
	}

	us.Load(user)

	LoadKey(user, ks)

	fmt.Println(string(user.EnrollmentCertificate))
	fmt.Println(user.PrivateKey)
}
