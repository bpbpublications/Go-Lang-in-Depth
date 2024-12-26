package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DBUtil struct {
}

func NewDBUtil() *DBUtil {
	return &DBUtil{}
}

func (dbUtil *DBUtil) readDB(path string) {

	database, error := sql.Open(path, ":memory:")

	if error != nil {
		log.Fatal(error)
	}

	defer database.Close()

	var version_number string
	error = database.QueryRow("SELECT SQLITE_VERSION()").Scan(&version_number)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(version_number)
}

func main() {

	var dbUtil *DBUtil

	dbUtil = NewDBUtil()

	dbUtil.readDB("sqlite3")

}
