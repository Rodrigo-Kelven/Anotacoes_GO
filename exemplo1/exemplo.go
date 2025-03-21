package main

// o main sera a funcao principal, ela sera executda primeiro ao rodar o programa

// fmt é um pacote que contem diversas funcoes pra diversas coisas
import "fmt"

// func é uma palavra reservada
// main é a funcao inicial/principal
func main() {
	// println é uma funcao da library fmt

	// println, o ln significa uma quebra de linha, igual ao java
	fmt.Println("Hello World!")
	fmt.Println("Pulou uma linha!")

	// sem quebra de linha
	//fmt.Print("Sem quebra de linha!")
	//fmt.Print("Sem quebra de linha com caractere especial de quebra de linha! \n")
	//fmt.Println("\tCom tabulacao")

	fmt.Println("###########################")
	var nome string

	println(("Digite seu nome:"))
	fmt.Scanf("%s\n", &nome)

	fmt.Printf("Ola: %s\n", nome)
}