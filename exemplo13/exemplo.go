package main

import "fmt"

// struct
// como comeca uma struct?
// type, nome da estrutura, e struct
// todo atributo que comeca com uma letra maiuscula, ele é publico, ele ta exportado
// atributos privados sao inicializados com letras minusculas, como essa struct esta dentro do packge main, pode ser acessado seus atributos privados
// caso esteja em outro packge, os atributos privados nao podem ser acessados
type Pessoa struct {
	Nome string
	Sobrenome string
	Idade uint8
	Status bool
}
// este metodo String já exite no go
func (p Pessoa) String() string {
	return fmt.Sprintf("Ola, meu nome e: %s e eu tenho %d anos", p.Nome, p.Idade) // Sprintf() basicamente formata uma string e retorna 
}

// esta struct esta referenciando a si mesma, por isso o atributo pai recebe *Categoria
type Categoria struct {
	Nome string
	Pai *Categoria
}

// metodos, sao basicamente funcoes que estao atreladas a um struct.
// func para declarar uma funcao, (c Categoria), para dizer que esta sendo atrelada á struct Categoria, HasParent() nome da funcao, bool retorno da funcao
// desta forma, eu passo a struct como valor
func (c Categoria) HasParent() bool {
	return c.Pai != nil
}



func main() {
	// slice
	// isto é um slice, ele tem tamanho dinamico, onde nao é declarado um valor para o seu temanho
	nomes := []string{"Rodrigo", "Kelven", "Rael","Santana", "Juan", "Macedo", "Tiago", "Jesus", "Wesley", "Nicolas", "Marcusi"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	// funcao appende que adiciona elemento no slice
	nomes = append(nomes, "Rafael")
	// funcoes que mostram o tamanho e capacidade da noss lista/slice
	// len() -> tamanho do array
	// cap() -> capacidade do array
	fmt.Println(nomes, len(nomes), cap(nomes))

	// usando make
	// nome := make([]string, 10, 20)

	// maps
	// muito simples, sao chave e valor, lembra um dictionary do python
	// key , value
	// key vai dentro dos colchetes
	// value ao lado do colchetes
	idade := make(map[string]uint8)
	idade["Rodrigo"] = 19
	idade["Rael"] = 24
	idade["Juan"] = 18

	// recuperando o valor de acordo com a chave
	fmt.Println(idade["Rael"])
	idadeRael := idade["Rael"]
	fmt.Printf("A idade de Rael é: %d\n", idadeRael)
	fmt.Println(idade)

	// buscando valor que nao exite
	fmt.Println(idade["Lucas"])
	// retorna 0, mas porque?, ao passar o indice é passado valores de retorno
	// ao retornar, dois parametros sao passados, o valor -> val, e o que é chamado de ok -> ok
	// exemplo
	fmt.Println("Valores que nao exite, teste")
	val, ok := idade["lucas"]
	fmt.Println(val, ok)

	// struct
	fmt.Println("")
	fmt.Println("Struct")
	// esta é a melhor forma de fazer structs, ja que caso tenha alguma mudanca, será de facil mudanca
	p := Pessoa{
		Nome: "Kelve",
		Sobrenome: "Kelven",
		Idade: 10,
		Status: true,
	}
	// retornando todos os dados instanciados da struct referenciado em p
	fmt.Println(p)

	// retornando somente idade e nome
	fmt.Println(p.Idade, p.Nome)

	// usando o metodo String atrelado a struct Pessoa
	fmt.Println(p)

	// segunda forma de fazer, nao recomendado, colocar os dados na ordem que estao instanciados
	p2 := Pessoa{"Rodrigo", "Kelven", 19, true}

	fmt.Println(p2)
	// retornando somente idade s nome
	fmt.Println(p2.Idade, p2.Nome)

	cate := Categoria{Nome: "Categoria 1"}
	if !cate.HasParent() {
		fmt.Println("Nao tem pai")
	}


}