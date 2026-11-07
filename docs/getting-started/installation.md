# Instalação e Configuração

O **MotionGo** é uma biblioteca de motion graphics e animação 3D programática em Go, construída sobre o [raylib-go](https://github.com/gen2brain/raylib-go).

---

## Requisitos do Sistema

- **Go**: Versão 1.22 ou superior recomendada.
- **Compilador C (CGO)**: Como o `raylib-go` utiliza bindings para a biblioteca gráfica Raylib em C/C++, um compilador C compatível é necessário dependendo da plataforma:
  - **Linux**: `gcc` ou `clang`, além dos pacotes de desenvolvimento do OpenGL e X11 / Wayland:
    ```bash
    # Debian / Ubuntu / Pop!_OS
    sudo apt install build-essential libasound2-dev libmesa-dev libx11-dev libxrandr-dev libxi-dev libxcursor-dev libxinerama-dev libgl1-mesa-dev
    
    # Arch Linux / Manjaro
    sudo pacman -S base-devel mesa libx11 libxrandr libxi libxcursor libxinerama alsa-lib
    
    # Fedora / RHEL
    sudo dnf install @development-tools mesa-libGL-devel libX11-devel libXrandr-devel libXi-devel libXcursor-devel libXinerama-devel alsa-lib-devel
    ```
  - **macOS**: Xcode Command Line Tools (`xcode-select --install`).
  - **Windows**: [MinGW-w64](https://www.mingw-w64.org/) ou [w64devkit](https://github.com/skeeto/w64devkit) configurado no `PATH`.

---

## Adicionando o MotionGo ao seu Projeto

Para inicializar um novo projeto Go e adicionar a dependência:

```bash
mkdir meu-motion
cd meu-motion
go mod init meu-motion
go get github.com/Xelckis/motiongo
```

Verifique se a dependência do Raylib foi baixada corretamente executando:

```bash
go mod tidy
```

---

## Verificação da Instalação

Crie um arquivo `main.go` básico para testar se a janela gráfica inicializa corretamente:

```go
package main

import (
	"github.com/Xelckis/motiongo/pkg/anim"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	engine := anim.NewEngine()
	engine.InitWindow(800, 600, 60, "Teste MotionGo")
	defer engine.CloseWindow()

	camera := rl.Camera3D{
		Position:   rl.Vector3{X: 0, Y: 2, Z: -5},
		Target:     rl.Vector3{X: 0, Y: 0, Z: 0},
		Up:         rl.Vector3{Y: 1},
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	engine.Render(rl.DarkGray, camera)
}
```

Execute o programa:

```bash
go run main.go
```

Se uma janela escura abrir e puder ser fechada normalmente, o ambiente está configurado com sucesso!
