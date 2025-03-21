package main

import "fmt"

func main() {
	fmt.Print(("Digite um numero inteiro: "))
	var numero int64
	fmt.Scanf("%d", &numero) // o numero inteiro digitado será armazenado no espaco de memoria alocado 'numero'
	fmt.Printf("Numero digitado: %d \n", numero)

	// for incremento
	// no for, a linguagem go realiza um 'cast' no tipo de dado armazenado, ele entende que cada processador tem uma arquitetura
	for x := 0; x <= int(numero); x++ {
		fmt.Printf("%d ", x)
		if x == 10 {
			fmt.Print("Seila ")
		}
	}
	fmt.Println("")
	// for decremento
	for y := numero; y >= 0; y-- {
		fmt.Printf("%d ", y)
	}

	fmt.Println("")

	// segunda forma de usar o for
	// usando  um range no slice string para percorrer os valores
	// e declarando e atribuindo a k o range, percorrendo o slice
	// caso queria omitir a saida do range, posso usar _
	nomes := []string{"Joao", "Lucas", "Carlos", "Marcos"} // isto é um slice, o que é um slice? é basicamente um array sem declaracao de tamanho, por isso nao foi passado nenhum valor entre []
	// omitindo a saida do range com _, e retirando o k do print
	// jogando assim o valor dessa variavel no lixo
	for _, nomes := range nomes {
		fmt.Println(nomes)
	}

	// for infinito
	//for z:=0; ;z++{
	//	fmt.Printf("%d ", z)
	//}

	// estrutura for, ' o for nao nescessariamente precisa ter toda a estrutura para funcionar, so se atente para nao ter um for infinito!'
	// declaracao e inicializacao da variavel,
	// condicao do for,
	// exeibicao com o print,
	// incremento ou decremento
	var i int
	for i < 10 {
		fmt.Println("Estrutura do for")
		i++
	}

}
