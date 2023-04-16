// Escreva uma função que calcule a média de uma lista de números e retorne um erro caso a lista esteja vazia.

package main // Chamamos o pacote que contém as bibliotecas que usaremos. Nesse caso, o pacote "main".

import "fmt" // Importamos a biblioteca "fmt" do pacote, pois ela tem as funções que usaremos.

func m(list []int) (int, error) { // Criamos a função "m", a qual retornará um valor do tipo inteiro e um erro, caso aconteça, do tipo error. Esta função criará uma lista de números inteiros.

	if len(list) == 0 { // Puxamos um controlador if para retornar o erro caso a lista esteja vazia.

		return 0, fmt.Errorf("a lista está vazia") // Retornamos o valor 0 e denominamos o erro como: "a list está vazia".

	} // Fechamos o if.

	soma := 0 // Declaramos uma variável "soma" com o valor inicial de 0.

	for _, x := range list { // Puxamos um loop for, que tem o seguinte papel: para cada item da lista criada, não interessando que item seja esse, declare a variável "x" e atualize o valor atribuído para o valor do elemento atual da lista. Faça isso até o último valor da lista.

		soma += x // Atribuímos à variável "soma", a soma de cada elemento da lista.

	} // Fechamos o for.

	return soma / len(list), nil // Retornamos o valor da soma de cada item divido pela quantidade de intens contidos na lista e um valor "nil" ou vazio.

} // Fechamos a função "m".

func main() { // Criamos a função "main", que não possui parâmetros.

	result, err := m([]int{1, 2, 3, 4, 5}) // Declaramos as varáveis "result" e "err" e chamamos a função "m", atribuindo os valores dos elementos da lista, para que eles sejam armazenados nas varáveis.

	if err != nil { // Puxamos um outro if para mostrar na tela do usuário o erro declarado na função anterior, caso acorra algum erro.

		fmt.Printf("Houve um erro ao executar o programa: %s.\n", err) // Caso a varável "err" seja diferente de "nil" ou vazio, esta linha mostrará para usuário que algum erro ocorreu e que erro foi esse.

	} // Fechamos o if.

	fmt.Printf("A média é : %d.", result) // Caso o if não seja atendido, o resultado da operação será "impresso" na tela do usuário.

} // Fechamos a função "main".
