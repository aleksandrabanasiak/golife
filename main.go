package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
	"time"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Wziuuum",
		Bounds: pixel.R(0, 0, 400, 400),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Aliceblue)

	grid := gridInit(20)
	grid.drawGrid(win)

	for !win.Closed() {
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			fmt.Print(win.MousePosition())
		}
		win.Update()
	}
}

func (c cell) drawCell(win *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Color = randomColor()
	imd.Push(pixel.V(c.x, c.y))
	imd.Push(pixel.V(c.x+20, c.y+20))
	imd.Rectangle(0)

	imd.Draw(win)
}

func (g grid) drawGrid(win *pixelgl.Window) {
	for row := range g {
		for col := range g[row] {
			if !g[row][col].alive {
				g[row][col].drawCell(win)
			}
		}
	}
}

func gridInit(size int64) grid {
	grid := make([][]cell, size)
	for i := range grid {
		grid[i] = make([]cell, size)
		for j := range grid[i] {
			grid[i][j].x = float64(i * 20)
			grid[i][j].y = float64(j * 20)
			grid[i][j].alive = false
		}
	}

	return grid
}

func randomColor() color.RGBA {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 255

	return color.RGBA{
		R: uint8(min + rand.Intn(max-min+1)),
		B: uint8(min + rand.Intn(max-min+1)),
		G: uint8(min + rand.Intn(max-min+1)),
		A: 200,
	}
}

func main() {
	pixelgl.Run(run)
}

type grid [][]cell

type cell struct {
	x     float64
	y     float64
	alive bool
}
