# cdomutil
--
    import "github.com/metaleap/go-collada/dom/util"

Provides utility functions encapsulating (otherwise potentially verbose)
recurring tasks for working with the go-collada/dom package.

Specifically, provides a wide range of useful constructor functions for
a variety of go-collada/dom package resource definitions and instances.

## Usage

#### func  FxAddProfileCommon

```go
func FxAddProfileCommon(def *cdom.FxEffectDef) (prof *cdom.FxProfile)
```
Adds a new common profile to the specified effect definition (even if it already
contains one).

#### func  FxAddProfileGlsl

```go
func FxAddProfileGlsl(def *cdom.FxEffectDef) (prof *cdom.FxProfile)
```
Adds a new GLSL profile to the specified effect definition (even if it already
contains one).

#### func  FxEnsureProfileCommon

```go
func FxEnsureProfileCommon(def *cdom.FxEffectDef) (prof *cdom.FxProfile)
```
Ensures the specified effect definition contains a common profile and returns
it.

#### func  FxEnsureProfileGlsl

```go
func FxEnsureProfileGlsl(def *cdom.FxEffectDef) (prof *cdom.FxProfile)
```
Ensures the specified effect definition contains a GLSL profile and returns it.

#### func  NewFxColor

```go
func NewFxColor(sid string, r, g, b, a float32) (me *cdom.FxColor)
```
Creates and returns a new cdom.FxColor initialized with the specified Sid and
red, green, blue and alpha channel values.

#### func  NewFxColorOrTexture

```go
func NewFxColorOrTexture(texture *cdom.FxTexture, color *cdom.FxColor, paramRef string) (ct *cdom.FxColorOrTexture)
```
Creates and returns a new cdom.FxColorOrTexture referring to either the
specified texture or the specified color. Only at least and at most one of the
arguments (texture, color, paramRef) must be non-nil/non-empty.

#### func  NewFxEffectDef

```go
func NewFxEffectDef(id string, ensureCommonProfile, ensureGlslProfile bool) (me *cdom.FxEffectDef)
```
Creates a new cdom.FxEffectDef with the specified Id, and optionally adds a
common profile, a GLSL profile, or both to it.

#### func  NewFxImageDef

```go
func NewFxImageDef(id, refUrl string) (me *cdom.FxImageDef)
```
Creates and returns a new cdom.FxImageDef with the specified Id, to be
initialized from the specified refUrl.

#### func  NewFxSampler2D

```go
func NewFxSampler2D(image *cdom.FxImageInst, filtering *cdom.FxSamplerFiltering, wrapping *cdom.FxSamplerWrapping) (me *cdom.FxSampler)
```
Creates and returns a new FxSampler of cdom.FxSamplerKind2D, sampling the
specified image with the specified filtering (or cdom.DefaultFxSamplerFiltering
if nil) and the specified wrapping (or cdom.DefaultFxSamplerWrapping if nil).

#### func  NewFxTexture

```go
func NewFxTexture(sampler2DParamRef, texCoord string) (me *cdom.FxTexture)
```
Creates and returns a new cdom.FxTexture sampling from the specified 2D sampler.

--
**godocdown** http://github.com/robertkrimen/godocdown
