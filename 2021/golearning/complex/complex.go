package main

import "fmt"

func main() {
	num := 1 + 2i
	scalar_res := num * 2
	fmt.Println(scalar_res)
	real_part := real(scalar_res)
	imaginary_part := imag(scalar_res)

	fmt.Printf("Real part: %v\nImaginary part: %v\n", real_part, imaginary_part)

	complex_res := num * complex(2, 0)
	fmt.Println(complex_res)
}