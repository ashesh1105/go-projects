package main

import "fmt"

func main() {

	/**
	For maps in GO, all keys must be of same type and all values also must be of same type.
	 */
	 // Below is one way to create a map. Like structs, you need comma at every line, including last one!
	colorMap := map[string]string{
		"a": "Red",
		"b": "Blue",
		"c": "Green",
	}

	// Here is another way:
	var hobbies map[int]string

	// Populating the map. Map takes square brace syntax unlike dot syntax with structs (like person.firstName)
	// This is because maps take typed keys, it has to be a string, int, etc. So dot syntax does not make sense with maps
	colorMap["d"] = "Orange"

	// And here's 3rd way of declaring a map using a built-in function in GO:
	bestPlaces := make(map[string]string)

	fmt.Println(colorMap)
	fmt.Println(colorMap["a"])

	fmt.Println(hobbies)	// will print an empty map

	fmt.Println(bestPlaces)	// will print an empty map

	// Delete a record from map
	delete(colorMap, "c")
	fmt.Println(colorMap)

	// Pass a map to a function that iterates over map and prints its elements
	printMap(colorMap)

	/**
	Difference between Maps and Structs:
	A) Keys are strongly typed all keys must be of same type in Map but not in Structs
	B) Values are strongly typed and all values must of be same type in Map but not so in Structs
	C) Maps are passed by reference but Structs are passed by value.
	D) Keys are indexed and we can iterate over them in Maps but not in Structs
	E) We don't need to know all keys while declaring a Map and we can add them later but this is not the case with Structs
	F) Both Maps and Structs are useful and which one to use when depends on a use case.
	 */



}

func printMap(m map[string]string) {

	fmt.Println("\n\nPrinting map:", m)
	for k, v := range m {
		fmt.Println("key:", k, " value:", v)
	}
}
