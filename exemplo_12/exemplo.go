package main

import (
	"fmt"
	"strconv"
)

// o go forca a colocar os tipos dos parametros por ser fortemente tipado!
// para que esta funcao seja usada em outros pacotes externos, coloca-se  a primeira letra como maiuscula Hello
func Hello(nome string) {

	fmt.Println("Hello", nome, "!")
}

// passar parametros e seus respectivos tipos,
// como será retornado valores, tem que especificar o tipo de dado retornado apos os parenteses
func soma(a int, b int) int { // tambem pode ser a,b int, go emtende que esta sequencia se parametros será de numeros do tipo inteiro
	return a + b
}

// convert string in inteiro e soma
func convertADDSun(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)

	// nil significa ausencia de erro, ou seja, correto, seguindo uma logica inversa
	// se o erro for diferente de certo, entao esta errado KKKK, complexo antes de entender, engracado após entender
	if err != nil {
		return
	}

	total = a + i
	return
}

func main() {
	fmt.Println("Primeira forma usando a funcao Hello()")
	// primeira forma usando a funcao hello
	Hello("ok") // passando o valor diretamente nela como string

	fmt.Println("Segunda forma usando a funcao Hello()")
	// segunda forma
	// usando uma interacao com o usuario com imput de dado,
	// declarando uma variavel para receber o valor passado,
	// armazenamento do valor no espaco de memoria,
	// passagem do valor da variavel
	fmt.Printf("Digite seu nome: ")
	var nome string
	fmt.Scanf("%s", &nome)
	Hello(nome)

	// primeira forma usando a funcao soma
	fmt.Println("Primeira forma usando a funcao soma()")
	fmt.Println("Total soma da primeira forma:", soma(4, 5))

	fmt.Println("Segunda forma usando a funcao soma()")
	// segunda forma,
	// usando inputs do usuario
	fmt.Print("Digite um numero inteiro: ")
	var num1 int64
	fmt.Scanf("%d", &num1)

	fmt.Print("Digite outro numero inteiro: ")
	var num2 int64
	fmt.Scanf("%d", &num2)

	fmt.Println("Valor da soma da segunda forma:", soma(int(num1), int(num2)))

	fmt.Println("")
	fmt.Println("Usando a funcao para converter valores de string para int!")
	total, err := convertADDSun(10, "23")

	fmt.Println("Total soma2:", total, err)

}
