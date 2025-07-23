package main

import "fmt"

// trabalhando com numeros

func main() {

	// inteiros
	// uint32 --> 0 a 4.294.967.295  | para valores que sejam somente positivos com 32 bits
	// uint64 --> 0 a 18.446.744.073.709.551.615 | para valores que sejam somente positivos com 64 bits
	// int32 --> -2.147.483.648 a 2.147.483.647
	// int64 --> -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807

	// Go armazena o tipo da variavel de acordo com a arquitetura do procesador !!!
	var quantidade uint = 34
	var temperatura int = -3

	fmt.Printf("Quantidade: %d\n", quantidade)
	fmt.Printf("Temperatura: %d\n", temperatura)


	// float
	// float32 --> -3.4e+38 a 3.4e+38.
	// float64 --> -1.7e+308 a 1.7e+308. 
	var peso float64 = 1.5

	fmt.Printf("Peso: %.2f Kg\n", peso) // %.0f para arrendondar o valor tanto pra cima tanto pra baixo
}