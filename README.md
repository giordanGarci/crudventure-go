# Crudventure-go

Crudventure-go is a project that implements api system with basic functionalities to manage data struct Product. This project was developed in Go (Golang) and aims to serve as a foundation for learning, implementing, or expanding applications that require CRUD operations.

## Features

- Modular structure for easy maintenance and expansion.
- Use of best practices in Go development.
- Integration with relational or non-relational databases.
- RESTful API for system communication.

## Requirements

Make sure you have the following prerequisites installed in your environment:

- [Go](https://go.dev/) (version 1.18 or higher)
- [Git](https://git-scm.com/)
- Compatible database (e.g., PostgreSQL, MySQL, SQLite)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/giordanGarci/crudventure-go.git
   ```

2. Navigate to the project directory:
   ```bash
   cd crudventure-go
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Configure the environment variables to connect to the database. Use the `.env.example` file as a reference:
   ```bash
   cp .env.example .env
   ```
   Edit the `.env` file with your database credentials.

5. Run the application:
   ```bash
   go run main.go
   ```

## Usage

Once the application is running, you can access the API endpoints to perform CRUD operations. For example:

- **GET** `/products`: List all items.
- **POST** `/product`: Create a new item.
- **GET** `/product/{id}`: Get an existing product.

## Contribution

Feel free to contribute improvements to the project. Follow these steps to contribute:

1. Fork the repository.
2. Create a branch for your contribution:
   ```bash
   git checkout -b my-contribution
   ```
3. Commit your changes:
   ```bash
   git commit -m "Description of changes"
   ```
4. Push the changes to your fork:
   ```bash
   git push origin my-contribution
   ```
5. Open a Pull Request on the original repository.

