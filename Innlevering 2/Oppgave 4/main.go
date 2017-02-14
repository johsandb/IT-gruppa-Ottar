package main

import ("IntelliJ/Algorithms"
	//"container/list"
	"fmt"
)

func main (){
	values := []int{9, 1, 20, 3, 6, 7}

	fmt.Println("Tallrekke:", values)
	//fmt.Println(values)
	algorithms.Bubble_sort_modified(values)
	//algorithms.Bubble_sort(values)
	fmt.Println("Sorted:  ", values)
}
