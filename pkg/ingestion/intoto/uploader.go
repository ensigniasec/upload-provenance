package intoto

import (
	"bytes"
	"encoding/json"
	"fmt"

	intoto "github.com/in-toto/in-toto-golang/in_toto"
	dsselib "github.com/secure-systems-lab/go-securesystemslib/dsse"
)

var ErrInvalidPayloadType = fmt.Errorf("invalid payload type")

const (
	InToToStatementType = "application/vnd.in-toto+json"
)

type (
	Envelope struct {
		*dsselib.Envelope
		decodedData []byte
		Statement   *intoto.ProvenanceStatementSLSA02
	}
)

func envelopeFromBytes(payload []byte) (env *dsselib.Envelope, err error) {
	env = &dsselib.Envelope{}
	err = json.Unmarshal(payload, env)
	return
}

func NewFromBytes(content []byte) (*Envelope, error) {
	env, err := envelopeFromBytes(content)
	if err != nil {
		return nil, err
	}

	if env.PayloadType != InToToStatementType {
		return nil, ErrInvalidPayloadType
	}

	data, err := env.DecodeB64Payload()
	if err != nil {
		return nil, err
	}

	prov := &intoto.ProvenanceStatementSLSA02{}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	err = dec.Decode(prov)
	if err != nil {
		return nil, err
	}

	// TODO: Verify signature

	return &Envelope{Envelope: env, decodedData: data, Statement: prov}, nil
}
