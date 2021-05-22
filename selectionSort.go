package main

import "github.com/veandco/go-sdl2/sdl"

func SelectionSort(arr []int32, n int, renderer *sdl.Renderer) {
	for i := 0; i < n; i++ {
		min_idx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}

		temp := arr[min_idx]
		arr[min_idx] = arr[i]
		arr[i] = temp

		PrintArray(arr, renderer)
	}
}
