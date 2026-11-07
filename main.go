package main

import (
	"math"

	"github.com/Xelckis/motiongo/pkg/anim"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	camera := rl.Camera3D{
		Position: rl.Vector3{
			X: 0.0,
			Y: 2.0,
			Z: -14.0,
		},
		Target:     rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
		Up:         rl.Vector3{Y: 1.0},
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	backgroundColor := rl.DarkGray

	root := anim.NewEngine()

	root.InitWindow(1000, 700, 60, "MotionGo - All Shapes")
	defer root.CloseWindow()

	// 1. Cube
	cube := anim.NewCube(rl.SkyBlue, 2.0, 2.0, 2.0)
	cube.Transform.Position = rl.Vector3{X: -5.0, Y: 3.5, Z: 0.0}

	// 2. Cone
	cone := anim.NewCone(rl.Red, 1.2, 2.5, 16)
	cone.Transform.Position = rl.Vector3{X: 0.0, Y: 3.5, Z: 0.0}

	// 3. Cylinder
	cylinder := anim.NewCylinder(rl.Green, 1.0, 2.5, 16)
	cylinder.Transform.Position = rl.Vector3{X: 5.0, Y: 3.5, Z: 0.0}

	// 4. HemiSphere
	hemiSphere := anim.NewHemiSphere(rl.Gold, 1.3, 16, 16)
	hemiSphere.Transform.Position = rl.Vector3{X: -5.0, Y: 0.0, Z: 0.0}

	// 5. Knot
	knot := anim.NewKnot(rl.Pink, 0.8, 0.3, 32, 16)
	knot.Transform.Position = rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0}

	// 6. Plane
	plane := anim.NewPlane(rl.Purple, 2.5, 2.5, 5, 5)
	plane.Transform.Position = rl.Vector3{X: 5.0, Y: 0.0, Z: 0.0}

	// 7. Poly
	poly := anim.NewPoly(rl.Orange, 6, 1.3)
	poly.Transform.Position = rl.Vector3{X: -5.0, Y: -3.5, Z: 0.0}

	// 8. Sphere
	sphere := anim.NewSphere(rl.Blue, 1.3, 16, 16)
	sphere.Transform.Position = rl.Vector3{X: 0.0, Y: -3.5, Z: 0.0}

	// 9. Torus
	torus := anim.NewTorus(rl.Lime, 1.0, 0.4, 16, 16)
	torus.Transform.Position = rl.Vector3{X: 5.0, Y: -3.5, Z: 0.0}

	shapes := []*anim.ShapeNode{
		cube,
		cone,
		cylinder,
		hemiSphere,
		knot,
		plane,
		poly,
		sphere,
		torus,
	}

	light := anim.LoadLights(camera, 0.1, 0.1, 0.1, 1.0)
	anim.NewLight(anim.LightTypePoint, rl.NewVector3(-2, 1, -2), rl.NewVector3(0, 0, 0), rl.Yellow, light)

	for _, shape := range shapes {
		shape.ApplyShaders(light)
		root.Root.AddChild(shape)
		root.Animate(&shape.Transform.Rotation).
			From(rl.QuaternionIdentity()).
			To(rl.QuaternionFromAxisAngle(rl.Vector3{X: 1, Y: 1, Z: 0}, float32(math.Pi*4))).
			In(0, 400).
			Lerp(func(_, _ rl.Quaternion, progress float32) rl.Quaternion {
				return rl.QuaternionFromAxisAngle(
					rl.Vector3{X: 1, Y: 1, Z: 0},
					progress*float32(2*math.Pi),
				)
			})
	}

	root.Render(backgroundColor, camera)
}
