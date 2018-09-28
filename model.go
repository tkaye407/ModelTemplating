package valmodels

import (
	"errors"
	"fmt"
)

// Value represents ... TODO
type Value interface {
	ID() bson.ObjectId
	Name() string
	Value() xjson.Value
	Private() bool
	Builder() *ValueBuilder
}

// ValueBuilder builds and validates values
type ValueBuilder struct {
	data value
}

type value struct {
	id      bson.ObjectId
	name    string
	value   xjson.Value
	private bool
}

// NewValue returns a new ValueBuilder
func NewValueBuilder() *ValueBuilder {
	return &ValueBuilder{}
}

// ID returns the id of this value
func (val *value) ID() bson.ObjectId {
	return val.id
}

// Name returns the name of this value
func (val *value) Name() string {
	return val.name
}

// Value returns the value of this value
func (val *value) Value() xjson.Value {
	return val.value
}

// Private returns the private of this value
func (val *value) Private() bool {
	return val.private
}

// Builder creates a shallow copy of the value and returns it as a ValueBuilder
func (val *value) Builder() *ValueBuilder {
	builder := NewValueBuilder().
		WithID(val.ID()).
		WithName(val.Name()).
		WithValue(val.Value()).
		WithPrivate(val.Private())

	// perform any necessary checks
	// if ....

	return builder
}

// WithID sets the ID for the ValueBuilder
func (builder *ValueBuilder) WithID(id bson.ObjectId) *ValueBuilder {
	builder.data.id = id
	return builder
}

// WithName sets the Name for the ValueBuilder
func (builder *ValueBuilder) WithName(name string) *ValueBuilder {
	builder.data.name = name
	return builder
}

// WithValue sets the Value for the ValueBuilder
func (builder *ValueBuilder) WithValue(value xjson.Value) *ValueBuilder {
	builder.data.value = value
	return builder
}

// WithPrivate sets the Private for the ValueBuilder
func (builder *ValueBuilder) WithPrivate(private bool) *ValueBuilder {
	builder.data.private = private
	return builder
}

// Build builds a new Value if it is validated
func (builder *ValueBuilder) Build() (Value, error) {
	built := &value{
		id:      builder.data.id,
		name:    builder.data.name,
		value:   builder.data.value,
		private: builder.data.private,
	}

	// Do relevant checks here
	// if val.data.id == "" {
	// 		return nil, errors.New("ERROR MESSAGE")
	// }

	return built, nil
}

// MustBuild calls Build() but panics if there is an error
func (builder *ValueBuilder) MustBuild() Value {
	data, err := builder.Build()
	if err != nil {
		panic(fmt.Errorf("failed to build value: %v", err))
	}
	return data
}
