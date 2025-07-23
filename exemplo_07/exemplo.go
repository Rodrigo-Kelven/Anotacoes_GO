package main

import (
	"fmt"
)

func main() {
	fmt.Println("Digite um numero:")
	var numero1 int 
	// o & significa um poteiro unitario, onde o valor que esta sendo passado sera armazenado na variavel no tipo int, onde memroia foi alocada
	fmt.Scanf("%d\n", &numero1) 

	fmt.Println("Digite outro numero:")
	var numero2 int
	fmt.Scanf("%d\n", &numero2)

	// operacoes matematicas com variaveis declaradas a inicalizadas
	soma := numero1 + numero2
	subtracao := numero1 - numero2
	mult := numero1 * numero2
	div := numero1/numero2
	fmt.Printf("A soma de %d + %d é %d\n", numero1,numero2,soma)
	fmt.Printf("A subtracao de %d - %d é %d\n", numero1,numero2,subtracao)
	fmt.Printf("A multiplicacao de %d * %d é %d\n", numero1,numero2,mult)
	fmt.Printf("A divisao de %d / %d é %d\n", numero1,numero2,div)
}