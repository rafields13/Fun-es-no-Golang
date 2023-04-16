// Crie uma função que receba dois parâmetros (a e b) e retorne a divisão entre eles. Caso o segundo parâmetro seja zero, retorne um erro.

package main // Chamamos o pacote que contém as bibliotecas que usaremos. Nesse caso, o pacote "main".

import "fmt" // Importamos a biblioteca "fmt" do pacote, pois ela tem as funções que usaremos.

func divide(a float64, b float64) (float64, error) { // Criamos a função "divide", a qual retornará um valor do tipo float e um erro, caso aconteça, do tipo error. Esta função vai dividir o primeiro parâmetro pelo segundo.

	if b == 0 { // Puxamos um controlador if para retornar o erro caso o segundo parâmetro seja igual a 0.

		return 0, fmt.Errorf("não pode dividir por 0") // Retornamos o valor 0 e denominamos o erro como: "não pode dividir por 0".

	} // Fechamos o if.

	return a / b, nil // Caso o if não seja atendido, a função retornará a operação "a / b" e o "nil", que é o vazio.

} // Fechamos a função "divide".

func main() { // Criamos a função "main", que não possui parâmetros.

	result, err := divide(21, 7) // Declaramos as varáveis "result" e "err" e chamamos a função "divide" com os valores dos parâmetros, para que eles sejam armazenados nas varáveis.

	if err != nil { // Puxamos um outro if para mostrar na tela do usuário o erro declarado na função anterior, caso acorra algum erro.

		fmt.Printf("Ocorreu um erro ao dividir a e b: %s.\n", err) // Caso a varável "err" seja diferente de "nil" ou vazio, esta linha mostrará para usuário que algum erro ocorreu e que erro foi esse.

	} // Fechamos o if.

	fmt.Printf("O resultado dessa divisão é: %.1f.", result) // Caso o if não seja atendido, o resultado da operação será "impresso" na tela do usuário.

} // Fechamos a função "main".
