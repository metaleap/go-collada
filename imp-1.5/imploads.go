package collimp

import (
	"encoding/hex"
	"strings"

	xmlx "github.com/goforks/go-pkg-xmlx"

	cdom "github.com/go3d/go-collada/dom"
	ugfx "github.com/metaleap/go-util/gfx"
	ustr "github.com/metaleap/go-util/str"
	xsdt "github.com/metaleap/go-xsd/types"
)

func load_Document(xn *xmlx.Node, obj *cdom.Document) {
	obj.Scene = obj_Scene(xn, "scene")
}

func load_SamplerWrapping(xn *xmlx.Node, obj *ugfx.SamplerWrapping) {
	if cn := xcn(xn, "border_color"); cn != nil {
		list_Rgba32(cn, &obj.BorderColor)
	}
	for n, i := range map[string]*ugfx.WrapKind{"wrap_s": &obj.WrapS, "wrap_t": &obj.WrapT, "wrap_p": &obj.WrapP} {
		switch strings.ToUpper(xs(xn, n)) {
		case "BORDER":
			*i = ugfx.WrapKindBorder
		case "CLAMP":
			*i = ugfx.WrapKindClamp
		case "MIRROR":
			*i = ugfx.WrapKindMirror
		case "MIRROR_ONCE":
			*i = ugfx.WrapKindMirrorOnce
		default:
			*i = ugfx.WrapKindRepeat
		}
	}
}

func load_FxEffectDef(xn *xmlx.Node, obj *cdom.FxEffectDef) {
	obj.Annotations = objs_FxAnnotation(xn, "annotate")
	for _, cn := range xn.Children {
		if (cn.Type == xmlx.NT_ELEMENT) && strings.HasPrefix(cn.Name.Local, "profile_") {
			obj.Profiles = append(obj.Profiles, obj_FxProfile(cn, ""))
		}
	}
}

func load_PxCylinder(xn *xmlx.Node, obj *cdom.PxCylinder) {
	if bc := obj_GeometryBrepCylinder(xn, ""); bc != nil {
		obj.GeometryBrepCylinder = *bc
	}
	obj.Height = xf64(xn, "height")
}

func load_Float4x4(xn *xmlx.Node, obj *cdom.Float4x4) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_LightDef(xn *xmlx.Node, obj *cdom.LightDef) {
	if tcn := node_TechCommon(xn); tcn != nil {
		obj.TC.Ambient = obj_LightAmbient(tcn, "ambient")
		obj.TC.Directional = obj_LightDirectional(tcn, "directional")
		obj.TC.Point = obj_LightPoint(tcn, "point")
		obj.TC.Spot = obj_LightSpot(tcn, "spot")
	}
}

func load_SourceArray(xn *xmlx.Node, obj *cdom.SourceArray) {
	if dn := xcn1(xn, "bool_array", "float_array", "IDREF_array", "idref_array", "int_array", "Name_array", "name_array", "SIDREF_array", "sidref_array", "token_array"); dn != nil {
		switch strings.ToLower(dn.Name.Local) {
		case "bool_array":
			obj.Bools = list_Bools(dn)
		case "float_array":
			obj.Floats = list_Floats(dn)
		case "idref_array":
			obj.IdRefs = list_Strings(dn)
		case "int_array":
			obj.Ints = list_Ints(dn)
		case "name_array":
			obj.Names = list_Strings(dn)
		case "sidref_array":
			obj.SidRefs = list_Strings(dn)
		case "token_array":
			obj.Tokens = list_Strings(dn)
		}
	}
}

func load_Source(xn *xmlx.Node, obj *cdom.Source) {
	if tcn := node_TechCommon(xn); tcn != nil {
		obj.TC.Accessor = obj_SourceAccessor(tcn, "accessor")
	}
	if arr := obj_SourceArray(xn, ""); arr != nil {
		obj.Array = *arr
	}
}

func load_FxPass(xn *xmlx.Node, obj *cdom.FxPass) {
	obj.Annotations = objs_FxAnnotation(xn, "annotate")
	obj.Evaluate = obj_FxPassEvaluation(xn, "evaluate")
	obj.Program = obj_FxPassProgram(xn, "program")
	if sn := xcn(xn, "states"); sn != nil {
		for _, scn := range sn.Children {
			if scn.Type == xmlx.NT_ELEMENT {
				obj.States[scn.Name.Local] = obj_FxPassState(scn, "")
			}
		}
	}
}

func load_FxPassEvaluationClearColor(xn *xmlx.Node, obj *cdom.FxPassEvaluationClearColor) {
	obj.Index = xau64(xn, "index")
	list_Rgba32(xn, &obj.Rgba32)
}

func load_GeometryBrepCylinder(xn *xmlx.Node, obj *cdom.GeometryBrepCylinder) {
	if f := obj_Float2(xn, "radius"); f != nil {
		obj.Radii = *f
	}
}

func load_PxModelDef(xn *xmlx.Node, obj *cdom.PxModelDef) {
	obj.Insts = objs_PxModelInst(xn, "instance_physics_model")
	for _, rb := range objs_PxRigidBodyDef(xn, "rigid_body") {
		obj.RigidBodies[rb.Sid] = rb
	}
	for _, rc := range objs_PxRigidConstraintDef(xn, "rigid_constraint") {
		obj.RigidConstraints[rc.Sid] = rc
	}
}

func load_FxCreateCubeInitFrom(xn *xmlx.Node, obj *cdom.FxCreateCubeInitFrom) {
	if i := obj_FxCreateInitFrom(xn, ""); i != nil {
		obj.FxCreateInitFrom = *i
	}
	obj.Face = get_CubeFace(xn)
}

func load_KxModelInst(xn *xmlx.Node, obj *cdom.KxModelInst) {
	obj.Bindings = objs_KxBinding(xn, "bind")
}

func load_GeometryBrepSphere(xn *xmlx.Node, obj *cdom.GeometryBrepSphere) {
	obj.Radius = xf64(xn, "radius")
}

func load_Float4(xn *xmlx.Node, obj *cdom.Float4) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_SourceAccessor(xn *xmlx.Node, obj *cdom.SourceAccessor) {
	obj.Count = xau64(xn, "count")
	obj.Offset = xau64(xn, "offset")
	setIdRef(&obj.Source, xas(xn, "source"))
	if u := xau64p(xn, "stride"); u != nil {
		obj.Stride = *u
	} else {
		obj.Stride = 1
	}
	obj.Params = objs_Param(xn, "param")
}

func load_FxPassEvaluation(xn *xmlx.Node, obj *cdom.FxPassEvaluation) {
	obj.Draw = xs(xn, "draw")
	obj.Color.Clear = obj_FxPassEvaluationClearColor(xn, "color_clear")
	obj.Color.Target = obj_FxPassEvaluationTarget(xn, "color_target")
	obj.Depth.Clear = obj_FxPassEvaluationClearDepth(xn, "depth_clear")
	obj.Depth.Target = obj_FxPassEvaluationTarget(xn, "depth_target")
	obj.Stencil.Clear = obj_FxPassEvaluationClearStencil(xn, "stencil_clear")
	obj.Stencil.Target = obj_FxPassEvaluationTarget(xn, "stencil_target")
}

func load_FxEffectInstTechniqueHint(xn *xmlx.Node, obj *cdom.FxEffectInstTechniqueHint) {
	obj.Platform = xas(xn, "platform")
	obj.Ref = xas(xn, "ref")
	obj.Profile = xas(xn, "profile")
}

func load_FxSamplerFiltering(xn *xmlx.Node, obj *cdom.FxSamplerFiltering) {
	obj.MipBias = xf64(xn, "mip_bias")
	obj.MipMinLevel = uint8(xu64(xn, "mip_min_level"))
	obj.MipMaxLevel = uint8(xu64(xn, "mip_max_level"))
	if u := xu64p(xn, "max_anisotropy"); u != nil {
		obj.MaxAnisotropy = uint32(*u)
	} else {
		obj.MaxAnisotropy = 1
	}
	for n, i := range map[string]*cdom.FxFilterKind{"minfilter": &obj.FilterMin, "magfilter": &obj.FilterMag, "mipfilter": &obj.FilterMip} {
		switch strings.ToUpper(xs(xn, n)) {
		case "ANISOTROPIC":
			*i = cdom.FxFilterKindAnisotropic
		case "NEAREST":
			*i = cdom.FxFilterKindNearest
		case "NONE":
			if n == "mipfilter" {
				*i = cdom.FxFilterKindMipNone
			} else {
				*i = cdom.FxFilterKindNearest
			}
		default:
			*i = cdom.FxFilterKindLinear
		}
	}
}

func load_FxPassProgram(xn *xmlx.Node, obj *cdom.FxPassProgram) {
	obj.Shaders = objs_FxPassProgramShader(xn, "shader")
	obj.BindAttributes = objs_FxPassProgramBindAttribute(xn, "bind_attribute")
	obj.BindUniforms = objs_FxPassProgramBindUniform(xn, "bind_uniform")
}

func load_CameraInst(xn *xmlx.Node, obj *cdom.CameraInst) {
}

func load_PxMaterialInst(xn *xmlx.Node, obj *cdom.PxMaterialInst) {
}

func load_FxImageDef(xn *xmlx.Node, obj *cdom.FxImageDef) {
	obj.Create2D = obj_FxCreate2D(xn, "create_2d")
	obj.Create3D = obj_FxCreate3D(xn, "create_3d")
	obj.CreateCube = obj_FxCreateCube(xn, "create_cube")
	obj.InitFrom = obj_FxImageInitFrom(xn, "init_from")
	if rn := xcn(xn, "renderable"); rn != nil {
		obj.Renderable.Is = true
		obj.Renderable.Shared = xab(rn, "share")
	}
}

func load_Float4x3(xn *xmlx.Node, obj *cdom.Float4x3) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_KxKinematicsAxis(xn *xmlx.Node, obj *cdom.KxKinematicsAxis) {
	var f *cdom.Formula
	obj.Axis.SetSidRef(xas(xn, "axis"))
	for _, cn := range xcns(xn, "formula", "instance_formula") {
		if f = obj_Formula(cn, ""); f != nil {
			obj.Formulas = append(obj.Formulas, *f)
		}
	}
	obj.Limits = obj_KxAxisLimits(xn, "limits")
	obj.Indices = objs_KxAxisIndex(xn, "index")
	if l := obj_ParamOrBool(xn, "locked"); l != nil {
		obj.Locked = *l
	}
	if pb := obj_ParamOrBool(xn, "active"); pb != nil {
		obj.Active = *pb
	} else {
		obj.Active.B = true
	}
}

func load_PxModelInst(xn *xmlx.Node, obj *cdom.PxModelInst) {
	setIdRef(&obj.Parent, xas(xn, "parent"))
	obj.ForceFields = objs_PxForceFieldInst(xn, "instance_force_field")
	obj.RigidBodies = objs_PxRigidBodyInst(xn, "instance_rigid_body")
	obj.RigidConstraints = objs_PxRigidConstraintInst(xn, "instance_rigid_constraint")
}

func load_Transform(xn *xmlx.Node, obj *cdom.Transform) {
	switch xn.Name.Local {
	case "lookat":
		obj.Kind = cdom.TransformKindLookat
	case "matrix":
		obj.Kind = cdom.TransformKindMatrix
	case "rotate":
		obj.Kind = cdom.TransformKindRotate
	case "scale":
		obj.Kind = cdom.TransformKindScale
	case "skew":
		obj.Kind = cdom.TransformKindSkew
	case "translate":
		obj.Kind = cdom.TransformKindTranslate
	}
	if obj.Kind > 0 {
		obj.F = list_Floats(xn)
	}
}

func load_KxJoint(xn *xmlx.Node, obj *cdom.KxJoint) {
	switch xn.Name.Local {
	case "revolute":
		obj.Kind = cdom.KxJointKindRevolute
	case "prismatic":
		obj.Kind = cdom.KxJointKindPrismatic
	}
	if obj.Kind > 0 {
		if xa := xcn(xn, "axis"); xa != nil {
			if sv3 := obj_SidVec3(xa, ""); sv3 != nil {
				obj.Axis.SidVec3 = *sv3
			}
			has_Name(xa, &obj.Axis.HasName)
		}
		obj.Limits = obj_KxJointLimits(xn, "limits")
	}
}

func load_ParamOrInt(xn *xmlx.Node, obj *cdom.ParamOrInt) {
	obj.Param.SetParamRef(xs(xn, "param"))
	obj.I = xi64(xn, "int")
}

func load_ParamDef(xn *xmlx.Node, obj *cdom.ParamDef) {
	for _, cn := range xn.Children {
		if (cn.Type == xmlx.NT_ELEMENT) && !ustr.IsOneOf(cn.Name.Local, "annotate", "semantic", "modifier") {
			obj.Value = xv(cn)
			break
		}
	}
}

func load_Int2x2(xn *xmlx.Node, obj *cdom.Int2x2) {
	arr_Ints(xn, len(obj), func(i int, n int64) {
		obj[i] = n
	})
}

func load_ControllerInst(xn *xmlx.Node, obj *cdom.ControllerInst) {
	obj.BindMaterial = obj_MaterialBinding(xn, "bind_material")
	obj.SkinSkeletons = list_StringsN(xn, "skeleton")
}

func load_GeometryPositioning(xn *xmlx.Node, obj *cdom.GeometryPositioning) {
	obj.Orientations = objs_GeometryBrepOrientation(xn, "orient")
	obj.Origin = xv3(xn, "origin")
}

func load_GeometryBrepCurve(xn *xmlx.Node, obj *cdom.GeometryBrepCurve) {
	if pos := obj_GeometryPositioning(xn, ""); pos != nil {
		obj.Location = *pos
	}
	obj.Element.Line = obj_GeometryBrepLine(xn, "line")
	obj.Element.Circle = obj_GeometryBrepCircle(xn, "circle")
	obj.Element.Ellipse = obj_GeometryBrepEllipse(xn, "ellipse")
	obj.Element.Parabola = obj_GeometryBrepParabola(xn, "parabola")
	obj.Element.Hyperbola = obj_GeometryBrepHyperbola(xn, "hyperbola")
	obj.Element.Nurbs = obj_GeometryBrepNurbs(xn, "nurbs")
}

func load_GeometryBrepNurbs(xn *xmlx.Node, obj *cdom.GeometryBrepNurbs) {
	obj.Degree = xau64(xn, "degree")
	obj.Closed = xab(xn, "closed")
	if cv := obj_GeometryControlVertices(xn, "control_vertices"); cv != nil {
		obj.ControlVertices = *cv
	}
}

func load_Formula(xn *xmlx.Node, obj *cdom.Formula) {
	switch xn.Name.Local {
	case "instance_formula":
		obj.Inst = obj_FormulaInst(xn, "")
	case "formula":
		obj.Def = obj_FormulaDef(xn, "")
	}
}

func load_FormulaDef(xn *xmlx.Node, obj *cdom.FormulaDef) {
	if pf := obj_ParamOrFloat(xn, "target"); pf != nil {
		obj.Target = *pf
	}
	if tcn := node_TechCommon(xn); tcn != nil {
		obj.TC.MathML = tcn.String()
	}
}

func load_LightInst(xn *xmlx.Node, obj *cdom.LightInst) {
}

func load_Asset(xn *xmlx.Node, obj *cdom.Asset) {
	obj.Contributors = objs_AssetContributor(xn, "contributor")
	if cn := xcn(xn, "coverage"); cn != nil {
		obj.Coverage = obj_AssetGeographicLocation(cn, "geographic_location")
	}
	if cn := xcn(xn, "unit"); cn != nil {
		obj.Unit.Meter, obj.Unit.Name = xaf64d(cn, "meter", obj.Unit.Meter), xasd(cn, "name", obj.Unit.Name)
	}
	// obj.Created = xs(xn, "created")
	// obj.Keywords = xs(xn, "keywords")
	// obj.Modified = xs(xn, "modified")
	// obj.Revision = xs(xn, "revision")
	// obj.Subject = xs(xn, "subject")
	// obj.Title = xs(xn, "title")
	if obj.UpAxis = xs(xn, "up_axis"); len(obj.UpAxis) == 0 {
		obj.UpAxis = "Y"
	} else if obj.UpAxis = strings.ToUpper(obj.UpAxis[:1]); (obj.UpAxis != "X") && (obj.UpAxis != "Z") {
		obj.UpAxis = "Y"
	}
}

func load_FxPassProgramShader(xn *xmlx.Node, obj *cdom.FxPassProgramShader) {
	switch strings.ToUpper(xas(xn, "stage")) {
	case "COMPUTE":
		obj.Stage = cdom.FxShaderStageCompute
	case "FRAGMENT":
		obj.Stage = cdom.FxShaderStageFragment
	case "GEOMETRY":
		obj.Stage = cdom.FxShaderStageGeometry
	case "TESSELATION", "TESSELLATION":
		obj.Stage = cdom.FxShaderStageTessellation
	case "VERTEX":
		obj.Stage = cdom.FxShaderStageVertex
	}
	if sn := xcn(xn, "sources"); (obj.Stage > 0) && (sn != nil) {
		pss := cdom.FxPassProgramShaderSources{}
		arr := make([]cdom.FxPassProgramShaderSources, 0, len(sn.Children))
		for _, scn := range sn.Children {
			if pss.S, pss.IsImportRef = scn.Value, (scn.Name.Local == "import"); pss.IsImportRef {
				pss.S = xas(scn, "ref")
			}
			arr = append(arr, pss)
		}
		obj.Sources = arr
	}
}

func load_GeometryBrepCapsule(xn *xmlx.Node, obj *cdom.GeometryBrepCapsule) {
	obj.Height = xf64(xn, "height")
	if v3 := xv3(xn, "radius"); v3 != nil {
		obj.Radii = *v3
	}
}

func load_FxProfileCommon(xn *xmlx.Node, obj *cdom.FxProfileCommon) {
	if tc := obj_FxTechniqueCommon(xn, "technique"); tc != nil {
		obj.Technique = *tc
	}
}

func load_ControllerMorph(xn *xmlx.Node, obj *cdom.ControllerMorph) {
	setIdRef(&obj.Source, xas(xn, "source"))
	obj.Relative = (strings.ToUpper(xas(xn, "method")) == "RELATIVE")
	if t := obj_ControllerInputs(xn, "targets"); t != nil {
		obj.Targets = *t
	}
}

func load_FxVertexInputBinding(xn *xmlx.Node, obj *cdom.FxVertexInputBinding) {
	obj.InputSemantic, obj.Semantic = xas(xn, "input_semantic"), xas(xn, "semantic")
	obj.InputSet = xau64p(xn, "input_set")
}

func load_Float3(xn *xmlx.Node, obj *cdom.Float3) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_PxSceneDef(xn *xmlx.Node, obj *cdom.PxSceneDef) {
	obj.ForceFields = objs_PxForceFieldInst(xn, "instance_force_field")
	obj.Models = objs_PxModelInst(xn, "instance_physics_model")
	if tcn := node_TechCommon(xn); tcn != nil {
		obj.TC.Gravity = obj_SidVec3(tcn, "gravity")
		obj.TC.TimeStep = obj_SidFloat(tcn, "time_step")
	}
}

func load_InputShared(xn *xmlx.Node, obj *cdom.InputShared) {
	if in := obj_Input(xn, ""); in != nil {
		obj.Input = *in
	}
	obj.Offset = xau64(xn, "offset")
	obj.Set = xau64p(xn, "set")
}

func load_SidBool(xn *xmlx.Node, obj *cdom.SidBool) {
	obj.B = xb(xn, "")
}

func load_Int3(xn *xmlx.Node, obj *cdom.Int3) {
	arr_Ints(xn, len(obj), func(i int, n int64) {
		obj[i] = n
	})
}

func load_Float2x2(xn *xmlx.Node, obj *cdom.Float2x2) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_GeometryBrepEdges(xn *xmlx.Node, obj *cdom.GeometryBrepEdges) {
	if ii := obj_IndexedInputs(xn, ""); ii != nil {
		obj.IndexedInputs = *ii
	}
}

func load_GeometryMesh(xn *xmlx.Node, obj *cdom.GeometryMesh) {
	setIdRef(&obj.ConvexHullOf, xas(xn, "convex_hull_of"))
	for _, pn := range xcns(xn, "lines", "linestrips", "polygons", "polylist", "triangles", "trifans", "tristrips") {
		if p := obj_GeometryPrimitives(pn, ""); p != nil {
			obj.Primitives = append(obj.Primitives, p)
		}
	}
	obj.Vertices = obj_GeometryVertices(xn, "vertices")
}

func load_MaterialBinding(xn *xmlx.Node, obj *cdom.MaterialBinding) {
	obj.Params = objs_Param(xn, "param")
	if tcn := node_TechCommon(xn); tcn != nil {
		obj.TC.Materials = objs_FxMaterialInst(tcn, "instance_material")
	}
}

func load_FxEffectInst(xn *xmlx.Node, obj *cdom.FxEffectInst) {
	obj.TechniqueHints = objs_FxEffectInstTechniqueHint(xn, "technique_hint")
}

func load_KxJointLimits(xn *xmlx.Node, obj *cdom.KxJointLimits) {
	obj.Max = obj_SidFloat(xn, "max")
	obj.Min = obj_SidFloat(xn, "min")
}

func load_AnimationInst(xn *xmlx.Node, obj *cdom.AnimationInst) {
}

func load_GeometryBrepWires(xn *xmlx.Node, obj *cdom.GeometryBrepWires) {
	if iiv := obj_IndexedInputs(xn, ""); iiv != nil {
		obj.IndexedInputs = *iiv
	}
}

func load_FxCreateFormat(xn *xmlx.Node, obj *cdom.FxCreateFormat) {
	obj.Exact = xs(xn, "exact")
	obj.Hint = obj_FxCreateFormatHint(xn, "hint")
}

func load_FormulaInst(xn *xmlx.Node, obj *cdom.FormulaInst) {
}

func load_CameraOptics(xn *xmlx.Node, obj *cdom.CameraOptics) {
	if tcn := node_TechCommon(xn); tcn != nil {
		cn, on, pn := tcn, xcn(tcn, "orthographic"), xcn(tcn, "perspective")
		if cn = pn; cn == nil {
			cn = on
		}
		if cn != nil {
			sf := obj_SidFloat(cn, "aspect_ratio")
			obj.TC.AspectRatio = sf
			if sf = obj_SidFloat(cn, "zfar"); sf != nil {
				obj.TC.Zfar = *sf
			}
			if sf = obj_SidFloat(cn, "znear"); sf != nil {
				obj.TC.Znear = *sf
			}
		}
		if pn != nil {
			obj.TC.Perspective = obj_CameraPerspective(pn, "")
		}
		if on != nil {
			obj.TC.Orthographic = obj_CameraOrthographic(on, "")
		}
	}
}

func load_PxRigidBodyInst(xn *xmlx.Node, obj *cdom.PxRigidBodyInst) {
	setIdRef(&obj.TargetNode, xas(xn, "target"))
	if tcn := node_TechCommon(xn); tcn != nil {
		if rbc := obj_PxRigidBodyCommon(tcn, ""); rbc != nil {
			obj.TC.PxRigidBodyCommon = *rbc
		}
		v3 := xv3(tcn, "angular_velocity")
		if v3 != nil {
			obj.TC.AngularVelocity = *v3
		}
		if v3 = xv3(tcn, "velocity"); v3 != nil {
			obj.TC.LinearVelocity = *v3
		}
	}
}

func load_FxAnnotation(xn *xmlx.Node, obj *cdom.FxAnnotation) {
	for _, cn := range xn.Children {
		if cn.Type == xmlx.NT_ELEMENT {
			if obj.Value = xv(cn); obj.Value != nil {
				break
			}
		}
	}
}

func load_FxCreateFormatHint(xn *xmlx.Node, obj *cdom.FxCreateFormatHint) {
	obj.Space = xas(xn, "space")
	switch strings.ToUpper(xas(xn, "channels")) {
	case "RGB":
		obj.Channels = cdom.FxFormatChannelsRgb
	case "RGBA":
		obj.Channels = cdom.FxFormatChannelsRgba
	case "RGBE":
		obj.Channels = cdom.FxFormatChannelsRgbe
	case "L":
		obj.Channels = cdom.FxFormatChannelsL
	case "LA":
		obj.Channels = cdom.FxFormatChannelsLa
	case "D":
		obj.Channels = cdom.FxFormatChannelsD
	}
	switch strings.ToUpper(xas(xn, "range")) {
	case "FLOAT":
		obj.Range = cdom.FxFormatRangeFloat
	case "SINT":
		obj.Range = cdom.FxFormatRangeSint
	case "SNORM":
		obj.Range = cdom.FxFormatRangeSnorm
	case "UINT":
		obj.Range = cdom.FxFormatRangeUint
	case "UNORM":
		obj.Range = cdom.FxFormatRangeUnorm
	}
	switch strings.ToUpper(xas(xn, "precision")) {
	case "LOW":
		obj.Precision = cdom.FxFormatPrecisionLow
	case "MID":
		obj.Precision = cdom.FxFormatPrecisionMid
	case "HIGH":
		obj.Precision = cdom.FxFormatPrecisionHigh
	case "MAX":
		obj.Precision = cdom.FxFormatPrecisionMax
	default:
		obj.Precision = cdom.FxFormatPrecisionDefault
	}
}

func load_GeometryBrepCircle(xn *xmlx.Node, obj *cdom.GeometryBrepCircle) {
	obj.Radius = xf64(xn, "radius")
}

func load_PxSceneInst(xn *xmlx.Node, obj *cdom.PxSceneInst) {
}

func load_FxParamDef(xn *xmlx.Node, obj *cdom.FxParamDef) {
	if pd := obj_ParamDef(xn, ""); pd != nil {
		obj.ParamDef = *pd
	}
	obj.Annotations = objs_FxAnnotation(xn, "annotate")
	obj.Modifier = xs(xn, "modifier")
	obj.Semantic = xs(xn, "semantic")
}

func load_FxCreateMips(xn *xmlx.Node, obj *cdom.FxCreateMips) {
	obj.Levels = xau64(xn, "levels")
	if b := xabp(xn, "auto_generate"); b != nil {
		obj.NoAutoGen = !*b
	}
}

func load_KxFrameObject(xn *xmlx.Node, obj *cdom.KxFrameObject) {
	if kf := obj_KxFrame(xn, ""); kf != nil {
		obj.KxFrame = *kf
	}
}

func load_KxFrameOrigin(xn *xmlx.Node, obj *cdom.KxFrameOrigin) {
	if kf := obj_KxFrame(xn, ""); kf != nil {
		obj.KxFrame = *kf
	}
}

func load_KxFrameTcp(xn *xmlx.Node, obj *cdom.KxFrameTcp) {
	if kf := obj_KxFrame(xn, ""); kf != nil {
		obj.KxFrame = *kf
	}
}

func load_KxFrameTip(xn *xmlx.Node, obj *cdom.KxFrameTip) {
	if kf := obj_KxFrame(xn, ""); kf != nil {
		obj.KxFrame = *kf
	}
}

func load_KxFrame(xn *xmlx.Node, obj *cdom.KxFrame) {
	obj.Link.SetSidRef(xas(xn, "link"))
	obj.Transforms = get_Transforms(xn)
}

func load_KxJointInst(xn *xmlx.Node, obj *cdom.KxJointInst) {
}

func load_Input(xn *xmlx.Node, obj *cdom.Input) {
	obj.Semantic = xas(xn, "semantic")
	setIdRef(&obj.Source, xas(xn, "source"))
}

func load_ControllerDef(xn *xmlx.Node, obj *cdom.ControllerDef) {
	obj.Morph = obj_ControllerMorph(xn, "morph")
	obj.Skin = obj_ControllerSkin(xn, "skin")
}

func load_FxCreate(xn *xmlx.Node, obj *cdom.FxCreate) {
	obj.Format = obj_FxCreateFormat(xn, "format")
	if an := xcn(xn, "array"); an != nil {
		obj.ArrayLength = xau64(an, "length")
	}
}

func load_PxRigidConstraintLimit(xn *xmlx.Node, obj *cdom.PxRigidConstraintLimit) {
	sv3 := obj_SidVec3(xn, "max")
	if sv3 != nil {
		obj.Max = *sv3
	}
	if sv3 = obj_SidVec3(xn, "min"); sv3 != nil {
		obj.Min = *sv3
	}
}

func load_SidVec3(xn *xmlx.Node, obj *cdom.SidVec3) {
	if v3 := xv3(xn, ""); v3 != nil {
		obj.Vec3 = *v3
	}
}

func load_NodeInst(xn *xmlx.Node, obj *cdom.NodeInst) {
	setIdRef(&obj.Proxy, xas(xn, "proxy"))
}

func load_CameraOrthographic(xn *xmlx.Node, obj *cdom.CameraOrthographic) {
	obj.MagX = obj_SidFloat(xn, "xmag")
	obj.MagY = obj_SidFloat(xn, "ymag")
}

func load_Float3x3(xn *xmlx.Node, obj *cdom.Float3x3) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_FxPassState(xn *xmlx.Node, obj *cdom.FxPassState) {
	obj.Value = xas(xn, "value")
	obj.Param.SetParamRef(xas(xn, "param"))
	obj.Index = xf64(xn, "index")
}

func load_GeometryBrepSurface(xn *xmlx.Node, obj *cdom.GeometryBrepSurface) {
	if pos := obj_GeometryPositioning(xn, ""); pos != nil {
		obj.Location = *pos
	}
	obj.Element.Cone = obj_GeometryBrepCone(xn, "cone")
	obj.Element.Plane = obj_GeometryBrepPlane(xn, "plane")
	obj.Element.Cylinder = obj_GeometryBrepCylinder(xn, "cylinder")
	obj.Element.NurbsSurface = obj_GeometryBrepNurbsSurface(xn, "nurbs_surface")
	obj.Element.Sphere = obj_GeometryBrepSphere(xn, "sphere")
	obj.Element.Torus = obj_GeometryBrepTorus(xn, "torus")
	obj.Element.SweptSurface = obj_GeometryBrepSweptSurface(xn, "swept_surface")
}

func load_AnimationSampler(xn *xmlx.Node, obj *cdom.AnimationSampler) {
	for n, pb := range map[string]*cdom.AnimSamplerBehavior{"pre_behavior": &obj.PreBehavior, "post_behavior": &obj.PostBehavior} {
		switch strings.ToUpper(xas(xn, n)) {
		case "CONSTANT":
			*pb = cdom.AnimSamplerBehaviorConstant
		case "CYCLE":
			*pb = cdom.AnimSamplerBehaviorCycle
		case "CYCLE_RELATIVE":
			*pb = cdom.AnimSamplerBehaviorCycleRelative
		case "GRADIENT":
			*pb = cdom.AnimSamplerBehaviorGradient
		case "OSCILLATE":
			*pb = cdom.AnimSamplerBehaviorOscillate
		default:
			*pb = cdom.AnimSamplerBehaviorUndefined
		}
	}
}

func load_ParamOrUint(xn *xmlx.Node, obj *cdom.ParamOrUint) {
	obj.Param.SetParamRef(xs(xn, "param"))
	obj.U = xu64(xn, "uint")
}

func load_ParamDefs(xn *xmlx.Node, obj *cdom.ParamDefs) {
}

func load_ParamInsts(xn *xmlx.Node, obj *cdom.ParamInsts) {
}

func load_AnimationChannel(xn *xmlx.Node, obj *cdom.AnimationChannel) {
	setIdRef(&obj.Source, xas(xn, "source"))
	obj.Target.SetSidRef(xas(xn, "target"))
}

func load_LightAmbient(xn *xmlx.Node, obj *cdom.LightAmbient) {
	get_LightColor(xn, &obj.Color)
}

func load_CameraPerspective(xn *xmlx.Node, obj *cdom.CameraPerspective) {
	obj.FovX = obj_SidFloat(xn, "xfov")
	obj.FovY = obj_SidFloat(xn, "yfov")
}

func load_PxForceFieldDef(xn *xmlx.Node, obj *cdom.PxForceFieldDef) {
}

func load_PxRigidConstraintSpring(xn *xmlx.Node, obj *cdom.PxRigidConstraintSpring) {
	sf := obj_SidFloat(xn, "damping")
	if sf != nil {
		obj.Damping = *sf
	}
	if sf = obj_SidFloat(xn, "target_value"); sf != nil {
		obj.TargetValue = *sf
	}
	if sf = obj_SidFloat(xn, "stiffness"); sf != nil {
		obj.Stiffness = *sf
	} else {
		obj.Stiffness.F = 1
	}
}

func load_GeometryBrepNurbsSurface(xn *xmlx.Node, obj *cdom.GeometryBrepNurbsSurface) {
	obj.U.Degree = xau64(xn, "degree_u")
	obj.U.Closed = xab(xn, "closed_u")
	obj.V.Degree = xau64(xn, "degree_v")
	obj.V.Closed = xab(xn, "closed_v")
	if cv := obj_GeometryControlVertices(xn, "control_vertices"); cv != nil {
		obj.ControlVertices = *cv
	}
}

func load_CameraDef(xn *xmlx.Node, obj *cdom.CameraDef) {
	obj.Imager = obj_CameraImager(xn, "imager")
	if op := obj_CameraOptics(xn, "optics"); op != nil {
		obj.Optics = *op
	}
}

func load_GeometryBrepSweptSurface(xn *xmlx.Node, obj *cdom.GeometryBrepSweptSurface) {
	obj.Curve = obj_GeometryBrepCurve(xn, "curve")
	obj.Extrusion.Direction = xv3(xn, "direction")
	obj.Revolution.Origin = xv3(xn, "origin")
	obj.Revolution.Direction = xv3(xn, "axis")
}

func load_FxPassProgramBindUniform(xn *xmlx.Node, obj *cdom.FxPassProgramBindUniform) {
	obj.Symbol = xas(xn, "symbol")
	obj.ParamRef.SetParamRef(get_ParamRef(xn, "param"))
	for _, cn := range xn.Children {
		if cn.Type == xmlx.NT_ELEMENT {
			if obj.Value = xv(cn); obj.Value != nil {
				break
			}
		}
	}
}

func load_KxMotionAxis(xn *xmlx.Node, obj *cdom.KxMotionAxis) {
	obj.Axis.SetSidRef(xas(xn, "axis"))
	obj.Bindings = objs_KxBinding(xn, "bind")
	obj.Speed = obj_ParamOrFloat(xn, "speed")
	obj.Acceleration = obj_ParamOrFloat(xn, "acceleration")
	obj.Deceleration = obj_ParamOrFloat(xn, "deceleration")
	obj.Jerk = obj_ParamOrFloat(xn, "jerk")
}

func load_AnimationClipInst(xn *xmlx.Node, obj *cdom.AnimationClipInst) {
}

func load_AnimationClipDef(xn *xmlx.Node, obj *cdom.AnimationClipDef) {
	obj.Start = xf64(xn, "start")
	obj.End = xf64(xn, "end")
	obj.Animations = objs_AnimationInst(xn, "instance_animation")
	obj.Formulas = objs_FormulaInst(xn, "instance_formula")
}

func load_Float7(xn *xmlx.Node, obj *cdom.Float7) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_AnimationDef(xn *xmlx.Node, obj *cdom.AnimationDef) {
	obj.AnimationDefs = objs_AnimationDef(xn, "animation")
	obj.Samplers = objs_AnimationSampler(xn, "sampler")
	obj.Channels = objs_AnimationChannel(xn, "channel")
}

func load_PxMaterial(xn *xmlx.Node, obj *cdom.PxMaterial) {
	obj.Def = obj_PxMaterialDef(xn, "physics_material")
	obj.Inst = obj_PxMaterialInst(xn, "instance_physics_material")
}

func load_PxShape(xn *xmlx.Node, obj *cdom.PxShape) {
	if sb := obj_SidBool(xn, "hollow"); sb != nil {
		obj.Hollow = *sb
	}
	obj.Density = obj_SidFloat(xn, "density")
	obj.Mass = obj_SidFloat(xn, "mass")
	obj.Transforms = get_Transforms(xn)
	if pm := obj_PxMaterial(xn, ""); pm != nil {
		obj.Material = *pm
	}
	obj.Geometry.Plane = obj_GeometryBrepPlane(xn, "plane")
	obj.Geometry.Box = obj_GeometryBrepBox(xn, "box")
	obj.Geometry.Sphere = obj_GeometryBrepSphere(xn, "sphere")
	obj.Geometry.Cylinder = obj_PxCylinder(xn, "cylinder")
	obj.Geometry.Capsule = obj_GeometryBrepCapsule(xn, "capsule")
	obj.Geometry.Inst = obj_GeometryInst(xn, "instance_geometry")
}

func load_LightPoint(xn *xmlx.Node, obj *cdom.LightPoint) {
	get_LightColor(xn, &obj.Color)
	if a := obj_LightAttenuation(xn, ""); a != nil {
		obj.Attenuation = *a
	}
}

func load_GeometrySpline(xn *xmlx.Node, obj *cdom.GeometrySpline) {
	obj.Closed = xab(xn, "closed")
	if cv := obj_GeometryControlVertices(xn, "control_vertices"); cv != nil {
		obj.ControlVertices = *cv
	}
}

func load_GeometryInst(xn *xmlx.Node, obj *cdom.GeometryInst) {
	obj.MaterialBinding = obj_MaterialBinding(xn, "bind_material")
}

func load_PxRigidConstraintDef(xn *xmlx.Node, obj *cdom.PxRigidConstraintDef) {
	a := obj_PxRigidConstraintAttachment(xn, "attachment")
	if a != nil {
		obj.Attachment = *a
	}
	if a = obj_PxRigidConstraintAttachment(xn, "ref_attachment"); a != nil {
		obj.RefAttachment = *a
	}
	if tcn := node_TechCommon(xn); tcn != nil {
		sb := obj_SidBool(tcn, "enabled")
		if sb != nil {
			obj.TC.Enabled = *sb
		}
		if sb = obj_SidBool(tcn, "interpenetrate"); sb != nil {
			obj.TC.Interpenetrate = *sb
		}
		cn := xcn(tcn, "limits")
		if cn != nil {
			obj.TC.Limits.Angular = obj_PxRigidConstraintLimit(cn, "swing_cone_and_twist")
			obj.TC.Limits.Linear = obj_PxRigidConstraintLimit(cn, "linear")
		}
		if cn = xcn(tcn, "spring"); cn != nil {
			obj.TC.Spring.Angular = obj_PxRigidConstraintSpring(cn, "angular")
			obj.TC.Spring.Linear = obj_PxRigidConstraintSpring(cn, "linear")
		}
	}
}

func load_GeometryBrepSurfaceCurves(xn *xmlx.Node, obj *cdom.GeometryBrepSurfaceCurves) {
	obj.All = objs_GeometryBrepCurve(xn, "curve")
}

func load_Int3x3(xn *xmlx.Node, obj *cdom.Int3x3) {
	arr_Ints(xn, len(obj), func(i int, n int64) {
		obj[i] = n
	})
}

func load_SidString(xn *xmlx.Node, obj *cdom.SidString) {
	obj.S = xn.Value
}

func load_GeometryDef(xn *xmlx.Node, obj *cdom.GeometryDef) {
	obj.Brep = obj_GeometryBrep(xn, "brep")
	obj.Spline = obj_GeometrySpline(xn, "spline")
	if obj.Mesh = obj_GeometryMesh(xn, "mesh"); obj.Mesh == nil {
		obj.Mesh = obj_GeometryMesh(xn, "convex_mesh")
	}
}

func load_FxCreate2DSizeExact(xn *xmlx.Node, obj *cdom.FxCreate2DSizeExact) {
	obj.Width = xau64(xn, "width")
	obj.Height = xau64(xn, "height")
}

func load_VisualSceneInst(xn *xmlx.Node, obj *cdom.VisualSceneInst) {
}

func load_Int2(xn *xmlx.Node, obj *cdom.Int2) {
	arr_Ints(xn, len(obj), func(i int, n int64) {
		obj[i] = n
	})
}

func load_Float2(xn *xmlx.Node, obj *cdom.Float2) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_KxBinding(xn *xmlx.Node, obj *cdom.KxBinding) {
	obj.Symbol = xas(xn, "symbol")
	obj.Param.SetParamRef(get_ParamRef(xn, "param"))
	for _, cn := range xn.Children {
		if cn.Type == xmlx.NT_ELEMENT {
			if obj.Value = xv(cn); obj.Value != nil {
				break
			}
		}
	}
}

func load_FxPassEvaluationTarget(xn *xmlx.Node, obj *cdom.FxPassEvaluationTarget) {
	if u := xau64p(xn, "index"); u != nil {
		obj.Index = *u
	} else {
		obj.Index = 1
	}
	obj.Slice = xau64(xn, "slice")
	obj.Mip = xau64(xn, "mip")
	obj.CubeFace = get_CubeFace(xn)
	obj.Sampler.SetParamRef(get_ParamRef(xn, "param"))
	obj.Image = obj_FxImageInst(xn, "instance_image")
}

func load_FxColor(xn *xmlx.Node, obj *cdom.FxColor) {
	list_Rgba32(xn, &obj.Rgba32)
}

func load_FxColorOrTexture(xn *xmlx.Node, obj *cdom.FxColorOrTexture) {
	obj.ParamRef.SetParamRef(get_ParamRef(xn, "param"))
	obj.Texture = obj_FxTexture(xn, "texture")
	obj.Color = obj_FxColor(xn, "color")
	switch strings.ToUpper(xas(xn, "opaque")) {
	case "RGB_ZERO":
		obj.Opaque = cdom.FxTextureOpaqueRgb0
	case "RGB_ONE":
		obj.Opaque = cdom.FxTextureOpaqueRgb1
	case "A_ZERO":
		obj.Opaque = cdom.FxTextureOpaqueA0
	default:
		obj.Opaque = cdom.FxTextureOpaqueA1
	}
}

func load_GeometryVertices(xn *xmlx.Node, obj *cdom.GeometryVertices) {
}

func load_KxMotionSystem(xn *xmlx.Node, obj *cdom.KxMotionSystem) {
	obj.ArticulatedSystem = obj_KxArticulatedSystemInst(xn, "instance_articulated_system")
	if tcn := node_TechCommon(xn); tcn != nil {
		obj.TC.AxisInfos = objs_KxMotionAxis(tcn, "axis_info")
		obj.TC.EffectorInfo = obj_KxEffector(tcn, "effector_info")
	}
}

func load_Int4x4(xn *xmlx.Node, obj *cdom.Int4x4) {
	arr_Ints(xn, len(obj), func(i int, n int64) {
		obj[i] = n
	})
}

func load_CameraImager(xn *xmlx.Node, obj *cdom.CameraImager) {
}

func load_FxPassEvaluationClearDepth(xn *xmlx.Node, obj *cdom.FxPassEvaluationClearDepth) {
	obj.Index = xau64(xn, "index")
	obj.F = xf64(xn, "")
}

func load_GeometryBrepEllipse(xn *xmlx.Node, obj *cdom.GeometryBrepEllipse) {
	if f2 := obj_Float2(xn, "radius"); f2 != nil {
		obj.Radii = *f2
	}
}

func load_FxBinding(xn *xmlx.Node, obj *cdom.FxBinding) {
	obj.Semantic = xas(xn, "semantic")
	obj.Target.SetSidRef(xas(xn, "target"))
}

func load_GeometryBrepParabola(xn *xmlx.Node, obj *cdom.GeometryBrepParabola) {
	obj.FocalLength = xf64(xn, "focal")
}

func load_SidFloat(xn *xmlx.Node, obj *cdom.SidFloat) {
	obj.F = xf64(xn, "")
}

func load_Bool3(xn *xmlx.Node, obj *cdom.Bool3) {
	arr_Bools(xn, len(obj), func(i int, b bool) {
		obj[i] = b
	})
}

func load_VisualSceneDef(xn *xmlx.Node, obj *cdom.VisualSceneDef) {
	obj.Evaluations = objs_VisualSceneEvaluation(xn, "evaluate_scene")
	obj.Nodes = objs_NodeDef(xn, "node")
}

func load_KxSceneInst(xn *xmlx.Node, obj *cdom.KxSceneInst) {
	obj.ModelBindings = objs_KxModelBinding(xn, "bind_kinematics_model")
	obj.JointAxisBindings = objs_KxJointAxisBinding(xn, "bind_joint_axis")
}

func load_LightDirectional(xn *xmlx.Node, obj *cdom.LightDirectional) {
	get_LightColor(xn, &obj.Color)
}

func load_KxAxisIndex(xn *xmlx.Node, obj *cdom.KxAxisIndex) {
	obj.Semantic = xas(xn, "semantic")
	if pi := obj_ParamOrInt(xn, ""); pi != nil {
		obj.I = *pi
	}
}

func load_VisualSceneRendering(xn *xmlx.Node, obj *cdom.VisualSceneRendering) {
	setIdRef(&obj.CameraNode, xas(xn, "camera_node"))
	for _, ln := range xcns(xn, "layer") {
		obj.Layers[ln.Value] = true
	}
	obj.MaterialInst = obj_VisualSceneRenderingMaterialInst(xn, "instance_material")
}

func load_GeometryPolygonHole(xn *xmlx.Node, obj *cdom.GeometryPolygonHole) {
	obj.Indices = listcn_Uints(xn, "p")
	for _, cn := range xcns(xn, "h") {
		obj.Holes = append(obj.Holes, list_Uints(cn))
	}
}

func load_Param(xn *xmlx.Node, obj *cdom.Param) {
	obj.Type = xas(xn, "type")
	obj.Semantic = xas(xn, "semantic")
}

func load_ParamOrSidFloat(xn *xmlx.Node, obj *cdom.ParamOrSidFloat) {
	obj.Param.SetParamRef(xs(xn, "param"))
	if f := obj_SidFloat(xn, "float"); f != nil {
		obj.F = *f
	}
}

func load_Float2x3(xn *xmlx.Node, obj *cdom.Float2x3) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_FxCreateInitFrom(xn *xmlx.Node, obj *cdom.FxCreateInitFrom) {
	if i := obj_FxInitFrom(xn, ""); i != nil {
		obj.FxInitFrom = *i
	}
	obj.ArrayIndex = xau64(xn, "array_index")
	obj.MipIndex = xau64(xn, "mip_index")
}

func load_LightAttenuation(xn *xmlx.Node, obj *cdom.LightAttenuation) {
	var sf *cdom.SidFloat
	for n, f := range map[string]*cdom.SidFloat{"constant_attenuation": &obj.Constant, "linear_attenuation": &obj.Linear, "quadratic_attenuation": &obj.Quadratic} {
		if sf = obj_SidFloat(xn, n); sf != nil {
			*f = *sf
		}
	}
}

func load_KxModelBinding(xn *xmlx.Node, obj *cdom.KxModelBinding) {
	setIdRef(&obj.Node, xas(xn, "node"))
	obj.Model.SidRef.SetSidRef(xss(xn, "SIDREF", "sidref"))
	obj.Model.ParamRef.SetParamRef(get_ParamRef(xn, "param"))
}

func load_GeometryBrepSolids(xn *xmlx.Node, obj *cdom.GeometryBrepSolids) {
	if iiv := obj_IndexedInputs(xn, ""); iiv != nil {
		obj.IndexedInputs = *iiv
	}
}

func load_FxMaterialDef(xn *xmlx.Node, obj *cdom.FxMaterialDef) {
	if e := obj_FxEffectInst(xn, "instance_effect"); e != nil {
		obj.Effect = *e
	}
}

func load_ParamOrFloat(xn *xmlx.Node, obj *cdom.ParamOrFloat) {
	obj.Param.SetParamRef(xs(xn, "param"))
	obj.F = xf64(xn, "float")
}

func load_SidFloat3(xn *xmlx.Node, obj *cdom.SidFloat3) {
	if f3 := obj_Float3(xn, ""); f3 != nil {
		obj.F = *f3
	}
}

func load_FxInitFrom(xn *xmlx.Node, obj *cdom.FxInitFrom) {
	obj.RefUrl = xs(xn, "ref")
	if hn := xcn(xn, "hex"); hn != nil {
		obj.Raw.Format = xas(hn, "format")
		obj.Raw.Data, _ = hex.DecodeString(xs(hn, ""))
	}
}

func load_FxTexture(xn *xmlx.Node, obj *cdom.FxTexture) {
	obj.TexCoord = xas(xn, "texcoord")
	obj.Sampler2D.SetParamRef(xas(xn, "texture"))
}

func load_Int4(xn *xmlx.Node, obj *cdom.Int4) {
	arr_Ints(xn, len(obj), func(i int, n int64) {
		obj[i] = n
	})
}

func load_GeometryBrepOrientation(xn *xmlx.Node, obj *cdom.GeometryBrepOrientation) {
	fs := list_Floats(xn)
	if len(fs) > 0 {
		if obj.Axis.X = fs[0]; len(fs) > 1 {
			if obj.Axis.Y = fs[1]; len(fs) > 2 {
				if obj.Axis.Z = fs[2]; len(fs) > 3 {
					obj.Angle = fs[3]
				}
			}
		}
	}
}

func load_GeometryPrimitives(xn *xmlx.Node, obj *cdom.GeometryPrimitives) {
	switch xn.Name.Local {
	case "lines":
		obj.Kind = cdom.GeometryPrimitiveKindLines
	case "linestrips":
		obj.Kind = cdom.GeometryPrimitiveKindLineStrips
	case "polygons":
		obj.Kind = cdom.GeometryPrimitiveKindPolygons
	case "polylist":
		obj.Kind = cdom.GeometryPrimitiveKindPolylist
	case "triangles":
		obj.Kind = cdom.GeometryPrimitiveKindTriangles
	case "trifans":
		obj.Kind = cdom.GeometryPrimitiveKindTrifans
	case "tristrips":
		obj.Kind = cdom.GeometryPrimitiveKindTristrips
	}
	if obj.Kind > 0 {
		if ii := obj_IndexedInputs(xn, ""); ii != nil {
			obj.IndexedInputs = *ii
		}
		obj.Material = xas(xn, "material")
		obj.PolyHoles = objs_GeometryPolygonHole(xn, "ph")
	}
}

func load_Sources(xn *xmlx.Node, obj *cdom.Sources) {
}

func load_PxRigidConstraintAttachment(xn *xmlx.Node, obj *cdom.PxRigidConstraintAttachment) {
	obj.RigidBody.SetSidRef(xas(xn, "rigid_body"))
	obj.Transforms = get_Transforms(xn)
}

func load_FxImageInitFrom(xn *xmlx.Node, obj *cdom.FxImageInitFrom) {
	if i := obj_FxInitFrom(xn, ""); i != nil {
		obj.FxInitFrom = *i
	}
	if b := xabp(xn, "mips_generate"); b != nil {
		obj.NoAutoMip = !*b
	}
}

func load_FxTechniqueCommon(xn *xmlx.Node, obj *cdom.FxTechniqueCommon) {
	if t := obj_FxTechnique(xn, ""); t != nil {
		obj.FxTechnique = *t
	}
	if cn := xcn1(xn, "blinn", "constant", "lambert", "phong"); cn != nil {
		switch cn.Name.Local {
		case "blinn":
			obj.Kind = cdom.FxTechniqueKindBlinn
		case "lambert":
			obj.Kind = cdom.FxTechniqueKindLambert
		case "phong":
			obj.Kind = cdom.FxTechniqueKindPhong
		default:
			obj.Kind = cdom.FxTechniqueKindConstant
		}
		//	constant
		obj.Emission = obj_FxColorOrTexture(cn, "emission")
		obj.IndexOfRefraction = obj_ParamOrSidFloat(cn, "index_of_refraction")
		obj.Reflective = obj_FxColorOrTexture(cn, "reflective")
		obj.Reflectivity = obj_ParamOrSidFloat(cn, "reflectivity")
		obj.Transparency = obj_ParamOrSidFloat(cn, "transparency")
		obj.Transparent = obj_FxColorOrTexture(cn, "transparent")
		//	lambert
		obj.Ambient = obj_FxColorOrTexture(cn, "ambient")
		obj.Diffuse = obj_FxColorOrTexture(cn, "diffuse")
		//	blinn or phong
		obj.Specular = obj_FxColorOrTexture(cn, "specular")
		obj.Shininess = obj_ParamOrSidFloat(cn, "shininess")

	}
}

func load_Float3x4(xn *xmlx.Node, obj *cdom.Float3x4) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_KxJointAxisBinding(xn *xmlx.Node, obj *cdom.KxJointAxisBinding) {
	obj.Target.SetSidRef(xas(xn, "target"))
	if pf := obj_ParamOrFloat(xn, "value"); pf != nil {
		obj.Value = *pf
	}
	if psr := obj_ParamOrRefSid(xn, "axis"); psr != nil {
		obj.Axis = *psr
	}
}

func load_KxJointDef(xn *xmlx.Node, obj *cdom.KxJointDef) {
	var j *cdom.KxJoint
	for _, cn := range xcns(xn, "prismatic", "revolute") {
		if j = obj_KxJoint(cn, ""); (j != nil) && (j.Kind > 0) {
			obj.All = append(obj.All, j)
		}
	}
}

func load_VisualSceneEvaluation(xn *xmlx.Node, obj *cdom.VisualSceneEvaluation) {
	if b := xabp(xn, "enable"); b != nil {
		obj.Disabled = !*b
	}
	obj.RenderPasses = objs_VisualSceneRendering(xn, "render")
}

func load_FxProfileGlslCodeInclude(xn *xmlx.Node, obj *cdom.FxProfileGlslCodeInclude) {
	if obj.S, obj.IsInclude = xn.Value, (xn.Name.Local == "include"); obj.IsInclude {
		obj.S = xas(xn, "url")
	}
}

func load_AssetGeographicLocation(xn *xmlx.Node, obj *cdom.AssetGeographicLocation) {
	// obj.Longitude = xf64(xn, "longitude")
	// obj.Latitude = xf64(xn, "latitude")
	// if an := xcn(xn, "altitude"); an != nil {
	// 	obj.Altitude = xf64(an, "")
	// 	obj.AltitudeAbsolute = xab(an, "absolute")
	// }
}

func load_KxEffector(xn *xmlx.Node, obj *cdom.KxEffector) {
	obj.Bindings = objs_KxBinding(xn, "bind")
	obj.Speed = obj_ParamOrFloat2(xn, "speed")
	obj.Acceleration = obj_ParamOrFloat2(xn, "acceleration")
	obj.Deceleration = obj_ParamOrFloat2(xn, "deceleration")
	obj.Jerk = obj_ParamOrFloat2(xn, "jerk")
}

func load_GeometryBrepTorus(xn *xmlx.Node, obj *cdom.GeometryBrepTorus) {
	if f2 := obj_Float2(xn, "radius"); f2 != nil {
		obj.Radii = *f2
	}
}

func load_Layers(xn *xmlx.Node, obj *cdom.Layers) {
}

func load_FxProfile(xn *xmlx.Node, obj *cdom.FxProfile) {
	switch xn.Name.Local {
	case "profile_GLSL":
		obj.Glsl = obj_FxProfileGlsl(xn, "")
	case "profile_COMMON":
		obj.Common = obj_FxProfileCommon(xn, "")
	}
}

func load_KxAxisLimits(xn *xmlx.Node, obj *cdom.KxAxisLimits) {
	pf := obj_ParamOrFloat(xn, "max")
	if pf != nil {
		obj.Max = *pf
	}
	if pf = obj_ParamOrFloat(xn, "min"); pf != nil {
		obj.Min = *pf
	}
}

func load_FxMaterialInst(xn *xmlx.Node, obj *cdom.FxMaterialInst) {
	obj.Symbol = xas(xn, "symbol")
	obj.VertexInputBindings = objs_FxVertexInputBinding(xn, "bind_vertex_input")
	obj.Bindings = objs_FxBinding(xn, "bind")
}

func load_GeometryBrepCone(xn *xmlx.Node, obj *cdom.GeometryBrepCone) {
	obj.Angle = xf64(xn, "angle")
	obj.Radius = xf64(xn, "radius")
}

func load_Extra(xn *xmlx.Node, obj *cdom.Extra) {
	obj.Type = xas(xn, "type")
}

func load_ControllerSkin(xn *xmlx.Node, obj *cdom.ControllerSkin) {
	setIdRef(&obj.Source, xas(xn, "source"))
	obj.BindShapeMatrix = *xm4(xn, "bind_shape_matrix")
	if ci := obj_ControllerInputs(xn, "joints"); ci != nil {
		obj.Joints = *ci
	}
	if iiv := obj_IndexedInputs(xn, "vertex_weights"); iiv != nil {
		obj.VertexWeights = *iiv
	}
}

func load_Bool2(xn *xmlx.Node, obj *cdom.Bool2) {
	arr_Bools(xn, len(obj), func(i int, b bool) {
		obj[i] = b
	})
}

func load_VisualSceneRenderingMaterialInst(xn *xmlx.Node, obj *cdom.VisualSceneRenderingMaterialInst) {
	obj.Bindings = objs_FxBinding(xn, "bind")
	if tn := xcn(xn, "technique_override"); tn != nil {
		obj.OverrideTechnique.Ref.SetSidRef(xas(tn, "ref"))
		obj.OverrideTechnique.Pass.SetSidRef(xas(tn, "pass"))
	}
}

func load_ChildNode(xn *xmlx.Node, obj *cdom.ChildNode) {
	switch xn.Name.Local {
	case "instance_node":
		obj.Inst = obj_NodeInst(xn, "")
	case "node":
		obj.Def = obj_NodeDef(xn, "")
	}
}

func load_NodeDef(xn *xmlx.Node, obj *cdom.NodeDef) {
	var nc *cdom.ChildNode
	obj.IsSkinJoint = (xas(xn, "type") == "JOINT")
	if l := xsdt.ListValues(xas(xn, "layer")); len(l) > 0 {
		for _, n := range l {
			obj.Layers[n] = true
		}
	}
	for _, cn := range xcns(xn, "node", "instance_node") {
		if nc = obj_ChildNode(cn, ""); nc != nil {
			obj.Nodes = append(obj.Nodes, *nc)
		}
	}
	obj.Transforms = get_Transforms(xn)
	obj.Insts.Camera = objs_CameraInst(xn, "instance_camera")
	obj.Insts.Controller = objs_ControllerInst(xn, "instance_controller")
	obj.Insts.Geometry = objs_GeometryInst(xn, "instance_geometry")
	obj.Insts.Light = objs_LightInst(xn, "instance_light")
}

func load_PxForceFieldInst(xn *xmlx.Node, obj *cdom.PxForceFieldInst) {
}

func load_KxKinematicsSystem(xn *xmlx.Node, obj *cdom.KxKinematicsSystem) {
	obj.Models = objs_KxModelInst(xn, "instance_kinematics_model")
	if tcn := node_TechCommon(xn); tcn != nil {
		obj.TC.AxisInfos = objs_KxKinematicsAxis(tcn, "axis_info")
		obj.TC.Frame.Object = obj_KxFrameObject(tcn, "frame_object")
		obj.TC.Frame.Tcp = obj_KxFrameTcp(tcn, "frame_tcp")
		if f := obj_KxFrameTip(tcn, "frame_tip"); f != nil {
			obj.TC.Frame.Tip = *f
		}
		if f := obj_KxFrameOrigin(tcn, "frame_origin"); f != nil {
			obj.TC.Frame.Origin = *f
		}
	}
}

func load_GeometryBrepLine(xn *xmlx.Node, obj *cdom.GeometryBrepLine) {
	v3 := xv3(xn, "origin")
	if v3 != nil {
		obj.Origin = *v3
	}
	if v3 = xv3(xn, "direction"); v3 != nil {
		obj.Direction = *v3
	}
}

func load_GeometryBrep(xn *xmlx.Node, obj *cdom.GeometryBrep) {
	if v := obj_GeometryVertices(xn, "vertices"); v != nil {
		obj.Vertices = *v
	}
	obj.Curves = obj_GeometryBrepCurves(xn, "curves")
	obj.Edges = obj_GeometryBrepEdges(xn, "edges")
	obj.Faces = obj_GeometryBrepFaces(xn, "faces")
	obj.Pcurves = obj_GeometryBrepPcurves(xn, "pcurves")
	obj.Shells = obj_GeometryBrepShells(xn, "shells")
	obj.Solids = obj_GeometryBrepSolids(xn, "solids")
	obj.SurfaceCurves = obj_GeometryBrepSurfaceCurves(xn, "surface_curves")
	obj.Wires = obj_GeometryBrepWires(xn, "wires")
}

func load_KxModelDef(xn *xmlx.Node, obj *cdom.KxModelDef) {
	var f *cdom.Formula
	if tcn := node_TechCommon(xn); tcn != nil {
		has_ParamDefs(tcn, &obj.TC.HasParamDefs)
		for _, cn := range xcns(tcn, "formula", "instance_formula") {
			if f = obj_Formula(cn, ""); f != nil {
				obj.TC.Formulas = append(obj.TC.Formulas, *f)
			}
		}
		obj.TC.Links = objs_KxLink(tcn, "link")
	}
}

func load_FxImageInst(xn *xmlx.Node, obj *cdom.FxImageInst) {
}

func load_GeometryBrepBox(xn *xmlx.Node, obj *cdom.GeometryBrepBox) {
	if v3 := xv3(xn, "half_extents"); v3 != nil {
		obj.HalfExtents = *v3
	}
}

func load_ParamOrFloat2(xn *xmlx.Node, obj *cdom.ParamOrFloat2) {
	obj.Param.SetParamRef(xs(xn, "param"))
	if fn := xcn(xn, "float2"); fn != nil {
		arr_Floats(fn, 2, func(i int, f float64) {
			obj.F[i] = f
		})
	}
}

func load_Technique(xn *xmlx.Node, obj *cdom.Technique) {
	obj.Profile = xas(xn, "profile")
	obj.Data = xn.String()
}

func load_FxPassEvaluationClearStencil(xn *xmlx.Node, obj *cdom.FxPassEvaluationClearStencil) {
	obj.Index = xau64(xn, "index")
	obj.B = byte(xu64(xn, ""))
}

func load_FxCreate2DSizeRatio(xn *xmlx.Node, obj *cdom.FxCreate2DSizeRatio) {
	obj.Width = xaf64(xn, "width")
	obj.Height = xaf64(xn, "height")
}

func load_GeometryBrepPlane(xn *xmlx.Node, obj *cdom.GeometryBrepPlane) {
	if f4 := obj_Float4(xn, "equation"); f4 != nil {
		obj.Equation = *f4
	}
}

func load_FxPassProgramBindAttribute(xn *xmlx.Node, obj *cdom.FxPassProgramBindAttribute) {
	obj.Symbol = xas(xn, "symbol")
	obj.Semantic = xs(xn, "semantic")
}

func load_FxParamDefs(xn *xmlx.Node, obj *cdom.FxParamDefs) {
}

func load_PxMaterialDef(xn *xmlx.Node, obj *cdom.PxMaterialDef) {
	var sf *cdom.SidFloat
	if tcn := node_TechCommon(xn); tcn != nil {
		for n, f := range map[string]*cdom.SidFloat{"dynamic_friction": &obj.TC.DynamicFriction, "restitution": &obj.TC.Restitution, "static_friction": &obj.TC.StaticFriction} {
			if sf = obj_SidFloat(tcn, n); sf != nil {
				*f = *sf
			}
		}
	}
}

func load_Float4x2(xn *xmlx.Node, obj *cdom.Float4x2) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_PxRigidConstraintInst(xn *xmlx.Node, obj *cdom.PxRigidConstraintInst) {
}

func load_GeometryBrepSurfaces(xn *xmlx.Node, obj *cdom.GeometryBrepSurfaces) {
	obj.All = objs_GeometryBrepSurface(xn, "surface")
}

func load_Bool4(xn *xmlx.Node, obj *cdom.Bool4) {
	arr_Bools(xn, len(obj), func(i int, b bool) {
		obj[i] = b
	})
}

func load_FxCreate3D(xn *xmlx.Node, obj *cdom.FxCreate3D) {
	if c := obj_FxCreate(xn, ""); c != nil {
		obj.FxCreate = *c
	}
	if m := obj_FxCreateMips(xn, "mips"); m != nil {
		obj.Mips = *m
	}
	obj.InitFrom = objs_FxCreate3DInitFrom(xn, "init_from")
	if sn := xcn(xn, "size"); sn != nil {
		obj.Size.Width = xau64(sn, "width")
		obj.Size.Height = xau64(sn, "height")
		obj.Size.Depth = xau64(sn, "depth")
	}
}

func load_GeometryBrepCurves(xn *xmlx.Node, obj *cdom.GeometryBrepCurves) {
	obj.All = objs_GeometryBrepCurve(xn, "curve")
}

func load_PxRigidBodyDef(xn *xmlx.Node, obj *cdom.PxRigidBodyDef) {
	if tcn := node_TechCommon(xn); tcn != nil {
		if rbc := obj_PxRigidBodyCommon(tcn, ""); rbc != nil {
			obj.TC.PxRigidBodyCommon = *rbc
		}
	}
}

func load_KxArticulatedSystemDef(xn *xmlx.Node, obj *cdom.KxArticulatedSystemDef) {
	obj.Kinematics = obj_KxKinematicsSystem(xn, "kinematics")
	obj.Motion = obj_KxMotionSystem(xn, "motion")
}

func load_GeometryBrepPcurves(xn *xmlx.Node, obj *cdom.GeometryBrepPcurves) {
	if iiv := obj_IndexedInputs(xn, ""); iiv != nil {
		obj.IndexedInputs = *iiv
	}
}

func load_GeometryBrepHyperbola(xn *xmlx.Node, obj *cdom.GeometryBrepHyperbola) {
	if f2 := obj_Float2(xn, "radius"); f2 != nil {
		obj.Radii = *f2
	}
}

func load_PxRigidBodyCommon(xn *xmlx.Node, obj *cdom.PxRigidBodyCommon) {
	if sb := obj_SidBool(xn, "dynamic"); sb != nil {
		obj.Dynamic = *sb
	}
	obj.Mass = obj_SidFloat(xn, "mass")
	if tn := xcn(xn, "mass_frame"); tn != nil {
		obj.MassFrame = get_Transforms(tn)
	}
	obj.Inertia = obj_SidFloat3(xn, "inertia")
	obj.Shapes = objs_PxShape(xn, "shape")
	if pm := obj_PxMaterial(xn, ""); pm != nil {
		obj.Material = *pm
	}
}

func load_KxArticulatedSystemInst(xn *xmlx.Node, obj *cdom.KxArticulatedSystemInst) {
	obj.Bindings = objs_KxBinding(xn, "bind")
}

func load_Float3x2(xn *xmlx.Node, obj *cdom.Float3x2) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_GeometryControlVertices(xn *xmlx.Node, obj *cdom.GeometryControlVertices) {
}

func load_Float2x4(xn *xmlx.Node, obj *cdom.Float2x4) {
	arr_Floats(xn, len(obj), func(i int, f float64) {
		obj[i] = f
	})
}

func load_ParamInst(xn *xmlx.Node, obj *cdom.ParamInst) {
	obj.Ref.SetSidRef(xas(xn, "ref"))
	for _, cn := range xn.Children {
		if cn.Type == xmlx.NT_ELEMENT {
			if cn.Name.Local == "connect_param" {
				obj.IsConnectParamRef = true
				obj.Value = xas(cn, "ref")
				break
			} else if obj.Value = xv(cn); obj.Value != nil {
				break
			}
		}
	}
}

func load_FxTechniqueGlsl(xn *xmlx.Node, obj *cdom.FxTechniqueGlsl) {
	if ft := obj_FxTechnique(xn, ""); ft != nil {
		obj.FxTechnique = *ft
	}
	obj.Annotations = objs_FxAnnotation(xn, "annotate")
	obj.Passes = objs_FxPass(xn, "pass")
}

func load_KxLink(xn *xmlx.Node, obj *cdom.KxLink) {
	var a *cdom.KxAttachment
	obj.Transforms = get_Transforms(xn)
	for _, cn := range xcns(xn, "attachment_full", "attachment_start", "attachment_end") {
		if a = obj_KxAttachment(cn, ""); a != nil {
			obj.Attachments = append(obj.Attachments, a)
		}
	}
}

func load_FxProfileGlsl(xn *xmlx.Node, obj *cdom.FxProfileGlsl) {
	var ci *cdom.FxProfileGlslCodeInclude
	obj.Platform = xasd(xn, "platform", obj.Platform)
	for _, gt := range objs_FxTechniqueGlsl(xn, "technique") {
		obj.Techniques[gt.Sid] = gt
	}
	for _, cn := range xcns(xn, "code", "include") {
		if ci = obj_FxProfileGlslCodeInclude(cn, ""); ci != nil {
			obj.CodesIncludes = append(obj.CodesIncludes, *ci)
		}
	}
}

func load_FxCreateCube(xn *xmlx.Node, obj *cdom.FxCreateCube) {
	if c := obj_FxCreate(xn, ""); c != nil {
		obj.FxCreate = *c
	}
	if m := obj_FxCreateMips(xn, "mips"); m != nil {
		obj.Mips = *m
	}
	obj.InitFrom = objs_FxCreateCubeInitFrom(xn, "init_from")
	if sn := xcn(xn, "size"); sn != nil {
		obj.Size.Width = xau64(sn, "width")
	}
}

func load_FxCreate3DInitFrom(xn *xmlx.Node, obj *cdom.FxCreate3DInitFrom) {
	if ci := obj_FxCreateInitFrom(xn, ""); ci != nil {
		obj.FxCreateInitFrom = *ci
	}
	obj.Depth = xau64(xn, "depth")
}

func load_KxSceneDef(xn *xmlx.Node, obj *cdom.KxSceneDef) {
	obj.ArticulatedSystems = objs_KxArticulatedSystemInst(xn, "instance_articulated_systems")
	obj.Models = objs_KxModelInst(xn, "instance_kinematics_model")
}

func load_ControllerInputs(xn *xmlx.Node, obj *cdom.ControllerInputs) {
}

func load_IndexedInputs(xn *xmlx.Node, obj *cdom.IndexedInputs) {
	obj.Count = xau64(xn, "count")
	obj.Inputs = objs_InputShared(xn, "input")
	obj.Indices = listcn_Uints(xn, "p")
	obj.Vcount = listcn_Ints(xn, "vcount")
}

func load_FxTechnique(xn *xmlx.Node, obj *cdom.FxTechnique) {
}

func load_GeometryBrepShells(xn *xmlx.Node, obj *cdom.GeometryBrepShells) {
}

func load_LightSpot(xn *xmlx.Node, obj *cdom.LightSpot) {
	get_LightColor(xn, &obj.Color)
	if a := obj_LightAttenuation(xn, ""); a != nil {
		obj.Attenuation = *a
	}
	sf := obj_SidFloat(xn, "falloff_exponent")
	if sf != nil {
		obj.Falloff.Exponent = *sf
	}
	if sf = obj_SidFloat(xn, "falloff_angle"); sf != nil {
		obj.Falloff.Angle = *sf
	}
}

func load_FxCreate2D(xn *xmlx.Node, obj *cdom.FxCreate2D) {
	if c := obj_FxCreate(xn, ""); c != nil {
		obj.FxCreate = *c
	}
	obj.Mips = obj_FxCreateMips(xn, "mips")
	obj.InitFrom = objs_FxCreateInitFrom(xn, "init_from")
	obj.Size.Exact = obj_FxCreate2DSizeExact(xn, "size_exact")
	obj.Size.Ratio = obj_FxCreate2DSizeRatio(xn, "size_ratio")
	obj.Unnormalized = (xcns(xn, "unnormalized") != nil)
}

func load_ParamOrBool(xn *xmlx.Node, obj *cdom.ParamOrBool) {
	obj.Param.SetParamRef(xs(xn, "param"))
	obj.B = xb(xn, "bool")
}

func load_GeometryBrepFaces(xn *xmlx.Node, obj *cdom.GeometryBrepFaces) {
	if iiv := obj_IndexedInputs(xn, ""); iiv != nil {
		obj.IndexedInputs = *iiv
	}
}

func load_KxAttachment(xn *xmlx.Node, obj *cdom.KxAttachment) {
	switch xn.Name.Local {
	case "attachment_full":
		obj.Kind = cdom.KxAttachmentKindFull
	case "attachment_start":
		obj.Kind = cdom.KxAttachmentKindStart
	case "attachment_end":
		obj.Kind = cdom.KxAttachmentKindEnd
	}
	if obj.Kind > 0 {
		obj.Joint.SetSidRef(xas(xn, "joint"))
		obj.Transforms = get_Transforms(xn)
		obj.Link = obj_KxLink(xn, "link")
	}
}

func load_AssetContributor(xn *xmlx.Node, obj *cdom.AssetContributor) {
	// obj.Author = xs(xn, "author")
	// obj.AuthorEmail = xs(xn, "author_email")
	// obj.AuthorWebsite = xs(xn, "author_website")
	// obj.AuthoringTool = xs(xn, "authoring_tool")
	// obj.Comments = xs(xn, "comments")
	// obj.Copyright = xs(xn, "copyright")
	// obj.SourceData = xs(xn, "source_data")
}

func load_ParamOrRefSid(xn *xmlx.Node, obj *cdom.ParamOrRefSid) {
	obj.Param.SetParamRef(xs(xn, "param"))
	obj.Sr.SetSidRef(xss(xn, "sidref", "SIDREF"))
}

func load_FxSampler(xn *xmlx.Node, obj *cdom.FxSampler) {
	switch xn.Name.Local {
	case "sampler1D":
		obj.Kind = cdom.FxSamplerKind1D
	case "sampler2D":
		obj.Kind = cdom.FxSamplerKind2D
	case "sampler3D":
		obj.Kind = cdom.FxSamplerKind3D
	case "samplerCUBE":
		obj.Kind = cdom.FxSamplerKindCube
	case "samplerDEPTH":
		obj.Kind = cdom.FxSamplerKindDepth
	case "samplerRECT":
		obj.Kind = cdom.FxSamplerKindRect
	}
	if obj.Kind > 0 {
		obj.Image = obj_FxImageInst(xn, "instance_image")
		if ss := obj_FxSamplerStates(xn, ""); ss != nil {
			obj.FxSamplerStates = *ss
		}
	}
}

func load_Scene(xn *xmlx.Node, obj *cdom.Scene) {
	obj.Kinematics = obj_KxSceneInst(xn, "instance_kinematics_scene")
	obj.Visual = obj_VisualSceneInst(xn, "instance_visual_scene")
	obj.Physics = objs_PxSceneInst(xn, "instance_physics_scene")
}

func load_FxPassProgramShaderSources(xn *xmlx.Node, obj *cdom.FxPassProgramShaderSources) {
}

func load_FxSamplerImage(xn *xmlx.Node, obj *cdom.FxSamplerImage) {
	if ii := obj_FxImageInst(xn, ""); ii != nil {
		obj.FxImageInst = *ii
	}
}

func load_FxSamplerStates(xn *xmlx.Node, obj *cdom.FxSamplerStates) {
	obj.Filtering = obj_FxSamplerFiltering(xn, "")
	obj.Wrapping = obj_SamplerWrapping(xn, "")
}

func load_PxRigidBodyDefs(xn *xmlx.Node, obj *cdom.PxRigidBodyDefs) {
}

func load_PxRigidConstraintDefs(xn *xmlx.Node, obj *cdom.PxRigidConstraintDefs) {
}

func load_FxGlslTechniques(xn *xmlx.Node, obj *cdom.FxGlslTechniques) {
}

func load_GeometryPrimitiveKind(xn *xmlx.Node, obj *cdom.GeometryPrimitiveKind) {
}

func load_TransformKind(xn *xmlx.Node, obj *cdom.TransformKind) {
}

func load_KxAttachmentKind(xn *xmlx.Node, obj *cdom.KxAttachmentKind) {
}

func load_FxFormatRange(xn *xmlx.Node, obj *cdom.FxFormatRange) {
}

func load_FxFilterKind(xn *xmlx.Node, obj *cdom.FxFilterKind) {
}

func load_FxFormatPrecision(xn *xmlx.Node, obj *cdom.FxFormatPrecision) {
}

func load_FxFormatChannels(xn *xmlx.Node, obj *cdom.FxFormatChannels) {
}

func load_FxShaderStage(xn *xmlx.Node, obj *cdom.FxShaderStage) {
}

func load_FxCubeFace(xn *xmlx.Node, obj *cdom.FxCubeFace) {
}

func load_KxJointKind(xn *xmlx.Node, obj *cdom.KxJointKind) {
}

func load_FxSamplerKind(xn *xmlx.Node, obj *cdom.FxSamplerKind) {
}

func load_FxTechniqueKind(xn *xmlx.Node, obj *cdom.FxTechniqueKind) {
}

func load_FxTextureOpaque(xn *xmlx.Node, obj *cdom.FxTextureOpaque) {
}

func load_AnimSamplerBehavior(xn *xmlx.Node, obj *cdom.AnimSamplerBehavior) {
}

func init_SamplerWrapping(xn *xmlx.Node) (obj *ugfx.SamplerWrapping) {
	obj = new(ugfx.SamplerWrapping)

	load_SamplerWrapping(xn, obj)
	return
}

func obj_SamplerWrapping(xn *xmlx.Node, n string) (obj *ugfx.SamplerWrapping) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_SamplerWrapping(xn)
	}
	return
}

func objs_SamplerWrapping(xn *xmlx.Node, n string) (objs []*ugfx.SamplerWrapping) {
	xns := xcns(xn, n)
	objs = make([]*ugfx.SamplerWrapping, len(xns))
	for i, xn := range xns {
		objs[i] = obj_SamplerWrapping(xn, "")
	}
	return
}
