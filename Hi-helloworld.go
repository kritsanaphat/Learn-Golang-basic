package main

import (
	"fmt"
)

const (
	PI  = 3.14
	PI2 = 3.141
)

func main() {

	var a string = "Hello"
	var b int = 15
	var c bool
	d := 8.0
	e := "FEE"

	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", c, c)
	fmt.Printf("%T %v\n", d, d)
	fmt.Printf("%T %v\n", e, e)

	//-----------------------------------------------------
	var a1, b1, c1, d1 int = 1, 3, 5, 7
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)

	//------------------------------------------------------
	fmt.Println(PI)
	fmt.Println(PI2)

	//------------------------------------------------------
	var txt = "Hello World!"
	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)

	//------------------------------------------------------
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	arr3 := [...]int{4, 5, 6, 7, 8}
	arr4 := [5]int{1: 10, 2: 40}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	//------------------------------------------------------
	mysliceBydatatype := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(mysliceBydatatype))
	fmt.Println(cap(mysliceBydatatype))
	fmt.Println(mysliceBydatatype)

	arr5 := [6]int{10, 11, 12, 13, 14, 15}
	mysliceByArray := arr5[2:4]

	fmt.Printf("myslice = %v\n", mysliceByArray)
	fmt.Printf("length = %d\n", len(mysliceByArray))
	fmt.Printf("capacity = %d\n", cap(mysliceByArray))

	myslicByMake1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslicByMake1)
	fmt.Printf("length = %d\n", len(myslicByMake1))
	fmt.Printf("capacity = %d\n", cap(myslicByMake1))

	// with omitted capacity
	mysliceByMake2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", mysliceByMake2)
	fmt.Printf("length = %d\n", len(mysliceByMake2))
	fmt.Printf("capacity = %d\n", cap(mysliceByMake2))
}
