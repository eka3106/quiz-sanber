package book

import (
	"errors"
	"os"
	"quiz/databases"
)

func GetAll() (result []Books, status int, err error) {
	getAllQuery := `SELECT * FROM books`
	rows, err := databases.DB.Query(getAllQuery)

	if err != nil {
		return nil, 500, err
	}

	defer rows.Close()
	for rows.Next() {
		var books Books
		err := rows.Scan(&books.Id, &books.Title, &books.Description, &books.Image_url, &books.Release_year, &books.Price, &books.Total_page, &books.Thickness, &books.Category_id, &books.Created_at, &books.Created_by, &books.Modified_at, &books.Modified_by)
		if err != nil {
			return nil, 500, err
		}

		result = append(result, books)
	}

	return result, 200, nil
}

func GetOne(id int) (result Books, status int, err error) {
	getOneQuery := `SELECT * FROM books WHERE id=$1`
	err = databases.DB.QueryRow(getOneQuery, id).Scan(&result.Id, &result.Title, &result.Description, &result.Image_url, &result.Release_year, &result.Price, &result.Total_page, &result.Thickness, &result.Category_id, &result.Created_at, &result.Created_by, &result.Modified_at, &result.Modified_by)

	if err != nil {
		return result, 404, errors.New("data book not found")
	}

	return result, 200, nil
}

func Create(books Books) (status int, err error) {
	createQuery := `INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id,  created_by, modified_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err = databases.DB.Exec(createQuery, books.Title, books.Description, books.Image_url, books.Release_year, books.Price, books.Total_page, books.Thickness, books.Category_id, books.Created_by, books.Modified_by)

	if err != nil {
		return 500, err
	}

	return 201, nil
}

func Update(books Books) (status int, err error) {
	result, status, err := GetOne(books.Id)
	if err != nil {
		return status, err
	}
	if result.Image_url != books.Image_url {
		err := os.Remove(result.Image_url)
		if err != nil {
			return 500, err
		}
	}
	println(books.Image_url)
	updateQuery := `UPDATE books SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, category_id=$8, modified_at=$9, modified_by=$10 WHERE id=$11`
	_, err = databases.DB.Exec(updateQuery, books.Title, books.Description, books.Image_url, books.Release_year, books.Price, books.Total_page, books.Thickness, books.Category_id, books.Modified_at, books.Modified_by, books.Id)

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func Delete(id int) (status int, err error) {
	result, status, err := GetOne(id)
	if err != nil {
		return status, err
	}
	deleteQuery := `DELETE FROM books WHERE id=$1`
	_, err = databases.DB.Exec(deleteQuery, id)
	if err != nil {
		return 500, err
	}
	os.Remove(result.Image_url)

	return 200, nil
}
