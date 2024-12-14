```go
package main

import "fmt"

func main() {
    var m map[string]int
    m = make(map[string]int)
    value, ok := m["key"]
    if ok {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key not found")
    }
}
```