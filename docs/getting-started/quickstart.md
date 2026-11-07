# Guia Rápido (Quickstart)

Neste guia, você criará sua primeira animação 3D com o **MotionGo**: um cubo girando e se movendo pelo espaço.

---

## 1. Estrutura do Programa

O fluxo básico de qualquer aplicação com MotionGo é:

1. **Configurar a Câmera 3D**: Definir posição, alvo e perspectiva.
2. **Instanciar o Engine**: Criar o contexto do motor de animação.
3. **Criar Nós e Formas**: Instanciar primitivas 3D (`ShapeNode`).
4. **Agendar Animações**: Definir transformações ao longo dos frames da timeline.
5. **Iniciar o Render Loop**: Abrir a janela e rodar a cena.

---

## 2. Código Completo

Crie um arquivo `main.go`:

```go
package main

import (
	"github.com/Xelckis/motiongo/pkg/anim"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// 1. Configurar câmera
	camera := rl.Camera3D{
		Position:   rl.Vector3{X: 0.0, Y: 3.0, Z: -10.0},
		Target:     rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
		Up:         rl.Vector3{Y: 1.0},
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	// 2. Inicializar o Engine do MotionGo
	engine := anim.NewEngine()
	engine.InitWindow(1000, 700, 60, "MotionGo - Quickstart")
	defer engine.CloseWindow()

	// 3. Criar uma forma 3D (Cubo azul céu com dimensões 2x2x2)
	cube := anim.NewCube(rl.SkyBlue, 2.0, 2.0, 2.0)
	cube.Transform.Position = rl.Vector3{X: -3.0, Y: 0.0, Z: 0.0}

	// Adicionar o cubo à raiz da cena
	engine.Root.AddChild(cube)

	// 4. Animar posição: mover de X=-3 para X=3 entre os frames 0 e 120 (2 segundos a 60 FPS)
	engine.Move(cube).
		To(3.0, 0.0, 0.0).
		In(0, 120)

	// Animar rotação: rotacionar 360 graus no eixo Y entre os frames 0 e 180
	engine.Rotate(cube).
		To(0.0, 360.0, 0.0).
		In(0, 180)

	// 5. Iniciar o render loop
	engine.Render(rl.DarkGray, camera)
}
```

---

## 3. Explicação Passo a Passo

### A. Criação de Formas
```go
cube := anim.NewCube(rl.SkyBlue, 2.0, 2.0, 2.0)
```
O método `anim.NewCube` instancia um `*anim.ShapeNode` contendo a malha (mesh) geométrica e o material com a cor especificada.

### B. Inserção no Grafo de Cena
```go
engine.Root.AddChild(cube)
```
Todos os elementos renderizáveis devem ser filhos diretos ou indiretos de `engine.Root`.

### C. Animações Fluentes
```go
engine.Move(cube).To(3.0, 0.0, 0.0).In(0, 120)
```
O MotionGo utiliza uma API fluente baseada em builders. 
- `.To(x, y, z)` especifica o valor final da transformação.
- `.In(startFrame, endFrame)` define em qual intervalo da timeline a interpolação linear acontecerá.

### D. Execução
```bash
go run main.go
```
Você verá o cubo azul se movimentar suavemente e rotacionar no espaço 3D.
