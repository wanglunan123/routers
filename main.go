package main

import (
	"fmt"
)

func main() {
	router := NewTrie()

	router.Insert("api/v1/java/doc")
	fmt.Println("-------------------")
	router.Insert("api/v1/go/doc")
	fmt.Println("-------------------")
	router.Insert("api/v1/php/doc")
	fmt.Println("-------------------")
	router.Insert("api/v1/js/doc")
	fmt.Println("-----------------")

	router.Insert("api/v2/js/doc")
	fmt.Println("-----------------")
	router.Insert("api/v3/js/doc")
	fmt.Println("-----------------")

	router.Insert("api/v4/js/doc")
	fmt.Println("-----------------")

	fmt.Println(router.Search("api/v1/js"))
	fmt.Println(router.Search("api/v1/java/doc"))
	fmt.Println(router.Search("api/v1/go/doc"))
	fmt.Println(router.Search("api/v4/js/doc"))
	fmt.Println(router.root)
}
