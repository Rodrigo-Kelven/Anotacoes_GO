package main

import "fmt"


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


func main() {
    resultado1 := 5
    resultado2 := 7

	resultado1 = resultado1 + 1 // resultado1 +1 ou resultado1++
	resultado2 = resultado2 - 1 // resultado2 -1 ou resultado2--

    fmt.Printf("Resultado 1: %d\n", resultado1)
    fmt.Printf("Resultado 2: %d\n", resultado2)

    resultado1 = calcularSoma(5, 2)
    resultado2 = calcularMultiplicacao(5, 2)
    fmt.Printf("Soma: %d\n", resultado1)
    fmt.Printf("Multiplicação: %d\n", resultado2)
}
