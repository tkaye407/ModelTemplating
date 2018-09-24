package INSERTPACKAGENAMEHERE

import (
	"errors"

	"github.com/10gen/stitch/utils/xjson"

	"gopkg.in/mgo.v2/bson"
)

// AssetMetadata provides a convenient interface for building assetMetadata
type AssetMetadata struct {
	data assetMetadata
}

type assetMetadata struct {
	ID    bson.ObjectID   `bson:"_id"`
	Name  string          `bson:"name"`
	Age   int             `bson:"age"`
	Sport string          `bson:"sport"`
	Weird []hosting.weird `bson:"weird"`
}

// Newhostingmodels.AssetMetadata returns a new AssetMetadata
func NewAssetMetadata() *AssetMetadata {
	return &AssetMetadata{}
}

// ID returns the id of this AssetMetadata
func (amd *AssetMetadata) ID() bson.ObjectID {
	return amd.id
}

// Name returns the name of this AssetMetadata
func (amd *AssetMetadata) Name() string {
	return amd.name
}

// Age returns the age of this AssetMetadata
func (amd *AssetMetadata) Age() int {
	return amd.age
}

// Sport returns the sport of this AssetMetadata
func (amd *AssetMetadata) Sport() string {
	return amd.sport
}

// Weird returns the weird of this AssetMetadata
func (amd *AssetMetadata) Weird() []hosting.weird {
	return amd.weird
}

// Builder creates a shallow copy of the AssetMetadata and returns it as a builder
func (amd *AssetMetadata) Builder() *AssetMetadata {
	return NewAssetMetadata().
		WithID(amd.ID()).
		WithName(amd.Name()).
		WithAge(amd.Age()).
		WithSport(amd.Sport()).
		WithWeird(amd.Weird())
}

// MarshalBSON marshals the AssetMetadata to BSON
func (amd, AssetMetadata) MarshalBSON() ([]byte, error) {
	return bson.Marshal(amd)
}

// UnmarshalBSON unmarshals the AssetMetadata from BSON
func (amd *AssetMetadata) UnmarshalBSON(data []byte) error {
	return bson.Unmarshal(data, amd)
}

// GetBSON returns the inner data for BSON marshaling
func (amd AssetMetadata) GetBSON() (interface{}, error) {
	return amd.data, nil
}

// SetBSON unmarshals BSON onto the AssetMetadata
func (amd *AssetMetadata) SetBSON(raw bson.Raw) error {
	return raw.Unmarshal(&amd.data)
}
