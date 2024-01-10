package main

import (
	"errors"
	"fmt"

	"github.com/mv-kan/go-error-inheritance/pkg/communication"
	myhttp "github.com/mv-kan/go-error-inheritance/pkg/http"
)

/*
//////////////////////////////////////////////

# HOW TO CREATE ERROR INHERITANCE IN GOLANG? #

//////////////////////////////////////////////

1. If you need to pass arguments to your error object, then you need to create
struct type and implement Error and Unwrap methods for this struct type like in
pkg/communication package. See TimeoutError in pkg/communication.
2. If you have one static message without arguments then you can use fmt.Errorf() function
like in pkg/http package. See ErrorNotImplemented in pkg/http.
3. Or you can just make up errors on the go without defining the variables like
in NoDefinitionErrorFunction function
*/

func ReallyBadFunction(timeout int) error {
	if timeout%2 == 0 {
		return &communication.TimeoutError{
			Timeout: fmt.Sprint(timeout),
		}
	} else {
		return &myhttp.ServerError{
			Msg: "strange kind of error happened",
		}
	}
}

func httpNotImplemented() error {
	return myhttp.ErrorNotImplemented
}

func NoDefinitionErrorFunction(msg string) error {
	return fmt.Errorf("%w: i am too lazy to define error, here is the string msg \"%s\"", communication.ErrorCommunication, msg)
}

func main() {
	err := ReallyBadFunction(10)
	// err = ReallyBadFunction(9) // try with odd number

	if errors.Is(err, communication.ErrorCommunication) {
		fmt.Println("Communication Error happened")
		fmt.Printf("%v\n", err)
	} else if errors.Is(err, myhttp.ErrorHttp) {
		fmt.Println("Http Error happened")
		fmt.Printf("%v\n", err)
	}
	fmt.Println()

	err = httpNotImplemented() // returns error that is ErrorHttp
	if errors.Is(err, communication.ErrorCommunication) {
		fmt.Println("Communication Error happened")
		fmt.Printf("%v\n", err)
	} else if errors.Is(err, myhttp.ErrorHttp) {
		fmt.Println("Http Error happened")
		fmt.Printf("%v\n", err)
	}
	fmt.Println()

	err = NoDefinitionErrorFunction("Hehehe")
	if errors.Is(err, communication.ErrorCommunication) {
		fmt.Println("Communication Error happened")
		fmt.Printf("%v\n", err)
	} else if errors.Is(err, myhttp.ErrorHttp) {
		fmt.Println("Http Error happened")
		fmt.Printf("%v\n", err)
	}

}
