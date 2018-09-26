package funcmodels

import (
	"errors"
	"fmt"
)

// Function represents ... TODO
type Function interface {
	ID() bson.ObjectId
	Name() string
	TranspiledSource() string
	SourceMap() json.RawMessage
	Private() bool
	CanEvaluate() *bson.D
	Builder() *FunctionBuilder
}

// FunctionBuilder builds and validates functions
type FunctionBuilder struct {
	data function
}

type function struct {
	id               bson.ObjectId
	name             string
	transpiledSource string
	sourceMap        json.RawMessage
	private          bool
	canEvaluate      *bson.D
}

// NewFunction returns a new FunctionBuilder
func NewFunctionBuilder() *FunctionBuilder {
	return &FunctionBuilder{}
}

// ID returns the id of this function
func (fn *function) ID() bson.ObjectId {
	return fn.id
}

// Name returns the name of this function
func (fn *function) Name() string {
	return fn.name
}

// TranspiledSource returns the transpiledSource of this function
func (fn *function) TranspiledSource() string {
	return fn.transpiledSource
}

// SourceMap returns the sourceMap of this function
func (fn *function) SourceMap() json.RawMessage {
	return fn.sourceMap
}

// Private returns the private of this function
func (fn *function) Private() bool {
	return fn.private
}

// CanEvaluate returns the canEvaluate of this function
func (fn *function) CanEvaluate() *bson.D {
	return fn.canEvaluate
}

// Builder creates a shallow copy of the function and returns it as a FunctionBuilder
func (fn *function) Builder() *FunctionBuilder {
	builder := NewFunctionBuilder().
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

// WithID sets the ID for the FunctionBuilder
func (builder *FunctionBuilder) WithID(id bson.ObjectId) *FunctionBuilder {
	builder.data.id = id
	return builder
}

// WithName sets the Name for the FunctionBuilder
func (builder *FunctionBuilder) WithName(name string) *FunctionBuilder {
	builder.data.name = name
	return builder
}

// WithTranspiledSource sets the TranspiledSource for the FunctionBuilder
func (builder *FunctionBuilder) WithTranspiledSource(transpiledSource string) *FunctionBuilder {
	builder.data.transpiledSource = transpiledSource
	return builder
}

// WithSourceMap sets the SourceMap for the FunctionBuilder
func (builder *FunctionBuilder) WithSourceMap(sourceMap json.RawMessage) *FunctionBuilder {
	builder.data.sourceMap = sourceMap
	return builder
}

// WithPrivate sets the Private for the FunctionBuilder
func (builder *FunctionBuilder) WithPrivate(private bool) *FunctionBuilder {
	builder.data.private = private
	return builder
}

// WithCanEvaluate sets the CanEvaluate for the FunctionBuilder
func (builder *FunctionBuilder) WithCanEvaluate(canEvaluate *bson.D) *FunctionBuilder {
	builder.data.canEvaluate = canEvaluate
	return builder
}

// Build builds a new Function if it is validated
func (builder *FunctionBuilder) Build() (Function, error) {
	built := &function{
		id:               builder.data.id,
		name:             builder.data.name,
		transpiledSource: builder.data.transpiledSource,
		sourceMap:        builder.data.sourceMap,
		private:          builder.data.private,
		canEvaluate:      builder.data.canEvaluate,
	}

	// Do relevant checks here
	// if fn.data.id == "" {
	// 		return nil, errors.New("ERROR MESSAGE")
	// }

	return built, nil
}

// MustBuild calls Build() but panics if there is an error
func (fn *FunctionBuilder) MustBuild() Function {
	data, err := fn.Build()
	if err != nil {
		panic(fmt.Errorf("failed to build function: %v", err))
	}
	return data
}
