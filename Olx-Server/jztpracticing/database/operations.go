package database

import (

	"gorm.io/gorm"
)

func Query(table Group_tables, query string, db *gorm.DB, params ...interface{}) {
	db.Raw(query, params...).Scan(&table)
}
