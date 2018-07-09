package main

import (
	"fmt"
)

type epyt struct {
	Name string
}

func main() {
	var input = []int{31, 27, 16,  7, 20, 44, 10, 11,  8, 41, 49, 48,  1, 40, 16,  9, 13,
		27, 42, 45, 41, 35, 49, 30, 39, 18, 25, 29, 38,  6, 47, 21, 44, 26,
		31,  7,  1, 17, 38, 40, 29,  1, 21, 30,  4, 28,  6, 30,  7,  8}
	fmt.Println(input)
	for i:=1; i < len(input); i++ {
		for j:=0; j < len(input)-i; j++{
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	fmt.Println(input)
	var input2 = []int{31, 27, 16,  7, 20, 44, 10, 11,  8, 41, 49, 48,  1, 40, 16,  9, 13,
		27, 42, 45, 41, 35, 49, 30, 39, 18, 25, 29, 38,  6, 47, 21, 44, 26,
		31,  7,  1, 17, 38, 40, 29,  1, 21, 30,  4, 28,  6, 30,  7,  8}
	for i:=0; i < len(input2)-1; i++ {
		current := input2[i+1]
		j := i
		for j>=0 && input2[j] >= current{
			input2[j+1] = input2[j]
			j--
		}
		input2[j+1] = current
	}
	fmt.Println(input2)
	input = []int{31, 27, 16,  7, 20, 44, 10, 11,  8, 41, 49, 48,  1, 40, 16,  9, 13,
		27, 42, 45, 41, 35, 49, 30, 39, 18, 25, 29, 38,  6, 47, 21, 44, 26,
		31,  7,  1, 17, 38, 40, 29,  1, 21, 30,  4, 28,  6, 30,  7,  8}
	fmt.Println(merge(input))
	input = []int{31, 27, 16,  7, 20, 44, 10, 11,  8, 41, 49, 48,  1, 40, 16,  9, 13,
		27, 42, 45, 41, 35, 49, 30, 39, 18, 25, 29, 38,  6, 47, 21, 44, 26,
		31,  7,  1, 17, 38, 40, 29,  1, 21, 30,  4, 28,  6, 30,  7,  8}
		quicksort(input, 0, len(input))
		fmt.Println(input)
	obj := Newstuct()
	fmt.Println(obj)
	}

func mergesort(left []int, right []int)  []int{
	results := make([]int, 0,len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			results = append(results, left[i])
			i ++
		}else {
			results = append(results, right[j])
			j ++
		}
	}
	if i == len(left) {
		results = append(results, right[j:]...)
	} else {
		results = append(results, left[i:]...)
	}
	return results
}

func merge(arr []int) []int{
	if len(arr) <= 1{
		return arr
	}
	mid := len(arr)/2
	leftarr := merge(arr[0: mid])
	rightarr := merge(arr[mid: ])
	return mergesort(leftarr, rightarr)
}

func aport(arr []int, left ,right int) int{
	//head := arr[left]
	i, j := left+1, left+1
	for j<right{
		if arr[j]<arr[left]{
			arr[i], arr[j] = arr[j], arr[i]
			i ++
		}
		j++
	}
	arr[i-1] ,arr[left] = arr[left], arr[i-1]
	return i -1
}

func quicksort(arr []int, low, high int){
	if low+1 < high {
		mark := aport(arr, low, len(arr))
		quicksort(arr, low, mark-1)
		quicksort(arr, mark+1, high)
	}
}



func partition(aris []int, begin int, end int) int {
	pvalue := aris[begin]
	i := begin
	j := begin + 1
	for j < end {
		if aris[j] < pvalue {
			i++
			aris[i], aris[j] = aris[j], aris[i]
		}
		j++
	}
	aris[i], aris[begin] = aris[begin], aris[i]
	return i
}

func quick_sort(aris []int, begin int, end int) {
	if begin+1 < end {
		mid := partition(aris, begin, end)
		quick_sort(aris, begin, mid)
		quick_sort(aris, mid+1, end)
	}
}


func Newstuct() epyt{
	obj := epyt{}
	obj.Name = "epyt"
	return obj
}