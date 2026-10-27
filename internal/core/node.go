package core

import rl "github.com/gen2brain/raylib-go/raylib"

type Node interface {
	calculateLocalMatrix()
	Update(globalMatrix rl.Matrix)
	AddChild(Node)
	Render()
	GetGlobalMatrix() rl.Matrix

	Pos() *rl.Vector3
	Rot() *rl.Quaternion
	Scl() *rl.Vector3
}
type Transform struct {
	Position rl.Vector3
	Rotation rl.Quaternion
	Scale    rl.Vector3
}

type BaseNode struct {
	LocalMatrix  rl.Matrix
	GlobalMatrix rl.Matrix
	Children     []Node
	Transform    *Transform
}

func (n *BaseNode) Pos() *rl.Vector3 {
	return &n.Transform.Position
}

func (n *BaseNode) Rot() *rl.Quaternion {
	return &n.Transform.Rotation
}

func (n *BaseNode) Scl() *rl.Vector3 {
	return &n.Transform.Scale
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

func NewTransform() *Transform {
	return &Transform{
		Position: rl.Vector3{0, 0, 0},
		Rotation: rl.Quaternion{0, 0, 0, 1},
		Scale:    rl.Vector3{1, 1, 1},
	}
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
