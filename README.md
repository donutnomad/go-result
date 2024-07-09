# go-result

A Go package that provides a functional approach to error handling using the `Result` type.

## Overview

`go-result` is a Go library designed to simplify error handling by using a functional programming pattern. It introduces a `Result` type that can encapsulate either a successful value or an error, allowing for more readable and maintainable code when dealing with operations that may fail.

## Features

- **Generic Result Type**: Supports any type for success and error values.
- **Immutability**: Once a `Result` is created, its state cannot change, promoting safe usage across functions.
- **Composability**: Easily chain operations with `MapOk` and `MapErr` to transform success and error values.
- **Inspection**: Use `Inspect` and `InspectErr` to perform actions based on the presence of a value or error.

## Getting Started

To get started with `go-result`, you'll first need to install the package:

```bash
go get github.com/donutnomad/go-result
```

After installation, you can import it into your Go project:

```go
import "github.com/donutnomad/go-result"
```

## Usage

Here's a simple example of how to use `go-result`:

```go
package main

import (
	"fmt"
	"github.com/donutnomad/go-result"
)

func main() {
	// Create a successful result
	okResult := result.NewOk[int, string](42)

	// Map the successful value to a new type
	newResult := result.MapOk[okResult, string](func(value int) string {
		return fmt.Sprintf("The answer is: %d", value)
	})

	// Handle the result
	if newResult.IsOk() {
		fmt.Println(newResult.Unwrap()) // Output: The answer is: 42
	} else {
		fmt.Println(newResult.UnwrapErr())
	}
}
```

## API Reference

- `NewOk[T, E](value T) Result[T, E]`: Create a new `Result` with a success value.
- `NewErr[T, E](err E) Result[T, E]`: Create a new `Result` with an error value.
- `MapOk[T, E, U](r Result[T, E], ok U) Result[U, E]`: Transform the success value to a new type.
- `MapErr[T, E, EU](r Result[T, E], err EU) Result[T, EU]`: Transform the error value to a new type.
- `Inspect[T, E](r Result[T, E], f func(T))`: Perform an action if the result is a success.
- `InspectErr[T, E](r Result[T, E], f func(E))`: Perform an action if the result is an error.

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

Thank you to the Go community for providing a robust and versatile language for systems programming.
