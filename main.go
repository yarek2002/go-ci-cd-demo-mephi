package main

import (
	"fmt"
	"go-ci-cd-demo/mathutils"
)

func main() {
	fmt.Println("2 + 3 =", mathutils.Add(2, 3))
	fmt.Println("5 - 3 =", mathutils.Sub(5, 3))
}
