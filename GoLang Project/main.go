package main

import "fmt"

func main() {

	var entrada int
	resultado := 0

	fmt.Scan(&entrada)

	for i := 2; i <= entrada/2; i++ {

		if entrada%i == 0 {
			resultado++
		}

	}

	fmt.Println("O número digitado é:", numeroPrimo(resultado))

}
func numeroPrimo(valor int) string {

	if valor == 0 {

		return "primo"
	} else {

		return "não primo"
	}

}
