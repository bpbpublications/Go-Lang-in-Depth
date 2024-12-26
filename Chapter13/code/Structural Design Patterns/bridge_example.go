package main

import "fmt"

type SQLCompiler interface {
	ExecuteSQL(string)
}

type PGSQLCompiler struct {
}

func (pgc *PGSQLCompiler) ExecuteSQL(sql string) {
	fmt.Println("PostGres executing sql- ", sql)
}

type PostGres struct {
	sqlCompiler SQLCompiler
}

func (pg *PostGres) GetDatabaseInfo() {
	fmt.Println("PostGres")
}

func (pg *PostGres) SetSQLCompiler(sql SQLCompiler) {
	fmt.Println("PostGres setting SQLCompiler")
	pg.sqlCompiler = sql
}
func (pg *PostGres) ExecuteSQL(sql string) {
	fmt.Println("PostGres executing sql")
	pg.sqlCompiler.ExecuteSQL(sql)
}

type Database interface {
	GetDatabaseInfo()
	SetSQLCompiler(SQLCompiler)
	ExecuteSQL(string)
}

func main() {

	pg := &PostGres{}

	pgsql := &PGSQLCompiler{}

	pg.SetSQLCompiler(pgsql)

	sql := "Select * FROM Customer"
	pg.ExecuteSQL(sql)
}
