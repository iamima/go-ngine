package core

import (
	ugl "github.com/go3d/go-opengl/util"
)

type RenderTechniqueQuad struct {
	Effect FxEffect

	fxTex *FxOpTex2D
	renderTechniqueBase
	glVao ugl.VertexArray
}

func newRenderTechniqueQuad(progName string) RenderTechnique {
	me := &RenderTechniqueQuad{}
	me.init("Quad")
	me.Effect.init()
	me.fxTex = me.Effect.Ops.EnableTex2D(-1)
	me.fxTex.Sampler = &Core.Rendering.Fx.Samplers.NoFilteringClamp
	me.Effect.UpdateRoutine()
	me.glVao.Create()
	me.glVao.Setup(nil)
	return me
}

func (me *RenderTechniqueQuad) dispose() {
	me.glVao.Dispose()
}
