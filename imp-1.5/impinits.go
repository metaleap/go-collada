package collimp

import (
	xmlx "github.com/jteeuwen/go-pkg-xmlx"

	cdom "github.com/metaleap/go-collada/dom"
)

func init_GeometryBrepOrientation(xn *xmlx.Node) (obj *cdom.GeometryBrepOrientation) {
	obj = new(cdom.GeometryBrepOrientation)

	load_GeometryBrepOrientation(xn, obj)
	return
}

func init_FxMaterialInst(xn *xmlx.Node) (obj *cdom.FxMaterialInst) {
	obj = new(cdom.FxMaterialInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_FxMaterialInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxFormatPrecision(xn *xmlx.Node) (obj *cdom.FxFormatPrecision) {
	obj = new(cdom.FxFormatPrecision)

	load_FxFormatPrecision(xn, obj)
	return
}

func init_GeometryPolygonHole(xn *xmlx.Node) (obj *cdom.GeometryPolygonHole) {
	obj = new(cdom.GeometryPolygonHole)

	load_GeometryPolygonHole(xn, obj)
	return
}

func init_KxJointDef(xn *xmlx.Node) (obj *cdom.KxJointDef) {
	obj = new(cdom.KxJointDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_KxJointDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxFilterKind(xn *xmlx.Node) (obj *cdom.FxFilterKind) {
	obj = new(cdom.FxFilterKind)

	load_FxFilterKind(xn, obj)
	return
}

func init_GeometryBrepCircle(xn *xmlx.Node) (obj *cdom.GeometryBrepCircle) {
	obj = new(cdom.GeometryBrepCircle)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepCircle(xn, obj)
	return
}

func init_ControllerInputs(xn *xmlx.Node) (obj *cdom.ControllerInputs) {
	obj = new(cdom.ControllerInputs)
	has_Extras(xn, &obj.HasExtras)
	has_Inputs(xn, &obj.HasInputs)

	load_ControllerInputs(xn, obj)
	return
}

func init_FxInitFrom(xn *xmlx.Node) (obj *cdom.FxInitFrom) {
	obj = new(cdom.FxInitFrom)

	load_FxInitFrom(xn, obj)
	return
}

func init_Asset(xn *xmlx.Node) (obj *cdom.Asset) {
	obj = cdom.NewAsset()
	has_Extras(xn, &obj.HasExtras)

	load_Asset(xn, obj)
	return
}

func init_LightPoint(xn *xmlx.Node) (obj *cdom.LightPoint) {
	obj = cdom.NewLightPoint()

	load_LightPoint(xn, obj)
	return
}

func init_GeometryBrepSphere(xn *xmlx.Node) (obj *cdom.GeometryBrepSphere) {
	obj = new(cdom.GeometryBrepSphere)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepSphere(xn, obj)
	return
}

func init_KxMotionSystem(xn *xmlx.Node) (obj *cdom.KxMotionSystem) {
	obj = new(cdom.KxMotionSystem)
	has_Techniques(xn, &obj.HasTechniques)

	load_KxMotionSystem(xn, obj)
	return
}

func init_KxEffector(xn *xmlx.Node) (obj *cdom.KxEffector) {
	obj = cdom.NewKxEffector()
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_KxEffector(xn, obj)
	return
}

func init_CameraOptics(xn *xmlx.Node) (obj *cdom.CameraOptics) {
	obj = new(cdom.CameraOptics)
	has_Extras(xn, &obj.HasExtras)
	has_Techniques(xn, &obj.HasTechniques)

	load_CameraOptics(xn, obj)
	return
}

func init_KxJointAxisBinding(xn *xmlx.Node) (obj *cdom.KxJointAxisBinding) {
	obj = new(cdom.KxJointAxisBinding)

	load_KxJointAxisBinding(xn, obj)
	return
}

func init_KxMotionAxis(xn *xmlx.Node) (obj *cdom.KxMotionAxis) {
	obj = cdom.NewKxMotionAxis()
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_KxMotionAxis(xn, obj)
	return
}

func init_FxTextureOpaque(xn *xmlx.Node) (obj *cdom.FxTextureOpaque) {
	obj = new(cdom.FxTextureOpaque)

	load_FxTextureOpaque(xn, obj)
	return
}

func init_Float7(xn *xmlx.Node) (obj *cdom.Float7) {
	obj = new(cdom.Float7)

	load_Float7(xn, obj)
	return
}

func init_FxCreateMips(xn *xmlx.Node) (obj *cdom.FxCreateMips) {
	obj = new(cdom.FxCreateMips)

	load_FxCreateMips(xn, obj)
	return
}

func init_KxLink(xn *xmlx.Node) (obj *cdom.KxLink) {
	obj = new(cdom.KxLink)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_KxLink(xn, obj)
	return
}

func init_FxEffectInst(xn *xmlx.Node) (obj *cdom.FxEffectInst) {
	obj = new(cdom.FxEffectInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_FxEffectInst(xn, obj)
	obj.SetDirty()
	return
}

func init_MaterialBinding(xn *xmlx.Node) (obj *cdom.MaterialBinding) {
	obj = new(cdom.MaterialBinding)
	has_Extras(xn, &obj.HasExtras)
	has_Techniques(xn, &obj.HasTechniques)

	load_MaterialBinding(xn, obj)
	return
}

func init_PxRigidConstraintDef(xn *xmlx.Node) (obj *cdom.PxRigidConstraintDef) {
	obj = new(cdom.PxRigidConstraintDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)
	has_Techniques(xn, &obj.HasTechniques)

	load_PxRigidConstraintDef(xn, obj)
	obj.SetDirty()
	return
}

func init_VisualSceneRendering(xn *xmlx.Node) (obj *cdom.VisualSceneRendering) {
	obj = cdom.NewVisualSceneRendering()
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_VisualSceneRendering(xn, obj)
	return
}

func init_InputShared(xn *xmlx.Node) (obj *cdom.InputShared) {
	obj = new(cdom.InputShared)

	load_InputShared(xn, obj)
	return
}

func init_FxVertexInputBinding(xn *xmlx.Node) (obj *cdom.FxVertexInputBinding) {
	obj = new(cdom.FxVertexInputBinding)

	load_FxVertexInputBinding(xn, obj)
	return
}

func init_Float2x2(xn *xmlx.Node) (obj *cdom.Float2x2) {
	obj = new(cdom.Float2x2)

	load_Float2x2(xn, obj)
	return
}

func init_CameraOrthographic(xn *xmlx.Node) (obj *cdom.CameraOrthographic) {
	obj = new(cdom.CameraOrthographic)

	load_CameraOrthographic(xn, obj)
	return
}

func init_PxRigidBodyInst(xn *xmlx.Node) (obj *cdom.PxRigidBodyInst) {
	obj = new(cdom.PxRigidBodyInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)
	has_Techniques(xn, &obj.HasTechniques)

	load_PxRigidBodyInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxCreateCubeInitFrom(xn *xmlx.Node) (obj *cdom.FxCreateCubeInitFrom) {
	obj = new(cdom.FxCreateCubeInitFrom)

	load_FxCreateCubeInitFrom(xn, obj)
	return
}

func init_GeometryMesh(xn *xmlx.Node) (obj *cdom.GeometryMesh) {
	obj = cdom.NewGeometryMesh()
	has_Extras(xn, &obj.HasExtras)
	has_Sources(xn, &obj.HasSources)

	load_GeometryMesh(xn, obj)
	return
}

func init_GeometryBrepCylinder(xn *xmlx.Node) (obj *cdom.GeometryBrepCylinder) {
	obj = new(cdom.GeometryBrepCylinder)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepCylinder(xn, obj)
	return
}

func init_FxPassEvaluationClearStencil(xn *xmlx.Node) (obj *cdom.FxPassEvaluationClearStencil) {
	obj = new(cdom.FxPassEvaluationClearStencil)

	load_FxPassEvaluationClearStencil(xn, obj)
	return
}

func init_FxCreateFormat(xn *xmlx.Node) (obj *cdom.FxCreateFormat) {
	obj = new(cdom.FxCreateFormat)

	load_FxCreateFormat(xn, obj)
	return
}

func init_KxModelBinding(xn *xmlx.Node) (obj *cdom.KxModelBinding) {
	obj = new(cdom.KxModelBinding)

	load_KxModelBinding(xn, obj)
	return
}

func init_FxPassProgramShader(xn *xmlx.Node) (obj *cdom.FxPassProgramShader) {
	obj = new(cdom.FxPassProgramShader)
	has_Extras(xn, &obj.HasExtras)

	load_FxPassProgramShader(xn, obj)
	return
}

func init_Float4x4(xn *xmlx.Node) (obj *cdom.Float4x4) {
	obj = new(cdom.Float4x4)

	load_Float4x4(xn, obj)
	return
}

func init_Param(xn *xmlx.Node) (obj *cdom.Param) {
	obj = new(cdom.Param)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_Param(xn, obj)
	return
}

func init_ParamOrInt(xn *xmlx.Node) (obj *cdom.ParamOrInt) {
	obj = new(cdom.ParamOrInt)

	load_ParamOrInt(xn, obj)
	return
}

func init_AssetGeographicLocation(xn *xmlx.Node) (obj *cdom.AssetGeographicLocation) {
	obj = new(cdom.AssetGeographicLocation)

	load_AssetGeographicLocation(xn, obj)
	return
}

func init_ControllerSkin(xn *xmlx.Node) (obj *cdom.ControllerSkin) {
	obj = cdom.NewControllerSkin()
	has_Sources(xn, &obj.HasSources)

	load_ControllerSkin(xn, obj)
	return
}

func init_AssetContributor(xn *xmlx.Node) (obj *cdom.AssetContributor) {
	obj = new(cdom.AssetContributor)

	load_AssetContributor(xn, obj)
	return
}

func init_VisualSceneInst(xn *xmlx.Node) (obj *cdom.VisualSceneInst) {
	obj = new(cdom.VisualSceneInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_VisualSceneInst(xn, obj)
	obj.SetDirty()
	return
}

func init_KxFrame(xn *xmlx.Node) (obj *cdom.KxFrame) {
	obj = new(cdom.KxFrame)

	load_KxFrame(xn, obj)
	return
}

func init_PxRigidConstraintDefs(xn *xmlx.Node) (obj *cdom.PxRigidConstraintDefs) {
	obj = new(cdom.PxRigidConstraintDefs)

	load_PxRigidConstraintDefs(xn, obj)
	return
}

func init_PxSceneInst(xn *xmlx.Node) (obj *cdom.PxSceneInst) {
	obj = new(cdom.PxSceneInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_PxSceneInst(xn, obj)
	obj.SetDirty()
	return
}

func init_Float2x3(xn *xmlx.Node) (obj *cdom.Float2x3) {
	obj = new(cdom.Float2x3)

	load_Float2x3(xn, obj)
	return
}

func init_GeometryBrepWires(xn *xmlx.Node) (obj *cdom.GeometryBrepWires) {
	obj = new(cdom.GeometryBrepWires)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryBrepWires(xn, obj)
	return
}

func init_LightAmbient(xn *xmlx.Node) (obj *cdom.LightAmbient) {
	obj = new(cdom.LightAmbient)

	load_LightAmbient(xn, obj)
	return
}

func init_GeometryBrepNurbs(xn *xmlx.Node) (obj *cdom.GeometryBrepNurbs) {
	obj = cdom.NewGeometryBrepNurbs()
	has_Extras(xn, &obj.HasExtras)
	has_Sources(xn, &obj.HasSources)

	load_GeometryBrepNurbs(xn, obj)
	return
}

func init_SourceArray(xn *xmlx.Node) (obj *cdom.SourceArray) {
	obj = new(cdom.SourceArray)
	has_Name(xn, &obj.HasName)

	load_SourceArray(xn, obj)
	return
}

func init_CameraInst(xn *xmlx.Node) (obj *cdom.CameraInst) {
	obj = new(cdom.CameraInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_CameraInst(xn, obj)
	obj.SetDirty()
	return
}

func init_PxCylinder(xn *xmlx.Node) (obj *cdom.PxCylinder) {
	obj = new(cdom.PxCylinder)
	has_Extras(xn, &obj.HasExtras)

	load_PxCylinder(xn, obj)
	return
}

func init_KxFrameTcp(xn *xmlx.Node) (obj *cdom.KxFrameTcp) {
	obj = new(cdom.KxFrameTcp)

	load_KxFrameTcp(xn, obj)
	return
}

func init_Int2x2(xn *xmlx.Node) (obj *cdom.Int2x2) {
	obj = new(cdom.Int2x2)

	load_Int2x2(xn, obj)
	return
}

func init_GeometryBrepEdges(xn *xmlx.Node) (obj *cdom.GeometryBrepEdges) {
	obj = new(cdom.GeometryBrepEdges)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryBrepEdges(xn, obj)
	return
}

func init_AnimSamplerBehavior(xn *xmlx.Node) (obj *cdom.AnimSamplerBehavior) {
	obj = new(cdom.AnimSamplerBehavior)

	load_AnimSamplerBehavior(xn, obj)
	return
}

func init_Input(xn *xmlx.Node) (obj *cdom.Input) {
	obj = new(cdom.Input)

	load_Input(xn, obj)
	return
}

func init_FxPass(xn *xmlx.Node) (obj *cdom.FxPass) {
	obj = cdom.NewFxPass()
	has_Extras(xn, &obj.HasExtras)
	has_Sid(xn, &obj.HasSid)

	load_FxPass(xn, obj)
	return
}

func init_GeometryBrepSweptSurface(xn *xmlx.Node) (obj *cdom.GeometryBrepSweptSurface) {
	obj = new(cdom.GeometryBrepSweptSurface)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepSweptSurface(xn, obj)
	return
}

func init_KxFrameTip(xn *xmlx.Node) (obj *cdom.KxFrameTip) {
	obj = new(cdom.KxFrameTip)

	load_KxFrameTip(xn, obj)
	return
}

func init_GeometryBrepPlane(xn *xmlx.Node) (obj *cdom.GeometryBrepPlane) {
	obj = new(cdom.GeometryBrepPlane)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepPlane(xn, obj)
	return
}

func init_FxSamplerKind(xn *xmlx.Node) (obj *cdom.FxSamplerKind) {
	obj = new(cdom.FxSamplerKind)

	load_FxSamplerKind(xn, obj)
	return
}

func init_KxJointInst(xn *xmlx.Node) (obj *cdom.KxJointInst) {
	obj = new(cdom.KxJointInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_KxJointInst(xn, obj)
	obj.SetDirty()
	return
}

func init_GeometryBrepPcurves(xn *xmlx.Node) (obj *cdom.GeometryBrepPcurves) {
	obj = new(cdom.GeometryBrepPcurves)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryBrepPcurves(xn, obj)
	return
}

func init_Technique(xn *xmlx.Node) (obj *cdom.Technique) {
	obj = new(cdom.Technique)

	load_Technique(xn, obj)
	return
}

func init_GeometryBrepSurfaceCurves(xn *xmlx.Node) (obj *cdom.GeometryBrepSurfaceCurves) {
	obj = new(cdom.GeometryBrepSurfaceCurves)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepSurfaceCurves(xn, obj)
	return
}

func init_Float2x4(xn *xmlx.Node) (obj *cdom.Float2x4) {
	obj = new(cdom.Float2x4)

	load_Float2x4(xn, obj)
	return
}

func init_FxCreate3D(xn *xmlx.Node) (obj *cdom.FxCreate3D) {
	obj = new(cdom.FxCreate3D)

	load_FxCreate3D(xn, obj)
	return
}

func init_SidVec3(xn *xmlx.Node) (obj *cdom.SidVec3) {
	obj = new(cdom.SidVec3)
	has_Sid(xn, &obj.HasSid)

	load_SidVec3(xn, obj)
	return
}

func init_PxRigidConstraintAttachment(xn *xmlx.Node) (obj *cdom.PxRigidConstraintAttachment) {
	obj = new(cdom.PxRigidConstraintAttachment)
	has_Extras(xn, &obj.HasExtras)

	load_PxRigidConstraintAttachment(xn, obj)
	return
}

func init_PxModelDef(xn *xmlx.Node) (obj *cdom.PxModelDef) {
	obj = new(cdom.PxModelDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_PxModelDef(xn, obj)
	obj.SetDirty()
	return
}

func init_Bool3(xn *xmlx.Node) (obj *cdom.Bool3) {
	obj = new(cdom.Bool3)

	load_Bool3(xn, obj)
	return
}

func init_GeometryBrepParabola(xn *xmlx.Node) (obj *cdom.GeometryBrepParabola) {
	obj = new(cdom.GeometryBrepParabola)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepParabola(xn, obj)
	return
}

func init_Document(xn *xmlx.Node) (obj *cdom.Document) {
	obj = new(cdom.Document)
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)

	load_Document(xn, obj)
	return
}

func init_SidString(xn *xmlx.Node) (obj *cdom.SidString) {
	obj = new(cdom.SidString)
	has_Sid(xn, &obj.HasSid)

	load_SidString(xn, obj)
	return
}

func init_LightSpot(xn *xmlx.Node) (obj *cdom.LightSpot) {
	obj = cdom.NewLightSpot()

	load_LightSpot(xn, obj)
	return
}

func init_ParamDef(xn *xmlx.Node) (obj *cdom.ParamDef) {
	obj = new(cdom.ParamDef)
	has_Sid(xn, &obj.HasSid)

	load_ParamDef(xn, obj)
	return
}

func init_PxSceneDef(xn *xmlx.Node) (obj *cdom.PxSceneDef) {
	obj = new(cdom.PxSceneDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_PxSceneDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxImageInst(xn *xmlx.Node) (obj *cdom.FxImageInst) {
	obj = new(cdom.FxImageInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_FxImageInst(xn, obj)
	obj.SetDirty()
	return
}

func init_PxRigidBodyCommon(xn *xmlx.Node) (obj *cdom.PxRigidBodyCommon) {
	obj = new(cdom.PxRigidBodyCommon)

	load_PxRigidBodyCommon(xn, obj)
	return
}

func init_NodeInst(xn *xmlx.Node) (obj *cdom.NodeInst) {
	obj = new(cdom.NodeInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_NodeInst(xn, obj)
	obj.SetDirty()
	return
}

func init_GeometryBrepLine(xn *xmlx.Node) (obj *cdom.GeometryBrepLine) {
	obj = new(cdom.GeometryBrepLine)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepLine(xn, obj)
	return
}

func init_KxAttachmentKind(xn *xmlx.Node) (obj *cdom.KxAttachmentKind) {
	obj = new(cdom.KxAttachmentKind)

	load_KxAttachmentKind(xn, obj)
	return
}

func init_CameraImager(xn *xmlx.Node) (obj *cdom.CameraImager) {
	obj = new(cdom.CameraImager)
	has_Extras(xn, &obj.HasExtras)
	has_Techniques(xn, &obj.HasTechniques)

	load_CameraImager(xn, obj)
	return
}

func init_FxTechniqueGlsl(xn *xmlx.Node) (obj *cdom.FxTechniqueGlsl) {
	obj = new(cdom.FxTechniqueGlsl)
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Sid(xn, &obj.HasSid)

	load_FxTechniqueGlsl(xn, obj)
	return
}

func init_LightInst(xn *xmlx.Node) (obj *cdom.LightInst) {
	obj = new(cdom.LightInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_LightInst(xn, obj)
	obj.SetDirty()
	return
}

func init_Formula(xn *xmlx.Node) (obj *cdom.Formula) {
	obj = new(cdom.Formula)

	load_Formula(xn, obj)
	return
}

func init_LightAttenuation(xn *xmlx.Node) (obj *cdom.LightAttenuation) {
	obj = cdom.NewLightAttenuation()

	load_LightAttenuation(xn, obj)
	return
}

func init_Bool4(xn *xmlx.Node) (obj *cdom.Bool4) {
	obj = new(cdom.Bool4)

	load_Bool4(xn, obj)
	return
}

func init_FxImageDef(xn *xmlx.Node) (obj *cdom.FxImageDef) {
	obj = new(cdom.FxImageDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_FxImageDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FormulaDef(xn *xmlx.Node) (obj *cdom.FormulaDef) {
	obj = new(cdom.FormulaDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_Sid(xn, &obj.HasSid)
	has_Techniques(xn, &obj.HasTechniques)

	load_FormulaDef(xn, obj)
	obj.SetDirty()
	return
}

func init_KxArticulatedSystemDef(xn *xmlx.Node) (obj *cdom.KxArticulatedSystemDef) {
	obj = new(cdom.KxArticulatedSystemDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_KxArticulatedSystemDef(xn, obj)
	obj.SetDirty()
	return
}

func init_PxForceFieldDef(xn *xmlx.Node) (obj *cdom.PxForceFieldDef) {
	obj = new(cdom.PxForceFieldDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_PxForceFieldDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxSamplerFiltering(xn *xmlx.Node) (obj *cdom.FxSamplerFiltering) {
	obj = new(cdom.FxSamplerFiltering)

	load_FxSamplerFiltering(xn, obj)
	return
}

func init_NodeDef(xn *xmlx.Node) (obj *cdom.NodeDef) {
	obj = new(cdom.NodeDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_NodeDef(xn, obj)
	obj.SetDirty()
	return
}

func init_CameraDef(xn *xmlx.Node) (obj *cdom.CameraDef) {
	obj = new(cdom.CameraDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_CameraDef(xn, obj)
	obj.SetDirty()
	return
}

func init_Float4x3(xn *xmlx.Node) (obj *cdom.Float4x3) {
	obj = new(cdom.Float4x3)

	load_Float4x3(xn, obj)
	return
}

func init_ParamOrSidFloat(xn *xmlx.Node) (obj *cdom.ParamOrSidFloat) {
	obj = new(cdom.ParamOrSidFloat)

	load_ParamOrSidFloat(xn, obj)
	return
}

func init_FxProfileGlsl(xn *xmlx.Node) (obj *cdom.FxProfileGlsl) {
	obj = cdom.NewFxProfileGlsl()

	load_FxProfileGlsl(xn, obj)
	return
}

func init_Int3(xn *xmlx.Node) (obj *cdom.Int3) {
	obj = new(cdom.Int3)

	load_Int3(xn, obj)
	return
}

func init_FormulaInst(xn *xmlx.Node) (obj *cdom.FormulaInst) {
	obj = new(cdom.FormulaInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_FormulaInst(xn, obj)
	obj.SetDirty()
	return
}

func init_ControllerDef(xn *xmlx.Node) (obj *cdom.ControllerDef) {
	obj = new(cdom.ControllerDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_ControllerDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxEffectDef(xn *xmlx.Node) (obj *cdom.FxEffectDef) {
	obj = new(cdom.FxEffectDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_FxParamDefs(xn, &obj.HasFxParamDefs)
	has_Name(xn, &obj.HasName)

	load_FxEffectDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxMaterialDef(xn *xmlx.Node) (obj *cdom.FxMaterialDef) {
	obj = new(cdom.FxMaterialDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_FxMaterialDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxGlslTechniques(xn *xmlx.Node) (obj *cdom.FxGlslTechniques) {
	obj = new(cdom.FxGlslTechniques)

	load_FxGlslTechniques(xn, obj)
	return
}

func init_Float2(xn *xmlx.Node) (obj *cdom.Float2) {
	obj = new(cdom.Float2)

	load_Float2(xn, obj)
	return
}

func init_FxCreateInitFrom(xn *xmlx.Node) (obj *cdom.FxCreateInitFrom) {
	obj = new(cdom.FxCreateInitFrom)

	load_FxCreateInitFrom(xn, obj)
	return
}

func init_KxJointLimits(xn *xmlx.Node) (obj *cdom.KxJointLimits) {
	obj = new(cdom.KxJointLimits)

	load_KxJointLimits(xn, obj)
	return
}

func init_FxEffectInstTechniqueHint(xn *xmlx.Node) (obj *cdom.FxEffectInstTechniqueHint) {
	obj = new(cdom.FxEffectInstTechniqueHint)

	load_FxEffectInstTechniqueHint(xn, obj)
	return
}

func init_Float3x3(xn *xmlx.Node) (obj *cdom.Float3x3) {
	obj = new(cdom.Float3x3)

	load_Float3x3(xn, obj)
	return
}

func init_ParamDefs(xn *xmlx.Node) (obj *cdom.ParamDefs) {
	obj = new(cdom.ParamDefs)

	load_ParamDefs(xn, obj)
	return
}

func init_FxCreateFormatHint(xn *xmlx.Node) (obj *cdom.FxCreateFormatHint) {
	obj = new(cdom.FxCreateFormatHint)

	load_FxCreateFormatHint(xn, obj)
	return
}

func init_SourceAccessor(xn *xmlx.Node) (obj *cdom.SourceAccessor) {
	obj = cdom.NewSourceAccessor()

	load_SourceAccessor(xn, obj)
	return
}

func init_ParamInsts(xn *xmlx.Node) (obj *cdom.ParamInsts) {
	obj = new(cdom.ParamInsts)

	load_ParamInsts(xn, obj)
	return
}

func init_ControllerMorph(xn *xmlx.Node) (obj *cdom.ControllerMorph) {
	obj = cdom.NewControllerMorph()
	has_Sources(xn, &obj.HasSources)

	load_ControllerMorph(xn, obj)
	return
}

func init_FxImageInitFrom(xn *xmlx.Node) (obj *cdom.FxImageInitFrom) {
	obj = new(cdom.FxImageInitFrom)

	load_FxImageInitFrom(xn, obj)
	return
}

func init_GeometryVertices(xn *xmlx.Node) (obj *cdom.GeometryVertices) {
	obj = new(cdom.GeometryVertices)
	has_Extras(xn, &obj.HasExtras)
	has_Inputs(xn, &obj.HasInputs)
	has_Name(xn, &obj.HasName)

	load_GeometryVertices(xn, obj)
	return
}

func init_GeometryBrepTorus(xn *xmlx.Node) (obj *cdom.GeometryBrepTorus) {
	obj = new(cdom.GeometryBrepTorus)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepTorus(xn, obj)
	return
}

func init_FxTexture(xn *xmlx.Node) (obj *cdom.FxTexture) {
	obj = new(cdom.FxTexture)
	has_Extras(xn, &obj.HasExtras)

	load_FxTexture(xn, obj)
	return
}

func init_FxPassEvaluation(xn *xmlx.Node) (obj *cdom.FxPassEvaluation) {
	obj = new(cdom.FxPassEvaluation)

	load_FxPassEvaluation(xn, obj)
	return
}

func init_ParamInst(xn *xmlx.Node) (obj *cdom.ParamInst) {
	obj = new(cdom.ParamInst)

	load_ParamInst(xn, obj)
	return
}

func init_FxProfile(xn *xmlx.Node) (obj *cdom.FxProfile) {
	obj = new(cdom.FxProfile)
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_FxParamDefs(xn, &obj.HasFxParamDefs)

	load_FxProfile(xn, obj)
	return
}

func init_FxPassProgramBindAttribute(xn *xmlx.Node) (obj *cdom.FxPassProgramBindAttribute) {
	obj = new(cdom.FxPassProgramBindAttribute)

	load_FxPassProgramBindAttribute(xn, obj)
	return
}

func init_FxCreate2D(xn *xmlx.Node) (obj *cdom.FxCreate2D) {
	obj = new(cdom.FxCreate2D)

	load_FxCreate2D(xn, obj)
	return
}

func init_ParamOrFloat2(xn *xmlx.Node) (obj *cdom.ParamOrFloat2) {
	obj = new(cdom.ParamOrFloat2)

	load_ParamOrFloat2(xn, obj)
	return
}

func init_FxCreateCube(xn *xmlx.Node) (obj *cdom.FxCreateCube) {
	obj = new(cdom.FxCreateCube)

	load_FxCreateCube(xn, obj)
	return
}

func init_KxBinding(xn *xmlx.Node) (obj *cdom.KxBinding) {
	obj = new(cdom.KxBinding)

	load_KxBinding(xn, obj)
	return
}

func init_GeometryBrepBox(xn *xmlx.Node) (obj *cdom.GeometryBrepBox) {
	obj = new(cdom.GeometryBrepBox)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepBox(xn, obj)
	return
}

func init_Float4x2(xn *xmlx.Node) (obj *cdom.Float4x2) {
	obj = new(cdom.Float4x2)

	load_Float4x2(xn, obj)
	return
}

func init_FxCreate3DInitFrom(xn *xmlx.Node) (obj *cdom.FxCreate3DInitFrom) {
	obj = new(cdom.FxCreate3DInitFrom)

	load_FxCreate3DInitFrom(xn, obj)
	return
}

func init_FxCreate2DSizeExact(xn *xmlx.Node) (obj *cdom.FxCreate2DSizeExact) {
	obj = new(cdom.FxCreate2DSizeExact)

	load_FxCreate2DSizeExact(xn, obj)
	return
}

func init_PxMaterialDef(xn *xmlx.Node) (obj *cdom.PxMaterialDef) {
	obj = new(cdom.PxMaterialDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_PxMaterialDef(xn, obj)
	obj.SetDirty()
	return
}

func init_AnimationDef(xn *xmlx.Node) (obj *cdom.AnimationDef) {
	obj = new(cdom.AnimationDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sources(xn, &obj.HasSources)

	load_AnimationDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxParamDef(xn *xmlx.Node) (obj *cdom.FxParamDef) {
	obj = new(cdom.FxParamDef)
	has_Sid(xn, &obj.HasSid)

	load_FxParamDef(xn, obj)
	return
}

func init_ChildNode(xn *xmlx.Node) (obj *cdom.ChildNode) {
	obj = new(cdom.ChildNode)

	load_ChildNode(xn, obj)
	return
}

func init_PxRigidBodyDefs(xn *xmlx.Node) (obj *cdom.PxRigidBodyDefs) {
	obj = new(cdom.PxRigidBodyDefs)

	load_PxRigidBodyDefs(xn, obj)
	return
}

func init_FxSamplerImage(xn *xmlx.Node) (obj *cdom.FxSamplerImage) {
	obj = new(cdom.FxSamplerImage)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_FxSamplerImage(xn, obj)
	return
}

func init_FxPassEvaluationClearColor(xn *xmlx.Node) (obj *cdom.FxPassEvaluationClearColor) {
	obj = new(cdom.FxPassEvaluationClearColor)

	load_FxPassEvaluationClearColor(xn, obj)
	return
}

func init_Layers(xn *xmlx.Node) (obj *cdom.Layers) {
	obj = new(cdom.Layers)

	load_Layers(xn, obj)
	return
}

func init_GeometryBrepFaces(xn *xmlx.Node) (obj *cdom.GeometryBrepFaces) {
	obj = new(cdom.GeometryBrepFaces)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryBrepFaces(xn, obj)
	return
}

func init_FxCreate2DSizeRatio(xn *xmlx.Node) (obj *cdom.FxCreate2DSizeRatio) {
	obj = new(cdom.FxCreate2DSizeRatio)

	load_FxCreate2DSizeRatio(xn, obj)
	return
}

func init_FxColor(xn *xmlx.Node) (obj *cdom.FxColor) {
	obj = new(cdom.FxColor)
	has_Sid(xn, &obj.HasSid)

	load_FxColor(xn, obj)
	return
}

func init_FxParamDefs(xn *xmlx.Node) (obj *cdom.FxParamDefs) {
	obj = new(cdom.FxParamDefs)

	load_FxParamDefs(xn, obj)
	return
}

func init_GeometryBrepCone(xn *xmlx.Node) (obj *cdom.GeometryBrepCone) {
	obj = new(cdom.GeometryBrepCone)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepCone(xn, obj)
	return
}

func init_LightDirectional(xn *xmlx.Node) (obj *cdom.LightDirectional) {
	obj = new(cdom.LightDirectional)

	load_LightDirectional(xn, obj)
	return
}

func init_KxSceneInst(xn *xmlx.Node) (obj *cdom.KxSceneInst) {
	obj = new(cdom.KxSceneInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_KxSceneInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxSampler(xn *xmlx.Node) (obj *cdom.FxSampler) {
	obj = cdom.NewFxSampler()
	has_Extras(xn, &obj.HasExtras)

	load_FxSampler(xn, obj)
	return
}

func init_GeometryBrep(xn *xmlx.Node) (obj *cdom.GeometryBrep) {
	obj = cdom.NewGeometryBrep()
	has_Extras(xn, &obj.HasExtras)
	has_Sources(xn, &obj.HasSources)

	load_GeometryBrep(xn, obj)
	return
}

func init_FxTechnique(xn *xmlx.Node) (obj *cdom.FxTechnique) {
	obj = new(cdom.FxTechnique)
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Sid(xn, &obj.HasSid)

	load_FxTechnique(xn, obj)
	return
}

func init_FxFormatRange(xn *xmlx.Node) (obj *cdom.FxFormatRange) {
	obj = new(cdom.FxFormatRange)

	load_FxFormatRange(xn, obj)
	return
}

func init_Int4x4(xn *xmlx.Node) (obj *cdom.Int4x4) {
	obj = new(cdom.Int4x4)

	load_Int4x4(xn, obj)
	return
}

func init_Sources(xn *xmlx.Node) (obj *cdom.Sources) {
	obj = new(cdom.Sources)

	load_Sources(xn, obj)
	return
}

func init_FxTechniqueKind(xn *xmlx.Node) (obj *cdom.FxTechniqueKind) {
	obj = new(cdom.FxTechniqueKind)

	load_FxTechniqueKind(xn, obj)
	return
}

func init_AnimationSampler(xn *xmlx.Node) (obj *cdom.AnimationSampler) {
	obj = new(cdom.AnimationSampler)
	has_Inputs(xn, &obj.HasInputs)

	load_AnimationSampler(xn, obj)
	return
}

func init_FxCreate(xn *xmlx.Node) (obj *cdom.FxCreate) {
	obj = new(cdom.FxCreate)

	load_FxCreate(xn, obj)
	return
}

func init_ParamOrBool(xn *xmlx.Node) (obj *cdom.ParamOrBool) {
	obj = new(cdom.ParamOrBool)

	load_ParamOrBool(xn, obj)
	return
}

func init_KxFrameObject(xn *xmlx.Node) (obj *cdom.KxFrameObject) {
	obj = new(cdom.KxFrameObject)

	load_KxFrameObject(xn, obj)
	return
}

func init_ParamOrFloat(xn *xmlx.Node) (obj *cdom.ParamOrFloat) {
	obj = new(cdom.ParamOrFloat)

	load_ParamOrFloat(xn, obj)
	return
}

func init_ControllerInst(xn *xmlx.Node) (obj *cdom.ControllerInst) {
	obj = new(cdom.ControllerInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_ControllerInst(xn, obj)
	obj.SetDirty()
	return
}

func init_Float3x4(xn *xmlx.Node) (obj *cdom.Float3x4) {
	obj = new(cdom.Float3x4)

	load_Float3x4(xn, obj)
	return
}

func init_Scene(xn *xmlx.Node) (obj *cdom.Scene) {
	obj = new(cdom.Scene)
	has_Extras(xn, &obj.HasExtras)

	load_Scene(xn, obj)
	return
}

func init_GeometryInst(xn *xmlx.Node) (obj *cdom.GeometryInst) {
	obj = new(cdom.GeometryInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_GeometryInst(xn, obj)
	obj.SetDirty()
	return
}

func init_IndexedInputs(xn *xmlx.Node) (obj *cdom.IndexedInputs) {
	obj = new(cdom.IndexedInputs)

	load_IndexedInputs(xn, obj)
	return
}

func init_PxRigidConstraintSpring(xn *xmlx.Node) (obj *cdom.PxRigidConstraintSpring) {
	obj = cdom.NewPxRigidConstraintSpring()

	load_PxRigidConstraintSpring(xn, obj)
	return
}

func init_GeometryBrepCapsule(xn *xmlx.Node) (obj *cdom.GeometryBrepCapsule) {
	obj = new(cdom.GeometryBrepCapsule)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepCapsule(xn, obj)
	return
}

func init_TransformKind(xn *xmlx.Node) (obj *cdom.TransformKind) {
	obj = new(cdom.TransformKind)

	load_TransformKind(xn, obj)
	return
}

func init_PxMaterial(xn *xmlx.Node) (obj *cdom.PxMaterial) {
	obj = new(cdom.PxMaterial)

	load_PxMaterial(xn, obj)
	return
}

func init_GeometryPrimitives(xn *xmlx.Node) (obj *cdom.GeometryPrimitives) {
	obj = new(cdom.GeometryPrimitives)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryPrimitives(xn, obj)
	return
}

func init_Transform(xn *xmlx.Node) (obj *cdom.Transform) {
	obj = new(cdom.Transform)
	has_Sid(xn, &obj.HasSid)

	load_Transform(xn, obj)
	return
}

func init_PxRigidBodyDef(xn *xmlx.Node) (obj *cdom.PxRigidBodyDef) {
	obj = new(cdom.PxRigidBodyDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)
	has_Techniques(xn, &obj.HasTechniques)

	load_PxRigidBodyDef(xn, obj)
	obj.SetDirty()
	return
}

func init_KxArticulatedSystemInst(xn *xmlx.Node) (obj *cdom.KxArticulatedSystemInst) {
	obj = new(cdom.KxArticulatedSystemInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_KxArticulatedSystemInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxProfileGlslCodeInclude(xn *xmlx.Node) (obj *cdom.FxProfileGlslCodeInclude) {
	obj = new(cdom.FxProfileGlslCodeInclude)
	has_Sid(xn, &obj.HasSid)

	load_FxProfileGlslCodeInclude(xn, obj)
	return
}

func init_GeometryDef(xn *xmlx.Node) (obj *cdom.GeometryDef) {
	obj = new(cdom.GeometryDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxPassProgramShaderSources(xn *xmlx.Node) (obj *cdom.FxPassProgramShaderSources) {
	obj = new(cdom.FxPassProgramShaderSources)

	load_FxPassProgramShaderSources(xn, obj)
	return
}

func init_VisualSceneEvaluation(xn *xmlx.Node) (obj *cdom.VisualSceneEvaluation) {
	obj = new(cdom.VisualSceneEvaluation)
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_VisualSceneEvaluation(xn, obj)
	obj.SetDirty()
	return
}

func init_KxSceneDef(xn *xmlx.Node) (obj *cdom.KxSceneDef) {
	obj = new(cdom.KxSceneDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_KxSceneDef(xn, obj)
	obj.SetDirty()
	return
}

func init_Float3(xn *xmlx.Node) (obj *cdom.Float3) {
	obj = new(cdom.Float3)

	load_Float3(xn, obj)
	return
}

func init_Int2(xn *xmlx.Node) (obj *cdom.Int2) {
	obj = new(cdom.Int2)

	load_Int2(xn, obj)
	return
}

func init_KxJoint(xn *xmlx.Node) (obj *cdom.KxJoint) {
	obj = new(cdom.KxJoint)
	has_Sid(xn, &obj.HasSid)

	load_KxJoint(xn, obj)
	return
}

func init_SidFloat3(xn *xmlx.Node) (obj *cdom.SidFloat3) {
	obj = new(cdom.SidFloat3)
	has_Sid(xn, &obj.HasSid)

	load_SidFloat3(xn, obj)
	return
}

func init_GeometryBrepCurve(xn *xmlx.Node) (obj *cdom.GeometryBrepCurve) {
	obj = new(cdom.GeometryBrepCurve)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_GeometryBrepCurve(xn, obj)
	return
}

func init_VisualSceneDef(xn *xmlx.Node) (obj *cdom.VisualSceneDef) {
	obj = new(cdom.VisualSceneDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_VisualSceneDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxPassState(xn *xmlx.Node) (obj *cdom.FxPassState) {
	obj = new(cdom.FxPassState)

	load_FxPassState(xn, obj)
	return
}

func init_GeometryPrimitiveKind(xn *xmlx.Node) (obj *cdom.GeometryPrimitiveKind) {
	obj = new(cdom.GeometryPrimitiveKind)

	load_GeometryPrimitiveKind(xn, obj)
	return
}

func init_GeometryBrepSurface(xn *xmlx.Node) (obj *cdom.GeometryBrepSurface) {
	obj = new(cdom.GeometryBrepSurface)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_GeometryBrepSurface(xn, obj)
	return
}

func init_LightDef(xn *xmlx.Node) (obj *cdom.LightDef) {
	obj = new(cdom.LightDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_LightDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxPassEvaluationTarget(xn *xmlx.Node) (obj *cdom.FxPassEvaluationTarget) {
	obj = cdom.NewFxPassEvaluationTarget()

	load_FxPassEvaluationTarget(xn, obj)
	return
}

func init_CameraPerspective(xn *xmlx.Node) (obj *cdom.CameraPerspective) {
	obj = new(cdom.CameraPerspective)

	load_CameraPerspective(xn, obj)
	return
}

func init_Float4(xn *xmlx.Node) (obj *cdom.Float4) {
	obj = new(cdom.Float4)

	load_Float4(xn, obj)
	return
}

func init_KxModelDef(xn *xmlx.Node) (obj *cdom.KxModelDef) {
	obj = new(cdom.KxModelDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_KxModelDef(xn, obj)
	obj.SetDirty()
	return
}

func init_AnimationInst(xn *xmlx.Node) (obj *cdom.AnimationInst) {
	obj = new(cdom.AnimationInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_AnimationInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxFormatChannels(xn *xmlx.Node) (obj *cdom.FxFormatChannels) {
	obj = new(cdom.FxFormatChannels)

	load_FxFormatChannels(xn, obj)
	return
}

func init_KxKinematicsSystem(xn *xmlx.Node) (obj *cdom.KxKinematicsSystem) {
	obj = new(cdom.KxKinematicsSystem)
	has_Techniques(xn, &obj.HasTechniques)

	load_KxKinematicsSystem(xn, obj)
	return
}

func init_KxKinematicsAxis(xn *xmlx.Node) (obj *cdom.KxKinematicsAxis) {
	obj = cdom.NewKxKinematicsAxis()
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_Sid(xn, &obj.HasSid)

	load_KxKinematicsAxis(xn, obj)
	return
}

func init_FxBinding(xn *xmlx.Node) (obj *cdom.FxBinding) {
	obj = new(cdom.FxBinding)

	load_FxBinding(xn, obj)
	return
}

func init_Float3x2(xn *xmlx.Node) (obj *cdom.Float3x2) {
	obj = new(cdom.Float3x2)

	load_Float3x2(xn, obj)
	return
}

func init_FxSamplerStates(xn *xmlx.Node) (obj *cdom.FxSamplerStates) {
	obj = cdom.NewFxSamplerStates()
	has_Extras(xn, &obj.HasExtras)

	load_FxSamplerStates(xn, obj)
	return
}

func init_PxForceFieldInst(xn *xmlx.Node) (obj *cdom.PxForceFieldInst) {
	obj = new(cdom.PxForceFieldInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_PxForceFieldInst(xn, obj)
	obj.SetDirty()
	return
}

func init_KxAxisLimits(xn *xmlx.Node) (obj *cdom.KxAxisLimits) {
	obj = new(cdom.KxAxisLimits)

	load_KxAxisLimits(xn, obj)
	return
}

func init_PxRigidConstraintLimit(xn *xmlx.Node) (obj *cdom.PxRigidConstraintLimit) {
	obj = new(cdom.PxRigidConstraintLimit)

	load_PxRigidConstraintLimit(xn, obj)
	return
}

func init_PxMaterialInst(xn *xmlx.Node) (obj *cdom.PxMaterialInst) {
	obj = new(cdom.PxMaterialInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_PxMaterialInst(xn, obj)
	obj.SetDirty()
	return
}

func init_VisualSceneRenderingMaterialInst(xn *xmlx.Node) (obj *cdom.VisualSceneRenderingMaterialInst) {
	obj = new(cdom.VisualSceneRenderingMaterialInst)
	has_Extras(xn, &obj.HasExtras)

	load_VisualSceneRenderingMaterialInst(xn, obj)
	return
}

func init_AnimationClipInst(xn *xmlx.Node) (obj *cdom.AnimationClipInst) {
	obj = new(cdom.AnimationClipInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_AnimationClipInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxAnnotation(xn *xmlx.Node) (obj *cdom.FxAnnotation) {
	obj = new(cdom.FxAnnotation)
	has_Name(xn, &obj.HasName)

	load_FxAnnotation(xn, obj)
	return
}

func init_GeometryBrepEllipse(xn *xmlx.Node) (obj *cdom.GeometryBrepEllipse) {
	obj = new(cdom.GeometryBrepEllipse)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepEllipse(xn, obj)
	return
}

func init_GeometryBrepSurfaces(xn *xmlx.Node) (obj *cdom.GeometryBrepSurfaces) {
	obj = new(cdom.GeometryBrepSurfaces)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepSurfaces(xn, obj)
	return
}

func init_GeometryBrepCurves(xn *xmlx.Node) (obj *cdom.GeometryBrepCurves) {
	obj = new(cdom.GeometryBrepCurves)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepCurves(xn, obj)
	return
}

func init_Int3x3(xn *xmlx.Node) (obj *cdom.Int3x3) {
	obj = new(cdom.Int3x3)

	load_Int3x3(xn, obj)
	return
}

func init_AnimationClipDef(xn *xmlx.Node) (obj *cdom.AnimationClipDef) {
	obj = new(cdom.AnimationClipDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_AnimationClipDef(xn, obj)
	obj.SetDirty()
	return
}

func init_PxRigidConstraintInst(xn *xmlx.Node) (obj *cdom.PxRigidConstraintInst) {
	obj = new(cdom.PxRigidConstraintInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_PxRigidConstraintInst(xn, obj)
	obj.SetDirty()
	return
}

func init_GeometrySpline(xn *xmlx.Node) (obj *cdom.GeometrySpline) {
	obj = cdom.NewGeometrySpline()
	has_Extras(xn, &obj.HasExtras)
	has_Sources(xn, &obj.HasSources)

	load_GeometrySpline(xn, obj)
	return
}

func init_Bool2(xn *xmlx.Node) (obj *cdom.Bool2) {
	obj = new(cdom.Bool2)

	load_Bool2(xn, obj)
	return
}

func init_GeometryBrepShells(xn *xmlx.Node) (obj *cdom.GeometryBrepShells) {
	obj = new(cdom.GeometryBrepShells)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryBrepShells(xn, obj)
	return
}

func init_KxFrameOrigin(xn *xmlx.Node) (obj *cdom.KxFrameOrigin) {
	obj = new(cdom.KxFrameOrigin)

	load_KxFrameOrigin(xn, obj)
	return
}

func init_FxColorOrTexture(xn *xmlx.Node) (obj *cdom.FxColorOrTexture) {
	obj = new(cdom.FxColorOrTexture)

	load_FxColorOrTexture(xn, obj)
	return
}

func init_ParamOrRefSid(xn *xmlx.Node) (obj *cdom.ParamOrRefSid) {
	obj = new(cdom.ParamOrRefSid)

	load_ParamOrRefSid(xn, obj)
	return
}

func init_PxModelInst(xn *xmlx.Node) (obj *cdom.PxModelInst) {
	obj = new(cdom.PxModelInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_PxModelInst(xn, obj)
	obj.SetDirty()
	return
}

func init_Extra(xn *xmlx.Node) (obj *cdom.Extra) {
	obj = new(cdom.Extra)
	has_Asset(xn, &obj.HasAsset)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_Extra(xn, obj)
	return
}

func init_GeometryControlVertices(xn *xmlx.Node) (obj *cdom.GeometryControlVertices) {
	obj = new(cdom.GeometryControlVertices)
	has_Extras(xn, &obj.HasExtras)
	has_Inputs(xn, &obj.HasInputs)

	load_GeometryControlVertices(xn, obj)
	return
}

func init_SidBool(xn *xmlx.Node) (obj *cdom.SidBool) {
	obj = new(cdom.SidBool)
	has_Sid(xn, &obj.HasSid)

	load_SidBool(xn, obj)
	return
}

func init_FxCubeFace(xn *xmlx.Node) (obj *cdom.FxCubeFace) {
	obj = new(cdom.FxCubeFace)

	load_FxCubeFace(xn, obj)
	return
}

func init_Source(xn *xmlx.Node) (obj *cdom.Source) {
	obj = new(cdom.Source)
	has_Asset(xn, &obj.HasAsset)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_Source(xn, obj)
	return
}

func init_KxJointKind(xn *xmlx.Node) (obj *cdom.KxJointKind) {
	obj = new(cdom.KxJointKind)

	load_KxJointKind(xn, obj)
	return
}

func init_KxAttachment(xn *xmlx.Node) (obj *cdom.KxAttachment) {
	obj = new(cdom.KxAttachment)

	load_KxAttachment(xn, obj)
	return
}

func init_PxShape(xn *xmlx.Node) (obj *cdom.PxShape) {
	obj = new(cdom.PxShape)
	has_Extras(xn, &obj.HasExtras)

	load_PxShape(xn, obj)
	return
}

func init_GeometryBrepHyperbola(xn *xmlx.Node) (obj *cdom.GeometryBrepHyperbola) {
	obj = new(cdom.GeometryBrepHyperbola)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepHyperbola(xn, obj)
	return
}

func init_FxProfileCommon(xn *xmlx.Node) (obj *cdom.FxProfileCommon) {
	obj = new(cdom.FxProfileCommon)

	load_FxProfileCommon(xn, obj)
	return
}

func init_GeometryBrepNurbsSurface(xn *xmlx.Node) (obj *cdom.GeometryBrepNurbsSurface) {
	obj = cdom.NewGeometryBrepNurbsSurface()
	has_Extras(xn, &obj.HasExtras)
	has_Sources(xn, &obj.HasSources)

	load_GeometryBrepNurbsSurface(xn, obj)
	return
}

func init_ParamOrUint(xn *xmlx.Node) (obj *cdom.ParamOrUint) {
	obj = new(cdom.ParamOrUint)

	load_ParamOrUint(xn, obj)
	return
}

func init_FxPassProgram(xn *xmlx.Node) (obj *cdom.FxPassProgram) {
	obj = new(cdom.FxPassProgram)

	load_FxPassProgram(xn, obj)
	return
}

func init_FxPassProgramBindUniform(xn *xmlx.Node) (obj *cdom.FxPassProgramBindUniform) {
	obj = new(cdom.FxPassProgramBindUniform)

	load_FxPassProgramBindUniform(xn, obj)
	return
}

func init_GeometryPositioning(xn *xmlx.Node) (obj *cdom.GeometryPositioning) {
	obj = new(cdom.GeometryPositioning)

	load_GeometryPositioning(xn, obj)
	return
}

func init_GeometryBrepSolids(xn *xmlx.Node) (obj *cdom.GeometryBrepSolids) {
	obj = new(cdom.GeometryBrepSolids)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryBrepSolids(xn, obj)
	return
}

func init_FxPassEvaluationClearDepth(xn *xmlx.Node) (obj *cdom.FxPassEvaluationClearDepth) {
	obj = new(cdom.FxPassEvaluationClearDepth)

	load_FxPassEvaluationClearDepth(xn, obj)
	return
}

func init_SidFloat(xn *xmlx.Node) (obj *cdom.SidFloat) {
	obj = new(cdom.SidFloat)
	has_Sid(xn, &obj.HasSid)

	load_SidFloat(xn, obj)
	return
}

func init_KxModelInst(xn *xmlx.Node) (obj *cdom.KxModelInst) {
	obj = new(cdom.KxModelInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_KxModelInst(xn, obj)
	obj.SetDirty()
	return
}

func init_Int4(xn *xmlx.Node) (obj *cdom.Int4) {
	obj = new(cdom.Int4)

	load_Int4(xn, obj)
	return
}

func init_KxAxisIndex(xn *xmlx.Node) (obj *cdom.KxAxisIndex) {
	obj = new(cdom.KxAxisIndex)

	load_KxAxisIndex(xn, obj)
	return
}

func init_FxTechniqueCommon(xn *xmlx.Node) (obj *cdom.FxTechniqueCommon) {
	obj = new(cdom.FxTechniqueCommon)
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Sid(xn, &obj.HasSid)

	load_FxTechniqueCommon(xn, obj)
	return
}

func init_AnimationChannel(xn *xmlx.Node) (obj *cdom.AnimationChannel) {
	obj = new(cdom.AnimationChannel)

	load_AnimationChannel(xn, obj)
	return
}

func init_FxShaderStage(xn *xmlx.Node) (obj *cdom.FxShaderStage) {
	obj = new(cdom.FxShaderStage)

	load_FxShaderStage(xn, obj)
	return
}
