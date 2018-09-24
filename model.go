package hostingmodels

import (
	"errors"
)

// AssetMetadata represents ... TODO
type AssetMetadata interface {
	ID() bson.ObjectID
	Name() string
	Age() int
	Sport() string
	Weird() []hosting.weird
	Builder() *AssetMetadataBuilder
}

// AssetMetadataBuilder provides a convenient interface for building assetMetadata
type AssetMetadataBuilder struct {
	data assetMetadata
}

type assetMetadata struct {
	id    bson.ObjectID
	name  string
	age   int
	sport string
	weird []hosting.weird
}

// NewAssetMetadata returns a new AssetMetadataBuilder
func NewAssetMetadataBuilder() *AssetMetadataBuilder {
	return &AssetMetadataBuilder{}
}

// ID returns the id of this basicAssetMetadata
func (amd *basicAssetMetadata) ID() bson.ObjectID {
	return amd.id
}

// Name returns the name of this basicAssetMetadata
func (amd *basicAssetMetadata) Name() string {
	return amd.name
}

// Age returns the age of this basicAssetMetadata
func (amd *basicAssetMetadata) Age() int {
	return amd.age
}

// Sport returns the sport of this basicAssetMetadata
func (amd *basicAssetMetadata) Sport() string {
	return amd.sport
}

// Weird returns the weird of this basicAssetMetadata
func (amd *basicAssetMetadata) Weird() []hosting.weird {
	return amd.weird
}

// Builder creates a shallow copy of the basicAssetMetadata and returns it as a builder
func (amd *basicAssetMetadata) Builder() *AssetMetadataBuilder {
	return NewAssetMetadataBuilder().
		WithID(amd.ID()).
		WithName(amd.Name()).
		WithAge(amd.Age()).
		WithSport(amd.Sport()).
		WithWeird(amd.Weird())
}

// WithID sets the ID for the AssetMetadataBuilder
func (amd *AssetMetadataBuilder) WithID() bson.ObjectID {
	amd.data.id = id
	return amd
}

// WithName sets the Name for the AssetMetadataBuilder
func (amd *AssetMetadataBuilder) WithName() string {
	amd.data.name = name
	return amd
}

// WithAge sets the Age for the AssetMetadataBuilder
func (amd *AssetMetadataBuilder) WithAge() int {
	amd.data.age = age
	return amd
}

// WithSport sets the Sport for the AssetMetadataBuilder
func (amd *AssetMetadataBuilder) WithSport() string {
	amd.data.sport = sport
	return amd
}

// WithWeird sets the Weird for the AssetMetadataBuilder
func (amd *AssetMetadataBuilder) WithWeird() []hosting.weird {
	amd.data.weird = weird
	return amd
}

// Build builds a new AssetMetadata if it is validated
func (amd *AssetMetadataBuilder) Build() (AssetMetadata, error) {
	// Do relevant checks here
	// if amd.data.id == "" {
	// 		return nil, errors.New("ERROR MESSAGE")
	// }
	return &amd.data, nil
}

// MustBuild calls Build() but panics if there is an error
func (amd *AssetMetadataBuilder) MustBuild() (AssetMetadata, error) {
	data, err := amd.Build()
	if err != nil {
		panic(err)
	}
	return data
}
