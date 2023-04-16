// Escreva uma função que receba um slice de inteiros como parâmetro e retorne o menor valor contido no slice. Caso o slice esteja vazio, retorne um erro.

package main // Chamamos o pacote que contém as bibliotecas que usaremos. Nesse caso, o pacote "main".

import "fmt" // Importamos a biblioteca "fmt" do pacote, pois ela tem as funções que usaremos.

func slice(nums []int) (int, error) { // Criamos a função "slice", que vai receber um slice de números inteiros. Ela vai retornar um número do tipo inteiro, e um erro, caso ele ocorra.

	if len(nums) == 0 { // Puxamos um controlador if para retornar o erro caso o slice esteja vazio.

		return 0, fmt.Errorf("o Slice está vazio") // Retornamos o valor 0 e denominamos o erro como: "o Slice está vazio".

	} // Fechamos o if.

	menor := nums[0] // Declaramos a varável "menor" com valor inicial igual ao do primeiro item do Slice.

	for i := 1; i < len(nums); i++ { // Criamos um loop for que vai passar por cada item do Slice e atribuir o respectivo valor à variável "i".

		if nums[i] < menor { // Puxamos um outro if para comparar o valor do item atual com o anterior.

			menor = nums[i] // E, caso o valor atual seja maior que o valor anterior, o menor valor é atribuído à variável "menor".

		} // Fechamos o if.

	} // Fechamos o for.

	return menor, nil // Retornamos a varável menor, com o menor valor e retornamos também o "nil", conhecido também como vazio.

} // Fechamos a função "slice".

func main() { // Criamos a função "main", que não possui parâmetros.

	result, err := slice([]int{1, 2, 3, 4, 5}) // Declaramos as varáveis "result" e "err" e chamamos a função "slice", atribuindo os valores dos elementos da lista, para que eles sejam armazenados nas varáveis.

	if err != nil { // Puxamos um outro if para mostrar na tela do usuário o erro declarado na função anterior, caso acorra algum erro.

		fmt.Printf("Houve um erro ao executar o programa: %s.", err) // Caso a varável "err" seja diferente de "nil" ou vazio, esta linha mostrará para usuário que algum erro ocorreu e que erro foi esse.

	} // Fechamos o if.

	fmt.Printf("O menor número do Slice é: %d.", result) // Caso o if não seja atendido, o resultado da operação será "impresso" na tela do usuário.

} // Fechamos a função "main".
