package collimp

import (
	xmlx "github.com/go-forks/go-pkg-xmlx"

	cdom "github.com/metaleap/go-collada/dom"
)

func libs_animation_clips(xn *xmlx.Node) {
	var (
		lib *cdom.LibAnimationClipDefs
		def *cdom.AnimationClipDef
		id  string
	)
	for _, ln := range xcns(xn, "library_animation_clips") {
		id = xas(ln, "id")
		if lib = cdom.AllAnimationClipDefLibs[id]; lib == nil {
			lib = cdom.AllAnimationClipDefLibs.AddNew(id)
		}
		for _, def = range objs_AnimationClipDef(ln, "animation_clip") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_animations(xn *xmlx.Node) {
	var (
		lib *cdom.LibAnimationDefs
		def *cdom.AnimationDef
		id  string
	)
	for _, ln := range xcns(xn, "library_animations") {
		id = xas(ln, "id")
		if lib = cdom.AllAnimationDefLibs[id]; lib == nil {
			lib = cdom.AllAnimationDefLibs.AddNew(id)
		}
		for _, def = range objs_AnimationDef(ln, "animation") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_cameras(xn *xmlx.Node) {
	var (
		lib *cdom.LibCameraDefs
		def *cdom.CameraDef
		id  string
	)
	for _, ln := range xcns(xn, "library_cameras") {
		id = xas(ln, "id")
		if lib = cdom.AllCameraDefLibs[id]; lib == nil {
			lib = cdom.AllCameraDefLibs.AddNew(id)
		}
		for _, def = range objs_CameraDef(ln, "camera") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_controllers(xn *xmlx.Node) {
	var (
		lib *cdom.LibControllerDefs
		def *cdom.ControllerDef
		id  string
	)
	for _, ln := range xcns(xn, "library_controllers") {
		id = xas(ln, "id")
		if lib = cdom.AllControllerDefLibs[id]; lib == nil {
			lib = cdom.AllControllerDefLibs.AddNew(id)
		}
		for _, def = range objs_ControllerDef(ln, "controller") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_formulas(xn *xmlx.Node) {
	var (
		lib *cdom.LibFormulaDefs
		def *cdom.FormulaDef
		id  string
	)
	for _, ln := range xcns(xn, "library_formulas") {
		id = xas(ln, "id")
		if lib = cdom.AllFormulaDefLibs[id]; lib == nil {
			lib = cdom.AllFormulaDefLibs.AddNew(id)
		}
		for _, def = range objs_FormulaDef(ln, "formula") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_geometries(xn *xmlx.Node) {
	var (
		lib *cdom.LibGeometryDefs
		def *cdom.GeometryDef
		id  string
	)
	for _, ln := range xcns(xn, "library_geometries") {
		id = xas(ln, "id")
		if lib = cdom.AllGeometryDefLibs[id]; lib == nil {
			lib = cdom.AllGeometryDefLibs.AddNew(id)
		}
		for _, def = range objs_GeometryDef(ln, "geometry") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_lights(xn *xmlx.Node) {
	var (
		lib *cdom.LibLightDefs
		def *cdom.LightDef
		id  string
	)
	for _, ln := range xcns(xn, "library_lights") {
		id = xas(ln, "id")
		if lib = cdom.AllLightDefLibs[id]; lib == nil {
			lib = cdom.AllLightDefLibs.AddNew(id)
		}
		for _, def = range objs_LightDef(ln, "light") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_nodes(xn *xmlx.Node) {
	var (
		lib *cdom.LibNodeDefs
		def *cdom.NodeDef
		id  string
	)
	for _, ln := range xcns(xn, "library_nodes") {
		id = xas(ln, "id")
		if lib = cdom.AllNodeDefLibs[id]; lib == nil {
			lib = cdom.AllNodeDefLibs.AddNew(id)
		}
		for _, def = range objs_NodeDef(ln, "node") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_visual_scenes(xn *xmlx.Node) {
	var (
		lib *cdom.LibVisualSceneDefs
		def *cdom.VisualSceneDef
		id  string
	)
	for _, ln := range xcns(xn, "library_visual_scenes") {
		id = xas(ln, "id")
		if lib = cdom.AllVisualSceneDefLibs[id]; lib == nil {
			lib = cdom.AllVisualSceneDefLibs.AddNew(id)
		}
		for _, def = range objs_VisualSceneDef(ln, "visual_scene") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_force_fields(xn *xmlx.Node) {
	var (
		lib *cdom.LibPxForceFieldDefs
		def *cdom.PxForceFieldDef
		id  string
	)
	for _, ln := range xcns(xn, "library_force_fields") {
		id = xas(ln, "id")
		if lib = cdom.AllPxForceFieldDefLibs[id]; lib == nil {
			lib = cdom.AllPxForceFieldDefLibs.AddNew(id)
		}
		for _, def = range objs_PxForceFieldDef(ln, "force_field") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_physics_materials(xn *xmlx.Node) {
	var (
		lib *cdom.LibPxMaterialDefs
		def *cdom.PxMaterialDef
		id  string
	)
	for _, ln := range xcns(xn, "library_physics_materials") {
		id = xas(ln, "id")
		if lib = cdom.AllPxMaterialDefLibs[id]; lib == nil {
			lib = cdom.AllPxMaterialDefLibs.AddNew(id)
		}
		for _, def = range objs_PxMaterialDef(ln, "physics_material") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_physics_models(xn *xmlx.Node) {
	var (
		lib *cdom.LibPxModelDefs
		def *cdom.PxModelDef
		id  string
	)
	for _, ln := range xcns(xn, "library_physics_models") {
		id = xas(ln, "id")
		if lib = cdom.AllPxModelDefLibs[id]; lib == nil {
			lib = cdom.AllPxModelDefLibs.AddNew(id)
		}
		for _, def = range objs_PxModelDef(ln, "physics_model") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_physics_scenes(xn *xmlx.Node) {
	var (
		lib *cdom.LibPxSceneDefs
		def *cdom.PxSceneDef
		id  string
	)
	for _, ln := range xcns(xn, "library_physics_scenes") {
		id = xas(ln, "id")
		if lib = cdom.AllPxSceneDefLibs[id]; lib == nil {
			lib = cdom.AllPxSceneDefLibs.AddNew(id)
		}
		for _, def = range objs_PxSceneDef(ln, "physics_scene") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_effects(xn *xmlx.Node) {
	var (
		lib *cdom.LibFxEffectDefs
		def *cdom.FxEffectDef
		id  string
	)
	for _, ln := range xcns(xn, "library_effects") {
		id = xas(ln, "id")
		if lib = cdom.AllFxEffectDefLibs[id]; lib == nil {
			lib = cdom.AllFxEffectDefLibs.AddNew(id)
		}
		for _, def = range objs_FxEffectDef(ln, "effect") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_images(xn *xmlx.Node) {
	var (
		lib *cdom.LibFxImageDefs
		def *cdom.FxImageDef
		id  string
	)
	for _, ln := range xcns(xn, "library_images") {
		id = xas(ln, "id")
		if lib = cdom.AllFxImageDefLibs[id]; lib == nil {
			lib = cdom.AllFxImageDefLibs.AddNew(id)
		}
		for _, def = range objs_FxImageDef(ln, "image") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_materials(xn *xmlx.Node) {
	var (
		lib *cdom.LibFxMaterialDefs
		def *cdom.FxMaterialDef
		id  string
	)
	for _, ln := range xcns(xn, "library_materials") {
		id = xas(ln, "id")
		if lib = cdom.AllFxMaterialDefLibs[id]; lib == nil {
			lib = cdom.AllFxMaterialDefLibs.AddNew(id)
		}
		for _, def = range objs_FxMaterialDef(ln, "material") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_articulated_systems(xn *xmlx.Node) {
	var (
		lib *cdom.LibKxArticulatedSystemDefs
		def *cdom.KxArticulatedSystemDef
		id  string
	)
	for _, ln := range xcns(xn, "library_articulated_systems") {
		id = xas(ln, "id")
		if lib = cdom.AllKxArticulatedSystemDefLibs[id]; lib == nil {
			lib = cdom.AllKxArticulatedSystemDefLibs.AddNew(id)
		}
		for _, def = range objs_KxArticulatedSystemDef(ln, "articulated_system") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_joints(xn *xmlx.Node) {
	var (
		lib *cdom.LibKxJointDefs
		def *cdom.KxJointDef
		id  string
	)
	for _, ln := range xcns(xn, "library_joints") {
		id = xas(ln, "id")
		if lib = cdom.AllKxJointDefLibs[id]; lib == nil {
			lib = cdom.AllKxJointDefLibs.AddNew(id)
		}
		for _, def = range objs_KxJointDef(ln, "joint") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_kinematics_models(xn *xmlx.Node) {
	var (
		lib *cdom.LibKxModelDefs
		def *cdom.KxModelDef
		id  string
	)
	for _, ln := range xcns(xn, "library_kinematics_models") {
		id = xas(ln, "id")
		if lib = cdom.AllKxModelDefLibs[id]; lib == nil {
			lib = cdom.AllKxModelDefLibs.AddNew(id)
		}
		for _, def = range objs_KxModelDef(ln, "kinematics_model") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_kinematics_scenes(xn *xmlx.Node) {
	var (
		lib *cdom.LibKxSceneDefs
		def *cdom.KxSceneDef
		id  string
	)
	for _, ln := range xcns(xn, "library_kinematics_scenes") {
		id = xas(ln, "id")
		if lib = cdom.AllKxSceneDefLibs[id]; lib == nil {
			lib = cdom.AllKxSceneDefLibs.AddNew(id)
		}
		for _, def = range objs_KxSceneDef(ln, "kinematics_scene") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_All(xn *xmlx.Node) {
	libs_animation_clips(xn)
	libs_animations(xn)
	libs_cameras(xn)
	libs_controllers(xn)
	libs_formulas(xn)
	libs_geometries(xn)
	libs_lights(xn)
	libs_nodes(xn)
	libs_visual_scenes(xn)
	libs_force_fields(xn)
	libs_physics_materials(xn)
	libs_physics_models(xn)
	libs_physics_scenes(xn)
	libs_effects(xn)
	libs_images(xn)
	libs_materials(xn)
	libs_articulated_systems(xn)
	libs_joints(xn)
	libs_kinematics_models(xn)
	libs_kinematics_scenes(xn)
}
