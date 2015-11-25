//usr/bin/env go run $0 "$@"; exit
package tasks

import (
	"fmt"

	"github.com/IIC2173-2015-2-Grupo2/news-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "user=gslopez dbname=newsapi sslmode=disable")
	// db, err := gorm.Open("foundation", "dbname=gorm") // FoundationDB.
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open("sqlite3", "/tmp/gorm.db")

	// You can also use an existing database connection handle
	// dbSql, _ := sql.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	// db, _ := gorm.Open("postgres", dbSql)

	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Printf("Se debe crear la base de datos 'newsapi'")

	}

	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.CreateTable(&models.User{})
}
