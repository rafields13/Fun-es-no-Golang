// Crie uma função que receba uma string como parâmetro e retorne o número de caracteres contidos nessa string. Caso a string seja vazia, retorne um erro.

package main // Chamamos o pacote que contém as bibliotecas que usaremos. Nesse caso, o pacote "main".

import (
	"fmt"
	"strings"
) // Importamos as bibliotecas "fmt" e "strings" do pacote, pois elas têm as funções que usaremos.

func c(s string) (int, error) { // Criamos a função "c", que vai receber uma string e vai retornar um valor do tipo inteiro e um erro, caso aconteça, do tipo error.

	if strings.Count(s, "") == 0 { // Puxamos um controlador if para retornar o erro caso a string esteja vazia.

		return 0, fmt.Errorf("não tem nenhum texto para ser lido") // Retornamos o valor 0 e declaramos o erro como: "não tem nenhum texto para ser lido".

	} // Fechamos o if.

	return strings.Count(s, "") - 1, nil // Retornamos a quantidade de espaços curtos que existem entre cada caractere. Porém, subtraímos 1 ao valor final, porque um espaço curto repete, dessa forma, temos que tirar um da contagem. E retornamos o "nil" ou vazio.

}

func main() { // Criamos a função "main", que não possui parâmetros.

	result, err := c("Hello, world!") // Declaramos as varáveis "result" e "err" e chamamos a função "c", atribuindo o texto que será armazenado nessas variáveis.

	if err != nil { // Puxamos um outro if para mostrar na tela do usuário o erro declarado na função anterior, caso acorra algum erro.

		fmt.Printf("Houve um erro ao executar o programa: %s.\n", err) // Caso a varável "err" seja diferente de "nil" ou vazio, esta linha mostrará para usuário que algum erro ocorreu e que erro foi esse.

	} // Fechamos o if.

	fmt.Printf("A quantidade de caracteres desse texto é: %d.", result) // Caso o if não seja atendido, o resultado da operação será "impresso" na tela do usuário.

} // Fechamos a função "main".
