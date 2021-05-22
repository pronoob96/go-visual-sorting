package main

import "github.com/veandco/go-sdl2/sdl"

func HeapSort(arr []int32, n int, renderer *sdl.Renderer) {

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i, renderer)
	}

	for i := n - 1; i > 0; i-- {
		temp := arr[0]
		arr[0] = arr[i]
		arr[i] = temp

		heapify(arr, i, 0, renderer)
	}
}

func heapify(arr []int32, n, i int, renderer *sdl.Renderer) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		swap := arr[i]
		arr[i] = arr[largest]
		arr[largest] = swap
		PrintArray(arr, renderer)

		heapify(arr, n, largest, renderer)
	}
}
