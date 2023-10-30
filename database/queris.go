package database

import (
	"FishRu/functions"
	"FishRu/types"
	"context"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"os"
	"strings"
)

func CreateProductTable(conn *pgx.Conn) error {
	query := `CREATE TABLE IF NOT EXISTS product (
				id serial primary key, 
				price text NOT NULL, 
				name varchar(64) NOT NULL, 
				description text NOT NULL, 
				photos text[] NOT NULL,
    			category varchar(64) NOT NULL,
				slug_name varchar(64) NOT NULL,
    			slug_category varchar(64) NOT NULL		
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
	rows, err := conn.Query(context.Background(), query)
	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Price, &product.Name, &product.Description, &product.Photos, &product.Category, &product.SlugName, &product.SlugCategory); err != nil {
			log.Fatalf("Can't scan data: %s", err)
		}
		product.Name = functions.ToUpperFirstSymbol(product.Name)
		product.Category = functions.ToUpperFirstSymbol(product.Category)
		productSlice = append(productSlice, product)
	}
	return productSlice, err
}

func InsertProduct(conn *pgx.Conn, product *types.ProductCard) (bool, error) {
	var ins bool
	query := `INSERT INTO product(price, name, description, photos, category, slug_name, slug_category)
		VALUES ($1, $2, $3, $4, $5, $6, $7);`
	msg, err := conn.Exec(context.Background(), query, product.Price, strings.ToLower(product.Name), product.Description, product.Photos, strings.ToLower(product.Category), slug.Make(product.Name), slug.Make(product.Category))
	if msg.String() == "INSERT 0 1" {
		ins = true
	}
	return ins, err
}

func DeleteProduct(conn *pgx.Conn, id int) (bool, error) {
	var del bool
	queryPhotosAddress := `SELECT photos FROM product WHERE id = $1;`
	var photosArray pgtype.Array[string]
	if err := conn.QueryRow(context.Background(), queryPhotosAddress, id).Scan(&photosArray); err != nil {
		return del, err
	}
	photos := photosArray.Elements

	for _, photo := range photos {
		if err := os.Remove(photo); err != nil {
			return false, err
		}
	}

	query := `DELETE FROM product WHERE id = $1;`
	msg, err := conn.Exec(context.Background(), query, id)
	if msg.String() == "DELETE 1" {
		del = true
	}
	return del, err
}

func ModifyProduct(conn *pgx.Conn, product *types.ProductCard, id int) (bool, error) {
	if product.Name != "" {
		query := `UPDATE product SET name = $1 WHERE id = $2;`
		_, err := conn.Exec(context.Background(), query, product.Name, id)
		if err != nil {
			return false, err
		}
	}
	if product.Price != "" {
		query := `UPDATE product SET price = $1 WHERE id = $2;`
		_, err := conn.Exec(context.Background(), query, product.Price, id)
		if err != nil {
			return false, err
		}
	}

	// todo
	//if len(product.Photos) > 0 {
	//	query := `UPDATE product SET photos = $1 WHERE id = $2;`
	//	_, err := conn.Exec(context.Background(), query, product.Photos, id)
	//	if err != nil {
	//		return upd0, err
	//	}
	//}

	if product.Description != "" {
		query := `UPDATE product SET description = $1 WHERE id = $2;`
		_, err := conn.Exec(context.Background(), query, product.Description, id)
		if err != nil {
			return false, err
		}
	}
	if product.Category != "" {
		query := `UPDATE product SET category = $1 WHERE id = $2;`
		_, err := conn.Exec(context.Background(), query, product.Category, id)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func SelectByID(conn *pgx.Conn, id int) (types.ProductCard, error) {
	query := `SELECT * FROM product WHERE id = $1;`
	product := types.ProductCard{}
	err := conn.QueryRow(context.Background(), query, id).Scan(&product.Id, &product.Price, &product.Name, &product.Description, &product.Photos, &product.Category, &product.SlugName, &product.SlugCategory)
	product.Name = functions.ToUpperFirstSymbol(product.Name)
	product.Category = functions.ToUpperFirstSymbol(product.Category)
	return product, err
}

func SelectBySlugName(conn *pgx.Conn, slug string) (types.ProductCard, error) {
	query := `SELECT * FROM product WHERE slug_name = $1;`
	product := types.ProductCard{}
	err := conn.QueryRow(context.Background(), query, slug).Scan(&product.Id, &product.Price, &product.Name, &product.Description, &product.Photos, &product.Category, &product.SlugName, &product.SlugCategory)
	product.Name = functions.ToUpperFirstSymbol(product.Name)
	product.Category = functions.ToUpperFirstSymbol(product.Category)
	return product, err
}

func SelectBySlugCategory(conn *pgx.Conn, slug string) ([]types.ProductCard, error) {
	query := `SELECT DISTINCT ON (slug_category) * 
			  FROM product 
			  WHERE slug_category = $1
			  ORDER BY slug_category;`
	product := types.ProductCard{}
	var sliceProduct []types.ProductCard
	raws, err := conn.Query(context.Background(), query, slug)
	for raws.Next() {
		if err := raws.Scan(&product.Id, &product.Price, &product.Name, &product.Description, &product.Photos, &product.Category, &product.SlugName, &product.SlugCategory); err != nil {
			return sliceProduct, err
		}
		product.Name = functions.ToUpperFirstSymbol(product.Name)
		product.Category = functions.ToUpperFirstSymbol(product.Category)
		sliceProduct = append(sliceProduct, product)
	}
	return sliceProduct, err
}

func GetCategory(conn *pgx.Conn) ([]string, error) {
	query := `SELECT category FROM product GROUP BY category;`
	var cat []string
	category := struct {
		c string
	}{}
	rows, err := conn.Query(context.Background(), query)
	for rows.Next() {
		if err := rows.Scan(&category.c); err != nil {
			return cat, err
		}
		category.c = functions.ToUpperFirstSymbol(category.c)
		cat = append(cat, category.c)
	}
	return cat, err
}
