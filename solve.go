package main

import ("fmt"; "unicode"; "strings")

func RemoveEven(slice []int) (newSlice []int) {
	newSlice = make([]int, 0, len(slice))
	for _, elem := range slice {
		if elem % 2 != 0 {
			newSlice = append(newSlice, elem)
		}
	}
	return
}

func DifferentWordsCount(s string) (count uint64) {
	count = 0
	s = strings.ToLower(s)
	dict := make(map[string]uint8)
	i := 0
	for i < len(s) {
		var curString string = ""
		for ; i < len(s) && unicode.IsLetter(rune(s[i])); i++ {
			curString += string(s[i])
		}
		if _, ok := dict[curString]; ok == false {
			dict[curString] = 1;
			count++
		}
		for ;i < len(s) && ! unicode.IsLetter(rune(s[i])); i++ {}
	}
	return
}


func PowerGenerator(x int) (func() int) {
	number := x
	answer := 1
	return func() (int) {
		answer *= number
		return answer
	}
}



/*func main(){
	input := []int{0, 3, 2, 5}
	result := RemoveEven(input)
	fmt.Println(result) // Должно напечататься [3 5]
	gen := PowerGenerator(2)
	fmt.Println(gen()) // Должно напечатать 2
	fmt.Println(gen()) // Должно напечатать 4
	fmt.Println(gen()) // Должно напечатать 8
	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
	// Должно напечатать 2
}*/