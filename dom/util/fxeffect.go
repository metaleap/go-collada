package cdomutil

import (
	cdom "github.com/metaleap/go-collada/dom"
)

//	Adds a new common profile to the specified effect definition (even if it already contains one).
func FxAddProfileCommon(def *cdom.FxEffectDef) (prof *cdom.FxProfile) {
	prof = cdom.NewProfile()
	prof.Common = &cdom.FxProfileCommon{}
	def.Profiles = append(def.Profiles, prof)
	return
}

//	Adds a new GLSL profile to the specified effect definition (even if it already contains one).
func FxAddProfileGlsl(def *cdom.FxEffectDef) (prof *cdom.FxProfile) {
	prof = cdom.NewProfile()
	prof.Glsl = cdom.NewFxProfileGlsl()
	def.Profiles = append(def.Profiles, prof)
	return
}

//	Ensures the specified effect definition contains a common profile and returns it.
func FxEnsureProfileCommon(def *cdom.FxEffectDef) (prof *cdom.FxProfile) {
	for _, prof = range def.Profiles {
		if prof.Common != nil {
			break
		}
		prof = nil
	}
	if prof == nil {
		prof = FxAddProfileCommon(def)
	}
	return
}

//	Ensures the specified effect definition contains a GLSL profile and returns it.
func FxEnsureProfileGlsl(def *cdom.FxEffectDef) (prof *cdom.FxProfile) {
	for _, prof = range def.Profiles {
		if prof.Glsl != nil {
			break
		}
		prof = nil
	}
	if prof == nil {
		prof = FxAddProfileGlsl(def)
	}
	return
}

//	Creates and returns a new cdom.FxColor initialized with the specified
//	Sid and red, green, blue and alpha channel values.
func NewFxColor(sid string, r, g, b, a float32) (me *cdom.FxColor) {
	me = &cdom.FxColor{}
	me.Sid, me.R, me.G, me.B, me.A = sid, r, g, b, a
	return me
}

//	Creates and returns a new cdom.FxColorOrTexture referring to either the
//	specified texture or the specified color. Only at least and at most one of
//	the arguments (texture, color, paramRef) must be non-nil/non-empty.
func NewFxColorOrTexture(texture *cdom.FxTexture, color *cdom.FxColor, paramRef string) (ct *cdom.FxColorOrTexture) {
	ct = &cdom.FxColorOrTexture{}
	ct.Texture = texture
	ct.Color = color
	ct.ParamRef.SetParamRef(paramRef)
	return
}

//	Creates a new cdom.FxEffectDef with the specified Id, and
//	optionally adds a common profile, a GLSL profile, or both to it.
func NewFxEffectDef(id string, ensureCommonProfile, ensureGlslProfile bool) (me *cdom.FxEffectDef) {
	me = cdom.FxEffectDefs.New(id)
	if ensureCommonProfile {
		FxEnsureProfileCommon(me)
	}
	if ensureGlslProfile {
		FxEnsureProfileGlsl(me)
	}
	return
}

//	Creates and returns a new cdom.FxTexture sampling from the specified 2D sampler.
func NewFxTexture(sampler2DParamRef, texCoord string) (me *cdom.FxTexture) {
	me = &cdom.FxTexture{}
	me.Sampler2D.SetParamRef(sampler2DParamRef)
	me.TexCoord = texCoord
	return me
}
