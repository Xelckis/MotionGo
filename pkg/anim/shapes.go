package anim

import (
	core "github.com/Xelckis/motiongo/internal/core"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShapeNode struct {
	*core.BaseNode
	Color    rl.Color
	Material rl.Material
	Mesh     rl.Mesh
}

func NewCube(color rl.Color, width, height, length float32) *ShapeNode {
	mat := rl.LoadMaterialDefault()
	mat.Maps.Color = color
	return &ShapeNode{core.NewNode(), color, mat, rl.GenMeshCube(width, height, length)}
}

func NewCone(color rl.Color, radius, height float32, slices int) *ShapeNode {
	mat := rl.LoadMaterialDefault()
	mat.Maps.Color = color
	return &ShapeNode{core.NewNode(), color, mat, rl.GenMeshCone(radius, height, slices)}
}

func NewCylinder(color rl.Color, radius, height float32, slices int) *ShapeNode {
	mat := rl.LoadMaterialDefault()
	mat.Maps.Color = color
	return &ShapeNode{core.NewNode(), color, mat, rl.GenMeshCylinder(radius, height, slices)}
}

func NewHemiSphere(color rl.Color, radius float32, rings, slices int) *ShapeNode {
	mat := rl.LoadMaterialDefault()
	mat.Maps.Color = color
	return &ShapeNode{core.NewNode(), color, mat, rl.GenMeshHemiSphere(radius, rings, slices)}
}

func NewKnot(color rl.Color, radius, size float32, radSeg, sides int) *ShapeNode {
	mat := rl.LoadMaterialDefault()
	mat.Maps.Color = color
	return &ShapeNode{core.NewNode(), color, mat, rl.GenMeshKnot(radius, size, radSeg, sides)}
}

func NewPlane(color rl.Color, width, height float32, resX, resZ int) *ShapeNode {
	mat := rl.LoadMaterialDefault()
	mat.Maps.Color = color
	return &ShapeNode{core.NewNode(), color, mat, rl.GenMeshPlane(width, height, resX, resZ)}
}

func NewPoly(color rl.Color, sides int, radius float32) *ShapeNode {
	mat := rl.LoadMaterialDefault()
	mat.Maps.Color = color
	return &ShapeNode{core.NewNode(), color, mat, rl.GenMeshPoly(sides, radius)}
}

func NewSphere(color rl.Color, radius float32, rings, slices int) *ShapeNode {
	mat := rl.LoadMaterialDefault()
	mat.Maps.Color = color
	return &ShapeNode{core.NewNode(), color, mat, rl.GenMeshSphere(radius, rings, slices)}
}

func NewTorus(color rl.Color, radius, size float32, radSeg, sides int) *ShapeNode {
	mat := rl.LoadMaterialDefault()
	mat.Maps.Color = color
	return &ShapeNode{core.NewNode(), color, mat, rl.GenMeshTorus(radius, size, radSeg, sides)}
}

func (s *ShapeNode) Render() {
	rl.BeginShaderMode(s.Material.Shader)
	rl.DrawMesh(s.Mesh, s.Material, s.GlobalMatrix)
	rl.EndShaderMode()
}
