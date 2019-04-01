package dynamic

import (
	"encoding/json"
	"reflect"
)

var DynamicType = reflect.TypeOf(Type{})

type Type struct {
	Value interface{}
	raw   json.RawMessage
}

func (t *Type) UnmarshalJSON(data []byte) error {
	t.raw = data
	return nil
}

func (t *Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Value)
}

func (t *Type) Unmarshal(v interface{}) error {
	return json.Unmarshal(t.raw, v)
}
