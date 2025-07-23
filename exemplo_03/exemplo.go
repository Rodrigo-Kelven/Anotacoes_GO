package main

import (
	"fmt"
	"strings"
)

// strings


func main() {
	str_test := ""
	fmt.Println(str_test)

	str_test = "ok"
	fmt.Println(str_test)

	str_test = "Ola mundo"
	fmt.Println(str_test)

	
	// %d -> numero inteiro
	// %f -> numero do tipo float
	// %.2f -> numero do tipo float com duas casas
	// %d -> numero do tipo double
	// %c -> sera impresso um caracter
	// %s -> sera impresso uma string


	// exemplo 2

	// declarando, inicializando e passando valor para uma variavel
	frase := "Ola, mundo!"

	primeiroCaractere := frase[0]
	//fmt.Printf("Primeiro caractere em ASCI %s\n", primeiroCaractere) // desta forma, será impresso o numero asci da tabela, 79
	fmt.Printf("Primeiro caractere: %c\n", primeiroCaractere) // por isso deve-se ser passado a instrucao que será passado um caractere %c

	indiceultimoCaractere := len(frase) - 1 // pegamos o valor do ultimo indice da frase

	ultimoCaractere := frase[indiceultimoCaractere] // pego o valor do ultimo indice e armazeno como caractere na variavel

	fmt.Printf("Ultimo caractere: %c\n", ultimoCaractere) // passa o valor do indice como caracter

	fmt.Println("###############################")


	// exemplo 3

	frase_2 := "Aprendendo Go"

	// tamanho da string
	tamanho_frase := len(frase_2)
	fmt.Printf("Tamanho da frase_2: %d caracteres.\n", tamanho_frase)

	// pegando a primeira palavra
	primeiraPalavra := frase_2[0:10] // o primeiro indice é inclusivo, o ultimo é exclusivo! ou seja, o ultimo indice sempre é excluso, entao pega 1 a mais.

	fmt.Printf("Primeira palavra: %s\n", primeiraPalavra)

	// pegando a segunda palavra
	segundaPalavra := frase_2[11:] // o primeiro indice é inclusivo, quando nao é passado valor, vai ate o final da frase/indice

	fmt.Printf("Segunda palavra: %s\n", segundaPalavra)


	fmt.Println("#########################")
	// exemplo 4
	// convertendo valores em string

	frase_3 := "Go é lindo demais mermao"

	// o Printf nao pula linha, por isso usa-se \n
	fmt.Printf("Frase original: %s\n", frase_3)

	// a variavel esta sendo declarada e inicializada por := 
	fraseMinuscula := strings.ToLower(frase_3) // a classe ToLower esta dentro do pacote strings, por isso esta sendo importada
	fmt.Printf("Frase minuscula: %s\n", fraseMinuscula)

	fraseMaiuscula := strings.ToUpper(frase_3)
	fmt.Printf("Frase Maiuscula: %s\n", fraseMaiuscula)

	// exemplo 5

	outraFrase := "   frase com espaco   "

	tamanho_outraFrase := len(outraFrase)// pegando o tamanho da variavel com os espacos

	fmt.Printf("Tamanho da variavel outraFrase com espacos: %d\n", tamanho_outraFrase)

	outraFrase = strings.Trim(outraFrase, " ") // a funcao Trim recebe dois parametros, a frase, e o que sera retirado, no caso os espacos

	tamanho_outraFrase = len(outraFrase)

	fmt.Printf("Tamanho da variavel outraFrase sem espacos: %d\n", tamanho_outraFrase)

	fmt.Println("#########################")

	// outra forma

	seila := "   seila mermao   "

	tamanho := len(seila) // aqui a variavel tamanho esta sendo declarada e inicializada, e contara os espacoes como tamanho

	fmt.Printf("Tamanho da frase com espacos: %d\n", tamanho)

	// sem espacos

	seila_dois := strings.Trim(seila, " ") // esta classe Trim remove os espacos, ela esat dentro de strings, lembra o PHP

	seila_dois_tamanho := len(seila_dois)

	fmt.Printf("Tamanho da frase sem espaco: %d\n", seila_dois_tamanho) // Printf nao pula linha, %d significa que será impresso um valor do tipo inteiro
}