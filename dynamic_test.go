package dynamic_test

import (
	"encoding/json"
	"testing"

	"github.com/k81/dynamic"
	"github.com/stretchr/testify/require"
)

func TestJSON(t *testing.T) {
	aObj := &jsonValue{
		Type:    "a",
		Content: dynamic.New(&aContent{16}),
	}
	aContent, err := json.Marshal(aObj)
	require.NoError(t, err)

	aObjParsed := &jsonValue{}
	err = dynamic.Parse(aContent, aObjParsed)
	require.NoError(t, err)
	require.Equal(t, aObj.Type, aObjParsed.Type)
	require.Equal(t, aObj.Content.Value, aObjParsed.Content.Value)

	bObj := &jsonValue{
		Type:    "b",
		Content: dynamic.New(&bContent{Values: []int{1, 2, 3}}),
	}
	bContent, err := json.Marshal(bObj)
	require.NoError(t, err)

	bObjParsed := &jsonValue{}
	err = dynamic.Parse(bContent, bObjParsed)
	require.NoError(t, err)
	require.Equal(t, bObj.Type, bObjParsed.Type)
	require.Equal(t, bObj.Content.Value, bObjParsed.Content.Value)
}

func TestNilMarshal(t *testing.T) {
	obj := &jsonValue{}
	input := []byte(`{"type": "c", "content": {"hello":"hello"}}`)
	err := dynamic.Parse(input, obj)
	require.NoError(t, err)
	output, err := json.Marshal(obj)
	require.NoError(t, err)
	require.NotContains(t, string(output), "content")
}
