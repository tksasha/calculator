To calculate simple arithmetical formulae.

Examples:

```
package main

import (
  "fmt"

  "github.com/tksasha/formula"
)

func main() {
  result, _ := formula.Calculate("7*8")

  fmt.Println(result) // 56
}
```
