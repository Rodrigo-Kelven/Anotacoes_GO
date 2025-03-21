package main

import "fmt"

func main() {
	// variavel numeros com tamanho 5 do tipo int | array
	// desta forma, esta sendo declarado um array homogenio do tipo int com tamanho 5
	// logo abaixo os indices do array com seus respectivos valores
	var numeros [5]int 
	numeros[0] = 1
	numeros[1] = 2
	numeros[2] = 3
	numeros[3] = 4
	numeros[4] = 5

	// imprime o array
	fmt.Println(numeros) 


	// segunda forma de criar e declarar array
	// desta forma, os valores ja sao armazenados de acordo com o indice/posicao, que esta sendo declarado
	numeros2 := [5]int{1,2,3,4,5} 

	// imprime o array
	fmt.Println(numeros2)

	// terceira forma
	// se o tamanho do array for maior que a quantidde de indices preenchidos, OK
	// senao, será identificado como erro de forma automatica e o codigo nao será executado
	numeros3 := [6]int{1,2,3,4,5,6}

	// for para realizar soma dos valores preenchidos no array
	soma := 0
	for i :=0; i < len(numeros3); i++{
		soma += numeros3[i] // a variavel recebe e soma o valor referente ao indice do array numeros3
	}
	fmt.Printf("Somatorio %d\n", soma)
}