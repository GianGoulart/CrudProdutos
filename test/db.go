package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetDB retorna uma instancia do db mocado
func GetDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.MonitorPingsOption(true))

	gdb, err := gorm.Open(mysql.New(
		mysql.Config{Conn: db},
	))

	if err != nil {
		logrus.Error(err)
	}

	return gdb, mock
}

// NewRows retorna um modelo para adicionar rows
func NewRows(columns ...string) *sqlmock.Rows {
	return sqlmock.NewRows(columns)
}
