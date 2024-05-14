package main

import (
	_ "fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var screenWidth int32 = 1280
var screenHeight int32 = 720
var isStartup bool = true
var prePos rl.Vector2
var penScale float32
var penColor rl.Color

func main() {
	prePos.X = 0
	prePos.Y = 0
	penColor = rl.Purple
	penScale = 3
	defer rl.CloseWindow()

	//rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(screenWidth, screenHeight, "DEBi")
	canvas := rl.LoadRenderTexture(screenWidth, screenHeight)
	rl.BeginTextureMode(canvas)
	rl.ClearBackground(rl.DarkGray)
	rl.EndTextureMode()

	rl.SetTargetFPS(240)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		displayStartupMessage(isStartup)
		inputHandler()
		mouseInputHandler(canvas)
		rl.ClearBackground(rl.DarkGray)
		rl.EndDrawing()

	}
}

func inputHandler() {
	//Check full screen
	//Changeing window size will require a canvase clear
	// if rl.IsKeyPressed(rl.KeyF) && rl.IsKeyDown(rl.KeyLeftAlt) || rl.IsKeyDown(rl.KeyRightAlt) {
	// 	display := rl.GetCurrentMonitor()
	// 	if rl.IsWindowFullscreen() {
	// 		rl.SetWindowSize(int(screenWidth), int(screenHeight))
	//		rl.ClearBackground(rl.DarkGray)
	// 	} else {
	// 		rl.SetWindowSize(rl.GetMonitorWidth(display), rl.GetMonitorHeight(display))
	//		rl.ClearBackground(rl.DarkGray)
	//}
	// 	rl.ToggleBorderlessWindowed()
	// }

	//Clear screen
	if rl.IsKeyPressed(rl.KeyC) && rl.IsKeyDown(rl.KeyLeftShift) {
		isStartup = false
		rl.ClearBackground(rl.DarkGray)
	}
	//Performance mode - high FPS - change to what target you want.
	if rl.IsKeyPressed(rl.KeyP) && rl.IsKeyDown(rl.KeyLeftShift) {
		rl.SetTargetFPS(240)
	}
	//Color change - Switch would look be nice
	if rl.IsKeyPressed(rl.KeyA) {
		penColor = rl.SkyBlue
	}
	if rl.IsKeyPressed(rl.KeyS) {
		penColor = rl.Yellow
	}
	if rl.IsKeyPressed(rl.KeyD) {
		penColor = rl.Purple
	}
	if rl.IsKeyPressed(rl.KeyF) {
		penColor = rl.Green
	}
}

func mouseInputHandler(c rl.RenderTexture2D) {
	// set to one input at a time only
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		rl.BeginTextureMode(c)
		rl.DrawCircleV(rl.GetMousePosition(), 10, penColor)
		rl.EndTextureMode()
	}

	if rl.IsMouseButtonDown(rl.MouseButtonRight) {
		rl.BeginTextureMode(c)
		rl.DrawCircleV(rl.GetMousePosition(), 10, rl.DarkGray)
		prePos = rl.GetMousePosition()
		rl.EndTextureMode()
	}
}

func displayStartupMessage(b bool) {
	if b {
		rl.DrawText("DEBi v0.1!", screenWidth/2-
			rl.MeasureText("DEBi v0.1!", 10), screenHeight/2-12, 20, rl.RayWhite)
		rl.DrawText("Draw on the screen", screenWidth/2-
			rl.MeasureText("Draw on the screen", 10), screenHeight/2+12, 20, rl.RayWhite)

	}

}
