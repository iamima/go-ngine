package exampleutils

import (
	ng "github.com/go3d/go-ngine/core"
	unum "github.com/metaleap/go-util/num"
)

//	A fake "2D GUI" concoction. There will be much better ngine-provided support for stuff like this.
//	Has two textured quads, a cat and a dog one, shows them both animated and overlapping
//	inside a red 64x64 px square in the bottom left canvas corner.
type Gui2D struct {
	Scene    ng.Scene
	Cam      *ng.Camera
	Cat, Dog *ng.Node
}

//	Adds a "2D camera" to the main render canvas, and sets up Cat and Dog.
func (me *Gui2D) Setup() (err error) {
	var (
		meshBuf  *ng.MeshBuffer
		quadMesh *ng.Mesh
	)
	me.Scene.Init()
	me.Cam = SceneCanvas.AddNewCamera2D(true)
	me.Cam.Rendering.States.ClearColor.Set(0.75, 0.25, 0.1, 1)
	me.Cam.Rendering.Viewport.SetAbs(8, 8, 64, 64) //SetRel(0.02, 0.04, 0.125, 0.222)
	me.Cam.SetScene(&me.Scene)

	if meshBuf, err = ng.Core.MeshBuffers.Add("buf_quad", ng.Core.MeshBuffers.NewParams(6, 6)); err != nil {
		return
	}
	if quadMesh, err = ng.Core.Libs.Meshes.AddLoad("mesh_quad", ng.MeshProviderPrefabQuad); err != nil {
		return
	}
	if err = meshBuf.Add(quadMesh); err != nil {
		return
	}

	quadMesh.Models.Default().SetMatID("mat_dog")
	me.Dog = me.Scene.RootNode.ChildNodes.AddNew("gui_dog", "mesh_quad", "")
	me.Dog.Transform.SetScale(0.85)
	me.Dog.Transform.Rot.Z = unum.DegToRad(90)

	quadMesh.Models.Default().Clone("model_cat").SetMatID("mat_cat")
	me.Cat = me.Scene.RootNode.ChildNodes.AddNew("gui_cat", "mesh_quad", "model_cat")
	me.Cat.Transform.SetScale(0.85)
	me.Cat.Transform.Rot.Z = unum.DegToRad(90)

	me.Dog.Transform.Pos.Z = 0.1
	me.Cat.Transform.Pos.Z = 0.11
	me.Dog.ApplyTransform()
	me.Cat.ApplyTransform()
	return
}
