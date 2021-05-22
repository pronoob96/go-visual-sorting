package main

import "github.com/veandco/go-sdl2/sdl"

func merge(arr []int32, l, m, r int, renderer *sdl.Renderer) {
	n1 := m - l + 1
	n2 := r - m

	L := make([]int32, n1)
	R := make([]int32, n2)

	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}

	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	i := 0
	j := 0
	k := l

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
		PrintArray(arr, renderer)
	}

	for i < n1 {
		arr[k] = L[i]
		i++
		k++
		PrintArray(arr, renderer)
	}

	for j < n2 {
		arr[k] = R[j]
		j++
		k++
		PrintArray(arr, renderer)
	}
}

func MergeSort(arr []int32, l, r int, renderer *sdl.Renderer) {
	if l >= r {
		return
	}
	m := l + (r-l)/2
	MergeSort(arr, l, m, renderer)
	MergeSort(arr, m+1, r, renderer)
	merge(arr, l, m, r, renderer)

}
