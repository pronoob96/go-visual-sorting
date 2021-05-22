package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 1366
	screenHeight = 768
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Sorting",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_FULLSCREEN_DESKTOP)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	arr := make([]int32, screenWidth)

	for i := 0; i < screenWidth; i++ {
		arr[i] = rand.Int31n(screenHeight)
	}

	SelectionSort(arr, screenWidth, renderer)

	for i := 0; i < screenWidth; i++ {
		arr[i] = rand.Int31n(screenHeight)
	}

	InsertionSort(arr, screenWidth, renderer)

	for i := 0; i < screenWidth; i++ {
		arr[i] = rand.Int31n(screenHeight)
	}

	BubbleSort(arr, screenWidth, renderer)

	for i := 0; i < screenWidth; i++ {
		arr[i] = rand.Int31n(screenHeight)
	}

	MergeSort(arr, 0, screenWidth-1, renderer)

	for i := 0; i < screenWidth; i++ {
		arr[i] = rand.Int31n(screenHeight)
	}

	HeapSort(arr, screenWidth, renderer)

	for i := 0; i < screenWidth; i++ {
		arr[i] = rand.Int31n(screenHeight)
	}

	RadixSort(arr, screenWidth, renderer)

	for i := 0; i < screenWidth; i++ {
		arr[i] = rand.Int31n(screenHeight)
	}

	CocktailSort(arr, renderer)

	for i := 0; i < screenWidth; i++ {
		arr[i] = rand.Int31n(screenHeight)
	}

	ShellSort(arr, screenWidth, renderer)

}

func PrintArray(arr []int32, renderer *sdl.Renderer) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			os.Exit(1)
		}
	}
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	renderer.SetDrawColor(255, 255, 255, 255)
	for i := 0; i < screenWidth; i++ {
		renderer.DrawLine(int32(i), screenHeight, int32(i), screenHeight-arr[i])
	}
	renderer.Present()
	//time.Sleep(15 * time.Millisecond)
}
