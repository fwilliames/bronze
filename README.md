# SuperMarket Tracker

SuperMarket Tracker é um aplicativo desenvolvido em Go com interface em Fyne, que permite ao usuário registrar produtos comprados no mercado, armazenando seu valor e a data da compra (mês/ano). Além disso, o programa gera relatórios em PDF com essas informações.

## Tecnologias Utilizadas

- **Linguagem:** Go
- **Interface:** Fyne
- **Banco de Dados:** SQLite
- **Bibliotecas externas:** `gofpdf` para a geração de relatórios em PDF

### Como buildar:
1. Na pasta raiz do projeto usar o comando `make build` para criar o binário do projeto (para Windows SO).
2. Usar o comando `make setup` para criar o instalador do projeto.
3. usar o comando `make release` para buildar e gerar o instalador `SuperMarketSetup.exe`. Obs: é necessário ter o wine e o inno setup instalados

## Instalação e Configuração

O SuperMarket Tracker é distribuído como um instalador Windows criado com Inno Setup, garantindo facilidade de instalação.

### Requisitos:
- Nenhum pré-requisito necessário. O instalador já inclui todas as dependências necessárias.

### Como instalar:
1. Execute o `SuperMarketSetup.exe` presente na pasta `build`.
2. Siga as instruções da instalação.
3. O aplicativo será instalado e estará pronto para uso.

### Como rodar:
- Basta abrir o executável `superMarket.exe` instalado pelo setup.

## Arquitetura do Projeto

O projeto segue a **Arquitetura Hexagonal**, garantindo um código modular e bem estruturado. Abaixo está a organização das pastas e suas responsabilidades:

```
├── Makefile
├── README.md
├── build
│   ├── SuperMarketSetup.exe
│   ├── setup.iss
│   └── superMarket.exe
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── adapters
│   │   ├── database
│   │   │   └── sqlite_repository.go
│   │   └── gui
│   │       ├── fyne_presenter.go
│   │       ├── gui_buttons.go
│   │       ├── insert_window.go
│   │       ├── list_window.go
│   │       └── main_window.go
│   ├── application
│   │   ├── services
│   │   │   ├── gui_service.go
│   │   │   └── user_service.go
│   │   └── usecases
│   │       └── user_repository.go
│   ├── config
│   │   └── colors
│   │       └── colors.go
│   └── domain
│       └── product.go
├── main
├── pkg
├── products.db
├── reports
└── temp
    ├── products.db
    └── report.pdf
```

## Funcionalidades

- **Cadastrar Produtos:** O usuário pode adicionar produtos com nome, valor e data (mês/ano).
- **Listar Produtos:** Exibir todos os produtos cadastrados.
- **Gerar Relatórios em PDF:** Criar um arquivo PDF com os produtos cadastrados e seus valores.

## Contribuição

## Licença
