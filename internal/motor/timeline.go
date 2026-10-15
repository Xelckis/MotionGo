package motor

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Interpolator[T any] func(start, end T, progress float32) T

func Linear[T any]() Interpolator[T] {
	return func(start, end T, progress float32) T {
		var startAny any = start
		var endAny any = end

		switch s := startAny.(type) {
		case float32:
			e := endAny.(float32)
			return any(s + (e-s)*progress).(T)

		case rl.Vector3:
			e := endAny.(rl.Vector3)
			return any(rl.Vector3{
				X: s.X + (e.X-s.X)*progress,
				Y: s.Y + (e.Y-s.Y)*progress,
				Z: s.Z + (e.Z-s.Z)*progress,
			}).(T)

		case rl.Quaternion:
			e := endAny.(rl.Quaternion)
			return any(rl.QuaternionSlerp(s, e, progress)).(T)

		case rl.Color:
			e := endAny.(rl.Color)
			return any(rl.Color{
				R: uint8(float32(s.R) + (float32(e.R)-float32(s.R))*progress),
				G: uint8(float32(s.G) + (float32(e.G)-float32(s.G))*progress),
				B: uint8(float32(s.B) + (float32(e.B)-float32(s.B))*progress),
				A: uint8(float32(s.A) + (float32(e.A)-float32(s.A))*progress),
			}).(T)
		}

		return end
	}
}

type Track[T any] struct {
	Target     *T
	StartVal   T
	EndVal     T
	StartFrame int
	EndFrame   int
	LerpFunc   Interpolator[T]
}

type Engine struct {
	Root   *BaseNode
	Tracks []func(frame int)
	Frame  int
}

func NewEngine() *Engine {
	return &Engine{NewNode(), []func(frame int){}, 0}

}

type AnimBuilder[T any] struct {
	engine *Engine
	track  *Track[T]
}

func (e *Engine) Animate[T any](target *T) *AnimBuilder[T] {
	t := &Track[T]{Target: target}

	updater := func(frame int) {
		if frame >= t.StartFrame && frame <= t.EndFrame {

			duracao := t.EndFrame - t.StartFrame
			if duracao == 0 {
				*t.Target = t.EndVal
				return
			}

			progresso := float32(frame-t.StartFrame) / float32(duracao)

			if t.LerpFunc != nil {
				*t.Target = t.LerpFunc(t.StartVal, t.EndVal, progresso)
			}
		}
	}

	e.Tracks = append(e.Tracks, updater)

	return &AnimBuilder[T]{engine: e, track: t}
}

func (b *AnimBuilder[T]) From(v T) *AnimBuilder[T] { b.track.StartVal = v; return b }

func (b *AnimBuilder[T]) To(v T) *AnimBuilder[T] { b.track.EndVal = v; return b }

func (b *AnimBuilder[T]) In(start, end int) *AnimBuilder[T] {
	b.track.StartFrame = start
	b.track.EndFrame = end

	return b
}

func (b *AnimBuilder[T]) Lerp(f Interpolator[T]) *AnimBuilder[T] { b.track.LerpFunc = f; return b }

func (e *Engine) Tick() {
	e.Frame++

	for _, atualizarTrack := range e.Tracks {
		atualizarTrack(e.Frame)
	}

	e.Root.Update(rl.MatrixIdentity())
}

func (e *Engine) Render(backgroundColor color.RGBA, camera rl.Camera3D) {

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)

		rl.BeginMode3D(camera)

		e.Tick()

		e.Root.Render()

		rl.EndMode3D()

		rl.EndDrawing()
	}
}

func (e *Engine) InitWindow(width, height, fps int32, title string) {
	rl.InitWindow(width, height, title)

	rl.SetTargetFPS(fps)
}

func (e *Engine) CloseWindow() {
	rl.CloseWindow()
}
