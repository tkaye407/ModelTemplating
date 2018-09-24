package funcmodels

import (
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// Function defines a function
type Function interface {
	ID() bson.ObjectId
	Name() string
	Source() string
	TranspiledSource() string
	SourceMap() json.RawMessage
	Private() bool
	CanEvaluate() *bson.D
	Builder() *FunctionBuilder
}

// FunctionBuilder builds and validates Functions
type FunctionBuilder struct {
	internal basicFunction
}

// NewFunctionBuilder returns a new, empty FunctionBuilder
func NewFunctionBuilder() *FunctionBuilder {
	return &FunctionBuilder{}
}

// WithID sets the ID for the function; optional
func (builder *FunctionBuilder) WithID(id bson.ObjectId) *FunctionBuilder {
	builder.internal.id = id
	return builder
}

// WithName sets the name for the function; required
func (builder *FunctionBuilder) WithName(name string) *FunctionBuilder {
	builder.internal.name = name
	return builder
}

// WithSource sets the internal source for the function; optional
func (builder *FunctionBuilder) WithSource(source string) *FunctionBuilder {
	builder.internal.source = source
	return builder
}

// WithTranspiledSource sets the internal source for the function; optional
func (builder *FunctionBuilder) WithTranspiledSource(transpiledSource string) *FunctionBuilder {
	builder.internal.transpiledSource = transpiledSource
	return builder
}

// WithSourceMap sets the internal source for the function; optional
func (builder *FunctionBuilder) WithSourceMap(sourceMap json.RawMessage) *FunctionBuilder {
	builder.internal.sourceMap = sourceMap
	return builder
}

// WithPrivate sets the internal source for the function; optional
func (builder *FunctionBuilder) WithPrivate(private bool) *FunctionBuilder {
	builder.internal.private = private
	return builder
}

// WithCanEvaluate sets the internal rule for the function; optional
func (builder *FunctionBuilder) WithCanEvaluate(canEvaluate bson.D) *FunctionBuilder {
	builder.internal.canEvaluate = &canEvaluate
	return builder
}

// Build builds a new Function if valid
func (builder *FunctionBuilder) Build() (Function, error) {
	// Make a shallow copy
	built := &basicFunction{
		id:               builder.internal.id,
		name:             builder.internal.name,
		source:           builder.internal.source,
		transpiledSource: builder.internal.transpiledSource,
		sourceMap:        builder.internal.sourceMap,
		private:          builder.internal.private,
		canEvaluate:      builder.internal.canEvaluate,
	}

	if !built.id.Valid() {
		built.id = bson.NewObjectId()
	}

	// for testing
	if built.sourceMap == nil {
		built.sourceMap = json.RawMessage{}
	}

	if built.name == "" {
		return nil, errors.New("a function requires a name")
	}

	return built, nil
}

// MustBuild builds and returns a Function if valid; panics otherwise
func (builder *FunctionBuilder) MustBuild() Function {
	function, err := builder.Build()
	if err != nil {
		panic(fmt.Errorf("failed to build function: %v", err))
	}
	return function
}


type basicFunction struct {
	id               bson.ObjectId
	name             string
	source           string
	transpiledSource string
	sourceMap        json.RawMessage
	private          bool
	canEvaluate      *bson.D
}

// ID returns the ID of this function
func (function *basicFunction) ID() bson.ObjectId {
	return function.id
}

// Name returns the name of this function
func (function *basicFunction) Name() string {
	return function.name
}

// Source returns the inner source of this function
func (function *basicFunction) Source() string {
	return function.source
}

// TranspiledSource returns the transpiled source of this function
func (function *basicFunction) TranspiledSource() string {
	return function.transpiledSource
}

// SourceMap returns the source map of this function
func (function *basicFunction) SourceMap() json.RawMessage {
	return function.sourceMap
}

// Private returns the private status of this function
func (function *basicFunction) Private() bool {
	return function.private
}

// CanEvaluate returns the CanEvaluate of this function
func (function *basicFunction) CanEvaluate() *bson.D {
	return function.canEvaluate
}

// Builder makes a shallow copy of this function and returns it
// as a builder
func (function *basicFunction) Builder() *FunctionBuilder {
	builder := NewFunctionBuilder().
		WithID(function.id).
		WithName(function.name).
		WithSource(function.source).
		WithTranspiledSource(function.transpiledSource).
		WithSourceMap(function.sourceMap).
		WithPrivate(function.private)
	if function.canEvaluate != nil {
		builder = builder.WithCanEvaluate(*function.canEvaluate)
	}
	return builder

}
