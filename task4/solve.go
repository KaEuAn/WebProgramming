package main

import "fmt"

func RemoveEven(slice []int) (newSlice []int) {
	newSlice = make([]int, 0, len(slice))
	for _, elem := range slice {
		if elem % 2 != 0 {
			newSlice = append(newSlice, elem)
		}
	}
	return
}

func main(){
	input := []int{0, 3, 2, 5}
	result := RemoveEven(input)
	fmt.Println(result) // Должно напечататься [3 5]
}