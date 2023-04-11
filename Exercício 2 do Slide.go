package main

import "fmt"

func m(list []int) (int, error) {

	if len(list) == 0 {

		return 0, fmt.Errorf("a list está vazia")

	}

	soma := 0

	for _, x := range list {

		soma += x

	}

	return soma / len(list), nil

}

func main() {

	result, err := m([]int{1, 2, 3, 4, 5})

	if err != nil {

		fmt.Printf("Houve um erro ao executar o programa: %s.\n", err)

	} else {

		fmt.Printf("A média é : %d.", result)

	}

}
