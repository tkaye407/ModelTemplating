package appstoremodels

import (
	"encoding/json"

	"github.com/10gen/stitch/function/models"
	"github.com/10gen/stitch/utils/xjson"

	"gopkg.in/mgo.v2/bson"
)

// Function is the store's representation of a funcmodels.Function
type Function struct {
	data functionData
}

type functionData struct {
	ID               bson.ObjectId   `bson:"_id"`
	Name             string          `bson:"name"`
	Source           string          `bson:"source"`
	TranspiledSource string          `bson:"transpiledSource"`
	SourceMap        json.RawMessage `bson:"sourceMap"`
	Private          bool            `bson:"private"`
	CanEvaluate *xjson.MarshalD `bson:"canEvaluate,omitempty"`
}

// ID returns the ID of this function
func (function *Function) ID() bson.ObjectId {
	return function.data.ID
}

// Name returns the name of this function
func (function *Function) Name() string {
	return function.data.Name
}

// Source returns the source of this function
func (function *Function) Source() string {
	return function.data.Source
}

// TranspiledSource returns the transpiled source of this function
func (function *Function) TranspiledSource() string {
	return function.data.TranspiledSource
}

// SourceMap returns the source map of this function
func (function *Function) SourceMap() json.RawMessage {
	return function.data.SourceMap
}

// Private returns the private status of this function
func (function *Function) Private() bool {
	return function.data.Private
}

// CanEvaluate returns an expression that determines whether or not
// this function can be executed
func (function *Function) CanEvaluate() *bson.D {
	if function.data.CanEvaluate == nil {
		return nil
	}
	rule := bson.D(*function.data.CanEvaluate)
	return &rule
}

// Builder makes a shallow copy of this function and returns it
// as a builder
func (function *Function) Builder() *funcmodels.FunctionBuilder {
	builder := funcmodels.NewFunctionBuilder().
		WithID(function.data.ID).
		WithName(function.data.Name).
		WithSource(function.data.Source).
		WithTranspiledSource(function.data.TranspiledSource).
		WithSourceMap(function.data.SourceMap).
		WithPrivate(function.data.Private)
	if function.data.CanEvaluate != nil {
		builder = builder.WithCanEvaluate(bson.D(*function.data.CanEvaluate))
	}
	return builder
}

// MarshalBSON marshals the function to BSON
func (function Function) MarshalBSON() ([]byte, error) {
	return bson.Marshal(function)
}

// UnmarshalBSON unmarshals the function from BSON
func (function *Function) UnmarshalBSON(data []byte) error {
	return bson.Unmarshal(data, function)
}

// GetBSON returns the inner data for BSON marshaling
func (function Function) GetBSON() (interface{}, error) {
	return function.data, nil
}

// SetBSON unmarshals BSON onto the Function
func (function *Function) SetBSON(raw bson.Raw) error {
	return raw.Unmarshal(&function.data)
}

// ToFunction converts a funcmodels.Function to a store Function
func ToFunction(function funcmodels.Function) *Function {
	data := functionData{
		ID:               function.ID(),
		Name:             function.Name(),
		Source:           function.Source(),
		TranspiledSource: function.TranspiledSource(),
		SourceMap:        function.SourceMap(),
		Private:          function.Private(),
	}
	if canEvaluate := function.CanEvaluate(); canEvaluate != nil {
		rule := xjson.MarshalD(*canEvaluate)
		data.CanEvaluate = &rule
	}
	return &Function{
		data: data,
	}
}
