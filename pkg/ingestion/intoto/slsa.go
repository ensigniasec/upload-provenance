package intoto

import (
	intoto "github.com/in-toto/in-toto-golang/in_toto"
	slsa1 "github.com/in-toto/in-toto-golang/in_toto/slsa_provenance/v1"
)

type ProvenanceV1 struct {
	intoto.StatementHeader
	Predicate     slsa1.ProvenancePredicate `json:"predicate"`
	predicateType string
}
