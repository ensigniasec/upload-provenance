package intoto

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromBytes(t *testing.T) {
	a := assert.New(t)

	content, err := os.ReadFile("testdata/ensignia-action-linux-amd64.intoto.jsonl")
	a.NoError(err)

	e, err := NewFromBytes(content)
	a.NoError(err)
	a.Equal("application/vnd.in-toto+json", e.PayloadType)

	t.Log(string(e.decodedData))
}
