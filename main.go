package main

import (
	m "math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)
	rl.InitWindow(screenWidth, screenHeight, "Simple Model .obj Viwer")
	rl.SetTargetFPS(60)

	//you can change for your desired file.
	model := rl.LoadModel("donut.obj")

	modelScale := float32(0.1)

	camera := rl.Camera3D{
		Position:   rl.NewVector3(0.0, 0.0, 4.0),
		Target:     rl.NewVector3(0.0, 0.5, 0.0),
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyW) {
			camera.Position.Z -= 0.2
		}
		if rl.IsKeyDown(rl.KeyS) {
			camera.Position.Z += 0.2
		}

		angle := float32(rl.GetTime())

		modelRotation := rl.MatrixMultiply(
			rl.MatrixRotateX(angle*m.Pi/4),
			rl.MatrixRotateY(angle*m.Pi/4),
		)
		model.Transform = rl.MatrixMultiply(modelRotation, rl.MatrixScale(modelScale, modelScale, modelScale))

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)

		pinkColor := rl.NewColor(255, 105, 180, 255)

		rl.DrawModel(model, rl.NewVector3(0.0, 0.0, 0.0), 1.0, pinkColor)

		rl.EndMode3D()
		rl.DrawText("Press W/S for Zoom In/Out", 10, 20, 20, rl.Black)

		rl.EndDrawing()
	}

	rl.UnloadModel(model)
	rl.CloseWindow()
}
