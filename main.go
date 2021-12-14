package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/tatskaari/pho/result"
)

func main() {
	// Wrap will create a result type from the usual error tuple returned from non-generic go code
	wd := result.Wrap(os.Getwd())

	// Map applies the function to the value in the result unless the result is not okay
	path := result.Map(wd, func(wd string) string {
		 return filepath.Join(wd, "file.txt")
	})

	// The wrap is like Map however the function can return an error. In which case this will be wrapped up nicely 
	// for you
	file := result.ThenWrap(path, os.Open)

	// Use applies a procedure to the value, or returns an error if the result is not okay
	defer file.Use(func(file *os.File) {
		file.Close()
	})

	// Cast can be useful to conver the type of the result. Go generics are not covarient so this is often necessary. 
	// In this case, ioutil.ReadAll expects a io.Reader, which *io.File implements. Go isn't happy with this though,
	// so we have to cast it here. 
	reader := result.Cast[*os.File, io.Reader](file)

	bytes := result.ThenWrap(reader, ioutil.ReadAll)

	// Once we've mapped all our functions to the result, we can unwrap it to get at the raw result type. 
	// There's also MustUnwrap which will panic if the result is not okay. 
	str, err := result.Map(bytes, func(b []byte) string {
		return string(b)
	}).Unwrap()

	if err != nil {
		fmt.Println("Error reading file: %v", err)
		return
	}
	fmt.Println(str)

}
