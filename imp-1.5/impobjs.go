package collimp

import (
	xmlx "github.com/jteeuwen/go-pkg-xmlx"

	cdom "github.com/go3d/go-collada/dom"
)

func obj_GeometryBrepOrientation(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepOrientation) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepOrientation(xn)
	}
	return
}

func objs_GeometryBrepOrientation(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepOrientation) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepOrientation, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepOrientation(xn, "")
	}
	return
}

func obj_FxMaterialInst(xn *xmlx.Node, n string) (obj *cdom.FxMaterialInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxMaterialInst(xn)
	}
	return
}

func objs_FxMaterialInst(xn *xmlx.Node, n string) (objs []*cdom.FxMaterialInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxMaterialInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxMaterialInst(xn, "")
	}
	return
}

func obj_FxFormatPrecision(xn *xmlx.Node, n string) (obj *cdom.FxFormatPrecision) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxFormatPrecision(xn)
	}
	return
}

func objs_FxFormatPrecision(xn *xmlx.Node, n string) (objs []*cdom.FxFormatPrecision) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxFormatPrecision, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxFormatPrecision(xn, "")
	}
	return
}

func obj_GeometryPolygonHole(xn *xmlx.Node, n string) (obj *cdom.GeometryPolygonHole) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryPolygonHole(xn)
	}
	return
}

func objs_GeometryPolygonHole(xn *xmlx.Node, n string) (objs []*cdom.GeometryPolygonHole) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryPolygonHole, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryPolygonHole(xn, "")
	}
	return
}

func obj_KxJointDef(xn *xmlx.Node, n string) (obj *cdom.KxJointDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxJointDef(xn)
	}
	return
}

func objs_KxJointDef(xn *xmlx.Node, n string) (objs []*cdom.KxJointDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxJointDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxJointDef(xn, "")
	}
	return
}

func obj_FxFilterKind(xn *xmlx.Node, n string) (obj *cdom.FxFilterKind) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxFilterKind(xn)
	}
	return
}

func objs_FxFilterKind(xn *xmlx.Node, n string) (objs []*cdom.FxFilterKind) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxFilterKind, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxFilterKind(xn, "")
	}
	return
}

func obj_GeometryBrepCircle(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepCircle) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepCircle(xn)
	}
	return
}

func objs_GeometryBrepCircle(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepCircle) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepCircle, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepCircle(xn, "")
	}
	return
}

func obj_ControllerInputs(xn *xmlx.Node, n string) (obj *cdom.ControllerInputs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ControllerInputs(xn)
	}
	return
}

func objs_ControllerInputs(xn *xmlx.Node, n string) (objs []*cdom.ControllerInputs) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ControllerInputs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ControllerInputs(xn, "")
	}
	return
}

func obj_FxInitFrom(xn *xmlx.Node, n string) (obj *cdom.FxInitFrom) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxInitFrom(xn)
	}
	return
}

func objs_FxInitFrom(xn *xmlx.Node, n string) (objs []*cdom.FxInitFrom) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxInitFrom, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxInitFrom(xn, "")
	}
	return
}

func obj_Asset(xn *xmlx.Node, n string) (obj *cdom.Asset) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Asset(xn)
	}
	return
}

func objs_Asset(xn *xmlx.Node, n string) (objs []*cdom.Asset) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Asset, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Asset(xn, "")
	}
	return
}

func obj_LightPoint(xn *xmlx.Node, n string) (obj *cdom.LightPoint) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_LightPoint(xn)
	}
	return
}

func objs_LightPoint(xn *xmlx.Node, n string) (objs []*cdom.LightPoint) {
	xns := xcns(xn, n)
	objs = make([]*cdom.LightPoint, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightPoint(xn, "")
	}
	return
}

func obj_GeometryBrepSphere(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepSphere) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepSphere(xn)
	}
	return
}

func objs_GeometryBrepSphere(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepSphere) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepSphere, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepSphere(xn, "")
	}
	return
}

func obj_KxMotionSystem(xn *xmlx.Node, n string) (obj *cdom.KxMotionSystem) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxMotionSystem(xn)
	}
	return
}

func objs_KxMotionSystem(xn *xmlx.Node, n string) (objs []*cdom.KxMotionSystem) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxMotionSystem, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxMotionSystem(xn, "")
	}
	return
}

func obj_KxEffector(xn *xmlx.Node, n string) (obj *cdom.KxEffector) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxEffector(xn)
	}
	return
}

func objs_KxEffector(xn *xmlx.Node, n string) (objs []*cdom.KxEffector) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxEffector, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxEffector(xn, "")
	}
	return
}

func obj_CameraOptics(xn *xmlx.Node, n string) (obj *cdom.CameraOptics) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_CameraOptics(xn)
	}
	return
}

func objs_CameraOptics(xn *xmlx.Node, n string) (objs []*cdom.CameraOptics) {
	xns := xcns(xn, n)
	objs = make([]*cdom.CameraOptics, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraOptics(xn, "")
	}
	return
}

func obj_KxJointAxisBinding(xn *xmlx.Node, n string) (obj *cdom.KxJointAxisBinding) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxJointAxisBinding(xn)
	}
	return
}

func objs_KxJointAxisBinding(xn *xmlx.Node, n string) (objs []*cdom.KxJointAxisBinding) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxJointAxisBinding, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxJointAxisBinding(xn, "")
	}
	return
}

func obj_KxMotionAxis(xn *xmlx.Node, n string) (obj *cdom.KxMotionAxis) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxMotionAxis(xn)
	}
	return
}

func objs_KxMotionAxis(xn *xmlx.Node, n string) (objs []*cdom.KxMotionAxis) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxMotionAxis, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxMotionAxis(xn, "")
	}
	return
}

func obj_FxTextureOpaque(xn *xmlx.Node, n string) (obj *cdom.FxTextureOpaque) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxTextureOpaque(xn)
	}
	return
}

func objs_FxTextureOpaque(xn *xmlx.Node, n string) (objs []*cdom.FxTextureOpaque) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxTextureOpaque, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTextureOpaque(xn, "")
	}
	return
}

func obj_Float7(xn *xmlx.Node, n string) (obj *cdom.Float7) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float7(xn)
	}
	return
}

func objs_Float7(xn *xmlx.Node, n string) (objs []*cdom.Float7) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float7, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float7(xn, "")
	}
	return
}

func obj_FxCreateMips(xn *xmlx.Node, n string) (obj *cdom.FxCreateMips) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCreateMips(xn)
	}
	return
}

func objs_FxCreateMips(xn *xmlx.Node, n string) (objs []*cdom.FxCreateMips) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCreateMips, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateMips(xn, "")
	}
	return
}

func obj_KxLink(xn *xmlx.Node, n string) (obj *cdom.KxLink) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxLink(xn)
	}
	return
}

func objs_KxLink(xn *xmlx.Node, n string) (objs []*cdom.KxLink) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxLink, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxLink(xn, "")
	}
	return
}

func obj_FxEffectInst(xn *xmlx.Node, n string) (obj *cdom.FxEffectInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxEffectInst(xn)
	}
	return
}

func objs_FxEffectInst(xn *xmlx.Node, n string) (objs []*cdom.FxEffectInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxEffectInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxEffectInst(xn, "")
	}
	return
}

func obj_MaterialBinding(xn *xmlx.Node, n string) (obj *cdom.MaterialBinding) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_MaterialBinding(xn)
	}
	return
}

func objs_MaterialBinding(xn *xmlx.Node, n string) (objs []*cdom.MaterialBinding) {
	xns := xcns(xn, n)
	objs = make([]*cdom.MaterialBinding, len(xns))
	for i, xn := range xns {
		objs[i] = obj_MaterialBinding(xn, "")
	}
	return
}

func obj_PxRigidConstraintDef(xn *xmlx.Node, n string) (obj *cdom.PxRigidConstraintDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxRigidConstraintDef(xn)
	}
	return
}

func objs_PxRigidConstraintDef(xn *xmlx.Node, n string) (objs []*cdom.PxRigidConstraintDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxRigidConstraintDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidConstraintDef(xn, "")
	}
	return
}

func obj_VisualSceneRendering(xn *xmlx.Node, n string) (obj *cdom.VisualSceneRendering) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_VisualSceneRendering(xn)
	}
	return
}

func objs_VisualSceneRendering(xn *xmlx.Node, n string) (objs []*cdom.VisualSceneRendering) {
	xns := xcns(xn, n)
	objs = make([]*cdom.VisualSceneRendering, len(xns))
	for i, xn := range xns {
		objs[i] = obj_VisualSceneRendering(xn, "")
	}
	return
}

func obj_InputShared(xn *xmlx.Node, n string) (obj *cdom.InputShared) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_InputShared(xn)
	}
	return
}

func objs_InputShared(xn *xmlx.Node, n string) (objs []*cdom.InputShared) {
	xns := xcns(xn, n)
	objs = make([]*cdom.InputShared, len(xns))
	for i, xn := range xns {
		objs[i] = obj_InputShared(xn, "")
	}
	return
}

func obj_FxVertexInputBinding(xn *xmlx.Node, n string) (obj *cdom.FxVertexInputBinding) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxVertexInputBinding(xn)
	}
	return
}

func objs_FxVertexInputBinding(xn *xmlx.Node, n string) (objs []*cdom.FxVertexInputBinding) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxVertexInputBinding, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxVertexInputBinding(xn, "")
	}
	return
}

func obj_Float2x2(xn *xmlx.Node, n string) (obj *cdom.Float2x2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float2x2(xn)
	}
	return
}

func objs_Float2x2(xn *xmlx.Node, n string) (objs []*cdom.Float2x2) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float2x2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float2x2(xn, "")
	}
	return
}

func obj_CameraOrthographic(xn *xmlx.Node, n string) (obj *cdom.CameraOrthographic) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_CameraOrthographic(xn)
	}
	return
}

func objs_CameraOrthographic(xn *xmlx.Node, n string) (objs []*cdom.CameraOrthographic) {
	xns := xcns(xn, n)
	objs = make([]*cdom.CameraOrthographic, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraOrthographic(xn, "")
	}
	return
}

func obj_PxRigidBodyInst(xn *xmlx.Node, n string) (obj *cdom.PxRigidBodyInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxRigidBodyInst(xn)
	}
	return
}

func objs_PxRigidBodyInst(xn *xmlx.Node, n string) (objs []*cdom.PxRigidBodyInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxRigidBodyInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidBodyInst(xn, "")
	}
	return
}

func obj_FxCreateCubeInitFrom(xn *xmlx.Node, n string) (obj *cdom.FxCreateCubeInitFrom) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCreateCubeInitFrom(xn)
	}
	return
}

func objs_FxCreateCubeInitFrom(xn *xmlx.Node, n string) (objs []*cdom.FxCreateCubeInitFrom) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCreateCubeInitFrom, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateCubeInitFrom(xn, "")
	}
	return
}

func obj_GeometryMesh(xn *xmlx.Node, n string) (obj *cdom.GeometryMesh) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryMesh(xn)
	}
	return
}

func objs_GeometryMesh(xn *xmlx.Node, n string) (objs []*cdom.GeometryMesh) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryMesh, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryMesh(xn, "")
	}
	return
}

func obj_GeometryBrepCylinder(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepCylinder) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepCylinder(xn)
	}
	return
}

func objs_GeometryBrepCylinder(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepCylinder) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepCylinder, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepCylinder(xn, "")
	}
	return
}

func obj_FxPassEvaluationClearStencil(xn *xmlx.Node, n string) (obj *cdom.FxPassEvaluationClearStencil) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxPassEvaluationClearStencil(xn)
	}
	return
}

func objs_FxPassEvaluationClearStencil(xn *xmlx.Node, n string) (objs []*cdom.FxPassEvaluationClearStencil) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxPassEvaluationClearStencil, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassEvaluationClearStencil(xn, "")
	}
	return
}

func obj_FxCreateFormat(xn *xmlx.Node, n string) (obj *cdom.FxCreateFormat) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCreateFormat(xn)
	}
	return
}

func objs_FxCreateFormat(xn *xmlx.Node, n string) (objs []*cdom.FxCreateFormat) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCreateFormat, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateFormat(xn, "")
	}
	return
}

func obj_KxModelBinding(xn *xmlx.Node, n string) (obj *cdom.KxModelBinding) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxModelBinding(xn)
	}
	return
}

func objs_KxModelBinding(xn *xmlx.Node, n string) (objs []*cdom.KxModelBinding) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxModelBinding, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxModelBinding(xn, "")
	}
	return
}

func obj_FxPassProgramShader(xn *xmlx.Node, n string) (obj *cdom.FxPassProgramShader) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxPassProgramShader(xn)
	}
	return
}

func objs_FxPassProgramShader(xn *xmlx.Node, n string) (objs []*cdom.FxPassProgramShader) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxPassProgramShader, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassProgramShader(xn, "")
	}
	return
}

func obj_Float4x4(xn *xmlx.Node, n string) (obj *cdom.Float4x4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float4x4(xn)
	}
	return
}

func objs_Float4x4(xn *xmlx.Node, n string) (objs []*cdom.Float4x4) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float4x4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float4x4(xn, "")
	}
	return
}

func obj_Param(xn *xmlx.Node, n string) (obj *cdom.Param) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Param(xn)
	}
	return
}

func objs_Param(xn *xmlx.Node, n string) (objs []*cdom.Param) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Param, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Param(xn, "")
	}
	return
}

func obj_ParamOrInt(xn *xmlx.Node, n string) (obj *cdom.ParamOrInt) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ParamOrInt(xn)
	}
	return
}

func objs_ParamOrInt(xn *xmlx.Node, n string) (objs []*cdom.ParamOrInt) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ParamOrInt, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamOrInt(xn, "")
	}
	return
}

func obj_AssetGeographicLocation(xn *xmlx.Node, n string) (obj *cdom.AssetGeographicLocation) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_AssetGeographicLocation(xn)
	}
	return
}

func objs_AssetGeographicLocation(xn *xmlx.Node, n string) (objs []*cdom.AssetGeographicLocation) {
	xns := xcns(xn, n)
	objs = make([]*cdom.AssetGeographicLocation, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AssetGeographicLocation(xn, "")
	}
	return
}

func obj_ControllerSkin(xn *xmlx.Node, n string) (obj *cdom.ControllerSkin) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ControllerSkin(xn)
	}
	return
}

func objs_ControllerSkin(xn *xmlx.Node, n string) (objs []*cdom.ControllerSkin) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ControllerSkin, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ControllerSkin(xn, "")
	}
	return
}

func obj_AssetContributor(xn *xmlx.Node, n string) (obj *cdom.AssetContributor) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_AssetContributor(xn)
	}
	return
}

func objs_AssetContributor(xn *xmlx.Node, n string) (objs []*cdom.AssetContributor) {
	xns := xcns(xn, n)
	objs = make([]*cdom.AssetContributor, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AssetContributor(xn, "")
	}
	return
}

func obj_VisualSceneInst(xn *xmlx.Node, n string) (obj *cdom.VisualSceneInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_VisualSceneInst(xn)
	}
	return
}

func objs_VisualSceneInst(xn *xmlx.Node, n string) (objs []*cdom.VisualSceneInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.VisualSceneInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_VisualSceneInst(xn, "")
	}
	return
}

func obj_KxFrame(xn *xmlx.Node, n string) (obj *cdom.KxFrame) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxFrame(xn)
	}
	return
}

func objs_KxFrame(xn *xmlx.Node, n string) (objs []*cdom.KxFrame) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxFrame, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxFrame(xn, "")
	}
	return
}

func obj_PxRigidConstraintDefs(xn *xmlx.Node, n string) (obj *cdom.PxRigidConstraintDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxRigidConstraintDefs(xn)
	}
	return
}

func objs_PxRigidConstraintDefs(xn *xmlx.Node, n string) (objs []*cdom.PxRigidConstraintDefs) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxRigidConstraintDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidConstraintDefs(xn, "")
	}
	return
}

func obj_PxSceneInst(xn *xmlx.Node, n string) (obj *cdom.PxSceneInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxSceneInst(xn)
	}
	return
}

func objs_PxSceneInst(xn *xmlx.Node, n string) (objs []*cdom.PxSceneInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxSceneInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxSceneInst(xn, "")
	}
	return
}

func obj_Float2x3(xn *xmlx.Node, n string) (obj *cdom.Float2x3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float2x3(xn)
	}
	return
}

func objs_Float2x3(xn *xmlx.Node, n string) (objs []*cdom.Float2x3) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float2x3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float2x3(xn, "")
	}
	return
}

func obj_GeometryBrepWires(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepWires) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepWires(xn)
	}
	return
}

func objs_GeometryBrepWires(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepWires) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepWires, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepWires(xn, "")
	}
	return
}

func obj_LightAmbient(xn *xmlx.Node, n string) (obj *cdom.LightAmbient) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_LightAmbient(xn)
	}
	return
}

func objs_LightAmbient(xn *xmlx.Node, n string) (objs []*cdom.LightAmbient) {
	xns := xcns(xn, n)
	objs = make([]*cdom.LightAmbient, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightAmbient(xn, "")
	}
	return
}

func obj_GeometryBrepNurbs(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepNurbs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepNurbs(xn)
	}
	return
}

func objs_GeometryBrepNurbs(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepNurbs) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepNurbs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepNurbs(xn, "")
	}
	return
}

func obj_SourceArray(xn *xmlx.Node, n string) (obj *cdom.SourceArray) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_SourceArray(xn)
	}
	return
}

func objs_SourceArray(xn *xmlx.Node, n string) (objs []*cdom.SourceArray) {
	xns := xcns(xn, n)
	objs = make([]*cdom.SourceArray, len(xns))
	for i, xn := range xns {
		objs[i] = obj_SourceArray(xn, "")
	}
	return
}

func obj_CameraInst(xn *xmlx.Node, n string) (obj *cdom.CameraInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_CameraInst(xn)
	}
	return
}

func objs_CameraInst(xn *xmlx.Node, n string) (objs []*cdom.CameraInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.CameraInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraInst(xn, "")
	}
	return
}

func obj_PxCylinder(xn *xmlx.Node, n string) (obj *cdom.PxCylinder) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxCylinder(xn)
	}
	return
}

func objs_PxCylinder(xn *xmlx.Node, n string) (objs []*cdom.PxCylinder) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxCylinder, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxCylinder(xn, "")
	}
	return
}

func obj_KxFrameTcp(xn *xmlx.Node, n string) (obj *cdom.KxFrameTcp) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxFrameTcp(xn)
	}
	return
}

func objs_KxFrameTcp(xn *xmlx.Node, n string) (objs []*cdom.KxFrameTcp) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxFrameTcp, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxFrameTcp(xn, "")
	}
	return
}

func obj_Int2x2(xn *xmlx.Node, n string) (obj *cdom.Int2x2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Int2x2(xn)
	}
	return
}

func objs_Int2x2(xn *xmlx.Node, n string) (objs []*cdom.Int2x2) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Int2x2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Int2x2(xn, "")
	}
	return
}

func obj_GeometryBrepEdges(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepEdges) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepEdges(xn)
	}
	return
}

func objs_GeometryBrepEdges(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepEdges) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepEdges, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepEdges(xn, "")
	}
	return
}

func obj_AnimSamplerBehavior(xn *xmlx.Node, n string) (obj *cdom.AnimSamplerBehavior) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_AnimSamplerBehavior(xn)
	}
	return
}

func objs_AnimSamplerBehavior(xn *xmlx.Node, n string) (objs []*cdom.AnimSamplerBehavior) {
	xns := xcns(xn, n)
	objs = make([]*cdom.AnimSamplerBehavior, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AnimSamplerBehavior(xn, "")
	}
	return
}

func obj_Input(xn *xmlx.Node, n string) (obj *cdom.Input) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Input(xn)
	}
	return
}

func objs_Input(xn *xmlx.Node, n string) (objs []*cdom.Input) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Input, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Input(xn, "")
	}
	return
}

func obj_FxPass(xn *xmlx.Node, n string) (obj *cdom.FxPass) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxPass(xn)
	}
	return
}

func objs_FxPass(xn *xmlx.Node, n string) (objs []*cdom.FxPass) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxPass, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPass(xn, "")
	}
	return
}

func obj_GeometryBrepSweptSurface(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepSweptSurface) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepSweptSurface(xn)
	}
	return
}

func objs_GeometryBrepSweptSurface(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepSweptSurface) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepSweptSurface, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepSweptSurface(xn, "")
	}
	return
}

func obj_KxFrameTip(xn *xmlx.Node, n string) (obj *cdom.KxFrameTip) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxFrameTip(xn)
	}
	return
}

func objs_KxFrameTip(xn *xmlx.Node, n string) (objs []*cdom.KxFrameTip) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxFrameTip, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxFrameTip(xn, "")
	}
	return
}

func obj_GeometryBrepPlane(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepPlane) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepPlane(xn)
	}
	return
}

func objs_GeometryBrepPlane(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepPlane) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepPlane, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepPlane(xn, "")
	}
	return
}

func obj_FxSamplerKind(xn *xmlx.Node, n string) (obj *cdom.FxSamplerKind) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxSamplerKind(xn)
	}
	return
}

func objs_FxSamplerKind(xn *xmlx.Node, n string) (objs []*cdom.FxSamplerKind) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxSamplerKind, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxSamplerKind(xn, "")
	}
	return
}

func obj_KxJointInst(xn *xmlx.Node, n string) (obj *cdom.KxJointInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxJointInst(xn)
	}
	return
}

func objs_KxJointInst(xn *xmlx.Node, n string) (objs []*cdom.KxJointInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxJointInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxJointInst(xn, "")
	}
	return
}

func obj_GeometryBrepPcurves(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepPcurves) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepPcurves(xn)
	}
	return
}

func objs_GeometryBrepPcurves(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepPcurves) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepPcurves, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepPcurves(xn, "")
	}
	return
}

func obj_Technique(xn *xmlx.Node, n string) (obj *cdom.Technique) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Technique(xn)
	}
	return
}

func objs_Technique(xn *xmlx.Node, n string) (objs []*cdom.Technique) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Technique, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Technique(xn, "")
	}
	return
}

func obj_GeometryBrepSurfaceCurves(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepSurfaceCurves) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepSurfaceCurves(xn)
	}
	return
}

func objs_GeometryBrepSurfaceCurves(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepSurfaceCurves) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepSurfaceCurves, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepSurfaceCurves(xn, "")
	}
	return
}

func obj_Float2x4(xn *xmlx.Node, n string) (obj *cdom.Float2x4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float2x4(xn)
	}
	return
}

func objs_Float2x4(xn *xmlx.Node, n string) (objs []*cdom.Float2x4) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float2x4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float2x4(xn, "")
	}
	return
}

func obj_FxCreate3D(xn *xmlx.Node, n string) (obj *cdom.FxCreate3D) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCreate3D(xn)
	}
	return
}

func objs_FxCreate3D(xn *xmlx.Node, n string) (objs []*cdom.FxCreate3D) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCreate3D, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreate3D(xn, "")
	}
	return
}

func obj_SidVec3(xn *xmlx.Node, n string) (obj *cdom.SidVec3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_SidVec3(xn)
	}
	return
}

func objs_SidVec3(xn *xmlx.Node, n string) (objs []*cdom.SidVec3) {
	xns := xcns(xn, n)
	objs = make([]*cdom.SidVec3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_SidVec3(xn, "")
	}
	return
}

func obj_PxRigidConstraintAttachment(xn *xmlx.Node, n string) (obj *cdom.PxRigidConstraintAttachment) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxRigidConstraintAttachment(xn)
	}
	return
}

func objs_PxRigidConstraintAttachment(xn *xmlx.Node, n string) (objs []*cdom.PxRigidConstraintAttachment) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxRigidConstraintAttachment, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidConstraintAttachment(xn, "")
	}
	return
}

func obj_PxModelDef(xn *xmlx.Node, n string) (obj *cdom.PxModelDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxModelDef(xn)
	}
	return
}

func objs_PxModelDef(xn *xmlx.Node, n string) (objs []*cdom.PxModelDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxModelDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxModelDef(xn, "")
	}
	return
}

func obj_Bool3(xn *xmlx.Node, n string) (obj *cdom.Bool3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Bool3(xn)
	}
	return
}

func objs_Bool3(xn *xmlx.Node, n string) (objs []*cdom.Bool3) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Bool3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Bool3(xn, "")
	}
	return
}

func obj_GeometryBrepParabola(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepParabola) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepParabola(xn)
	}
	return
}

func objs_GeometryBrepParabola(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepParabola) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepParabola, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepParabola(xn, "")
	}
	return
}

func obj_Document(xn *xmlx.Node, n string) (obj *cdom.Document) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Document(xn)
	}
	return
}

func objs_Document(xn *xmlx.Node, n string) (objs []*cdom.Document) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Document, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Document(xn, "")
	}
	return
}

func obj_SidString(xn *xmlx.Node, n string) (obj *cdom.SidString) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_SidString(xn)
	}
	return
}

func objs_SidString(xn *xmlx.Node, n string) (objs []*cdom.SidString) {
	xns := xcns(xn, n)
	objs = make([]*cdom.SidString, len(xns))
	for i, xn := range xns {
		objs[i] = obj_SidString(xn, "")
	}
	return
}

func obj_LightSpot(xn *xmlx.Node, n string) (obj *cdom.LightSpot) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_LightSpot(xn)
	}
	return
}

func objs_LightSpot(xn *xmlx.Node, n string) (objs []*cdom.LightSpot) {
	xns := xcns(xn, n)
	objs = make([]*cdom.LightSpot, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightSpot(xn, "")
	}
	return
}

func obj_ParamDef(xn *xmlx.Node, n string) (obj *cdom.ParamDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ParamDef(xn)
	}
	return
}

func objs_ParamDef(xn *xmlx.Node, n string) (objs []*cdom.ParamDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ParamDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamDef(xn, "")
	}
	return
}

func obj_PxSceneDef(xn *xmlx.Node, n string) (obj *cdom.PxSceneDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxSceneDef(xn)
	}
	return
}

func objs_PxSceneDef(xn *xmlx.Node, n string) (objs []*cdom.PxSceneDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxSceneDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxSceneDef(xn, "")
	}
	return
}

func obj_FxImageInst(xn *xmlx.Node, n string) (obj *cdom.FxImageInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxImageInst(xn)
	}
	return
}

func objs_FxImageInst(xn *xmlx.Node, n string) (objs []*cdom.FxImageInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxImageInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxImageInst(xn, "")
	}
	return
}

func obj_PxRigidBodyCommon(xn *xmlx.Node, n string) (obj *cdom.PxRigidBodyCommon) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxRigidBodyCommon(xn)
	}
	return
}

func objs_PxRigidBodyCommon(xn *xmlx.Node, n string) (objs []*cdom.PxRigidBodyCommon) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxRigidBodyCommon, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidBodyCommon(xn, "")
	}
	return
}

func obj_NodeInst(xn *xmlx.Node, n string) (obj *cdom.NodeInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_NodeInst(xn)
	}
	return
}

func objs_NodeInst(xn *xmlx.Node, n string) (objs []*cdom.NodeInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.NodeInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_NodeInst(xn, "")
	}
	return
}

func obj_GeometryBrepLine(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepLine) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepLine(xn)
	}
	return
}

func objs_GeometryBrepLine(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepLine) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepLine, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepLine(xn, "")
	}
	return
}

func obj_KxAttachmentKind(xn *xmlx.Node, n string) (obj *cdom.KxAttachmentKind) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxAttachmentKind(xn)
	}
	return
}

func objs_KxAttachmentKind(xn *xmlx.Node, n string) (objs []*cdom.KxAttachmentKind) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxAttachmentKind, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxAttachmentKind(xn, "")
	}
	return
}

func obj_CameraImager(xn *xmlx.Node, n string) (obj *cdom.CameraImager) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_CameraImager(xn)
	}
	return
}

func objs_CameraImager(xn *xmlx.Node, n string) (objs []*cdom.CameraImager) {
	xns := xcns(xn, n)
	objs = make([]*cdom.CameraImager, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraImager(xn, "")
	}
	return
}

func obj_FxTechniqueGlsl(xn *xmlx.Node, n string) (obj *cdom.FxTechniqueGlsl) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxTechniqueGlsl(xn)
	}
	return
}

func objs_FxTechniqueGlsl(xn *xmlx.Node, n string) (objs []*cdom.FxTechniqueGlsl) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxTechniqueGlsl, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTechniqueGlsl(xn, "")
	}
	return
}

func obj_LightInst(xn *xmlx.Node, n string) (obj *cdom.LightInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_LightInst(xn)
	}
	return
}

func objs_LightInst(xn *xmlx.Node, n string) (objs []*cdom.LightInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.LightInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightInst(xn, "")
	}
	return
}

func obj_Formula(xn *xmlx.Node, n string) (obj *cdom.Formula) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Formula(xn)
	}
	return
}

func objs_Formula(xn *xmlx.Node, n string) (objs []*cdom.Formula) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Formula, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Formula(xn, "")
	}
	return
}

func obj_LightAttenuation(xn *xmlx.Node, n string) (obj *cdom.LightAttenuation) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_LightAttenuation(xn)
	}
	return
}

func objs_LightAttenuation(xn *xmlx.Node, n string) (objs []*cdom.LightAttenuation) {
	xns := xcns(xn, n)
	objs = make([]*cdom.LightAttenuation, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightAttenuation(xn, "")
	}
	return
}

func obj_Bool4(xn *xmlx.Node, n string) (obj *cdom.Bool4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Bool4(xn)
	}
	return
}

func objs_Bool4(xn *xmlx.Node, n string) (objs []*cdom.Bool4) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Bool4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Bool4(xn, "")
	}
	return
}

func obj_FxImageDef(xn *xmlx.Node, n string) (obj *cdom.FxImageDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxImageDef(xn)
	}
	return
}

func objs_FxImageDef(xn *xmlx.Node, n string) (objs []*cdom.FxImageDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxImageDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxImageDef(xn, "")
	}
	return
}

func obj_FormulaDef(xn *xmlx.Node, n string) (obj *cdom.FormulaDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FormulaDef(xn)
	}
	return
}

func objs_FormulaDef(xn *xmlx.Node, n string) (objs []*cdom.FormulaDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FormulaDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FormulaDef(xn, "")
	}
	return
}

func obj_KxArticulatedSystemDef(xn *xmlx.Node, n string) (obj *cdom.KxArticulatedSystemDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxArticulatedSystemDef(xn)
	}
	return
}

func objs_KxArticulatedSystemDef(xn *xmlx.Node, n string) (objs []*cdom.KxArticulatedSystemDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxArticulatedSystemDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxArticulatedSystemDef(xn, "")
	}
	return
}

func obj_PxForceFieldDef(xn *xmlx.Node, n string) (obj *cdom.PxForceFieldDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxForceFieldDef(xn)
	}
	return
}

func objs_PxForceFieldDef(xn *xmlx.Node, n string) (objs []*cdom.PxForceFieldDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxForceFieldDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxForceFieldDef(xn, "")
	}
	return
}

func obj_FxSamplerFiltering(xn *xmlx.Node, n string) (obj *cdom.FxSamplerFiltering) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxSamplerFiltering(xn)
	}
	return
}

func objs_FxSamplerFiltering(xn *xmlx.Node, n string) (objs []*cdom.FxSamplerFiltering) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxSamplerFiltering, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxSamplerFiltering(xn, "")
	}
	return
}

func obj_NodeDef(xn *xmlx.Node, n string) (obj *cdom.NodeDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_NodeDef(xn)
	}
	return
}

func objs_NodeDef(xn *xmlx.Node, n string) (objs []*cdom.NodeDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.NodeDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_NodeDef(xn, "")
	}
	return
}

func obj_CameraDef(xn *xmlx.Node, n string) (obj *cdom.CameraDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_CameraDef(xn)
	}
	return
}

func objs_CameraDef(xn *xmlx.Node, n string) (objs []*cdom.CameraDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.CameraDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraDef(xn, "")
	}
	return
}

func obj_Float4x3(xn *xmlx.Node, n string) (obj *cdom.Float4x3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float4x3(xn)
	}
	return
}

func objs_Float4x3(xn *xmlx.Node, n string) (objs []*cdom.Float4x3) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float4x3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float4x3(xn, "")
	}
	return
}

func obj_ParamOrSidFloat(xn *xmlx.Node, n string) (obj *cdom.ParamOrSidFloat) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ParamOrSidFloat(xn)
	}
	return
}

func objs_ParamOrSidFloat(xn *xmlx.Node, n string) (objs []*cdom.ParamOrSidFloat) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ParamOrSidFloat, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamOrSidFloat(xn, "")
	}
	return
}

func obj_FxProfileGlsl(xn *xmlx.Node, n string) (obj *cdom.FxProfileGlsl) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxProfileGlsl(xn)
	}
	return
}

func objs_FxProfileGlsl(xn *xmlx.Node, n string) (objs []*cdom.FxProfileGlsl) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxProfileGlsl, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxProfileGlsl(xn, "")
	}
	return
}

func obj_Int3(xn *xmlx.Node, n string) (obj *cdom.Int3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Int3(xn)
	}
	return
}

func objs_Int3(xn *xmlx.Node, n string) (objs []*cdom.Int3) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Int3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Int3(xn, "")
	}
	return
}

func obj_FormulaInst(xn *xmlx.Node, n string) (obj *cdom.FormulaInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FormulaInst(xn)
	}
	return
}

func objs_FormulaInst(xn *xmlx.Node, n string) (objs []*cdom.FormulaInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FormulaInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FormulaInst(xn, "")
	}
	return
}

func obj_ControllerDef(xn *xmlx.Node, n string) (obj *cdom.ControllerDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ControllerDef(xn)
	}
	return
}

func objs_ControllerDef(xn *xmlx.Node, n string) (objs []*cdom.ControllerDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ControllerDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ControllerDef(xn, "")
	}
	return
}

func obj_FxEffectDef(xn *xmlx.Node, n string) (obj *cdom.FxEffectDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxEffectDef(xn)
	}
	return
}

func objs_FxEffectDef(xn *xmlx.Node, n string) (objs []*cdom.FxEffectDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxEffectDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxEffectDef(xn, "")
	}
	return
}

func obj_FxMaterialDef(xn *xmlx.Node, n string) (obj *cdom.FxMaterialDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxMaterialDef(xn)
	}
	return
}

func objs_FxMaterialDef(xn *xmlx.Node, n string) (objs []*cdom.FxMaterialDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxMaterialDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxMaterialDef(xn, "")
	}
	return
}

func obj_FxGlslTechniques(xn *xmlx.Node, n string) (obj *cdom.FxGlslTechniques) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxGlslTechniques(xn)
	}
	return
}

func objs_FxGlslTechniques(xn *xmlx.Node, n string) (objs []*cdom.FxGlslTechniques) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxGlslTechniques, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxGlslTechniques(xn, "")
	}
	return
}

func obj_Float2(xn *xmlx.Node, n string) (obj *cdom.Float2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float2(xn)
	}
	return
}

func objs_Float2(xn *xmlx.Node, n string) (objs []*cdom.Float2) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float2(xn, "")
	}
	return
}

func obj_FxCreateInitFrom(xn *xmlx.Node, n string) (obj *cdom.FxCreateInitFrom) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCreateInitFrom(xn)
	}
	return
}

func objs_FxCreateInitFrom(xn *xmlx.Node, n string) (objs []*cdom.FxCreateInitFrom) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCreateInitFrom, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateInitFrom(xn, "")
	}
	return
}

func obj_KxJointLimits(xn *xmlx.Node, n string) (obj *cdom.KxJointLimits) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxJointLimits(xn)
	}
	return
}

func objs_KxJointLimits(xn *xmlx.Node, n string) (objs []*cdom.KxJointLimits) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxJointLimits, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxJointLimits(xn, "")
	}
	return
}

func obj_FxEffectInstTechniqueHint(xn *xmlx.Node, n string) (obj *cdom.FxEffectInstTechniqueHint) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxEffectInstTechniqueHint(xn)
	}
	return
}

func objs_FxEffectInstTechniqueHint(xn *xmlx.Node, n string) (objs []*cdom.FxEffectInstTechniqueHint) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxEffectInstTechniqueHint, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxEffectInstTechniqueHint(xn, "")
	}
	return
}

func obj_Float3x3(xn *xmlx.Node, n string) (obj *cdom.Float3x3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float3x3(xn)
	}
	return
}

func objs_Float3x3(xn *xmlx.Node, n string) (objs []*cdom.Float3x3) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float3x3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float3x3(xn, "")
	}
	return
}

func obj_ParamDefs(xn *xmlx.Node, n string) (obj *cdom.ParamDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ParamDefs(xn)
	}
	return
}

func objs_ParamDefs(xn *xmlx.Node, n string) (objs []*cdom.ParamDefs) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ParamDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamDefs(xn, "")
	}
	return
}

func obj_FxCreateFormatHint(xn *xmlx.Node, n string) (obj *cdom.FxCreateFormatHint) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCreateFormatHint(xn)
	}
	return
}

func objs_FxCreateFormatHint(xn *xmlx.Node, n string) (objs []*cdom.FxCreateFormatHint) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCreateFormatHint, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateFormatHint(xn, "")
	}
	return
}

func obj_SourceAccessor(xn *xmlx.Node, n string) (obj *cdom.SourceAccessor) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_SourceAccessor(xn)
	}
	return
}

func objs_SourceAccessor(xn *xmlx.Node, n string) (objs []*cdom.SourceAccessor) {
	xns := xcns(xn, n)
	objs = make([]*cdom.SourceAccessor, len(xns))
	for i, xn := range xns {
		objs[i] = obj_SourceAccessor(xn, "")
	}
	return
}

func obj_ParamInsts(xn *xmlx.Node, n string) (obj *cdom.ParamInsts) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ParamInsts(xn)
	}
	return
}

func objs_ParamInsts(xn *xmlx.Node, n string) (objs []*cdom.ParamInsts) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ParamInsts, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamInsts(xn, "")
	}
	return
}

func obj_ControllerMorph(xn *xmlx.Node, n string) (obj *cdom.ControllerMorph) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ControllerMorph(xn)
	}
	return
}

func objs_ControllerMorph(xn *xmlx.Node, n string) (objs []*cdom.ControllerMorph) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ControllerMorph, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ControllerMorph(xn, "")
	}
	return
}

func obj_FxImageInitFrom(xn *xmlx.Node, n string) (obj *cdom.FxImageInitFrom) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxImageInitFrom(xn)
	}
	return
}

func objs_FxImageInitFrom(xn *xmlx.Node, n string) (objs []*cdom.FxImageInitFrom) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxImageInitFrom, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxImageInitFrom(xn, "")
	}
	return
}

func obj_GeometryVertices(xn *xmlx.Node, n string) (obj *cdom.GeometryVertices) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryVertices(xn)
	}
	return
}

func objs_GeometryVertices(xn *xmlx.Node, n string) (objs []*cdom.GeometryVertices) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryVertices, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryVertices(xn, "")
	}
	return
}

func obj_GeometryBrepTorus(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepTorus) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepTorus(xn)
	}
	return
}

func objs_GeometryBrepTorus(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepTorus) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepTorus, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepTorus(xn, "")
	}
	return
}

func obj_FxTexture(xn *xmlx.Node, n string) (obj *cdom.FxTexture) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxTexture(xn)
	}
	return
}

func objs_FxTexture(xn *xmlx.Node, n string) (objs []*cdom.FxTexture) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxTexture, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTexture(xn, "")
	}
	return
}

func obj_FxPassEvaluation(xn *xmlx.Node, n string) (obj *cdom.FxPassEvaluation) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxPassEvaluation(xn)
	}
	return
}

func objs_FxPassEvaluation(xn *xmlx.Node, n string) (objs []*cdom.FxPassEvaluation) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxPassEvaluation, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassEvaluation(xn, "")
	}
	return
}

func obj_ParamInst(xn *xmlx.Node, n string) (obj *cdom.ParamInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ParamInst(xn)
	}
	return
}

func objs_ParamInst(xn *xmlx.Node, n string) (objs []*cdom.ParamInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ParamInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamInst(xn, "")
	}
	return
}

func obj_FxProfile(xn *xmlx.Node, n string) (obj *cdom.FxProfile) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxProfile(xn)
	}
	return
}

func objs_FxProfile(xn *xmlx.Node, n string) (objs []*cdom.FxProfile) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxProfile, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxProfile(xn, "")
	}
	return
}

func obj_FxPassProgramBindAttribute(xn *xmlx.Node, n string) (obj *cdom.FxPassProgramBindAttribute) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxPassProgramBindAttribute(xn)
	}
	return
}

func objs_FxPassProgramBindAttribute(xn *xmlx.Node, n string) (objs []*cdom.FxPassProgramBindAttribute) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxPassProgramBindAttribute, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassProgramBindAttribute(xn, "")
	}
	return
}

func obj_FxCreate2D(xn *xmlx.Node, n string) (obj *cdom.FxCreate2D) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCreate2D(xn)
	}
	return
}

func objs_FxCreate2D(xn *xmlx.Node, n string) (objs []*cdom.FxCreate2D) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCreate2D, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreate2D(xn, "")
	}
	return
}

func obj_ParamOrFloat2(xn *xmlx.Node, n string) (obj *cdom.ParamOrFloat2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ParamOrFloat2(xn)
	}
	return
}

func objs_ParamOrFloat2(xn *xmlx.Node, n string) (objs []*cdom.ParamOrFloat2) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ParamOrFloat2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamOrFloat2(xn, "")
	}
	return
}

func obj_FxCreateCube(xn *xmlx.Node, n string) (obj *cdom.FxCreateCube) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCreateCube(xn)
	}
	return
}

func objs_FxCreateCube(xn *xmlx.Node, n string) (objs []*cdom.FxCreateCube) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCreateCube, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateCube(xn, "")
	}
	return
}

func obj_KxBinding(xn *xmlx.Node, n string) (obj *cdom.KxBinding) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxBinding(xn)
	}
	return
}

func objs_KxBinding(xn *xmlx.Node, n string) (objs []*cdom.KxBinding) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxBinding, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxBinding(xn, "")
	}
	return
}

func obj_GeometryBrepBox(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepBox) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepBox(xn)
	}
	return
}

func objs_GeometryBrepBox(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepBox) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepBox, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepBox(xn, "")
	}
	return
}

func obj_Float4x2(xn *xmlx.Node, n string) (obj *cdom.Float4x2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float4x2(xn)
	}
	return
}

func objs_Float4x2(xn *xmlx.Node, n string) (objs []*cdom.Float4x2) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float4x2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float4x2(xn, "")
	}
	return
}

func obj_FxCreate3DInitFrom(xn *xmlx.Node, n string) (obj *cdom.FxCreate3DInitFrom) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCreate3DInitFrom(xn)
	}
	return
}

func objs_FxCreate3DInitFrom(xn *xmlx.Node, n string) (objs []*cdom.FxCreate3DInitFrom) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCreate3DInitFrom, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreate3DInitFrom(xn, "")
	}
	return
}

func obj_FxCreate2DSizeExact(xn *xmlx.Node, n string) (obj *cdom.FxCreate2DSizeExact) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCreate2DSizeExact(xn)
	}
	return
}

func objs_FxCreate2DSizeExact(xn *xmlx.Node, n string) (objs []*cdom.FxCreate2DSizeExact) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCreate2DSizeExact, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreate2DSizeExact(xn, "")
	}
	return
}

func obj_PxMaterialDef(xn *xmlx.Node, n string) (obj *cdom.PxMaterialDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxMaterialDef(xn)
	}
	return
}

func objs_PxMaterialDef(xn *xmlx.Node, n string) (objs []*cdom.PxMaterialDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxMaterialDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxMaterialDef(xn, "")
	}
	return
}

func obj_AnimationDef(xn *xmlx.Node, n string) (obj *cdom.AnimationDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_AnimationDef(xn)
	}
	return
}

func objs_AnimationDef(xn *xmlx.Node, n string) (objs []*cdom.AnimationDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.AnimationDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AnimationDef(xn, "")
	}
	return
}

func obj_FxParamDef(xn *xmlx.Node, n string) (obj *cdom.FxParamDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxParamDef(xn)
	}
	return
}

func objs_FxParamDef(xn *xmlx.Node, n string) (objs []*cdom.FxParamDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxParamDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxParamDef(xn, "")
	}
	return
}

func obj_ChildNode(xn *xmlx.Node, n string) (obj *cdom.ChildNode) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ChildNode(xn)
	}
	return
}

func objs_ChildNode(xn *xmlx.Node, n string) (objs []*cdom.ChildNode) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ChildNode, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ChildNode(xn, "")
	}
	return
}

func obj_PxRigidBodyDefs(xn *xmlx.Node, n string) (obj *cdom.PxRigidBodyDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxRigidBodyDefs(xn)
	}
	return
}

func objs_PxRigidBodyDefs(xn *xmlx.Node, n string) (objs []*cdom.PxRigidBodyDefs) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxRigidBodyDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidBodyDefs(xn, "")
	}
	return
}

func obj_FxSamplerImage(xn *xmlx.Node, n string) (obj *cdom.FxSamplerImage) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxSamplerImage(xn)
	}
	return
}

func objs_FxSamplerImage(xn *xmlx.Node, n string) (objs []*cdom.FxSamplerImage) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxSamplerImage, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxSamplerImage(xn, "")
	}
	return
}

func obj_FxPassEvaluationClearColor(xn *xmlx.Node, n string) (obj *cdom.FxPassEvaluationClearColor) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxPassEvaluationClearColor(xn)
	}
	return
}

func objs_FxPassEvaluationClearColor(xn *xmlx.Node, n string) (objs []*cdom.FxPassEvaluationClearColor) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxPassEvaluationClearColor, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassEvaluationClearColor(xn, "")
	}
	return
}

func obj_Layers(xn *xmlx.Node, n string) (obj *cdom.Layers) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Layers(xn)
	}
	return
}

func objs_Layers(xn *xmlx.Node, n string) (objs []*cdom.Layers) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Layers, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Layers(xn, "")
	}
	return
}

func obj_GeometryBrepFaces(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepFaces) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepFaces(xn)
	}
	return
}

func objs_GeometryBrepFaces(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepFaces) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepFaces, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepFaces(xn, "")
	}
	return
}

func obj_FxCreate2DSizeRatio(xn *xmlx.Node, n string) (obj *cdom.FxCreate2DSizeRatio) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCreate2DSizeRatio(xn)
	}
	return
}

func objs_FxCreate2DSizeRatio(xn *xmlx.Node, n string) (objs []*cdom.FxCreate2DSizeRatio) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCreate2DSizeRatio, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreate2DSizeRatio(xn, "")
	}
	return
}

func obj_FxColor(xn *xmlx.Node, n string) (obj *cdom.FxColor) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxColor(xn)
	}
	return
}

func objs_FxColor(xn *xmlx.Node, n string) (objs []*cdom.FxColor) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxColor, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxColor(xn, "")
	}
	return
}

func obj_FxParamDefs(xn *xmlx.Node, n string) (obj *cdom.FxParamDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxParamDefs(xn)
	}
	return
}

func objs_FxParamDefs(xn *xmlx.Node, n string) (objs []*cdom.FxParamDefs) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxParamDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxParamDefs(xn, "")
	}
	return
}

func obj_GeometryBrepCone(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepCone) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepCone(xn)
	}
	return
}

func objs_GeometryBrepCone(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepCone) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepCone, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepCone(xn, "")
	}
	return
}

func obj_LightDirectional(xn *xmlx.Node, n string) (obj *cdom.LightDirectional) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_LightDirectional(xn)
	}
	return
}

func objs_LightDirectional(xn *xmlx.Node, n string) (objs []*cdom.LightDirectional) {
	xns := xcns(xn, n)
	objs = make([]*cdom.LightDirectional, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightDirectional(xn, "")
	}
	return
}

func obj_KxSceneInst(xn *xmlx.Node, n string) (obj *cdom.KxSceneInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxSceneInst(xn)
	}
	return
}

func objs_KxSceneInst(xn *xmlx.Node, n string) (objs []*cdom.KxSceneInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxSceneInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxSceneInst(xn, "")
	}
	return
}

func obj_FxSampler(xn *xmlx.Node, n string) (obj *cdom.FxSampler) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxSampler(xn)
	}
	return
}

func objs_FxSampler(xn *xmlx.Node, n string) (objs []*cdom.FxSampler) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxSampler, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxSampler(xn, "")
	}
	return
}

func obj_GeometryBrep(xn *xmlx.Node, n string) (obj *cdom.GeometryBrep) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrep(xn)
	}
	return
}

func objs_GeometryBrep(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrep) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrep, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrep(xn, "")
	}
	return
}

func obj_FxTechnique(xn *xmlx.Node, n string) (obj *cdom.FxTechnique) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxTechnique(xn)
	}
	return
}

func objs_FxTechnique(xn *xmlx.Node, n string) (objs []*cdom.FxTechnique) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxTechnique, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTechnique(xn, "")
	}
	return
}

func obj_FxFormatRange(xn *xmlx.Node, n string) (obj *cdom.FxFormatRange) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxFormatRange(xn)
	}
	return
}

func objs_FxFormatRange(xn *xmlx.Node, n string) (objs []*cdom.FxFormatRange) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxFormatRange, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxFormatRange(xn, "")
	}
	return
}

func obj_Int4x4(xn *xmlx.Node, n string) (obj *cdom.Int4x4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Int4x4(xn)
	}
	return
}

func objs_Int4x4(xn *xmlx.Node, n string) (objs []*cdom.Int4x4) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Int4x4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Int4x4(xn, "")
	}
	return
}

func obj_Sources(xn *xmlx.Node, n string) (obj *cdom.Sources) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Sources(xn)
	}
	return
}

func objs_Sources(xn *xmlx.Node, n string) (objs []*cdom.Sources) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Sources, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Sources(xn, "")
	}
	return
}

func obj_FxTechniqueKind(xn *xmlx.Node, n string) (obj *cdom.FxTechniqueKind) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxTechniqueKind(xn)
	}
	return
}

func objs_FxTechniqueKind(xn *xmlx.Node, n string) (objs []*cdom.FxTechniqueKind) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxTechniqueKind, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTechniqueKind(xn, "")
	}
	return
}

func obj_AnimationSampler(xn *xmlx.Node, n string) (obj *cdom.AnimationSampler) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_AnimationSampler(xn)
	}
	return
}

func objs_AnimationSampler(xn *xmlx.Node, n string) (objs []*cdom.AnimationSampler) {
	xns := xcns(xn, n)
	objs = make([]*cdom.AnimationSampler, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AnimationSampler(xn, "")
	}
	return
}

func obj_FxCreate(xn *xmlx.Node, n string) (obj *cdom.FxCreate) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCreate(xn)
	}
	return
}

func objs_FxCreate(xn *xmlx.Node, n string) (objs []*cdom.FxCreate) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCreate, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreate(xn, "")
	}
	return
}

func obj_ParamOrBool(xn *xmlx.Node, n string) (obj *cdom.ParamOrBool) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ParamOrBool(xn)
	}
	return
}

func objs_ParamOrBool(xn *xmlx.Node, n string) (objs []*cdom.ParamOrBool) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ParamOrBool, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamOrBool(xn, "")
	}
	return
}

func obj_KxFrameObject(xn *xmlx.Node, n string) (obj *cdom.KxFrameObject) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxFrameObject(xn)
	}
	return
}

func objs_KxFrameObject(xn *xmlx.Node, n string) (objs []*cdom.KxFrameObject) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxFrameObject, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxFrameObject(xn, "")
	}
	return
}

func obj_ParamOrFloat(xn *xmlx.Node, n string) (obj *cdom.ParamOrFloat) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ParamOrFloat(xn)
	}
	return
}

func objs_ParamOrFloat(xn *xmlx.Node, n string) (objs []*cdom.ParamOrFloat) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ParamOrFloat, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamOrFloat(xn, "")
	}
	return
}

func obj_ControllerInst(xn *xmlx.Node, n string) (obj *cdom.ControllerInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ControllerInst(xn)
	}
	return
}

func objs_ControllerInst(xn *xmlx.Node, n string) (objs []*cdom.ControllerInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ControllerInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ControllerInst(xn, "")
	}
	return
}

func obj_Float3x4(xn *xmlx.Node, n string) (obj *cdom.Float3x4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float3x4(xn)
	}
	return
}

func objs_Float3x4(xn *xmlx.Node, n string) (objs []*cdom.Float3x4) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float3x4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float3x4(xn, "")
	}
	return
}

func obj_Scene(xn *xmlx.Node, n string) (obj *cdom.Scene) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Scene(xn)
	}
	return
}

func objs_Scene(xn *xmlx.Node, n string) (objs []*cdom.Scene) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Scene, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Scene(xn, "")
	}
	return
}

func obj_GeometryInst(xn *xmlx.Node, n string) (obj *cdom.GeometryInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryInst(xn)
	}
	return
}

func objs_GeometryInst(xn *xmlx.Node, n string) (objs []*cdom.GeometryInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryInst(xn, "")
	}
	return
}

func obj_IndexedInputs(xn *xmlx.Node, n string) (obj *cdom.IndexedInputs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_IndexedInputs(xn)
	}
	return
}

func objs_IndexedInputs(xn *xmlx.Node, n string) (objs []*cdom.IndexedInputs) {
	xns := xcns(xn, n)
	objs = make([]*cdom.IndexedInputs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_IndexedInputs(xn, "")
	}
	return
}

func obj_PxRigidConstraintSpring(xn *xmlx.Node, n string) (obj *cdom.PxRigidConstraintSpring) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxRigidConstraintSpring(xn)
	}
	return
}

func objs_PxRigidConstraintSpring(xn *xmlx.Node, n string) (objs []*cdom.PxRigidConstraintSpring) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxRigidConstraintSpring, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidConstraintSpring(xn, "")
	}
	return
}

func obj_GeometryBrepCapsule(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepCapsule) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepCapsule(xn)
	}
	return
}

func objs_GeometryBrepCapsule(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepCapsule) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepCapsule, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepCapsule(xn, "")
	}
	return
}

func obj_TransformKind(xn *xmlx.Node, n string) (obj *cdom.TransformKind) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_TransformKind(xn)
	}
	return
}

func objs_TransformKind(xn *xmlx.Node, n string) (objs []*cdom.TransformKind) {
	xns := xcns(xn, n)
	objs = make([]*cdom.TransformKind, len(xns))
	for i, xn := range xns {
		objs[i] = obj_TransformKind(xn, "")
	}
	return
}

func obj_PxMaterial(xn *xmlx.Node, n string) (obj *cdom.PxMaterial) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxMaterial(xn)
	}
	return
}

func objs_PxMaterial(xn *xmlx.Node, n string) (objs []*cdom.PxMaterial) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxMaterial, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxMaterial(xn, "")
	}
	return
}

func obj_GeometryPrimitives(xn *xmlx.Node, n string) (obj *cdom.GeometryPrimitives) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryPrimitives(xn)
	}
	return
}

func objs_GeometryPrimitives(xn *xmlx.Node, n string) (objs []*cdom.GeometryPrimitives) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryPrimitives, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryPrimitives(xn, "")
	}
	return
}

func obj_Transform(xn *xmlx.Node, n string) (obj *cdom.Transform) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Transform(xn)
	}
	return
}

func objs_Transform(xn *xmlx.Node, n string) (objs []*cdom.Transform) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Transform, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Transform(xn, "")
	}
	return
}

func obj_PxRigidBodyDef(xn *xmlx.Node, n string) (obj *cdom.PxRigidBodyDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxRigidBodyDef(xn)
	}
	return
}

func objs_PxRigidBodyDef(xn *xmlx.Node, n string) (objs []*cdom.PxRigidBodyDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxRigidBodyDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidBodyDef(xn, "")
	}
	return
}

func obj_KxArticulatedSystemInst(xn *xmlx.Node, n string) (obj *cdom.KxArticulatedSystemInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxArticulatedSystemInst(xn)
	}
	return
}

func objs_KxArticulatedSystemInst(xn *xmlx.Node, n string) (objs []*cdom.KxArticulatedSystemInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxArticulatedSystemInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxArticulatedSystemInst(xn, "")
	}
	return
}

func obj_FxProfileGlslCodeInclude(xn *xmlx.Node, n string) (obj *cdom.FxProfileGlslCodeInclude) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxProfileGlslCodeInclude(xn)
	}
	return
}

func objs_FxProfileGlslCodeInclude(xn *xmlx.Node, n string) (objs []*cdom.FxProfileGlslCodeInclude) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxProfileGlslCodeInclude, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxProfileGlslCodeInclude(xn, "")
	}
	return
}

func obj_GeometryDef(xn *xmlx.Node, n string) (obj *cdom.GeometryDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryDef(xn)
	}
	return
}

func objs_GeometryDef(xn *xmlx.Node, n string) (objs []*cdom.GeometryDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryDef(xn, "")
	}
	return
}

func obj_FxPassProgramShaderSources(xn *xmlx.Node, n string) (obj *cdom.FxPassProgramShaderSources) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxPassProgramShaderSources(xn)
	}
	return
}

func objs_FxPassProgramShaderSources(xn *xmlx.Node, n string) (objs []*cdom.FxPassProgramShaderSources) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxPassProgramShaderSources, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassProgramShaderSources(xn, "")
	}
	return
}

func obj_VisualSceneEvaluation(xn *xmlx.Node, n string) (obj *cdom.VisualSceneEvaluation) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_VisualSceneEvaluation(xn)
	}
	return
}

func objs_VisualSceneEvaluation(xn *xmlx.Node, n string) (objs []*cdom.VisualSceneEvaluation) {
	xns := xcns(xn, n)
	objs = make([]*cdom.VisualSceneEvaluation, len(xns))
	for i, xn := range xns {
		objs[i] = obj_VisualSceneEvaluation(xn, "")
	}
	return
}

func obj_KxSceneDef(xn *xmlx.Node, n string) (obj *cdom.KxSceneDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxSceneDef(xn)
	}
	return
}

func objs_KxSceneDef(xn *xmlx.Node, n string) (objs []*cdom.KxSceneDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxSceneDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxSceneDef(xn, "")
	}
	return
}

func obj_Float3(xn *xmlx.Node, n string) (obj *cdom.Float3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float3(xn)
	}
	return
}

func objs_Float3(xn *xmlx.Node, n string) (objs []*cdom.Float3) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float3(xn, "")
	}
	return
}

func obj_Int2(xn *xmlx.Node, n string) (obj *cdom.Int2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Int2(xn)
	}
	return
}

func objs_Int2(xn *xmlx.Node, n string) (objs []*cdom.Int2) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Int2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Int2(xn, "")
	}
	return
}

func obj_KxJoint(xn *xmlx.Node, n string) (obj *cdom.KxJoint) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxJoint(xn)
	}
	return
}

func objs_KxJoint(xn *xmlx.Node, n string) (objs []*cdom.KxJoint) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxJoint, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxJoint(xn, "")
	}
	return
}

func obj_SidFloat3(xn *xmlx.Node, n string) (obj *cdom.SidFloat3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_SidFloat3(xn)
	}
	return
}

func objs_SidFloat3(xn *xmlx.Node, n string) (objs []*cdom.SidFloat3) {
	xns := xcns(xn, n)
	objs = make([]*cdom.SidFloat3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_SidFloat3(xn, "")
	}
	return
}

func obj_GeometryBrepCurve(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepCurve) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepCurve(xn)
	}
	return
}

func objs_GeometryBrepCurve(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepCurve) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepCurve, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepCurve(xn, "")
	}
	return
}

func obj_VisualSceneDef(xn *xmlx.Node, n string) (obj *cdom.VisualSceneDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_VisualSceneDef(xn)
	}
	return
}

func objs_VisualSceneDef(xn *xmlx.Node, n string) (objs []*cdom.VisualSceneDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.VisualSceneDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_VisualSceneDef(xn, "")
	}
	return
}

func obj_FxPassState(xn *xmlx.Node, n string) (obj *cdom.FxPassState) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxPassState(xn)
	}
	return
}

func objs_FxPassState(xn *xmlx.Node, n string) (objs []*cdom.FxPassState) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxPassState, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassState(xn, "")
	}
	return
}

func obj_GeometryPrimitiveKind(xn *xmlx.Node, n string) (obj *cdom.GeometryPrimitiveKind) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryPrimitiveKind(xn)
	}
	return
}

func objs_GeometryPrimitiveKind(xn *xmlx.Node, n string) (objs []*cdom.GeometryPrimitiveKind) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryPrimitiveKind, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryPrimitiveKind(xn, "")
	}
	return
}

func obj_GeometryBrepSurface(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepSurface) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepSurface(xn)
	}
	return
}

func objs_GeometryBrepSurface(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepSurface) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepSurface, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepSurface(xn, "")
	}
	return
}

func obj_LightDef(xn *xmlx.Node, n string) (obj *cdom.LightDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_LightDef(xn)
	}
	return
}

func objs_LightDef(xn *xmlx.Node, n string) (objs []*cdom.LightDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.LightDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightDef(xn, "")
	}
	return
}

func obj_FxPassEvaluationTarget(xn *xmlx.Node, n string) (obj *cdom.FxPassEvaluationTarget) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxPassEvaluationTarget(xn)
	}
	return
}

func objs_FxPassEvaluationTarget(xn *xmlx.Node, n string) (objs []*cdom.FxPassEvaluationTarget) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxPassEvaluationTarget, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassEvaluationTarget(xn, "")
	}
	return
}

func obj_CameraPerspective(xn *xmlx.Node, n string) (obj *cdom.CameraPerspective) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_CameraPerspective(xn)
	}
	return
}

func objs_CameraPerspective(xn *xmlx.Node, n string) (objs []*cdom.CameraPerspective) {
	xns := xcns(xn, n)
	objs = make([]*cdom.CameraPerspective, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraPerspective(xn, "")
	}
	return
}

func obj_Float4(xn *xmlx.Node, n string) (obj *cdom.Float4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float4(xn)
	}
	return
}

func objs_Float4(xn *xmlx.Node, n string) (objs []*cdom.Float4) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float4(xn, "")
	}
	return
}

func obj_KxModelDef(xn *xmlx.Node, n string) (obj *cdom.KxModelDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxModelDef(xn)
	}
	return
}

func objs_KxModelDef(xn *xmlx.Node, n string) (objs []*cdom.KxModelDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxModelDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxModelDef(xn, "")
	}
	return
}

func obj_AnimationInst(xn *xmlx.Node, n string) (obj *cdom.AnimationInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_AnimationInst(xn)
	}
	return
}

func objs_AnimationInst(xn *xmlx.Node, n string) (objs []*cdom.AnimationInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.AnimationInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AnimationInst(xn, "")
	}
	return
}

func obj_FxFormatChannels(xn *xmlx.Node, n string) (obj *cdom.FxFormatChannels) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxFormatChannels(xn)
	}
	return
}

func objs_FxFormatChannels(xn *xmlx.Node, n string) (objs []*cdom.FxFormatChannels) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxFormatChannels, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxFormatChannels(xn, "")
	}
	return
}

func obj_KxKinematicsSystem(xn *xmlx.Node, n string) (obj *cdom.KxKinematicsSystem) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxKinematicsSystem(xn)
	}
	return
}

func objs_KxKinematicsSystem(xn *xmlx.Node, n string) (objs []*cdom.KxKinematicsSystem) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxKinematicsSystem, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxKinematicsSystem(xn, "")
	}
	return
}

func obj_KxKinematicsAxis(xn *xmlx.Node, n string) (obj *cdom.KxKinematicsAxis) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxKinematicsAxis(xn)
	}
	return
}

func objs_KxKinematicsAxis(xn *xmlx.Node, n string) (objs []*cdom.KxKinematicsAxis) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxKinematicsAxis, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxKinematicsAxis(xn, "")
	}
	return
}

func obj_FxBinding(xn *xmlx.Node, n string) (obj *cdom.FxBinding) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxBinding(xn)
	}
	return
}

func objs_FxBinding(xn *xmlx.Node, n string) (objs []*cdom.FxBinding) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxBinding, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxBinding(xn, "")
	}
	return
}

func obj_Float3x2(xn *xmlx.Node, n string) (obj *cdom.Float3x2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Float3x2(xn)
	}
	return
}

func objs_Float3x2(xn *xmlx.Node, n string) (objs []*cdom.Float3x2) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Float3x2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float3x2(xn, "")
	}
	return
}

func obj_FxSamplerStates(xn *xmlx.Node, n string) (obj *cdom.FxSamplerStates) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxSamplerStates(xn)
	}
	return
}

func objs_FxSamplerStates(xn *xmlx.Node, n string) (objs []*cdom.FxSamplerStates) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxSamplerStates, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxSamplerStates(xn, "")
	}
	return
}

func obj_PxForceFieldInst(xn *xmlx.Node, n string) (obj *cdom.PxForceFieldInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxForceFieldInst(xn)
	}
	return
}

func objs_PxForceFieldInst(xn *xmlx.Node, n string) (objs []*cdom.PxForceFieldInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxForceFieldInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxForceFieldInst(xn, "")
	}
	return
}

func obj_KxAxisLimits(xn *xmlx.Node, n string) (obj *cdom.KxAxisLimits) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxAxisLimits(xn)
	}
	return
}

func objs_KxAxisLimits(xn *xmlx.Node, n string) (objs []*cdom.KxAxisLimits) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxAxisLimits, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxAxisLimits(xn, "")
	}
	return
}

func obj_PxRigidConstraintLimit(xn *xmlx.Node, n string) (obj *cdom.PxRigidConstraintLimit) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxRigidConstraintLimit(xn)
	}
	return
}

func objs_PxRigidConstraintLimit(xn *xmlx.Node, n string) (objs []*cdom.PxRigidConstraintLimit) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxRigidConstraintLimit, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidConstraintLimit(xn, "")
	}
	return
}

func obj_PxMaterialInst(xn *xmlx.Node, n string) (obj *cdom.PxMaterialInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxMaterialInst(xn)
	}
	return
}

func objs_PxMaterialInst(xn *xmlx.Node, n string) (objs []*cdom.PxMaterialInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxMaterialInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxMaterialInst(xn, "")
	}
	return
}

func obj_VisualSceneRenderingMaterialInst(xn *xmlx.Node, n string) (obj *cdom.VisualSceneRenderingMaterialInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_VisualSceneRenderingMaterialInst(xn)
	}
	return
}

func objs_VisualSceneRenderingMaterialInst(xn *xmlx.Node, n string) (objs []*cdom.VisualSceneRenderingMaterialInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.VisualSceneRenderingMaterialInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_VisualSceneRenderingMaterialInst(xn, "")
	}
	return
}

func obj_AnimationClipInst(xn *xmlx.Node, n string) (obj *cdom.AnimationClipInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_AnimationClipInst(xn)
	}
	return
}

func objs_AnimationClipInst(xn *xmlx.Node, n string) (objs []*cdom.AnimationClipInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.AnimationClipInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AnimationClipInst(xn, "")
	}
	return
}

func obj_FxAnnotation(xn *xmlx.Node, n string) (obj *cdom.FxAnnotation) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxAnnotation(xn)
	}
	return
}

func objs_FxAnnotation(xn *xmlx.Node, n string) (objs []*cdom.FxAnnotation) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxAnnotation, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxAnnotation(xn, "")
	}
	return
}

func obj_GeometryBrepEllipse(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepEllipse) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepEllipse(xn)
	}
	return
}

func objs_GeometryBrepEllipse(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepEllipse) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepEllipse, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepEllipse(xn, "")
	}
	return
}

func obj_GeometryBrepSurfaces(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepSurfaces) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepSurfaces(xn)
	}
	return
}

func objs_GeometryBrepSurfaces(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepSurfaces) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepSurfaces, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepSurfaces(xn, "")
	}
	return
}

func obj_GeometryBrepCurves(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepCurves) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepCurves(xn)
	}
	return
}

func objs_GeometryBrepCurves(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepCurves) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepCurves, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepCurves(xn, "")
	}
	return
}

func obj_Int3x3(xn *xmlx.Node, n string) (obj *cdom.Int3x3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Int3x3(xn)
	}
	return
}

func objs_Int3x3(xn *xmlx.Node, n string) (objs []*cdom.Int3x3) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Int3x3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Int3x3(xn, "")
	}
	return
}

func obj_AnimationClipDef(xn *xmlx.Node, n string) (obj *cdom.AnimationClipDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_AnimationClipDef(xn)
	}
	return
}

func objs_AnimationClipDef(xn *xmlx.Node, n string) (objs []*cdom.AnimationClipDef) {
	xns := xcns(xn, n)
	objs = make([]*cdom.AnimationClipDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AnimationClipDef(xn, "")
	}
	return
}

func obj_PxRigidConstraintInst(xn *xmlx.Node, n string) (obj *cdom.PxRigidConstraintInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxRigidConstraintInst(xn)
	}
	return
}

func objs_PxRigidConstraintInst(xn *xmlx.Node, n string) (objs []*cdom.PxRigidConstraintInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxRigidConstraintInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidConstraintInst(xn, "")
	}
	return
}

func obj_GeometrySpline(xn *xmlx.Node, n string) (obj *cdom.GeometrySpline) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometrySpline(xn)
	}
	return
}

func objs_GeometrySpline(xn *xmlx.Node, n string) (objs []*cdom.GeometrySpline) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometrySpline, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometrySpline(xn, "")
	}
	return
}

func obj_Bool2(xn *xmlx.Node, n string) (obj *cdom.Bool2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Bool2(xn)
	}
	return
}

func objs_Bool2(xn *xmlx.Node, n string) (objs []*cdom.Bool2) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Bool2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Bool2(xn, "")
	}
	return
}

func obj_GeometryBrepShells(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepShells) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepShells(xn)
	}
	return
}

func objs_GeometryBrepShells(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepShells) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepShells, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepShells(xn, "")
	}
	return
}

func obj_KxFrameOrigin(xn *xmlx.Node, n string) (obj *cdom.KxFrameOrigin) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxFrameOrigin(xn)
	}
	return
}

func objs_KxFrameOrigin(xn *xmlx.Node, n string) (objs []*cdom.KxFrameOrigin) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxFrameOrigin, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxFrameOrigin(xn, "")
	}
	return
}

func obj_FxColorOrTexture(xn *xmlx.Node, n string) (obj *cdom.FxColorOrTexture) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxColorOrTexture(xn)
	}
	return
}

func objs_FxColorOrTexture(xn *xmlx.Node, n string) (objs []*cdom.FxColorOrTexture) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxColorOrTexture, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxColorOrTexture(xn, "")
	}
	return
}

func obj_ParamOrRefSid(xn *xmlx.Node, n string) (obj *cdom.ParamOrRefSid) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ParamOrRefSid(xn)
	}
	return
}

func objs_ParamOrRefSid(xn *xmlx.Node, n string) (objs []*cdom.ParamOrRefSid) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ParamOrRefSid, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamOrRefSid(xn, "")
	}
	return
}

func obj_PxModelInst(xn *xmlx.Node, n string) (obj *cdom.PxModelInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxModelInst(xn)
	}
	return
}

func objs_PxModelInst(xn *xmlx.Node, n string) (objs []*cdom.PxModelInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxModelInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxModelInst(xn, "")
	}
	return
}

func obj_Extra(xn *xmlx.Node, n string) (obj *cdom.Extra) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Extra(xn)
	}
	return
}

func objs_Extra(xn *xmlx.Node, n string) (objs []*cdom.Extra) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Extra, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Extra(xn, "")
	}
	return
}

func obj_GeometryControlVertices(xn *xmlx.Node, n string) (obj *cdom.GeometryControlVertices) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryControlVertices(xn)
	}
	return
}

func objs_GeometryControlVertices(xn *xmlx.Node, n string) (objs []*cdom.GeometryControlVertices) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryControlVertices, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryControlVertices(xn, "")
	}
	return
}

func obj_SidBool(xn *xmlx.Node, n string) (obj *cdom.SidBool) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_SidBool(xn)
	}
	return
}

func objs_SidBool(xn *xmlx.Node, n string) (objs []*cdom.SidBool) {
	xns := xcns(xn, n)
	objs = make([]*cdom.SidBool, len(xns))
	for i, xn := range xns {
		objs[i] = obj_SidBool(xn, "")
	}
	return
}

func obj_FxCubeFace(xn *xmlx.Node, n string) (obj *cdom.FxCubeFace) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxCubeFace(xn)
	}
	return
}

func objs_FxCubeFace(xn *xmlx.Node, n string) (objs []*cdom.FxCubeFace) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxCubeFace, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCubeFace(xn, "")
	}
	return
}

func obj_Source(xn *xmlx.Node, n string) (obj *cdom.Source) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Source(xn)
	}
	return
}

func objs_Source(xn *xmlx.Node, n string) (objs []*cdom.Source) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Source, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Source(xn, "")
	}
	return
}

func obj_KxJointKind(xn *xmlx.Node, n string) (obj *cdom.KxJointKind) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxJointKind(xn)
	}
	return
}

func objs_KxJointKind(xn *xmlx.Node, n string) (objs []*cdom.KxJointKind) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxJointKind, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxJointKind(xn, "")
	}
	return
}

func obj_KxAttachment(xn *xmlx.Node, n string) (obj *cdom.KxAttachment) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxAttachment(xn)
	}
	return
}

func objs_KxAttachment(xn *xmlx.Node, n string) (objs []*cdom.KxAttachment) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxAttachment, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxAttachment(xn, "")
	}
	return
}

func obj_PxShape(xn *xmlx.Node, n string) (obj *cdom.PxShape) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_PxShape(xn)
	}
	return
}

func objs_PxShape(xn *xmlx.Node, n string) (objs []*cdom.PxShape) {
	xns := xcns(xn, n)
	objs = make([]*cdom.PxShape, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxShape(xn, "")
	}
	return
}

func obj_GeometryBrepHyperbola(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepHyperbola) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepHyperbola(xn)
	}
	return
}

func objs_GeometryBrepHyperbola(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepHyperbola) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepHyperbola, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepHyperbola(xn, "")
	}
	return
}

func obj_FxProfileCommon(xn *xmlx.Node, n string) (obj *cdom.FxProfileCommon) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxProfileCommon(xn)
	}
	return
}

func objs_FxProfileCommon(xn *xmlx.Node, n string) (objs []*cdom.FxProfileCommon) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxProfileCommon, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxProfileCommon(xn, "")
	}
	return
}

func obj_GeometryBrepNurbsSurface(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepNurbsSurface) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepNurbsSurface(xn)
	}
	return
}

func objs_GeometryBrepNurbsSurface(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepNurbsSurface) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepNurbsSurface, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepNurbsSurface(xn, "")
	}
	return
}

func obj_ParamOrUint(xn *xmlx.Node, n string) (obj *cdom.ParamOrUint) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_ParamOrUint(xn)
	}
	return
}

func objs_ParamOrUint(xn *xmlx.Node, n string) (objs []*cdom.ParamOrUint) {
	xns := xcns(xn, n)
	objs = make([]*cdom.ParamOrUint, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamOrUint(xn, "")
	}
	return
}

func obj_FxPassProgram(xn *xmlx.Node, n string) (obj *cdom.FxPassProgram) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxPassProgram(xn)
	}
	return
}

func objs_FxPassProgram(xn *xmlx.Node, n string) (objs []*cdom.FxPassProgram) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxPassProgram, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassProgram(xn, "")
	}
	return
}

func obj_FxPassProgramBindUniform(xn *xmlx.Node, n string) (obj *cdom.FxPassProgramBindUniform) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxPassProgramBindUniform(xn)
	}
	return
}

func objs_FxPassProgramBindUniform(xn *xmlx.Node, n string) (objs []*cdom.FxPassProgramBindUniform) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxPassProgramBindUniform, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassProgramBindUniform(xn, "")
	}
	return
}

func obj_GeometryPositioning(xn *xmlx.Node, n string) (obj *cdom.GeometryPositioning) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryPositioning(xn)
	}
	return
}

func objs_GeometryPositioning(xn *xmlx.Node, n string) (objs []*cdom.GeometryPositioning) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryPositioning, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryPositioning(xn, "")
	}
	return
}

func obj_GeometryBrepSolids(xn *xmlx.Node, n string) (obj *cdom.GeometryBrepSolids) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_GeometryBrepSolids(xn)
	}
	return
}

func objs_GeometryBrepSolids(xn *xmlx.Node, n string) (objs []*cdom.GeometryBrepSolids) {
	xns := xcns(xn, n)
	objs = make([]*cdom.GeometryBrepSolids, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepSolids(xn, "")
	}
	return
}

func obj_FxPassEvaluationClearDepth(xn *xmlx.Node, n string) (obj *cdom.FxPassEvaluationClearDepth) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxPassEvaluationClearDepth(xn)
	}
	return
}

func objs_FxPassEvaluationClearDepth(xn *xmlx.Node, n string) (objs []*cdom.FxPassEvaluationClearDepth) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxPassEvaluationClearDepth, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassEvaluationClearDepth(xn, "")
	}
	return
}

func obj_SidFloat(xn *xmlx.Node, n string) (obj *cdom.SidFloat) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_SidFloat(xn)
	}
	return
}

func objs_SidFloat(xn *xmlx.Node, n string) (objs []*cdom.SidFloat) {
	xns := xcns(xn, n)
	objs = make([]*cdom.SidFloat, len(xns))
	for i, xn := range xns {
		objs[i] = obj_SidFloat(xn, "")
	}
	return
}

func obj_KxModelInst(xn *xmlx.Node, n string) (obj *cdom.KxModelInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxModelInst(xn)
	}
	return
}

func objs_KxModelInst(xn *xmlx.Node, n string) (objs []*cdom.KxModelInst) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxModelInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxModelInst(xn, "")
	}
	return
}

func obj_Int4(xn *xmlx.Node, n string) (obj *cdom.Int4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_Int4(xn)
	}
	return
}

func objs_Int4(xn *xmlx.Node, n string) (objs []*cdom.Int4) {
	xns := xcns(xn, n)
	objs = make([]*cdom.Int4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Int4(xn, "")
	}
	return
}

func obj_KxAxisIndex(xn *xmlx.Node, n string) (obj *cdom.KxAxisIndex) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_KxAxisIndex(xn)
	}
	return
}

func objs_KxAxisIndex(xn *xmlx.Node, n string) (objs []*cdom.KxAxisIndex) {
	xns := xcns(xn, n)
	objs = make([]*cdom.KxAxisIndex, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxAxisIndex(xn, "")
	}
	return
}

func obj_FxTechniqueCommon(xn *xmlx.Node, n string) (obj *cdom.FxTechniqueCommon) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxTechniqueCommon(xn)
	}
	return
}

func objs_FxTechniqueCommon(xn *xmlx.Node, n string) (objs []*cdom.FxTechniqueCommon) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxTechniqueCommon, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTechniqueCommon(xn, "")
	}
	return
}

func obj_AnimationChannel(xn *xmlx.Node, n string) (obj *cdom.AnimationChannel) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_AnimationChannel(xn)
	}
	return
}

func objs_AnimationChannel(xn *xmlx.Node, n string) (objs []*cdom.AnimationChannel) {
	xns := xcns(xn, n)
	objs = make([]*cdom.AnimationChannel, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AnimationChannel(xn, "")
	}
	return
}

func obj_FxShaderStage(xn *xmlx.Node, n string) (obj *cdom.FxShaderStage) {
	if (xn != nil) && (len(n) > 0) {
		xn = xcn(xn, n)
	}
	if xn != nil {
		obj = init_FxShaderStage(xn)
	}
	return
}

func objs_FxShaderStage(xn *xmlx.Node, n string) (objs []*cdom.FxShaderStage) {
	xns := xcns(xn, n)
	objs = make([]*cdom.FxShaderStage, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxShaderStage(xn, "")
	}
	return
}
