package main

import (
	"fmt"
	//"github.com/BriceLucfier/golox/chunk"
	//"github.com/BriceLucfier/golox/memory"
)

type hello struct{
	name string
}

func main() {
	tmep := hello {
		"world",
	}
	fmt.Println("hello wdsjadljasld")
	for i := 12; i < 20; i++ {
		fmt.Printf("%v", i)
		fmt.Printf("%T", i)
		fmt.Println()
		fmt.Printf("%#v",tmep)
		fmt.Println()
	}
	fmt.Println("hello world")
}
