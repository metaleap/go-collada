//	The go-collada/dom package provides logic-less base data structures
//	for all resource types used in COLLADA 1.5 or 1.4.1 documents.
//	
//	These data structures only describe resource definitions and instances, but do not provide
//	any specific logic or algorithms, such as physics or visual rendering.
//	
//	Context: every graphics app uses a variety of both simple and complex resource types, including for
//	example geometry definitions, image textures, material and light descriptions and many others.
//	
//	While for all these resource types, a graphics app most likely manages their respective run-time
//	representation in its entirety,	some basic part of that representation is "merely descriptive" and
//	should be readable from or writable to a binary stream ("design-time"). That "merely descriptive"
//	sub-set of resource definitions is fully contained in this go-collada/dom package.
//	
//	This package thus allows for things such as server-side procedural asset generators, networked
//	resource streaming or simple custom asset-import/export/conversion tools, all of which shouldn't
//	have to needlessly depend on the graphics, windowing etc. stacks.
//	
//	NOTE: there are essentially two distinct "modes" or use-cases in which this package will be active:
//	
//	1. in graphics-independent, server-side or non-interactive "toolage", dealing only with the raw resource
//	data where it can be generated, loaded, stored, manipulated at will with no particular repercussions.
//	
//	2. in an interactive graphical app:
//	
//	All "Sync"-related functions pertain to use-case #2, where this package essentially becomes the
//	live repository for all resource definitions loaded, used, or manipulated by the parent package at runtime.
//	So now every image definition in go-collada/dom may have a corresponding GPU-bound texture object in the app,
//	every dom mesh definition may be bound to an app's "MeshBuffer" or some such construction, etc.
//	
//	Structure: generally speaking, all resource types are organized in a consistent fashion as follows ---
//	fully consistent with the COLLADA specification in terminology and resource organization:
//	
//	1. First, there is a FooDef struct for the one-time definition of a unique resource:
//	GeometryDef, FxImageDef, LightDef, FxMaterialDef etc.
//	
//	2. Next, there is a smaller FooInst struct for handling many individual (sometimes parameterized) instantiations of
//	a FooDef: GeometryInst, FxImageInst, LightInst, FxMaterialInst etc.
//	
//	3. Finally, there is a light-weight LibFooDefs struct type (encapsulating a typed hash-table) containing Defs
//	associated with their Id: LibGeometryDefs, LibFxImageDefs, LibLightDefs, LibFxMaterialDefs etc.
//	
//	4. The package also provides a pre-initialized global FooDefs variable for each such Lib type, in simple apps
//	considered the "default / main / only Lib you'll need": GeometryDefs, FxImageDefs, LightDefs, FxMaterialDefs etc.
//	
//	5. For more complex use-cases, you can also organize multiple libs for any given resource type in the global
//	AllFooDefLibs variable, essentially a hash-table of Libs: AllGeometryDefLibs (of type
//	LibsGeometryDef), AllFxImageDefLibs (of type LibsFxImageDef), AllLightDefLibs (of type LibsLightDef),
//	AllFxMaterialDefLibs (of type LibsFxMaterialDef) etc.
//	
//	Any exported types in this package not following the above pattern should be
//	considered "auxiliary helpers" rather than primary / "first-class" resource types.
package cdom

//	Encapsulates a complete, fully self-contained scene graph.
//	Note, resource definition libraries are organized in "globals" (exported package variables)
//	rather than Document struct instances.
type Document struct {
	//	Asset
	HasAsset

	//	Extras
	HasExtras

	//	Describes a complete, fully self-contained scene graph.
	Scene *Scene
}
