# Product API — Clean Architecture com Golang

Uma API REST de Produtos desenvolvida em **Golang**, criada como **projeto de estudo aprofundado em Clean Architecture**, com foco em **design orientado ao domínio**, **baixo acoplamento** e **alta testabilidade**.

O objetivo deste repositório é servir como **referência prática** para a construção de aplicações backend em Go, aplicando princípios arquiteturais sólidos e boas práticas de Clean Code.

## Objetivo

- Estudar e aplicar **Clean Architecture** em Go
- Construir uma API simples, porém arquiteturalmente correta
- Demonstrar **separação clara de responsabilidades**
- Facilitar manutenção, testes e evolução do sistema
- Servir como base para projetos maiores

## Princípios Arquiteturais

Este projeto segue os seguintes princípios:

- **Domínio no centro da aplicação**
- **Casos de uso orquestram regras de negócio**
- **Interfaces definidas pelo domínio**
- **Infraestrutura como detalhe**
- **Dependências sempre apontam para dentro**
- **Frameworks não controlam o fluxo da aplicação**

## Estrutura do Projeto

```text
.
├── cmd/
│   └── api/              # Ponto de entrada da aplicação
│       └── main.go
├── internal/
│   ├── domain/           # Regras de negócio (núcleo)
│   │   └── product/
│   ├── usecase/          # Casos de uso da aplicação
│   │   └── product/
│   ├── interfaces/       # Adaptadores de entrada (HTTP)
│   │   └── http/
│   └── infra/            # Implementações externas (DB, etc)
│       └── repository/
└── go.mod
```

> A estrutura prioriza **organização por responsabilidade**, não por tipo técnico.

## Fluxo da Aplicação

```text
HTTP Request
   ↓
Handler (Interface)
   ↓
Use Case
   ↓
Domain Entity
   ↓
Repository (Interface)
   ↓
Infrastructure (Implementação)
```

## Testes

O design do projeto facilita:

- Testes unitários de **entidades**
- Testes de **casos de uso** sem dependência de banco ou HTTP
- Uso de repositórios em memória ou mocks simples

> A testabilidade é consequência direta das decisões arquiteturais.

## Tecnologias

- Golang
- net/http
- Clean Architecture
- UUID
- Repositório em memória (inicialmente)

## 📈 Evolução planejada

Este projeto será evoluído incrementalmente para incluir:

- [ ] Testes unitários
- [ ] DTOs e mapeamento de respostas
- [ ] Persistência com banco de dados
- [ ] Versionamento de API
- [ ] Middlewares (log, recovery)
- [ ] Observabilidade
- [ ] Autenticação (conceitual)

## 📌 Observações

Este projeto tem caráter **educacional**, com foco em **clareza, legibilidade e decisões explícitas**, evitando abstrações desnecessárias ou complexidade prematura.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests com melhorias, correções ou sugestões.

## Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## Contato

Para dúvidas ou sugestões, entre em contato:

- Nome: Vitor Valente
- LinkedIn: https://www.linkedin.com/in/ovitorvalente/
