package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func CreateTable(conn *pgx.Conn) error {
	query := `CREATE TABLE IF NOT EXISTS product (
				id serial primary key, 
				price int NOT NULL, 
				name varchar(50) NOT NULL, 
				description text NOT NULL, 
				photos text[] NOT NULL, 
				available bool NOT NULL
    );`
	_, err := conn.Exec(context.Background(), query)
	return err
}

func DropTable(conn *pgx.Conn) error {
	query := `DROP TABLE product;`
	_, err := conn.Exec(context.Background(), query)
	return err
}

func SelectAll(conn *pgx.Conn) ([]ProductCard, error) {
	query := `SELECT * FROM product;`
	product := ProductCard{}
	var productSlice []ProductCard
	raws, err := conn.Query(context.Background(), query)
	for raws.Next() {
		if err := raws.Scan(&product.Id, &product.Price, &product.Name, &product.Description, &product.Photos, &product.Available); err != nil {
			log.Fatalf("Can't scan data: %s", err)
		}
		productSlice = append(productSlice, product)
	}
	return productSlice, err
}

func InsertProduct(conn *pgx.Conn, product *ProductCard) (string, error) {
	query := `INSERT INTO product(price, name, description, photos, available)
		VALUES ($1, $2, $3, $4, true);`
	msg, err := conn.Exec(context.Background(), query, int(product.Price), product.Name, product.Description, product.Photos)
	return msg.String(), err
}

func DeleteProduct(conn *pgx.Conn, product *ProductCard) (string, error) {
	query := `DELETE FROM product WHERE id = $1;`
	msg, err := conn.Exec(context.Background(), query, product.Id)
	return msg.String(), err
}

func ModifyProduct(conn *pgx.Conn, product *ProductCard) (string, error) {
	upd0, upd1 := "UPDATE 0", "UPDATE 1"
	if product.Name != "" {
		query := `UPDATE product SET name = $1 WHERE id = $2;`
		_, err := conn.Exec(context.Background(), query, product.Name, product.Id)
		if err != nil {
			return upd0, err
		}
	}
	if product.Price != 0 {
		query := `UPDATE product SET price = $1 WHERE id = $2;`
		_, err := conn.Exec(context.Background(), query, product.Price, product.Id)
		if err != nil {
			return upd0, err
		}
	}
	if len(product.Photos) > 0 {
		query := `UPDATE product SET photos = $1 WHERE id = $2;`
		_, err := conn.Exec(context.Background(), query, product.Photos, product.Id)
		if err != nil {
			return upd0, err
		}
	}
	if product.Description != "" {
		query := `UPDATE product SET description = $1 WHERE id = $2;`
		_, err := conn.Exec(context.Background(), query, product.Description, product.Id)
		if err != nil {
			return upd0, err
		}
	}

	return upd1, nil
}
