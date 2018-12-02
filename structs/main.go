package main

import "fmt"

func main() {

	/**
	Note: In GO, you need a comma at end of every line in between the statement below.
	If you do this:
	alex := Person{
		firstName: "Alex",
		lastName:  "Wong"
	}
	GO will make this statement as:
	alex := Person{
		firstName: "Alex",
		lastName:  "Wong";
	};
	Which will obviously be a compile time error! So you must add a comma at the end last line too!
	*/
	alex := Person{
		firstName: "Alex",
		lastName:  "Wong",
		contact: ContanctInfo{
			phone:   "408-123-4567",
			address: "123 San Tomas Ave, San Jose, CA 95129",
		}, // Notice, you need a comma here as well, just so GO does not insert a semicolon here!
	}

	// Change a specific field in struct
	alex.lastName = "Anderson"
	alex.contact.phone = "408-111-2222"

	// Call the receiver function for Person defined in person.go
	//alex.print()

	// Let's change the person alex using receiver function
	fmt.Println("\nPrinting Alex after updating its first name via receiver function:")

	/**
	As you can see, below print statement won't print first name as Terry and will print what it originally had.
	This is because:
	1) When you call setFirstName on alex, its "pass by value" there, GO copies original person to another space first.
	2) Then GO updates the first name of new person value at a brand new address.
	3) When you say alex.print(), you tell GO to print the original person, which was never updated!
	 */
	alex.setFirstName("Terry")
	alex.print()

	/**
	In order to tell GO to modify original object (person alex), we need to:
	1) Change the function to accept a type of pointer to Person type of data (p *Person).
	2) Then call that function instead.
	3) GO allows you a flexibility there to directly call this on alex object instead of getting the address of it
	and then call this new function on that address data. This is a short cut GO provides.
	 */

	// You can either do below
	pointerToAlex := &alex
	fmt.Println("\n\nCalling receiver function using address of person object to function that takes type as pointer to Person type to update firstName:")
	pointerToAlex.setFirstNameWithPointer("Prashant")
	alex.print()

	// Or this will work too:
	fmt.Println("\n\nCalling receiver function using Person data itself to function that takes type as pointer to to Person type to update firstName:")
	alex.setFirstNameWithPointer("Terry")
	alex.print()

	fmt.Println("\n\nUpdating this person by passing address this data to a function:")
	updatePersonFirstName(&alex)
	alex.print()

}

func updatePersonFirstName(p *Person) {
	(*p).firstName = "Ashesh"
}
