package motor

import rl "github.com/gen2brain/raylib-go/raylib"

type Node interface {
	calculateLocalMatrix()
	Update(globalMatrix rl.Matrix)
	AddChild(Node)
	Render()
	GetGlobalMatrix() rl.Matrix
}
type transform struct {
	Position rl.Vector3
	Rotation rl.Quaternion
	Scale    rl.Vector3
}

type BaseNode struct {
	LocalMatrix  rl.Matrix
	GlobalMatrix rl.Matrix
	Children     []Node
	Transform    *transform
}

func (n *BaseNode) GetGlobalMatrix() rl.Matrix {
	return n.GlobalMatrix
}

func NewNode() *BaseNode {
	return &BaseNode{
		LocalMatrix:  rl.MatrixIdentity(),
		GlobalMatrix: rl.MatrixIdentity(),
		Transform:    NewTransform(),
	}
}

func (n *BaseNode) AddChild(child Node) {
	n.Children = append(n.Children, child)
}

func (n *BaseNode) Render() {
	for _, child := range n.Children {
		child.Render()
	}
}

func NewTransform() *transform {
	return &transform{
		Position: rl.Vector3{0, 0, 0},
		Rotation: rl.Quaternion{0, 0, 0, 1},
		Scale:    rl.Vector3{1, 1, 1},
	}
}

type CubeNode struct {
	*BaseNode
	Color    rl.Color
	Material rl.Material
	Mesh     rl.Mesh
}

func NewCube(color rl.Color, width, height, length float32) *CubeNode {
	mat := rl.LoadMaterialDefault()
	mat.Maps.Color = color
	return &CubeNode{NewNode(), color, mat, rl.GenMeshCube(width, height, length)}
}

func (c *CubeNode) Render() {
	rl.DrawMesh(c.Mesh, c.Material, c.GlobalMatrix)
}

func (n *BaseNode) calculateLocalMatrix() {
	matPos := rl.MatrixTranslate(n.Transform.Position.X, n.Transform.Position.Y, n.Transform.Position.Z)
	matRot := rl.QuaternionToMatrix(n.Transform.Rotation)
	matScale := rl.MatrixScale(n.Transform.Scale.X, n.Transform.Scale.Y, n.Transform.Scale.Z)

	// Position x Rotation x Scale
	n.LocalMatrix = rl.MatrixMultiply(matPos, rl.MatrixMultiply(matRot, matScale))
}
func (n *BaseNode) Update(globalMatrix rl.Matrix) {
	n.calculateLocalMatrix()

	n.GlobalMatrix = rl.MatrixMultiply(globalMatrix, n.LocalMatrix)

	for _, child := range n.Children {
		child.Update(n.GlobalMatrix)
	}

}
