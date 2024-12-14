```go
package main

import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["key"]) // this will not cause a runtime error
}
```