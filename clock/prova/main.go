package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func main() {
	// Create a new image with white background
	width := 400
	height := 400
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Center of the clock
	centerX := float64(width) / 2
	centerY := float64(height) / 2
	radius := 150.0

	// Load a basic font
	face := basicfont.Face7x13

	// Set up the font drawer
	drawer := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: face,
	}

	for hour := 1; hour <= 12; hour++ {
		// Calculate the angle for each hour
		angle := 2 * math.Pi / 12 * float64(hour-3)
		x := centerX + radius*math.Cos(angle)
		y := centerY + radius*math.Sin(angle)

		// Calculate the position for the text (adjust for centering)
		text := fmt.Sprintf("%d", hour)
		textX := x - float64(drawer.MeasureString(text).Ceil())*0.5
		textY := y + float64(face.Metrics().Ascent.Ceil())*0.5

		// Draw the text
		drawer.Dot = fixed.Point26_6{
			X: fixed.I(int(textX)),
			Y: fixed.I(int(textY)),
		}
		drawer.DrawString(text)
	}

	// Save the image to a file
	file, err := os.Create("clock.png")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()
	if err := png.Encode(file, img); err != nil {
		log.Fatalf("failed to encode png: %v", err)
	}

	log.Println("Clock image saved as clock.png")
}
