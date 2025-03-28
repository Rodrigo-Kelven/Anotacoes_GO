package main

import (
	"fmt"
)

// 1 operadores uniario (++ e --)
// 2 modulo, divisao e multiplicacao (%, / e *)
// 3 soma e subtracao (+ e -)

// a linguagem Go é fortemente tipada, onde é nescesario passar o tipo do da variavel dentro da funcao,
// tanto na entrada tando na saida da funcao, 'no retorno'
func calcularSoma(a int, b int) int {
    return a + b
}
// int fora dos parenteses significa que sera retornado um valor do tipo int
func calcularMultiplicacao(a int, b int) int { 
    return a * b
}
// funcao para calcular a subtracao de dois parametros do tipo int passado para a funcao, onde o retorno dos mesmos seram do tipo int
func calcularSubtracao(a int, b int) int {
    return a - b
}

// nesta funcao, ela recebe um unico parametro do tipo int e retorna dois parametros
// em go, para informar que será retornado dois parametros, é declarado entre () os tipos de dados retornado
// como nesta funcao, dois parametros estao sendo retornados, (int, int) foi declarado como tipos dos dados de saida 
func elevado(a int) (int, int) {
    var quadrado int = a * a
    var cubo int = a * a * a

    return quadrado, cubo
}
// desta forma, eu posso nomear e especificar o parametro e seu tipo de retorno, por exemplo
// (quadrado int, cubo int) significa que essas variaveis/parametros de retorno sao do declaradamente do tipo int
// assim eu posso retirar o var, o tipo da variavel e as variaveis no return, ja que esta explicitamente declarado no retorno da funcao
// ESSE TIPO DE FUNCAO COM VALORES DE RETORNO VARIAVEL SO DEVE SER USADA EM FUNCOES CURTAS, POIS O USO DAS MESMAS EM FUNCOES GRANDES PODEM ACARRETAR 
// EM DIFICULDADE DE LEITURA E ENTENDIMENTO DO CODIGO
func elevado2(a int) (quadrado int, cubo int) {
    quadrado = a * a
    cubo = a * a * a

    return
}

func main() {
    resultado1 := 5
    resultado2 := 7

	resultado1 = resultado1 + 1 // resultado1 +1 ou resultado1++
	resultado2 = resultado2 - 1 // resultado2 -1 ou resultado2--

    fmt.Printf("Resultado 1: %d\n", resultado1)
    fmt.Printf("Resultado 2: %d\n", resultado2)

    resultado1 = calcularSoma(5, 2)
    resultado2 = calcularMultiplicacao(5, 2)
    var resultado3 = calcularSubtracao(5, 2)

    fmt.Printf("Soma: %d\n", resultado1)
    fmt.Printf("Multiplicação: %d\n", resultado2)
    fmt.Printf("Subtracao de %d\n", resultado3)
    fmt.Println("")
    fmt.Print("Primeira forma usnado a funcao elevado:")
    fmt.Println(elevado(2))
    fmt.Print("Segunda forma usando a funcao elevado:")
    fmt.Println(elevado2(2))
}
