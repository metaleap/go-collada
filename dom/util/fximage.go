package cdomutil

import (
	cdom "github.com/metaleap/go-collada/dom"
)

//	Creates and returns a new cdom.FxImageDef with the specified Id,
//	to be initialized from the specified refUrl.
func NewFxImageDef(id, refUrl string) (me *cdom.FxImageDef) {
	me = cdom.FxImageDefs.New(id)
	me.InitFrom = cdom.NewFxImageInitFrom(refUrl)
	return
}
