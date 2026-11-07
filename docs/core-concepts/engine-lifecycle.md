# Ciclo de Vida da Engine (Engine Lifecycle)

O motor do MotionGo encapsula a inicialização da janela, o avanço da timeline e o loop de renderização 3D.

---

## 1. Estrutura da Engine

A struct `Engine` (em `pkg/anim/core.go`) centraliza o estado da cena:

```go
type Engine struct {
	Root   *core.BaseNode      // Raiz da árvore hierárquica de nós
	Tracks []func(frame int)  // Funções atualizadoras de animação por frame
	Frame  int                 // Contador de frames atual
}
```

---

## 2. Inicialização e Encerramento

A inicialização e a finalização da janela gráfica utilizam wrappers diretos sobre a Raylib:

```go
// 1. Criar a instância do motor
engine := anim.NewEngine()

// 2. Abrir a janela e definir a taxa de quadros (FPS)
engine.InitWindow(1280, 720, 60, "Minha Animação 3D")

// 3. Garantir o fechamento correto dos recursos da GPU ao sair
defer engine.CloseWindow()
```

---

## 3. O Loop Principal (`Tick` e `Render`)

A execução do render loop é acionada por `engine.Render(backgroundColor, camera)`:

```go
func (e *Engine) Render(backgroundColor color.RGBA, camera rl.Camera3D) {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)

		rl.BeginMode3D(camera)

		e.Tick()        // 1. Atualiza lógica e transformações
		e.Root.Render() // 2. Desenha todos os nós da árvore

		rl.EndMode3D()
		rl.EndDrawing()
	}
}
```

### O que acontece a cada frame no `Tick()`:

```go
func (e *Engine) Tick() {
	e.Frame++ // Avança a timeline

	// Executa as animações cadastradas para o frame atual
	for _, atualizarTrack := range e.Tracks {
		atualizarTrack(e.Frame)
	}

	// Recalcula matrizes locais e globais em cascata por toda a árvore
	e.Root.Update(rl.MatrixIdentity())
}
```

---

## 4. Diagrama de Execução por Frame

```text
+-------------------------------------------------------------+
|                     Frame N Loop                           |
+-------------------------------------------------------------+
                              |
                     [ engine.Tick() ]
                              |
            +-----------------+-----------------+
            |                                   |
    1. engine.Frame++                3. e.Root.Update()
            |                                   |
    2. Executa Tracks             Recalcula LocalMatrix & GlobalMatrix
       (Atualiza Transforms)           na hierarquia de nós
            |                                   |
            +-----------------+-----------------+
                              |
                    [ e.Root.Render() ]
                              |
                  Desenha Meshes com Shaders
                              |
                  [ rl.EndDrawing() / VSync ]
```
