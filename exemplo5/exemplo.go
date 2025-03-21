package main
// package main indica que esta sendo criada uma funcao main, ela será a funcao principal do sistema

// imports para usar as funcoes especificas
import (
	"fmt"
	"math"
)

// funccao principal
func main() {

	// o Println indica que será pulada uma linha automaticamente
	fmt.Println("Digite uma palvra:")
	var palavra string // declarando a variavel palavra como string, (var pode ser passado opcionalmente para indicar uma variavel)
	
	// a funcao Scanf esta dentro do pacote fmt, por isso deve ser importado, ela aloca um espaco de memoria para armazenar o valor que esta sendo passado, e armazena na variavel palavra
	fmt.Scanf("%s", &palavra) 

	// desta forma, nao precisa declarar a variavel antes, o := indica declaracao e inicializacao da variavel
	tamanho := len(palavra)  // mas so pode usar := na primeira vez, ja que voce esta declarando, significa que ja esta declarado e inicializado

	metade := tamanho / 2 // mesmo esquema da variavel tamanho

	primeiraParte := palavra[0:metade] // mesmo esquema, declaracao e inicializacao 

	fmt.Printf("Primeira metade: %s\n", primeiraParte) // fmt.Printf("Primeira metade: %s\n", palavra[0:metade])
	fmt.Printf("Segunda metade: %s\n", palavra[metade:])

	// Fazendo IMC

	fmt.Println("Digite o seu peso:")
	var peso float64
	// o valor que esta sendo passado, a funcao scanf esta esperando o valor ser passado, para alocar no espaco de memoria que ja foi 'definido'
	fmt.Scanf("%f\n", &peso)

	fmt.Println("Digite a sua altura:")
	var altura float64
	fmt.Scanf("%f\n", &altura)

	//imc := peso / (altura * altura)
	imc := peso / math.Pow(altura, 2)

	fmt.Printf("Seu IMC é: %.2f\n", imc)

}
