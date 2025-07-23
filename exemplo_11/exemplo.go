package main

import "fmt"

func main() {
	// arrays tem tamanho fixo
	// arrays sao homogeneo = seus itens/valores devem ter o mesmo tipo
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
	// numeros2 = [5]int{1,2,3,4,5} // declarando um array de forma normal
	// numeros2 = []int{1,2,3,4,5} // declarando um array de forma normal sem o tamanho, go entende e declara o tamanho do array pelo numero de elementos passados
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

	// quarta forma
	// variavel numeros4 de tipo int com tamanho 8
	var numeros4 [8]int
	fmt.Println(numeros4) // o return é [0 0 0 0 0 0 0 0] porque o array foi alocado sem valor passados, ou seja, zeros 
	// agora o for ira percorrer o array e alocar valores de acordo que os indices do for sao percorridos
	for indice:=0; indice <8; indice++{
		numeros4[indice]=indice
	}
	fmt.Println(numeros4)
	fmt.Printf("Tipo de array: %T\n", numeros4) // [8]int é o tamanho, o tamanho do array tambem faz parte do tipo do array

	// pegando partes de array, 0 - 2
	fmt.Println("Do indice 0 ate 2", numeros4[0:3]) // o ultimo indice é exclusivo, entao sempre pege um a mais
	fmt.Println("Do indice 0 ate 2", numeros4[:3]) // pegando do indice 0 ate o indice 2, o ultimo é exclusivo
	// indices de 2 - 5
	fmt.Println(numeros4[2:5])
	fmt.Println(numeros4[2:]) // pegando do indice 2 ate o ultimo,
}