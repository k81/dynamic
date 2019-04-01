package dynamic_test

import (
	"encoding/json"
	"fmt"

	"github.com/k81/dynamic"
)

type aContent struct {
	Value int `json:"value"`
}

type bContent struct {
	Values []int `json:"items"`
}

type jsonValue struct {
	Type    string       `json:"type"`
	Content dynamic.Type `json:"content"`
}

func (jc *jsonValue) NewDynamicField(fieldName string) interface{} {
	switch jc.Type {
	case "a":
		return &aContent{}
	case "b":
		return &bContent{}
	}
	return nil
}

func ExampleMarshalA() {
	obj := &jsonValue{
		Type:    "a",
		Content: dynamic.Type{Value: &aContent{16}},
	}
	data, _ := json.Marshal(obj)
	fmt.Println(string(data))

	obj = &jsonValue{
		Type:    "b",
		Content: dynamic.Type{Value: &bContent{Values: []int{1, 2, 3}}},
	}
	data, _ = json.Marshal(obj)
	fmt.Println(string(data))
	// Output:
	// {"type":"a","content":{"value":16}}
	// {"type":"b","content":{"items":[1,2,3]}}
}

func ExampleUnmarshal() {
	input := []byte(`{"type":"b","content":{"items":[1,2,3]}}`)
	obj := &jsonValue{}
	_ = dynamic.Parse(input, obj)
	content, ok := obj.Content.Value.(*bContent)
	fmt.Println(obj.Type)
	fmt.Println(ok)
	fmt.Println(content.Values)
	// Output:
	// b
	// true
	// [1 2 3]
}
