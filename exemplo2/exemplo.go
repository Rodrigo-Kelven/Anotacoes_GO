package main

import "fmt"

func main() {
	// outra forma de declarar variaveis em conjunto
	var (
		username string
		full_name = "Rodrigo Kelven"
	)
	// desta forma, a variavelteste esta sendo declarada e inicializada com valor/conteudo "ok"
	teste := "ok"
	username = "Kelven" // esta varivael ja foi declarada, agora esta sendo atribuido uma valor a ela
	
	fmt.Println("Username", username, "Full_Name", full_name, "Test", teste) // impressao das variaveis usando Println

	// outra forma de declarar e atribuir valor á variaveis, percebe-se que os valores serao atribuidos na mesma posicao que as variaveis foram declaradas 
	var b, f, s = true, 2.32, "Ola"

	fmt.Println(b,f,s)

	// trocando os valores de variaveis
	var x, y = 10, 20
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

	fmt.Println("Digite o seu nome:")
	// o dado inserido sera armazenado na memoria ram
	// será armazenado na variavel == espaco na memoria ram

	// variavel, para declarar, coloca a palavra reservada var, que significa dizer que e uma variavel, o nome da varival, e o tipo da variavel
	var nome string


	fmt.Scanf("%s", &nome) // a funcao scanf recebe o valor, e armazena em nome, que é o endereco de memoria
	// o & diz que o valor sera armazenado no endereco de memoria chamado nome

	fmt.Printf("Ola: %s\n", nome) // usar o printf, porque? porque ele é uma saida formatada! formatada com variavel!
	fmt.Println("############################")

	// declarando e inicializando variaveis

	// exemplo 1
	var test string
	test = "Valor teste"
	fmt.Printf("String declara com o valor: %s\n", test)

	// exemplo 2
	// podemos declarar e atribuir valor á variavel, fazendo inicializacao da mesma
	// o operador := ja faz a declaracao e atribuicao da variavel ao mesmo tempo
	// mas so pode ser usado uma unica vez pela variavel, na primeira vez, porque o := faz a declaracao e atribuicao somente na inicializacao

	test2 := "String com variavel declarada!" 
	fmt.Printf("Exemplo e de declaracao de variavel: %s\n", test2)

}