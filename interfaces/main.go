package main

import "fmt"

/**
Notice the declaration of an interface here.
Basically, it says that any type in the scope of this program that implements a function
getGreeting that returns a string is an "honorary member" of my type. In this case, the two
struts englishBot and spanishBot, both have a receiver functions getGreeting functions that
return a string, so they are of type bot!
Basically, this is how interfaces are linked to a specific type in GO.

Note that since getGreeting function for both the type of bots do not accept any argument,
we're okay to define this method signature in interface like this but say, the getGreeting
functions in two type of bots would take an int argument (getGreeting(i int)), then the
interface had to has the method as: getGreeting(int) string

In general, an interface looks like this:
type <interface_name> interface {
	<method_name>(<argument_type>, <argument_type>..) (<return_type>, <return_type>..)
}
e.g.:
type car interface {
	drive(string, int) (string, error)
	getModel(string) string
	...
}

Interface in GO is implicit and GO figures out which classes can be covered by it based on method(s) they have.
We do not have to say SomeClass implements SomeInterface.. etc., in GO. This can be a blessing as well as curse. Easy to
code but can be hard to figure out.



 */
type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func main() {

	var eb englishBot
	var sb spanishBot

	printBot(eb)
	printBot(sb)

}

func printBot(b bot) {
	msg := b.getGreeting()
	fmt.Println(msg)
}

// You don't have to do this. Use above function that takes an interface rather!
//func printBot(eb englishBot) {
//	msg := eb.getGreeting()
//	fmt.Println(msg)
//}

// Below will give you an error even if type of argument it takes is different from englishBot
//func printBot(sb spanishBot) {
//	msg := sb.getGreeting()
//	fmt.Println(msg)
//}

// Receiver function for type englishBot
func (eb englishBot) getGreeting() string {

	// Custom logic goes here
	return "Hello people!"
}

// Receiver function for type spanishBot
func (sb spanishBot) getGreeting() string {

	// Custom logic goes here
	return "Hola!"
}
