package entities

import (
	core "github.com/Xelckis/motiongo/internal/core"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CubeNode struct {
	*core.BaseNode
	Color    rl.Color
	Material rl.Material
	Mesh     rl.Mesh
}

func NewCube(color rl.Color, width, height, length float32) *CubeNode {
	mat := rl.LoadMaterialDefault()
	mat.Maps.Color = color
	return &CubeNode{core.NewNode(), color, mat, rl.GenMeshCube(width, height, length)}
}

func (c *CubeNode) Render() {
	rl.DrawMesh(c.Mesh, c.Material, c.GlobalMatrix)
}
