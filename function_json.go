package apimodelsv3

import (
	"encoding/json"

	"github.com/10gen/stitch/function/models"
	"github.com/10gen/stitch/utils/xjson"

	"gopkg.in/mgo.v2/bson"
)

// Function is the API's representation of a funcmodels.Function
type Function struct {
	data functionData
}

type functionData struct {
	ID               bson.ObjectId   `json:"_id"`
	Name             string          `json:"name"`
	Source           string          `json:"source"`
	TranspiledSource string          `json:"transpiled_source,omitempty"`
	SourceMap        json.RawMessage `json:"source_map,omitempty"`
	CanEvaluate      *xjson.MarshalD `json:"can_evaluate,omitempty"`
	Private          bool            `json:"private"`
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

// CanEvaluate returns an expression that determines whether or not
// this function can be executed
func (function *Function) CanEvaluate() *bson.D {
	if function.data.CanEvaluate == nil {
		return nil
	}
	rule := bson.D(*function.data.CanEvaluate)
	return &rule
}

// Private returns whether or not this function is permitted to
// be called by clients
func (function *Function) Private() bool {
	return function.data.Private
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

// MarshalJSON marshals the inner data out to JSON
func (function Function) MarshalJSON() ([]byte, error) {
	return json.Marshal(function.data)
}

// UnmarshalJSON unmarshals JSON into the Function
func (function *Function) UnmarshalJSON(data []byte) error {
	// for testing
	if function.data.SourceMap == nil {
		function.data.SourceMap = json.RawMessage{}
	}

	return json.Unmarshal(data, &function.data)
}
