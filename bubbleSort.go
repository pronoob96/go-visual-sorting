package main

import "github.com/veandco/go-sdl2/sdl"

func BubbleSort(arr []int32, n int, renderer *sdl.Renderer) {
	for i := 0; i < n; i++ {

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}

		PrintArray(arr, renderer)
	}
}
