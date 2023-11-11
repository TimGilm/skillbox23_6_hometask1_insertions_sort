/*
Задание 1. Сортировка вставками
Что нужно сделать
Напишите функцию, сортирующую массив длины 10 вставками.
*/
package main

import "fmt"

const size = 10

func main() {
	a := [size]int{8, 9, 3, 5, 2, 6, 1, 7, 4, 0}
	fmt.Println(a)
	fmt.Println(insertionSort(a))
}

func insertionSort(a [size]int) [size]int {
	for i := 1; i < size; i++ {
		for j := i - 1; j >= 0; j-- {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
				i--
			}
		}
	}
	return a
}
