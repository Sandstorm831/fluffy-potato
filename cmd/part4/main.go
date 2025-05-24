package main

import "fmt"

func main() {
	var intArr [3]int = [3]int{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length of extended slice is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("The length of slice3 is %v and the capacity is %v\n", len(intSlice3), cap(intSlice3))

	//////////////////
	var myMap map[string]uint32 = make(map[string]uint32)
	fmt.Println(myMap)

	var myMap2 map[string]uint16 = map[string]uint16{"Adam": 13, "ramesh": 23}
	fmt.Println(myMap2["Adam"])

	var age, ok = myMap2["Jason"]
	if ok{
		fmt.Printf("The age of jason is %v\n", age)
	}else{
		fmt.Printf("Jason doesn't exist in the map\n")
	}

	delete(myMap2, "Adam")
	fmt.Println(myMap2)


	for name, age := range myMap2{
		fmt.Printf("Name is %v and the age is %v \n", name, age)
	}

	for i, v := range intArr{
		fmt.Printf("index is %v, and the val is %v \n", i, v)
	}

	for i:=0; i<10; i++{
		fmt.Printf("For loop with i = %v \n", i)
	}

}
