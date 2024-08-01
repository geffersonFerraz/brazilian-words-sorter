package main

import (
	"bws"
	"fmt"
)

func main() {
	bwords := bws.BrazilianWords(3, "-")
	fmt.Println(bwords.Sort())
}
