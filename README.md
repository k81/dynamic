Parse json with dynamic field

```go
package dynamicjson_test

import (
	"encoding/json"
	"fmt"

	"github.com/k81/dynamicjson"
)

type aContent struct {
	Value int `json:"value"`
}

type bContent struct {
	Values []int `json:"items"`
}

type jsonValue struct {
	dynamicjson.DynamicJSON
}

func (jc *jsonValue) NewDynamicContent(typ string) interface{} {
	switch typ {
	case "a":
		return &aContent{}
	case "b":
		return &bContent{}
	}
	return nil
}

func ExampleMarshalA() {
	obj := &jsonValue{}
	obj.SetType("a")
	obj.SetContent(&aContent{16})
	data, _ := json.Marshal(obj)
	fmt.Println(string(data))

	obj = &jsonValue{}
	obj.SetType("b")
	obj.SetContent(&bContent{Values: []int{1, 2, 3}})
	data, _ = json.Marshal(obj)
	fmt.Println(string(data))
	// Output:
	// {"type":"a","content":{"value":16}}
	// {"type":"b","content":{"items":[1,2,3]}}
}

func ExampleUnmarshal() {
	input := []byte(`{"type":"b","content":{"items":[1,2,3]}}`)
	obj := &jsonValue{}
	_ = dynamicjson.Parse(input, obj)
	content, ok := obj.GetContent().(*bContent)
	fmt.Println(obj.GetType())
	fmt.Println(ok)
	fmt.Println(content.Values)
	// Output:
	// b
	// true
	// [1 2 3]
}
```
