package hostingmodels

import (
	"errors"
	"strings"
)

// AssetMetadata represents a single asset in the store
type AssetMetadata interface {
	AppID() string
	FilePath() string
	FileHash() string
	FileSize() int64
	Attrs() []AssetAttribute
	AttrsHash() string
	IsDirectory() bool
	Builder() *AssetMetadataBuilder
}

// AssetAttribute represents an attribute of some hosted asset, e.g. Content-Type
type AssetAttribute interface {
	Name() string
	Value() string
}

// AssetMetadataBuilder provides a convenient interface for building
// AssetMetadata structs
type AssetMetadataBuilder struct {
	data basicAssetMetadata
}

// NewAssetAttribute returns a new AssetAttribute
func NewAssetAttribute(name, value string) AssetAttribute {
	return assetAttribute{name, value}
}

type assetAttribute struct {
	name  string
	value string
}

type basicAssetMetadata struct {
	appID     string
	filePath  string
	fileHash  string
	fileSize  int64
	attrs     []assetAttribute
	attrsHash string
}

// Name returns the name of this attribute
func (attr assetAttribute) Name() string {
	return attr.name
}

// Value returns the value of this attribute
func (attr assetAttribute) Value() string {
	return attr.value
}

// AppID returns the AppID of this asset
func (amd *basicAssetMetadata) AppID() string {
	return amd.appID
}

// FilePath returns the FilePath for this asset
func (amd *basicAssetMetadata) FilePath() string {
	return amd.filePath
}

// FileHash returns the FileHash of this asset
func (amd *basicAssetMetadata) FileHash() string {
	return amd.fileHash
}

// FileSize returns the FileSize of this asset
func (amd *basicAssetMetadata) FileSize() int64 {
	return amd.fileSize
}

// Attrs returns the slice of AssetAttributes for this asset
func (amd *basicAssetMetadata) Attrs() []AssetAttribute {
	var newAttrs = make([]AssetAttribute, 0, len(amd.attrs))
	for _, newAttr := range amd.attrs {
		newAttrs = append(newAttrs, newAttr)
	}
	return newAttrs
}

// AttrsHash returns the hash of the attributes for this asset
func (amd *basicAssetMetadata) AttrsHash() string {
	return amd.attrsHash
}

// IsDirectory is true if the asset represents a directory
func (amd *basicAssetMetadata) IsDirectory() bool {
	return strings.HasSuffix(amd.filePath, "/")
}

// Builder creates a shallow copy of the asset metadata and returns it as a Builder
func (amd *basicAssetMetadata) Builder() *AssetMetadataBuilder {
	return NewAssetMetadataBuilder().
		WithAppID(amd.AppID()).
		WithFilePath(amd.FilePath()).
		WithFileHash(amd.FileHash()).
		WithFileSize(amd.FileSize()).
		WithAttrs(amd.Attrs()).
		WithAttrsHash(amd.AttrsHash())
}

// NewAssetMetadataBuilder returns a new, empty NewAssetMetadataBuilder
func NewAssetMetadataBuilder() *AssetMetadataBuilder {
	return &AssetMetadataBuilder{}
}

// WithAppID sets the AppID for the AssetMetadataBuilder
func (amd *AssetMetadataBuilder) WithAppID(id string) *AssetMetadataBuilder {
	amd.data.appID = id
	return amd
}

// WithFilePath sets the FilePath for the AssetMetadataBuilder
func (amd *AssetMetadataBuilder) WithFilePath(path string) *AssetMetadataBuilder {
	amd.data.filePath = path
	return amd
}

// WithFileHash sets the FileHash for the AssetMetadataBuilder
func (amd *AssetMetadataBuilder) WithFileHash(hash string) *AssetMetadataBuilder {
	amd.data.fileHash = hash
	return amd
}

// WithFileSize sets the FileSize for the AssetMetadataBuilder
func (amd *AssetMetadataBuilder) WithFileSize(size int64) *AssetMetadataBuilder {
	amd.data.fileSize = size
	return amd
}

// WithAttrs sets the Attrs for the AssetMetadataBuilder
func (amd *AssetMetadataBuilder) WithAttrs(attrs []AssetAttribute) *AssetMetadataBuilder {
	var newAttrs = make([]assetAttribute, 0, len(attrs))
	for _, newAttr := range attrs {
		newAttrs = append(newAttrs, assetAttribute{name: newAttr.Name(), value: newAttr.Value()})
	}
	amd.data.attrs = newAttrs
	return amd
}

// WithAttrsHash sets the AttrsHash for the AssetMetadataBuilder
func (amd *AssetMetadataBuilder) WithAttrsHash(attrsHash string) *AssetMetadataBuilder {
	amd.data.attrsHash = attrsHash
	return amd
}

// Build builds a new AssetMetadata if it is validated
func (amd *AssetMetadataBuilder) Build() (AssetMetadata, error) {

	// Do relevant checks here
	if amd.data.appID == "" {
		return nil, errors.New("an asset requires an appID associated with it")
	}

	if amd.data.filePath == "" {
		return nil, errors.New("an asset requires a non-empty file path")
	}

	return &amd.data, nil
}

// MustBuild calls Build() but panics if there is an error.
func (amd *AssetMetadataBuilder) MustBuild() AssetMetadata {
	md, err := amd.Build()
	if err != nil {
		panic(err)
	}
	return md
}
