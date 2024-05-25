package main

import "fmt"

type LowCouplingAndHighCohesion struct{}

func (l *LowCouplingAndHighCohesion) Method1() {
	fmt.Println("Method 1 called")
}

func (l *LowCouplingAndHighCohesion) Method2() {
	fmt.Println("Method 2 caled")
}

func main() {
	obj := LowCouplingAndHighCohesion{}
	obj.Method1()
	obj.Method2()
}
