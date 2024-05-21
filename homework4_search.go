package main

import "fmt"

func main() {

	var shoppingList []string = make([]string, 5)
	for i := 0; i < 5; i++ {
		fmt.Printf("Input product %v  -  ", i)
		var product string
		fmt.Scan(&product)

		shoppingList[i] = product

	}
	fmt.Println("Your shopping list is:", shoppingList)

	var findWord string
	fmt.Printf("Input product to find:   ")
	fmt.Scan(&findWord)

	for n := 0; n < 5; n++ {
		if findWord == shoppingList[n] {
			fmt.Printf("FOUND MATCH  --  %v   \n", shoppingList[n])
			//} else {
			//	fmt.Println("MATCH was not found")
		}
	}

}
