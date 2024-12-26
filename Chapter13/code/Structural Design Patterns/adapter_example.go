package main

import "fmt"

type MySQL struct{}

func (mysql *MySQL) GetDBInfo() {
	fmt.Println("MySQL DB")
}

type MySQLAdapter struct {
	mysqlServer *MySQL
}

func (mysql *MySQLAdapter) GetDatabaseInfo() {
	fmt.Println("MySQL adapter gets DB Info")
	mysql.mysqlServer.GetDBInfo()
}

type PostGres struct {
}

func (pg *PostGres) GetDatabaseInfo() {
	fmt.Println("PostGres")
}

type DBClient struct {
}

func (client *DBClient) GetDatabaseInfo(db Database) {
	fmt.Println("Client  gets Database Info from the specific DB")
	db.GetDatabaseInfo()
}

type Database interface {
	GetDatabaseInfo()
}

func main() {

	client := &DBClient{}
	pg := &PostGres{}

	client.GetDatabaseInfo(pg)

	mysqlServer := &MySQL{}
	mysqlAdapter := &MySQLAdapter{
		mysqlServer: mysqlServer,
	}

	client.GetDatabaseInfo(mysqlAdapter)
}
