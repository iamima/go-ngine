// File generated by github.com/ungerik/pkgreflect
package pkgreflect

import "reflect"
import nga "github.com/go3d/go-ngine/assets"

var Types = map[string]reflect.Type{
	"LibKxModelDefs": reflect.TypeOf((*nga.LibKxModelDefs)(nil)).Elem(),
	"KxLink": reflect.TypeOf((*nga.KxLink)(nil)).Elem(),
	"KxModelDef": reflect.TypeOf((*nga.KxModelDef)(nil)).Elem(),
	"KxModelInst": reflect.TypeOf((*nga.KxModelInst)(nil)).Elem(),
	"LibsKxModelDef": reflect.TypeOf((*nga.LibsKxModelDef)(nil)).Elem(),
	"KxAttachment": reflect.TypeOf((*nga.KxAttachment)(nil)).Elem(),
	"KxAttachmentKind": reflect.TypeOf((*nga.KxAttachmentKind)(nil)).Elem(),
	"LibAnimationClipDefs": reflect.TypeOf((*nga.LibAnimationClipDefs)(nil)).Elem(),
	"AnimationClipDef": reflect.TypeOf((*nga.AnimationClipDef)(nil)).Elem(),
	"LibsAnimationClipDef": reflect.TypeOf((*nga.LibsAnimationClipDef)(nil)).Elem(),
	"AnimationClipInst": reflect.TypeOf((*nga.AnimationClipInst)(nil)).Elem(),
	"KxFrameTcp": reflect.TypeOf((*nga.KxFrameTcp)(nil)).Elem(),
	"KxFrameObject": reflect.TypeOf((*nga.KxFrameObject)(nil)).Elem(),
	"KxFrameTip": reflect.TypeOf((*nga.KxFrameTip)(nil)).Elem(),
	"KxBinding": reflect.TypeOf((*nga.KxBinding)(nil)).Elem(),
	"KxKinematicsSystem": reflect.TypeOf((*nga.KxKinematicsSystem)(nil)).Elem(),
	"KxAxisIndex": reflect.TypeOf((*nga.KxAxisIndex)(nil)).Elem(),
	"KxArticulatedSystemInst": reflect.TypeOf((*nga.KxArticulatedSystemInst)(nil)).Elem(),
	"KxMotionAxis": reflect.TypeOf((*nga.KxMotionAxis)(nil)).Elem(),
	"KxAxisLimits": reflect.TypeOf((*nga.KxAxisLimits)(nil)).Elem(),
	"LibKxArticulatedSystemDefs": reflect.TypeOf((*nga.LibKxArticulatedSystemDefs)(nil)).Elem(),
	"KxArticulatedSystemDef": reflect.TypeOf((*nga.KxArticulatedSystemDef)(nil)).Elem(),
	"LibsKxArticulatedSystemDef": reflect.TypeOf((*nga.LibsKxArticulatedSystemDef)(nil)).Elem(),
	"KxFrameOrigin": reflect.TypeOf((*nga.KxFrameOrigin)(nil)).Elem(),
	"KxMotionSystem": reflect.TypeOf((*nga.KxMotionSystem)(nil)).Elem(),
	"KxEffector": reflect.TypeOf((*nga.KxEffector)(nil)).Elem(),
	"KxKinematicsAxis": reflect.TypeOf((*nga.KxKinematicsAxis)(nil)).Elem(),
	"KxFrame": reflect.TypeOf((*nga.KxFrame)(nil)).Elem(),
	"Sources": reflect.TypeOf((*nga.Sources)(nil)).Elem(),
	"SourceArray": reflect.TypeOf((*nga.SourceArray)(nil)).Elem(),
	"SourceAccessor": reflect.TypeOf((*nga.SourceAccessor)(nil)).Elem(),
	"Source": reflect.TypeOf((*nga.Source)(nil)).Elem(),
	"AnimationDef": reflect.TypeOf((*nga.AnimationDef)(nil)).Elem(),
	"AnimSamplerBehavior": reflect.TypeOf((*nga.AnimSamplerBehavior)(nil)).Elem(),
	"LibAnimationDefs": reflect.TypeOf((*nga.LibAnimationDefs)(nil)).Elem(),
	"LibsAnimationDef": reflect.TypeOf((*nga.LibsAnimationDef)(nil)).Elem(),
	"AnimationInst": reflect.TypeOf((*nga.AnimationInst)(nil)).Elem(),
	"AnimationChannel": reflect.TypeOf((*nga.AnimationChannel)(nil)).Elem(),
	"AnimationSampler": reflect.TypeOf((*nga.AnimationSampler)(nil)).Elem(),
	"FxBinding": reflect.TypeOf((*nga.FxBinding)(nil)).Elem(),
	"FxMaterialDef": reflect.TypeOf((*nga.FxMaterialDef)(nil)).Elem(),
	"LibsFxMaterialDef": reflect.TypeOf((*nga.LibsFxMaterialDef)(nil)).Elem(),
	"LibFxMaterialDefs": reflect.TypeOf((*nga.LibFxMaterialDefs)(nil)).Elem(),
	"FxVertexInputBinding": reflect.TypeOf((*nga.FxVertexInputBinding)(nil)).Elem(),
	"FxMaterialInst": reflect.TypeOf((*nga.FxMaterialInst)(nil)).Elem(),
	"LibPxModelDefs": reflect.TypeOf((*nga.LibPxModelDefs)(nil)).Elem(),
	"LibsPxModelDef": reflect.TypeOf((*nga.LibsPxModelDef)(nil)).Elem(),
	"PxModelDef": reflect.TypeOf((*nga.PxModelDef)(nil)).Elem(),
	"PxModelInst": reflect.TypeOf((*nga.PxModelInst)(nil)).Elem(),
	"Float2x3": reflect.TypeOf((*nga.Float2x3)(nil)).Elem(),
	"Bool4": reflect.TypeOf((*nga.Bool4)(nil)).Elem(),
	"SidVec3": reflect.TypeOf((*nga.SidVec3)(nil)).Elem(),
	"Int2x2": reflect.TypeOf((*nga.Int2x2)(nil)).Elem(),
	"Float3": reflect.TypeOf((*nga.Float3)(nil)).Elem(),
	"Int3": reflect.TypeOf((*nga.Int3)(nil)).Elem(),
	"Float3x3": reflect.TypeOf((*nga.Float3x3)(nil)).Elem(),
	"Int3x3": reflect.TypeOf((*nga.Int3x3)(nil)).Elem(),
	"Bool3": reflect.TypeOf((*nga.Bool3)(nil)).Elem(),
	"Int4x4": reflect.TypeOf((*nga.Int4x4)(nil)).Elem(),
	"SidString": reflect.TypeOf((*nga.SidString)(nil)).Elem(),
	"ParamOrFloat2": reflect.TypeOf((*nga.ParamOrFloat2)(nil)).Elem(),
	"ParamOrBool": reflect.TypeOf((*nga.ParamOrBool)(nil)).Elem(),
	"ParamOrRefSid": reflect.TypeOf((*nga.ParamOrRefSid)(nil)).Elem(),
	"ParamOrUint": reflect.TypeOf((*nga.ParamOrUint)(nil)).Elem(),
	"Float4": reflect.TypeOf((*nga.Float4)(nil)).Elem(),
	"SidFloat3": reflect.TypeOf((*nga.SidFloat3)(nil)).Elem(),
	"Int4": reflect.TypeOf((*nga.Int4)(nil)).Elem(),
	"Float3x2": reflect.TypeOf((*nga.Float3x2)(nil)).Elem(),
	"Float4x3": reflect.TypeOf((*nga.Float4x3)(nil)).Elem(),
	"Int2": reflect.TypeOf((*nga.Int2)(nil)).Elem(),
	"Float4x2": reflect.TypeOf((*nga.Float4x2)(nil)).Elem(),
	"Bool2": reflect.TypeOf((*nga.Bool2)(nil)).Elem(),
	"ParamOrInt": reflect.TypeOf((*nga.ParamOrInt)(nil)).Elem(),
	"Float4x4": reflect.TypeOf((*nga.Float4x4)(nil)).Elem(),
	"ParamOrFloat": reflect.TypeOf((*nga.ParamOrFloat)(nil)).Elem(),
	"Float7": reflect.TypeOf((*nga.Float7)(nil)).Elem(),
	"Float3x4": reflect.TypeOf((*nga.Float3x4)(nil)).Elem(),
	"Float2x4": reflect.TypeOf((*nga.Float2x4)(nil)).Elem(),
	"ParamOrSidFloat": reflect.TypeOf((*nga.ParamOrSidFloat)(nil)).Elem(),
	"Float2x2": reflect.TypeOf((*nga.Float2x2)(nil)).Elem(),
	"Float2": reflect.TypeOf((*nga.Float2)(nil)).Elem(),
	"SidBool": reflect.TypeOf((*nga.SidBool)(nil)).Elem(),
	"SidFloat": reflect.TypeOf((*nga.SidFloat)(nil)).Elem(),
	"Param": reflect.TypeOf((*nga.Param)(nil)).Elem(),
	"HasAsset": reflect.TypeOf((*nga.HasAsset)(nil)).Elem(),
	"Layers": reflect.TypeOf((*nga.Layers)(nil)).Elem(),
	"Technique": reflect.TypeOf((*nga.Technique)(nil)).Elem(),
	"ParamDef": reflect.TypeOf((*nga.ParamDef)(nil)).Elem(),
	"Scene": reflect.TypeOf((*nga.Scene)(nil)).Elem(),
	"AssetGeographicLocation": reflect.TypeOf((*nga.AssetGeographicLocation)(nil)).Elem(),
	"HasExtras": reflect.TypeOf((*nga.HasExtras)(nil)).Elem(),
	"ParamDefs": reflect.TypeOf((*nga.ParamDefs)(nil)).Elem(),
	"InputShared": reflect.TypeOf((*nga.InputShared)(nil)).Elem(),
	"HasInputs": reflect.TypeOf((*nga.HasInputs)(nil)).Elem(),
	"AssetContributor": reflect.TypeOf((*nga.AssetContributor)(nil)).Elem(),
	"HasId": reflect.TypeOf((*nga.HasId)(nil)).Elem(),
	"HasParamInsts": reflect.TypeOf((*nga.HasParamInsts)(nil)).Elem(),
	"HasParamDefs": reflect.TypeOf((*nga.HasParamDefs)(nil)).Elem(),
	"HasSid": reflect.TypeOf((*nga.HasSid)(nil)).Elem(),
	"IndexedInputs": reflect.TypeOf((*nga.IndexedInputs)(nil)).Elem(),
	"HasFxParamDefs": reflect.TypeOf((*nga.HasFxParamDefs)(nil)).Elem(),
	"TransformKind": reflect.TypeOf((*nga.TransformKind)(nil)).Elem(),
	"HasName": reflect.TypeOf((*nga.HasName)(nil)).Elem(),
	"MaterialBinding": reflect.TypeOf((*nga.MaterialBinding)(nil)).Elem(),
	"HasSources": reflect.TypeOf((*nga.HasSources)(nil)).Elem(),
	"ParamInst": reflect.TypeOf((*nga.ParamInst)(nil)).Elem(),
	"ParamInsts": reflect.TypeOf((*nga.ParamInsts)(nil)).Elem(),
	"HasTechniques": reflect.TypeOf((*nga.HasTechniques)(nil)).Elem(),
	"Input": reflect.TypeOf((*nga.Input)(nil)).Elem(),
	"Transform": reflect.TypeOf((*nga.Transform)(nil)).Elem(),
	"Extra": reflect.TypeOf((*nga.Extra)(nil)).Elem(),
	"Asset": reflect.TypeOf((*nga.Asset)(nil)).Elem(),
	"GeometryPrimitives": reflect.TypeOf((*nga.GeometryPrimitives)(nil)).Elem(),
	"LibGeometryDefs": reflect.TypeOf((*nga.LibGeometryDefs)(nil)).Elem(),
	"GeometryPrimitiveKind": reflect.TypeOf((*nga.GeometryPrimitiveKind)(nil)).Elem(),
	"GeometryVertices": reflect.TypeOf((*nga.GeometryVertices)(nil)).Elem(),
	"GeometryMesh": reflect.TypeOf((*nga.GeometryMesh)(nil)).Elem(),
	"LibsGeometryDef": reflect.TypeOf((*nga.LibsGeometryDef)(nil)).Elem(),
	"GeometryControlVertices": reflect.TypeOf((*nga.GeometryControlVertices)(nil)).Elem(),
	"GeometryDef": reflect.TypeOf((*nga.GeometryDef)(nil)).Elem(),
	"GeometrySpline": reflect.TypeOf((*nga.GeometrySpline)(nil)).Elem(),
	"GeometryPolygonHole": reflect.TypeOf((*nga.GeometryPolygonHole)(nil)).Elem(),
	"GeometryInst": reflect.TypeOf((*nga.GeometryInst)(nil)).Elem(),
	"PxMaterial": reflect.TypeOf((*nga.PxMaterial)(nil)).Elem(),
	"PxRigidBodyDefs": reflect.TypeOf((*nga.PxRigidBodyDefs)(nil)).Elem(),
	"PxRigidBodyDef": reflect.TypeOf((*nga.PxRigidBodyDef)(nil)).Elem(),
	"PxRigidBodyInst": reflect.TypeOf((*nga.PxRigidBodyInst)(nil)).Elem(),
	"PxShape": reflect.TypeOf((*nga.PxShape)(nil)).Elem(),
	"PxRigidBodyCommon": reflect.TypeOf((*nga.PxRigidBodyCommon)(nil)).Elem(),
	"PxCylinder": reflect.TypeOf((*nga.PxCylinder)(nil)).Elem(),
	"MeshVert": reflect.TypeOf((*nga.MeshVert)(nil)).Elem(),
	"MeshFace3": reflect.TypeOf((*nga.MeshFace3)(nil)).Elem(),
	"MeshProvider": reflect.TypeOf((*nga.MeshProvider)(nil)).Elem(),
	"MeshRaw": reflect.TypeOf((*nga.MeshRaw)(nil)).Elem(),
	"MeshVertAtt2": reflect.TypeOf((*nga.MeshVertAtt2)(nil)).Elem(),
	"MeshRawFace": reflect.TypeOf((*nga.MeshRawFace)(nil)).Elem(),
	"MeshData": reflect.TypeOf((*nga.MeshData)(nil)).Elem(),
	"MeshVertAtt3": reflect.TypeOf((*nga.MeshVertAtt3)(nil)).Elem(),
	"RefSidRoot": reflect.TypeOf((*nga.RefSidRoot)(nil)).Elem(),
	"RefParam": reflect.TypeOf((*nga.RefParam)(nil)).Elem(),
	"RefSid": reflect.TypeOf((*nga.RefSid)(nil)).Elem(),
	"VisualSceneRenderingMaterialInst": reflect.TypeOf((*nga.VisualSceneRenderingMaterialInst)(nil)).Elem(),
	"VisualSceneDef": reflect.TypeOf((*nga.VisualSceneDef)(nil)).Elem(),
	"VisualSceneEvaluation": reflect.TypeOf((*nga.VisualSceneEvaluation)(nil)).Elem(),
	"LibVisualSceneDefs": reflect.TypeOf((*nga.LibVisualSceneDefs)(nil)).Elem(),
	"LibsVisualSceneDef": reflect.TypeOf((*nga.LibsVisualSceneDef)(nil)).Elem(),
	"VisualSceneRendering": reflect.TypeOf((*nga.VisualSceneRendering)(nil)).Elem(),
	"VisualSceneInst": reflect.TypeOf((*nga.VisualSceneInst)(nil)).Elem(),
	"NodeInst": reflect.TypeOf((*nga.NodeInst)(nil)).Elem(),
	"NodeDef": reflect.TypeOf((*nga.NodeDef)(nil)).Elem(),
	"ChildNode": reflect.TypeOf((*nga.ChildNode)(nil)).Elem(),
	"LibsNodeDef": reflect.TypeOf((*nga.LibsNodeDef)(nil)).Elem(),
	"LibNodeDefs": reflect.TypeOf((*nga.LibNodeDefs)(nil)).Elem(),
	"LibsPxSceneDef": reflect.TypeOf((*nga.LibsPxSceneDef)(nil)).Elem(),
	"LibPxSceneDefs": reflect.TypeOf((*nga.LibPxSceneDefs)(nil)).Elem(),
	"PxSceneDef": reflect.TypeOf((*nga.PxSceneDef)(nil)).Elem(),
	"PxSceneInst": reflect.TypeOf((*nga.PxSceneInst)(nil)).Elem(),
	"LightAttenuation": reflect.TypeOf((*nga.LightAttenuation)(nil)).Elem(),
	"LightInst": reflect.TypeOf((*nga.LightInst)(nil)).Elem(),
	"LightDirectional": reflect.TypeOf((*nga.LightDirectional)(nil)).Elem(),
	"LightBase": reflect.TypeOf((*nga.LightBase)(nil)).Elem(),
	"LibLightDefs": reflect.TypeOf((*nga.LibLightDefs)(nil)).Elem(),
	"LightAmbient": reflect.TypeOf((*nga.LightAmbient)(nil)).Elem(),
	"LibsLightDef": reflect.TypeOf((*nga.LibsLightDef)(nil)).Elem(),
	"LightSpot": reflect.TypeOf((*nga.LightSpot)(nil)).Elem(),
	"LightDef": reflect.TypeOf((*nga.LightDef)(nil)).Elem(),
	"LightPoint": reflect.TypeOf((*nga.LightPoint)(nil)).Elem(),
	"LibsCameraDef": reflect.TypeOf((*nga.LibsCameraDef)(nil)).Elem(),
	"CameraDef": reflect.TypeOf((*nga.CameraDef)(nil)).Elem(),
	"CameraOrthographic": reflect.TypeOf((*nga.CameraOrthographic)(nil)).Elem(),
	"CameraImager": reflect.TypeOf((*nga.CameraImager)(nil)).Elem(),
	"CameraPerspective": reflect.TypeOf((*nga.CameraPerspective)(nil)).Elem(),
	"CameraOptics": reflect.TypeOf((*nga.CameraOptics)(nil)).Elem(),
	"CameraInst": reflect.TypeOf((*nga.CameraInst)(nil)).Elem(),
	"LibCameraDefs": reflect.TypeOf((*nga.LibCameraDefs)(nil)).Elem(),
	"RefId": reflect.TypeOf((*nga.RefId)(nil)).Elem(),
	"GeometryBrepEdges": reflect.TypeOf((*nga.GeometryBrepEdges)(nil)).Elem(),
	"GeometryBrepOrientation": reflect.TypeOf((*nga.GeometryBrepOrientation)(nil)).Elem(),
	"GeometryBrep": reflect.TypeOf((*nga.GeometryBrep)(nil)).Elem(),
	"GeometryBrepWires": reflect.TypeOf((*nga.GeometryBrepWires)(nil)).Elem(),
	"GeometryBrepPlane": reflect.TypeOf((*nga.GeometryBrepPlane)(nil)).Elem(),
	"GeometryBrepSphere": reflect.TypeOf((*nga.GeometryBrepSphere)(nil)).Elem(),
	"GeometryBrepEllipse": reflect.TypeOf((*nga.GeometryBrepEllipse)(nil)).Elem(),
	"GeometryBrepParabola": reflect.TypeOf((*nga.GeometryBrepParabola)(nil)).Elem(),
	"GeometryBrepLine": reflect.TypeOf((*nga.GeometryBrepLine)(nil)).Elem(),
	"GeometryBrepTorus": reflect.TypeOf((*nga.GeometryBrepTorus)(nil)).Elem(),
	"GeometryBrepNurbs": reflect.TypeOf((*nga.GeometryBrepNurbs)(nil)).Elem(),
	"GeometryBrepBox": reflect.TypeOf((*nga.GeometryBrepBox)(nil)).Elem(),
	"GeometryBrepCone": reflect.TypeOf((*nga.GeometryBrepCone)(nil)).Elem(),
	"GeometryPositioning": reflect.TypeOf((*nga.GeometryPositioning)(nil)).Elem(),
	"GeometryBrepCircle": reflect.TypeOf((*nga.GeometryBrepCircle)(nil)).Elem(),
	"GeometryBrepHyperbola": reflect.TypeOf((*nga.GeometryBrepHyperbola)(nil)).Elem(),
	"GeometryBrepPcurves": reflect.TypeOf((*nga.GeometryBrepPcurves)(nil)).Elem(),
	"GeometryBrepFaces": reflect.TypeOf((*nga.GeometryBrepFaces)(nil)).Elem(),
	"GeometryBrepSurfaceCurves": reflect.TypeOf((*nga.GeometryBrepSurfaceCurves)(nil)).Elem(),
	"GeometryBrepSweptSurface": reflect.TypeOf((*nga.GeometryBrepSweptSurface)(nil)).Elem(),
	"GeometryBrepCapsule": reflect.TypeOf((*nga.GeometryBrepCapsule)(nil)).Elem(),
	"GeometryBrepNurbsSurface": reflect.TypeOf((*nga.GeometryBrepNurbsSurface)(nil)).Elem(),
	"GeometryBrepShells": reflect.TypeOf((*nga.GeometryBrepShells)(nil)).Elem(),
	"GeometryBrepSurfaces": reflect.TypeOf((*nga.GeometryBrepSurfaces)(nil)).Elem(),
	"GeometryBrepSolids": reflect.TypeOf((*nga.GeometryBrepSolids)(nil)).Elem(),
	"GeometryBrepCurve": reflect.TypeOf((*nga.GeometryBrepCurve)(nil)).Elem(),
	"GeometryBrepCurves": reflect.TypeOf((*nga.GeometryBrepCurves)(nil)).Elem(),
	"GeometryBrepCylinder": reflect.TypeOf((*nga.GeometryBrepCylinder)(nil)).Elem(),
	"GeometryBrepSurface": reflect.TypeOf((*nga.GeometryBrepSurface)(nil)).Elem(),
	"ControllerMorph": reflect.TypeOf((*nga.ControllerMorph)(nil)).Elem(),
	"LibsControllerDef": reflect.TypeOf((*nga.LibsControllerDef)(nil)).Elem(),
	"ControllerInst": reflect.TypeOf((*nga.ControllerInst)(nil)).Elem(),
	"ControllerSkin": reflect.TypeOf((*nga.ControllerSkin)(nil)).Elem(),
	"LibControllerDefs": reflect.TypeOf((*nga.LibControllerDefs)(nil)).Elem(),
	"ControllerDef": reflect.TypeOf((*nga.ControllerDef)(nil)).Elem(),
	"ControllerInputs": reflect.TypeOf((*nga.ControllerInputs)(nil)).Elem(),
	"FxInitFrom": reflect.TypeOf((*nga.FxInitFrom)(nil)).Elem(),
	"FxCubeFace": reflect.TypeOf((*nga.FxCubeFace)(nil)).Elem(),
	"FxFormatChannels": reflect.TypeOf((*nga.FxFormatChannels)(nil)).Elem(),
	"FxCreate3DInitFrom": reflect.TypeOf((*nga.FxCreate3DInitFrom)(nil)).Elem(),
	"LibFxImageDefs": reflect.TypeOf((*nga.LibFxImageDefs)(nil)).Elem(),
	"FxImageInst": reflect.TypeOf((*nga.FxImageInst)(nil)).Elem(),
	"FxCreate2DSizeRatio": reflect.TypeOf((*nga.FxCreate2DSizeRatio)(nil)).Elem(),
	"FxFormatRange": reflect.TypeOf((*nga.FxFormatRange)(nil)).Elem(),
	"FxCreateMips": reflect.TypeOf((*nga.FxCreateMips)(nil)).Elem(),
	"LibsFxImageDef": reflect.TypeOf((*nga.LibsFxImageDef)(nil)).Elem(),
	"FxCreate2D": reflect.TypeOf((*nga.FxCreate2D)(nil)).Elem(),
	"FxImageDef": reflect.TypeOf((*nga.FxImageDef)(nil)).Elem(),
	"FxCreateFormatHint": reflect.TypeOf((*nga.FxCreateFormatHint)(nil)).Elem(),
	"FxCreate3D": reflect.TypeOf((*nga.FxCreate3D)(nil)).Elem(),
	"FxCreate2DSizeExact": reflect.TypeOf((*nga.FxCreate2DSizeExact)(nil)).Elem(),
	"FxCreateInitFrom": reflect.TypeOf((*nga.FxCreateInitFrom)(nil)).Elem(),
	"FxFormatPrecision": reflect.TypeOf((*nga.FxFormatPrecision)(nil)).Elem(),
	"FxCreate": reflect.TypeOf((*nga.FxCreate)(nil)).Elem(),
	"FxCreateCube": reflect.TypeOf((*nga.FxCreateCube)(nil)).Elem(),
	"FxCreateFormat": reflect.TypeOf((*nga.FxCreateFormat)(nil)).Elem(),
	"FxImageInitFrom": reflect.TypeOf((*nga.FxImageInitFrom)(nil)).Elem(),
	"FxCreateCubeInitFrom": reflect.TypeOf((*nga.FxCreateCubeInitFrom)(nil)).Elem(),
	"BaseInst": reflect.TypeOf((*nga.BaseInst)(nil)).Elem(),
	"BaseSync": reflect.TypeOf((*nga.BaseSync)(nil)).Elem(),
	"BaseDef": reflect.TypeOf((*nga.BaseDef)(nil)).Elem(),
	"BaseLib": reflect.TypeOf((*nga.BaseLib)(nil)).Elem(),
	"KxJointKind": reflect.TypeOf((*nga.KxJointKind)(nil)).Elem(),
	"KxJoint": reflect.TypeOf((*nga.KxJoint)(nil)).Elem(),
	"LibKxJointDefs": reflect.TypeOf((*nga.LibKxJointDefs)(nil)).Elem(),
	"KxJointLimits": reflect.TypeOf((*nga.KxJointLimits)(nil)).Elem(),
	"KxJointDef": reflect.TypeOf((*nga.KxJointDef)(nil)).Elem(),
	"LibsKxJointDef": reflect.TypeOf((*nga.LibsKxJointDef)(nil)).Elem(),
	"KxJointInst": reflect.TypeOf((*nga.KxJointInst)(nil)).Elem(),
	"PxRigidConstraintSpring": reflect.TypeOf((*nga.PxRigidConstraintSpring)(nil)).Elem(),
	"PxRigidConstraintLimit": reflect.TypeOf((*nga.PxRigidConstraintLimit)(nil)).Elem(),
	"PxRigidConstraintDefs": reflect.TypeOf((*nga.PxRigidConstraintDefs)(nil)).Elem(),
	"PxRigidConstraintInst": reflect.TypeOf((*nga.PxRigidConstraintInst)(nil)).Elem(),
	"PxRigidConstraintAttachment": reflect.TypeOf((*nga.PxRigidConstraintAttachment)(nil)).Elem(),
	"PxRigidConstraintDef": reflect.TypeOf((*nga.PxRigidConstraintDef)(nil)).Elem(),
	"Document": reflect.TypeOf((*nga.Document)(nil)).Elem(),
	"FxParamDef": reflect.TypeOf((*nga.FxParamDef)(nil)).Elem(),
	"FxTechniqueGlsl": reflect.TypeOf((*nga.FxTechniqueGlsl)(nil)).Elem(),
	"FxEffectInstTechniqueHint": reflect.TypeOf((*nga.FxEffectInstTechniqueHint)(nil)).Elem(),
	"FxEffectDef": reflect.TypeOf((*nga.FxEffectDef)(nil)).Elem(),
	"FxProfile": reflect.TypeOf((*nga.FxProfile)(nil)).Elem(),
	"FxPassProgramBindUniform": reflect.TypeOf((*nga.FxPassProgramBindUniform)(nil)).Elem(),
	"FxAnnotation": reflect.TypeOf((*nga.FxAnnotation)(nil)).Elem(),
	"FxTechniqueCommonBlinn": reflect.TypeOf((*nga.FxTechniqueCommonBlinn)(nil)).Elem(),
	"FxGlslTechniques": reflect.TypeOf((*nga.FxGlslTechniques)(nil)).Elem(),
	"FxEffectInst": reflect.TypeOf((*nga.FxEffectInst)(nil)).Elem(),
	"FxPassEvaluationTarget": reflect.TypeOf((*nga.FxPassEvaluationTarget)(nil)).Elem(),
	"FxColor": reflect.TypeOf((*nga.FxColor)(nil)).Elem(),
	"FxPassEvaluationClearDepth": reflect.TypeOf((*nga.FxPassEvaluationClearDepth)(nil)).Elem(),
	"FxTechnique": reflect.TypeOf((*nga.FxTechnique)(nil)).Elem(),
	"FxPassProgram": reflect.TypeOf((*nga.FxPassProgram)(nil)).Elem(),
	"FxPassEvaluation": reflect.TypeOf((*nga.FxPassEvaluation)(nil)).Elem(),
	"FxPassProgramBindAttribute": reflect.TypeOf((*nga.FxPassProgramBindAttribute)(nil)).Elem(),
	"FxTechniqueCommonLambert": reflect.TypeOf((*nga.FxTechniqueCommonLambert)(nil)).Elem(),
	"FxColorOrTexture": reflect.TypeOf((*nga.FxColorOrTexture)(nil)).Elem(),
	"FxProfileGlsl": reflect.TypeOf((*nga.FxProfileGlsl)(nil)).Elem(),
	"FxTextureOpaque": reflect.TypeOf((*nga.FxTextureOpaque)(nil)).Elem(),
	"FxProfileCommon": reflect.TypeOf((*nga.FxProfileCommon)(nil)).Elem(),
	"FxProfileGlslCodeInclude": reflect.TypeOf((*nga.FxProfileGlslCodeInclude)(nil)).Elem(),
	"FxTechniqueCommonConstant": reflect.TypeOf((*nga.FxTechniqueCommonConstant)(nil)).Elem(),
	"FxPassEvaluationClearStencil": reflect.TypeOf((*nga.FxPassEvaluationClearStencil)(nil)).Elem(),
	"FxPass": reflect.TypeOf((*nga.FxPass)(nil)).Elem(),
	"FxTechniqueCommon": reflect.TypeOf((*nga.FxTechniqueCommon)(nil)).Elem(),
	"LibFxEffectDefs": reflect.TypeOf((*nga.LibFxEffectDefs)(nil)).Elem(),
	"FxTexture": reflect.TypeOf((*nga.FxTexture)(nil)).Elem(),
	"FxPassProgramShaderSources": reflect.TypeOf((*nga.FxPassProgramShaderSources)(nil)).Elem(),
	"FxShaderStage": reflect.TypeOf((*nga.FxShaderStage)(nil)).Elem(),
	"FxPassEvaluationClearColor": reflect.TypeOf((*nga.FxPassEvaluationClearColor)(nil)).Elem(),
	"FxPassState": reflect.TypeOf((*nga.FxPassState)(nil)).Elem(),
	"LibsFxEffectDef": reflect.TypeOf((*nga.LibsFxEffectDef)(nil)).Elem(),
	"FxParamDefs": reflect.TypeOf((*nga.FxParamDefs)(nil)).Elem(),
	"FxTechniqueCommonPhong": reflect.TypeOf((*nga.FxTechniqueCommonPhong)(nil)).Elem(),
	"FxPassProgramShader": reflect.TypeOf((*nga.FxPassProgramShader)(nil)).Elem(),
	"LibsFormulaDef": reflect.TypeOf((*nga.LibsFormulaDef)(nil)).Elem(),
	"FormulaDef": reflect.TypeOf((*nga.FormulaDef)(nil)).Elem(),
	"FormulaInst": reflect.TypeOf((*nga.FormulaInst)(nil)).Elem(),
	"Formula": reflect.TypeOf((*nga.Formula)(nil)).Elem(),
	"LibFormulaDefs": reflect.TypeOf((*nga.LibFormulaDefs)(nil)).Elem(),
	"PxMaterialInst": reflect.TypeOf((*nga.PxMaterialInst)(nil)).Elem(),
	"PxMaterialDef": reflect.TypeOf((*nga.PxMaterialDef)(nil)).Elem(),
	"LibPxMaterialDefs": reflect.TypeOf((*nga.LibPxMaterialDefs)(nil)).Elem(),
	"LibsPxMaterialDef": reflect.TypeOf((*nga.LibsPxMaterialDef)(nil)).Elem(),
	"FxSamplerWrapping": reflect.TypeOf((*nga.FxSamplerWrapping)(nil)).Elem(),
	"FxSamplerFiltering": reflect.TypeOf((*nga.FxSamplerFiltering)(nil)).Elem(),
	"FxSamplerKind": reflect.TypeOf((*nga.FxSamplerKind)(nil)).Elem(),
	"FxSamplerImage": reflect.TypeOf((*nga.FxSamplerImage)(nil)).Elem(),
	"FxSamplerStates": reflect.TypeOf((*nga.FxSamplerStates)(nil)).Elem(),
	"FxSampler": reflect.TypeOf((*nga.FxSampler)(nil)).Elem(),
	"FxFilterKind": reflect.TypeOf((*nga.FxFilterKind)(nil)).Elem(),
	"FxWrapKind": reflect.TypeOf((*nga.FxWrapKind)(nil)).Elem(),
	"KxJointAxisBinding": reflect.TypeOf((*nga.KxJointAxisBinding)(nil)).Elem(),
	"KxSceneDef": reflect.TypeOf((*nga.KxSceneDef)(nil)).Elem(),
	"KxSceneInst": reflect.TypeOf((*nga.KxSceneInst)(nil)).Elem(),
	"KxModelBinding": reflect.TypeOf((*nga.KxModelBinding)(nil)).Elem(),
	"LibKxSceneDefs": reflect.TypeOf((*nga.LibKxSceneDefs)(nil)).Elem(),
	"LibsKxSceneDef": reflect.TypeOf((*nga.LibsKxSceneDef)(nil)).Elem(),
	"PxForceFieldDef": reflect.TypeOf((*nga.PxForceFieldDef)(nil)).Elem(),
	"PxForceFieldInst": reflect.TypeOf((*nga.PxForceFieldInst)(nil)).Elem(),
	"LibPxForceFieldDefs": reflect.TypeOf((*nga.LibPxForceFieldDefs)(nil)).Elem(),
	"LibsPxForceFieldDef": reflect.TypeOf((*nga.LibsPxForceFieldDef)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"NewControllerMorph": reflect.ValueOf(nga.NewControllerMorph),
	"NewControllerSkin": reflect.ValueOf(nga.NewControllerSkin),
	"NewFxImageInitFrom": reflect.ValueOf(nga.NewFxImageInitFrom),
	"NewPxRigidConstraintSpring": reflect.ValueOf(nga.NewPxRigidConstraintSpring),
	"NewFxPass": reflect.ValueOf(nga.NewFxPass),
	"NewFxProfileGlsl": reflect.ValueOf(nga.NewFxProfileGlsl),
	"NewFxPassEvaluationTarget": reflect.ValueOf(nga.NewFxPassEvaluationTarget),
	"NewProfile": reflect.ValueOf(nga.NewProfile),
	"NewFxSamplerStates": reflect.ValueOf(nga.NewFxSamplerStates),
	"NewFxSampler": reflect.ValueOf(nga.NewFxSampler),
	"NewKxMotionAxis": reflect.ValueOf(nga.NewKxMotionAxis),
	"NewKxEffector": reflect.ValueOf(nga.NewKxEffector),
	"NewKxKinematicsAxis": reflect.ValueOf(nga.NewKxKinematicsAxis),
	"NewSourceAccessor": reflect.ValueOf(nga.NewSourceAccessor),
	"NewAsset": reflect.ValueOf(nga.NewAsset),
	"NewGeometrySpline": reflect.ValueOf(nga.NewGeometrySpline),
	"NewGeometryMesh": reflect.ValueOf(nga.NewGeometryMesh),
	"NewMeshData": reflect.ValueOf(nga.NewMeshData),
	"NewMeshRawFace": reflect.ValueOf(nga.NewMeshRawFace),
	"NewRefParam": reflect.ValueOf(nga.NewRefParam),
	"NewRefSid": reflect.ValueOf(nga.NewRefSid),
	"NewVisualSceneRendering": reflect.ValueOf(nga.NewVisualSceneRendering),
	"NewLightSpot": reflect.ValueOf(nga.NewLightSpot),
	"NewLightPoint": reflect.ValueOf(nga.NewLightPoint),
	"NewLightAttenuation": reflect.ValueOf(nga.NewLightAttenuation),
	"NewGeometryBrepNurbs": reflect.ValueOf(nga.NewGeometryBrepNurbs),
	"NewGeometryBrep": reflect.ValueOf(nga.NewGeometryBrep),
	"NewGeometryBrepNurbsSurface": reflect.ValueOf(nga.NewGeometryBrepNurbsSurface),
	"SyncChanges": reflect.ValueOf(nga.SyncChanges),
	"SidF": reflect.ValueOf(nga.SidF),
}

