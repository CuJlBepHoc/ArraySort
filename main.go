package main

import (
	"fmt"
	"sort"
)

// 1. Написать функцию, которая производит слияние отсортированных массивов.
// Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
func sortArrays(arr1 [4]int, arr2 [5]int) ([4]int, [5]int) {
	sort.Ints(arr1[:])
	sort.Ints(arr2[:])

	return arr1, arr2
}

func mergeArrays(arr1 [4]int, arr2 [5]int) (merged [9]int) {
	i, j, k := 0, 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			merged[k] = arr1[i]
			i++
		} else {
			merged[k] = arr2[j]
			j++
		}
		k++
	}

	for i < len(arr1) {
		merged[k] = arr1[i]
		i++
		k++
	}

	for j < len(arr2) {
		merged[k] = arr2[j]
		j++
		k++
	}

	return merged
}

// 2. Отсортировать массив пузырьком.
// Отсортируйте массив длиной шесть пузырьком.
func bubbleSort(arr *[6]int) {
	n := len(*arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Выполнение первой задачи
	firstArray := [4]int{34, 76, 12, 43}
	secondArray := [5]int{56, 78, 32, 42, 9}

	sortedArray1, sortedArray2 := sortArrays(firstArray, secondArray)
	result := mergeArrays(sortedArray1, sortedArray2)
	fmt.Println("Отсортированный список длинной 9:", result)

	fmt.Println("-----------------------------------------------------")
	// Выполнение второй задачи
	array := [6]int{34, 56, 21, 89, 67, 2}
	bubbleSort(&array)
	fmt.Println("Сортировка пузырьком:", array)
}
