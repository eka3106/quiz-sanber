package category

import (
	"errors"
	"quiz/databases"
	books "quiz/modules/book"
)

func GetAll() (result []Category, status int, err error) {
	getAllQuery := `SELECT * FROM categories`
	rows, err := databases.DB.Query(getAllQuery)

	if err != nil {
		return nil, 500, err
	}

	defer rows.Close()
	for rows.Next() {
		var categories Category
		err := rows.Scan(&categories.Id, &categories.Name, &categories.Created_at, &categories.Created_by, &categories.Modified_at, &categories.Modified_by)
		if err != nil {
			return nil, 500, err
		}

		result = append(result, categories)
	}

	return result, 200, nil
}

func GetOne(id int) (result Category, status int, err error) {
	getOneQuery := `SELECT * FROM categories WHERE id=$1`
	err = databases.DB.QueryRow(getOneQuery, id).Scan(&result.Id, &result.Name, &result.Created_at, &result.Created_by, &result.Modified_at, &result.Modified_by)

	if err != nil {
		return result, 404, errors.New("data category not found")
	}

	return result, 200, nil
}

func Create(category Category) (status int, err error) {
	createQuery := `INSERT INTO categories (name,  created_by, modified_by) VALUES ($1, $2, $3 )`
	_, err = databases.DB.Exec(createQuery, category.Name, category.Created_by, category.Modified_by)

	if err != nil {
		return 500, err
	}

	return 201, nil
}

func Update(category Category) (status int, err error) {
	_, status, err = GetOne(category.Id)
	if err != nil {
		return status, err
	}

	updateQuery := `UPDATE categories SET name=$1,   modified_at=$2, modified_by=$3 WHERE id=$4`
	_, err = databases.DB.Exec(updateQuery, category.Name, category.Modified_at, category.Modified_by, category.Id)

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func Delete(id int) (status int, err error) {
	_, status, err = GetOne(id)
	if err != nil {
		return status, err
	}
	deleteQuery := `DELETE FROM Categories WHERE id=$1`
	_, err = databases.DB.Exec(deleteQuery, id)

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func GetBooks(id int) (result []books.Books, status int, err error) {
	getBooksQuery := `SELECT * FROM books WHERE category_id=$1`
	rows, err := databases.DB.Query(getBooksQuery, id)

	if err != nil {
		return nil, 500, err
	}

	defer rows.Close()
	for rows.Next() {
		var books books.Books
		err := rows.Scan(&books.Id, &books.Title, &books.Description, &books.Image_url, &books.Release_year, &books.Price, &books.Total_page, &books.Thickness, &books.Category_id, &books.Created_at, &books.Created_by, &books.Modified_at, &books.Modified_by)
		if err != nil {
			return nil, 500, err
		}

		result = append(result, books)
	}

	if len(result) == 0 {
		return nil, 404, errors.New("buku tidak ditemukan")
	}

	return result, 200, nil
}
