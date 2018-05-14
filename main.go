package main

import (
	"fmt"

	"github.com/ynozue/lib_sample/subpackage_A"
	"github.com/ynozue/lib_sample/subpackage_B"
)

func main() {
	fmt.Println("Hello! [https://github.com/ynozue/lib_sample]")
	a := subpackage_A.SubA{Name: "サブA"}
	a.PrintA()
	b := subpackage_B.SubB{Name: "サブB"}
	b.PrintA()
}
