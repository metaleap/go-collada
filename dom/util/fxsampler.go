package cdomutil

import (
	cdom "github.com/metaleap/go-collada/dom"
)

//	Creates and returns a new FxSampler of cdom.FxSamplerKind2D, sampling the specified
//	image with the specified filtering (or cdom.DefaultFxSamplerFiltering if nil) and the
//	specified wrapping (or cdom.DefaultFxSamplerWrapping if nil).
func NewFxSampler2D(image *cdom.FxImageInst, filtering *cdom.FxSamplerFiltering, wrapping *cdom.FxSamplerWrapping) (me *cdom.FxSampler) {
	me = cdom.NewFxSampler()
	if filtering != nil {
		me.Filtering = filtering
	}
	if wrapping != nil {
		me.Wrapping = wrapping
	}
	me.Kind = cdom.FxSamplerKind2D
	me.Image = image
	return
}
