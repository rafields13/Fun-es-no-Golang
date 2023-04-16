// Crie uma função que receba um número inteiro como parâmetro e retorne a sequência de Fibonacci correspondente a esse número. Caso o número seja negativo, retorne um erro.

package main // Chamamos o pacote que contém as bibliotecas que usaremos. Nesse caso, o pacote "main".

import "fmt" // Importamos a biblioteca "fmt" do pacote, pois ela tem as funções que usaremos.

func num(x int) ([]int, error) { // Criamos a função "num", que receberá um valor inteiro e retornará outros valores inteiros, caso dê erro, retornará um erro do tipo error.

	if x < 0 { // Puxamos um controlador if para retornar o erro caso o valor atribuído à variável "x" seja menor que 0, ou seja, um número negativo.

		return nil, fmt.Errorf("o número não pode ser negativo") // Retornamos o "nil", ou vazio, e denominamos o erro como: "o número não pode ser negativo".

	} // Fechamos o if.

	sequence := []int{0, 1} // Declaramos uma lista "sequence", contendo dois números inteiros, 0 e 1.

	for i := 2; i < x; i++ { // Criamos um loop for que vai ler a partir do próximo elemento da lista, até um número que seja meno ou igual ao valor da varável "x".

		next := sequence[i-1] + sequence[i-2] // Declaramos uma variável "next", que representa o próximo elemento da lista, que será a soma dos dois números anteriores à ela.

		sequence = append(sequence, next) // Depois disso, adicionamos esse novo valor à nossa lista de números do Fibonacci.

	}

	return sequence, nil // Quando pronta, retornamos essa sequência Fibonacci e retornamos também o "nil", ou vazio.

}

func main() { // Criamos a função "main", que não possui parâmetros.

	result, err := num(10) // Declaramos as varáveis "result" e "err" e chamamos a função "num", atribuindo o valor correspondente da sequência Fibonacci para que ele seja armazenado nas varáveis.

	if err != nil { // Puxamos um outro if para mostrar na tela do usuário o erro declarado na função anterior, caso acorra algum erro.

		fmt.Printf("Houve um erro ao executar o programa: %s.", err) // Caso a varável "err" seja diferente de "nil" ou vazio, esta linha mostrará para usuário que algum erro ocorreu e que erro foi esse.

	} //Fechamos o if

	fmt.Printf("A sequência Fibonacci correspondente a esse número é: %d.", result) // Caso o if não seja atendido, o resultado da operação será "impresso" na tela do usuário.

} // Fechamos a função "main".
