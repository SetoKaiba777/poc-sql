package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type SqlDatabase struct {
	db *sql.DB
}

func NewSqlDatabase(user, password, url, dbName string) *SqlDatabase {
	cfg := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 url,
		DBName:               dbName,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	return &SqlDatabase{db: db}
}

func (s *SqlDatabase) SelectQuery(query string, params ...interface{}) ([]map[string]interface{}, error) {
	var items []map[string]interface{}
	rows, err := s.db.Query("SELECT "+query, params...)

	if err != nil {
		return nil, err
	}

	col, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		values := make([]interface{}, len(col))

		for i := range values {
			var v interface{}
			values[i] = &v
		}

		if rows.Scan(values...) != nil {
			return nil, fmt.Errorf("scan error")
		}
		item := make(map[string]interface{})
		for i, col := range col {
			val := *(values[i].(*interface{}))

			// Tratamento de valores null
			if b, ok := val.([]byte); ok {
				item[col] = string(b)
			} else {
				item[col] = val
			}
		}
		items = append(items, item)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return items, nil
}
