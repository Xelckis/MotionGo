# Grafo de Cena (Scene Graph)

O MotionGo organiza todos os elementos renderizáveis e transformáveis em uma estrutura em árvore hierárquica conhecida como **Grafo de Cena (Scene Graph)**.

---

## 1. Nós (`core.Node` e `core.BaseNode`)

A interface fundamental é `core.Node` (definida em `internal/core/node.go`):

```go
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
```

A estrutura `BaseNode` fornece a implementação padrão dessa interface, gerenciando:
- Lista de nós filhos (`Children []Node`).
- Dados de transformação espacial (`Transform *Transform`).
- Matriz local (`LocalMatrix rl.Matrix`) e matriz global combinada (`GlobalMatrix rl.Matrix`).

---

## 2. Transformações Espaciais (`Transform`)

Cada nó possui uma struct `Transform`:

```go
type Transform struct {
	Position rl.Vector3     // Posição espacial (X, Y, Z)
	Rotation rl.Quaternion  // Rotação via Quatérnions (X, Y, Z, W)
	Scale    rl.Vector3     // Escala nos eixos (X, Y, Z)
}
```

Valores padrão ao instanciar `core.NewTransform()`:
- `Position`: `{0, 0, 0}`
- `Rotation`: `{0, 0, 0, 1}` (Quatérnion identidade)
- `Scale`: `{1, 1, 1}`

---

## 3. Cálculo de Matrizes (TRS)

Em computação gráfica 3D, a matriz local de um nó é calculada combinando Translação, Rotação e Escala na seguinte ordem:

$$\text{LocalMatrix} = \text{Translation} \times (\text{Rotation} \times \text{Scale})$$

No código do `BaseNode`:
```go
func (n *BaseNode) calculateLocalMatrix() {
	matPos := rl.MatrixTranslate(n.Transform.Position.X, n.Transform.Position.Y, n.Transform.Position.Z)
	matRot := rl.QuaternionToMatrix(n.Transform.Rotation)
	matScale := rl.MatrixScale(n.Transform.Scale.X, n.Transform.Scale.Y, n.Transform.Scale.Z)

	n.LocalMatrix = rl.MatrixMultiply(matPos, rl.MatrixMultiply(matRot, matScale))
}
```

A **Matriz Global** propaga a transformação do nó pai para todos os nós filhos:
```go
func (n *BaseNode) Update(globalMatrix rl.Matrix) {
	n.calculateLocalMatrix()
	n.GlobalMatrix = rl.MatrixMultiply(globalMatrix, n.LocalMatrix)

	for _, child := range n.Children {
		child.Update(n.GlobalMatrix)
	}
}
```

---

## 4. Hierarquia Pai-Filho na Prática

Quando um nó filho é anexado a um nó pai usando `AddChild`, todas as transformações de posição, rotação e escala aplicadas ao pai afetam automaticamente os filhos.

### Exemplo: Sistema Solar (Planeta e Lua)

```go
package main

import (
	"github.com/Xelckis/motiongo/pkg/anim"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	engine := anim.NewEngine()
	engine.InitWindow(800, 600, 60, "Hierarquia de Nós")
	defer engine.CloseWindow()

	camera := rl.Camera3D{
		Position:   rl.Vector3{X: 0, Y: 5, Z: -15},
		Target:     rl.Vector3{X: 0, Y: 0, Z: 0},
		Up:         rl.Vector3{Y: 1},
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	// 1. Criar o planeta (nó pai)
	planet := anim.NewSphere(rl.Blue, 1.5, 16, 16)
	planet.Transform.Position = rl.Vector3{X: 0, Y: 0, Z: 0}

	// 2. Criar a lua (nó filho) com deslocamento local relativo ao planeta
	moon := anim.NewSphere(rl.LightGray, 0.5, 16, 16)
	moon.Transform.Position = rl.Vector3{X: 3.5, Y: 0, Z: 0}

	// Anexar a lua como filha do planeta
	planet.AddChild(moon)

	// Anexar o planeta à raiz da cena
	engine.Root.AddChild(planet)

	// 3. Ao rotacionar o planeta, a lua orbitará automaticamente em torno dele!
	engine.Rotate(planet).
		To(0, 360, 0).
		In(0, 240)

	engine.Render(rl.Black, camera)
}
```
