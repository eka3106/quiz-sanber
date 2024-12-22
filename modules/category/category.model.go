package category

import "database/sql"

type Category struct {
	Id          int
	Name        string
	Created_at  string
	Created_by  string
	Modified_at string
	Modified_by string
}

func (c *Category) Scan(rows *sql.Rows) error {
	return rows.Scan(&c.Id, &c.Name, &c.Created_at, &c.Created_by, &c.Modified_at, &c.Modified_by)
}
