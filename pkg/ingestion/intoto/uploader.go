package intoto

import (
	"encoding/json"
	"fmt"

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
		prov        *ProvenanceV1
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

	// prov := &ProvenanceV1{
	// 	predicateType: slsa1.PredicateSLSAProvenance,
	// }

	// dec := json.NewDecoder(bytes.NewReader(data))
	// dec.DisallowUnknownFields()
	// err = dec.Decode(prov)
	// if err != nil {
	// 	return nil, err
	// }

	// intoto := &intoto.Envelope{}
	// err = intoto.SetPayload(data)
	// if err != nil {
	// 	return nil, err
	// }

	// TODO: Verify signature

	return &Envelope{Envelope: env, decodedData: data, prov: prov}, nil
}
