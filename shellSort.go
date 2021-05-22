package main

import "github.com/veandco/go-sdl2/sdl"

func ShellSort(arr []int32, n int, renderer *sdl.Renderer) {

	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i += 1 {
			temp := arr[i]

			var j int
			for j = i; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
			}

			arr[j] = temp
			PrintArray(arr, renderer)
		}
	}
}
