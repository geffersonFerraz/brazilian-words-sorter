# brazilian-words-sorter

A simple Go program that sort a predefined list of Brazilian words.

## How to use

```
package your_package

import (
    "fmt"
    "github.com/geffersonFerraz/brazilian-words-sorter"
)

func main() {
	bws := BrazilianWords(3, ",")
	fmt.Println(bws.Sort())
}

```