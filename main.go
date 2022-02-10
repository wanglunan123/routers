package main

import (
	"fmt"
)

func main() {
	router := NewTrie()

	router.Insert("api/v1/:lang/doc")
	fmt.Println("-------------------")
	router.Insert("api/v1/:language/doc")
	fmt.Println("-------------------")
	fmt.Println(router.Search("api/v1/java/doc"))
	fmt.Println(router.Search("api/v1/go/doc"))
	fmt.Println(router.Search("api/v1/php/doc"))
	fmt.Println(router.Search("api/v1/donet/doc/1"))
}
