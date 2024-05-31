package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Choose how to input data: press 1 to input from terminal, press 2 to input from file")
	var dataOrigin int
	fmt.Scan(&dataOrigin)

	var stopAsking bool

	for !stopAsking {

		if dataOrigin == 1 {
			fmt.Printf("How manu products are you going to input to shopping list?  ")
			var number int
			fmt.Scan(&number)
			shoppingList := createShoppingList(number)
			fmt.Println("Your shopping list is:", shoppingList)
			stopAsking = true
			searchProduct(shoppingList)
			printRangeShoppingList(shoppingList)
		} else if dataOrigin == 2 {
			// input fom file
			shoppingList := createListFromFile()
			stopAsking = true
			searchProduct(shoppingList)
			printRangeShoppingList(shoppingList)
		} else {
			fmt.Println("Wrong input. Try again")
			stopAsking = false
		}
	}
}

func createListFromFile() []string {
	myFile, err := os.Open("Shoppinglist.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer myFile.Close()
	var shoppingList []string
	scanner := bufio.NewScanner(myFile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		shoppingList = append(shoppingList, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
	fmt.Println("Your Shopping list slice is:  ", shoppingList)
	for b, rangedShoppingList := range shoppingList {
		fmt.Println(b, rangedShoppingList)
	}
	return shoppingList
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

func searchProduct(shoppingList []string) {
	var findWord string
	fmt.Printf("To search, type name of the product:   ")
	fmt.Scan(&findWord)
	for n := 0; n < len(shoppingList); n++ {
		if //findWord == shoppingList[n] {
		strings.Contains(shoppingList[n], findWord) {
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
