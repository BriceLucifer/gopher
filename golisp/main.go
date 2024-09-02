package main

import (
	"fmt"
	"github.com/BriceLucifer/golisp/repl"
)

func main(){
	fmt.Printf("GoLisp 👋 Version 0.1\n")
	fmt.Printf("Press Ctrl + C or Type exit to Exit \n")
	repl.Repl()
}
