package main

import "fmt"

func main() {

	const numberOfProductsInList = 4

	shoppingList := createShoppingList(numberOfProductsInList)

	fmt.Println("Your shopping list is:", shoppingList)

	searchProduct(numberOfProductsInList, shoppingList)

	printRangeShoppingList(shoppingList)

}

func createShoppingList(m int) []string {

	var shoppingList []string = make([]string, m)

	for i := 0; i < m; i++ {
		fmt.Printf("Here is your shopping list. Input product #%v  -  ", i)
		var product string
		fmt.Scan(&product)

		shoppingList[i] = product

	}

	return shoppingList

}

func searchProduct(numberOfProductsInList int, shoppingList []string) {

	var findWord string
	fmt.Printf("To search, type name of the product:   ")
	fmt.Scan(&findWord)

	for n := 0; n < numberOfProductsInList; n++ {
		if findWord == shoppingList[n] {
			fmt.Printf("FOUND MATCH -- product #%v,  %v   \n", n, shoppingList[n])
		} else {
			fmt.Println("MATCH was not found")
		}
	}
}

func printRangeShoppingList(shoppingList []string) {
	for b, rangedShoppingList := range shoppingList {
		fmt.Println(b, rangedShoppingList)
	}
}
