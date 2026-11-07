# Sistema de Animação

O sistema de animação do MotionGo é declarativo, fluente e baseado em **amostragem por frames**. 

---

## 1. Como Funciona a Linha do Tempo (Timeline)

Ao contrário de motores baseados em delta time contínuo (`dt`), o MotionGo trabalha com uma linha do tempo discreta baseada em **número de frames**:

- A cada iteração do loop, o `Engine.Tick()` incrementa o frame atual (`engine.Frame++`).
- Cada animação (chamada internamente de `Track`) possui um intervalo: `StartFrame` e `EndFrame`.
- O progresso normalizado varia de $0.0$ a $1.0$:
  
  $$\text{progresso} = \frac{\text{frame} - \text{StartFrame}}{\text{EndFrame} - \text{StartFrame}}$$

---

## 2. API Fluente (Builders Específicos)

Para conveniência no dia a dia, a biblioteca fornece métodos auxiliares diretos para transformar nós que implementam `core.Node`:

### A. Mover (`engine.Move`)
Modifica o vetor de posição (`Transform.Position`) do nó:

```go
engine.Move(node).
	To(5.0, 2.0, -1.0). // Posição final (X, Y, Z)
	In(0, 60)           // Dos frames 0 a 60 (1 segundo a 60 FPS)
```

### B. Rotacionar (`engine.Rotate`)
Recebe ângulos em **graus (Euler)** e converte internamente para **Quatérnions** com interpolação esférica (Slerp):

```go
engine.Rotate(node).
	To(0.0, 180.0, 0.0). // Gira 180 graus no eixo Y
	In(60, 120)          // Dos frames 60 a 120
```

### C. Escalar (`engine.Scale`)
Modifica a escala tridimensional (`Transform.Scale`) do nó:

```go
engine.Scale(node).
	To(2.0, 2.0, 2.0).  // Dobra o tamanho do objeto
	In(120, 180)        // Dos frames 120 a 180
```

---

## 3. Animação Genérica (`engine.Animate`)

Quando você precisa animar qualquer variável ou propriedade por ponteiro (incluindo cores, números ou quatérnions customizados), utilize `engine.Animate(&variavel)`:

```go
var opacity float32 = 0.0

engine.Animate(&opacity).
	From(0.0).
	To(1.0).
	In(0, 100)
```

### Tipos Suportados Nativamente pelo `Linear[T]()`:
- `float32` (Interpolação linear numérica)
- `rl.Vector3` (Interpolação linear nos componentes X, Y, Z)
- `rl.Quaternion` (Interpolação esférica `QuaternionSlerp`)
- `rl.Color` (Interpolação linear nos canais RGBA)

---

## 4. Interpolação Personalizada (`.Lerp()`)

Você pode substituir o interpolador padrão passando uma função customizada com a assinatura `anim.Interpolator[T]`:

```go
type Interpolator[T any] func(start, end T, progress float32) T
```

### Exemplo: Rotação contínua customizada

```go
engine.Animate(&shape.Transform.Rotation).
	From(rl.QuaternionIdentity()).
	To(rl.QuaternionFromAxisAngle(rl.Vector3{X: 1, Y: 1, Z: 0}, float32(math.Pi*4))).
	In(0, 400).
	Lerp(func(_, _ rl.Quaternion, progress float32) rl.Quaternion {
		return rl.QuaternionFromAxisAngle(
			rl.Vector3{X: 1, Y: 1, Z: 0},
			progress*float32(2*math.Pi),
		)
	})
```

---

## 5. Encadeamento e Paralelismo de Animações

Como as animações são agendadas por intervalos de frames, você pode executá-las em paralelo ou em sequência apenas definindo seus frames de início e término:

```go
// Animações simultâneas (frames 0 a 60):
engine.Move(node).To(2, 0, 0).In(0, 60)
engine.Rotate(node).To(0, 90, 0).In(0, 60)

// Animação sequencial (frames 60 a 120):
engine.Scale(node).To(1.5, 1.5, 1.5).In(60, 120)
```
