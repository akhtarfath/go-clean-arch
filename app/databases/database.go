package databases

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/url"
)

type DatabaseMySQL Database
type DatabasePgSQL Database
type DatabaseMSSQL Database

type Database struct {
	Name string
}

type DatabaseRepository struct {
	DatabaseMySQL *DatabaseMySQL
	DatabasePgSQL *DatabasePgSQL
	DatabaseMSSQL *DatabaseMSSQL
}

func NewDatabaseMySQL() *DatabaseMySQL {
	return (*DatabaseMySQL)(&Database{
		Name: "mysql",
	})
}

func (d *DatabaseMySQL) GetDatabaseConnection() *sql.DB {
	dbHost, dbPort, dbUser, dbPass, dbName := dbEnv(d.Name)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal("Error: Driver MySQL not found")
	}

	return dbConn
}

func NewDatabasePgSQL() *DatabasePgSQL {
	return (*DatabasePgSQL)(&Database{
		Name: "pgsql",
	})
}

func (d *DatabasePgSQL) GetDatabaseConnection() *sql.DB {
	dbHost, dbPort, dbUser, dbPass, dbName := dbEnv(d.Name)
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	dbConn, err := sql.Open(`postgres`, connection)
	if err != nil {
		log.Fatal("Error: Driver PostgresSQL not found")
	}

	return dbConn
}

func (d *DatabaseMSSQL) GetDatabaseConnection() *sql.DB {
	dbHost, dbPort, dbUser, dbPass, dbName := dbEnv(d.Name)
	connection := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", dbHost, dbUser, dbPass, dbPort, dbName)
	dbConn, err := sql.Open(`mssql`, connection)
	if err != nil {
		log.Fatal("Error: Driver MSSQL not found")
	}

	return dbConn
}

func NewDatabaseMSSQL() *DatabaseMSSQL {
	return (*DatabaseMSSQL)(&Database{
		Name: "mssql",
	})
}

func NewDatabaseRepository(MySQL *DatabaseMySQL, PgSQL *DatabasePgSQL, MSSQL *DatabaseMSSQL) *DatabaseRepository {
	return &DatabaseRepository{DatabaseMySQL: MySQL, DatabasePgSQL: PgSQL, DatabaseMSSQL: MSSQL}
}

func dbEnv(driverName string) (host, port, user, pass, name string) {
	dbHost := viper.GetString(`database.driver.` + driverName + `.host`)
	dbPort := viper.GetString(`database.driver.` + driverName + `.port`)
	dbUser := viper.GetString(`database.driver.` + driverName + `.user`)
	dbPass := viper.GetString(`database.driver.` + driverName + `.pass`)
	dbName := viper.GetString(`database.driver.` + driverName + `.name`)

	return dbHost, dbPort, dbUser, dbPass, dbName
}
