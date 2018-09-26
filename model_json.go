package apimodelsv3

import (
	"encoding/json"
	"errors"

	"github.com/10gen/stitch/utils/xjson"

	"gopkg.in/mgo.v2/bson"
)

// Function is the representation of a funcmodels.Function interface
type Function struct {
	data function
}

type function struct {
	ID               bson.ObjectId   `json:"_id"`
	Name             string          `json:"name"`
	TranspiledSource string          `json:"transpiledSource"`
	SourceMap        json.RawMessage `json:"sourceMap"`
	Private          bool            `json:"private"`
	CanEvaluate      *bson.D         `json:"canEvaluate"`
}

// Newfuncmodels.Function returns a new Function
func NewFunction() *Function {
	return &Function{}
}

// ID returns the id of this Function
func (fn *Function) ID() bson.ObjectId {
	return fn.data.id
}

// Name returns the name of this Function
func (fn *Function) Name() string {
	return fn.data.name
}

// TranspiledSource returns the transpiledSource of this Function
func (fn *Function) TranspiledSource() string {
	return fn.data.transpiledSource
}

// SourceMap returns the sourceMap of this Function
func (fn *Function) SourceMap() json.RawMessage {
	return fn.data.sourceMap
}

// Private returns the private of this Function
func (fn *Function) Private() bool {
	return fn.data.private
}

// CanEvaluate returns the canEvaluate of this Function
func (fn *Function) CanEvaluate() *bson.D {
	return fn.data.canEvaluate
}

// Builder creates a shallow copy of the Function and returns it as a funcmodels.FunctionBuilder
func (fn *Function) Builder() *funcmodels.FunctionBuilder {
	builder := funcmodels.NewFunctionBuilder().
		WithID(fn.ID()).
		WithName(fn.Name()).
		WithTranspiledSource(fn.TranspiledSource()).
		WithSourceMap(fn.SourceMap()).
		WithPrivate(fn.Private()).
		WithCanEvaluate(fn.CanEvaluate())

	// perform any necessary checks
	// if ....

	return builder
}

// MarshalBSON marshals the Function to BSON
func (fn Function) MarshalBSON() ([]byte, error) {
	return bson.Marshal(fn)
}

// UnmarshalBSON unmarshals the Function from BSON
func (fn *Function) UnmarshalBSON(data []byte) error {
	return bson.Unmarshal(data, fn)
}

// GetBSON returns the inner data for BSON marshaling
func (fn Function) GetBSON() (interface{}, error) {
	return fn.data, nil
}

// SetBSON unmarshals BSON onto the Function
func (fn *Function) SetBSON(raw bson.Raw) error {
	return raw.Unmarshal(&fn.data)
}

// ToFunction converts a funcmodels.Function to a Function
func ToFunction(fn funcmodels.Function) *Function {
	data := function{
		ID:               fn.ID(),
		Name:             fn.Name(),
		TranspiledSource: fn.TranspiledSource(),
		SourceMap:        fn.SourceMap(),
		Private:          fn.Private(),
		CanEvaluate:      fn.CanEvaluate(),
	}

	// Perform some checks
	// if data.... ==

	return &Function{
		data: data,
	}
}
