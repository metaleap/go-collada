package cdom

import (
	ugfx "github.com/metaleap/go-util/gfx"
)

//	Categorizes the kinds of filtering used in FxSamplerFiltering.
type FxFilterKind int

const (
	//	Bilinear filtering.
	FxFilterKindNearest FxFilterKind = 0x2600

	//	Trilinear filtering.
	FxFilterKindLinear FxFilterKind = 0x2601

	//	Compensates for distortion caused by the difference in angle between a polygon and the view plane.
	FxFilterKindAnisotropic FxFilterKind = 21

	//	No MIP-mapped minification.
	FxFilterKindMipNone FxFilterKind = 22
)

//	Categorizes the kind of an FxSampler.
type FxSamplerKind int

const (
	//	Declares a one-dimensional texture sampler.
	FxSamplerKind1D FxSamplerKind = 0x8B5D

	//	Declares a two-dimensional texture sampler.
	FxSamplerKind2D FxSamplerKind = 0x8B5E

	//	Declares a three-dimensional texture sampler.
	FxSamplerKind3D FxSamplerKind = 0x8B5F

	//	Declares a texture sampler for cube maps.
	FxSamplerKindCube FxSamplerKind = 0x8B60

	//	Declares a texture sampler for depth maps.
	FxSamplerKindDepth FxSamplerKind = 31

	//	Declares a rectangular texture sampler.
	FxSamplerKindRect FxSamplerKind = 32
)

var (
	//	Default texture minification, magnification and MIP-mapping.
	DefaultFxSamplerFiltering = &FxSamplerFiltering{
		FilterMag:     FxFilterKindLinear,
		FilterMin:     FxFilterKindLinear,
		FilterMip:     FxFilterKindLinear,
		MaxAnisotropy: 1,
	}
	//	Default texture repeating and clamping.
	DefaultFxSamplerWrapping = &ugfx.SamplerWrapping{
		BorderColor: ugfx.Rgba32{R: 0, G: 0, B: 0, A: 1},
		WrapS:       ugfx.WrapKindRepeat,
		WrapT:       ugfx.WrapKindRepeat,
		WrapP:       ugfx.WrapKindRepeat,
	}
)

//	Declares a texture sampler.
type FxSampler struct {
	//	Extras
	HasExtras

	//	Filtering, Wrapping
	FxSamplerStates

	//	If set, instantiates a default image from which the sampler is to consume.
	Image *FxImageInst

	//	Indicates the kind of this texture sampler.
	//	Must be one of the FxSamplerKind* enumerated constants.
	Kind FxSamplerKind
}

//	Constructor
func NewFxSampler() (me *FxSampler) {
	me = &FxSampler{}
	me.FxSamplerStates.Filtering = DefaultFxSamplerFiltering
	me.FxSamplerStates.Wrapping = DefaultFxSamplerWrapping
	return
}

//	Controls texture minification, magnification and MIP-mapping.
type FxSamplerFiltering struct {
	//	Magnification filter. Must be one of the FxFilterKind* enumerated constants.
	FilterMag FxFilterKind

	//	Minification filter. Must be one of the FxFilterKind* enumerated constants.
	FilterMin FxFilterKind

	//	Mip-mapping filter. Must be one of the FxFilterKind* enumerated constants.
	FilterMip FxFilterKind

	//	The number of samples that can be used durring anisotropic filtering.
	MaxAnisotropy uint32

	//	The maximum number of progressive levels that the sampler will evaluate.
	MipMaxLevel uint8

	//	The minimum progressive levels to begin to evaluate.
	MipMinLevel uint8

	//	Biases the gamma (level of detail parameter) used by the sampler to evaluate the MIPmap chain.
	MipBias float64
}

//	Instantiates an image targeted for samplers.
type FxSamplerImage struct {
	//	Sid, Name, Extras, DefRef
	FxImageInst
}

//	Allows users to modify an effect's sampler state from a material.
type FxSamplerStates struct {
	//	Extras
	HasExtras

	//	Controls texture minification, magnification and MIP-mapping.
	Filtering *FxSamplerFiltering

	//	Controls texture repeating and clamping.
	Wrapping *ugfx.SamplerWrapping
}

//	Constructor
func NewFxSamplerStates() (me *FxSamplerStates) {
	me = &FxSamplerStates{Filtering: DefaultFxSamplerFiltering, Wrapping: DefaultFxSamplerWrapping}
	return
}
