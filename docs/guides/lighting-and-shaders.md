# Iluminação e Shaders

O MotionGo integra um pipeline de shaders embutidos via `go:embed` e um sistema de iluminação dinâmica compatível com modelos Phong/Blinn-Phong e efeitos de pós-processamento como renderização em estilo ASCII.

---

## 1. Shaders Embutidos (`assets.ShadersFS`)

Os arquivos de shader GLSL ficam localizados no diretório `assets/shaders/` e são compilados diretamente no binário Go através do pacote `assets`:

- `lighting.vs`: Vertex shader com cálculo de normais e posições no espaço do mundo.
- `lighting.fs`: Fragment shader com cálculo de luz ambiente, difusa e especular.
- `ascii.fs`: Fragment shader para renderização em estilo caractere ASCII.

---

## 2. Configurando Iluminação Dinâmica

Para iluminar sua cena 3D:

### Passo 1: Carregar o Shader de Luzes
Use `anim.LoadLights` passando a câmera da cena e a cor da luz ambiente (valores RGBA normalizados entre 0.0 e 1.0):

```go
// Luz ambiente suave (10% de intensidade)
lightShader := anim.LoadLights(camera, 0.1, 0.1, 0.1, 1.0)
```

### Passo 2: Criar Fontes de Luz
Use `anim.NewLight` especificando o tipo (`LightTypePoint` ou `LightTypeDirectional`), posição, alvo, cor e o shader:

```go
// 1. Luz pontual amarela localizada em (-2, 1, -2) apontando para a origem
anim.NewLight(
	anim.LightTypePoint,
	rl.NewVector3(-2, 1, -2),
	rl.NewVector3(0, 0, 0),
	rl.Yellow,
	lightShader,
)

// 2. Luz direcional branca
anim.NewLight(
	anim.LightTypeDirectional,
	rl.NewVector3(5, 10, 5),
	rl.NewVector3(0, 0, 0),
	rl.White,
	lightShader,
)
```

### Passo 3: Aplicar o Shader às Formas
Atribua o shader aos objetos desejados com `.ApplyShaders()`:

```go
cube := anim.NewCube(rl.SkyBlue, 2.0, 2.0, 2.0)
cube.ApplyShaders(lightShader)
```

---

## 3. Shader de Efeito ASCII (`LoadAsciiShaders`)

Para transformar o visual 3D em arte ASCII baseada em grade de texto:

```go
screenWidth := 1000
screenHeight := 700
fontSize := 12

asciiShader := anim.LoadAsciiShaders(fontSize, screenWidth, screenHeight)

// Aplicar ao objeto
sphere := anim.NewSphere(rl.Green, 2.0, 32, 32)
sphere.ApplyShaders(asciiShader)
```

---

## 4. Exemplo Completo de Cena com Iluminação

```go
package main

import (
	"github.com/Xelckis/motiongo/pkg/anim"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	camera := rl.Camera3D{
		Position:   rl.Vector3{X: 0, Y: 3, Z: -8},
		Target:     rl.Vector3{X: 0, Y: 0, Z: 0},
		Up:         rl.Vector3{Y: 1},
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	engine := anim.NewEngine()
	engine.InitWindow(1000, 700, 60, "Iluminação MotionGo")
	defer engine.CloseWindow()

	// 1. Carregar shader de iluminação com ambiente escuro
	lightShader := anim.LoadLights(camera, 0.15, 0.15, 0.15, 1.0)

	// 2. Adicionar luz pontual
	anim.NewLight(
		anim.LightTypePoint,
		rl.NewVector3(2, 4, -2),
		rl.NewVector3(0, 0, 0),
		rl.White,
		lightShader,
	)

	// 3. Criar formas e aplicar o shader
	torus := anim.NewTorus(rl.Gold, 1.5, 0.5, 32, 32)
	torus.ApplyShaders(lightShader)
	engine.Root.AddChild(torus)

	engine.Rotate(torus).To(360, 360, 0).In(0, 300)

	engine.Render(rl.Black, camera)
}
```
