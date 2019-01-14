package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 =", 2)
	fmt.Println(true || false)
	fmt.Println(!true)

	s := "0x056bc75e2d63100000"
	a := &big.Int{}
	b := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)  // 10^19
	fmt.Println(b)
	a.SetString(s, 0)
	fmt.Println(a)
	fmt.Println(new(big.Int).Div(a, b))
}
