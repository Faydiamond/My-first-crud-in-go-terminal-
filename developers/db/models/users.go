package models

import (
	"fmt"
	"godevps/db"

	"github.com/TwiN/go-color"
)

type User struct {
	Id_user        int
	Nombres        string
	Apellidos      string
	Age            int
	fk_ciudad      int
	fk_genero      int
	fecha_creacion string
}

type Users []User

/* Insertar usuario */
func newuser(nombres string, apellidos string, age int, numbercity int, numbergender int) *User {
	usercreated := &User{Nombres: nombres, Apellidos: apellidos, Age: age, fk_ciudad: numbercity, fk_genero: numbergender}
	return usercreated
}

func (user *User) Insert() {
	sql := "INSERT INTO users (nombres,apellidos,age,fk_ciudad,fk_genero) values (?,?,?,?,?)"
	db.Exec(sql, user.Nombres, user.Apellidos, user.Age, user.fk_ciudad, user.fk_genero)
}

func CreateUser(nombres string, apellidos string, age int, numbercity int, numbergender int) string {
	userr := newuser(nombres, apellidos, age, numbercity, numbergender)
	userr.Insert()
	res := fmt.Sprintf(" Se creo el usuario %s con exito. ", nombres)
	return res
}

/* Fin Insertar usuario */

/*  Obtener user por nombres y apellidos */

func GetUserName(name string, lastname string) *User {
	userr := newuser("", "", 0, 0, 0)
	sql := "SELECT * FROM developers.users where nombres = ? and apellidos=?"
	row, _ := db.Query(sql, name, lastname)
	for row.Next() {
		row.Scan(&userr.Id_user, &userr.Nombres, &userr.Apellidos, &userr.Age, &userr.fk_ciudad, &userr.fk_genero, &userr.fecha_creacion)
	}
	return userr
}

/* Fin Obtener user por nombres y apellidos */

/*  Update user */
func (us *User) Update() {
	sql := " Update developers.users set nombres =?, apellidos =?, age =? where id_user=? "
	db.Exec(sql, us.Nombres, us.Apellidos, us.Age, us.Id_user)

	res := fmt.Sprintf("El usuario %s se actualizo de manera correcta!", us.Nombres)
	fmt.Println(color.Ize(color.Green, res))
}

/*  Fin Update user */

/* listar usuarios */

func ListUsers() Users {
	users := Users{}
	sql := " select * from users "
	rows, _ := db.Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id_user, &user.Nombres, &user.Apellidos, &user.Age, &user.fk_ciudad, &user.fk_genero, &user.fecha_creacion)
		users = append(users, user)
	}
	return users
}

/* fin listar usuarios */

/* Eliminar usuario */
func (us *User) DeleteUser() {
	sql := "DELETE from users where id_user=? "
	db.Exec(sql, us.Id_user)
	res := fmt.Sprintf("El usuario %s se elimino de manera correcta!", us.Nombres)
	fmt.Println(color.Ize(color.Green, res))
}

/* Fin Eliminar usuario  */
