package dynamic

import (
	"encoding/json"
	"reflect"
)

var DynamicType = reflect.TypeOf(&Type{})

type Type struct {
	Value interface{}     `json:"-"`
	raw   json.RawMessage `json:"-"`
}

func New(v interface{}) *Type {
	return &Type{Value: v}
}

func (t *Type) UnmarshalJSON(data []byte) error {
	t.raw = data
	return nil
}

func (t *Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Value)
}
