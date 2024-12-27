package repository

import (
	"database/sql"
	"fmt"

	"github.com/giordanGarci/crudventure-go/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			// Nenhum resultado encontrado
			return []model.Product{}, nil
		}
		// Outro erro ocorreu
		fmt.Println(err)
		return []model.Product{}, err
	}
	var productList []model.Product
	var productObject model.Product

	for rows.Next() {
		err := rows.Scan(&productObject.ID, &productObject.Name, &productObject.Price)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObject)
	}
	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int64, error) {
	query := "INSERT INTO product (product_name, price) VALUES (?, ?)"

	// Executa a query e obtém o resultado
	result, err := pr.connection.Exec(query, product.Name, product.Price)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Obtém o ID do último registro inserido
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func (pr *ProductRepository) GetProductById(id int64) (*model.Product, error) {
	query := "SELECT id_product, product_name, price FROM product WHERE id_product = ?"

	var product model.Product
	err := pr.connection.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			// Nenhum resultado encontrado
			return nil, nil
		}
		// Outro erro ocorreu
		fmt.Println(err)
		return nil, err
	}

	return &product, nil
}
