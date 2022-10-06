package models

import (
	"fmt"
	"godevps/db"

	"github.com/TwiN/go-color"
)

type City struct {
	Id_Ciudad      int
	Ciudad         string
	Fk_pais        int
	Fecha_creacion string
}

type CytiCountry struct {
	Id_pais   int
	Pais      string
	Id_ciudad int
	Ciudad    string
}

type Citys []City
type CytisCountrys []CytiCountry

//Listar Ciudades
func ListCities() CytisCountrys {
	sql := "SELECT Co.id_pais,Co.pais,C.id_ciudad ,C.Ciudad FROM cities C INNER JOIN countries Co on C.fk_pais = Co.id_pais"
	cytisCountrys := CytisCountrys{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		cytiCountry := CytiCountry{}
		rows.Scan(&cytiCountry.Id_pais, &cytiCountry.Pais, &cytiCountry.Id_ciudad, &cytiCountry.Ciudad)
		cytisCountrys = append(cytisCountrys, cytiCountry)
	}
	return cytisCountrys
}

/* insertar city */
func newcity(fPais int, cityIn string) *City {
	citiCreated := &City{Ciudad: cityIn, Fk_pais: fPais}
	return citiCreated
}

func (citii *City) InsertCity() {
	sql := fmt.Sprintf("insert cities set Ciudad='%s', Fk_pais='%d'", citii.Ciudad, citii.Fk_pais)
	db.Exec(sql)

}

func CreateCity(fPais int, cityIn string) string {
	cityy := newcity(fPais, cityIn)
	cityy.InsertCity()
	res := fmt.Sprintf(" Se creo la ciudad '%s' con exito. ", cityIn)
	return res
}

/* fin insertar city */

func ExistCity(name string) bool {
	sql := "SELECT * FROM cities WHERE ciudad=?"
	row, err := db.Query(sql, name)
	if err != nil {
		panic(err)
	}
	return row.Next()
}

func GetCityName(name string) *City {
	divi := newcity(0, "")
	sql := "select id_ciudad,ciudad ,fk_pais,fecha_creacion from cities where  ciudad=?"
	//reee := fmt.Sprintf("select id_ciudad,ciudad ,fk_pais,fecha_creacion from cities where  ciudad='%s'", name)
	//fmt.Println(reee)
	row, _ := db.Query(sql, name)

	for row.Next() {
		row.Scan(&divi.Id_Ciudad, &divi.Ciudad, &divi.Fk_pais, &divi.Fecha_creacion)
	}
	//fmt.Println("id del pais: ", divi.Fk_pais ,"Pais : ", ,divi.Id_Ciudad)
	return divi
}

func (ciu *City) Update() {
	sql := "update  cities set  id_ciudad =?,ciudad =?,fk_pais=?,fecha_creacion=?  where id_ciudad=?"
	db.Exec(sql, ciu.Id_Ciudad, ciu.Ciudad, ciu.Fk_pais, ciu.Fecha_creacion, ciu.Id_Ciudad)
	//division.id, _ = result.LastInsertId()
	res := fmt.Sprintf("La ciudad %s se actualizo de manera correcta!", ciu.Ciudad)
	fmt.Println(color.Ize(color.Green, res))
}

func (ciu *City) Delete() {
	sql := "DELETE from cities where id_ciudad=?"
	db.Exec(sql, ciu.Id_Ciudad)
	res := fmt.Sprintf("La ciudad %s se elimino de manera correcta!", ciu.Ciudad)
	fmt.Println(color.Ize(color.Green, res))
}
