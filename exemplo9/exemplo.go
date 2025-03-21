package main

import "fmt"

func main() {
	// Operadores lógicos
	var valor1 bool = true
	var valor2 bool = false

	resultado := valor1 && valor2
	fmt.Printf("%t && %t = %t\n", valor1, valor2, resultado)

	numero1 := 3
	numero2 := 5

	resultado1 := numero1 == numero2
	resultado2 := numero1 > numero2
	resultado3 := numero1 < numero2
	resultado4 := numero1 >= numero2
	resultado5 := numero1 <= numero2
	resultado6 := numero1 != numero2

	fmt.Printf("Resultado 1: %t\n", resultado1)
	fmt.Printf("Resultado 2: %t\n", resultado2)
	fmt.Printf("Resultado 3: %t\n", resultado3)
	fmt.Printf("Resultado 4: %t\n", resultado4)
	fmt.Printf("Resultado 5: %t\n", resultado5)
	fmt.Printf("Resultado 6: %t\n", resultado6)

	// Entrada de números
	fmt.Println("Digite um número:")
	var um_numero float64
	fmt.Scanf("%f\n", &um_numero)

	fmt.Println("Digite outro número:")
	var outro_numero float64
	fmt.Scanf("%f\n", &outro_numero)

	var resultado_calc float64

	// Escolha da operação
	fmt.Println("Escolha uma operação:")
	fmt.Println("Soma: 1")
	fmt.Println("Subtração: 2")
	fmt.Println("Divisão: 3")
	fmt.Println("Multiplicação: 4")
	var opcao string
	fmt.Scanf("%s", &opcao)

	// Usando if para calcular
	if opcao == "1" {
		resultado_calc = um_numero + outro_numero
		fmt.Printf("A soma de %.2f + %.2f é %.2f\n", um_numero, outro_numero, resultado_calc)
	} else if opcao == "2" {
		resultado_calc = um_numero - outro_numero
		fmt.Printf("A subtração de %.2f - %.2f é %.2f\n", um_numero, outro_numero, resultado_calc)
	} else if opcao == "3" {
		if outro_numero != 0 {
			resultado_calc = um_numero / outro_numero
			fmt.Printf("A divisão de %.2f / %.2f é %.2f\n", um_numero, outro_numero, resultado_calc)
		} else {
			fmt.Println("Erro: Divisão por zero não é permitida.")
		}
	} else if opcao == "4" {
		resultado_calc = um_numero * outro_numero
		fmt.Printf("A multiplicação de %.2f * %.2f é %.2f\n", um_numero, outro_numero, resultado_calc)
	} else {
		fmt.Println("Opção inválida.")
	}

	// Usando switch para calcular
	fmt.Println()
	fmt.Println("Digite um número:")
	var umm_numero float64
	fmt.Scanf("%f\n", &umm_numero)

	fmt.Println("Digite outro número:")
	var outroo_numero float64
	fmt.Scanf("%f\n", &outroo_numero)

	var resultado_calcu float64

	fmt.Println("Escolha uma operação:")
	fmt.Println("Soma: 1")
	fmt.Println("Subtração: 2")
	fmt.Println("Divisão: 3")
	fmt.Println("Multiplicação: 4")
	var oopcao string
	fmt.Scanf("%s", &oopcao)

	// usando switch case, fica mais limpo e de facio entendimento
	// usando funcoes para cada operacao, ficaria muito mais limpo.
	switch oopcao {
	case "1":
		resultado_calcu = umm_numero + outroo_numero
		fmt.Printf("A soma de %.2f + %.2f é %.2f\n", umm_numero, outroo_numero, resultado_calcu)
	case "2":
		resultado_calcu = umm_numero - outroo_numero
		fmt.Printf("A subtração de %.2f - %.2f é %.2f\n", umm_numero, outroo_numero, resultado_calcu)
	case "3":
		if outroo_numero != 0 {
			resultado_calcu = umm_numero / outroo_numero
			fmt.Printf("A divisão de %.2f / %.2f é %.2f\n", umm_numero, outroo_numero, resultado_calcu)
		} else {
			fmt.Println("Erro : Divisão por zero não é permitida.")
		}
	case "4":
		resultado_calcu = umm_numero * outroo_numero
		fmt.Printf("A multiplicação de %.2f * %.2f é %.2f\n", umm_numero, outroo_numero, resultado_calcu)
	default:
		fmt.Println("Opção inválida.")
	}
}
