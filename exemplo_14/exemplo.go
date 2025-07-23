package main

import "fmt"

// struct criada com atributos publicos
type Pessoa struct {
	Nome string
	Cpf string
	Idade uint8
}

// struct que herda struct Pessoa
type PessoaRegister struct {
	Pessoa
	password string
	Email string
	Status bool
}


func main() {
	// instanciacao da struct PessoaRegister
	p := PessoaRegister{
		Pessoa{
			Nome: "Rodrigo",
			Cpf: "000.000.000-00",
			Idade: 19,
		},
		"seilamermao",
		"email@email.com",
		true,

	}

	// exibindo struct
	fmt.Println("Struct PessoaRegister",p)

	// somente o struct Pessoa
	
	p1 := Pessoa{
		Nome: "Kelven",
		Cpf: "000.000.000-00",
		Idade: 19,
	}
	// exibindo todos os dados instanciados no struct Pessoa
	fmt.Println("Todas as informacoes setadas de p1",p1)

	// exibindo somente o cpf de p1
	fmt.Println("Somente o nome de p1:", p1.Nome)
	// exibindo somente o cpf de p1
	fmt.Println("Somente o cpf de p1:", p1.Cpf)
	// exibindo somente o cpf de p1
	fmt.Println("Somente o idade de p1:", p1.Idade)

	

}