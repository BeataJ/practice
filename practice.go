package main

import "fmt"

type Product struct {
	title string
	price float64
	id    int
}

func main() {
	// 1
	hobbies := [3]string{"cooking", "reading", "coding"}
	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[0])
	new_hobbies := hobbies[1:]
	fmt.Println(new_hobbies)

	// 3
	slice_hobbies := hobbies[:2]
	slice_hobbies2 := hobbies[0:2]
	fmt.Println(slice_hobbies)
	fmt.Println(slice_hobbies2)
	fmt.Println(cap(slice_hobbies))

	// 4
	slice_hobbies = slice_hobbies[1:3]
	fmt.Println(slice_hobbies)

	// 5
	dynamicArray := []string{"coding Go", "Be good in go"}
	fmt.Println(dynamicArray)

	// 6
	dynamicArray[1] = "Go Go Go"
	fmt.Println(dynamicArray)
	addSlice := "new goal"
	dynamicArray = append(dynamicArray, addSlice)

	fmt.Println(dynamicArray)

	// 7
	products := []Product{
		{
			title: "glass",
			price: 8.99,
			id:    1,
		},
		{
			title: "bottle",
			price: 99,
			id:    2,
		},
	}

	fmt.Println(products)
	p := Product{
		"stone",
		3.99,
		3,
	}

	products = append(products, p)
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
