package empty_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/googleapis/gnostic/OpenAPIv2"
	"github.com/googleapis/gnostic/compiler"
)

const docEmpty = `swagger: "2.0"
info:
  title: ""
  version: v9
  description: ""
paths: {}
`

func TestEmpty(t *testing.T) {
	t.Log("Original decode")
	doc, err := decodeDoc(t, docEmpty)
	require.NoError(t, err)
	encoded := encodeDoc(t, doc)
	t.Log("Decode decoded and encoded")
	_, err = decodeDoc(t, encoded)
	assert.NoError(t, err, "decode encoded failed :(")
	assert.Equal(t, docEmpty, encoded)
}

func encodeDoc(t *testing.T, doc *openapi_v2.Document) string {
	bytes, err := yaml.Marshal(doc.ToRawInfo())
	require.NoError(t, err)
	return string(bytes)
}

func decodeDoc(t *testing.T, data string) (*openapi_v2.Document, error) {
	var info yaml.MapSlice
	err := yaml.Unmarshal([]byte(data), &info)
	require.NoError(t, err)
	t.Log("After yaml parse: ", info)
	return openapi_v2.NewDocument(info, compiler.NewContextWithExtensions("$root", nil, nil))
}
