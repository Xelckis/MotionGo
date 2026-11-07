# Interpolações e Easing Personalizados

O MotionGo permite estender ou substituir a interpolação linear padrão (`Linear[T]()`) com qualquer função de curvatura temporal (*easing functions*).

---

## 1. A Assinatura de um `Interpolator[T]`

Qualquer função que corresponda à assinatura genérica `anim.Interpolator[T]` pode ser passada para `.Lerp(...)`:

```go
type Interpolator[T any] func(start, end T, progress float32) T
```

- `start`: Valor no frame inicial (`StartFrame`).
- `end`: Valor no frame final (`EndFrame`).
- `progress`: Valor normalizado entre `0.0` (início) e `1.0` (fim).

---

## 2. Criando Funções de Easing (Curvas de Animação)

As funções de easing transformam a taxa de variação do `progress` linear em aceleração ou desaceleração suave:

```go
// Ease In Quad (Aceleração quadrática)
func EaseInQuad(t float32) float32 {
	return t * t
}

// Ease Out Quad (Desaceleração quadrática)
func EaseOutQuad(t float32) float32 {
	return t * (2 - t)
}

// Ease In-Out Quad (Aceleração no início e desaceleração no final)
func EaseInOutQuad(t float32) float32 {
	if t < 0.5 {
		return 2 * t * t
	}
	return -1 + (4-2*t)*t
}
```

---

## 3. Aplicando Easing em `Vector3`

Podemos combinar a função de easing com interpolação vetorial:

```go
func EaseInOutVector3(start, end rl.Vector3, progress float32) rl.Vector3 {
	// 1. Aplica a curva no progresso
	easedProgress := EaseInOutQuad(progress)

	// 2. Interpola linearmente usando o progresso curvado
	return rl.Vector3{
		X: start.X + (end.X-start.X)*easedProgress,
		Y: start.Y + (end.Y-start.Y)*easedProgress,
		Z: start.Z + (end.Z-start.Z)*easedProgress,
	}
}
```

### Uso no Builder:
```go
engine.Animate(shape.Pos()).
	From(rl.Vector3{X: -5, Y: 0, Z: 0}).
	To(rl.Vector3{X: 5, Y: 0, Z: 0}).
	In(0, 120).
	Lerp(EaseInOutVector3)
```

---

## 4. Easing para Cores (`rl.Color`)

```go
func SmoothColorLerp(start, end rl.Color, progress float32) rl.Color {
	p := EaseInOutQuad(progress)
	return rl.Color{
		R: uint8(float32(start.R) + (float32(end.R)-float32(start.R))*p),
		G: uint8(float32(start.G) + (float32(end.G)-float32(start.G))*p),
		B: uint8(float32(start.B) + (float32(end.B)-float32(start.B))*p),
		A: uint8(float32(start.A) + (float32(end.A)-float32(start.A))*p),
	}
}
```

---

## 5. Exemplo Completo de Movimento Suave com Easing

```go
package main

import (
	"github.com/Xelckis/motiongo/pkg/anim"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func EaseOutBounce(t float32) float32 {
	if t < (1 / 2.75) {
		return 7.5625 * t * t
	} else if t < (2 / 2.75) {
		t -= (1.5 / 2.75)
		return 7.5625*t*t + 0.75
	} else if t < (2.5 / 2.75) {
		t -= (2.25 / 2.75)
		return 7.5625*t*t + 0.9375
	}
	t -= (2.625 / 2.75)
	return 7.5625*t*t + 0.984375
}

func BounceVector3(start, end rl.Vector3, progress float32) rl.Vector3 {
	p := EaseOutBounce(progress)
	return rl.Vector3{
		X: start.X + (end.X-start.X)*p,
		Y: start.Y + (end.Y-start.Y)*p,
		Z: start.Z + (end.Z-start.Z)*p,
	}
}

func main() {
	camera := rl.Camera3D{
		Position:   rl.Vector3{X: 0, Y: 4, Z: -12},
		Target:     rl.Vector3{X: 0, Y: 0, Z: 0},
		Up:         rl.Vector3{Y: 1},
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	engine := anim.NewEngine()
	engine.InitWindow(800, 600, 60, "Animação com Bounce Easing")
	defer engine.CloseWindow()

	sphere := anim.NewSphere(rl.Orange, 1.2, 32, 32)
	sphere.Transform.Position = rl.Vector3{X: 0, Y: 6, Z: 0}
	engine.Root.AddChild(sphere)

	// Animar queda com efeito bounce (quicando no chão)
	engine.Animate(sphere.Pos()).
		From(rl.Vector3{X: 0, Y: 6, Z: 0}).
		To(rl.Vector3{X: 0, Y: -2, Z: 0}).
		In(0, 120).
		Lerp(BounceVector3)

	engine.Render(rl.DarkGray, camera)
}
```
