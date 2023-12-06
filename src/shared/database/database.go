package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
)

type Database struct {
	*gorm.DB
}

func New(config ConfigDatabase) *Database {
	fmt.Println("Try NewDatabase ...")

	db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/alchemy?multiStatements=true&charset=utf8&parseTime=True")

	if err != nil {
		fmt.Println("failed to connect database")
		panic(err)
	}

	db.SingularTable(false)

	err = db.DB().Ping()
	if err != nil {
		fmt.Println("failed to connect database")
		panic(err)
	}

	db.DB().SetMaxIdleConns(config.MinIdleConnections)
	db.DB().SetMaxOpenConns(config.MaxOpenConnections)
	db.DB().SetConnMaxLifetime(config.ConnMaxLifetime)

	db.LogMode(config.DebugMode)

	s := &Database{
		db,
	}

	if err = s.MigrateUP("./db/migration"); err != nil {
		fmt.Println("failed to migrate database")
		panic(err)
	}

	return s
}

func (s *Database) MigrateUP(path string) error {
	driver, _ := mysql.WithInstance(s.DB.DB(), &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migration",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}
	_ = m.Steps(2)

	return nil
}
