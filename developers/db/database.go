package db

import (
	"database/sql"
	"fmt"

	// Import the Azure AD driver module (also imports the regular driver package)
	_ "github.com/go-sql-driver/mysql"
)

const url = "Testt:2022*@tcp(localhost:3306)/developers" //user:password

var db *sql.DB

func Connect() {
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)

	}
	//fmt.Println("Conexion Exitosa")
	db = connection
}

//cerrar conexion
func Close() {
	db.Close()
}

//revisar conexion
func Ping() {

	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//verificar si existe tabla
func verifyTable(tableName string) bool {
	sql := fmt.Sprintf("show tables like '%s'", tableName)
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return rows.Next()
}

//crear tablas
func CreateTables(sheme string, table string) {
	if !verifyTable(table) {
		_, err := db.Exec(sheme)
		if err != nil {
			fmt.Println(" Error al momento de crear la tabla: ", err)
		}
	} else {
		fmt.Println("La tabla", table, "ya existe", table)
	}
}

//Eliminar tablas
func DeleteTable(tableName string) bool {
	sql := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
	fmt.Println("Truncar ", sql)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("error al momento de truncar la tabla ", tableName, "    ", err)
	}
	return rows.Next()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println("El erro es : ", err)
	}
	return rows, err
}
