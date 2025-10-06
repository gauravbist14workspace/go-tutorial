# Consider a project structure like:

|___main.go
    |___ds_algo
    |   |___alien_language.go (this has "main" as package & a main() )
    |
    |___bit_bool
    |   |___bit_bool_alternate.go   (this doesnt have "main" as package & will be imported by ROOT main() )


# To import bit_bool_alternate.go inside ROOT main.go
```
package main

import (
    "go_tutorial/bit_bool"
)

func main() {
    // here we can import bit_bool functions
}
```
## This import line imports all files inside the "/bit_bool" directory
## And if they have function starting with CAPITAL LETTERS, they will be usable from main.go#