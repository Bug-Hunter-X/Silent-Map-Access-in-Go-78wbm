# Silent Map Access in Go

This example demonstrates a subtle bug in Go related to map access. Accessing a non-existent key in a map does not trigger a runtime error, instead it returns the zero value for the map's value type.  This can lead to unexpected behavior if not carefully considered.

The `bug.go` file shows the problematic code.  The solution, found in `bugSolution.go`, demonstrates how to check for the key's existence before accessing it.

## How to reproduce
1. Save the code in `bug.go`.
2. Run the code using `go run bug.go`. Observe the output.

## Solution
The solution in `bugSolution.go` demonstrates the correct way to access map values, preventing potential errors by first checking if the key exists using the `value, ok := m["key"]` idiom. This approach avoids unexpected behavior by explicitly handling cases where the key is absent.