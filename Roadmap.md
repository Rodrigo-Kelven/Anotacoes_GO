
# 🧠 Roadmap Completo de Go – Do Básico ao Avançado (Intensivo de 2 Semanas)


### 📅 Semana 1 — Fundamentos Sólidos da Linguagem Go

#### Dia 1 – Fundamentos, Tipos e Variáveis
- Instalação e configuração do Go
- Estrutura básica de um programa (package main, import, func main())
- Comentários
- Variáveis, Tipos e Constantes
    - Tipos primitivos (int, string, bool, float64)
    - var, :=, const
    - Conversão de tipos
- Pacotes e importações
- Escopo de variáveis


#### Dia 2 – Controle de Fluxo e Funções
- Operadores: aritméticos, relacionais, lógicos
- Condicionais: if, else, switch
- Laços: for, range
- Funções
    - Nomeadas, variádicas
    - Parâmetros e múltiplos retornos
    - Funções como valor (first-class functions)
    - Funções anônimas
    - defer, panic, recover


#### Dia 3 – Estruturas de Dados + Recursão
- Arrays e Slices (capacidade, comprimento, ponteiros internos)
- Maps (chave/valor)
- Recursão
- Structs
    - Definição e uso prático
    - Structs aninhadas


#### Dia 4 – Métodos, Ponteiros e Interfaces
- Ponteiros
    - Referência e desreferência
    - Ponteiros em funções
    - Receptores de métodos com ponteiros
    - Indireção de ponteiros
- Métodos em structs
- Interfaces
    - Definição, implementação implícita (duck typing)
    - Type assertions
    - Type switches

#### Dia 5 – Concorrência com Goroutines e Channels
- Goroutines
- Channels (buffered e unbuffered)
- select statement
- sync.WaitGroup, sync.Mutex, sync.Once
- Worker Pool
- context.Context para cancelamento e timeout

#### Dia 6 – Manipulação de Erros + Boas Práticas
- Interface error
- Criando erros personalizados
- Wrapping com fmt.Errorf, errors.Is, errors.As, errors.Unwrap
- Sentinel errors
- Tratamento com defer, panic, recover
- Go idiomático: boas práticas, naming conventions, init(), iota, estrutura de pacotes


#### Dia 7 – Generics + Testes
- Generics (Go 1.18+)
    - Funções e tipos genéricos
    - Constraints (comparable, interfaces)
    - Casos de uso práticos
- Testes com testing package
    - Testes unitários e de integração
    - Table-driven tests
    - go test, benchmark, cobertura
    - t.Errorf() vs assert (testify)


### 📅 Semana 2 — Profundidade Profissional com Go

#### Dia 8 – I/O, Arquivos e Entrada do Usuário
Leitura e escrita de arquivos (os, io, ioutil, bufio)
Entrada de dados (fmt.Scanln, bufio.Scanner, os.Stdin)
Manipulação de diretórios

#### Dia 9 – JSON, HTTP e Web com Go
Leitura e escrita de JSON (encoding/json)
Servidores HTTP com net/http
Handlers personalizados
Middlewares
Routers externos: chi, mux, gin


#### Dia 10 – Banco de Dados com Go
- database/sql e drivers
- Conectando com PostgreSQL, MySQL ou SQLite
- Queries, transações
    - ORMs
    - sqlx (leve)
    - gorm (completo)

#### Dia 11 – Ferramentas Profissionais e Go Idiomático
- go mod, go fmt, go vet, golint, staticcheck
- godoc e documentação idiomática
- go generate
- Boas práticas para código limpo


#### Dia 12 – Build, Deploy e Docker
go build, go install, go run
Compilação cruzada (cross-compilation)
Binaries estáticos
Dockerizando aplicações Go
Compilação para ARM/Linux


#### Dia 13 – Padrões e Arquitetura Profissional
- Padrões de projeto em Go
    - Factory, Singleton, Strategy, Adapter, etc.
- Arquitetura
    - Clean Architecture
    - Hexagonal Architecture
    - DDD aplicado a Go
    - Monolito vs microserviços


#### Dia 14 – Revisão, Projeto Final e Extras
- Revisar todos os conceitos aprendidos
- Criar um projeto final (ex: API RESTful completa com DB + goroutines + - testes)
- Benchmarking e profiling de performance
- Extras (avançado e opcional):
    - CLI apps (cobra, urfave/cli)
    - gRPC (protobuf, grpc-go)
    - Mensageria (Kafka, RabbitMQ, NATS)
    - Logging avançado (zap, logrus)
    - DevOps: Prometheus, Grafana


### ✅ Resumo da Ordem Recomendada (por conceito)
- Fundamentos (variáveis, tipos, fluxo)
- Funções → variádicas → closures → recursão
- Ponteiros e métodos
- Structs + métodos
- Erros e tratamento avançado
- Goroutines, channels e concorrência
- Interfaces, type assertions/switches, generics
- I/O, JSON, entrada de usuário
- HTTP, DB, testes, deploy, arquitetura
