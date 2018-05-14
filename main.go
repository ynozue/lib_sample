package main

import (
	"fmt"

	"github.com/ynozue/lib_sample/subpackage_A"
	"github.com/ynozue/lib_sample/subpackage_B"
)

type Main struct {
	Name string
}

func (m Main) PrintMain() {
	fmt.Println("Main [" + m.Name + "]")
}

func main() {
	fmt.Println("Hello! [https://github.com/ynozue/lib_sample]")
	m := Main{Name: "メイン"}
	m.PrintMain()
	a := subpackage_A.SubA{Name: "サブA"}
	a.PrintA()
	b := subpackage_B.SubB{Name: "サブB"}
	b.PrintA()
}
