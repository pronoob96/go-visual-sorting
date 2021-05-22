package main

import "github.com/veandco/go-sdl2/sdl"

func CocktailSort(a []int32, renderer *sdl.Renderer) {
	swapped := true
	start := 0
	end := len(a)

	for swapped {
		swapped = false

		for i := start; i < end-1; i++ {
			if a[i] > a[i+1] {
				temp := a[i]
				a[i] = a[i+1]
				a[i+1] = temp
				swapped = true
				//PrintArray(a, renderer)
			}
		}

		if !swapped {
			break
		}

		swapped = false
		end = end - 1

		for i := end - 1; i >= start; i-- {
			if a[i] > a[i+1] {
				temp := a[i]
				a[i] = a[i+1]
				a[i+1] = temp
				swapped = true
				//PrintArray(a, renderer)
			}
		}
		PrintArray(a, renderer)

		start = start + 1
	}
}
