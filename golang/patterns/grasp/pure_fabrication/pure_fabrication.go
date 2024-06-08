package main

import "fmt"

type Helper struct{}

func (h *Helper) UsefulMethod() {
	fmt.Println("This is a useful method provided by Helper entity.")
	
}

func main() {
	helper := Helper{}
	helper.UsefulMethod()
}
