package main // Packages have a directory namespace

// Is the function that's run the same as the package name by default ?

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
