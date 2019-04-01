package dynamic

import (
	"encoding/json"
	"reflect"
)

// DynamicFielder is the dynamic fielder interface
// Struct which implement this interface will have dynamic field support
type DynamicFielder interface {
	NewDynamicField(fieldName string) interface{}
}

var DynamicType = reflect.TypeOf(&Type{})

func IsDynamic(typ reflect.Type) bool {
	return typ == DynamicType
}

type Type struct {
	Value interface{}     `json:"-"`
	raw   json.RawMessage `json:"-"`
}

func New(v interface{}) *Type {
	return &Type{Value: v}
}

func GetValue(t *Type) interface{} {
	if t != nil {
		return t.Value
	}
	return nil
}

func (t *Type) UnmarshalJSON(data []byte) error {
	t.raw = data
	return nil
}

func (t *Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Value)
}

func (t *Type) GetRawMessage() json.RawMessage {
	return t.raw
}
