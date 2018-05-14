package subpackage_A

import "fmt"

type SubA struct {
	Name string
}

func (sa SubA) PrintA() {
	fmt.Println("subpackage A [" + sa.Name + "]")
}
