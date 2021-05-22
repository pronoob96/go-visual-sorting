package main

import "github.com/veandco/go-sdl2/sdl"

const (
	ex = 3
)

func getMax(arr []int32, n int) int32 {
	mx := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > mx {
			mx = arr[i]
		}
	}
	return mx
}

func countSort(arr []int32, n int, exp int32, renderer *sdl.Renderer) {
	output := make([]int32, n)
	i := 0
	count := make([]int32, ex)

	for i = 0; i < n; i++ {
		count[(arr[i]/exp)%ex]++
	}

	for i = 1; i < ex; i++ {
		count[i] += count[i-1]
	}

	for i = n - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%ex]-1] = arr[i]
		count[(arr[i]/exp)%ex]--
	}

	for i = 0; i < n; i++ {
		arr[i] = output[i]
		PrintArray(arr, renderer)
	}
}

func RadixSort(arr []int32, n int, renderer *sdl.Renderer) {
	m := getMax(arr, n)

	for exp := int32(1); m/exp > 0; exp *= ex {
		countSort(arr, n, exp, renderer)
	}

}
