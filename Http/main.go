package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www.google.com/")

	if err != nil {
		log.Fatal("Error in request", err)
	}

	// bs := make([]byte, 8*1024)
	// fmt.Println(string(bs))

	// body, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", body)

	// stephens way
	io.Copy(os.Stdout, res.Body)
	// check out source code for io.Copy! Fascinating!
}

// http.Get returns (res *Response, err error)
// The Response is a struct with a field of body which is of type io.ReadCloser
// io.ReadCloser is actually an interface!
/*
ReadCloser is the interface that groups the basic Read and Close methods.
type ReadCloser interface {
	Reader
	Closer
}
*/
/*
Reader is the interface that wraps the basic Read method.
type Reader interface {
	Read(p []byte) (n int, err error)
}
*/

// *** A field inside of a struct can be an interface
// Body field inside the Response Struct can have any value so long as it fulfills the
// io.ReadCloser interface

// A type satisfies an interface when that type implements all of the functions contained in the interface definition
