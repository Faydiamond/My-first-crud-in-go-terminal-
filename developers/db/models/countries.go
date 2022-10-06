package models

import (
	"godevps/db"
)

type Country struct {
	Id_pais        int
	Pais           string
	Fecha_creacion string
}

type Countries []Country

func ListCountries() Countries {
	countries := Countries{}
	sql := "SELECT id_pais,pais,fecha_creacion FROM countries"
	rows, _ := db.Query(sql)

	for rows.Next() {
		country := Country{}
		rows.Scan(&country.Id_pais, &country.Pais, &country.Fecha_creacion)
		countries = append(countries, country)
	}
	return countries

}

//Listar Ciudades
func ListConts() Countries {
	sql := "SELECT id_pais,pais,fecha_creacion FROM countries"
	citys := Countries{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		city := Country{}
		rows.Scan(&city.Id_pais, &city.Pais, &city.Fecha_creacion)
		citys = append(citys, city)
	}
	return citys
}
