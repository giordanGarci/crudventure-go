# Crudventure-go

Crudventure-go é um projeto que implementa uma api com funcionalidades básicas para gerenciar dados de um Produto. Este projeto foi desenvolvido em Go (Golang) e tem como objetivo servir como base para aprender, implementar ou expandir aplicações que requerem operações de CRUD.

## Recursos

- Implementação de uma api de criar e buscar dados.
- Estrutura modular para facilitar manutenção e expansão.
- Uso de boas práticas de desenvolvimento com Go.
- Integração com bancos de dados relacionais ou não relacionais.
- API RESTful para comunicação com o sistema.

## Requisitos

Certifique-se de ter os seguintes requisitos instalados no seu ambiente:

- [Go](https://go.dev/) (versão 1.18 ou superior)
- [Git](https://git-scm.com/)
- Banco de dados compatível (ex.: PostgreSQL, MySQL, SQLite)

## Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/giordanGarci/crudventure-go.git
   ```

2. Acesse o diretório do projeto:
   ```bash
   cd crudventure-go
   ```

3. Instale as dependências:
   ```bash
   go mod tidy
   ```

4. Configure as variáveis de ambiente para conectar ao banco de dados. Use o arquivo `.env.example` como referência:
   ```bash
   cp .env.example .env
   ```
   Edite o arquivo `.env` com as credenciais do seu banco de dados.

5. Execute a aplicação:
   ```bash
   go run main.go
   ```

## Uso

Uma vez que a aplicação estiver em execução, você pode acessar os endpoints da API para realizar operações CRUD. Por exemplo:

- **GET** `/produtcs`: Listar todos os itens.
- **POST** `/produtcs`: Criar um novo item.
- **GET** `/produtc/{id}`: Atualizar um item existente.

## Contribuição

Fique à vontade para contribuir com melhorias para o projeto. Siga estas etapas para contribuir:

1. Faça um fork do repositório.
2. Crie uma branch para sua contribuição:
   ```bash
   git checkout -b minha-contribuicao
   ```
3. Faça commit das suas alterações:
   ```bash
   git commit -m "Descrição das alterações"
   ```
4. Envie as alterações para o seu fork:
   ```bash
   git push origin minha-contribuicao
   ```
5. Abra um Pull Request no repositório original.

