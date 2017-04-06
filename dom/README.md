# cdom
--
    import "github.com/metaleap/go-collada/dom"

The go-collada/dom package provides logic-less base data structures for all
resource types used in COLLADA 1.5 or 1.4.1 documents.

These data structures only describe resource definitions and instances, but do
not provide any specific logic or algorithms, such as physics or visual
rendering.

Context: every graphics app uses a variety of both simple and complex resource
types, including for example geometry definitions, image textures, material and
light descriptions and many others.

While for all these resource types, a graphics app most likely manages their
respective run-time representation in its entirety, some basic part of that
representation is "merely descriptive" and should be readable from or writable
to a binary stream ("design-time"). That "merely descriptive" sub-set of
resource definitions is fully contained in this go-collada/dom package.

This package thus allows for things such as server-side procedural asset
generators, networked resource streaming or simple custom
asset-import/export/conversion tools, all of which shouldn't have to needlessly
depend on the graphics, windowing etc. stacks.

NOTE: there are essentially two distinct "modes" or use-cases in which this
package will be active:

1. in graphics-independent, server-side or non-interactive "toolage", dealing
only with the raw resource data where it can be generated, loaded, stored,
manipulated at will with no particular repercussions.

2. in an interactive graphical app:

All "Sync"-related functions pertain to use-case #2, where this package
essentially becomes the live repository for all resource definitions loaded,
used, or manipulated by the parent package at runtime. So now every image
definition in go-collada/dom may have a corresponding GPU-bound texture object
in the app, every dom mesh definition may be bound to an app's "MeshBuffer" or
some such construction, etc.

Structure: generally speaking, all resource types are organized in a consistent
fashion as follows --- fully consistent with the COLLADA specification in
terminology and resource organization:

1. First, there is a FooDef struct for the one-time definition of a unique
resource: GeometryDef, FxImageDef, LightDef, FxMaterialDef etc.

2. Next, there is a smaller FooInst struct for handling many individual
(sometimes parameterized) instantiations of a FooDef: GeometryInst, FxImageInst,
LightInst, FxMaterialInst etc.

3. Finally, there is a light-weight LibFooDefs struct type (encapsulating a
typed hash-table) containing Defs associated with their Id: LibGeometryDefs,
LibFxImageDefs, LibLightDefs, LibFxMaterialDefs etc.

4. The package also provides a pre-initialized global FooDefs variable for each
such Lib type, in simple apps considered the "default / main / only Lib you'll
need": GeometryDefs, FxImageDefs, LightDefs, FxMaterialDefs etc.

5. For more complex use-cases, you can also organize multiple libs for any given
resource type in the global AllFooDefLibs variable, essentially a hash-table of
Libs: AllGeometryDefLibs (of type LibsGeometryDef), AllFxImageDefLibs (of type
LibsFxImageDef), AllLightDefLibs (of type LibsLightDef), AllFxMaterialDefLibs
(of type LibsFxMaterialDef) etc.

Any exported types in this package not following the above pattern should be
considered "auxiliary helpers" rather than primary / "first-class" resource
types.

## Usage

```go
var (
	//	A hash-table that contains LibAnimationDefs libraries associated by their Id.
	AllAnimationDefLibs = LibsAnimationDef{}

	//	The "default" LibAnimationDefs library for AnimationDefs.
	AnimationDefs = AllAnimationDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibAnimationClipDefs libraries associated by their Id.
	AllAnimationClipDefLibs = LibsAnimationClipDef{}

	//	The "default" LibAnimationClipDefs library for AnimationClipDefs.
	AnimationClipDefs = AllAnimationClipDefLibs.AddNew("")
)
```

```go
var (
	//	This callback, set by the core package (or your custom package),
	//	gets called before SyncChanges() proceeds with syncing.
	OnBeforeSyncAll func()

	//	This callback, set by the core package (or your custom package),
	//	gets called after SyncChanges() has finished syncing.
	OnAfterSyncAll func()
)
```

```go
var (
	//	A hash-table that contains LibCameraDefs libraries associated by their Id.
	AllCameraDefLibs = LibsCameraDef{}

	//	The "default" LibCameraDefs library for CameraDefs.
	CameraDefs = AllCameraDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibControllerDefs libraries associated by their Id.
	AllControllerDefLibs = LibsControllerDef{}

	//	The "default" LibControllerDefs library for ControllerDefs.
	ControllerDefs = AllControllerDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibFormulaDefs libraries associated by their Id.
	AllFormulaDefLibs = LibsFormulaDef{}

	//	The "default" LibFormulaDefs library for FormulaDefs.
	FormulaDefs = AllFormulaDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibFxEffectDefs libraries associated by their Id.
	AllFxEffectDefLibs = LibsFxEffectDef{}

	//	The "default" LibFxEffectDefs library for FxEffectDefs.
	FxEffectDefs = AllFxEffectDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibFxImageDefs libraries associated by their Id.
	AllFxImageDefLibs = LibsFxImageDef{}

	//	The "default" LibFxImageDefs library for FxImageDefs.
	FxImageDefs = AllFxImageDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibFxMaterialDefs libraries associated by their Id.
	AllFxMaterialDefLibs = LibsFxMaterialDef{}

	//	The "default" LibFxMaterialDefs library for FxMaterialDefs.
	FxMaterialDefs = AllFxMaterialDefLibs.AddNew("")
)
```

```go
var (
	//	Default texture minification, magnification and MIP-mapping.
	DefaultFxSamplerFiltering = &FxSamplerFiltering{
		FilterMag:     FxFilterKindLinear,
		FilterMin:     FxFilterKindLinear,
		FilterMip:     FxFilterKindLinear,
		MaxAnisotropy: 1,
	}
	//	Default texture repeating and clamping.
	DefaultFxSamplerWrapping = &FxSamplerWrapping{
		BorderColor: ugfx.Rgba32{R: 0, G: 0, B: 0, A: 1},
		WrapS:       FxWrapKindRepeat,
		WrapT:       FxWrapKindRepeat,
		WrapP:       FxWrapKindRepeat,
	}
)
```

```go
var (
	//	A hash-table that contains LibGeometryDefs libraries associated by their Id.
	AllGeometryDefLibs = LibsGeometryDef{}

	//	The "default" LibGeometryDefs library for GeometryDefs.
	GeometryDefs = AllGeometryDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibKxArticulatedSystemDefs libraries associated by their Id.
	AllKxArticulatedSystemDefLibs = LibsKxArticulatedSystemDef{}

	//	The "default" LibKxArticulatedSystemDefs library for KxArticulatedSystemDefs.
	KxArticulatedSystemDefs = AllKxArticulatedSystemDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibKxJointDefs libraries associated by their Id.
	AllKxJointDefLibs = LibsKxJointDef{}

	//	The "default" LibKxJointDefs library for KxJointDefs.
	KxJointDefs = AllKxJointDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibKxModelDefs libraries associated by their Id.
	AllKxModelDefLibs = LibsKxModelDef{}

	//	The "default" LibKxModelDefs library for KxModelDefs.
	KxModelDefs = AllKxModelDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibKxSceneDefs libraries associated by their Id.
	AllKxSceneDefLibs = LibsKxSceneDef{}

	//	The "default" LibKxSceneDefs library for KxSceneDefs.
	KxSceneDefs = AllKxSceneDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibLightDefs libraries associated by their Id.
	AllLightDefLibs = LibsLightDef{}

	//	The "default" LibLightDefs library for LightDefs.
	LightDefs = AllLightDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibNodeDefs libraries associated by their Id.
	AllNodeDefLibs = LibsNodeDef{}

	//	The "default" LibNodeDefs library for NodeDefs.
	NodeDefs = AllNodeDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibPxForceFieldDefs libraries associated by their Id.
	AllPxForceFieldDefLibs = LibsPxForceFieldDef{}

	//	The "default" LibPxForceFieldDefs library for PxForceFieldDefs.
	PxForceFieldDefs = AllPxForceFieldDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibPxMaterialDefs libraries associated by their Id.
	AllPxMaterialDefLibs = LibsPxMaterialDef{}

	//	The "default" LibPxMaterialDefs library for PxMaterialDefs.
	PxMaterialDefs = AllPxMaterialDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibPxModelDefs libraries associated by their Id.
	AllPxModelDefLibs = LibsPxModelDef{}

	//	The "default" LibPxModelDefs library for PxModelDefs.
	PxModelDefs = AllPxModelDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibPxSceneDefs libraries associated by their Id.
	AllPxSceneDefLibs = LibsPxSceneDef{}

	//	The "default" LibPxSceneDefs library for PxSceneDefs.
	PxSceneDefs = AllPxSceneDefLibs.AddNew("")
)
```

```go
var (
	//	A hash-table that contains LibVisualSceneDefs libraries associated by their Id.
	AllVisualSceneDefLibs = LibsVisualSceneDef{}

	//	The "default" LibVisualSceneDefs library for VisualSceneDefs.
	VisualSceneDefs = AllVisualSceneDefLibs.AddNew("")
)
```

#### func  SyncChanges

```go
func SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
that need to be picked up. Call this after you have made any number of changes
to your Defs, Insts or Libs.

#### type AnimSamplerBehavior

```go
type AnimSamplerBehavior int
```

Categorizes the possible behaviors for animation samplers.

```go
const (
	//	The before and after behaviors are not defined.
	AnimSamplerBehaviorUndefined AnimSamplerBehavior = iota

	//	The value for the first (PreBehavior) or last (PostBehavior) is returned.
	AnimSamplerBehaviorConstant

	//	The key is mapped in the [first_key , last_key] interval so that the animation cycles.
	AnimSamplerBehaviorCycle

	//	The animation continues indefinitely.
	AnimSamplerBehaviorCycleRelative

	//	The value follows the line given by the last two keys in the sample.
	AnimSamplerBehaviorGradient

	//	The key is mapped in the [first_key , last_key] interval so that the animation oscillates.
	AnimSamplerBehaviorOscillate
)
```

#### type AnimationChannel

```go
type AnimationChannel struct {
	//	Refers to the source animation sampler.
	Source RefId

	//	Refers to the resource property bound to the output of the sampler.
	Target RefSid
}
```

Declares an output channel of an animation.

#### type AnimationClipDef

```go
type AnimationClipDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	The time in seconds of the beginning of the clip.
	Start float64

	//	The time in seconds of the end of the clip.
	End float64

	//	The animation instances contributing to this animation clip.
	Animations []*AnimationInst

	//	Any formulas used in this animation clip.
	Formulas []*FormulaInst
}
```

Defines a section of a set of animation curves and/or formulas to be used
together as an animation clip.

#### func (*AnimationClipDef) DefaultInst

```go
func (me *AnimationClipDef) DefaultInst() (inst *AnimationClipInst)
```
Returns "the default AnimationClipInst instance" referencing this
AnimationClipDef definition. That instance is created once when this method is
first called on me, and will have its Def field readily set to me.

#### func (*AnimationClipDef) Init

```go
func (me *AnimationClipDef) Init()
```
Initialization

#### func (*AnimationClipDef) NewInst

```go
func (me *AnimationClipDef) NewInst() (inst *AnimationClipInst)
```
Creates and returns a new AnimationClipInst instance referencing this
AnimationClipDef definition. Any AnimationClipInst created by this method will
have its Def field readily set to me.

#### type AnimationClipInst

```go
type AnimationClipInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *AnimationClipDef
}
```

Instantiates an animation clip resource.

#### func (*AnimationClipInst) EnsureDef

```go
func (me *AnimationClipInst) EnsureDef() *AnimationClipDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct AnimationClipDef
according to the current me.DefRef value (by searching AllAnimationClipDefLibs).
Then returns me.Def. (Note, every AnimationClipInst's Def is nil initially,
unless it was created via AnimationClipDef.NewInst().)

#### func (*AnimationClipInst) Init

```go
func (me *AnimationClipInst) Init()
```
Initialization

#### type AnimationDef

```go
type AnimationDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Sources
	HasSources

	//	Allows the formation of a hierarchy of related animations.
	AnimationDefs []*AnimationDef

	//	Describes output channels for the animation.
	Channels []*AnimationChannel

	//	Describes the interpolation sampling functions for the animation.
	Samplers []*AnimationSampler
}
```

Categorizes the declaration of animation information.

#### func (*AnimationDef) DefaultInst

```go
func (me *AnimationDef) DefaultInst() (inst *AnimationInst)
```
Returns "the default AnimationInst instance" referencing this AnimationDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*AnimationDef) Init

```go
func (me *AnimationDef) Init()
```
Initialization

#### func (*AnimationDef) NewInst

```go
func (me *AnimationDef) NewInst() (inst *AnimationInst)
```
Creates and returns a new AnimationInst instance referencing this AnimationDef
definition. Any AnimationInst created by this method will have its Def field
readily set to me.

#### type AnimationInst

```go
type AnimationInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *AnimationDef
}
```

Instantiates an Animation resource.

#### func (*AnimationInst) EnsureDef

```go
func (me *AnimationInst) EnsureDef() *AnimationDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct AnimationDef
according to the current me.DefRef value (by searching AllAnimationDefLibs).
Then returns me.Def. (Note, every AnimationInst's Def is nil initially, unless
it was created via AnimationDef.NewInst().)

#### func (*AnimationInst) Init

```go
func (me *AnimationInst) Init()
```
Initialization

#### type AnimationSampler

```go
type AnimationSampler struct {
	//	Id
	HasId

	//	Inputs. These describe sampling points.
	//	At least one of the Inputs must have its Semantic set to "INTERPOLATION".
	HasInputs

	//	Indicates what the sampled value should be before the first key.
	//	Must be one of the AnimSamplerBehavior* enumerated constants (defaulting to AnimSamplerBehaviorUndefined).
	PreBehavior AnimSamplerBehavior

	//	Indicates what the sampled value should be after the last key.
	//	Must be one of the AnimSamplerBehavior* enumerated constants (defaulting to AnimSamplerBehaviorUndefined).
	PostBehavior AnimSamplerBehavior
}
```

Declares an interpolation sampling function for an animation.

#### type Asset

```go
type Asset struct {
	//	Extras
	HasExtras

	//	Contains the date and time that the parent element was created.
	Created string

	//	Contains the date and time that the parent element was last modified.
	Modified string

	//	Contains a list of words used as search criteria.
	Keywords string

	//	Contains revision information.
	Revision string

	//	Contains a description of the topical subject.
	Subject string

	//	Contains title information.
	Title string

	//	Contains descriptive information about the coordinate system of the geometric data. All
	//	coordinates are right-handed by definition. Valid values are "X", "Y" (the default), or "Z".
	UpAxis string

	//	The unit of distance that applies to all spatial measurements within the scope of this resource.
	Unit struct {
		//	How many real-world meters in one distance unit as a floating-point number.
		//	1.0 for meter, 0.01 for centimeter, 1000 for kilometer etc.
		Meter float64

		//	Name of the distance unit, such as "centimeter", "kilometer", "meter", "inch".
		//	Default is "meter".
		Name string
	}
	//	Contributor information.
	Contributors []*AssetContributor

	//	Provides information about the location of the visual scene in physical space.
	Coverage *AssetGeographicLocation
}
```

Resource-specific asset-management information and meta-data.

#### func  NewAsset

```go
func NewAsset() (me *Asset)
```
Constructor

#### type AssetContributor

```go
type AssetContributor struct {
	Author        string
	AuthorEmail   string
	AuthorWebsite string
	AuthoringTool string
	Comments      string
	Copyright     string
	SourceData    string
}
```

Defines authoring information for asset management.

#### type AssetGeographicLocation

```go
type AssetGeographicLocation struct {
	Longitude        float64
	Latitude         float64
	Altitude         float64
	AltitudeAbsolute bool
}
```

Provides information about the location of the visual scene in physical space.

#### type BaseDef

```go
type BaseDef struct {
	//	Syncability
	BaseSync

	//	Id
	HasId

	//	Name
	HasName

	//	Asset
	HasAsset

	//	Extras
	HasExtras
}
```

Provides a common base for resource definitions.

#### type BaseInst

```go
type BaseInst struct {
	//	Syncability
	BaseSync

	//	Name
	HasName

	//	Sid
	HasSid

	//	Extras
	HasExtras

	//	The Id of the resource definition referenced by this instance.
	DefRef RefId
}
```

Provides a common base for resource instantiations.

#### type BaseLib

```go
type BaseLib struct {
	//	Syncability
	BaseSync

	//	Id
	HasId

	//	Name
	HasName
}
```

Provides a common base for resource libraries.

#### type BaseSync

```go
type BaseSync struct {
	//	This callback, set by the core package (or your custom package), gets called by the
	//	SyncChanges() method. This is the ultimate point in the sync chain where the core package
	//	(or your custom package) picks up the changed contents of this resource.
	//	If the parent is a Lib then this gets called after all its Defs have synced.
	OnSync func()
}
```

Allows for notifying outside packages of changes.

#### func (*BaseSync) SetDirty

```go
func (me *BaseSync) SetDirty()
```
You NEED to call this method if you modified this Def or Inst by setting its
fields directly (instead of using the provided SetFoo() or SetFieldX() methods)
for your changes to be picked up by the core package (or your custom package).

#### func (*BaseSync) SetFieldB

```go
func (me *BaseSync) SetFieldB(field *bool, val bool)
```
If field does not equal val, sets field to val and calls SetDirty().

#### func (*BaseSync) SetFieldF

```go
func (me *BaseSync) SetFieldF(field *float64, val float64)
```
If field does not equal val, sets field to val and calls SetDirty().

#### func (*BaseSync) SyncChanges

```go
func (me *BaseSync) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this Def, Inst or Lib resource that need to be picked up. Call this after you
have made a number of changes to this this resource. Also called by the global
SyncChanges() function.

#### type Bool2

```go
type Bool2 [2]bool
```

Contains two bool values.

#### func (*Bool2) AccessIndex

```go
func (me *Bool2) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Bool3

```go
type Bool3 [3]bool
```

Contains three bool values.

#### func (*Bool3) AccessIndex

```go
func (me *Bool3) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Bool4

```go
type Bool4 [4]bool
```

Contains four bool values.

#### func (*Bool4) AccessIndex

```go
func (me *Bool4) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type CameraDef

```go
type CameraDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Describes the field of view and viewing frustum using canonical parameters.
	Optics CameraOptics

	//	Represents the image sensor of a camera.
	Imager *CameraImager
}
```

Declares a view of the visual scene hierarchy or scene graph.

#### func (*CameraDef) DefaultInst

```go
func (me *CameraDef) DefaultInst() (inst *CameraInst)
```
Returns "the default CameraInst instance" referencing this CameraDef definition.
That instance is created once when this method is first called on me, and will
have its Def field readily set to me.

#### func (*CameraDef) Init

```go
func (me *CameraDef) Init()
```
Initialization

#### func (*CameraDef) NewInst

```go
func (me *CameraDef) NewInst() (inst *CameraInst)
```
Creates and returns a new CameraInst instance referencing this CameraDef
definition. Any CameraInst created by this method will have its Def field
readily set to me.

#### type CameraImager

```go
type CameraImager struct {
	//	Extras
	HasExtras

	//	Techniques
	HasTechniques
}
```

Represents the image sensor of a camera (for example, film or CCD).

#### type CameraInst

```go
type CameraInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *CameraDef
}
```

Instantiates a camera resource.

#### func (*CameraInst) EnsureDef

```go
func (me *CameraInst) EnsureDef() *CameraDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct CameraDef
according to the current me.DefRef value (by searching AllCameraDefLibs). Then
returns me.Def. (Note, every CameraInst's Def is nil initially, unless it was
created via CameraDef.NewInst().)

#### func (*CameraInst) Init

```go
func (me *CameraInst) Init()
```
Initialization

#### type CameraOptics

```go
type CameraOptics struct {
	//	Extras
	HasExtras

	//	Techniques
	HasTechniques

	//	Common-technique profile.
	TC struct {
		//	Aspect ratio of the field of view.
		AspectRatio *SidFloat

		//	Distance to the far clipping plane.
		Zfar SidFloat

		//	Distance to the near clipping plane.
		Znear SidFloat

		//	Orthographic projection type. To use Perspective instead, also set this to nil.
		Orthographic *CameraOrthographic

		//	Perspective projection type. To use Orthographic instead, also set this to nil.
		Perspective *CameraPerspective
	}
}
```

Represents the apparatus on a camera that projects the image onto the image
sensor.

#### type CameraOrthographic

```go
type CameraOrthographic struct {
	//	Horizontal magnification of the view.
	MagX *SidFloat

	//	Vertical magnification of the view.
	MagY *SidFloat
}
```

Describes the field of view of an orthographic camera.

#### type CameraPerspective

```go
type CameraPerspective struct {
	//	Horizontal field of view in degrees.
	FovX *SidFloat

	//	Vertical field of view in degrees.
	FovY *SidFloat
}
```

Describes the field of view of a perspective camera.

#### type ChildNode

```go
type ChildNode struct {
	//	If set, Inst must be nil. An inline node definition.
	Def *NodeDef

	//	If set, Def must be nil. Instantiates a previously defined node.
	Inst *NodeInst
}
```

Used to recursively define hierarchies of nodes.

#### type ControllerDef

```go
type ControllerDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	If set, Skin must be nil; declares this a mesh-morphing controller that deforms meshes and blends them.
	Morph *ControllerMorph

	//	If set, Morph must be nil; declares this a vertex-skinning controller that transforms vertices
	//	based on weighted influences to produce a smoothly changing mesh.
	Skin *ControllerSkin
}
```

Defines generic control information for dynamic content.

#### func (*ControllerDef) DefaultInst

```go
func (me *ControllerDef) DefaultInst() (inst *ControllerInst)
```
Returns "the default ControllerInst instance" referencing this ControllerDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*ControllerDef) Init

```go
func (me *ControllerDef) Init()
```
Initialization

#### func (*ControllerDef) NewInst

```go
func (me *ControllerDef) NewInst() (inst *ControllerInst)
```
Creates and returns a new ControllerInst instance referencing this ControllerDef
definition. Any ControllerInst created by this method will have its Def field
readily set to me.

#### type ControllerInputs

```go
type ControllerInputs struct {
	//	Extras
	HasExtras

	//	Inputs
	HasInputs
}
```

Used to declare skinning joints or morph targets.

#### type ControllerInst

```go
type ControllerInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *ControllerDef

	//	Binds a specific material to this controller instantiation.
	BindMaterial *MaterialBinding

	//	Indicates where a Skin controller is to start to search for the joint nodes it needs.
	//	This element is meaningless for Morph controllers.
	SkinSkeletons []string
}
```

Instantiates a controller resource.

#### func (*ControllerInst) EnsureDef

```go
func (me *ControllerInst) EnsureDef() *ControllerDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct ControllerDef
according to the current me.DefRef value (by searching AllControllerDefLibs).
Then returns me.Def. (Note, every ControllerInst's Def is nil initially, unless
it was created via ControllerDef.NewInst().)

#### func (*ControllerInst) Init

```go
func (me *ControllerInst) Init()
```
Initialization

#### type ControllerMorph

```go
type ControllerMorph struct {
	//	Sources
	HasSources

	//	Which blending method to use: true for relative blending, false for normalized blending.
	Relative bool

	//	Refers to the Geometry that describes the base mesh.
	Source RefId

	//	Input meshes (morph targets) to be blended.
	Targets ControllerInputs
}
```

Describes the data required to blend between sets of static meshes.

#### func  NewControllerMorph

```go
func NewControllerMorph() (me *ControllerMorph)
```
Constructor

#### type ControllerSkin

```go
type ControllerSkin struct {
	//	Sources
	HasSources

	//	Provides extra information about the position and orientation of the base mesh before binding.
	BindShapeMatrix unum.Mat4

	//	Describes a per-vertex combination of joints and weights used in this skin.
	//	An index of –1 into the array of joints refers to the bind shape.
	//	Weights should be normalized before use.
	VertexWeights IndexedInputs

	//	Aggregates the per-joint information needed for this skin.
	Joints ControllerInputs

	//	Refers to the base mesh (a static mesh or a morphed mesh).
	//	This also provides the bind-shape of the skinned mesh.
	Source RefId
}
```

Contains vertex and primitive information sufficient to describe blend-weight
skinning.

#### func  NewControllerSkin

```go
func NewControllerSkin() (me *ControllerSkin)
```
Constructor

#### type Document

```go
type Document struct {
	//	Asset
	HasAsset

	//	Extras
	HasExtras

	//	Describes a complete, fully self-contained scene graph.
	Scene *Scene
}
```

Encapsulates a complete, fully self-contained scene graph. Note, resource
definition libraries are organized in "globals" (exported package variables)
rather than Document struct instances.

#### type Extra

```go
type Extra struct {
	//	Id
	HasId

	//	Name
	HasName

	//	Asset
	HasAsset

	//	Techniques
	HasTechniques

	//	A hint as to the type of information that this particular Extra represents.
	Type string
}
```

Provides arbitrary additional information about or related to its parent
resource.

#### type Float2

```go
type Float2 [2]float64
```

Contains two float64 values.

#### func (*Float2) AccessIndex

```go
func (me *Float2) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Float2x2

```go
type Float2x2 [4]float64
```

Contains four float64 values.

#### func (*Float2x2) AccessIndex

```go
func (me *Float2x2) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional or two-dimensional
indices.

#### type Float2x3

```go
type Float2x3 [6]float64
```

Contains six float64 values.

#### func (*Float2x3) AccessIndex

```go
func (me *Float2x3) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Float2x4

```go
type Float2x4 [8]float64
```

Contains eight float64 values.

#### func (*Float2x4) AccessIndex

```go
func (me *Float2x4) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Float3

```go
type Float3 [3]float64
```

Contains three float64 values.

#### func (*Float3) AccessIndex

```go
func (me *Float3) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Float3x2

```go
type Float3x2 [6]float64
```

Contains six float64 values.

#### func (*Float3x2) AccessIndex

```go
func (me *Float3x2) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Float3x3

```go
type Float3x3 [9]float64
```

Contains nine float64 values.

#### func (*Float3x3) AccessIndex

```go
func (me *Float3x3) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Float3x4

```go
type Float3x4 [12]float64
```

Contains twelve float64 values.

#### func (*Float3x4) AccessIndex

```go
func (me *Float3x4) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Float4

```go
type Float4 [4]float64
```

Contains four float64 values.

#### func (*Float4) AccessIndex

```go
func (me *Float4) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Float4x2

```go
type Float4x2 [8]float64
```

Contains eight float64 values.

#### func (*Float4x2) AccessIndex

```go
func (me *Float4x2) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Float4x3

```go
type Float4x3 [12]float64
```

Contains twelve float64 values.

#### func (*Float4x3) AccessIndex

```go
func (me *Float4x3) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Float4x4

```go
type Float4x4 [16]float64
```

Contains sixteen float64 values.

#### func (*Float4x4) AccessIndex

```go
func (me *Float4x4) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Float7

```go
type Float7 [7]float64
```

Contains seven float64 values.

#### func (*Float7) AccessIndex

```go
func (me *Float7) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Formula

```go
type Formula struct {
	//	If set, Inst must be nil.
	Def *FormulaDef

	//	If set, Def must be nil.
	Inst *FormulaInst
}
```

Provides either a formula definition or a formula instance.

#### type FormulaDef

```go
type FormulaDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Sid
	HasSid

	//	NewParams
	HasParamDefs

	//	Techniques
	HasTechniques

	//	A parameter that specifies the result variable of the formula.
	Target ParamOrFloat

	//	Common-technique profile.
	TC struct {
		//	Any valid MathML (content) XML defining this formula.
		MathML string
	}
}
```

There are many ways to describe a formula. COLLADA uses MathML as its
common-technique.

#### func (*FormulaDef) DefaultInst

```go
func (me *FormulaDef) DefaultInst() (inst *FormulaInst)
```
Returns "the default FormulaInst instance" referencing this FormulaDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*FormulaDef) Init

```go
func (me *FormulaDef) Init()
```
Initialization

#### func (*FormulaDef) NewInst

```go
func (me *FormulaDef) NewInst() (inst *FormulaInst)
```
Creates and returns a new FormulaInst instance referencing this FormulaDef
definition. Any FormulaInst created by this method will have its Def field
readily set to me.

#### type FormulaInst

```go
type FormulaInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	SetParams
	HasParamInsts

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *FormulaDef
}
```

Instantiates a formula resource.

#### func (*FormulaInst) EnsureDef

```go
func (me *FormulaInst) EnsureDef() *FormulaDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct FormulaDef
according to the current me.DefRef value (by searching AllFormulaDefLibs). Then
returns me.Def. (Note, every FormulaInst's Def is nil initially, unless it was
created via FormulaDef.NewInst().)

#### func (*FormulaInst) Init

```go
func (me *FormulaInst) Init()
```
Initialization

#### type FxAnnotation

```go
type FxAnnotation struct {
	//	Name
	HasName

	//	Annotation value.
	Value interface{}
}
```

Annotations communicate metadata from the Effect Runtime to the application only
and are not otherwise interpreted within the go-collada/dom package.

#### type FxBinding

```go
type FxBinding struct {
	//	Which effect parameter to bind.
	Semantic string

	//	Refers to the value to bind to the specified semantic.
	Target RefSid
}
```

Binds values to uniform inputs of a shader or binds values to effect parameters
upon instantiation.

#### type FxColor

```go
type FxColor struct {
	//	Sid
	HasSid

	//	Describes the literal color of the parent FxColorOrTexture.
	ugfx.Rgba32
}
```

Used to describe the literal color of an FxColorOrTexture.

#### func (*FxColor) AccessField

```go
func (me *FxColor) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "R", "G", "B", "A".

#### func (*FxColor) AccessIndex

```go
func (me *FxColor) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices 0 through 3.

#### type FxColorOrTexture

```go
type FxColorOrTexture struct {
	//	Specifies from which channel to take transparency information.
	//	Must be one of the FxTextureOpaque* enumerated constants.
	Opaque FxTextureOpaque

	//	If set, refers to a previously-defined parameter in the current scope that provides
	//	four float values describing the literal color of this value.
	ParamRef RefParam

	//	If set, describes he literal color of this value.
	Color *FxColor

	//	If set, refers to a previously-defined FxSampler with a Kind of FxSamplerKind2D.
	Texture *FxTexture
}
```

Describes color attributes of fixed-function shaders inside FxProfileCommon
effects.

#### type FxCreate

```go
type FxCreate struct {
	//	Specifies the length of the 2D array, 3D array or cube-map array.
	ArrayLength uint64

	//	Specifies an image's pixel or compression format.
	//	If not present, the format is assumed to be R8G8B8A8 linear.
	Format *FxCreateFormat
}
```

Fields shared by FxCreate2D, FxCreate3D and FxCreateCube

#### type FxCreate2D

```go
type FxCreate2D struct {
	//	ArrayLength and Format
	FxCreate

	//	Either Exact or Ratio, but not both, must be present.
	Size struct {
		//	Specifies that the surface should be sized to these exact dimensions.
		Exact *FxCreate2DSizeExact

		//	Specifies that the image size should be relative to the size of the viewport.
		Ratio *FxCreate2DSizeRatio
	}
	//	MIP information. Ignored if Unnormalized is true.
	Mips *FxCreateMips

	//	Unnormalized addressing of texels. (0-W, 0-H).
	Unnormalized bool

	//	Specifies which 2D image to initialize and which MIP level to initialize.
	InitFrom []*FxCreateInitFrom
}
```

Assists in the manual creation of a 2D FxImageDef asset.

#### type FxCreate2DSizeExact

```go
type FxCreate2DSizeExact struct {
	//	width in pixels
	Width uint64

	//	height in pixels
	Height uint64
}
```

Specifies that the surface should be sized to these exact dimensions.

#### type FxCreate2DSizeRatio

```go
type FxCreate2DSizeRatio struct {
	//	Relative width where 1.0 represents viewport width.
	Width float64

	//	Relative height where 1.0 represents viewport height.
	Height float64
}
```

Specifies that the image size should be relative to the size of the viewport.
For example, 1,1 is the same size as the viewport; 0.5,0.5 is 1/4 the size of
the viewport and half as long in either direction.

#### type FxCreate3D

```go
type FxCreate3D struct {
	//	ArrayLength and Format
	FxCreate

	//	Specifies that the surface should be sized to these exact dimensions.
	Size struct {
		//	Width in pixels for this 3D texture.
		Width uint64

		//	Height in pixels for this 3D texture.
		Height uint64

		//	Depth in pixels for this 3D texture.
		Depth uint64
	}
	//	MIP information.
	Mips FxCreateMips

	//	Specifies which 3D image to initialize and which MIP level to initialize.
	InitFrom []*FxCreate3DInitFrom
}
```

Assists in the manual creation of a 3D FxImageDef asset.

#### type FxCreate3DInitFrom

```go
type FxCreate3DInitFrom struct {
	//	Raw or RefUrl, ArrayIndex and MipIndex
	FxCreateInitFrom

	//	Specifies the slice (depth level) within the MIP that is to be initialized.
	Depth uint64
}
```

Initializes an entire 3D texture or portions of a 3D texture from referenced or
embedded data.

#### type FxCreateCube

```go
type FxCreateCube struct {
	//	ArrayLength and Format
	FxCreate

	//	Specifies that the cube surfaces should be sized to these exact dimensions.
	Size struct {
		//	Width and height are identical across all faces in a cube-map.
		Width uint64
	}
	//	MIP information.
	Mips FxCreateMips

	//	Specifies which cube image to initialize, which MIP level to initialize,
	//	and which cube face within the MIP that is to be initialized.
	InitFrom []*FxCreateCubeInitFrom
}
```

Assists in the manual creation of a cube-map FxImageDef asset.

#### type FxCreateCubeInitFrom

```go
type FxCreateCubeInitFrom struct {
	//	Raw or RefUrl, ArrayIndex and MipIndex
	FxCreateInitFrom

	//	Specifies the cube-map face within the MIP that is to be initialized.
	//	Must be one of the FxCubeFace* enumerated constants.
	Face FxCubeFace
}
```

Initializes an entire cube-map or portions of a cube-map from referenced or
embedded data.

#### type FxCreateFormat

```go
type FxCreateFormat struct {
	//	Contains a string representing the profile- and platform-specific texel format
	//	that the author would like this surface to use. If this element is not specified,
	//	or if it is specified but the application cannot process the specified format,
	//	then the application uses the Hint.
	Exact string

	//	If this is not set, then use a common format R8G8B8A8 with linear color gradient, not sRGB.
	Hint *FxCreateFormatHint
}
```

Describes the formatting or memory layout expected of an FxImageDef asset.

#### type FxCreateFormatHint

```go
type FxCreateFormatHint struct {
	//	Describes the per-texel layout of the format.
	//	Must be one of the FxFormatChannels* enumerated constants.
	Channels FxFormatChannels

	//	Describes the range of texel channel values. Each channel represents a range of values.
	//	Some example ranges are signed or unsigned integers, or are within a clamped range
	//	such as 0.0f to 1.0f, or are a high dynamic range via floating point.
	//	Must be one of the FxFormatRange* enumerated constants.
	Range FxFormatRange

	//	Identifies the precision of the texel channel value.
	//	Must be one of the FxFormatPrecision* enumerated constants.
	Precision FxFormatPrecision

	//	Optional custom / application-specific color-space information.
	Space string
}
```

Describes features and characteristics to select an appropriate format for image
creation.

#### type FxCreateInitFrom

```go
type FxCreateInitFrom struct {
	//	Raw and RefUrl
	FxInitFrom

	//	Specifies which array element in the image to initialize (fill).
	ArrayIndex uint64

	//	Specifies which MIP level in the image to initialize.
	MipIndex uint64
}
```

Initializes an entire image or portions of an image from referenced or embedded
data.

#### type FxCreateMips

```go
type FxCreateMips struct {
	//	Desired number of MIP levels. Special values: 1 is "no MIP levels", 0 is "all MIP levels".
	Levels uint64

	//	If false, initializes higher MIP levels if data does not exist in a file.
	//	If true, no MIP levels are ever automatically initialized.
	NoAutoGen bool
}
```

MIP information

#### type FxCubeFace

```go
type FxCubeFace int
```

Categorizes one of the six sub-images (faces) in a cube-map.

```go
const (
	//	Cube-map face "X negative"
	FxCubeFaceNx FxCubeFace = 0x8516

	//	Cube-map face "Y negative"
	FxCubeFaceNy FxCubeFace = 0x8518

	//	Cube-map face "Z negative"
	FxCubeFaceNz FxCubeFace = 0x851A

	//	Cube-map face "X positive"
	FxCubeFacePx FxCubeFace = 0x8515

	//	Cube-map face "Y positive"
	FxCubeFacePy FxCubeFace = 0x8517

	//	Cube-map face "Z positive"
	FxCubeFacePz FxCubeFace = 0x8519
)
```

#### type FxEffectDef

```go
type FxEffectDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	NewParams
	HasFxParamDefs

	//	Application-specific FX metadata
	Annotations []*FxAnnotation

	//	Rendering pipeline(s).
	Profiles []*FxProfile
}
```

Defines the equations necessary for the visual appearance of geometry or
screen-space image processing.

#### func (*FxEffectDef) Common

```go
func (me *FxEffectDef) Common() (prof *FxProfile)
```
Returns the first FxProfile with a Common in me.Profiles.

#### func (*FxEffectDef) DefaultInst

```go
func (me *FxEffectDef) DefaultInst() (inst *FxEffectInst)
```
Returns "the default FxEffectInst instance" referencing this FxEffectDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*FxEffectDef) Init

```go
func (me *FxEffectDef) Init()
```
Initialization

#### func (*FxEffectDef) NewInst

```go
func (me *FxEffectDef) NewInst() (inst *FxEffectInst)
```
Creates and returns a new FxEffectInst instance referencing this FxEffectDef
definition. Any FxEffectInst created by this method will have its Def field
readily set to me.

#### type FxEffectInst

```go
type FxEffectInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	SetParams
	HasParamInsts

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *FxEffectDef

	//	Platform-specific hints of which techniques to use in this effect.
	TechniqueHints []*FxEffectInstTechniqueHint
}
```

Instantiates an effect resource.

#### func (*FxEffectInst) EnsureDef

```go
func (me *FxEffectInst) EnsureDef() *FxEffectDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct FxEffectDef
according to the current me.DefRef value (by searching AllFxEffectDefLibs). Then
returns me.Def. (Note, every FxEffectInst's Def is nil initially, unless it was
created via FxEffectDef.NewInst().)

#### func (*FxEffectInst) Init

```go
func (me *FxEffectInst) Init()
```
Initialization

#### type FxEffectInstTechniqueHint

```go
type FxEffectInstTechniqueHint struct {
	//	Defines a string that specifies for which platform this hint is intended. Optional.
	Platform string

	//	A reference to the name of the platform. Required.
	Ref string

	//	Specifies for which API profile this hint is intended.
	//	Optional. If set, must be either "COMMON" or "GLSL".
	Profile string
}
```

Adds a hint for a platform of which technique to use in this effect.

#### type FxFilterKind

```go
type FxFilterKind int
```

Categorizes the kinds of filtering used in FxSamplerFiltering.

```go
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
```

#### type FxFormatChannels

```go
type FxFormatChannels int
```

Categorizes the Channels of an FxCreateFormatHint.

```go
const (
	//	Depth map, often used for displacement, parellax, relief, or shadow mapping.
	FxFormatChannelsD FxFormatChannels = iota + 1

	//	Luminance map, often used for light mapping.
	FxFormatChannelsL

	//	Luminance with alpha map, often used for light mapping.
	FxFormatChannelsLa

	//	RGB color map
	FxFormatChannelsRgb

	//	RGB color with alpha map. Often used for color plus transparency
	//	or other things packed into channel A, such as specular power.
	FxFormatChannelsRgba

	//	RGB color with shared exponent for HDR.
	FxFormatChannelsRgbe
)
```

#### type FxFormatPrecision

```go
type FxFormatPrecision int
```

Categorizes the Precision of an FxCreateFormatHint.

```go
const (
	//	Designer does not care as long as it provides "reasonable" precision and performance.
	FxFormatPrecisionDefault FxFormatPrecision = iota

	//	For integers, this typically represents 16 to 32 bits. For floating points, typically
	//	24 to 32 bits.
	FxFormatPrecisionHigh

	//	For integers, this typically represents 8 bits. For floating points, typically 16 bits.
	FxFormatPrecisionLow

	//	Typically 32 bits or 64 bits if available. 64 bits has been separated into its own category
	//	beyond HIGH because it typically has significant performance impact.
	FxFormatPrecisionMax

	//	For integers, this typically represents 8 to 24 bits.
	//	For floating points, typically 16 to 32 bits.
	FxFormatPrecisionMid
)
```

#### type FxFormatRange

```go
type FxFormatRange int
```

Categorizes the Range of an FxCreateFormatHint.

```go
const (
	//	Format should support full floating-point ranges.
	//	High precision is expected to be 32 bits.
	//	Mid precision may be 16 to 32 bits.
	//	Low precision is expected to be 16 bits.
	FxFormatRangeFloat FxFormatRange = iota + 1

	//	Format represents signed integer numbers. For example, 8 bits is -128 to 127.
	FxFormatRangeSint

	//	Format represents a decimal value that remains within the -1 to 1 range.
	//	Implementation could be integer-fixed-point or floating point.
	FxFormatRangeSnorm

	//	Format represent unsigned integer numbers. For example, 8 bits is 0 to 255.
	FxFormatRangeUint

	//	Format represents a decimal value that remains within the 0 to 1 range.
	//	Implementation could be integer-fixed-point or floating point.
	FxFormatRangeUnorm
)
```

#### type FxGlslTechniques

```go
type FxGlslTechniques map[string]*FxTechniqueGlsl
```

A hash-table of GLSL techniques mapped to their scoped identifiers.

#### type FxImageDef

```go
type FxImageDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Indicates whether this image represents a render target.
	Renderable struct {
		//	If true, defines the image as a render target, meaning the image can be rendered to.
		Is bool

		//	Indicates whether, when instantiated, the render target is to be shared among all
		//	instances instead of being cloned.
		Shared bool
	}
	//	If set, initializes a custom 2D image by specifying its size, viewport ratio, MIP levels,
	//	normalization, pixel format, and data sources. It also supports arrays of 2D images.
	Create2D *FxCreate2D

	//	If set, initializes a custom 3D image (a volumetric image) by specifying its size, MIP level,
	//	pixel format, and data sources. It also supports arrays of 3D images.
	Create3D *FxCreate3D

	//	If set, initializes the six faces of a cube by specifying its size, MIP level, pixel format,
	//	and data sources. It also supports arrays of images on each of the cube faces.
	//	It also supports arrays of cube images.
	CreateCube *FxCreateCube

	//	If set, initializes the image from a URL (for example, a file) or binary image data.
	InitFrom *FxImageInitFrom
}
```

Declares the storage for the graphical representation of an object.

#### func (*FxImageDef) DefaultInst

```go
func (me *FxImageDef) DefaultInst() (inst *FxImageInst)
```
Returns "the default FxImageInst instance" referencing this FxImageDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*FxImageDef) Init

```go
func (me *FxImageDef) Init()
```
Initialization

#### func (*FxImageDef) NewInst

```go
func (me *FxImageDef) NewInst() (inst *FxImageInst)
```
Creates and returns a new FxImageInst instance referencing this FxImageDef
definition. Any FxImageInst created by this method will have its Def field
readily set to me.

#### type FxImageInitFrom

```go
type FxImageInitFrom struct {
	//	Raw and RefUrl
	FxInitFrom

	//	If false, initializes higher MIP levels if data does not exist in a file.
	//	If true, no MIP levels are ever automatically initialized.
	NoAutoMip bool
}
```

Initializes an entire image or portions of an image from referenced or embedded
data.

#### func  NewFxImageInitFrom

```go
func NewFxImageInitFrom(refUrl string) (me *FxImageInitFrom)
```
Constructor

#### type FxImageInst

```go
type FxImageInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *FxImageDef
}
```

Instantiates an image resource.

#### func (*FxImageInst) EnsureDef

```go
func (me *FxImageInst) EnsureDef() *FxImageDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct FxImageDef
according to the current me.DefRef value (by searching AllFxImageDefLibs). Then
returns me.Def. (Note, every FxImageInst's Def is nil initially, unless it was
created via FxImageDef.NewInst().)

#### func (*FxImageInst) Init

```go
func (me *FxImageInst) Init()
```
Initialization

#### type FxInitFrom

```go
type FxInitFrom struct {
	//	Embedded binary image data; used if RefUrl is empty and Raw.Data is not.
	Raw struct {
		//	Contains the embedded binary image data as a sequence of bytes. Typically contains all
		//	the necessary information including header info such as width and height.
		Data []byte

		//	Specifies which codec decodes the image's descriptions and data.
		//	This is usually its typical file extension, such as "BMP", "JPG", "DDS", "TGA".
		Format string
	}
	//	Contains the URL of a file from which to take initialization data. Assumes the characteristics of the
	//	file: if it is a complex format such as DDS, this might include cube maps, volumes, MIPs, and so on.
	RefUrl string
}
```

Initializes an entire image or portions of an image from referenced or embedded
data.

#### type FxMaterialDef

```go
type FxMaterialDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	The parameterized effect instantiation that fully describes and defines this material.
	Effect FxEffectInst
}
```

Defines the equations necessary for the visual appearance of geometry or
screen-space image processing.

#### func (*FxMaterialDef) DefaultInst

```go
func (me *FxMaterialDef) DefaultInst() (inst *FxMaterialInst)
```
Returns "the default FxMaterialInst instance" referencing this FxMaterialDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*FxMaterialDef) Init

```go
func (me *FxMaterialDef) Init()
```
Initialization

#### func (*FxMaterialDef) NewInst

```go
func (me *FxMaterialDef) NewInst() (inst *FxMaterialInst)
```
Creates and returns a new FxMaterialInst instance referencing this FxMaterialDef
definition. Any FxMaterialInst created by this method will have its Def field
readily set to me.

#### type FxMaterialInst

```go
type FxMaterialInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *FxMaterialDef

	//	Which symbol defined from within the geometry this material binds to.
	Symbol string

	//	Binds values to uniform inputs of a shader or binds values to effect parameters upon instantiation.
	Bindings []*FxBinding

	//	Binds vertex inputs to effect parameters upon instantiation.
	VertexInputBindings []*FxVertexInputBinding
}
```

Instantiates a material resource.

#### func (*FxMaterialInst) AccessField

```go
func (me *FxMaterialInst) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "Symbol".

#### func (*FxMaterialInst) EnsureDef

```go
func (me *FxMaterialInst) EnsureDef() *FxMaterialDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct FxMaterialDef
according to the current me.DefRef value (by searching AllFxMaterialDefLibs).
Then returns me.Def. (Note, every FxMaterialInst's Def is nil initially, unless
it was created via FxMaterialDef.NewInst().)

#### func (*FxMaterialInst) Init

```go
func (me *FxMaterialInst) Init()
```
Initialization

#### type FxParamDef

```go
type FxParamDef struct {
	//	Sid and Value
	ParamDef

	//	Application-specific FX metadata
	Annotations []*FxAnnotation

	//	Specifies constant, external, or uniform parameters.
	Modifier string

	//	Provides metadata that describes the purpose of a parameter declaration.
	Semantic string
}
```

Declares a new parameter for its parent FX-related resource, and assigns it an
initial value.

#### func (*FxParamDef) AccessField

```go
func (me *FxParamDef) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "Modifier", "Semantic".

#### type FxParamDefs

```go
type FxParamDefs map[string]*FxParamDef
```

A hash-table containing parameter declarations of this FX-related resource.

#### func (FxParamDefs) Set

```go
func (me FxParamDefs) Set(sid string, val interface{})
```
If me does not contain an FxParamDef with the specified Sid, adds it. Next, sets
the value of the FxParamDef with the specified Sid in me to val.

#### type FxPass

```go
type FxPass struct {
	//	Sid
	HasSid

	//	Extras
	HasExtras

	//	Application-specific FX metadata
	Annotations []*FxAnnotation

	//	Contains all rendering states to set up for this pass.
	States map[string]*FxPassState

	//	Links multiple shaders together to produce a pipeline for geometry processing.
	Program *FxPassProgram

	//	Contains evaluation elements for this rendering pass.
	Evaluate *FxPassEvaluation
}
```

Provides a static declaration of all the render states, shaders, and settings
for one rendering pipeline.

#### func  NewFxPass

```go
func NewFxPass() (me *FxPass)
```
Constructor

#### type FxPassEvaluation

```go
type FxPassEvaluation struct {
	//	Instructs the FX Runtime what kind of geometry to submit.
	Draw string

	//	Color-information render target
	Color struct {
		//	Specifies whether this render target image is to be cleared, and which value to use.
		Clear *FxPassEvaluationClearColor

		//	Specifies which FxImageDef will receive the color information from the output of this pass.
		Target *FxPassEvaluationTarget
	}
	//	Depth-information render target
	Depth struct {
		//	Specifies whether this render target image is to be cleared, and which value to use.
		Clear *FxPassEvaluationClearDepth

		//	Specifies which FxImageDef will receive the depth information from the output of this pass.
		Target *FxPassEvaluationTarget
	}
	//	Stencil-information render target
	Stencil struct {
		//	Specifies whether this render target image is to be cleared, and which value to use.
		Clear *FxPassEvaluationClearStencil

		//	Specifies which FxImageDef will receive the stencil information from the output of this pass.
		Target *FxPassEvaluationTarget
	}
}
```

Contains evaluation elements for a rendering pass.

#### type FxPassEvaluationClearColor

```go
type FxPassEvaluationClearColor struct {
	//	Default clear-color value
	ugfx.Rgba32

	//	Which of the multiple render targets is being set. The default is 0.
	Index uint64
}
```

Specifies whether a color-information render target image is to be cleared, and
which value to use.

#### type FxPassEvaluationClearDepth

```go
type FxPassEvaluationClearDepth struct {
	//	Default clear-depth value
	F float64

	//	Which of the multiple render targets is being set. The default is 0.
	Index uint64
}
```

Specifies whether a depth-information render target image is to be cleared, and
which value to use.

#### type FxPassEvaluationClearStencil

```go
type FxPassEvaluationClearStencil struct {
	//	Default clear-stencil value
	B byte

	//	Which of the multiple render targets is being set. The default is 0.
	Index uint64
}
```

Specifies whether a stencil-information render target image is to be cleared,
and which value to use.

#### type FxPassEvaluationTarget

```go
type FxPassEvaluationTarget struct {
	//	Indexes one of the Multiple Render Targets.
	Index uint64

	//	Indexes a sub-image inside a target surface, specifically: a layer of a 3D texture.
	Slice uint64

	//	Indexes a sub-image inside a target surface, specifically: a single MIP-map level.
	Mip uint64

	//	Indexes a sub-image inside a target surface, specifically: a unique cube face.
	//	Must be one of the FxCubeFace* enumerated constants.
	CubeFace FxCubeFace

	//	If set, Image is ignored; this render target references a sampler parameter to determine which image to use.
	Sampler RefParam

	//	If set (and Sampler is empty), this render target directly instantiates a renderable image.
	Image *FxImageInst
}
```

Specifies which FxImageDef will receive the information from the output of a
pass.

#### func  NewFxPassEvaluationTarget

```go
func NewFxPassEvaluationTarget() (me *FxPassEvaluationTarget)
```
Constructor

#### type FxPassProgram

```go
type FxPassProgram struct {
	//	Information for binding the shader variables to effect parameters.
	BindAttributes []*FxPassProgramBindAttribute

	//	Binds a uniform shader variable to a parameter or a value.
	BindUniforms []*FxPassProgramBindUniform

	//	Setup and compilation information for shaders such as vertex and pixel shaders.
	Shaders []*FxPassProgramShader
}
```

Links multiple shaders together to produce a pipeline for geometry processing.

#### type FxPassProgramBindAttribute

```go
type FxPassProgramBindAttribute struct {
	//	The identifier for a vertex attribute variable in the shader
	//	(a formal function parameter or in-scope global).
	Symbol string

	//	Contains an alternative name to the attribute variable
	//	for semantic binding to geometry vertex inputs.
	Semantic string
}
```

Binds semantics to vertex attribute inputs of a shader.

#### type FxPassProgramBindUniform

```go
type FxPassProgramBindUniform struct {
	//	The identifier for a uniform input parameter to the shader
	//	(a formal function parameter or in-scope global) that will be bound to an external resource.
	Symbol string

	//	If set, refers to a previously defined parameter providing the uniform value to be bound.
	ParamRef RefParam

	//	If set, the uniform value to be bound.
	Value interface{}
}
```

Binds values to uniform inputs of a shader or binds values to effect parameters
upon instantiation.

#### type FxPassProgramShader

```go
type FxPassProgramShader struct {
	//	Extras
	HasExtras

	//	In which pipeline stage this programmable shader is designed to execute.
	//	Must be one of the FxShaderStage* enumerated constants.
	Stage FxShaderStage

	//	Concatenates the source code for the shader from one or more sources.
	Sources []FxPassProgramShaderSources
}
```

Declares and prepares a shader for execution in the rendering pipeline of a
pass.

#### type FxPassProgramShaderSources

```go
type FxPassProgramShaderSources struct {
	//	If true, S is an import reference; otherwise, S is code.
	IsImportRef bool

	//	The code or import reference.
	S string
}
```

Contains either code or an import reference.

#### type FxPassState

```go
type FxPassState struct {
	//	If set, Value is ignored; refers to a previously defined parameter providing the value for this rendering state.
	Param RefParam

	//	If set (and Param is empty), the value for this rendering state.
	Value string

	//	State-specific optional index attribute.
	Index float64
}
```

Represents a rendering state for a pass.

#### type FxProfile

```go
type FxProfile struct {
	//	Id
	HasId

	//	Asset
	HasAsset

	//	Extras
	HasExtras

	//	NewParams
	HasFxParamDefs

	//	If set, Glsl must be nil, and this FxProfile represents a common, fixed-function shader pipeline.
	Common *FxProfileCommon

	//	If set, Common must be nil, and this FxProfile represents an OpenGL Shading Language pipeline.
	Glsl *FxProfileGlsl
}
```

An FX profile represents a shader-based rendering pipeline.

#### func  NewProfile

```go
func NewProfile() (me *FxProfile)
```

#### type FxProfileCommon

```go
type FxProfileCommon struct {
	//	Declares the only technique for this effect.
	Technique FxTechniqueCommon
}
```

This FX profile provides platform-independent declarations for the common,
fixed-function shader.

#### type FxProfileGlsl

```go
type FxProfileGlsl struct {
	//	The type of platform. This is a vendor-defined character string that
	//	indicates the platform or capability target for the technique. Defaults to "PC".
	Platform string

	//	GLSL shader sources
	CodesIncludes []FxProfileGlslCodeInclude

	//	Declares the techniques for this effect.
	Techniques FxGlslTechniques
}
```

This FX profile provides platform-specific declarations for the OpenGL Shading
Language.

#### func  NewFxProfileGlsl

```go
func NewFxProfileGlsl() (me *FxProfileGlsl)
```
Constructor

#### type FxProfileGlslCodeInclude

```go
type FxProfileGlslCodeInclude struct {
	//	Source code or include reference
	SidString

	//	Indicates whether SidString is an import reference (true) or source code (false).
	IsInclude bool
}
```

GLSL shader sources

#### func (*FxProfileGlslCodeInclude) AccessField

```go
func (me *FxProfileGlslCodeInclude) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "IsInclude".

#### type FxSampler

```go
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
```

Declares a texture sampler.

#### func  NewFxSampler

```go
func NewFxSampler() (me *FxSampler)
```
Constructor

#### type FxSamplerFiltering

```go
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
```

Controls texture minification, magnification and MIP-mapping.

#### type FxSamplerImage

```go
type FxSamplerImage struct {
	//	Sid, Name, Extras, DefRef
	FxImageInst
}
```

Instantiates an image targeted for samplers.

#### type FxSamplerKind

```go
type FxSamplerKind int
```

Categorizes the kind of an FxSampler.

```go
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
```

#### type FxSamplerStates

```go
type FxSamplerStates struct {
	//	Extras
	HasExtras

	//	Controls texture minification, magnification and MIP-mapping.
	Filtering *FxSamplerFiltering

	//	Controls texture repeating and clamping.
	Wrapping *FxSamplerWrapping
}
```

Allows users to modify an effect's sampler state from a material.

#### func  NewFxSamplerStates

```go
func NewFxSamplerStates() (me *FxSamplerStates)
```
Constructor

#### type FxSamplerWrapping

```go
type FxSamplerWrapping struct {
	//	When reading past the edge of the texture address space
	//	based on the wrap modes involving clamps, this color takes over.
	BorderColor ugfx.Rgba32

	//	Controls texture repeating and clamping of the S coordinate.
	//	Must be one of the WrapKind* enumerated constants.
	WrapS FxWrapKind

	//	Controls texture repeating and clamping of the T coordinate.
	//	Must be one of the WrapKind* enumerated constants.
	WrapT FxWrapKind

	//	Controls texture repeating and clamping of the P coordinate.
	//	Must be one of the WrapKind* enumerated constants.
	WrapP FxWrapKind
}
```

Controls texture repeating and clamping.

#### type FxShaderStage

```go
type FxShaderStage int
```

Categorizes the kind of Stage in an FxPassProgramShader.

```go
const (
	//	This programmable shader is designed to execute in the Tessellation pipeline stage.
	FxShaderStageTessellation FxShaderStage = iota + 1

	//	This programmable shader is designed to execute in the Vertex pipeline stage.
	FxShaderStageVertex

	//	This programmable shader is designed to execute in the Geometry pipeline stage.
	FxShaderStageGeometry

	//	This programmable shader is designed to execute in the Fragment pipeline stage.
	FxShaderStageFragment

	//	This programmable shader is designed to execute in the Compute pipeline stage.
	FxShaderStageCompute
)
```

#### type FxTechnique

```go
type FxTechnique struct {
	//	Id
	HasId

	//	Sid
	HasSid

	//	Asset
	HasAsset

	//	Extras
	HasExtras
}
```

Holds a description of the textures, samplers, shaders, parameters, and passes
necessary for rendering this effect using one method.

#### type FxTechniqueCommon

```go
type FxTechniqueCommon struct {
	//	Id, Sid, Asset, Extras
	FxTechnique

	//	Must be one of the FxTechniqueKind* enumerated constants.
	Kind FxTechniqueKind

	//	Declares the amount of light emitted from the surface of this object.
	Emission *FxColorOrTexture

	//	Declares the color of a perfect mirror reflection.
	Reflective *FxColorOrTexture

	//	Declares the amount of perfect mirror reflection to be added to the reflected light
	//	as a value between 0.0 and 1.0.
	Reflectivity *ParamOrSidFloat

	//	Declares the color of perfectly refracted light.
	Transparent *FxColorOrTexture

	//	Declares the amount of perfectly refracted light added to the reflected color
	//	as a scalar value between 0.0 and 1.0.
	Transparency *ParamOrSidFloat

	//	Declares the index of refraction for perfectly refracted light as a single scalar index.
	IndexOfRefraction *ParamOrSidFloat

	//	Declares the amount of ambient light reflected from the surface of this object.
	//	Ignored if Kind is FxTechniqueKindConstant.
	Ambient *FxColorOrTexture

	//	Declares the amount of light diffusely reflected from the surface of this object.
	//	Ignored if Kind is FxTechniqueKindConstant.
	Diffuse *FxColorOrTexture

	//	Declares the color of light specularly reflected from the surface of this object.
	//	Ignored if Kind is FxTechniqueKindConstant or FxTechniqueKindLambert.
	Specular *FxColorOrTexture

	//	Declares the specularity or roughness of the specular reflection lobe.
	//	Ignored if Kind is FxTechniqueKindConstant or FxTechniqueKindLambert.
	Shininess *ParamOrSidFloat
}
```

Holds a description of the textures, samplers, shaders, parameters, and passes
necessary for rendering this effect within an FxProfileCommon.

#### type FxTechniqueGlsl

```go
type FxTechniqueGlsl struct {
	//	Id, Sid, Asset, Extras
	FxTechnique

	//	Application-specific FX metadata
	Annotations []*FxAnnotation

	//	Static declarations of all the render states, shaders, and settings for the rendering pipeline.
	Passes []*FxPass
}
```

Holds a description of the textures, samplers, shaders, parameters, and passes
necessary for rendering this effect within an FxProfileGlsl.

#### type FxTechniqueKind

```go
type FxTechniqueKind int
```

Categorizes the Kind of an FxTechniqueCommon.

```go
const (
	//	Produces a constantly shaded surface that is independent of lighting.
	FxTechniqueKindConstant FxTechniqueKind = iota

	//	Produces a constantly shaded surface that is independent of lighting.
	FxTechniqueKindLambert

	//	Produces a shaded surface with a Blinn BRDF approximation.
	FxTechniqueKindBlinn

	//	Produces a shaded surface with a Phong BRDF approximation.
	FxTechniqueKindPhong
)
```

#### type FxTexture

```go
type FxTexture struct {
	//	Extras
	HasExtras

	//	References a previously defined FxSampler of Kind FxSamplerKind2D.
	Sampler2D RefParam

	//	A semantic token, which will be referenced within FxMaterialBinding
	//	to bind an array of texture-coordinates from a geometry instance to the sampler.
	TexCoord string
}
```

Used in FxColorOrTexture instances that refer to a texture image instead of a
literal color value.

#### type FxTextureOpaque

```go
type FxTextureOpaque int
```

Categorizes transparency access in an FxColorOrTexture.

```go
const (
	//	Takes the transparency information from the color's alpha channel,
	//	where the value 1.0 is opaque. This is the default.
	FxTextureOpaqueA1 FxTextureOpaque = iota

	//	Takes the transparency information from the color's alpha channel,
	//	where the value 0.0 is opaque.
	FxTextureOpaqueA0

	//	Takes the transparency information from the color's red, green, and blue channels,
	//	where the value 0.0 is opaque, with each channel modulated independently.
	FxTextureOpaqueRgb0

	//	Takes the transparency information from the color's red, green, and blue channels,
	//	where the value 1.0 is opaque, with each channel modulated independently.
	FxTextureOpaqueRgb1
)
```

#### type FxVertexInputBinding

```go
type FxVertexInputBinding struct {
	//	Which effect parameter to bind.
	Semantic string

	//	Which input semantic to bind.
	InputSemantic string

	//	Which input set to bind. Optional.
	InputSet *uint64
}
```

Binds geometry vertex inputs to effect vertex inputs upon instantiation.

#### type FxWrapKind

```go
type FxWrapKind int
```

Categorizes the kinds of wrapping used in FxSamplerWrapping.

```go
const (
	//	Ignores the integer part of texture coordinates, using only the fractional part and tiling the
	//	texture at every integer junction. For example, for u values between 0 and 3, the texture is
	//	repeated three times; no mirroring is performed.
	FxWrapKindRepeat FxWrapKind = 0x2901

	//	First mirrors the texture coordinate. The mirrored coordinate is then clamped as described for
	//	FxWrapKindClamp. Flips the texture at every integer junction. For u values between 0 and 1,
	//	for example, the texture is addressed normally; between 1 and 2, the texture is flipped (mirrored);
	//	between 2 and 3, the texture is normal again; and so on.
	FxWrapKindMirror FxWrapKind = 0x8370

	//	Clamps texture coordinates at all MIPmap levels such that
	//	the texture filter never samples a border texel.
	FxWrapKindClamp FxWrapKind = 0x812F

	//	Clamps texture coordinates at all MIPmaps such that the texture filter always samples border
	//	texels for fragments whose corresponding texture coordinate is sufficiently far outside
	//	the range [0, 1]. Much like FxWrapKindClamp, except texture coordinates outside
	//	the range [0.0, 1.0] are set to the border color.
	FxWrapKindBorder FxWrapKind = 0x812D

	//	Takes the absolute value of the texture coordinate (thus, mirroring around 0),
	//	and then clamps to the maximum value.
	FxWrapKindMirrorOnce FxWrapKind = 41
)
```

#### type GeometryBrep

```go
type GeometryBrep struct {
	//	Extras
	HasExtras

	//	Sources
	HasSources

	//	Describes all vertices of the B-rep.
	//	Vertices are the base topological entity for all B-rep structures.
	Vertices GeometryVertices

	//	Contains all curves used in this B-rep. Required if Edges are present.
	Curves *GeometryBrepCurves

	//	Contains all 2D curves used in this B-rep.
	//	This includes surfaces that describe the kind of the face. Required if Faces are present.
	SurfaceCurves *GeometryBrepSurfaceCurves

	//	Contains all surfaces used in this B-rep.
	Surfaces *GeometryBrepSurfaces

	//	Describes all edges of the B-rep.
	Edges *GeometryBrepEdges

	//	Describes all wires of the B-rep.
	Wires *GeometryBrepWires

	//	Describes all faces of the B-rep.
	Faces *GeometryBrepFaces

	//	Describes all pcurves of the B-rep.
	Pcurves *GeometryBrepPcurves

	//	Describes all shells of the B-rep.
	Shells *GeometryBrepShells

	//	Describes all solids of the B-rep.
	Solids *GeometryBrepSolids
}
```

Describes a boundary representation (B-rep) structure.

#### func  NewGeometryBrep

```go
func NewGeometryBrep() (me *GeometryBrep)
```
Constructor

#### type GeometryBrepBox

```go
type GeometryBrepBox struct {
	//	Extras
	HasExtras

	//	Represents the extents of the box. The dimensions of the box are double the half-extents.
	HalfExtents unum.Vec3
}
```

Declares an axis-aligned box that is centered around its local origin.

#### type GeometryBrepCapsule

```go
type GeometryBrepCapsule struct {
	//	Extras
	HasExtras

	//	The length of the line segment connecting the centers of the capping hemispheres (ellipsoids).
	Height float64

	//	The x, y, and z radii of the capsule (it may be elliptical).
	Radii unum.Vec3
}
```

A capsule is a cylinder with rounded caps.

#### type GeometryBrepCircle

```go
type GeometryBrepCircle struct {
	//	Extras
	HasExtras

	//	The radius of the circle.
	Radius float64
}
```

Describes a circle in 3D space.

#### type GeometryBrepCone

```go
type GeometryBrepCone struct {
	//	Extras
	HasExtras

	//	The conical surface semi-angle.
	Angle float64

	//	Radius of the cone.
	Radius float64
}
```

Describes a conical surface. A cone is defined by the half-angle at its apex.

#### type GeometryBrepCurve

```go
type GeometryBrepCurve struct {
	//	Sid
	HasSid

	//	Name
	HasName

	//	Optional positioning of this surface to its correct location.
	Location GeometryPositioning

	//	The curve element. At least and at most one of these fields must be set (non-nil).
	Element struct {
		//	If set, curve element is a line.
		Line *GeometryBrepLine

		//	If set, curve element is a circle.
		Circle *GeometryBrepCircle

		//	If set, curve element is an ellipse.
		Ellipse *GeometryBrepEllipse

		//	If set, curve element is a parabola.
		Parabola *GeometryBrepParabola

		//	If set, curve element is a hyperbola.
		Hyperbola *GeometryBrepHyperbola

		//	If set, curve element is a NURBS curve.
		Nurbs *GeometryBrepNurbs
	}
}
```

Describes a specific curve.

#### type GeometryBrepCurves

```go
type GeometryBrepCurves struct {
	//	Extras
	HasExtras

	//	A container for all 3D curves used by the edges of the parent B-rep structure.
	All []*GeometryBrepCurve
}
```

Contains all curves that are used in the parent B-rep structure.

#### type GeometryBrepCylinder

```go
type GeometryBrepCylinder struct {
	//	Extras
	HasExtras

	//	The first value is the major radius, the second is the minor radius (cylinder may be elliptical).
	Radii Float2
}
```

Describes an unlimited cylindrical surface. An unlimited cylinder has a radius
but is assumed to extend to an infinite length.

#### type GeometryBrepEdges

```go
type GeometryBrepEdges struct {
	//	Id
	HasId

	//	Name
	HasName

	//	Extras
	HasExtras

	//	Four inputs are needed to define an edge:
	//	One with Semantic "CURVE" to reference the corresponding geometric element for the edge.
	//	Two with Semantic "VERTEX" to reference the two vertices that limit each edge.
	//	One with Semantic "PARAM" to set the parametric values (start and end parameters) of the curve.
	IndexedInputs
}
```

Describes the edges of a B-rep structure.

#### type GeometryBrepEllipse

```go
type GeometryBrepEllipse struct {
	//	Extras
	HasExtras

	//	The first value is the major radius, the second is the minor radius.
	Radii Float2
}
```

Describes an ellipse in 3D space.

#### type GeometryBrepFaces

```go
type GeometryBrepFaces struct {
	//	Id
	HasId

	//	Name
	HasName

	//	Extras
	HasExtras

	//	There must be at least three inputs:
	//	One with Semantic "SURFACE" to reference the corresponding geometric element for the face.
	//	One with Semantic "WIRE" to reference the wires for each face.
	//	One with Semantic "ORIENTATION" defining the orientation of the referenced wire within the face.
	IndexedInputs
}
```

Describes the faces of a B-rep structure.

#### type GeometryBrepHyperbola

```go
type GeometryBrepHyperbola struct {
	//	Extras
	HasExtras

	//	The first value is the major radius, the second is the minor radius.
	Radii Float2
}
```

Describes a hyperbola in 3D space.

#### type GeometryBrepLine

```go
type GeometryBrepLine struct {
	//	Extras
	HasExtras

	//	The origin of the line.
	Origin unum.Vec3

	//	The direction of the line as a unit vector.
	Direction unum.Vec3
}
```

Describes a single line in 3D space.

#### type GeometryBrepNurbs

```go
type GeometryBrepNurbs struct {
	//	Extras
	HasExtras

	//	Sources
	HasSources

	//	Specifies the degree of the NURBS curve.
	Degree uint64

	//	Specifies whether this NURBS curve is closed.
	Closed bool

	//	Control vertices for curve interpolation.
	ControlVertices GeometryControlVertices
}
```

Describes a NURBS curve in 3D space.

#### func  NewGeometryBrepNurbs

```go
func NewGeometryBrepNurbs() (me *GeometryBrepNurbs)
```
Constructor

#### type GeometryBrepNurbsSurface

```go
type GeometryBrepNurbsSurface struct {
	//	Extras
	HasExtras

	//	Sources
	HasSources

	//	The u and v directions for the NURBS curve.
	U, V struct {
		//	Specifies the degree of the NURBS curve for this direction.
		Degree uint64

		//	Specifies whether a NURBS curve is closed for this direction.
		Closed bool
	}
	//	Control vertices for curve interpolation.
	ControlVertices GeometryControlVertices
}
```

Describes a NURBS surface in 3D space.

#### func  NewGeometryBrepNurbsSurface

```go
func NewGeometryBrepNurbsSurface() (me *GeometryBrepNurbsSurface)
```
Constructor

#### type GeometryBrepOrientation

```go
type GeometryBrepOrientation struct {
	//	The axis of rotation.
	Axis unum.Vec3

	//	The rotational angle in degrees.
	Angle float64
}
```

Describes the orientation of an object frame.

#### type GeometryBrepParabola

```go
type GeometryBrepParabola struct {
	//	Extras
	HasExtras

	//	The distance between the parabola's focus and its apex.
	FocalLength float64
}
```

Describes a parabola in 3D space.

#### type GeometryBrepPcurves

```go
type GeometryBrepPcurves struct {
	//	Id
	HasId

	//	Name
	HasName

	//	Extras
	HasExtras

	//	There must be at least three inputs:
	//	One with Semantic "CURVE2D" referencing a pcurve.
	//	One with Semantic "FACE" and one with Semantic "EDGE"
	//	to specify the connection between the edge and the face.
	IndexedInputs
}
```

Specifies how an edge is represented in a face's parametric space.

#### type GeometryBrepPlane

```go
type GeometryBrepPlane struct {
	//	Extras
	HasExtras

	//	The four coefficients for the plane's equation: Ax + By + Cz + D = 0
	Equation Float4
}
```

Defines an infinite plane.

#### type GeometryBrepShells

```go
type GeometryBrepShells struct {
	//	Id
	HasId

	//	Name
	HasName

	//	Extras
	HasExtras

	//	There must be at least two inputs:
	//	One with Semantic "FACE" to reference the faces for each shell.
	//	One with Semantic "ORIENTATION" defining the orientation of the referenced face within the shell.
	IndexedInputs
}
```

Describes the shells of a B-rep structure. A shell is the union of one or more
faces. A closed shell can limit a solid.

#### type GeometryBrepSolids

```go
type GeometryBrepSolids struct {
	//	Id
	HasId

	//	Name
	HasName

	//	Extras
	HasExtras

	//	There must be at least two inputs:
	//	One with Semantic "SHELL" to reference the shells for each solid.
	//	One with Semantic "ORIENTATION" defining the orientation of the referenced shell within the solid.
	IndexedInputs
}
```

Describes the solids of a B-rep structure.

#### type GeometryBrepSphere

```go
type GeometryBrepSphere struct {
	//	Extras
	HasExtras

	//	The radius of this sphere.
	Radius float64
}
```

Describes a perfectly round sphere that is centered around its local origin.

#### type GeometryBrepSurface

```go
type GeometryBrepSurface struct {
	//	Sid
	HasSid

	//	Name
	HasName

	//	Optional positioning of this surface to its correct location.
	Location GeometryPositioning

	//	The surface element. At least and at most one of these fields must be set (non-nil).
	Element struct {
		//	Surface is described by a cone.
		Cone *GeometryBrepCone

		//	Surface is described by a plane.
		Plane *GeometryBrepPlane

		//	Surface is described by a cylinder.
		Cylinder *GeometryBrepCylinder

		//	Surface is described by a NURBS surface.
		NurbsSurface *GeometryBrepNurbsSurface

		//	Surface is described by a sphere.
		Sphere *GeometryBrepSphere

		//	Surface is described by a torus.
		Torus *GeometryBrepTorus

		//	Surface is described by an extruded or revolved curve.
		SweptSurface *GeometryBrepSweptSurface
	}
}
```

Describes a specific surface.

#### type GeometryBrepSurfaceCurves

```go
type GeometryBrepSurfaceCurves struct {
	//	Extras
	HasExtras

	//	Pcurves are curves in the parametric space of the surface on which they lie.
	All []*GeometryBrepCurve
}
```

Contains all parametric curves (pcurves) that are used in a B-rep structure.

#### type GeometryBrepSurfaces

```go
type GeometryBrepSurfaces struct {
	//	Extras
	HasExtras

	//	A container for all surfaces used by the faces of the parent B-rep structure.
	All []*GeometryBrepSurface
}
```

Contains all surfaces that are used in a B-rep structure.

#### type GeometryBrepSweptSurface

```go
type GeometryBrepSweptSurface struct {
	//	Extras
	HasExtras

	//	Describes the base curve being extruded or revolved.
	Curve *GeometryBrepCurve

	//	If Direction is set (non-nil), Revolution is ignored and this surface extrudes Curve.
	Extrusion struct {
		//	The direction of this curve extrusion.
		Direction *unum.Vec3
	}
	//	Only used if Extrusion.Direction is nil; then this surface revolves Curve.
	Revolution struct {
		//	The origin of the axis for revolution.
		Origin *unum.Vec3

		//	The axis' direction for revolution.
		Direction *unum.Vec3
	}
}
```

Describes a surface by extruding or revolving a curve.

#### func (*GeometryBrepSweptSurface) IsExtrusion

```go
func (me *GeometryBrepSweptSurface) IsExtrusion() bool
```
Returns true if this surface is described by extruding a curve.

#### func (*GeometryBrepSweptSurface) IsRevolution

```go
func (me *GeometryBrepSweptSurface) IsRevolution() bool
```
Returns true if this surface is described by revolving a curve.

#### type GeometryBrepTorus

```go
type GeometryBrepTorus struct {
	//	Extras
	HasExtras

	//	The first value is the major radius, the second is the minor radius.
	Radii Float2
}
```

Describes a torus in 3D space.

#### type GeometryBrepWires

```go
type GeometryBrepWires struct {
	//	Id
	HasId

	//	Name
	HasName

	//	Extras
	HasExtras

	//	There must be at least inputs:
	//	One with Semantic "EDGE" to reference the edges for each wire.
	//	One with Semantic "ORIENTATION" defining the orientation of the referenced edge within the wire.
	IndexedInputs
}
```

Describes the wires of a B-rep structure.

#### type GeometryControlVertices

```go
type GeometryControlVertices struct {
	//	Extras
	HasExtras

	//	Inputs
	HasInputs
}
```

Describes the control vertices of a spline.

#### type GeometryDef

```go
type GeometryDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	If set, Mesh and Spline must be nil, and the GeometryDef is described by this B-rep structure.
	Brep *GeometryBrep

	//	If set, Brep and Spline must be nil, and the GeometryDef is described by this mesh structure.
	Mesh *GeometryMesh

	//	If set, Mesh and Brep must be nil, and the GeometryDef is described by this multi-segment spline.
	Spline *GeometrySpline
}
```

Describes the visual shape and appearance of an object in a scene.

#### func (*GeometryDef) DefaultInst

```go
func (me *GeometryDef) DefaultInst() (inst *GeometryInst)
```
Returns "the default GeometryInst instance" referencing this GeometryDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*GeometryDef) Init

```go
func (me *GeometryDef) Init()
```
Initialization

#### func (*GeometryDef) NewInst

```go
func (me *GeometryDef) NewInst() (inst *GeometryInst)
```
Creates and returns a new GeometryInst instance referencing this GeometryDef
definition. Any GeometryInst created by this method will have its Def field
readily set to me.

#### type GeometryInst

```go
type GeometryInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *GeometryDef

	//	Binds material symbols to material instances. This allows a single geometry
	//	to be instantiated into a scene multiple times each with a different appearance.
	MaterialBinding *MaterialBinding
}
```

Instantiates a geometry resource.

#### func (*GeometryInst) EnsureDef

```go
func (me *GeometryInst) EnsureDef() *GeometryDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct GeometryDef
according to the current me.DefRef value (by searching AllGeometryDefLibs). Then
returns me.Def. (Note, every GeometryInst's Def is nil initially, unless it was
created via GeometryDef.NewInst().)

#### func (*GeometryInst) Init

```go
func (me *GeometryInst) Init()
```
Initialization

#### type GeometryMesh

```go
type GeometryMesh struct {
	//	Extras
	HasExtras

	//	Sources
	HasSources

	//	Refers to a GeometryDef described by a GeometryMesh.
	//	If specified, compute the convex hull of the specified mesh.
	ConvexHullOf RefId

	//	Describes the mesh-vertex attributes and establishes their topological identity.
	//	Required if ConvexHullOf is empty.
	Vertices *GeometryVertices

	//	Geometric primitives, which assemble values from the inputs into vertex attribute data.
	Primitives []*GeometryPrimitives
}
```

Describes basic geometric meshes using vertex and primitive information.

#### func  NewGeometryMesh

```go
func NewGeometryMesh() (me *GeometryMesh)
```
Constructor

#### type GeometryPolygonHole

```go
type GeometryPolygonHole struct {
	//	Specifies the vertex attributes (indices) for an individual polygon.
	Indices []uint64

	//	Specifies the indices of a hole in the polygon specified by Indices.
	Holes [][]uint64
}
```

Describes a polygon that contains one or more holes.

#### type GeometryPositioning

```go
type GeometryPositioning struct {
	//	If set, describes the origin of the object frame.
	Origin *unum.Vec3

	//	If set, these describe the orientation of the object frame.
	Orientations []*GeometryBrepOrientation
}
```

Used to position a surface or curve to its correct location.

#### type GeometryPrimitiveKind

```go
type GeometryPrimitiveKind int
```

Categorizes the kind of a GeometryPrimitives.

```go
const (
	//	Organizes vertices into individual lines.
	GeometryPrimitiveKindLines GeometryPrimitiveKind = 0x0001

	//	Organizes vertices into connected line-strips.
	GeometryPrimitiveKindLineStrips GeometryPrimitiveKind = 0x0003

	//	Organizes vertices into individual polygons that may contain holes.
	GeometryPrimitiveKindPolygons GeometryPrimitiveKind = 2

	//	Organizes vertices into individual polygons that cannot contain holes.
	GeometryPrimitiveKindPolylist GeometryPrimitiveKind = 7

	//	Organizes vertices into individual triangles.
	GeometryPrimitiveKindTriangles GeometryPrimitiveKind = 0x0004

	//	Organizes vertices into fan-connected triangles.
	GeometryPrimitiveKindTrifans GeometryPrimitiveKind = 0x0006

	//	Organizes vertices into strip-connected triangles.
	GeometryPrimitiveKindTristrips GeometryPrimitiveKind = 0x0005
)
```

#### type GeometryPrimitives

```go
type GeometryPrimitives struct {
	//	Extras
	HasExtras

	//	Name
	HasName

	//	When at least one input is present, one input must specify its Semantic as "VERTEX".
	IndexedInputs

	//	Must be one of the GeometryPrimitiveKind* enumerated constants.
	Kind GeometryPrimitiveKind

	//	Declares a symbol for a material. This symbol is bound to a material at the time of instantiation.
	//	Optional. If not specified then the lighting and shading results are application defined.
	Material string

	//	If Kind is GeometryPrimitiveKindPolygons, describes the polygons that contain one or more holes.
	PolyHoles []*GeometryPolygonHole
}
```

Geometric primitives, which assemble values from inputs into vertex attribute
data.

#### type GeometrySpline

```go
type GeometrySpline struct {
	//	Extras
	HasExtras

	//	Sources
	HasSources

	//	Whether there is a segment connecting the first and last control vertices.
	//	The default is false, indicating that the spline is open.
	Closed bool

	//	Describes the control vertices of the spline.
	ControlVertices GeometryControlVertices
}
```

Describes a multisegment spline with control vertex and segment information.

#### func  NewGeometrySpline

```go
func NewGeometrySpline() (me *GeometrySpline)
```
Constructor

#### type GeometryVertices

```go
type GeometryVertices struct {
	//	Id
	HasId

	//	Name
	HasName

	//	Extras
	HasExtras

	//	Inputs
	HasInputs
}
```

Declares the attributes and identity of mesh-vertices. The mesh-vertices
represent the position (identity) of the vertices comprising the mesh and other
vertex attributes that are invariant to tessellation.

#### type HasAsset

```go
type HasAsset struct {
	//	Resource-specific asset-management information and meta-data.
	Asset *Asset
}
```

Used in all resources that require asset-management information.

#### type HasExtras

```go
type HasExtras struct {
	//	Custom-technique/foreign-profile meta-data.
	Extras []*Extra
}
```

Used in all resources that support custom techniques / foreign profiles.

#### type HasFxParamDefs

```go
type HasFxParamDefs struct {
	//	A hash-table containing all parameter declarations of this resource.
	NewParams FxParamDefs
}
```

Used in all FX resources that declare their own parameters.

#### type HasId

```go
type HasId struct {
	//	The unique identifier of this resource.
	Id string
}
```

Used in all resources that declare their own unique identifier.

#### type HasInputs

```go
type HasInputs struct {
	//	Declares the input semantics of a data Source and connects a consumer to that Source.
	Inputs []*Input
}
```

Used in all data consumers that require input connections into a data Source.

#### type HasName

```go
type HasName struct {
	//	The optional pretty-print name/title of this resource.
	Name string
}
```

Used in all resources that support arbitrary pretty-print names/titles.

#### type HasParamDefs

```go
type HasParamDefs struct {
	//	A hash-table containing all parameter declarations of this resource.
	NewParams ParamDefs
}
```

Used in all resources that declare their own parameters.

#### type HasParamInsts

```go
type HasParamInsts struct {
	//	A hash-table containing all parameter values assigned by this resource.
	SetParams ParamInsts
}
```

Used in all resources that assign values to other parameters.

#### type HasSid

```go
type HasSid struct {
	//	The Scoped identifier of this resource.
	Sid string
}
```

Used in all resources that declare their own scoped identifier.

#### type HasSources

```go
type HasSources struct {
	//	Provides the bulk of this resource's data.
	Sources Sources
}
```

Used in all resources that provide data arrays.

#### type HasTechniques

```go
type HasTechniques struct {
	//	Custom-technique/foreign-profile content or data.
	Techniques []*Technique
}
```

Used in all resources that support custom techniques / foreign profiles.

#### type IndexedInputs

```go
type IndexedInputs struct {
	//	Number of primitives
	Count uint64

	//	Inputs specify how to read data from Sources.
	Inputs []*InputShared

	//	Indices that describe the attributes for a number of primitives.
	//	The indices reference into the Sources that are referenced by the Inputs.
	Indices []uint64

	//	Number of sub-primitives, if used.
	Vcount []int64
}
```

Used in various geometry primitives and b-rep resources.

#### type Input

```go
type Input struct {
	//	The user-defined meaning of the input connection.
	Semantic string

	//	Refers to the Source for this Input.
	Source RefId
}
```

Declares unshared input semantics of a data source and connects a consumer to
that source.

#### type InputShared

```go
type InputShared struct {
	//	Semantic and Source
	Input

	//	The offset into the list of indices.
	Offset uint64

	//	Which inputs to group as a single set.
	//	This is helpful when multiple inputs share the same semantics.
	Set *uint64
}
```

Declares shared input semantics of a data source and connects a consumer to that
source.

#### type Int2

```go
type Int2 [2]int64
```

Contains two int64 values.

#### func (*Int2) AccessIndex

```go
func (me *Int2) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Int2x2

```go
type Int2x2 [4]int64
```

Contains four int64 values.

#### func (*Int2x2) AccessIndex

```go
func (me *Int2x2) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Int3

```go
type Int3 [3]int64
```

Contains three int64 values.

#### func (*Int3) AccessIndex

```go
func (me *Int3) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Int3x3

```go
type Int3x3 [9]int64
```

Contains nine int64 values.

#### func (*Int3x3) AccessIndex

```go
func (me *Int3x3) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Int4

```go
type Int4 [4]int64
```

Contains four int64 values.

#### func (*Int4) AccessIndex

```go
func (me *Int4) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type Int4x4

```go
type Int4x4 [16]int64
```

Contains sixteen int64 values.

#### func (*Int4x4) AccessIndex

```go
func (me *Int4x4) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type KxArticulatedSystemDef

```go
type KxArticulatedSystemDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	If set, Motion must be nil, and this articulated system describes a kinematics system.
	Kinematics *KxKinematicsSystem

	//	If set, Kinematics must be nil, and this articulated system describes a motion system.
	Motion *KxMotionSystem
}
```

Categorizes the declaration of generic control information for kinematics
systems.

#### func (*KxArticulatedSystemDef) DefaultInst

```go
func (me *KxArticulatedSystemDef) DefaultInst() (inst *KxArticulatedSystemInst)
```
Returns "the default KxArticulatedSystemInst instance" referencing this
KxArticulatedSystemDef definition. That instance is created once when this
method is first called on me, and will have its Def field readily set to me.

#### func (*KxArticulatedSystemDef) Init

```go
func (me *KxArticulatedSystemDef) Init()
```
Initialization

#### func (*KxArticulatedSystemDef) NewInst

```go
func (me *KxArticulatedSystemDef) NewInst() (inst *KxArticulatedSystemInst)
```
Creates and returns a new KxArticulatedSystemInst instance referencing this
KxArticulatedSystemDef definition. Any KxArticulatedSystemInst created by this
method will have its Def field readily set to me.

#### type KxArticulatedSystemInst

```go
type KxArticulatedSystemInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	NewParams
	HasParamDefs

	//	SetParams
	HasParamInsts

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *KxArticulatedSystemDef

	//	Bindings of inputs to kinematics parameters.
	Bindings []*KxBinding
}
```

Instantiates a kinematics articulated system resource.

#### func (*KxArticulatedSystemInst) EnsureDef

```go
func (me *KxArticulatedSystemInst) EnsureDef() *KxArticulatedSystemDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct
KxArticulatedSystemDef according to the current me.DefRef value (by searching
AllKxArticulatedSystemDefLibs). Then returns me.Def. (Note, every
KxArticulatedSystemInst's Def is nil initially, unless it was created via
KxArticulatedSystemDef.NewInst().)

#### func (*KxArticulatedSystemInst) Init

```go
func (me *KxArticulatedSystemInst) Init()
```
Initialization

#### type KxAttachment

```go
type KxAttachment struct {
	//	Must be one of the KxAttachmentKind* enumerated constants.
	Kind KxAttachmentKind

	//	Refers to the KxJoint that connects the parent with the child link. Required.
	Joint RefSid

	//	Zero or more TransformKindRotate and/or TransformKindTranslate transformations.
	Transforms []*Transform

	//	If Kind is KxAttachmentKindFull, specifies the child link in this parent-child dependency.
	Link *KxLink
}
```

Connects links or define ends of closed loops.

#### type KxAttachmentKind

```go
type KxAttachmentKind int
```

Categorizes the kind of a KxAttachment.

```go
const (
	//	Connects two links, describing a real parent-child dependency between them.
	KxAttachmentKindFull KxAttachmentKind = iota + 1

	//	Connects two links and defines one end of a closed loop.
	KxAttachmentKindStart

	//	Defines one end of the closed loop in an attachment.
	KxAttachmentKindEnd
)
```

#### type KxAxisIndex

```go
type KxAxisIndex struct {
	//	If set, specifies the special use of this index.
	Semantic string

	//	If not set, the parent axis will not appear in the jointmap.
	I ParamOrInt
}
```

Specifies the parent axis' index in the jointmap.

#### type KxAxisLimits

```go
type KxAxisLimits struct {
	//	The "minimum" portion of this limits descriptor.
	Min ParamOrFloat

	//	The "maximum" portion of this limits descriptor.
	Max ParamOrFloat
}
```

Specifies the parent axis' soft limits.

#### type KxBinding

```go
type KxBinding struct {
	//	The identifier of the parameter to bind to the new symbol name. Required.
	Symbol string

	//	If set, Value is ignored.
	Param RefParam

	//	Only used if Param is empty.
	Value interface{}
}
```

Binds inputs to kinematics parameters upon instantiation.

#### type KxEffector

```go
type KxEffector struct {
	//	Sid
	HasSid

	//	Name
	HasName

	//	NewParams
	HasParamDefs

	//	SetParams
	HasParamInsts

	//	Bindings of inputs to kinematics parameters.
	Bindings []*KxBinding

	//	Specifies maximum speed.
	//	The first value is translational (m/sec), the second is rotational (°/sec).
	Speed *ParamOrFloat2

	//	Specifies maximum acceleration.
	//	The first value is translational (m/sec²), the second is rotational (°/sec²).
	Acceleration *ParamOrFloat2

	//	Specifies the maximum deceleration.
	//	The first value is translational (m/sec²), the second is rotational (°/sec²).
	Deceleration *ParamOrFloat2

	//	Specifies the maximum jerk (also called jolt or surge).
	//	The first value is translational (m/sec³), the second is rotational (°/sec³).
	Jerk *ParamOrFloat2
}
```

Specifies additional dynamics information for an effector.

#### func  NewKxEffector

```go
func NewKxEffector() (me *KxEffector)
```
Constructor

#### type KxFrame

```go
type KxFrame struct {
	//	References a KxLink defined in the kinematics model. Optional.
	Link RefSid

	//	Zero or more TransformKindTranslate and/or TransformKindRotate transformations.
	Transforms []*Transform
}
```

Contains information for a frame used for kinematics calculation.

#### type KxFrameObject

```go
type KxFrameObject struct {
	//	Link, Transforms
	KxFrame
}
```

Defines the offset frame from the KxFrameOrigin; this offset usually represents
the transformation to a work piece.

#### type KxFrameOrigin

```go
type KxFrameOrigin struct {
	//	Link, Transforms
	KxFrame
}
```

Defines the base frame for kinematics calculation.

#### type KxFrameTcp

```go
type KxFrameTcp struct {
	//	Link, Transforms
	KxFrame
}
```

Defines the offset frame from the KxFrameTip, which usually represents the work
point of the end effector (for example, a welding gun).

#### type KxFrameTip

```go
type KxFrameTip struct {
	//	Link, Transforms
	KxFrame
}
```

Defines the frame at the end of the kinematics chain.

#### type KxJoint

```go
type KxJoint struct {
	//	Sid
	HasSid

	//	Must be one of the KxJointKind* enumerated constants.
	Kind KxJointKind

	//	Specifies the axis of the degree of freedom.
	Axis struct {
		//	Name
		HasName

		//	Sid, V
		SidVec3
	}
	//	If set, these specified limits are physical limits.
	Limits *KxJointLimits
}
```

Primitive (simple) joints are joints with one degree of freedom (one given axis)
and are used to construct more complex joint types (compound joints) that
consist of multiple primitives, each representing an axis.

#### type KxJointAxisBinding

```go
type KxJointAxisBinding struct {
	//	A reference to a transformation of a node.
	Target RefSid

	//	If set, Value is ignored. Specifies an axis of a kinematics model.
	Axis ParamOrRefSid

	//	Only used if Axis is empty. Specifies a value of the axis.
	Value ParamOrFloat
}
```

Binds a joint axis of a kinematics model to a single transformation of a node.
By binding a joint axis to a transformation of a node, it is possible to
synchronize a kinematics scene with a visual scene.

#### type KxJointDef

```go
type KxJointDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Sid
	HasSid

	//	Primitive (simple) joints are joints with one degree of freedom (one given axis) and are
	//	used to construct more complex joint types (compound joints) that
	//	consist of multiple primitives, each representing an axis.
	All []*KxJoint
}
```

Defines a single complex/compound joint with one or more degrees of freedom.

#### func (*KxJointDef) DefaultInst

```go
func (me *KxJointDef) DefaultInst() (inst *KxJointInst)
```
Returns "the default KxJointInst instance" referencing this KxJointDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*KxJointDef) Init

```go
func (me *KxJointDef) Init()
```
Initialization

#### func (*KxJointDef) NewInst

```go
func (me *KxJointDef) NewInst() (inst *KxJointInst)
```
Creates and returns a new KxJointInst instance referencing this KxJointDef
definition. Any KxJointInst created by this method will have its Def field
readily set to me.

#### type KxJointInst

```go
type KxJointInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *KxJointDef
}
```

Instantiates a kinematics joint resource.

#### func (*KxJointInst) EnsureDef

```go
func (me *KxJointInst) EnsureDef() *KxJointDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct KxJointDef
according to the current me.DefRef value (by searching AllKxJointDefLibs). Then
returns me.Def. (Note, every KxJointInst's Def is nil initially, unless it was
created via KxJointDef.NewInst().)

#### func (*KxJointInst) Init

```go
func (me *KxJointInst) Init()
```
Initialization

#### type KxJointKind

```go
type KxJointKind int
```

Categorizes the kind of a KxJoint.

```go
const (
	//	Defines a single translational degree of freedom of a joint.
	KxJointKindPrismatic KxJointKind = iota + 1

	//	Defines a single rotational degree of freedom of a joint.
	KxJointKindRevolute
)
```

#### type KxJointLimits

```go
type KxJointLimits struct {
	//	If set, the "minimum" portion of this joint limitation.
	Min *SidFloat

	//	If set, the "maximum" portion of this joint limitation.
	Max *SidFloat
}
```

Declares a primitive/simple joint as fully limited (if Min and Max are both
set), partially limited (if either Min or Max is nil, but not both) or unlimited
(if Min and Max are nil).

#### type KxKinematicsAxis

```go
type KxKinematicsAxis struct {
	//	Sid
	HasSid

	//	Name
	HasName

	//	NewParams
	HasParamDefs

	//	The joint axis of an instantiated kinematics model.
	Axis RefSid

	//	Defaults to true.
	Active ParamOrBool

	//	Specifies this axis' indices in the jointmap. If empty, this axis will not appear in the jointmap.
	Indices []*KxAxisIndex

	//	Specifies the soft limits. If not set, the axis is limited only by its physical limits.
	Limits *KxAxisLimits

	//	Defaults to false.
	Locked ParamOrBool

	//	Formulas can be useful to define the behavior of a passive link according to one or more
	//	active axes, or to define dependencies of the soft limits and another joint, for example.
	Formulas []Formula
}
```

Contains axis information to describe the kinematics behavior of an articulated
model.

#### func  NewKxKinematicsAxis

```go
func NewKxKinematicsAxis() (me *KxKinematicsAxis)
```
Constructor

#### type KxKinematicsSystem

```go
type KxKinematicsSystem struct {
	//	Techniques
	HasTechniques

	//	The kinematics models to be enhanced with kinematics information.
	Models []*KxModelInst

	//	Common-technique profile
	TC struct {
		//	Kinematics-related information for all axes.
		AxisInfos []*KxKinematicsAxis

		//	Kinematics calculation chain frames
		Frame struct {
			//	Defines the base frame for kinematics calculation.
			Origin KxFrameOrigin

			//	Defines the frame at the end of the kinematics chain.
			Tip KxFrameTip

			//	If set, defines the offset frame from the Tip frame,
			//	which usually represents the work point of the end effector (for example, a welding gun).
			Tcp *KxFrameTcp

			//	If set, defines the offset frame from the Origin frame;
			//	this offset usually represents the transformation to a work piece.
			Object *KxFrameObject
		}
	}
}
```

Contains additional information to describe the kinematical behavior of an
articulated model.

#### type KxLink

```go
type KxLink struct {
	//	Sid
	HasSid

	//	Name
	HasName

	//	Zero or more TransformKindRotate and/or TransformKindTranslate transformations.
	Transforms []*Transform

	//	The attachments that make up this link.
	Attachments []*KxAttachment
}
```

Represents a rigid kinematical object without mass whose motion is constrained
by one or more joints.

#### type KxModelBinding

```go
type KxModelBinding struct {
	//	A reference to a node.
	Node RefId

	//	Refers to the kinematics model being bound.
	//	Only either SidRef or ParamRef, but not both, must be specified.
	Model struct {
		//	If set, ParamRef must be empty.
		//	The Sid path to the kinematics model to bind to the node.
		SidRef RefSid

		//	If set, SidRef must be empty.
		//	The parameter of the kinematics model that is defined in the instantiated kinematics scene.
		ParamRef RefParam
	}
}
```

Binds a kinematics model to a node. The description of a kinematics model is
completely independent of any visual information, but for calculation the
position is important.

#### type KxModelDef

```go
type KxModelDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Techniques
	HasTechniques

	//	Common-technique profile
	TC struct {
		//	NewParams
		HasParamDefs

		//	The kinematics chain.
		Links []*KxLink

		//	Specifies dependencies among the joints.
		Formulas []Formula
	}
}
```

Categorizes the declaration of kinematical information, containing declarations
of joints, links, and attachment points. A kinematics model is focused on strict
kinematics description "in zero position", without any additional physical
descriptions.

#### func (*KxModelDef) DefaultInst

```go
func (me *KxModelDef) DefaultInst() (inst *KxModelInst)
```
Returns "the default KxModelInst instance" referencing this KxModelDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*KxModelDef) Init

```go
func (me *KxModelDef) Init()
```
Initialization

#### func (*KxModelDef) NewInst

```go
func (me *KxModelDef) NewInst() (inst *KxModelInst)
```
Creates and returns a new KxModelInst instance referencing this KxModelDef
definition. Any KxModelInst created by this method will have its Def field
readily set to me.

#### type KxModelInst

```go
type KxModelInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	NewParams
	HasParamDefs

	//	SetParams
	HasParamInsts

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *KxModelDef

	//	Bindings of inputs to kinematics parameters.
	Bindings []*KxBinding
}
```

Instantiates a kinematics model resource.

#### func (*KxModelInst) EnsureDef

```go
func (me *KxModelInst) EnsureDef() *KxModelDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct KxModelDef
according to the current me.DefRef value (by searching AllKxModelDefLibs). Then
returns me.Def. (Note, every KxModelInst's Def is nil initially, unless it was
created via KxModelDef.NewInst().)

#### func (*KxModelInst) Init

```go
func (me *KxModelInst) Init()
```
Initialization

#### type KxMotionAxis

```go
type KxMotionAxis struct {
	//	Sid
	HasSid

	//	Name
	HasName

	//	NewParams
	HasParamDefs

	//	SetParams
	HasParamInsts

	//	References the KxKinematicsAxis of an instantiated kinematics system.
	Axis RefSid

	//	Bindings of inputs to kinematics parameters.
	Bindings []*KxBinding

	//	The maximum permitted speed of the axis in meters per second (m/sec).
	Speed *ParamOrFloat

	//	The maximum permitted acceleration of the axis in m/sec².
	Acceleration *ParamOrFloat

	//	The maximum permitted deceleration of an axis.
	//	If not set, acceleration and deceleration have the same value in m/sec².
	Deceleration *ParamOrFloat

	//	The maximum permitted jerk of an axis in m/sec³.
	Jerk *ParamOrFloat
}
```

Contains axis information to describe the motion behavior of an articulated
model.

#### func  NewKxMotionAxis

```go
func NewKxMotionAxis() (me *KxMotionAxis)
```
Constructor

#### type KxMotionSystem

```go
type KxMotionSystem struct {
	//	Techniques
	HasTechniques

	//	The articulated system to be enhanced with dynamics information.
	ArticulatedSystem *KxArticulatedSystemInst

	//	Common-technique profile
	TC struct {
		//	Dynamics-related information for all axes.
		AxisInfos []*KxMotionAxis

		//	Additional dynamics information
		EffectorInfo *KxEffector
	}
}
```

Contains additional information to describe the dynamics behaviour of an
articulated model.

#### type KxSceneDef

```go
type KxSceneDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Zero or more kinematics models participating in this kinematics scene.
	Models []*KxModelInst

	//	Zero or more articulated systems participating in this kinematics scene.
	ArticulatedSystems []*KxArticulatedSystemInst
}
```

Embodies the entire set of kinematics information that can be articulated from a
resource.

#### func (*KxSceneDef) DefaultInst

```go
func (me *KxSceneDef) DefaultInst() (inst *KxSceneInst)
```
Returns "the default KxSceneInst instance" referencing this KxSceneDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*KxSceneDef) Init

```go
func (me *KxSceneDef) Init()
```
Initialization

#### func (*KxSceneDef) NewInst

```go
func (me *KxSceneDef) NewInst() (inst *KxSceneInst)
```
Creates and returns a new KxSceneInst instance referencing this KxSceneDef
definition. Any KxSceneInst created by this method will have its Def field
readily set to me.

#### type KxSceneInst

```go
type KxSceneInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	NewParams
	HasParamDefs

	//	SetParams
	HasParamInsts

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *KxSceneDef

	//	Zero or more bindings of kinematics models to nodes.
	ModelBindings []*KxModelBinding

	//	Zero or more bindings of kinematics models' joint axes to single node transformations.
	JointAxisBindings []*KxJointAxisBinding
}
```

Instantiates a kinematics scene resource.

#### func (*KxSceneInst) EnsureDef

```go
func (me *KxSceneInst) EnsureDef() *KxSceneDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct KxSceneDef
according to the current me.DefRef value (by searching AllKxSceneDefLibs). Then
returns me.Def. (Note, every KxSceneInst's Def is nil initially, unless it was
created via KxSceneDef.NewInst().)

#### func (*KxSceneInst) Init

```go
func (me *KxSceneInst) Init()
```
Initialization

#### type Layers

```go
type Layers map[string]bool
```

Allows simple association of resources with custom named layers.

#### type LibAnimationClipDefs

```go
type LibAnimationClipDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*AnimationClipDef
}
```

A library that contains AnimationClipDefs associated by their Id. To create a
new LibAnimationClipDefs library, ONLY use the LibsAnimationClipDef.New() or
LibsAnimationClipDef.AddNew() methods.

#### func (*LibAnimationClipDefs) Add

```go
func (me *LibAnimationClipDefs) Add(d *AnimationClipDef) (n *AnimationClipDef)
```
Adds the specified AnimationClipDef definition to this LibAnimationClipDefs, and
returns it. If this LibAnimationClipDefs already contains a AnimationClipDef
definition with the same Id, does nothing and returns nil.

#### func (*LibAnimationClipDefs) AddNew

```go
func (me *LibAnimationClipDefs) AddNew(id string) *AnimationClipDef
```
Creates a new AnimationClipDef definition with the specified Id, adds it to this
LibAnimationClipDefs, and returns it. If this LibAnimationClipDefs already
contains a AnimationClipDef definition with the specified Id, does nothing and
returns nil.

#### func (*LibAnimationClipDefs) Len

```go
func (me *LibAnimationClipDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibAnimationClipDefs) New

```go
func (me *LibAnimationClipDefs) New(id string) (def *AnimationClipDef)
```
Creates a new AnimationClipDef definition with the specified Id and returns it,
but does not add it to this LibAnimationClipDefs.

#### func (*LibAnimationClipDefs) Remove

```go
func (me *LibAnimationClipDefs) Remove(id string)
```
Removes the AnimationClipDef with the specified Id from this
LibAnimationClipDefs.

#### func (*LibAnimationClipDefs) SyncChanges

```go
func (me *LibAnimationClipDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibAnimationClipDefs that need to be picked up. Call this after you have
made a number of changes to this LibAnimationClipDefs library or its
AnimationClipDef definitions. Also called by the global SyncChanges() function.

#### type LibAnimationDefs

```go
type LibAnimationDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*AnimationDef
}
```

A library that contains AnimationDefs associated by their Id. To create a new
LibAnimationDefs library, ONLY use the LibsAnimationDef.New() or
LibsAnimationDef.AddNew() methods.

#### func (*LibAnimationDefs) Add

```go
func (me *LibAnimationDefs) Add(d *AnimationDef) (n *AnimationDef)
```
Adds the specified AnimationDef definition to this LibAnimationDefs, and returns
it. If this LibAnimationDefs already contains a AnimationDef definition with the
same Id, does nothing and returns nil.

#### func (*LibAnimationDefs) AddNew

```go
func (me *LibAnimationDefs) AddNew(id string) *AnimationDef
```
Creates a new AnimationDef definition with the specified Id, adds it to this
LibAnimationDefs, and returns it. If this LibAnimationDefs already contains a
AnimationDef definition with the specified Id, does nothing and returns nil.

#### func (*LibAnimationDefs) Len

```go
func (me *LibAnimationDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibAnimationDefs) New

```go
func (me *LibAnimationDefs) New(id string) (def *AnimationDef)
```
Creates a new AnimationDef definition with the specified Id and returns it, but
does not add it to this LibAnimationDefs.

#### func (*LibAnimationDefs) Remove

```go
func (me *LibAnimationDefs) Remove(id string)
```
Removes the AnimationDef with the specified Id from this LibAnimationDefs.

#### func (*LibAnimationDefs) SyncChanges

```go
func (me *LibAnimationDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibAnimationDefs that need to be picked up. Call this after you have
made a number of changes to this LibAnimationDefs library or its AnimationDef
definitions. Also called by the global SyncChanges() function.

#### type LibCameraDefs

```go
type LibCameraDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*CameraDef
}
```

A library that contains CameraDefs associated by their Id. To create a new
LibCameraDefs library, ONLY use the LibsCameraDef.New() or
LibsCameraDef.AddNew() methods.

#### func (*LibCameraDefs) Add

```go
func (me *LibCameraDefs) Add(d *CameraDef) (n *CameraDef)
```
Adds the specified CameraDef definition to this LibCameraDefs, and returns it.
If this LibCameraDefs already contains a CameraDef definition with the same Id,
does nothing and returns nil.

#### func (*LibCameraDefs) AddNew

```go
func (me *LibCameraDefs) AddNew(id string) *CameraDef
```
Creates a new CameraDef definition with the specified Id, adds it to this
LibCameraDefs, and returns it. If this LibCameraDefs already contains a
CameraDef definition with the specified Id, does nothing and returns nil.

#### func (*LibCameraDefs) Len

```go
func (me *LibCameraDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibCameraDefs) New

```go
func (me *LibCameraDefs) New(id string) (def *CameraDef)
```
Creates a new CameraDef definition with the specified Id and returns it, but
does not add it to this LibCameraDefs.

#### func (*LibCameraDefs) Remove

```go
func (me *LibCameraDefs) Remove(id string)
```
Removes the CameraDef with the specified Id from this LibCameraDefs.

#### func (*LibCameraDefs) SyncChanges

```go
func (me *LibCameraDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibCameraDefs that need to be picked up. Call this after you have made a
number of changes to this LibCameraDefs library or its CameraDef definitions.
Also called by the global SyncChanges() function.

#### type LibControllerDefs

```go
type LibControllerDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*ControllerDef
}
```

A library that contains ControllerDefs associated by their Id. To create a new
LibControllerDefs library, ONLY use the LibsControllerDef.New() or
LibsControllerDef.AddNew() methods.

#### func (*LibControllerDefs) Add

```go
func (me *LibControllerDefs) Add(d *ControllerDef) (n *ControllerDef)
```
Adds the specified ControllerDef definition to this LibControllerDefs, and
returns it. If this LibControllerDefs already contains a ControllerDef
definition with the same Id, does nothing and returns nil.

#### func (*LibControllerDefs) AddNew

```go
func (me *LibControllerDefs) AddNew(id string) *ControllerDef
```
Creates a new ControllerDef definition with the specified Id, adds it to this
LibControllerDefs, and returns it. If this LibControllerDefs already contains a
ControllerDef definition with the specified Id, does nothing and returns nil.

#### func (*LibControllerDefs) Len

```go
func (me *LibControllerDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibControllerDefs) New

```go
func (me *LibControllerDefs) New(id string) (def *ControllerDef)
```
Creates a new ControllerDef definition with the specified Id and returns it, but
does not add it to this LibControllerDefs.

#### func (*LibControllerDefs) Remove

```go
func (me *LibControllerDefs) Remove(id string)
```
Removes the ControllerDef with the specified Id from this LibControllerDefs.

#### func (*LibControllerDefs) SyncChanges

```go
func (me *LibControllerDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibControllerDefs that need to be picked up. Call this after you have
made a number of changes to this LibControllerDefs library or its ControllerDef
definitions. Also called by the global SyncChanges() function.

#### type LibFormulaDefs

```go
type LibFormulaDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*FormulaDef
}
```

A library that contains FormulaDefs associated by their Id. To create a new
LibFormulaDefs library, ONLY use the LibsFormulaDef.New() or
LibsFormulaDef.AddNew() methods.

#### func (*LibFormulaDefs) Add

```go
func (me *LibFormulaDefs) Add(d *FormulaDef) (n *FormulaDef)
```
Adds the specified FormulaDef definition to this LibFormulaDefs, and returns it.
If this LibFormulaDefs already contains a FormulaDef definition with the same
Id, does nothing and returns nil.

#### func (*LibFormulaDefs) AddNew

```go
func (me *LibFormulaDefs) AddNew(id string) *FormulaDef
```
Creates a new FormulaDef definition with the specified Id, adds it to this
LibFormulaDefs, and returns it. If this LibFormulaDefs already contains a
FormulaDef definition with the specified Id, does nothing and returns nil.

#### func (*LibFormulaDefs) Len

```go
func (me *LibFormulaDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibFormulaDefs) New

```go
func (me *LibFormulaDefs) New(id string) (def *FormulaDef)
```
Creates a new FormulaDef definition with the specified Id and returns it, but
does not add it to this LibFormulaDefs.

#### func (*LibFormulaDefs) Remove

```go
func (me *LibFormulaDefs) Remove(id string)
```
Removes the FormulaDef with the specified Id from this LibFormulaDefs.

#### func (*LibFormulaDefs) SyncChanges

```go
func (me *LibFormulaDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibFormulaDefs that need to be picked up. Call this after you have made
a number of changes to this LibFormulaDefs library or its FormulaDef
definitions. Also called by the global SyncChanges() function.

#### type LibFxEffectDefs

```go
type LibFxEffectDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*FxEffectDef
}
```

A library that contains FxEffectDefs associated by their Id. To create a new
LibFxEffectDefs library, ONLY use the LibsFxEffectDef.New() or
LibsFxEffectDef.AddNew() methods.

#### func (*LibFxEffectDefs) Add

```go
func (me *LibFxEffectDefs) Add(d *FxEffectDef) (n *FxEffectDef)
```
Adds the specified FxEffectDef definition to this LibFxEffectDefs, and returns
it. If this LibFxEffectDefs already contains a FxEffectDef definition with the
same Id, does nothing and returns nil.

#### func (*LibFxEffectDefs) AddNew

```go
func (me *LibFxEffectDefs) AddNew(id string) *FxEffectDef
```
Creates a new FxEffectDef definition with the specified Id, adds it to this
LibFxEffectDefs, and returns it. If this LibFxEffectDefs already contains a
FxEffectDef definition with the specified Id, does nothing and returns nil.

#### func (*LibFxEffectDefs) Len

```go
func (me *LibFxEffectDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibFxEffectDefs) New

```go
func (me *LibFxEffectDefs) New(id string) (def *FxEffectDef)
```
Creates a new FxEffectDef definition with the specified Id and returns it, but
does not add it to this LibFxEffectDefs.

#### func (*LibFxEffectDefs) Remove

```go
func (me *LibFxEffectDefs) Remove(id string)
```
Removes the FxEffectDef with the specified Id from this LibFxEffectDefs.

#### func (*LibFxEffectDefs) SyncChanges

```go
func (me *LibFxEffectDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibFxEffectDefs that need to be picked up. Call this after you have made
a number of changes to this LibFxEffectDefs library or its FxEffectDef
definitions. Also called by the global SyncChanges() function.

#### type LibFxImageDefs

```go
type LibFxImageDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*FxImageDef
}
```

A library that contains FxImageDefs associated by their Id. To create a new
LibFxImageDefs library, ONLY use the LibsFxImageDef.New() or
LibsFxImageDef.AddNew() methods.

#### func (*LibFxImageDefs) Add

```go
func (me *LibFxImageDefs) Add(d *FxImageDef) (n *FxImageDef)
```
Adds the specified FxImageDef definition to this LibFxImageDefs, and returns it.
If this LibFxImageDefs already contains a FxImageDef definition with the same
Id, does nothing and returns nil.

#### func (*LibFxImageDefs) AddNew

```go
func (me *LibFxImageDefs) AddNew(id string) *FxImageDef
```
Creates a new FxImageDef definition with the specified Id, adds it to this
LibFxImageDefs, and returns it. If this LibFxImageDefs already contains a
FxImageDef definition with the specified Id, does nothing and returns nil.

#### func (*LibFxImageDefs) Len

```go
func (me *LibFxImageDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibFxImageDefs) New

```go
func (me *LibFxImageDefs) New(id string) (def *FxImageDef)
```
Creates a new FxImageDef definition with the specified Id and returns it, but
does not add it to this LibFxImageDefs.

#### func (*LibFxImageDefs) Remove

```go
func (me *LibFxImageDefs) Remove(id string)
```
Removes the FxImageDef with the specified Id from this LibFxImageDefs.

#### func (*LibFxImageDefs) SyncChanges

```go
func (me *LibFxImageDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibFxImageDefs that need to be picked up. Call this after you have made
a number of changes to this LibFxImageDefs library or its FxImageDef
definitions. Also called by the global SyncChanges() function.

#### type LibFxMaterialDefs

```go
type LibFxMaterialDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*FxMaterialDef
}
```

A library that contains FxMaterialDefs associated by their Id. To create a new
LibFxMaterialDefs library, ONLY use the LibsFxMaterialDef.New() or
LibsFxMaterialDef.AddNew() methods.

#### func (*LibFxMaterialDefs) Add

```go
func (me *LibFxMaterialDefs) Add(d *FxMaterialDef) (n *FxMaterialDef)
```
Adds the specified FxMaterialDef definition to this LibFxMaterialDefs, and
returns it. If this LibFxMaterialDefs already contains a FxMaterialDef
definition with the same Id, does nothing and returns nil.

#### func (*LibFxMaterialDefs) AddNew

```go
func (me *LibFxMaterialDefs) AddNew(id string) *FxMaterialDef
```
Creates a new FxMaterialDef definition with the specified Id, adds it to this
LibFxMaterialDefs, and returns it. If this LibFxMaterialDefs already contains a
FxMaterialDef definition with the specified Id, does nothing and returns nil.

#### func (*LibFxMaterialDefs) Len

```go
func (me *LibFxMaterialDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibFxMaterialDefs) New

```go
func (me *LibFxMaterialDefs) New(id string) (def *FxMaterialDef)
```
Creates a new FxMaterialDef definition with the specified Id and returns it, but
does not add it to this LibFxMaterialDefs.

#### func (*LibFxMaterialDefs) Remove

```go
func (me *LibFxMaterialDefs) Remove(id string)
```
Removes the FxMaterialDef with the specified Id from this LibFxMaterialDefs.

#### func (*LibFxMaterialDefs) SyncChanges

```go
func (me *LibFxMaterialDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibFxMaterialDefs that need to be picked up. Call this after you have
made a number of changes to this LibFxMaterialDefs library or its FxMaterialDef
definitions. Also called by the global SyncChanges() function.

#### type LibGeometryDefs

```go
type LibGeometryDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*GeometryDef
}
```

A library that contains GeometryDefs associated by their Id. To create a new
LibGeometryDefs library, ONLY use the LibsGeometryDef.New() or
LibsGeometryDef.AddNew() methods.

#### func (*LibGeometryDefs) Add

```go
func (me *LibGeometryDefs) Add(d *GeometryDef) (n *GeometryDef)
```
Adds the specified GeometryDef definition to this LibGeometryDefs, and returns
it. If this LibGeometryDefs already contains a GeometryDef definition with the
same Id, does nothing and returns nil.

#### func (*LibGeometryDefs) AddNew

```go
func (me *LibGeometryDefs) AddNew(id string) *GeometryDef
```
Creates a new GeometryDef definition with the specified Id, adds it to this
LibGeometryDefs, and returns it. If this LibGeometryDefs already contains a
GeometryDef definition with the specified Id, does nothing and returns nil.

#### func (*LibGeometryDefs) Len

```go
func (me *LibGeometryDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibGeometryDefs) New

```go
func (me *LibGeometryDefs) New(id string) (def *GeometryDef)
```
Creates a new GeometryDef definition with the specified Id and returns it, but
does not add it to this LibGeometryDefs.

#### func (*LibGeometryDefs) Remove

```go
func (me *LibGeometryDefs) Remove(id string)
```
Removes the GeometryDef with the specified Id from this LibGeometryDefs.

#### func (*LibGeometryDefs) SyncChanges

```go
func (me *LibGeometryDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibGeometryDefs that need to be picked up. Call this after you have made
a number of changes to this LibGeometryDefs library or its GeometryDef
definitions. Also called by the global SyncChanges() function.

#### type LibKxArticulatedSystemDefs

```go
type LibKxArticulatedSystemDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*KxArticulatedSystemDef
}
```

A library that contains KxArticulatedSystemDefs associated by their Id. To
create a new LibKxArticulatedSystemDefs library, ONLY use the
LibsKxArticulatedSystemDef.New() or LibsKxArticulatedSystemDef.AddNew() methods.

#### func (*LibKxArticulatedSystemDefs) Add

```go
func (me *LibKxArticulatedSystemDefs) Add(d *KxArticulatedSystemDef) (n *KxArticulatedSystemDef)
```
Adds the specified KxArticulatedSystemDef definition to this
LibKxArticulatedSystemDefs, and returns it. If this LibKxArticulatedSystemDefs
already contains a KxArticulatedSystemDef definition with the same Id, does
nothing and returns nil.

#### func (*LibKxArticulatedSystemDefs) AddNew

```go
func (me *LibKxArticulatedSystemDefs) AddNew(id string) *KxArticulatedSystemDef
```
Creates a new KxArticulatedSystemDef definition with the specified Id, adds it
to this LibKxArticulatedSystemDefs, and returns it. If this
LibKxArticulatedSystemDefs already contains a KxArticulatedSystemDef definition
with the specified Id, does nothing and returns nil.

#### func (*LibKxArticulatedSystemDefs) Len

```go
func (me *LibKxArticulatedSystemDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibKxArticulatedSystemDefs) New

```go
func (me *LibKxArticulatedSystemDefs) New(id string) (def *KxArticulatedSystemDef)
```
Creates a new KxArticulatedSystemDef definition with the specified Id and
returns it, but does not add it to this LibKxArticulatedSystemDefs.

#### func (*LibKxArticulatedSystemDefs) Remove

```go
func (me *LibKxArticulatedSystemDefs) Remove(id string)
```
Removes the KxArticulatedSystemDef with the specified Id from this
LibKxArticulatedSystemDefs.

#### func (*LibKxArticulatedSystemDefs) SyncChanges

```go
func (me *LibKxArticulatedSystemDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibKxArticulatedSystemDefs that need to be picked up. Call this after
you have made a number of changes to this LibKxArticulatedSystemDefs library or
its KxArticulatedSystemDef definitions. Also called by the global SyncChanges()
function.

#### type LibKxJointDefs

```go
type LibKxJointDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*KxJointDef
}
```

A library that contains KxJointDefs associated by their Id. To create a new
LibKxJointDefs library, ONLY use the LibsKxJointDef.New() or
LibsKxJointDef.AddNew() methods.

#### func (*LibKxJointDefs) Add

```go
func (me *LibKxJointDefs) Add(d *KxJointDef) (n *KxJointDef)
```
Adds the specified KxJointDef definition to this LibKxJointDefs, and returns it.
If this LibKxJointDefs already contains a KxJointDef definition with the same
Id, does nothing and returns nil.

#### func (*LibKxJointDefs) AddNew

```go
func (me *LibKxJointDefs) AddNew(id string) *KxJointDef
```
Creates a new KxJointDef definition with the specified Id, adds it to this
LibKxJointDefs, and returns it. If this LibKxJointDefs already contains a
KxJointDef definition with the specified Id, does nothing and returns nil.

#### func (*LibKxJointDefs) Len

```go
func (me *LibKxJointDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibKxJointDefs) New

```go
func (me *LibKxJointDefs) New(id string) (def *KxJointDef)
```
Creates a new KxJointDef definition with the specified Id and returns it, but
does not add it to this LibKxJointDefs.

#### func (*LibKxJointDefs) Remove

```go
func (me *LibKxJointDefs) Remove(id string)
```
Removes the KxJointDef with the specified Id from this LibKxJointDefs.

#### func (*LibKxJointDefs) SyncChanges

```go
func (me *LibKxJointDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibKxJointDefs that need to be picked up. Call this after you have made
a number of changes to this LibKxJointDefs library or its KxJointDef
definitions. Also called by the global SyncChanges() function.

#### type LibKxModelDefs

```go
type LibKxModelDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*KxModelDef
}
```

A library that contains KxModelDefs associated by their Id. To create a new
LibKxModelDefs library, ONLY use the LibsKxModelDef.New() or
LibsKxModelDef.AddNew() methods.

#### func (*LibKxModelDefs) Add

```go
func (me *LibKxModelDefs) Add(d *KxModelDef) (n *KxModelDef)
```
Adds the specified KxModelDef definition to this LibKxModelDefs, and returns it.
If this LibKxModelDefs already contains a KxModelDef definition with the same
Id, does nothing and returns nil.

#### func (*LibKxModelDefs) AddNew

```go
func (me *LibKxModelDefs) AddNew(id string) *KxModelDef
```
Creates a new KxModelDef definition with the specified Id, adds it to this
LibKxModelDefs, and returns it. If this LibKxModelDefs already contains a
KxModelDef definition with the specified Id, does nothing and returns nil.

#### func (*LibKxModelDefs) Len

```go
func (me *LibKxModelDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibKxModelDefs) New

```go
func (me *LibKxModelDefs) New(id string) (def *KxModelDef)
```
Creates a new KxModelDef definition with the specified Id and returns it, but
does not add it to this LibKxModelDefs.

#### func (*LibKxModelDefs) Remove

```go
func (me *LibKxModelDefs) Remove(id string)
```
Removes the KxModelDef with the specified Id from this LibKxModelDefs.

#### func (*LibKxModelDefs) SyncChanges

```go
func (me *LibKxModelDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibKxModelDefs that need to be picked up. Call this after you have made
a number of changes to this LibKxModelDefs library or its KxModelDef
definitions. Also called by the global SyncChanges() function.

#### type LibKxSceneDefs

```go
type LibKxSceneDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*KxSceneDef
}
```

A library that contains KxSceneDefs associated by their Id. To create a new
LibKxSceneDefs library, ONLY use the LibsKxSceneDef.New() or
LibsKxSceneDef.AddNew() methods.

#### func (*LibKxSceneDefs) Add

```go
func (me *LibKxSceneDefs) Add(d *KxSceneDef) (n *KxSceneDef)
```
Adds the specified KxSceneDef definition to this LibKxSceneDefs, and returns it.
If this LibKxSceneDefs already contains a KxSceneDef definition with the same
Id, does nothing and returns nil.

#### func (*LibKxSceneDefs) AddNew

```go
func (me *LibKxSceneDefs) AddNew(id string) *KxSceneDef
```
Creates a new KxSceneDef definition with the specified Id, adds it to this
LibKxSceneDefs, and returns it. If this LibKxSceneDefs already contains a
KxSceneDef definition with the specified Id, does nothing and returns nil.

#### func (*LibKxSceneDefs) Len

```go
func (me *LibKxSceneDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibKxSceneDefs) New

```go
func (me *LibKxSceneDefs) New(id string) (def *KxSceneDef)
```
Creates a new KxSceneDef definition with the specified Id and returns it, but
does not add it to this LibKxSceneDefs.

#### func (*LibKxSceneDefs) Remove

```go
func (me *LibKxSceneDefs) Remove(id string)
```
Removes the KxSceneDef with the specified Id from this LibKxSceneDefs.

#### func (*LibKxSceneDefs) SyncChanges

```go
func (me *LibKxSceneDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibKxSceneDefs that need to be picked up. Call this after you have made
a number of changes to this LibKxSceneDefs library or its KxSceneDef
definitions. Also called by the global SyncChanges() function.

#### type LibLightDefs

```go
type LibLightDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*LightDef
}
```

A library that contains LightDefs associated by their Id. To create a new
LibLightDefs library, ONLY use the LibsLightDef.New() or LibsLightDef.AddNew()
methods.

#### func (*LibLightDefs) Add

```go
func (me *LibLightDefs) Add(d *LightDef) (n *LightDef)
```
Adds the specified LightDef definition to this LibLightDefs, and returns it. If
this LibLightDefs already contains a LightDef definition with the same Id, does
nothing and returns nil.

#### func (*LibLightDefs) AddNew

```go
func (me *LibLightDefs) AddNew(id string) *LightDef
```
Creates a new LightDef definition with the specified Id, adds it to this
LibLightDefs, and returns it. If this LibLightDefs already contains a LightDef
definition with the specified Id, does nothing and returns nil.

#### func (*LibLightDefs) Len

```go
func (me *LibLightDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibLightDefs) New

```go
func (me *LibLightDefs) New(id string) (def *LightDef)
```
Creates a new LightDef definition with the specified Id and returns it, but does
not add it to this LibLightDefs.

#### func (*LibLightDefs) Remove

```go
func (me *LibLightDefs) Remove(id string)
```
Removes the LightDef with the specified Id from this LibLightDefs.

#### func (*LibLightDefs) SyncChanges

```go
func (me *LibLightDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibLightDefs that need to be picked up. Call this after you have made a
number of changes to this LibLightDefs library or its LightDef definitions. Also
called by the global SyncChanges() function.

#### type LibNodeDefs

```go
type LibNodeDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*NodeDef
}
```

A library that contains NodeDefs associated by their Id. To create a new
LibNodeDefs library, ONLY use the LibsNodeDef.New() or LibsNodeDef.AddNew()
methods.

#### func (*LibNodeDefs) Add

```go
func (me *LibNodeDefs) Add(d *NodeDef) (n *NodeDef)
```
Adds the specified NodeDef definition to this LibNodeDefs, and returns it. If
this LibNodeDefs already contains a NodeDef definition with the same Id, does
nothing and returns nil.

#### func (*LibNodeDefs) AddNew

```go
func (me *LibNodeDefs) AddNew(id string) *NodeDef
```
Creates a new NodeDef definition with the specified Id, adds it to this
LibNodeDefs, and returns it. If this LibNodeDefs already contains a NodeDef
definition with the specified Id, does nothing and returns nil.

#### func (*LibNodeDefs) Len

```go
func (me *LibNodeDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibNodeDefs) New

```go
func (me *LibNodeDefs) New(id string) (def *NodeDef)
```
Creates a new NodeDef definition with the specified Id and returns it, but does
not add it to this LibNodeDefs.

#### func (*LibNodeDefs) Remove

```go
func (me *LibNodeDefs) Remove(id string)
```
Removes the NodeDef with the specified Id from this LibNodeDefs.

#### func (*LibNodeDefs) SyncChanges

```go
func (me *LibNodeDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibNodeDefs that need to be picked up. Call this after you have made a
number of changes to this LibNodeDefs library or its NodeDef definitions. Also
called by the global SyncChanges() function.

#### type LibPxForceFieldDefs

```go
type LibPxForceFieldDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*PxForceFieldDef
}
```

A library that contains PxForceFieldDefs associated by their Id. To create a new
LibPxForceFieldDefs library, ONLY use the LibsPxForceFieldDef.New() or
LibsPxForceFieldDef.AddNew() methods.

#### func (*LibPxForceFieldDefs) Add

```go
func (me *LibPxForceFieldDefs) Add(d *PxForceFieldDef) (n *PxForceFieldDef)
```
Adds the specified PxForceFieldDef definition to this LibPxForceFieldDefs, and
returns it. If this LibPxForceFieldDefs already contains a PxForceFieldDef
definition with the same Id, does nothing and returns nil.

#### func (*LibPxForceFieldDefs) AddNew

```go
func (me *LibPxForceFieldDefs) AddNew(id string) *PxForceFieldDef
```
Creates a new PxForceFieldDef definition with the specified Id, adds it to this
LibPxForceFieldDefs, and returns it. If this LibPxForceFieldDefs already
contains a PxForceFieldDef definition with the specified Id, does nothing and
returns nil.

#### func (*LibPxForceFieldDefs) Len

```go
func (me *LibPxForceFieldDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibPxForceFieldDefs) New

```go
func (me *LibPxForceFieldDefs) New(id string) (def *PxForceFieldDef)
```
Creates a new PxForceFieldDef definition with the specified Id and returns it,
but does not add it to this LibPxForceFieldDefs.

#### func (*LibPxForceFieldDefs) Remove

```go
func (me *LibPxForceFieldDefs) Remove(id string)
```
Removes the PxForceFieldDef with the specified Id from this LibPxForceFieldDefs.

#### func (*LibPxForceFieldDefs) SyncChanges

```go
func (me *LibPxForceFieldDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibPxForceFieldDefs that need to be picked up. Call this after you have
made a number of changes to this LibPxForceFieldDefs library or its
PxForceFieldDef definitions. Also called by the global SyncChanges() function.

#### type LibPxMaterialDefs

```go
type LibPxMaterialDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*PxMaterialDef
}
```

A library that contains PxMaterialDefs associated by their Id. To create a new
LibPxMaterialDefs library, ONLY use the LibsPxMaterialDef.New() or
LibsPxMaterialDef.AddNew() methods.

#### func (*LibPxMaterialDefs) Add

```go
func (me *LibPxMaterialDefs) Add(d *PxMaterialDef) (n *PxMaterialDef)
```
Adds the specified PxMaterialDef definition to this LibPxMaterialDefs, and
returns it. If this LibPxMaterialDefs already contains a PxMaterialDef
definition with the same Id, does nothing and returns nil.

#### func (*LibPxMaterialDefs) AddNew

```go
func (me *LibPxMaterialDefs) AddNew(id string) *PxMaterialDef
```
Creates a new PxMaterialDef definition with the specified Id, adds it to this
LibPxMaterialDefs, and returns it. If this LibPxMaterialDefs already contains a
PxMaterialDef definition with the specified Id, does nothing and returns nil.

#### func (*LibPxMaterialDefs) Len

```go
func (me *LibPxMaterialDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibPxMaterialDefs) New

```go
func (me *LibPxMaterialDefs) New(id string) (def *PxMaterialDef)
```
Creates a new PxMaterialDef definition with the specified Id and returns it, but
does not add it to this LibPxMaterialDefs.

#### func (*LibPxMaterialDefs) Remove

```go
func (me *LibPxMaterialDefs) Remove(id string)
```
Removes the PxMaterialDef with the specified Id from this LibPxMaterialDefs.

#### func (*LibPxMaterialDefs) SyncChanges

```go
func (me *LibPxMaterialDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibPxMaterialDefs that need to be picked up. Call this after you have
made a number of changes to this LibPxMaterialDefs library or its PxMaterialDef
definitions. Also called by the global SyncChanges() function.

#### type LibPxModelDefs

```go
type LibPxModelDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*PxModelDef
}
```

A library that contains PxModelDefs associated by their Id. To create a new
LibPxModelDefs library, ONLY use the LibsPxModelDef.New() or
LibsPxModelDef.AddNew() methods.

#### func (*LibPxModelDefs) Add

```go
func (me *LibPxModelDefs) Add(d *PxModelDef) (n *PxModelDef)
```
Adds the specified PxModelDef definition to this LibPxModelDefs, and returns it.
If this LibPxModelDefs already contains a PxModelDef definition with the same
Id, does nothing and returns nil.

#### func (*LibPxModelDefs) AddNew

```go
func (me *LibPxModelDefs) AddNew(id string) *PxModelDef
```
Creates a new PxModelDef definition with the specified Id, adds it to this
LibPxModelDefs, and returns it. If this LibPxModelDefs already contains a
PxModelDef definition with the specified Id, does nothing and returns nil.

#### func (*LibPxModelDefs) Len

```go
func (me *LibPxModelDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibPxModelDefs) New

```go
func (me *LibPxModelDefs) New(id string) (def *PxModelDef)
```
Creates a new PxModelDef definition with the specified Id and returns it, but
does not add it to this LibPxModelDefs.

#### func (*LibPxModelDefs) Remove

```go
func (me *LibPxModelDefs) Remove(id string)
```
Removes the PxModelDef with the specified Id from this LibPxModelDefs.

#### func (*LibPxModelDefs) SyncChanges

```go
func (me *LibPxModelDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibPxModelDefs that need to be picked up. Call this after you have made
a number of changes to this LibPxModelDefs library or its PxModelDef
definitions. Also called by the global SyncChanges() function.

#### type LibPxSceneDefs

```go
type LibPxSceneDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*PxSceneDef
}
```

A library that contains PxSceneDefs associated by their Id. To create a new
LibPxSceneDefs library, ONLY use the LibsPxSceneDef.New() or
LibsPxSceneDef.AddNew() methods.

#### func (*LibPxSceneDefs) Add

```go
func (me *LibPxSceneDefs) Add(d *PxSceneDef) (n *PxSceneDef)
```
Adds the specified PxSceneDef definition to this LibPxSceneDefs, and returns it.
If this LibPxSceneDefs already contains a PxSceneDef definition with the same
Id, does nothing and returns nil.

#### func (*LibPxSceneDefs) AddNew

```go
func (me *LibPxSceneDefs) AddNew(id string) *PxSceneDef
```
Creates a new PxSceneDef definition with the specified Id, adds it to this
LibPxSceneDefs, and returns it. If this LibPxSceneDefs already contains a
PxSceneDef definition with the specified Id, does nothing and returns nil.

#### func (*LibPxSceneDefs) Len

```go
func (me *LibPxSceneDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibPxSceneDefs) New

```go
func (me *LibPxSceneDefs) New(id string) (def *PxSceneDef)
```
Creates a new PxSceneDef definition with the specified Id and returns it, but
does not add it to this LibPxSceneDefs.

#### func (*LibPxSceneDefs) Remove

```go
func (me *LibPxSceneDefs) Remove(id string)
```
Removes the PxSceneDef with the specified Id from this LibPxSceneDefs.

#### func (*LibPxSceneDefs) SyncChanges

```go
func (me *LibPxSceneDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibPxSceneDefs that need to be picked up. Call this after you have made
a number of changes to this LibPxSceneDefs library or its PxSceneDef
definitions. Also called by the global SyncChanges() function.

#### type LibVisualSceneDefs

```go
type LibVisualSceneDefs struct {
	//	Id, Name
	BaseLib

	//	The underlying hash-table. NOTE -- this is for easier read-access and range-iteration:
	//	DO NOT write to M, instead use the Add(), AddNew(), Remove() methods ONLY or bugs WILL ensue.
	M map[string]*VisualSceneDef
}
```

A library that contains VisualSceneDefs associated by their Id. To create a new
LibVisualSceneDefs library, ONLY use the LibsVisualSceneDef.New() or
LibsVisualSceneDef.AddNew() methods.

#### func (*LibVisualSceneDefs) Add

```go
func (me *LibVisualSceneDefs) Add(d *VisualSceneDef) (n *VisualSceneDef)
```
Adds the specified VisualSceneDef definition to this LibVisualSceneDefs, and
returns it. If this LibVisualSceneDefs already contains a VisualSceneDef
definition with the same Id, does nothing and returns nil.

#### func (*LibVisualSceneDefs) AddNew

```go
func (me *LibVisualSceneDefs) AddNew(id string) *VisualSceneDef
```
Creates a new VisualSceneDef definition with the specified Id, adds it to this
LibVisualSceneDefs, and returns it. If this LibVisualSceneDefs already contains
a VisualSceneDef definition with the specified Id, does nothing and returns nil.

#### func (*LibVisualSceneDefs) Len

```go
func (me *LibVisualSceneDefs) Len() int
```
Convenience short-hand for len(lib.M)

#### func (*LibVisualSceneDefs) New

```go
func (me *LibVisualSceneDefs) New(id string) (def *VisualSceneDef)
```
Creates a new VisualSceneDef definition with the specified Id and returns it,
but does not add it to this LibVisualSceneDefs.

#### func (*LibVisualSceneDefs) Remove

```go
func (me *LibVisualSceneDefs) Remove(id string)
```
Removes the VisualSceneDef with the specified Id from this LibVisualSceneDefs.

#### func (*LibVisualSceneDefs) SyncChanges

```go
func (me *LibVisualSceneDefs) SyncChanges()
```
Signals to the core package (or your custom package) that changes have been made
to this LibVisualSceneDefs that need to be picked up. Call this after you have
made a number of changes to this LibVisualSceneDefs library or its
VisualSceneDef definitions. Also called by the global SyncChanges() function.

#### type LibsAnimationClipDef

```go
type LibsAnimationClipDef map[string]*LibAnimationClipDefs
```

The underlying type of the global AllAnimationClipDefLibs variable: a hash-table
that contains LibAnimationClipDefs libraries associated by their Id.

#### func (LibsAnimationClipDef) AddNew

```go
func (me LibsAnimationClipDef) AddNew(id string) (lib *LibAnimationClipDefs)
```
Creates a new LibAnimationClipDefs library with the specified Id, adds it to
this LibsAnimationClipDef, and returns it. If this LibsAnimationClipDef already
contains a LibAnimationClipDefs library with the specified Id, does nothing and
returns nil.

#### type LibsAnimationDef

```go
type LibsAnimationDef map[string]*LibAnimationDefs
```

The underlying type of the global AllAnimationDefLibs variable: a hash-table
that contains LibAnimationDefs libraries associated by their Id.

#### func (LibsAnimationDef) AddNew

```go
func (me LibsAnimationDef) AddNew(id string) (lib *LibAnimationDefs)
```
Creates a new LibAnimationDefs library with the specified Id, adds it to this
LibsAnimationDef, and returns it. If this LibsAnimationDef already contains a
LibAnimationDefs library with the specified Id, does nothing and returns nil.

#### type LibsCameraDef

```go
type LibsCameraDef map[string]*LibCameraDefs
```

The underlying type of the global AllCameraDefLibs variable: a hash-table that
contains LibCameraDefs libraries associated by their Id.

#### func (LibsCameraDef) AddNew

```go
func (me LibsCameraDef) AddNew(id string) (lib *LibCameraDefs)
```
Creates a new LibCameraDefs library with the specified Id, adds it to this
LibsCameraDef, and returns it. If this LibsCameraDef already contains a
LibCameraDefs library with the specified Id, does nothing and returns nil.

#### type LibsControllerDef

```go
type LibsControllerDef map[string]*LibControllerDefs
```

The underlying type of the global AllControllerDefLibs variable: a hash-table
that contains LibControllerDefs libraries associated by their Id.

#### func (LibsControllerDef) AddNew

```go
func (me LibsControllerDef) AddNew(id string) (lib *LibControllerDefs)
```
Creates a new LibControllerDefs library with the specified Id, adds it to this
LibsControllerDef, and returns it. If this LibsControllerDef already contains a
LibControllerDefs library with the specified Id, does nothing and returns nil.

#### type LibsFormulaDef

```go
type LibsFormulaDef map[string]*LibFormulaDefs
```

The underlying type of the global AllFormulaDefLibs variable: a hash-table that
contains LibFormulaDefs libraries associated by their Id.

#### func (LibsFormulaDef) AddNew

```go
func (me LibsFormulaDef) AddNew(id string) (lib *LibFormulaDefs)
```
Creates a new LibFormulaDefs library with the specified Id, adds it to this
LibsFormulaDef, and returns it. If this LibsFormulaDef already contains a
LibFormulaDefs library with the specified Id, does nothing and returns nil.

#### type LibsFxEffectDef

```go
type LibsFxEffectDef map[string]*LibFxEffectDefs
```

The underlying type of the global AllFxEffectDefLibs variable: a hash-table that
contains LibFxEffectDefs libraries associated by their Id.

#### func (LibsFxEffectDef) AddNew

```go
func (me LibsFxEffectDef) AddNew(id string) (lib *LibFxEffectDefs)
```
Creates a new LibFxEffectDefs library with the specified Id, adds it to this
LibsFxEffectDef, and returns it. If this LibsFxEffectDef already contains a
LibFxEffectDefs library with the specified Id, does nothing and returns nil.

#### type LibsFxImageDef

```go
type LibsFxImageDef map[string]*LibFxImageDefs
```

The underlying type of the global AllFxImageDefLibs variable: a hash-table that
contains LibFxImageDefs libraries associated by their Id.

#### func (LibsFxImageDef) AddNew

```go
func (me LibsFxImageDef) AddNew(id string) (lib *LibFxImageDefs)
```
Creates a new LibFxImageDefs library with the specified Id, adds it to this
LibsFxImageDef, and returns it. If this LibsFxImageDef already contains a
LibFxImageDefs library with the specified Id, does nothing and returns nil.

#### type LibsFxMaterialDef

```go
type LibsFxMaterialDef map[string]*LibFxMaterialDefs
```

The underlying type of the global AllFxMaterialDefLibs variable: a hash-table
that contains LibFxMaterialDefs libraries associated by their Id.

#### func (LibsFxMaterialDef) AddNew

```go
func (me LibsFxMaterialDef) AddNew(id string) (lib *LibFxMaterialDefs)
```
Creates a new LibFxMaterialDefs library with the specified Id, adds it to this
LibsFxMaterialDef, and returns it. If this LibsFxMaterialDef already contains a
LibFxMaterialDefs library with the specified Id, does nothing and returns nil.

#### type LibsGeometryDef

```go
type LibsGeometryDef map[string]*LibGeometryDefs
```

The underlying type of the global AllGeometryDefLibs variable: a hash-table that
contains LibGeometryDefs libraries associated by their Id.

#### func (LibsGeometryDef) AddNew

```go
func (me LibsGeometryDef) AddNew(id string) (lib *LibGeometryDefs)
```
Creates a new LibGeometryDefs library with the specified Id, adds it to this
LibsGeometryDef, and returns it. If this LibsGeometryDef already contains a
LibGeometryDefs library with the specified Id, does nothing and returns nil.

#### type LibsKxArticulatedSystemDef

```go
type LibsKxArticulatedSystemDef map[string]*LibKxArticulatedSystemDefs
```

The underlying type of the global AllKxArticulatedSystemDefLibs variable: a
hash-table that contains LibKxArticulatedSystemDefs libraries associated by
their Id.

#### func (LibsKxArticulatedSystemDef) AddNew

```go
func (me LibsKxArticulatedSystemDef) AddNew(id string) (lib *LibKxArticulatedSystemDefs)
```
Creates a new LibKxArticulatedSystemDefs library with the specified Id, adds it
to this LibsKxArticulatedSystemDef, and returns it. If this
LibsKxArticulatedSystemDef already contains a LibKxArticulatedSystemDefs library
with the specified Id, does nothing and returns nil.

#### type LibsKxJointDef

```go
type LibsKxJointDef map[string]*LibKxJointDefs
```

The underlying type of the global AllKxJointDefLibs variable: a hash-table that
contains LibKxJointDefs libraries associated by their Id.

#### func (LibsKxJointDef) AddNew

```go
func (me LibsKxJointDef) AddNew(id string) (lib *LibKxJointDefs)
```
Creates a new LibKxJointDefs library with the specified Id, adds it to this
LibsKxJointDef, and returns it. If this LibsKxJointDef already contains a
LibKxJointDefs library with the specified Id, does nothing and returns nil.

#### type LibsKxModelDef

```go
type LibsKxModelDef map[string]*LibKxModelDefs
```

The underlying type of the global AllKxModelDefLibs variable: a hash-table that
contains LibKxModelDefs libraries associated by their Id.

#### func (LibsKxModelDef) AddNew

```go
func (me LibsKxModelDef) AddNew(id string) (lib *LibKxModelDefs)
```
Creates a new LibKxModelDefs library with the specified Id, adds it to this
LibsKxModelDef, and returns it. If this LibsKxModelDef already contains a
LibKxModelDefs library with the specified Id, does nothing and returns nil.

#### type LibsKxSceneDef

```go
type LibsKxSceneDef map[string]*LibKxSceneDefs
```

The underlying type of the global AllKxSceneDefLibs variable: a hash-table that
contains LibKxSceneDefs libraries associated by their Id.

#### func (LibsKxSceneDef) AddNew

```go
func (me LibsKxSceneDef) AddNew(id string) (lib *LibKxSceneDefs)
```
Creates a new LibKxSceneDefs library with the specified Id, adds it to this
LibsKxSceneDef, and returns it. If this LibsKxSceneDef already contains a
LibKxSceneDefs library with the specified Id, does nothing and returns nil.

#### type LibsLightDef

```go
type LibsLightDef map[string]*LibLightDefs
```

The underlying type of the global AllLightDefLibs variable: a hash-table that
contains LibLightDefs libraries associated by their Id.

#### func (LibsLightDef) AddNew

```go
func (me LibsLightDef) AddNew(id string) (lib *LibLightDefs)
```
Creates a new LibLightDefs library with the specified Id, adds it to this
LibsLightDef, and returns it. If this LibsLightDef already contains a
LibLightDefs library with the specified Id, does nothing and returns nil.

#### type LibsNodeDef

```go
type LibsNodeDef map[string]*LibNodeDefs
```

The underlying type of the global AllNodeDefLibs variable: a hash-table that
contains LibNodeDefs libraries associated by their Id.

#### func (LibsNodeDef) AddNew

```go
func (me LibsNodeDef) AddNew(id string) (lib *LibNodeDefs)
```
Creates a new LibNodeDefs library with the specified Id, adds it to this
LibsNodeDef, and returns it. If this LibsNodeDef already contains a LibNodeDefs
library with the specified Id, does nothing and returns nil.

#### type LibsPxForceFieldDef

```go
type LibsPxForceFieldDef map[string]*LibPxForceFieldDefs
```

The underlying type of the global AllPxForceFieldDefLibs variable: a hash-table
that contains LibPxForceFieldDefs libraries associated by their Id.

#### func (LibsPxForceFieldDef) AddNew

```go
func (me LibsPxForceFieldDef) AddNew(id string) (lib *LibPxForceFieldDefs)
```
Creates a new LibPxForceFieldDefs library with the specified Id, adds it to this
LibsPxForceFieldDef, and returns it. If this LibsPxForceFieldDef already
contains a LibPxForceFieldDefs library with the specified Id, does nothing and
returns nil.

#### type LibsPxMaterialDef

```go
type LibsPxMaterialDef map[string]*LibPxMaterialDefs
```

The underlying type of the global AllPxMaterialDefLibs variable: a hash-table
that contains LibPxMaterialDefs libraries associated by their Id.

#### func (LibsPxMaterialDef) AddNew

```go
func (me LibsPxMaterialDef) AddNew(id string) (lib *LibPxMaterialDefs)
```
Creates a new LibPxMaterialDefs library with the specified Id, adds it to this
LibsPxMaterialDef, and returns it. If this LibsPxMaterialDef already contains a
LibPxMaterialDefs library with the specified Id, does nothing and returns nil.

#### type LibsPxModelDef

```go
type LibsPxModelDef map[string]*LibPxModelDefs
```

The underlying type of the global AllPxModelDefLibs variable: a hash-table that
contains LibPxModelDefs libraries associated by their Id.

#### func (LibsPxModelDef) AddNew

```go
func (me LibsPxModelDef) AddNew(id string) (lib *LibPxModelDefs)
```
Creates a new LibPxModelDefs library with the specified Id, adds it to this
LibsPxModelDef, and returns it. If this LibsPxModelDef already contains a
LibPxModelDefs library with the specified Id, does nothing and returns nil.

#### type LibsPxSceneDef

```go
type LibsPxSceneDef map[string]*LibPxSceneDefs
```

The underlying type of the global AllPxSceneDefLibs variable: a hash-table that
contains LibPxSceneDefs libraries associated by their Id.

#### func (LibsPxSceneDef) AddNew

```go
func (me LibsPxSceneDef) AddNew(id string) (lib *LibPxSceneDefs)
```
Creates a new LibPxSceneDefs library with the specified Id, adds it to this
LibsPxSceneDef, and returns it. If this LibsPxSceneDef already contains a
LibPxSceneDefs library with the specified Id, does nothing and returns nil.

#### type LibsVisualSceneDef

```go
type LibsVisualSceneDef map[string]*LibVisualSceneDefs
```

The underlying type of the global AllVisualSceneDefLibs variable: a hash-table
that contains LibVisualSceneDefs libraries associated by their Id.

#### func (LibsVisualSceneDef) AddNew

```go
func (me LibsVisualSceneDef) AddNew(id string) (lib *LibVisualSceneDefs)
```
Creates a new LibVisualSceneDefs library with the specified Id, adds it to this
LibsVisualSceneDef, and returns it. If this LibsVisualSceneDef already contains
a LibVisualSceneDefs library with the specified Id, does nothing and returns
nil.

#### type LightAmbient

```go
type LightAmbient struct {
	//	Color
	LightBase
}
```

Describes an ambient light source. An ambient light is one that lights
everything evenly, regardless of location or orientation.

#### type LightAttenuation

```go
type LightAttenuation struct {
	//	Constant light attenuation. Defaults to 1.
	Constant SidFloat

	//	Linear light attenuation.
	Linear SidFloat

	//	Quadratic light attenuation.
	Quadratic SidFloat
}
```

Describes how the intensity of a light source is attenuated.

#### func  NewLightAttenuation

```go
func NewLightAttenuation() (me *LightAttenuation)
```
Constructor

#### func (*LightAttenuation) AccessField

```go
func (me *LightAttenuation) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "Constant", "Linear",
"Quadratic".

#### type LightBase

```go
type LightBase struct {
	//	Three floating-point numbers specifying the color of this light.
	Color Float3
}
```

Contains three floating-point numbers specifying the color of a light.

#### type LightDef

```go
type LightDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Techniques
	HasTechniques

	//	Common-technique profile. At least and at most one of its fields should ever be set.
	TC struct {
		//	If set, this light declares an ambient light.
		Ambient *LightAmbient

		//	If set, this light declares a directional light.
		Directional *LightDirectional

		//	If set, this light declares a point light.
		Point *LightPoint

		//	If set, this light declares a spot light.
		Spot *LightSpot
	}
}
```

Declares a light source that illuminates a scene.

#### func (*LightDef) DefaultInst

```go
func (me *LightDef) DefaultInst() (inst *LightInst)
```
Returns "the default LightInst instance" referencing this LightDef definition.
That instance is created once when this method is first called on me, and will
have its Def field readily set to me.

#### func (*LightDef) Init

```go
func (me *LightDef) Init()
```
Initialization

#### func (*LightDef) NewInst

```go
func (me *LightDef) NewInst() (inst *LightInst)
```
Creates and returns a new LightInst instance referencing this LightDef
definition. Any LightInst created by this method will have its Def field readily
set to me.

#### type LightDirectional

```go
type LightDirectional struct {
	//	Color
	LightBase
}
```

Describes a directional light source. A directional light is one that lights
everything from the same direction, regardless of location. The light's default
direction vector in local coordinates is [0,0,-1], pointing down the negative z
axis. The actual direction of the light is defined by the transform of the node
where the light is instantiated.

#### type LightInst

```go
type LightInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *LightDef
}
```

Instantiates a light resource.

#### func (*LightInst) EnsureDef

```go
func (me *LightInst) EnsureDef() *LightDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct LightDef according
to the current me.DefRef value (by searching AllLightDefLibs). Then returns
me.Def. (Note, every LightInst's Def is nil initially, unless it was created via
LightDef.NewInst().)

#### func (*LightInst) Init

```go
func (me *LightInst) Init()
```
Initialization

#### type LightPoint

```go
type LightPoint struct {
	//	Color
	LightBase

	//	The intensity of a point light source is attenuated as the distance to the light source increases.
	Attenuation LightAttenuation
}
```

Describes a point light source. A point light source radiates light in all
directions from a known location in space. The position of the light is defined
by the transform of the node in which it is instantiated.

#### func  NewLightPoint

```go
func NewLightPoint() (me *LightPoint)
```
Constructor

#### type LightSpot

```go
type LightSpot struct {
	//	Color
	LightBase

	//	 The intensity of a spot light is also attenuated as the distance to the light source increases.
	Attenuation LightAttenuation

	//	The light's intensity is also attenuated as the radiation angle increases away from the direction of the light source.
	Falloff struct {
		//	Fall-off angle. Defaults to 180.
		Angle SidFloat

		//	Fall-off exponent.
		Exponent SidFloat
	}
}
```

Describes a spot light source. A spot light source radiates light in one
direction in a cone shape from a known location in space. The light's default
direction vector in local coordinates is [0,0,-1], pointing down the negative z
axis. The actual direction of the light is defined by the transform of the node
in which the light is instantiated.

#### func  NewLightSpot

```go
func NewLightSpot() (me *LightSpot)
```
Constructor

#### type MaterialBinding

```go
type MaterialBinding struct {
	//	Extras
	HasExtras

	//	Techniques
	HasTechniques

	//	Targets for animation
	Params []*Param

	//	Common-technique profile.
	TC struct {
		//	References to the materials included in this material binding.
		Materials []*FxMaterialInst
	}
}
```

Binds a specific material to a piece of geometry, binding varying and uniform
parameters at the same time.

#### type NodeDef

```go
type NodeDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Sid
	HasSid

	//	Indicates whether this node is a joint for a skin controller.
	IsSkinJoint bool

	//	The names of the layers to which this node belongs.
	Layers Layers

	//	Any combination of zero or more transformations of any type.
	Transforms []*Transform

	//	Content resources participating in this node.
	Insts struct {
		//	Cameras participating in this node.
		Camera []*CameraInst

		//	Controllers participating in this node.
		Controller []*ControllerInst

		//	Geometries participating in this node.
		Geometry []*GeometryInst

		//	Lights participating in this node.
		Light []*LightInst
	}
	//	Child nodes to recursively define a hierarchy.
	Nodes []ChildNode
}
```

Declares a point of interest in a scene.

#### func (*NodeDef) AccessField

```go
func (me *NodeDef) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "IsSkinJoint".

#### func (*NodeDef) DefaultInst

```go
func (me *NodeDef) DefaultInst() (inst *NodeInst)
```
Returns "the default NodeInst instance" referencing this NodeDef definition.
That instance is created once when this method is first called on me, and will
have its Def field readily set to me.

#### func (*NodeDef) Init

```go
func (me *NodeDef) Init()
```
Initialization

#### func (*NodeDef) NewInst

```go
func (me *NodeDef) NewInst() (inst *NodeInst)
```
Creates and returns a new NodeInst instance referencing this NodeDef definition.
Any NodeInst created by this method will have its Def field readily set to me.

#### type NodeInst

```go
type NodeInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *NodeDef

	//	Optional. The mechanism and use of this attribute is application-defined.
	//	For example, it can be used for bounding boxes or level of detail.
	Proxy RefId
}
```

Instantiates a node resource.

#### func (*NodeInst) AccessField

```go
func (me *NodeInst) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "Proxy".

#### func (*NodeInst) EnsureDef

```go
func (me *NodeInst) EnsureDef() *NodeDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct NodeDef according
to the current me.DefRef value (by searching AllNodeDefLibs). Then returns
me.Def. (Note, every NodeInst's Def is nil initially, unless it was created via
NodeDef.NewInst().)

#### func (*NodeInst) Init

```go
func (me *NodeInst) Init()
```
Initialization

#### type Param

```go
type Param struct {
	//	Name
	HasName

	//	Sid
	HasSid

	//	The user-defined meaning of the parameter.
	Semantic string

	//	The type of the value data. This text string must be understood by the application.
	Type string
}
```

Declares parametric information for its parent resource.

#### func (*Param) AccessField

```go
func (me *Param) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "Semantic", "Type".

#### type ParamDef

```go
type ParamDef struct {
	//	Sid
	HasSid

	//	Initial value for this parameter
	Value interface{}
}
```

Declares a new parameter for its parent resource, and assigns it an initial
value.

#### func (*ParamDef) AccessField

```go
func (me *ParamDef) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "Value".

#### type ParamDefs

```go
type ParamDefs map[string]*ParamDef
```

A hash-table containing parameter declarations of this resource.

#### func (ParamDefs) Set

```go
func (me ParamDefs) Set(sid string, val interface{})
```
If me does not contain a ParamDef with the specified Sid, adds it. Next, sets
the value of the ParamDef with the specified Sid in me to val.

#### type ParamInst

```go
type ParamInst struct {
	//	References the identifier of the pre-defined parameter (ParamDef) that will have its value set.
	Ref RefSid

	//	Indicates if the Value is a string referencing the identifier of a connected parameter.
	IsConnectParamRef bool

	//	The new value for the referenced parameter.
	Value interface{}
}
```

Assigns a new value to a previously defined parameter.

#### type ParamInsts

```go
type ParamInsts map[string]*ParamInst
```

A hash-table containing parameter values assigned by this resource.

#### type ParamOrBool

```go
type ParamOrBool struct {
	//	The value provided if Param is empty.
	B bool

	//	If set, refers to a previously defined parameter providing the value.
	Param RefParam
}
```

Provides a bool value.

#### type ParamOrFloat

```go
type ParamOrFloat struct {
	//	The value provided if Param is empty.
	F float64

	//	If set, refers to a previously defined parameter providing the value.
	Param RefParam
}
```

Provides a float64 value.

#### type ParamOrFloat2

```go
type ParamOrFloat2 struct {
	//	The values provided if Param is empty.
	F Float2

	//	If set, refers to a previously defined parameter providing the values.
	Param RefParam
}
```

Provides two float64 values.

#### func (*ParamOrFloat2) AccessIndex

```go
func (me *ParamOrFloat2) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type ParamOrInt

```go
type ParamOrInt struct {
	//	The value provided if Param is empty.
	I int64

	//	If set, refers to a previously defined parameter providing the value.
	Param RefParam
}
```

Provides a int64 value.

#### type ParamOrRefSid

```go
type ParamOrRefSid struct {
	//	The value provided if Param is empty.
	Sr RefSid

	//	If set, refers to a previously defined parameter providing the value.
	Param RefParam
}
```

Provides a RefSid value.

#### type ParamOrSidFloat

```go
type ParamOrSidFloat struct {
	//	The value provided if Param is empty.
	F SidFloat

	//	If set, refers to a previously defined parameter providing the value.
	Param RefParam
}
```

Provides a scoped float64 value.

#### type ParamOrUint

```go
type ParamOrUint struct {
	//	The value provided if Param is empty.
	U uint64

	//	If set, refers to a previously defined parameter providing the value.
	Param RefParam
}
```

Provides a uint64 value.

#### type PxCylinder

```go
type PxCylinder struct {
	//	Radii, Extras
	GeometryBrepCylinder

	//	Contains a floating-point value that represents the length of the cylinder along the y axis.
	Height float64
}
```

Declares a cylinder primitive that is centered around its local origin and
aligned along its local y axis.

#### type PxForceFieldDef

```go
type PxForceFieldDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Techniques
	HasTechniques
}
```

Provides a general container for force fields. Force fields affect physical
objects, such as rigid bodies, and may be instantiated under a physics scene or
a physics model instance.

#### func (*PxForceFieldDef) DefaultInst

```go
func (me *PxForceFieldDef) DefaultInst() (inst *PxForceFieldInst)
```
Returns "the default PxForceFieldInst instance" referencing this PxForceFieldDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*PxForceFieldDef) Init

```go
func (me *PxForceFieldDef) Init()
```
Initialization

#### func (*PxForceFieldDef) NewInst

```go
func (me *PxForceFieldDef) NewInst() (inst *PxForceFieldInst)
```
Creates and returns a new PxForceFieldInst instance referencing this
PxForceFieldDef definition. Any PxForceFieldInst created by this method will
have its Def field readily set to me.

#### type PxForceFieldInst

```go
type PxForceFieldInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *PxForceFieldDef
}
```

Instantiates a force field resource.

#### func (*PxForceFieldInst) EnsureDef

```go
func (me *PxForceFieldInst) EnsureDef() *PxForceFieldDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct PxForceFieldDef
according to the current me.DefRef value (by searching AllPxForceFieldDefLibs).
Then returns me.Def. (Note, every PxForceFieldInst's Def is nil initially,
unless it was created via PxForceFieldDef.NewInst().)

#### func (*PxForceFieldInst) Init

```go
func (me *PxForceFieldInst) Init()
```
Initialization

#### type PxMaterial

```go
type PxMaterial struct {
	//	An inline physics material definition.
	Def *PxMaterialDef

	//	Instantiation of a previously defined physics material.
	Inst *PxMaterialInst
}
```

Describes the physical surface properties for a rigid body or one of its shapes.
Either Def or Inst (or none), but not both, should be specified.

#### type PxMaterialDef

```go
type PxMaterialDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Techniques
	HasTechniques

	//	Common-technique profile
	TC struct {
		//	The dynamic friction coefficient.
		DynamicFriction SidFloat

		//	The proportion of the kinetic energy preserved in the impact
		//	(typically ranges from 0.0 to 1.0). Also known as "bounciness" or "elasticity."
		Restitution SidFloat

		//	The static friction coefficient.
		StaticFriction SidFloat
	}
}
```

Defines the physical properties of an object.

#### func (*PxMaterialDef) DefaultInst

```go
func (me *PxMaterialDef) DefaultInst() (inst *PxMaterialInst)
```
Returns "the default PxMaterialInst instance" referencing this PxMaterialDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*PxMaterialDef) Init

```go
func (me *PxMaterialDef) Init()
```
Initialization

#### func (*PxMaterialDef) NewInst

```go
func (me *PxMaterialDef) NewInst() (inst *PxMaterialInst)
```
Creates and returns a new PxMaterialInst instance referencing this PxMaterialDef
definition. Any PxMaterialInst created by this method will have its Def field
readily set to me.

#### type PxMaterialInst

```go
type PxMaterialInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *PxMaterialDef
}
```

Lets a shape specify its surface properties using a previously defined physics
material.

#### func (*PxMaterialInst) EnsureDef

```go
func (me *PxMaterialInst) EnsureDef() *PxMaterialDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct PxMaterialDef
according to the current me.DefRef value (by searching AllPxMaterialDefLibs).
Then returns me.Def. (Note, every PxMaterialInst's Def is nil initially, unless
it was created via PxMaterialDef.NewInst().)

#### func (*PxMaterialInst) Init

```go
func (me *PxMaterialInst) Init()
```
Initialization

#### type PxModelDef

```go
type PxModelDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Contains zero or more rigid bodies participating in this physics model.
	RigidBodies PxRigidBodyDefs

	//	Contains zero or more rigid constraints participating in this physics model.
	RigidConstraints PxRigidConstraintDefs

	//	Child physics models participating in this physics model, with optional property overrides.
	Insts []*PxModelInst
}
```

Allows for building complex combinations of rigid bodies and constraints that
may be instantiated multiple times.

#### func (*PxModelDef) DefaultInst

```go
func (me *PxModelDef) DefaultInst() (inst *PxModelInst)
```
Returns "the default PxModelInst instance" referencing this PxModelDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*PxModelDef) Init

```go
func (me *PxModelDef) Init()
```
Initialization

#### func (*PxModelDef) NewInst

```go
func (me *PxModelDef) NewInst() (inst *PxModelInst)
```
Creates and returns a new PxModelInst instance referencing this PxModelDef
definition. Any PxModelInst created by this method will have its Def field
readily set to me.

#### type PxModelInst

```go
type PxModelInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *PxModelDef

	//	Points to the Id of a node in the visual scene. This allows a physics model to be instantiated
	//	under a specific transform node, which will dictate the initial position and orientation,
	//	and could be animated to influence kinematic rigid bodies. Optional.
	//	By default, the physics model is instantiated under the world, rather than a specific transform node.
	//	This parameter is only meaningful when the parent element of the current physics model is a physics scene.
	Parent RefId

	//	Zero or more force fields influencing this physics model.
	ForceFields []*PxForceFieldInst

	//	Contains instances of those rigid bodies included in the instantiated physics model that should
	//	have some properties overridden, or should be linked with transform nodes in the visual scene.
	RigidBodies []*PxRigidBodyInst

	//	Contains instances of those rigid constraints included in the instantiated
	//	physics model that should have some properties overridden.
	RigidConstraints []*PxRigidConstraintInst
}
```

Embeds a physics model inside another physics model or instantiates a physics
model within a physics scene.

#### func (*PxModelInst) AccessField

```go
func (me *PxModelInst) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "Parent".

#### func (*PxModelInst) EnsureDef

```go
func (me *PxModelInst) EnsureDef() *PxModelDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct PxModelDef
according to the current me.DefRef value (by searching AllPxModelDefLibs). Then
returns me.Def. (Note, every PxModelInst's Def is nil initially, unless it was
created via PxModelDef.NewInst().)

#### func (*PxModelInst) Init

```go
func (me *PxModelInst) Init()
```
Initialization

#### type PxRigidBodyCommon

```go
type PxRigidBodyCommon struct {
	//	Specifies whether this rigid body is movable. Defaults to true.
	Dynamic SidBool

	//	If set, specifies the total mass of this rigid body.
	Mass *SidFloat

	//	Zero or more TransformKindRotate and/or TransformKindTranslate transformations defining the
	//	center and orientation of mass of the rigid-body relative to the local origin of the "root" shape.
	//	This makes the off-diagonal elements of the inertia tensor (products of inertia) all 0
	//	and allows us to just store the diagonal elements (moments of inertia).
	MassFrame []*Transform

	//	The diagonal elements of the inertia tensor (moments of inertia),
	//	represented in the local frame of the center of mass.
	Inertia *SidFloat3

	//	Describes the physical surface properties for this rigid body.
	Material PxMaterial

	//	Zero or more shapes for collision detection.
	Shapes []*PxShape
}
```

Common-technique profile for rigid body definitions and instances.

#### func (*PxRigidBodyCommon) AccessField

```go
func (me *PxRigidBodyCommon) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "Dynamic".

#### type PxRigidBodyDef

```go
type PxRigidBodyDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Sid
	HasSid

	//	Techniques
	HasTechniques

	//	Common-technique profile
	TC struct {
		//	Dynamic, Mass, MassFrame, Inertia, Material, Shapes
		PxRigidBodyCommon
	}
}
```

Describes simulated bodies that do not deform.

#### func (*PxRigidBodyDef) Init

```go
func (me *PxRigidBodyDef) Init()
```
Initialization

#### type PxRigidBodyDefs

```go
type PxRigidBodyDefs map[string]*PxRigidBodyDef
```

A hash-table of rigid body definitions mapped to their scoped identifiers.

#### type PxRigidBodyInst

```go
type PxRigidBodyInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *PxRigidBodyDef

	//	Techniques
	HasTechniques

	//	Refers to the NodeDef influenced by this rigid body instance.
	TargetNode RefId

	//	Common-technique profile
	TC struct {
		//	Dynamic, Mass, MassFrame, Inertia, Material, Shapes
		PxRigidBodyCommon

		//	Initial spin or angular velocity, also known as the axis of rotation,
		//	with a magnitude equal to the rate of rotation in radians per second.
		AngularVelocity unum.Vec3

		//	Initial linear velocity.
		LinearVelocity unum.Vec3
	}
}
```

Instantiates a rigid-body resource.

#### func (*PxRigidBodyInst) AccessField

```go
func (me *PxRigidBodyInst) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "TargetNode".

#### func (*PxRigidBodyInst) Init

```go
func (me *PxRigidBodyInst) Init()
```
Initialization

#### type PxRigidConstraintAttachment

```go
type PxRigidConstraintAttachment struct {
	//	Extras
	HasExtras

	//	Refers to a RigidBodyDef or NodeDef.
	RigidBody RefSid

	//	Zero or more translation and/or rotation transformations:
	//	The position of a TransformKindTranslate Transform indicates
	//	the attachment point on the corresponding RigidBodyDef.
	//	The orientation of a TransformKindRotate Transform indicates
	//	the alignment of the joint frame for that RigidBodyDef.
	Transforms []*Transform
}
```

Defines an attachment frame (or attachment frame of reference), to a rigid body
or a node, within a rigid constraint.

#### type PxRigidConstraintDef

```go
type PxRigidConstraintDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Sid
	HasSid

	//	Techniques
	HasTechniques

	//	Defines the attachment frame of reference (to a rigid body or a node) within this rigid constraint.
	RefAttachment PxRigidConstraintAttachment

	//	Defines an attachment frame (to a rigid body or a node) within this rigid constraint.
	Attachment PxRigidConstraintAttachment

	//	Common-technique profile
	TC struct {
		//	If false, this rigid constraint doesn't exert any force or influence on the rigid bodies.
		Enabled SidBool

		//	If true, the attached rigid bodies may interpenetrate.
		Interpenetrate SidBool

		//	Degrees of freedom and ranges.
		Limits struct {
			//	Describes the angular limits along each rotation axis in degrees.
			Angular *PxRigidConstraintLimit

			//	Describes linear (translational) limits along each axis.
			Linear *PxRigidConstraintLimit
		}
		//	Spring is based based on either distance (Linear) or angle (Angular), or both.
		Spring struct {
			//	Angle-based spring.
			Angular *PxRigidConstraintSpring

			//	Distance-based spring.
			Linear *PxRigidConstraintSpring
		}
	}
}
```

Constrains rigid bodies to each other or to the world.

#### func (*PxRigidConstraintDef) Init

```go
func (me *PxRigidConstraintDef) Init()
```
Initialization

#### type PxRigidConstraintDefs

```go
type PxRigidConstraintDefs map[string]*PxRigidConstraintDef
```

A hash-table of rigid constraint definitions mapped to their scoped identifiers.

#### type PxRigidConstraintInst

```go
type PxRigidConstraintInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *PxRigidConstraintDef
}
```

Instantiates a rigid constraint resource.

#### func (*PxRigidConstraintInst) Init

```go
func (me *PxRigidConstraintInst) Init()
```
Initialization

#### type PxRigidConstraintLimit

```go
type PxRigidConstraintLimit struct {
	//	Lower bounds for this limit across all 3 axes.
	Min SidVec3

	//	Upper bounds for this limit across all 3 axes.
	Max SidVec3
}
```

Degrees of freedom and ranges.

#### func (*PxRigidConstraintLimit) AccessField

```go
func (me *PxRigidConstraintLimit) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "Min", "Max".

#### type PxRigidConstraintSpring

```go
type PxRigidConstraintSpring struct {
	//	Also called spring coefficient.
	//	Has units of force/distance (for Linear) or force/angle (for Angular).
	Stiffness SidFloat

	//	How this spring is damped.
	Damping SidFloat

	//	Target value for this spring.
	TargetValue SidFloat
}
```

Spring is based based on either distance (Linear) or angle (Angular), or both.

#### func  NewPxRigidConstraintSpring

```go
func NewPxRigidConstraintSpring() (me *PxRigidConstraintSpring)
```
Constructor

#### func (*PxRigidConstraintSpring) AccessField

```go
func (me *PxRigidConstraintSpring) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "Stiffness", "Damping",
"TargetValue".

#### type PxSceneDef

```go
type PxSceneDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Techniques
	HasTechniques

	//	Force fields influencing this physics scene.
	ForceFields []*PxForceFieldInst

	//	Refers to the rigid bodies and constraints participating in this scene.
	Models []*PxModelInst

	//	Common-technique profile
	TC struct {
		//	If set, a vector representation of this physics scene's gravity force field.
		//	It is given as a denormalized direction vector of three floating-point values that
		//	indicate both the magnitude and direction of acceleration caused by the field.
		Gravity *SidVec3

		//	If set, the integration time step, measured in seconds, of the physics scene.
		//	This value is engine-specific. If omitted, the physics engine's default is used.
		TimeStep *SidFloat
	}
}
```

Specifies an environment in which physical objects are instantiated and
simulated.

#### func (*PxSceneDef) DefaultInst

```go
func (me *PxSceneDef) DefaultInst() (inst *PxSceneInst)
```
Returns "the default PxSceneInst instance" referencing this PxSceneDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*PxSceneDef) Init

```go
func (me *PxSceneDef) Init()
```
Initialization

#### func (*PxSceneDef) NewInst

```go
func (me *PxSceneDef) NewInst() (inst *PxSceneInst)
```
Creates and returns a new PxSceneInst instance referencing this PxSceneDef
definition. Any PxSceneInst created by this method will have its Def field
readily set to me.

#### type PxSceneInst

```go
type PxSceneInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *PxSceneDef
}
```

Instantiates a physics scene resource.

#### func (*PxSceneInst) EnsureDef

```go
func (me *PxSceneInst) EnsureDef() *PxSceneDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct PxSceneDef
according to the current me.DefRef value (by searching AllPxSceneDefLibs). Then
returns me.Def. (Note, every PxSceneInst's Def is nil initially, unless it was
created via PxSceneDef.NewInst().)

#### func (*PxSceneInst) Init

```go
func (me *PxSceneInst) Init()
```
Initialization

#### type PxShape

```go
type PxShape struct {
	//	Extras
	HasExtras

	//	If true, the mass is distributed along the surface of this shape.
	Hollow SidBool

	//	The mass of this shape.
	Mass *SidFloat

	//	The density of this shape.
	Density *SidFloat

	//	Describes the physical surface properties for this shape.
	Material PxMaterial

	//	Geometry of the shape. At least and at most one of its fields should ever be set.
	Geometry struct {
		//	A flat plane.
		Plane *GeometryBrepPlane

		//	Axis-aligned box.
		Box *GeometryBrepBox

		//	A perfectly round sphere.
		Sphere *GeometryBrepSphere

		//	A cylinder primitive that is centered on its local origin and aligned along its y axis.
		Cylinder *PxCylinder

		//	A capsule primitive that is centered on the local origin and aligned along the y axis.
		Capsule *GeometryBrepCapsule

		//	Refers to a previously defined mesh or spline geometric primitive.
		Inst *GeometryInst
	}
	//	Zero or more TransformKindRotate and/or TransformKindTranslate transformations for the shape.
	Transforms []*Transform
}
```

A component part of a rigid body's collection of bounding shapes for collision
detection.

#### func (*PxShape) AccessField

```go
func (me *PxShape) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "Hollow".

#### type RefId

```go
type RefId string
```

References a resource by its unique identifier (Id).

#### func (RefId) AnimationClipDef

```go
func (me RefId) AnimationClipDef() (def *AnimationClipDef)
```
Searches (all LibAnimationClipDefs contained in AllAnimationClipDefLibs) for the
AnimationClipDef whose Id is referenced by me, returning the first match found.

#### func (RefId) AnimationDef

```go
func (me RefId) AnimationDef() (def *AnimationDef)
```
Searches (all LibAnimationDefs contained in AllAnimationDefLibs) for the
AnimationDef whose Id is referenced by me, returning the first match found.

#### func (RefId) AnimationSampler

```go
func (me RefId) AnimationSampler() (as *AnimationSampler)
```
Searches (all LibAnimationDefs contained in AllAnimationDefLibs) for the
AnimationSampler whose Id is referenced by me, returning the first match found.

#### func (RefId) ArrayInAnimationDef

```go
func (me RefId) ArrayInAnimationDef() *SourceArray
```
Searches (all LibAnimationDefs contained in AllAnimationDefLibs) for the
SourceArray whose Id is referenced by me, returning the first match found.

#### func (RefId) ArrayInAnyDef

```go
func (me RefId) ArrayInAnyDef() (srcArr *SourceArray)
```
Calls the ArrayInAnimationDef(), ArrayInControllerDef() and ArrayInGeometryDef()
methods in that order to find srcArr.

#### func (RefId) ArrayInControllerDef

```go
func (me RefId) ArrayInControllerDef() *SourceArray
```
Searches (all LibControllerDefs contained in AllControllerDefLibs) for the
SourceArray whose Id is referenced by me, returning the first match found.

#### func (RefId) ArrayInGeometryDef

```go
func (me RefId) ArrayInGeometryDef() (sa *SourceArray)
```
Searches (all LibGeometryDefs contained in AllGeometryDefLibs) for the
SourceArray whose Id is referenced by me, returning the first match found.

#### func (RefId) CameraDef

```go
func (me RefId) CameraDef() (def *CameraDef)
```
Searches (all LibCameraDefs contained in AllCameraDefLibs) for the CameraDef
whose Id is referenced by me, returning the first match found.

#### func (RefId) ControllerDef

```go
func (me RefId) ControllerDef() (def *ControllerDef)
```
Searches (all LibControllerDefs contained in AllControllerDefLibs) for the
ControllerDef whose Id is referenced by me, returning the first match found.

#### func (RefId) FormulaDef

```go
func (me RefId) FormulaDef() (def *FormulaDef)
```
Searches (all LibFormulaDefs contained in AllFormulaDefLibs) for the FormulaDef
whose Id is referenced by me, returning the first match found.

#### func (RefId) FxEffectDef

```go
func (me RefId) FxEffectDef() (def *FxEffectDef)
```
Searches (all LibFxEffectDefs contained in AllFxEffectDefLibs) for the
FxEffectDef whose Id is referenced by me, returning the first match found.

#### func (RefId) FxImageDef

```go
func (me RefId) FxImageDef() (def *FxImageDef)
```
Searches (all LibFxImageDefs contained in AllFxImageDefLibs) for the FxImageDef
whose Id is referenced by me, returning the first match found.

#### func (RefId) FxMaterialDef

```go
func (me RefId) FxMaterialDef() (def *FxMaterialDef)
```
Searches (all LibFxMaterialDefs contained in AllFxMaterialDefLibs) for the
FxMaterialDef whose Id is referenced by me, returning the first match found.

#### func (RefId) FxProfile

```go
func (me RefId) FxProfile() (fp *FxProfile)
```
Searches (all LibFxEffectDefs contained in AllFxEffectDefLibs) for the FxProfile
whose Id is referenced by me, returning the first match found.

#### func (RefId) FxTechniqueCommon

```go
func (me RefId) FxTechniqueCommon() *FxTechniqueCommon
```
Searches (all LibFxEffectDefs contained in AllFxEffectDefLibs) for the
FxTechniqueCommon whose Id is referenced by me, returning the first match found.

#### func (RefId) FxTechniqueGlsl

```go
func (me RefId) FxTechniqueGlsl() (t *FxTechniqueGlsl)
```
Searches (all LibFxEffectDefs contained in AllFxEffectDefLibs) for the
FxTechniqueGlsl whose Id is referenced by me, returning the first match found.

#### func (RefId) GeometryBrepEdges

```go
func (me RefId) GeometryBrepEdges() *GeometryBrepEdges
```
Searches (all LibGeometryDefs contained in AllGeometryDefLibs) for the
GeometryBrepEdges whose Id is referenced by me, returning the first match found.

#### func (RefId) GeometryBrepFaces

```go
func (me RefId) GeometryBrepFaces() *GeometryBrepFaces
```
Searches (all LibGeometryDefs contained in AllGeometryDefLibs) for the
GeometryBrepFaces whose Id is referenced by me, returning the first match found.

#### func (RefId) GeometryBrepPcurves

```go
func (me RefId) GeometryBrepPcurves() *GeometryBrepPcurves
```
Searches (all LibGeometryDefs contained in AllGeometryDefLibs) for the
GeometryBrepPcurves whose Id is referenced by me, returning the first match
found.

#### func (RefId) GeometryBrepShells

```go
func (me RefId) GeometryBrepShells() *GeometryBrepShells
```
Searches (all LibGeometryDefs contained in AllGeometryDefLibs) for the
GeometryBrepShells whose Id is referenced by me, returning the first match
found.

#### func (RefId) GeometryBrepSolids

```go
func (me RefId) GeometryBrepSolids() *GeometryBrepSolids
```
Searches (all LibGeometryDefs contained in AllGeometryDefLibs) for the
GeometryBrepSolids whose Id is referenced by me, returning the first match
found.

#### func (RefId) GeometryBrepWires

```go
func (me RefId) GeometryBrepWires() *GeometryBrepWires
```
Searches (all LibGeometryDefs contained in AllGeometryDefLibs) for the
GeometryBrepWires whose Id is referenced by me, returning the first match found.

#### func (RefId) GeometryDef

```go
func (me RefId) GeometryDef() (def *GeometryDef)
```
Searches (all LibGeometryDefs contained in AllGeometryDefLibs) for the
GeometryDef whose Id is referenced by me, returning the first match found.

#### func (RefId) GeometryMesh

```go
func (me RefId) GeometryMesh() (gm *GeometryMesh)
```
Searches (all LibGeometryDefs contained in AllGeometryDefLibs) for the
GeometryDef whose Id is referenced by me, returning the Mesh of the first match
found.

#### func (RefId) GeometryVertices

```go
func (me RefId) GeometryVertices() *GeometryVertices
```
Searches (all LibGeometryDefs contained in AllGeometryDefLibs) for the
GeometryVertices whose Id is referenced by me, returning the first match found.

#### func (RefId) KxArticulatedSystemDef

```go
func (me RefId) KxArticulatedSystemDef() (def *KxArticulatedSystemDef)
```
Searches (all LibKxArticulatedSystemDefs contained in
AllKxArticulatedSystemDefLibs) for the KxArticulatedSystemDef whose Id is
referenced by me, returning the first match found.

#### func (RefId) KxJointDef

```go
func (me RefId) KxJointDef() (def *KxJointDef)
```
Searches (all LibKxJointDefs contained in AllKxJointDefLibs) for the KxJointDef
whose Id is referenced by me, returning the first match found.

#### func (RefId) KxModelDef

```go
func (me RefId) KxModelDef() (def *KxModelDef)
```
Searches (all LibKxModelDefs contained in AllKxModelDefLibs) for the KxModelDef
whose Id is referenced by me, returning the first match found.

#### func (RefId) KxSceneDef

```go
func (me RefId) KxSceneDef() (def *KxSceneDef)
```
Searches (all LibKxSceneDefs contained in AllKxSceneDefLibs) for the KxSceneDef
whose Id is referenced by me, returning the first match found.

#### func (RefId) LightDef

```go
func (me RefId) LightDef() (def *LightDef)
```
Searches (all LibLightDefs contained in AllLightDefLibs) for the LightDef whose
Id is referenced by me, returning the first match found.

#### func (RefId) NodeDef

```go
func (me RefId) NodeDef() (def *NodeDef)
```
Searches (all LibNodeDefs contained in AllNodeDefLibs) for the NodeDef whose Id
is referenced by me, returning the first match found.

#### func (RefId) PxForceFieldDef

```go
func (me RefId) PxForceFieldDef() (def *PxForceFieldDef)
```
Searches (all LibPxForceFieldDefs contained in AllPxForceFieldDefLibs) for the
PxForceFieldDef whose Id is referenced by me, returning the first match found.

#### func (RefId) PxMaterialDef

```go
func (me RefId) PxMaterialDef() (def *PxMaterialDef)
```
Searches (all LibPxMaterialDefs contained in AllPxMaterialDefLibs) for the
PxMaterialDef whose Id is referenced by me, returning the first match found.

#### func (RefId) PxModelDef

```go
func (me RefId) PxModelDef() (def *PxModelDef)
```
Searches (all LibPxModelDefs contained in AllPxModelDefLibs) for the PxModelDef
whose Id is referenced by me, returning the first match found.

#### func (RefId) PxSceneDef

```go
func (me RefId) PxSceneDef() (def *PxSceneDef)
```
Searches (all LibPxSceneDefs contained in AllPxSceneDefLibs) for the PxSceneDef
whose Id is referenced by me, returning the first match found.

#### func (RefId) S

```go
func (me RefId) S() string
```
Returns the Id currently referenced by me.

#### func (*RefId) SetIdRef

```go
func (me *RefId) SetIdRef(v string)
```
Modifies the Id currently referenced by me. This is a mere string assignment.

#### func (RefId) SourceInAnimationDef

```go
func (me RefId) SourceInAnimationDef() (s *Source)
```
Searches (all LibAnimationDefs contained in AllAnimationDefLibs) for the Source
whose Id is referenced by me, returning the first match found.

#### func (RefId) SourceInAnyDef

```go
func (me RefId) SourceInAnyDef() (src *Source)
```
Calls the SourceInAnimationDef(), SourceInControllerDef() and
SourceInGeometryDef() methods in that order to find src.

#### func (RefId) SourceInControllerDef

```go
func (me RefId) SourceInControllerDef() (s *Source)
```
Searches (all LibControllerDefs contained in AllControllerDefLibs) for the
Source whose Id is referenced by me, returning the first match found.

#### func (RefId) SourceInGeometryDef

```go
func (me RefId) SourceInGeometryDef() (s *Source)
```
Searches (all LibGeometryDefs contained in AllGeometryDefLibs) for the Source
whose Id is referenced by me, returning the first match found.

#### func (RefId) VisualSceneDef

```go
func (me RefId) VisualSceneDef() (def *VisualSceneDef)
```
Searches (all LibVisualSceneDefs contained in AllVisualSceneDefLibs) for the
VisualSceneDef whose Id is referenced by me, returning the first match found.

#### type RefParam

```go
type RefParam struct {
	//	A parameter reference technically always refers to a Sid.
	RefSid
}
```

References a previously defined parameter.

#### func  NewRefParam

```go
func NewRefParam(paramRef string) (rs *RefParam)
```
Creates and returns a new RefParam initialized with the specified paramRef.

#### func (*RefParam) SetParamRef

```go
func (me *RefParam) SetParamRef(sidRef string)
```
Convenience short-hand for me.RefSid.SetSidRef(sidRef)

#### type RefSid

```go
type RefSid struct {
	//	The Sid path currently referenced.
	//	To be set ONLY through the NewRefSid() constructor or SetSidRef() method!
	S string

	//	The resolved value referenced by this Sid path.
	//	This is always a pointer: so V may be a *SidFloat but it will never be a SidFloat.
	//	To be set ONLY through the Resolve() method! Reset to nil by the SetSidRef() method.
	V interface{}
}
```

References a resource by its scoped identifier (Sid).

#### func  NewRefSid

```go
func NewRefSid(sidRef string) (rs *RefSid)
```
Creates and returns a new RefSid, its S initialized with the specified sidRef.

#### func (*RefSid) Resolve

```go
func (me *RefSid) Resolve(root RefSidRoot, force bool)
```
If me.V is nil or force is true: resolves the Sid path in me.S and sets V to the
result. For possible root arguments, see RefSidRoot. If no match is found for
the full path, V will become nil (rather than, say, a partial-path-match
result-value).

Sid path examples:

    -	foo/bar/doodad
    	either: root is a lib that finds object with Id "foo", which resolves path "bar/doodad"
    	or: root is a non-lib object with Id "foo" and resolves path "bar/doodad"
    -	./bar/doodad
    	root is a non-lib object with its own arbitrary Id and resolves path "bar/doodad"
    -	bar
    	gets rewritten to "./bar", then: see above
    -	foo/bar/doodad.Hollow
    	root resolves foo/bar/doodad, then returns pointer to its Hollow field
    	(if doodad supports named-field access by implementing RefSidFielder)
    -	foo/bar/doodad(2)
    	root resolves foo/bar/doodad, then returns pointer to a "slot" at index 2
    	(if doodad supports indexed-slot access by implementing RefSidIndexer)

#### func (*RefSid) SetSidRef

```go
func (me *RefSid) SetSidRef(sidRef string)
```
Sets S to sidRef and resets V to nil.

#### type RefSidFielder

```go
type RefSidFielder interface {
	AccessField(fieldName string) interface{}
}
```

Implemented by select types that embed HasSid to aid resolving Sid paths with a
tailing named-field accessor, as in "some/sid/path.fieldName".

#### type RefSidIndexer

```go
type RefSidIndexer interface {
	AccessIndex(i, j int) interface{}
}
```

Implemented by select types that embed HasSid to aid resolving Sid paths with a
tailing indexed-slot accessor, as in "some/sid/path(2)".

#### type RefSidRoot

```go
type RefSidRoot interface {
	// contains filtered or unexported methods
}
```

This interface needs to be passed to the RefSid.Resolve() method to resolve a
Sid path: Implemented by almost all "LibFooDefs" types, plus all types that
embed HasId and directly or indirectly lead to fields of types that embed HasSid
-- this includes almost all "FooDef" types.

#### type Scene

```go
type Scene struct {
	//	Extras
	HasExtras

	//	Embodies the entire set of information that can be visualized from the contents of a resource.
	Visual *VisualSceneInst

	//	Embodies the entire set of information that can be articulated kinematically from a resource.
	Kinematics *KxSceneInst

	//	Specifies an environment in which physical objects are instantiated and simulated.
	Physics []*PxSceneInst
}
```

Declares a complete, self-contained base of a scene hierarchy or scene graph.

#### type SidBool

```go
type SidBool struct {
	//	Sid
	HasSid

	//	The value.
	B bool
}
```

A bool value that has a scoped identifier.

#### type SidFloat

```go
type SidFloat struct {
	//	Sid
	HasSid

	//	The value.
	F float64
}
```

A float64 value that has a scoped identifier.

#### func  SidF

```go
func SidF(f float64) (sf *SidFloat)
```
Returns a ScopedFloat with the specified value and no Sid.

#### type SidFloat3

```go
type SidFloat3 struct {
	//	Sid
	HasSid

	//	The values.
	F Float3
}
```

Three float64 values that have a scoped identifier.

#### func (*SidFloat3) AccessIndex

```go
func (me *SidFloat3) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices.

#### type SidString

```go
type SidString struct {
	//	Sid
	HasSid

	//	The value.
	S string
}
```

A string value that has a scoped identifier.

#### type SidVec3

```go
type SidVec3 struct {
	//	Sid
	HasSid

	//	X, Y, Z
	unum.Vec3
}
```

A 3D vector that has a scoped identifier.

#### func (*SidVec3) AccessField

```go
func (me *SidVec3) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "X", "Y", "Z".

#### func (*SidVec3) AccessIndex

```go
func (me *SidVec3) AccessIndex(i, _ int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional indices 0 through 2.

#### type Source

```go
type Source struct {
	//	Id
	HasId

	//	Name
	HasName

	//	Asset
	HasAsset

	//	Techniques
	HasTechniques

	//	The data array of this Source.
	Array SourceArray

	//	Common-technique profile
	TC struct {
		//	Describes a stream of values from this array data source.
		Accessor *SourceAccessor
	}
}
```

Declares a data repository that provides values according to the semantics of an
Input that refers to it.

#### type SourceAccessor

```go
type SourceAccessor struct {
	//	The number of times the array is accessed. Required.
	Count uint64

	//	The index of the first value to be read from the array. The default is 0. Optional.
	Offset uint64

	//	The Id of the array to access. Required.
	Source RefId

	//	The number of values that are to be considered a unit during each access to the array.
	//	The default is 1, indicating that a single value is accessed. Optional.
	Stride uint64

	//	The number and order of Params define the output of the accessor.
	//	Parameters are bound to values in the order in which both are specified. No reordering of the data can occur.
	//	A Param without a Name indicates that the value is not part of the output, so the Param is unbound.
	Params []*Param
}
```

Describes a stream of values from an array data source.

#### func  NewSourceAccessor

```go
func NewSourceAccessor() (me *SourceAccessor)
```
Allocates, initializes and returns a new SourceAccessor.

#### type SourceArray

```go
type SourceArray struct {
	//	Id
	HasId

	//	Name
	HasName

	//	A slice into the array of bools that this Source represents, if any.
	Bools []bool

	//	A slice into the array of floats that this Source represents, if any.
	Floats []float64

	//	A slice into the array of RefIds that this Source represents, if any.
	IdRefs []string

	//	A slice into the array of ints that this Source represents, if any.
	Ints []int64

	//	A slice into the array of names that this Source represents, if any.
	Names []string

	//	A slice into the array of RefSids that this Source represents, if any.
	SidRefs []string

	//	A slice into the array of tokens that this Source represents, if any.
	Tokens []string
}
```

The data array of a Source. Of all its []slice fields, only ONE should ever be
non-nil/non-empty at any time.

#### type Sources

```go
type Sources map[string]*Source
```

A hash-table of Sources, each keyed with its Id.

#### type Technique

```go
type Technique struct {
	//	The type of profile. This is a vendor-defined character string
	//	that indicates the platform or capability target for the technique.
	Profile string

	//	Arbitrary XML content or meta-data for this Technique.
	Data string
}
```

Declares platform-specific or program-specific information used to process some
portion of the content.

#### type Transform

```go
type Transform struct {
	//	Sid
	HasSid

	//	The type of this transformation (rotation, skewing, scaling, translation, "look-at", or matrix).
	//	Must be one of the TransformKind* enumerated constants.
	Kind TransformKind

	//	Contains one or more vectors and values representing this transformation.
	//	For TransformKindLookat:
	//		9 values representing three 3D vectors (eye position, interest point, up-axis).
	//	For TransformKindMatrix:
	//		16 values representing one column-major 4x4 matrix.
	//	For TransformKindSkew:
	//		7 values -- one angle in degrees, then two 3D vectors for the axes of rotation and translation.
	//	For TransformKindRotate:
	//		4 values -- one 3D vector specifying the axis of rotation, then one angle in degrees.
	//	For TransformKindTranslate or TransformKindScale:
	//		3 values representing a single 3D vector.
	F []float64
}
```

Represents a single transformation of a specific kind.

#### func (*Transform) AccessField

```go
func (me *Transform) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "X", "Y", "Z" (mapped to
the first three values in me.F).

#### func (*Transform) AccessIndex

```go
func (me *Transform) AccessIndex(i, j int) interface{}
```
RefSidIndexer implementation. Supports one-dimensional and two-dimensional
indices.

#### type TransformKind

```go
type TransformKind int
```

Categorizes the kind of a Transform.

```go
const (
	//	A position and orientation transformation suitable for aiming a camera.
	TransformKindLookat TransformKind = iota + 1

	//	A transformation that embodies mathematical changes to points within a coordinate system
	//	or the coordinate system itself.
	TransformKindMatrix

	//	A transformation that specifies how to rotate an object around an axis.
	TransformKindRotate

	//	A transformation that specifies how to deform an object along one axis.
	TransformKindSkew

	//	A transformation that specifies how to change an object's size.
	TransformKindScale

	//	A transformation that changes the position of an object in a local coordinate system.
	TransformKindTranslate
)
```

#### type VisualSceneDef

```go
type VisualSceneDef struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	A scene graph containing nodes of visual information and related data.
	Nodes []*NodeDef

	//	Specifies how to evaluate this visual scene.
	Evaluations []*VisualSceneEvaluation
}
```

Embodies the entire set of information that can be visualized from the contents
of a resource.

#### func (*VisualSceneDef) DefaultInst

```go
func (me *VisualSceneDef) DefaultInst() (inst *VisualSceneInst)
```
Returns "the default VisualSceneInst instance" referencing this VisualSceneDef
definition. That instance is created once when this method is first called on
me, and will have its Def field readily set to me.

#### func (*VisualSceneDef) Init

```go
func (me *VisualSceneDef) Init()
```
Initialization

#### func (*VisualSceneDef) NewInst

```go
func (me *VisualSceneDef) NewInst() (inst *VisualSceneInst)
```
Creates and returns a new VisualSceneInst instance referencing this
VisualSceneDef definition. Any VisualSceneInst created by this method will have
its Def field readily set to me.

#### type VisualSceneEvaluation

```go
type VisualSceneEvaluation struct {
	//	Id, Name, Asset, Extras
	BaseDef

	//	Sid
	HasSid

	//	Whether evaluation is enabled. Disabling evaluation can be useful for debugging.
	Disabled bool

	//	Describes effects passes to render a scene.
	RenderPasses []*VisualSceneRendering
}
```

Declares information specifying how to evaluate a visual scene.

#### func (*VisualSceneEvaluation) AccessField

```go
func (me *VisualSceneEvaluation) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "Disabled".

#### type VisualSceneInst

```go
type VisualSceneInst struct {
	//	Sid, Name, Extras, DefRef
	BaseInst

	//	A pointer to the resource definition referenced by this instance.
	//	Is nil by default (unless created via Def.NewInst()) and meant to be set ONLY by
	//	the EnsureDef() method (which uses BaseInst.DefRef to find it).
	Def *VisualSceneDef
}
```

Instantiates a visual scene resource.

#### func (*VisualSceneInst) EnsureDef

```go
func (me *VisualSceneInst) EnsureDef() *VisualSceneDef
```
If me is "dirty" or me.Def is nil, sets me.Def to the correct VisualSceneDef
according to the current me.DefRef value (by searching AllVisualSceneDefLibs).
Then returns me.Def. (Note, every VisualSceneInst's Def is nil initially, unless
it was created via VisualSceneDef.NewInst().)

#### func (*VisualSceneInst) Init

```go
func (me *VisualSceneInst) Init()
```
Initialization

#### type VisualSceneRendering

```go
type VisualSceneRendering struct {
	//	Name
	HasName

	//	Sid
	HasSid

	//	Extras
	HasExtras

	//	Refers to a NodeDef that contains a camera describing
	//	the viewpoint from which to render this compositing step. Optional.
	CameraNode RefId

	//	Specifies which layer or layers to render in this compositing step while evaluating the scene.
	Layers Layers

	//	If set, specifies which effect to render in this compositing step while evaluating the scene.
	MaterialInst *VisualSceneRenderingMaterialInst
}
```

Describes one effect pass to evaluate a scene.

#### func  NewVisualSceneRendering

```go
func NewVisualSceneRendering() (me *VisualSceneRendering)
```
Constructor

#### func (*VisualSceneRendering) AccessField

```go
func (me *VisualSceneRendering) AccessField(fn string) interface{}
```
RefSidFielder implementation. Supported field names: "CameraNode".

#### type VisualSceneRenderingMaterialInst

```go
type VisualSceneRenderingMaterialInst struct {
	//	Extras
	HasExtras

	//	Binds values to effect parameters upon instantiation.
	Bindings []*FxBinding

	//	Target specific techniques and passes inside a material
	//	rather than having to split the effects techniques and passes into multiple effects.
	OverrideTechnique struct {
		//	Specifies the Sid of a Technique
		Ref RefSid

		//	Specifies the Sid of one FxPass to execute.
		//	If not specified, then all of the Technique's passes are used.
		Pass RefSid
	}
}
```

Instantiates a material resource for a screen effect.

--
**godocdown** http://github.com/robertkrimen/godocdown
