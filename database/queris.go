package database

import (
	"FishRu/types"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func CreateProductTable(conn *pgx.Conn) error {
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

func CreateOwnerTable(conn *pgx.Conn) error {
	query := `CREATE TABLE IF NOT EXISTS owner (id serial primary key, login varchar(64) NOT NULL, password varchar(64) NOT NULL);`
	_, err := conn.Exec(context.Background(), query)
	return err
}

func VerifyUser(conn *pgx.Conn, user *types.User) (string, error) {
	query := `SELECT login, password FROM owner WHERE login = $1 AND password = $2;`
	msg, err := conn.Exec(context.Background(), query, &user.Login, &user.Password)
	return msg.String(), err
}

func SelectAll(conn *pgx.Conn) ([]types.ProductCard, error) {
	query := `SELECT * FROM product;`
	product := types.ProductCard{}
	var productSlice []types.ProductCard
	raws, err := conn.Query(context.Background(), query)
	for raws.Next() {
		if err := raws.Scan(&product.Id, &product.Price, &product.Name, &product.Description, &product.Photos, &product.Available); err != nil {
			log.Fatalf("Can't scan data: %s", err)
		}
		productSlice = append(productSlice, product)
	}
	return productSlice, err
}

func InsertProduct(conn *pgx.Conn, product *types.ProductCard) (bool, error) {
	var ins bool
	query := `INSERT INTO product(price, name, description, photos, available)
		VALUES ($1, $2, $3, $4, true);`
	msg, err := conn.Exec(context.Background(), query, int(product.Price), product.Name, product.Description, product.Photos)
	if msg.String() == "INSERT 0 1" {
		ins = true
	}
	return ins, err
}

func DeleteProduct(conn *pgx.Conn, id int) (bool, error) {
	var del bool
	query := `DELETE FROM product WHERE id = $1;`
	msg, err := conn.Exec(context.Background(), query, id)
	if msg.String() == "DELETE 1" {
		del = true
	}
	return del, err
}

func ModifyProduct(conn *pgx.Conn, product *types.ProductCard, id int) (bool, error) {
	upd0, upd1 := false, true
	if product.Name != "" {
		query := `UPDATE product SET name = $1 WHERE id = $2;`
		_, err := conn.Exec(context.Background(), query, product.Name, id)
		if err != nil {
			return upd0, err
		}
	}
	if product.Price != 0 {
		query := `UPDATE product SET price = $1 WHERE id = $2;`
		_, err := conn.Exec(context.Background(), query, product.Price, id)
		if err != nil {
			return upd0, err
		}
	}
	if len(product.Photos) > 0 {
		query := `UPDATE product SET photos = $1 WHERE id = $2;`
		_, err := conn.Exec(context.Background(), query, product.Photos, id)
		if err != nil {
			return upd0, err
		}
	}
	if product.Description != "" {
		query := `UPDATE product SET description = $1 WHERE id = $2;`
		_, err := conn.Exec(context.Background(), query, product.Description, id)
		if err != nil {
			return upd0, err
		}
	}

	return upd1, nil
}

func SelectByID(conn *pgx.Conn, id int) (types.ProductCard, error) {
	query := `SELECT * FROM product WHERE id = $1;`
	raws, err := conn.Query(context.Background(), query, id)
	product := types.ProductCard{}
	for raws.Next() {
		if err := raws.Scan(&product.Id, &product.Price, &product.Name, &product.Description, &product.Photos, &product.Available); err != nil {
			log.Fatalf("Can't scan data: %s", err)
		}
	}
	return product, err
}
