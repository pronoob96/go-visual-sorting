package main

import "github.com/veandco/go-sdl2/sdl"

func InsertionSort(arr []int32, n int, renderer *sdl.Renderer) {
	for i := 0; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key

		PrintArray(arr, renderer)
	}
}
