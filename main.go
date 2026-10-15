package main

import (
	"math"

	"github.com/Xelckis/motiongo/internal/motor"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	camera := rl.Camera3D{
		Position: rl.Vector3{
			X: 6.0,
			Y: 6.0,
			Z: -10.0,
		},
		Target:     rl.Vector3{},
		Up:         rl.Vector3{Y: 1.0},
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	backgroundColor := rl.Black

	root := motor.NewEngine()

	root.InitWindow(800, 500, 60, "test")
	defer root.CloseWindow()

	cube := motor.NewCube(rl.SkyBlue, 2, 2, 2)
	root.Root.AddChild(cube)

	root.Animate(&cube.Transform.Rotation).From(rl.QuaternionIdentity()).To(rl.QuaternionFromAxisAngle(rl.Vector3{Y: 1}, float32(math.Pi))).In(0, 100).Lerp(motor.Linear[rl.Quaternion]())

	root.Render(backgroundColor, camera)
}
