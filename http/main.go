package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type MyCustomWriter struct {}

func main() {

	// resp is of type *Response, i.e., type of pointer to Response object
	// You'll need more processing of resp using ioutil, as shown below
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// Don't do this, follow the steps below this line
	//fmt.Println(resp)

	// Per GO documentation, Read is defined at Reader interface level that takes a byte slice and fills
	// it with data, in this case, body of the html response (struct). So we can print this byte slice
	// after converting it into a string to see the body of the html call response
	// Reader interface saves us from writing our own custom implementation for data reads like
	// images, stream of analytics data, etc, etc.
	//
	// Good news is, we don't need to do below commented code as well since GO helps us on it
	//
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	// This one line of code does it for us!
	//io.Copy(os.Stdout, resp.Body)	// Commented it just to write our custom Writer

	/**
	What actually is happening in above code is, there are standard GO library that implements
	Writer interface. Write interface has method that takes inputs like from HTTP Response, Text,
	Image(s), etc and sends them to standard outputs like terminal, hard drive, outgoing http requests
	etc. If you check the Writer interface, it expects the type implementing it to implement following
	function:
		type Writer interface {
    		Write(p []byte) (n int, err error)
		}
	In above, it copies the bytes from a type implementing Reader interface and returns number of bytes
	copied and any error occurred. os.Stdout is one such type that implements Writer. We can write our own
	as below.
	 */

	 // Get a variable out of custom Writer defined above called MyCustomWriter
	 cw := MyCustomWriter{}

	 // Now you can pass your own writer in io.Copy instead of built-in ones like Stdout.
	 // This is possible just because we are adding our own receiver function to this writer below as Writer
	 io.Copy(cw, resp.Body)


}

// MyCustomWriter will never be a Writer till it implements this function!
// Also note that since we aren't really using a variable of type MyCustomWriter in this function,
// we do not need to write it as func (cw MyCustomWriter) Write...
func (MyCustomWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	n := len(bs)
	fmt.Printf("\nJust wrote %d number of bytes onto console from MyCustomWriter!", n, "\n")

	return n, nil
}