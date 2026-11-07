# Catálogo de Primitivas 3D (Shapes)

O MotionGo disponibiliza 9 primitivas geométricas 3D prontas para uso. Cada forma é instanciada como um `*anim.ShapeNode`, herdando todas as propriedades de transformação (`Transform`), hierarquia de nós (`BaseNode`) e renderização com suporte a shaders.

---

## Estrutura do `ShapeNode`

```go
type ShapeNode struct {
	*core.BaseNode
	Color    rl.Color
	Material rl.Material
	Mesh     rl.Mesh
}
```

---

## Tabela de Primitivas e Parâmetros

| Função Construtora | Parâmetros | Descrição |
| :--- | :--- | :--- |
| `anim.NewCube` | `color, width, height, length` | Cubo/Paralelepípedo com largura, altura e profundidade. |
| `anim.NewCone` | `color, radius, height, slices` | Cone com raio de base, altura e contagem de fatias angulares. |
| `anim.NewCylinder` | `color, radius, height, slices` | Cilindro com raio, altura e fatias angulares. |
| `anim.NewHemiSphere` | `color, radius, rings, slices` | Meia-esfera com raio, número de anéis e fatias. |
| `anim.NewKnot` | `color, radius, size, radSeg, sides` | Nó tórico (Knot 3D) com raio, espessura e segmentos. |
| `anim.NewPlane` | `color, width, height, resX, resZ` | Plano com largura, comprimento e resolução de subdivisões. |
| `anim.NewPoly` | `color, sides, radius` | Polígono regular extrudado com N lados e raio. |
| `anim.NewSphere` | `color, radius, rings, slices` | Esfera completa com raio, anéis e fatias. |
| `anim.NewTorus` | `color, radius, size, radSeg, sides` | Toro (Donut) com raio do anel, raio do tubo e resoluções. |

---

## Exemplos de Criação

### 1. Cubo
```go
cube := anim.NewCube(rl.SkyBlue, 2.0, 2.0, 2.0)
cube.Transform.Position = rl.Vector3{X: -5.0, Y: 3.5, Z: 0.0}
```

### 2. Cone
```go
cone := anim.NewCone(rl.Red, 1.2, 2.5, 16)
cone.Transform.Position = rl.Vector3{X: 0.0, Y: 3.5, Z: 0.0}
```

### 3. Cilindro
```go
cylinder := anim.NewCylinder(rl.Green, 1.0, 2.5, 16)
cylinder.Transform.Position = rl.Vector3{X: 5.0, Y: 3.5, Z: 0.0}
```

### 4. Meia-Esfera (HemiSphere)
```go
hemiSphere := anim.NewHemiSphere(rl.Gold, 1.3, 16, 16)
hemiSphere.Transform.Position = rl.Vector3{X: -5.0, Y: 0.0, Z: 0.0}
```

### 5. Nó Tórico (Knot)
```go
knot := anim.NewKnot(rl.Pink, 0.8, 0.3, 32, 16)
knot.Transform.Position = rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0}
```

### 6. Plano (Plane)
```go
plane := anim.NewPlane(rl.Purple, 2.5, 2.5, 5, 5)
plane.Transform.Position = rl.Vector3{X: 5.0, Y: 0.0, Z: 0.0}
```

### 7. Polígono Regular (Poly)
```go
// Polígono hexagonal (6 lados)
poly := anim.NewPoly(rl.Orange, 6, 1.3)
poly.Transform.Position = rl.Vector3{X: -5.0, Y: -3.5, Z: 0.0}
```

### 8. Esfera (Sphere)
```go
sphere := anim.NewSphere(rl.Blue, 1.3, 16, 16)
sphere.Transform.Position = rl.Vector3{X: 0.0, Y: -3.5, Z: 0.0}
```

### 9. Toro (Torus)
```go
torus := anim.NewTorus(rl.Lime, 1.0, 0.4, 16, 16)
torus.Transform.Position = rl.Vector3{X: 5.0, Y: -3.5, Z: 0.0}
```

---

## Renderização de Malhas

O método `Render()` do `ShapeNode` ativa o shader do material e desenha a malha geométrica transformada pela matriz global calculada:

```go
func (s *ShapeNode) Render() {
	rl.BeginShaderMode(s.Material.Shader)
	rl.DrawMesh(s.Mesh, s.Material, s.GlobalMatrix)
	rl.EndShaderMode()
}
```
