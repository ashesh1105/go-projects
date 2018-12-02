package main

import "fmt"

func main() {

	/**
	As you can see below, mySlice got updated when we pass it to updateSlice function.
	We did not however, see this behavior with slice, as in GO project, "structs" in same parent folder
	The reason here is:
	A) GO has a "Pass by Reference" behavior with slices, maps, channels, pointers, functions, etc.
		- We do not need to worry about pointers for these data types.
	B) GO has "Pass by Value" behavior with primitive types like int, float, string, bool. structs, etc.
		- We need to manage these via pointers for updates, etc.

	In case of slice, it does has "Pass by Value" behaviour with reference to slice, where GO does copy the reference of
	slice passed to function to another address space, but both the references actually point to same underlying Array.
	Slices are backed by Arrays, same as ArrayLists in Java. So this allows slices to get updated when passed on example
	function as below!
	 */
	mySlice := []string {"Hello", "dude!", "how", "are", "you?"}

	fmt.Println(mySlice)

	updateSlice(mySlice)
	fmt.Println(mySlice)


}

func updateSlice(s []string) {
	s[0] = "Hey"
}
