package apimodelsv3

import (
	"encoding/json"
	"errors"

	"github.com/10gen/stitch/utils/xjson"

	"gopkg.in/mgo.v2/bson"
)

// Value is the representation of a valmodels.Value interface
type Value struct {
	data value
}

type value struct {
	ID      bson.ObjectId `json:"_id"`
	Name    string        `json:"name"`
	Value   xjson.Value   `json:"value"`
	Private bool          `json:"private"`
}

// Newvalmodels.Value returns a new Value
func NewValue() *Value {
	return &Value{}
}

// ID returns the id of this Value
func (val *Value) ID() bson.ObjectId {
	return val.data.ID
}

// Name returns the name of this Value
func (val *Value) Name() string {
	return val.data.Name
}

// Value returns the value of this Value
func (val *Value) Value() xjson.Value {
	return val.data.Value
}

// Private returns the private of this Value
func (val *Value) Private() bool {
	return val.data.Private
}

// Builder creates a shallow copy of the Value and returns it as a valmodels.ValueBuilder
func (val *Value) Builder() *valmodels.ValueBuilder {
	builder := valmodels.NewValueBuilder().
		WithID(val.ID()).
		WithName(val.Name()).
		WithValue(val.Value()).
		WithPrivate(val.Private())

	// perform any necessary checks
	// if ....

	return builder
}

// JSON marshals the Value to JSON
func (val Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(val)
}

// UnmarshalJSON unmarshals the Value from JSON
func (val *Value) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, val)
}

// GetJSON returns the inner data for JSON marshaling
func (val Value) GetJSON() (interface{}, error) {
	return val.data, nil
}

// SetJSON unmarshals JSON onto the Value
func (val *Value) SetJSON(raw json.Raw) error {
	return raw.Unmarshal(&val.data)
}

// ToValue converts a valmodels.Value to a Value
func ToValue(val valmodels.Value) *Value {
	data := value{
		ID:      val.ID(),
		Name:    val.Name(),
		Value:   val.Value(),
		Private: val.Private(),
	}

	// Perform some checks
	// if data.... ==

	return &Value{
		data: data,
	}
}
