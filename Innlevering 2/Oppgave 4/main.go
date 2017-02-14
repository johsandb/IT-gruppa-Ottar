package main

import ("./Algorithms"
	"fmt"
)

func main (){
	values := []int{9, 1, 20, 3, 6, 7}

	fmt.Println("Tallrekke:", values)
	algorithms.Bubble_sort_modified(values)
	//algorithms.Bubble_sort(values)
	fmt.Println("Sorted:  ", values)
}
