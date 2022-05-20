package main

import (
	"fmt"

	"os"

	"github.com/m4k0m1/qeq"
)

func main() {

	s1, s2, err := qeq.Solveqeq(1, -6, 5)

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	fmt.Printf("Solutions calculated: x = %f or x = %f\n", s1, s2)

}
