package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var screenWidth int32 = 1280
var screenHeight int32 = 720
var isStartup bool = true

var points []rl.Vector2


func main() {
	//rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(screenWidth, screenHeight, "DEBi")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	//Draw startup message on screen
	//When user input starts remove text


	for !rl.WindowShouldClose() {

		
		inputHandler()
		mouseInputHandler()

		rl.BeginDrawing()
		drawStrokes()
		displayStartupMessage(isStartup)
		rl.ClearBackground(rl.DarkGray)
		rl.EndDrawing()
	}
}

func inputHandler() {
	//Check full screen
	if rl.IsKeyPressed(rl.KeyF) && rl.IsKeyDown(rl.KeyLeftAlt) || rl.IsKeyDown(rl.KeyRightAlt) {
		display := rl.GetCurrentMonitor()
		if rl.IsWindowFullscreen() {
			rl.SetWindowSize(int(screenWidth), int(screenHeight))
		} else {
			rl.SetWindowSize(rl.GetMonitorWidth(display), rl.GetMonitorHeight(display))
		}
		rl.ToggleBorderlessWindowed()
	}
	if rl.IsKeyPressed(rl.KeyC) && rl.IsKeyDown(rl.KeyLeftShift) {
		isStartup = false;
	}
	if rl.IsKeyPressed(rl.KeyP) && rl.IsKeyDown(rl.KeyLeftShift) {
		rl.SetTargetFPS(240)

	}
}

func mouseInputHandler() {
	// set to one input at a time only
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		points = append(points, rl.GetMousePosition())	
	}
	
	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		fmt.Println(points)
		//add points to stroke list
	}

	if rl.IsMouseButtonDown(rl.MouseButtonRight) {
		fmt.Println("right button pressed")

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

func drawStrokes() {
	var tmpP rl.Vector2 
	tmpP.X = 0
	tmpP.Y = 0
	for _, p := range points {
		rl.DrawPixelV(p, rl.Purple)
		if tmpP.X != 0 && tmpP.Y != 0{
			rl.DrawLineV(tmpP, p, rl.Purple)
		} 
		tmpP = p
	}
}
