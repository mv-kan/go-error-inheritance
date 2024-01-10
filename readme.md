# HOW TO CREATE ERROR INHERITANCE IN GOLANG? 

1. If you need to pass arguments to your error object, then you need to create
struct type and implement Error and Unwrap methods for this struct type like in
`pkg/communication` package. See `TimeoutError` in `pkg/communication`.
2. If you have one static message without arguments then you can use `fmt.Errorf()` function
like in `pkg/http` package. See `ErrorNotImplemented` in pkg/http.
3. Or you can just make up errors on the go without defining the variables like
in `NoDefinitionErrorFunction` function in `main.go`