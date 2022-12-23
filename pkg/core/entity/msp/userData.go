package msp

import (
	"crypto/ecdsa"

	"github.com/golang/protobuf/proto"
	"github.com/rts-gordon/bsn-sdk-go/pkg/common/errors"
	pb_msp "github.com/rts-gordon/bsn-sdk-go/third_party/github.com/hyperledger/fabric/protos/msp"
)

type UserData struct {
	UserName              string
	AppCode               string
	MspId                 string
	EnrollmentCertificate []byte

	PrivateKey *ecdsa.PrivateKey
}

func (u *UserData) Serialize() ([]byte, error) {
	serializedIdentity := &pb_msp.SerializedIdentity{
		Mspid:   u.MspId,
		IdBytes: u.EnrollmentCertificate,
	}
	identity, err := proto.Marshal(serializedIdentity)
	if err != nil {
		return nil, errors.New("marshal serializedIdentity failed")
	}
	return identity, nil
}
