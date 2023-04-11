package main

import "fmt"

func divide(a float64, b float64) (float64, error) {

	if b == 0 {

		return 0, fmt.Errorf("não pode dividir por 0")
	}

	return a / b, nil

}

func main() {

	result, err := divide(21, 0)

	if err != nil {

		fmt.Printf("Ocorreu um erro ao dividir a e b: %s.\n", err)

	}

	fmt.Println(result)

}
