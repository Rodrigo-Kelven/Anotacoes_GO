package main

import (
	"fmt"
)

func main() {
	
	fmt.Println("Digite seu nome:")
	var nome string
	fmt.Scanf("%s",&nome) // o & significa que esta sendo alocado e armazenado um valor do tipo string na variavel nome

	fmt.Println("Digite um numero:")
	var numero1 float64
	fmt.Scanf("%f", &numero1) // o & significa que esta sendo alocado e armazenado um valor do tipo float64 na variavel numero1

	fmt.Println("Digite outro numero:")
	var numero2 float64
	fmt.Scanf("%f", &numero2) // o & significa que esta sendo alocado e armazenado um valor do tipo float64 na variavel numero2

	soma := numero1+numero2 // a variavel soma esta sendo declarada e inicalizada com o :=

	fmt.Printf("Ola %s, seja bem vindo ao Onion Chain\n", nome)
	
	fmt.Printf("A soma de %.2f + %.2f é %.2f\n", numero1,numero2,soma)

}