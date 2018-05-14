package subpackage_B

import "fmt"

type SubB struct {
	Name string
}

func (sb SubB) PrintA() {
	fmt.Println("subpackage B [" + sb.Name + "]")
}
